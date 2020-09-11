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
	makeHttpRequest(url, http.MethodPut, XML, appUpdatePlanXml, XML, token)
}

func PlanCreate(createPlanRequest ApiCreatePlanReq) *ApiCreatePlanResp {
	panCreateJson, _ := json.Marshal(createPlanRequest)
	fmt.Print(panCreateJson)
	url := os.Getenv("CV_CSIP") + "v3/ServerPlan"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodPost, JSON, panCreateJson, JSON, token)
	var apiCreatePlanResp ApiCreatePlanResp
	json.Unmarshal(respBody, &apiCreatePlanResp)
	return &apiCreatePlanResp
}

func PlanDelete(id string) *GenericResp {
	url := os.Getenv("CV_CSIP") + "/plan/" + id + "?confirmDelete=yes"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodDelete, JSON, nil, JSON, token)
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
	makeHttpRequest(url, http.MethodPut, XML, associateVmToPlanXml, XML, token)
}
