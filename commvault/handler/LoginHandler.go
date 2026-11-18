package handler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var tokenRenewMutex sync.Mutex

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
}

type AccessTokenCreateResp struct {
	Error     AccessTokenError `json:"error"`
	TokenInfo AccessTokenInfo  `json:"tokenInfo"`
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

func setRuntimeTokens(accessToken string) {
	if accessToken != "" {
		os.Setenv("AuthToken", accessToken)
	}
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

	if resp.TokenInfo.AccessToken == "" {
		return fmt.Errorf("create access token returned empty access token")
	}

	setRuntimeTokens(resp.TokenInfo.AccessToken)
	return nil
}

// RenewAccessToken renews the active access token by re-creating a new one from the bootstrap token.
// Refresh-token-based renewal is not supported; renewal always re-runs the bootstrap token exchange.
func RenewAccessToken(expiredAccessToken string) error {
	tokenRenewMutex.Lock()
	defer tokenRenewMutex.Unlock()

	currentToken := normalizeAuthToken(os.Getenv("AuthToken"))
	if currentToken != "" && currentToken != normalizeAuthToken(expiredAccessToken) {
		return nil
	}

	bootstrapToken := normalizeAuthToken(os.Getenv("CV_BOOTSTRAP_TOKEN"))
	if bootstrapToken == "" {
		return fmt.Errorf("bootstrap token is empty, cannot renew access token")
	}
	if err := CreateAccessTokenWithBootstrapToken(bootstrapToken); err != nil {
		return fmt.Errorf("renew access token failed: %w", err)
	}
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
