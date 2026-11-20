package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func s3ObjectStorageRequest(method, path string, request, response interface{}) error {
	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return fmt.Errorf("encode S3 request: %w", err)
		}
	}
	data, err := makeHttpRequestErr(os.Getenv("CV_CSIP")+path, method, JSON, body, JSON, os.Getenv("AuthToken"), 0)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, response); err != nil {
		return fmt.Errorf("decode S3 response for %s %s: %w", method, path, err)
	}
	return nil
}

// S3 API envelopes alternate between a single object and an array.
func DecodeS3Objects(data json.RawMessage, target interface{}) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		return fmt.Errorf("missing S3 response properties")
	}
	if data[0] == '{' {
		data = append(append([]byte{'['}, data...), ']')
	}
	return json.Unmarshal(data, target)
}

func s3ResponseError(code *int, message, errorString *string) error {
	if code == nil || *code == 0 {
		return nil
	}
	detail := ""
	if message != nil {
		detail = *message
	} else if errorString != nil {
		detail = *errorString
	}
	return fmt.Errorf("S3 API error %d: %s", *code, detail)
}

func s3Operation(method, path string, req interface{}) ([]MsgS3OperationResult, error) {
	var resp MsgS3OperationResponse
	if err := s3ObjectStorageRequest(method, path, req, &resp); err != nil {
		return nil, err
	}
	if err := s3ResponseError(resp.ErrorCode, resp.ErrorMessage, resp.ErrorString); err != nil {
		return nil, err
	}
	if len(resp.Response) == 0 && resp.ErrorCode != nil {
		return []MsgS3OperationResult{{ErrorCode: resp.ErrorCode}}, nil
	}
	var results []MsgS3OperationResult
	if err := DecodeS3Objects(resp.Response, &results); err != nil {
		return nil, fmt.Errorf("decode S3 operation result: %w", err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("S3 operation returned no results")
	}
	for _, result := range results {
		if err := s3ResponseError(result.ErrorCode, result.ErrorMessage, result.ErrorString); err != nil {
			return nil, err
		}
		if result.ErrorCode == nil {
			return nil, fmt.Errorf("S3 operation result is missing errorCode")
		}
		if result.WarningCode != nil && *result.WarningCode != 0 {
			message := ""
			if result.WarningMessage != nil {
				message = *result.WarningMessage
			}
			log.Printf("[WARN] S3 API warning %d: %s", *result.WarningCode, message)
		}
	}
	return results, nil
}

func CvCreateS3ObjectStorage(req MsgCreateS3ObjectStorageRequest) ([]MsgS3OperationResult, error) {
	return s3Operation(http.MethodPost, "/Client", req)
}

func CvGetS3ObjectStorage(id string) ([]MsgS3InstanceProperties, error) {
	var resp MsgGetS3ObjectStorageResponse
	if err := s3ObjectStorageRequest(http.MethodGet, "/instance/"+url.PathEscape(id), nil, &resp); err != nil {
		return nil, err
	}
	if err := s3ResponseError(resp.ErrorCode, resp.ErrorMessage, nil); err != nil {
		return nil, err
	}
	var properties []MsgS3InstanceProperties
	err := DecodeS3Objects(resp.InstanceProperties, &properties)
	return properties, err
}

func CvFindS3ObjectStorage(clientName, instanceName string) (*MsgS3Entity, error) {
	query := url.Values{"clientName": {clientName}, "appName": {"Cloud Apps"}}
	var resp MsgGetS3ObjectStorageResponse
	if err := s3ObjectStorageRequest(http.MethodGet, "/instance?"+query.Encode(), nil, &resp); err != nil {
		return nil, err
	}
	if err := s3ResponseError(resp.ErrorCode, resp.ErrorMessage, nil); err != nil {
		return nil, err
	}
	var properties []MsgS3InstanceProperties
	if err := DecodeS3Objects(resp.InstanceProperties, &properties); err != nil {
		return nil, err
	}
	var match *MsgS3Entity
	for _, prop := range properties {
		entity := prop.Instance
		if entity != nil && entity.InstanceName != nil && *entity.InstanceName == instanceName &&
			entity.InstanceId != nil && *entity.InstanceId > 0 {
			if match != nil {
				return nil, fmt.Errorf("multiple Cloud Apps instances named %q on client %q", instanceName, clientName)
			}
			match = entity
		}
	}
	if match == nil {
		return nil, fmt.Errorf("Cloud Apps instance %q on client %q not found", instanceName, clientName)
	}
	return match, nil
}

func CvModifyS3ObjectStorage(id string, req MsgModifyS3ObjectStorageRequest) ([]MsgS3OperationResult, error) {
	return s3Operation(http.MethodPost, "/Instance/"+url.PathEscape(id), req)
}

func CvDeleteS3ObjectStorage(id string, req MsgDeleteS3ObjectStorageRequest) ([]MsgS3OperationResult, error) {
	return s3Operation(http.MethodPost, "/Instance/"+url.PathEscape(id)+"/Deconfigure", req)
}

func CvCreateS3Subclient(req MsgS3SubclientRequest) ([]MsgS3OperationResult, error) {
	return s3Operation(http.MethodPost, "/Subclient", req)
}

func CvGetS3Subclient(id string) ([]MsgS3SubclientProperties, error) {
	var resp MsgGetS3SubclientResponse
	if err := s3ObjectStorageRequest(http.MethodGet, "/subclient/"+url.PathEscape(id), nil, &resp); err != nil {
		return nil, err
	}
	if err := s3ResponseError(resp.ErrorCode, resp.ErrorMessage, nil); err != nil {
		return nil, err
	}
	var properties []MsgS3SubclientProperties
	err := DecodeS3Objects(resp.SubClientProperties, &properties)
	return properties, err
}

func CvModifyS3Subclient(id string, req MsgS3SubclientRequest) ([]MsgS3OperationResult, error) {
	return s3Operation(http.MethodPost, "/subclient/"+url.PathEscape(id), req)
}

func CvModifyS3SubclientContent(id string, req MsgS3SubclientRequest) ([]MsgS3OperationResult, error) {
	return s3Operation(http.MethodPost, "/Subclient/"+url.PathEscape(id), req)
}

func CvDeleteS3Subclient(id string) ([]MsgS3OperationResult, error) {
	return s3Operation(http.MethodDelete, "/subclient/"+url.PathEscape(id), nil)
}

func CvGetS3AccessNodes(excludeGroups, excludeClients bool, regionID int) (*MsgGetS3AccessNodesResponse, error) {
	query := url.Values{
		"instanceType":        {"AMAZON_S3"},
		"excludeClientGroups": {strconv.FormatBool(excludeGroups)},
		"excludeClients":      {strconv.FormatBool(excludeClients)},
	}
	if regionID > 0 {
		query.Set("regionId", strconv.Itoa(regionID))
	}
	var resp MsgGetS3AccessNodesResponse
	if err := s3ObjectStorageRequest(http.MethodGet, "/CloudApps/AccessNodes?"+query.Encode(), nil, &resp); err != nil {
		return nil, err
	}
	if err := s3ResponseError(resp.ErrorCode, resp.ErrorMessage, nil); err != nil {
		return nil, err
	}
	if resp.AccessNodes == nil {
		return nil, fmt.Errorf("S3 access-node response is missing accessNodes")
	}
	return &resp, nil
}
