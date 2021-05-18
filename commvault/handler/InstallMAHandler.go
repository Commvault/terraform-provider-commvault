package handler

import (
	"encoding/json"
	"net/http"
	"os"
)

func InstallMA(installMAReq InstallMARequest, companyID int) *InstallMAResp {
	InstallMAReqJSON, _ := json.Marshal(installMAReq)
	url := os.Getenv("CV_CSIP") + "/InstallClient"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodPost, JSON, InstallMAReqJSON, JSON, token, companyID)
	var installMAResp InstallMAResp
	json.Unmarshal(respBody, &installMAResp)
	return &installMAResp
}

func JobStatus(jobId string) *JobSummaryResponse {
	url := os.Getenv("CV_CSIP") + "/Job/" + jobId
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var jobSummaryResponse JobSummaryResponse
	json.Unmarshal(respBody, &jobSummaryResponse)
	return &jobSummaryResponse
}

func GetClientID(clientName string) *ClientDetails {
	url := os.Getenv("CV_CSIP") + "/GetId?clientname=" + clientName
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var clientDetailsResp ClientDetails
	json.Unmarshal(respBody, &clientDetailsResp)
	return &clientDetailsResp
}

func UninstallMA(id string) *UninstalResp {
	url := os.Getenv("CV_CSIP") + "/Client/" + id + "/Retire"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodDelete, JSON, nil, JSON, token, 0)
	var uninstalResp UninstalResp
	json.Unmarshal(respBody, &uninstalResp)
	return &uninstalResp
}
