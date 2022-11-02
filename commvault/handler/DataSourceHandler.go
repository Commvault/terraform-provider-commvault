package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)

type MsgReadUserDS struct {
	Users []struct {
		UserEntity struct {
			UserId int `json:"userId"`
		} `json:"userEntity"`
	} `json:"users"`
}

func CvGetUserByName(name string) (*MsgReadUserDS, error) {
	url := os.Getenv("CV_CSIP") + "/User?fq=" + url.QueryEscape("name:eq:") + url.QueryEscape(name)
	token := os.Getenv("AuthToken")
	//respBody, err := makeHttpRequest(url, http.MethodGet, JSON, nil, JSON, token, 0)
	req := buildHttpReq(url, http.MethodGet, JSON, nil, JSON, token, 0)
	req.Header.Set("mode", "EdgeMode")
	respBody, err := executeHttpReq(req)
	//LogEntry("Response: ", string(respBody))
	var respObj MsgReadUserDS
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

type MsgReadCredentialDS struct {
	Id int `json:"id"`
}

func CvCredentialByName(name string) (*MsgReadCredentialDS, error) {
	url := os.Getenv("CV_CSIP") + "/V4/Credential/" + url.QueryEscape(name)
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgReadCredentialDS
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

type MsgReadClientDS struct {
	ClientProperties []struct {
		Client struct {
			ClientEntity struct {
				ClientId int `json:"clientId"`
			} `json:"clientEntity"`
		} `json:"client"`
	} `json:"clientProperties"`
}

func CvClientIdByName(name string) (*MsgReadClientDS, error) {
	url := os.Getenv("CV_CSIP") + "/Client?fq=" + url.QueryEscape("name:eq:") + url.QueryEscape(name)
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgReadClientDS
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

type MsgReadClientGroupDS struct {
	Groups []struct {
		ClientGroup struct {
			ClientGroupId int `json:"clientGroupId"`
		} `json:"clientGroup"`
	} `json:"groups"`
}

func CvClientGroupIdByName(name string) (*MsgReadClientGroupDS, error) {
	url := os.Getenv("CV_CSIP") + "/ClientGroup?fq=" + url.QueryEscape("name:eq:") + url.QueryEscape(name)
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgReadClientGroupDS
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

type MsgReadCompanyDS struct {
	Providers []struct {
		ShortName struct {
			Id int `json:"id"`
		} `json:"shortName"`
	} `json:"providers"`
}

func CvCompanyIdByName(name string) (*MsgReadCompanyDS, error) {
	url := os.Getenv("CV_CSIP") + "/Organization?fq=" + url.QueryEscape("name:eq:") + url.QueryEscape(name)
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgReadCompanyDS
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

type MsgReadPlanDS struct {
	Plans []struct {
		Plan struct {
			PlanId int `json:"planId"`
		} `json:"plan"`
	} `json:"plans"`
}

func CvPlanIdByName(name string) (*MsgReadPlanDS, error) {
	url := os.Getenv("CV_CSIP") + "/Plan?fq=" + url.QueryEscape("name:eq:") + url.QueryEscape(name)
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgReadPlanDS
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}
