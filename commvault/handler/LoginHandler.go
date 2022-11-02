package handler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
)

func GenerateAuthToken(username string, password string) string {
	url := os.Getenv("CV_CSIP") + "/Login"
	loginReq := LoginReq{Mode: 5, Username: username, Password: password}
	loginJson, _ := json.Marshal(&loginReq)
	respBody, err := makeHttpRequestErr(url, http.MethodPost, XML, loginJson, JSON, "", 0)
	if err != nil {
		panic(err)
	}
	var loginResponse DM2ContentIndexingCheckCredentialResp
	xml.Unmarshal(respBody, &loginResponse)
	os.Setenv("AuthToken", loginResponse.Token)
	fmt.Println(loginResponse.Token)
	return loginResponse.Token
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
		panic(err)
	}
	var loginResponse DM2ContentIndexingCheckCredentialResp
	xml.Unmarshal(respBody, &loginResponse)
	os.Setenv("AuthToken", loginResponse.Token)
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
