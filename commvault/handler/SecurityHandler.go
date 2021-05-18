package handler

import (
	"encoding/json"
	"net/http"
	"os"
)

func SecurityAssociation(securityAssociations SecurityAssociations) []byte {
	securityAssociationJSON, _ := json.Marshal(securityAssociations)
	url := os.Getenv("CV_CSIP") + "/Security"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodPost, JSON, securityAssociationJSON, JSON, token, 0)
	return respBody
}
