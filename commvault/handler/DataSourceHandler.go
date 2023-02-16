package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// func urlEscape(name string) string {
// 	p := &url.URL{Path: name}
// 	return p.String()
// }

func urlEscape(name string) string {
	s := url.QueryEscape(name)
	return strings.ReplaceAll(s, "+", "%20")
}

type MsgReadUserDS struct {
	Users []struct {
		UserEntity struct {
			UserId int `json:"userId"`
		} `json:"userEntity"`
	} `json:"users"`
}

func CvGetUserByName(name string) (*MsgReadUserDS, error) {
	url := os.Getenv("CV_CSIP") + "/User?fq=" + urlEscape("name:eq:") + urlEscape(name)
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
	url := os.Getenv("CV_CSIP") + "/V4/Credential/" + urlEscape(name)
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
	url := os.Getenv("CV_CSIP") + "/Client?fq=" + urlEscape("name:eq:") + urlEscape(name)
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
	url := os.Getenv("CV_CSIP") + "/ClientGroup?fq=" + urlEscape("name:eq:") + urlEscape(name)
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
	url := os.Getenv("CV_CSIP") + "/Organization?fq=" + urlEscape("name:eq:") + urlEscape(name)
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
	url := os.Getenv("CV_CSIP") + "/Plan?fq=" + urlEscape("name:eq:") + urlEscape(name)
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgReadPlanDS
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

type MsgReadUserGroupDS struct {
	UserGroups []struct {
		UserGroupEntity struct {
			UserGroupId int `json:"userGroupId"`
		} `json:"userGroupEntity"`
	} `json:"userGroups"`
}

func CvUserGroupIdByName(name string) (*MsgReadUserGroupDS, error) {
	url := os.Getenv("CV_CSIP") + "/UserGroup?fq=" + urlEscape("name:eq:") + urlEscape(name)
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgReadUserGroupDS
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

type MsgReadRoleDSResp struct {
	Roles []MsgReadRoleDS `json:"roles"`
}

type MsgReadRoleDS struct {
	RoleId   int    `json:"id"`
	RoleName string `json:"name"`
}

func CvRoleIdByName(name string) (MsgReadRoleDS, error) {
	url := os.Getenv("CV_CSIP") + "/V4/Role"
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgReadRoleDSResp
	json.Unmarshal(respBody, &respObj)
	var obj MsgReadRoleDS

	for _, r := range respObj.Roles {
		if strings.EqualFold(r.RoleName, name) {
			obj = r
		}
	}

	return obj, err
}

type MsgReadRegionDSResp struct {
	Regions []MsgReadRegionDS `json:"regions"`
}

type MsgReadRegionDS struct {
	RegionId   int    `json:"regionId"`
	RegionName string `json:"regionName"`
}

func CvRegionIdByName(name string) (MsgReadRegionDS, error) {
	url := os.Getenv("CV_CSIP") + "/regions"
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgReadRegionDSResp
	json.Unmarshal(respBody, &respObj)
	var obj MsgReadRegionDS

	for _, r := range respObj.Regions {
		if strings.EqualFold(r.RegionName, name) {
			obj = r
		}
	}

	return obj, err
}

type MsgReadTimezoneDSResp struct {
	Timezones []MsgReadTimezoneDS `json:"timezones"`
}

type MsgReadTimezoneDS struct {
	TimezoneId   int    `json:"tzId"`
	TimezoneName string `json:"timezoneStdName"`
}

func CvTimezoneIdByName(name string) (MsgReadTimezoneDS, error) {
	url := os.Getenv("CV_CSIP") + "/GetCommServTimeZones"
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgReadTimezoneDSResp
	json.Unmarshal(respBody, &respObj)
	var obj MsgReadTimezoneDS

	for _, r := range respObj.Timezones {
		if strings.EqualFold(r.TimezoneName, name) {
			obj = r
		}
	}

	return obj, err
}

type MsgReadStoragePoolDSResp struct {
	StoragePoolList []MsgReadStoragePoolDS `json:"storagePoolList"`
}

type MsgReadStoragePoolDS struct {
	StoragePolicyEntity struct {
		StoragePolicyName string `json:"storagePolicyName"`
		StoragePolicyId   int    `json:"storagePolicyId"`
	} `json:"storagePolicyEntity"`
}

func CvStoragePoolIdByName(name string) (MsgReadStoragePoolDS, error) {
	url := os.Getenv("CV_CSIP") + "/StoragePool?storageSubType=2"
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgReadStoragePoolDSResp
	json.Unmarshal(respBody, &respObj)
	var obj MsgReadStoragePoolDS

	for _, r := range respObj.StoragePoolList {
		if strings.EqualFold(r.StoragePolicyEntity.StoragePolicyName, name) {
			obj = r
		}
	}

	return obj, err
}
