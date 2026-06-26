package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var tokenRenewMutex sync.Mutex

type TokenCacheEntry struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	CreatedAt    int64  `json:"createdAt"`
	ExpiresAt    int64  `json:"expiresAt"`
}

type EncryptedTokenCache struct {
	Version    int    `json:"version"`
	Nonce      string `json:"nonce"`
	Ciphertext string `json:"ciphertext"`
}

type AccessTokenCreateReq struct {
	TokenName               string `json:"tokenName"`
	RenewableUntilTimestamp int64  `json:"renewableUntilTimestamp"`
	TokenExpiryTimestamp    int64  `json:"tokenExpiryTimestamp"`
	EnableIPAllowList       bool   `json:"enableIPAllowListFeature"`
}

type AccessTokenError struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    int    `json:"errorCode"`
}

type AccessTokenInfo struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AccessTokenCreateResp struct {
	Error     AccessTokenError `json:"error"`
	TokenInfo AccessTokenInfo  `json:"tokenInfo"`
}

type AccessTokenRenewReq struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AccessTokenRenewResp struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// getCacheFilePath returns the path to the token cache file: ~/.commvault/tokens.json
func getCacheFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	cacheDir := filepath.Join(homeDir, ".commvault")
	return filepath.Join(cacheDir, "tokens.json"), nil
}

// ensureCacheDir creates ~/.commvault directory if it doesn't exist
func ensureCacheDir() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}
	cacheDir := filepath.Join(homeDir, ".commvault")
	if err := os.MkdirAll(cacheDir, 0700); err != nil {
		return fmt.Errorf("failed to create cache directory %s: %w", cacheDir, err)
	}
	return nil
}

// getCachedTokens reads and validates the token cache file.
// Returns nil if cache doesn't exist, is invalid JSON, or tokens are expired.
func getCachedTokens() *TokenCacheEntry {
	cacheFile, err := getCacheFilePath()
	if err != nil {
		return nil
	}

	data, err := os.ReadFile(cacheFile)
	if err != nil {
		// File doesn't exist or can't be read - cache miss
		return nil
	}

	cache, decrypted := decryptCachedTokens(data)
	if !decrypted {
		// Backward compatibility: accept legacy plaintext cache and migrate it to encrypted format.
		if err := json.Unmarshal(data, &cache); err != nil {
			return nil
		}
		_ = saveCachedTokenEntry(&cache)
		if !isCacheValid(&cache) {
			return nil
		}
		return &cache
	}

	if !isCacheValid(&cache) {
		return nil
	}

	return &cache
}

// saveCachedTokens writes tokens to ~/.commvault/tokens.json
// Token expiry is set to 2 hours from now (Commvault's access token lifetime)
func saveCachedTokens(accessToken, refreshToken string) error {
	if err := ensureCacheDir(); err != nil {
		// Non-blocking: if cache fails, just continue without caching
		return nil
	}

	cache := TokenCacheEntry{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		CreatedAt:    time.Now().Unix(),
		ExpiresAt:    time.Now().Add(2 * time.Hour).Unix(), // Access tokens valid ~2 hours
	}

	return saveCachedTokenEntry(&cache)
}

func saveCachedTokenEntry(cache *TokenCacheEntry) error {
	cacheFile, err := getCacheFilePath()
	if err != nil {
		return nil
	}

	data, err := encryptCachedTokens(cache)
	if err != nil {
		// Non-blocking: if encryption fails, continue without cache.
		return nil
	}

	// Write with restricted permissions (user only)
	if err := os.WriteFile(cacheFile, data, 0600); err != nil {
		// Non-blocking: cache write failure doesn't break provider
		return nil
	}

	return nil
}

func isCacheValid(cache *TokenCacheEntry) bool {
	if cache == nil {
		return false
	}

	now := time.Now().Unix()
	if cache.ExpiresAt > 0 && now > cache.ExpiresAt-300 {
		return false
	}

	if cache.AccessToken == "" || cache.RefreshToken == "" {
		return false
	}

	return true
}

