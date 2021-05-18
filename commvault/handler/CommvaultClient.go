package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var XML = "application/xml"
var JSON = "application/json"

func makeHttpRequest(url string, method string, accept string, requestBody []byte, contentType string, authToken string, companyID int) []byte {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", accept)
	req.Header.Set("AuthToken", authToken)
	if method == "POST" {
		req.Header.Set("operatorCompanyId", strconv.Itoa(companyID))
	}
	if err != nil {
		panic(err)
	}
	client := &http.Client{Timeout: time.Second * 1000}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return data
}
