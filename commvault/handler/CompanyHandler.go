package handler

import (
	"encoding/json"
	"net/http"
	"os"
)

func CompanyUpdate(rpoinmin string, slaInMinutes string, id string) {
}

func CompanyCreate(createCompanyReq CreateCompanyReq, companyID int) *ApiCreateCompanyResp {
	companyCreateJSON, _ := json.Marshal(createCompanyReq)
	url := os.Getenv("CV_CSIP") + "/Organization"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodPost, JSON, companyCreateJSON, JSON, token, companyID)
	var apiCreateCompanyResp ApiCreateCompanyResp
	json.Unmarshal(respBody, &apiCreateCompanyResp)
	return &apiCreateCompanyResp
}

func CompanyDeactivate(id string) *ApiCreateCompanyResp {
	url := os.Getenv("CV_CSIP") + "/Organization/" + id + "/action/deactivate"
	token := os.Getenv("AuthToken")
	var deactivateCompany DeactivateCompany
	deactivateCompany.DeactivateOptions.DisableBackup = true
	deactivateCompany.DeactivateOptions.DisableLogin = true
	deactivateCompany.DeactivateOptions.DisableRestore = true
	companyDeactivateJSON, _ := json.Marshal(deactivateCompany)
	respBody := makeHttpRequest(url, http.MethodPost, JSON, companyDeactivateJSON, JSON, token, 0)
	var genericResp ApiCreateCompanyResp
	json.Unmarshal(respBody, &genericResp)
	return &genericResp
}

func CompanyDelete(id string) *GenericResp {
	url := os.Getenv("CV_CSIP") + "/Organization/" + id
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodDelete, JSON, nil, JSON, token, 0)
	var genericResp GenericResp
	json.Unmarshal(respBody, &genericResp)
	return &genericResp
}