func cacheEncryptionKey() ([]byte, error) {
	bootstrapToken := normalizeAuthToken(os.Getenv("CV_BOOTSTRAP_TOKEN"))
	if bootstrapToken == "" {
		bootstrapToken = normalizeAuthToken(os.Getenv("CV_API_TOKEN"))
	}
	if bootstrapToken == "" {
		return nil, fmt.Errorf("cache encryption key unavailable")
	}

	baseURL := strings.TrimSpace(strings.TrimRight(strings.ToLower(os.Getenv("CV_CSIP")), "/"))
	raw := "commvault-cache-v1|" + baseURL + "|" + bootstrapToken
	sum := sha256.Sum256([]byte(raw))
	return sum[:], nil
}

func encryptCachedTokens(cache *TokenCacheEntry) ([]byte, error) {
	key, err := cacheEncryptionKey()
	if err != nil {
		return nil, err
	}

	plain, err := json.Marshal(cache)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, plain, nil)
	enc := EncryptedTokenCache{
		Version:    1,
		Nonce:      base64.StdEncoding.EncodeToString(nonce),
		Ciphertext: base64.StdEncoding.EncodeToString(ciphertext),
	}

	return json.Marshal(enc)
}

func decryptCachedTokens(data []byte) (TokenCacheEntry, bool) {
	var enc EncryptedTokenCache
	if err := json.Unmarshal(data, &enc); err != nil {
		return TokenCacheEntry{}, false
	}
	if enc.Version != 1 || enc.Nonce == "" || enc.Ciphertext == "" {
		return TokenCacheEntry{}, false
	}

	key, err := cacheEncryptionKey()
	if err != nil {
		return TokenCacheEntry{}, false
	}

	nonce, err := base64.StdEncoding.DecodeString(enc.Nonce)
	if err != nil {
		return TokenCacheEntry{}, false
	}
	ciphertext, err := base64.StdEncoding.DecodeString(enc.Ciphertext)
	if err != nil {
		return TokenCacheEntry{}, false
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return TokenCacheEntry{}, false
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return TokenCacheEntry{}, false
	}

	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return TokenCacheEntry{}, false
	}

	var cache TokenCacheEntry
	if err := json.Unmarshal(plain, &cache); err != nil {
		return TokenCacheEntry{}, false
	}

	return cache, true
}

// TryUseCachedTokens attempts to use cached tokens if they're still valid.
// Returns (true, nil) if cache was loaded and set in environment.
// Returns (false, nil) if cache doesn't exist or is expired.
// Returns (false, err) if there's an error reading cache.
func TryUseCachedTokens() (bool, error) {
	cache := getCachedTokens()
	if cache == nil {
		return false, nil
	}

	// Cache is valid - load into environment
	setRuntimeTokens(cache.AccessToken, cache.RefreshToken)
	return true, nil
}

func createAccessTokenPair(bootstrapToken string, tokenName string) (*AccessTokenCreateResp, error) {
	reqBodyObj := AccessTokenCreateReq{
		TokenName:               tokenName,
		RenewableUntilTimestamp: time.Now().Add(30 * 24 * time.Hour).Unix(),
		TokenExpiryTimestamp:    0,
		EnableIPAllowList:       false,
	}
	reqBody, _ := json.Marshal(&reqBodyObj)

	url := buildV4TokenURL("/AccessToken")
	respBody, err := execHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, bootstrapToken, 0)
	if err != nil {
		return nil, fmt.Errorf("create access token failed: %w", err)
	}

	var resp AccessTokenCreateResp
	if jsonErr := json.Unmarshal(respBody, &resp); jsonErr != nil {
		return nil, fmt.Errorf("create access token parse failed: %w", jsonErr)
	}
	if resp.Error.ErrorCode != 0 {
		return nil, fmt.Errorf("create access token error: code=%d message=%s", resp.Error.ErrorCode, resp.Error.ErrorMessage)
	}

	return &resp, nil
}

func buildV4TokenURL(path string) string {
	base := strings.TrimRight(os.Getenv("CV_CSIP"), "/")
	if strings.HasSuffix(strings.ToLower(base), "/v4") {
		return base + path
	}
	return base + "/V4" + path
}

func setRuntimeTokens(accessToken string, refreshToken string) {
	if accessToken != "" {
		os.Setenv("AuthToken", accessToken)
	}
	if refreshToken != "" {
		os.Setenv("CV_REFRESH_TOKEN", refreshToken)
	}
}

func fallbackBootstrapOnRenewFailure() error {
	bootstrapToken := normalizeAuthToken(os.Getenv("CV_BOOTSTRAP_TOKEN"))
	if bootstrapToken == "" {
		return fmt.Errorf("bootstrap token is empty")
	}
	if err := CreateAccessTokenWithBootstrapToken(bootstrapToken); err != nil {
		return fmt.Errorf("bootstrap fallback failed: %w", err)
	}
	return nil
}

