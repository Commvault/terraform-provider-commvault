package handler

import (
	"encoding/json"
	"net/http"
	"os"
)

func CreateStorage(createStorageReq CreateStorageReq, companyID int) *CreateStorageResp {
	storageCreate, _ := json.Marshal(createStorageReq)
	url := os.Getenv("CV_CSIP") + "/StoragePool?Action=create"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodPost, JSON, storageCreate, JSON, token, companyID)
	var apiCreateStorageResp CreateStorageResp
	json.Unmarshal(respBody, &apiCreateStorageResp)
	return &apiCreateStorageResp
}

func DeleteStorage(id string) *GenericResp {
	url := os.Getenv("CV_CSIP") + "/StoragePool/" + id
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodDelete, JSON, nil, JSON, token, 0)
	var genericResp GenericResp
	json.Unmarshal(respBody, &genericResp)
	return &genericResp
}
