package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
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
	req, _ := buildHttpReq(url, http.MethodGet, JSON, nil, JSON, token, 0)
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

type MsgReadHyperscaleDSResp struct {
	Hyperscale []MsgReadHyperscaleDS `json:"hyperscale"`
}

type MsgReadHyperscaleDS struct {
	HyperscaleId   int    `json:"id"`
	HyperscaleName string `json:"name"`
}

func CvHyperscaleIdByName(name string) (MsgReadHyperscaleDS, error) {
	url := os.Getenv("CV_CSIP") + "/V4/Storage/Hyperscale"
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgReadHyperscaleDSResp
	json.Unmarshal(respBody, &respObj)
	var obj MsgReadHyperscaleDS

	for _, r := range respObj.Hyperscale {
		if strings.EqualFold(r.HyperscaleName, name) {
			obj = r
		}
	}

	return obj, err
}

type MsgReadRegionDSResp struct {
	Regions []MsgReadRegionEntity `json:"regions"`
}

type MsgReadRegionEntity struct {
	RegionEntity MsgReadRegionDS `json:"regionEntity"`
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
		if strings.EqualFold(r.RegionEntity.RegionName, name) {
			obj = r.RegionEntity
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

type MsgReadPermissionDSResp struct {
	Permissions []MsgReadPermissionGroupDSResp `json:"permissions"`
}

type MsgReadPermissionGroupDSResp struct {
	Permissions []MsgReadPermissionDS       `json:"permissions"`
	Category    MsgReadPermissionCategoryDS `json:"category"`
}

type MsgReadPermissionCategoryDS struct {
	CategoryId   int    `json:"id"`
	CategoryName string `json:"name"`
}

type MsgReadPermissionDS struct {
	PermissionName string `json:"name"`
	PermissionId   int    `json:"id"`
}

func CvPermissionIdByName(name string) (int, int) {
	url := os.Getenv("CV_CSIP") + "/v4/Permissions"
	token := os.Getenv("AuthToken")
	respBody, _ := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgReadPermissionDSResp
	json.Unmarshal(respBody, &respObj)
	//var obj MsgReadPermissionDS

	for _, r := range respObj.Permissions {
		for _, p := range r.Permissions {
			if strings.EqualFold(p.PermissionName, name) {
				return p.PermissionId, r.Category.CategoryId
			}
		}
	}

	return 0, 0
}

type MsgAccessPathDSResp struct {
	DiskAccessPaths []MsgAccessPathDS `json:"diskAccessPaths"`
}

type MsgAccessPathDS struct {
	AccessPathId int             `json:"id"`
	MediaAgent   MsgMediaAgentDS `json:"mediaAgent"`
}

type MsgMediaAgentDS struct {
	MediaAgentId   int    `json:"id"`
	MediaAgentName string `json:"name"`
}

func CvGetAccessPathForMediaAgent(storagePoolId string, backupLocationId string, mediaAgentId int) (int, error) {
	url := os.Getenv("CV_CSIP") + "/V4/Storage/Disk/" + storagePoolId + "/BackupLocation/" + backupLocationId
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgAccessPathDSResp
	json.Unmarshal(respBody, &respObj)
	resp := 0

	for _, r := range respObj.DiskAccessPaths {
		if r.MediaAgent.MediaAgentId == mediaAgentId {
			resp = r.AccessPathId
		}
	}

	if resp == 0 {
		LogEntry("ERROR", "could not find media agent id ["+strconv.Itoa(mediaAgentId)+"] from response ["+url+"]")
	}

	return resp, err
}

type MsgCloudAccessPathDSResp struct {
	CloudAccessPaths []MsgCloudAccessPathDS `json:"cloudAccessPaths"`
}

type MsgCloudAccessPathDS struct {
	AccessPathId int                  `json:"accessPathId"`
	MediaAgent   MsgCloudMediaAgentDS `json:"mediaAgent"`
}

type MsgCloudMediaAgentDS struct {
	MediaAgentId   int    `json:"id"`
	MediaAgentName string `json:"name"`
}

func CvGetCloudAccessPathForMediaAgent(cloudStorageId string, bucketId string, mediaAgentId int) (int, error) {
	url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud/" + cloudStorageId + "/Bucket/" + bucketId
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgCloudAccessPathDSResp
	json.Unmarshal(respBody, &respObj)
	resp := 0

	for _, r := range respObj.CloudAccessPaths {
		if r.MediaAgent.MediaAgentId == mediaAgentId {
			resp = r.AccessPathId
		}
	}

	if resp == 0 {
		LogEntry("ERROR", "could not find media agent id ["+strconv.Itoa(mediaAgentId)+"] from response ["+url+"]")
	}

	return resp, err
}

type MsgReadKubernetesDSResp struct {
	Items []MsgReadKubernetesDS `json:"items"`
}

type MsgReadKubernetesDS struct {
	KubernetesGuid string `json:"GUID"`
	KubernetesName string `json:"name"`
}

func CvKubernetesGUIDByName(query string, name string) (MsgReadKubernetesDS, error) {
	url := os.Getenv("CV_CSIP") + query
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgReadKubernetesDSResp
	json.Unmarshal(respBody, &respObj)
	var obj MsgReadKubernetesDS

	for _, r := range respObj.Items {
		if strings.EqualFold(r.KubernetesName, name) {
			obj = r
		}
	}

	return obj, err
}

func CvKubernetesNamespacesByName(clusterId int, name string) (MsgReadKubernetesDS, error) {
	url := "/V5/Kubernetes/Cluster/" + strconv.Itoa(clusterId) + "/Content/Namespace"
	return CvKubernetesGUIDByName(url, name)
}

func CvKubernetesStorageClassesByName(clusterId int, name string) (MsgReadKubernetesDS, error) {
	url := "/V5/Kubernetes/Cluster/" + strconv.Itoa(clusterId) + "/Content/StorageClass"
	return CvKubernetesGUIDByName(url, name)
}

func CvKubernetesApplicationsByName(clusterId int, namespace string, name string) (MsgReadKubernetesDS, error) {
	url := "/V5/Kubernetes/Cluster/" + strconv.Itoa(clusterId) + "/Content/Namespace/" + urlEscape(namespace) + "/Applications"
	return CvKubernetesGUIDByName(url, name)
}

func CvKubernetesVolumesByName(clusterId int, namespace string, name string) (MsgReadKubernetesDS, error) {
	url := "/V5/Kubernetes/Cluster/" + strconv.Itoa(clusterId) + "/Content/Namespace/" + urlEscape(namespace) + "/Volumes"
	return CvKubernetesGUIDByName(url, name)
}

func CvKubernetesLabelsByName(clusterId int, namespace string, name string) (MsgReadKubernetesDS, error) {
	url := "/V5/Kubernetes/Cluster/" + strconv.Itoa(clusterId) + "/Content/Namespace/" + urlEscape(namespace) + "/Labels"
	return CvKubernetesGUIDByName(url, name)
}
