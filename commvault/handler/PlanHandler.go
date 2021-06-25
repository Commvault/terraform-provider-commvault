package handler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func readFile(fileName string) *os.File {
	filePath, err := filepath.Abs(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return file
}

func PlanUpdate(rpoinmin string, slaInMinutes string, id string) {
	url := os.Getenv("CV_CSIP") + "/plan/" + id
	token := os.Getenv("AuthToken")
	var appUpdatePlan ApiUpdatePlanReq
	appUpdatePlan.Database.RpoInMinutes = rpoinmin
	appUpdatePlan.Database.SlaInMinutes = slaInMinutes
	appUpdatePlan.Summary.Addons.Database = "1"
	appUpdatePlan.Summary.Plan.PlanId = id
	appUpdatePlanXml, _ := xml.Marshal(appUpdatePlan)
	makeHttpRequest(url, http.MethodPut, XML, appUpdatePlanXml, XML, token, 0)
}

/*  func PlanCreate(createPlanRequest ApiCreatePlanReq, companyID int) *ApiCreatePlanResp {
	planCreateJSON, _ := json.Marshal(createPlanRequest)
	url := os.Getenv("CV_CSIP") + "/v3/ServerPlan"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodPost, JSON, planCreateJSON, JSON, token, companyID)
	var apiCreatePlanResp ApiCreatePlanResp
	json.Unmarshal(respBody, &apiCreatePlanResp)
	return &apiCreatePlanResp
} */

func PlanCreate(planCreateJSON string, companyID int) *V2CreatePlanResp {
	url := os.Getenv("CV_CSIP") + "/v2/Plan"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodPost, JSON, []byte(planCreateJSON), JSON, token, companyID)
	var v2CreatePlanResp V2CreatePlanResp
	json.Unmarshal(respBody, &v2CreatePlanResp)
	return &v2CreatePlanResp
}

func GetStoragePools() *GetStoragePoolListResp {
	url := os.Getenv("CV_CSIP") + "/StoragePool"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var getStoragePoolListResp GetStoragePoolListResp
	json.Unmarshal(respBody, &getStoragePoolListResp)
	return &getStoragePoolListResp
}

func PlanDelete(id string) *GenericResp {
	url := os.Getenv("CV_CSIP") + "/v2/plan/" + id + "?confirmDelete=yes"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodDelete, JSON, nil, JSON, token, 0)
	var genericResp GenericResp
	json.Unmarshal(respBody, &genericResp)
	return &genericResp
}

func AssociatePlanToVM(planName string, vmname string) {
	url := os.Getenv("CV_CSIP") + "/vm/plan"
	token := os.Getenv("AuthToken")
	var associateVMToPlanRequest AppAssociateVMToPlanRequest
	associateVMToPlanRequest.VmInfo.Plan.PlanName = planName
	associateVMToPlanRequest.VmInfo.VmClients.ClientName = vmname
	associateVmToPlanXml, _ := xml.Marshal(associateVMToPlanRequest)
	makeHttpRequest(url, http.MethodPut, XML, associateVmToPlanXml, XML, token, 0)
}
