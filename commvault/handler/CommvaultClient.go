package handler

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var XML = "application/xml"
var JSON = "application/json"

func buildHttpReq(url string, method string, accept string, requestBody []byte, contentType string, authToken string, companyID int) *http.Request {
	if !strings.HasSuffix(url, "login") {
		LogEntry("REQUEST: ", "("+method+") "+url+"\nBODY: "+string(requestBody))
	}
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
	return req
}

func executeHttpReq(req *http.Request) ([]byte, error) {
	client := &http.Client{Timeout: time.Second * 1000}
	if os.Getenv("IGNORE_CERT") == "true" {
		config := &tls.Config{InsecureSkipVerify: true}
		transport := &http.Transport{TLSClientConfig: config}
		client.Transport = transport
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		dst := &bytes.Buffer{}
		json.Compact(dst, data)
		return nil, fmt.Errorf("status: %d, body: %s", resp.StatusCode, dst)
	}
	if !strings.HasSuffix(req.URL.String(), "login") {
		LogEntry("RESPONSE: ", string(data))
	}
	return data, nil
}

func makeHttpRequestErr(url string, method string, accept string, requestBody []byte, contentType string, authToken string, companyID int) ([]byte, error) {
	req := buildHttpReq(url, method, accept, requestBody, contentType, authToken, companyID)
	return executeHttpReq(req)
}

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
	if os.Getenv("IGNORE_CERT") == "true" {
		config := &tls.Config{InsecureSkipVerify: true}
		transport := &http.Transport{TLSClientConfig: config}
		client.Transport = transport
	}
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

func ToStringArray(itemsRaw []interface{}) []string {
	if itemsRaw != nil {
		items := make([]string, len(itemsRaw))
		for i, raw := range itemsRaw {
			items[i] = raw.(string)
		}
		return items
	} else {
		return nil
	}
}

func ToIntArray(itemsRaw []interface{}) []int {
	if itemsRaw != nil {
		items := make([]int, len(itemsRaw))
		for i, raw := range itemsRaw {
			items[i] = raw.(int)
		}
		return items
	} else {
		return nil
	}
}

func ToLongArray(itemsRaw []interface{}) []int64 {
	if itemsRaw != nil {
		items := make([]int64, len(itemsRaw))
		for i, raw := range itemsRaw {
			items[i] = raw.(int64)
		}
		return items
	} else {
		return nil
	}
}

func ToFloatArray(itemsRaw []interface{}) []float32 {
	if itemsRaw != nil {
		items := make([]float32, len(itemsRaw))
		for i, raw := range itemsRaw {
			items[i] = raw.(float32)
		}
		return items
	} else {
		return nil
	}
}

func ToDoubleArray(itemsRaw []interface{}) []float64 {
	if itemsRaw != nil {
		items := make([]float64, len(itemsRaw))
		for i, raw := range itemsRaw {
			items[i] = raw.(float64)
		}
		return items
	} else {
		return nil
	}
}

func ToBooleanArray(itemsRaw []interface{}) []bool {
	if itemsRaw != nil {
		items := make([]bool, len(itemsRaw))
		for i, raw := range itemsRaw {
			items[i] = raw.(bool)
		}
		return items
	} else {
		return nil
	}
}

func ToStringValue(val interface{}, omitempty bool) *string {
	tmp := val.(string)
	if omitempty && tmp == "" {
		return nil
	}
	return &tmp
}

func ToIntValue(val interface{}, omitempty bool) *int {
	tmp := val.(int)
	if omitempty && tmp == 0 {
		return nil
	}
	return &tmp
}

func ToLongValue(val interface{}, omitempty bool) *int64 {
	tmp := int64(val.(int))
	if omitempty && tmp == 0 {
		return nil
	}
	return &tmp
}

func ToFloatValue(val interface{}, omitempty bool) *float32 {
	tmp := val.(float32)
	if omitempty && tmp == 0 {
		return nil
	}
	return &tmp
}

func ToDoubleValue(val interface{}, omitempty bool) *float64 {
	tmp := val.(float64)
	if omitempty && tmp == 0 {
		return nil
	}
	return &tmp
}

// func ToBooleanValue(val interface{}, omitempty bool) *bool {
// 	tmp := val.(bool)
// 	return &tmp
// }

func ToBooleanValue(val interface{}, omitempty bool) *bool {
	tmp := val.(string)
	t := true
	f := false
	if strings.EqualFold(tmp, "true") {
		return &t
	} else if strings.EqualFold(tmp, "false") {
		return &f
	}
	return nil
}

func LogEntry(prefix string, entry string) {
	if os.Getenv("CV_LOGGING") == "true" {
		f, err := os.OpenFile("terraform.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		logger := log.New(f, "", log.LstdFlags)
		logger.Println(prefix + ": " + entry)
	}
}