// CreateAccessTokenWithBootstrapToken creates a new access token pair from a bootstrap token.
// It calls POST /V4/AccessToken with a renewable-until timestamp set to 30 days from now.
func CreateAccessTokenWithBootstrapToken(bootstrapToken string) error {
	bootstrapToken = normalizeAuthToken(bootstrapToken)
	if bootstrapToken == "" {
		return fmt.Errorf("api_token is empty")
	}

	resp, err := createAccessTokenPair(bootstrapToken, "AdminToken")
	if err != nil {
		if strings.Contains(err.Error(), "Token name already present") {
			fallbackName := fmt.Sprintf("AdminToken-%d", time.Now().Unix())
			resp, err = createAccessTokenPair(bootstrapToken, fallbackName)
		}
		if err != nil {
			return err
		}
	}

	if resp.TokenInfo.AccessToken == "" || resp.TokenInfo.RefreshToken == "" {
		return fmt.Errorf("create access token returned empty access or refresh token")
	}

	setRuntimeTokens(resp.TokenInfo.AccessToken, resp.TokenInfo.RefreshToken)
	// Save to cache for next provider run (non-blocking)
	saveCachedTokens(resp.TokenInfo.AccessToken, resp.TokenInfo.RefreshToken)
	return nil
}

// TryRenewWithExpiredTokenAndRefresh attempts to renew an expired bootstrap token using a refresh token.
// This is used on provider startup when bootstrap token creation fails with 401.
func TryRenewWithExpiredTokenAndRefresh(expiredToken string, refreshToken string) error {
	expiredToken = normalizeAuthToken(expiredToken)
	refreshToken = normalizeAuthToken(refreshToken)

	if expiredToken == "" || refreshToken == "" {
		return fmt.Errorf("both expired token and refresh token are required for renewal")
	}

	reqBodyObj := AccessTokenRenewReq{AccessToken: expiredToken, RefreshToken: refreshToken}
	reqBody, _ := json.Marshal(&reqBodyObj)

	url := buildV4TokenURL("/AccessToken/Renew")
	respBody, err := execHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, expiredToken, 0)
	if err != nil {
		return fmt.Errorf("renew with expired token failed: %w", err)
	}

	var resp AccessTokenRenewResp
	if jsonErr := json.Unmarshal(respBody, &resp); jsonErr != nil {
		return fmt.Errorf("renew with expired token parse failed: %w", jsonErr)
	}
	if resp.AccessToken == "" || resp.RefreshToken == "" {
		return fmt.Errorf("renew with expired token returned empty access or refresh token")
	}

	setRuntimeTokens(resp.AccessToken, resp.RefreshToken)
	// Save to cache for next provider run (non-blocking)
	saveCachedTokens(resp.AccessToken, resp.RefreshToken)
	return nil
}

// RenewAccessToken renews the active token pair using POST /V4/AccessToken/Renew.
// It updates both access and refresh token in process environment after successful renewal.
func RenewAccessToken(expiredAccessToken string) error {
	tokenRenewMutex.Lock()
	defer tokenRenewMutex.Unlock()

	currentToken := normalizeAuthToken(os.Getenv("AuthToken"))
	if currentToken != "" && currentToken != normalizeAuthToken(expiredAccessToken) {
		return nil
	}

	refreshToken := normalizeAuthToken(os.Getenv("CV_REFRESH_TOKEN"))
	if refreshToken == "" {
		if fallbackErr := fallbackBootstrapOnRenewFailure(); fallbackErr == nil {
			return nil
		}
		return fmt.Errorf("refresh token is empty and bootstrap fallback unavailable")
	}

	accessToken := normalizeAuthToken(expiredAccessToken)
	if accessToken == "" {
		accessToken = currentToken
	}
	if accessToken == "" {
		return fmt.Errorf("access token is empty")
	}

	reqBodyObj := AccessTokenRenewReq{AccessToken: accessToken, RefreshToken: refreshToken}
	reqBody, _ := json.Marshal(&reqBodyObj)

	url := buildV4TokenURL("/AccessToken/Renew")
	respBody, err := execHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, accessToken, 0)
	if err != nil {
		if fallbackErr := fallbackBootstrapOnRenewFailure(); fallbackErr == nil {
			return nil
		}
		return fmt.Errorf("renew access token failed: %w", err)
	}

	var resp AccessTokenRenewResp
	if jsonErr := json.Unmarshal(respBody, &resp); jsonErr != nil {
		return fmt.Errorf("renew access token parse failed: %w", jsonErr)
	}
	if resp.AccessToken == "" || resp.RefreshToken == "" {
		return fmt.Errorf("renew access token returned empty access or refresh token")
	}

	setRuntimeTokens(resp.AccessToken, resp.RefreshToken)
	// Save to cache for next provider run (non-blocking)
	saveCachedTokens(resp.AccessToken, resp.RefreshToken)
	return nil
}

