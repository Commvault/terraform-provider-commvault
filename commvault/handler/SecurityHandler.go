package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

func SecurityAssociation(securityAssociations SecurityAssociations) []byte {
	securityAssociationJSON, _ := json.Marshal(securityAssociations)
	url := os.Getenv("CV_CSIP") + "/Security"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodPost, JSON, securityAssociationJSON, JSON, token, 0)
	return respBody
}

func CreateSecurityAssociationV2(securityAssociations SecurityAssociationsV2) ([]byte, error) {
	securityAssociationJSON, _ := json.Marshal(securityAssociations)
	url := os.Getenv("CV_CSIP") + "/Security"
	token := os.Getenv("AuthToken")
	//respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, securityAssociationJSON, JSON, token, 0)
	req := buildHttpReq(url, http.MethodPost, JSON, securityAssociationJSON, JSON, token, 0)
	req.Header.Set("LookupNames", "1")
	respBody, err := executeHttpReq(req)
	return respBody, err
}

func GetSecurityAssociationV2(entityType int, entityId int) ([]byte, error) {
	url := os.Getenv("CV_CSIP") + "/Security/" + strconv.Itoa(entityType) + "/" + strconv.Itoa(entityId)
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	return respBody, err
}

func DeleteSecurityAssociationV2(securityAssociations SecurityAssociationsV2) ([]byte, error) {
	securityAssociationJSON, _ := json.Marshal(securityAssociations)
	url := os.Getenv("CV_CSIP") + "/Security"
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, securityAssociationJSON, JSON, token, 0)
	return respBody, err
}