func GenerateAuthToken(username string, password string) string {
	url := os.Getenv("CV_CSIP") + "/Login"
	loginReq := LoginReq{Mode: 5, Username: username, Password: password}
	loginJson, _ := json.Marshal(&loginReq)
	respBody, err := makeHttpRequestErr(url, http.MethodPost, XML, loginJson, JSON, "", 0)
	if err != nil {
		os.Setenv("AuthToken", "")
		return ""
	} else {
		var loginResponse DM2ContentIndexingCheckCredentialResp
		xml.Unmarshal(respBody, &loginResponse)
		os.Setenv("AuthToken", loginResponse.Token)
		fmt.Println(loginResponse.Token)
		return loginResponse.Token
	}
}

func LoginWithProviderCredentials(username string, password string) {
	url := os.Getenv("CV_CSIP") + "/login"
	loginReq := DM2ContentIndexingCheckCredentialReq{
		Username: username,
		Password: password,
		TimeOut:  "10000",
	}
	loginXML, _ := xml.Marshal(&loginReq)
	respBody, err := makeHttpRequestErr(url, http.MethodPost, XML, loginXML, XML, "", 0)
	if err != nil {
		LogEntry("LoginWithProviderCredentials", "Error: "+err.Error())
	} else {
		var loginResponse DM2ContentIndexingCheckCredentialResp
		xml.Unmarshal(respBody, &loginResponse)
		os.Setenv("AuthToken", loginResponse.Token)
	}
}

type LoginReq struct {
	Mode     int    `json:"mode"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type DM2ContentIndexingCheckCredentialReq struct {
	XMLName  xml.Name `xml:"DM2ContentIndexing_CheckCredentialReq"`
	Text     string   `xml:",chardata"`
	Username string   `xml:"username,attr"`
	Password string   `xml:"password,attr"`
	TimeOut  string   `xml:"timeOut,attr"`
}

type DM2ContentIndexingCheckCredentialResp struct {
	XMLName             xml.Name `xml:"DM2ContentIndexing_CheckCredentialResp"`
	Text                string   `xml:",chardata"`
	AliasName           string   `xml:"aliasName,attr"`
	UserGUID            string   `xml:"userGUID,attr"`
	LoginAttempts       string   `xml:"loginAttempts,attr"`
	RemainingLockTime   string   `xml:"remainingLockTime,attr"`
	SmtpAddress         string   `xml:"smtpAddress,attr"`
	UserName            string   `xml:"userName,attr"`
	ProviderType        string   `xml:"providerType,attr"`
	Ccn                 string   `xml:"ccn,attr"`
	Token               string   `xml:"token,attr"`
	Capability          string   `xml:"capability,attr"`
	ForcePasswordChange string   `xml:"forcePasswordChange,attr"`
	IsAccountLocked     string   `xml:"isAccountLocked,attr"`
	OwnerOrganization   struct {
		Text               string `xml:",chardata"`
		ProviderId         string `xml:"providerId,attr"`
		GUID               string `xml:"GUID,attr"`
		ProviderDomainName string `xml:"providerDomainName,attr"`
	} `xml:"ownerOrganization"`
	AdditionalResp struct {
		Text       string `xml:",chardata"`
		NameValues struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Value string `xml:"value,attr"`
		} `xml:"nameValues"`
	} `xml:"additionalResp"`
	ProviderOrganization struct {
		Text               string `xml:",chardata"`
		ProviderId         string `xml:"providerId,attr"`
		GUID               string `xml:"GUID,attr"`
		ProviderDomainName string `xml:"providerDomainName,attr"`
	} `xml:"providerOrganization"`
}
