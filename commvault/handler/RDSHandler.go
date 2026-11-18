package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func rdsOperation(method, path string, request interface{}) (*MsgRDSOperationResponse, error) {
	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("encode RDS request: %w", err)
		}
	}
	data, err := makeHttpRequestErr(os.Getenv("CV_CSIP")+path, method, JSON, body, JSON, os.Getenv("AuthToken"), 0)
	if err != nil {
		return nil, err
	}
	var response MsgRDSOperationResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("decode RDS response for %s: %w", path, err)
	}
	if response.ErrorCode != nil && *response.ErrorCode != 0 {
		return nil, rdsAPIError(*response.ErrorCode, response.ErrorMessage)
	}
	if len(response.Response) == 0 {
		return nil, fmt.Errorf("RDS %s %s returned no operation results", method, path)
	}
	for _, result := range response.Response {
		if result.ErrorCode == nil {
			return nil, fmt.Errorf("RDS %s %s result is missing errorCode", method, path)
		}
		if *result.ErrorCode != 0 {
			message := result.ErrorMessage
			if message == nil {
				message = result.ErrorString
			}
			return nil, rdsAPIError(*result.ErrorCode, message)
		}
		if result.WarningCode != nil && *result.WarningCode != 0 {
			message := ""
			if result.WarningMessage != nil {
				message = *result.WarningMessage
			}
			log.Printf("[WARN] RDS API warning %d: %s", *result.WarningCode, message)
		}
	}
	return &response, nil
}

func rdsAPIError(code int, message *string) error {
	detail := ""
	if message != nil {
		detail = *message
	}
	return fmt.Errorf("RDS API error %d: %s", code, detail)
}

func rdsPathID(id string) error {
	value, err := strconv.Atoi(id)
	if err != nil || value <= 0 {
		return fmt.Errorf("RDS entity ID must be a positive integer, got %q", id)
	}
	return nil
}

func CvCreateRDSInstance(req MsgCreateRDSInstanceRequest) (*MsgRDSOperationResponse, error) {
	response, err := rdsOperation(http.MethodPost, "/Instance", req)
	if err != nil {
		return nil, err
	}
	for _, result := range response.Response {
		if result.Entity != nil && result.Entity.InstanceId != nil && *result.Entity.InstanceId > 0 {
			return response, nil
		}
	}
	return nil, fmt.Errorf("RDS instance creation returned no positive instanceId")
}

func CvModifyRDSInstance(id string, req MsgModifyRDSInstanceRequest) (*MsgRDSOperationResponse, error) {
	if err := rdsPathID(id); err != nil {
		return nil, err
	}
	return rdsOperation(http.MethodPost, "/Instance/"+id, req)
}

func CvCreateRDSInstanceGroup(req MsgCreateRDSInstanceGroupRequest) (*MsgRDSOperationResponse, error) {
	response, err := rdsOperation(http.MethodPost, "/subclient", req)
	if err != nil {
		return nil, err
	}
	for _, result := range response.Response {
		if result.Entity != nil && result.Entity.SubclientId != nil && *result.Entity.SubclientId > 0 {
			return response, nil
		}
	}
	return nil, fmt.Errorf("RDS instance group creation returned no positive subclientId")
}

func CvModifyRDSInstanceGroup(id string, req MsgModifyRDSInstanceGroupRequest) (*MsgRDSOperationResponse, error) {
	if err := rdsPathID(id); err != nil {
		return nil, err
	}
	return rdsOperation(http.MethodPost, "/Subclient/"+id, req)
}

func rdsRead(path string, response interface{}) error {
	data, err := makeHttpRequestErr(os.Getenv("CV_CSIP")+path, http.MethodGet, JSON, nil, JSON, os.Getenv("AuthToken"), 0)
	if err != nil {
		return err
	}
	var envelope struct {
		ErrorCode    *int                    `json:"errorCode"`
		ErrorMessage *string                 `json:"errorMessage"`
		Errors       []MsgRDSOperationResult `json:"errors"`
	}
	if err := json.Unmarshal(data, &envelope); err != nil {
		return fmt.Errorf("decode RDS GET %s: %w", path, err)
	}
	if envelope.ErrorCode != nil && *envelope.ErrorCode != 0 {
		return rdsAPIError(*envelope.ErrorCode, envelope.ErrorMessage)
	}
	for _, result := range envelope.Errors {
		if result.ErrorCode == nil {
			return fmt.Errorf("RDS GET %s error result is missing errorCode", path)
		}
		if *result.ErrorCode != 0 {
			message := result.ErrorMessage
			if message == nil {
				message = result.ErrorString
			}
			return rdsAPIError(*result.ErrorCode, message)
		}
	}
	if err := json.Unmarshal(data, response); err != nil {
		return fmt.Errorf("decode RDS GET %s properties: %w", path, err)
	}
	return nil
}

func CvGetRDSInstance(id string) ([]MsgRDSInstanceProperties, error) {
	if err := rdsPathID(id); err != nil {
		return nil, err
	}
	var response MsgGetRDSInstanceResponse
	if err := rdsRead("/Instance/"+id, &response); err != nil {
		return nil, err
	}
	if response.InstanceProperties == nil {
		return nil, fmt.Errorf("RDS GET instance %s is missing instanceProperties", id)
	}
	return *response.InstanceProperties, nil
}

func CvGetRDSInstanceGroup(id string) ([]MsgRDSReadSubclientProperties, error) {
	if err := rdsPathID(id); err != nil {
		return nil, err
	}
	var response MsgGetRDSInstanceGroupResponse
	if err := rdsRead("/Subclient/"+id, &response); err != nil {
		return nil, err
	}
	if response.SubclientProperties == nil {
		return nil, fmt.Errorf("RDS GET group %s is missing subClientProperties", id)
	}
	return *response.SubclientProperties, nil
}

func CvDeleteRDSInstance(id string) (*MsgRDSOperationResponse, error) {
	if err := rdsPathID(id); err != nil {
		return nil, err
	}
	return rdsOperation(http.MethodDelete, "/Instance/"+id, nil)
}

func CvDeleteRDSInstanceGroup(id string) (*MsgRDSOperationResponse, error) {
	if err := rdsPathID(id); err != nil {
		return nil, err
	}
	return rdsOperation(http.MethodDelete, "/Subclient/"+id, nil)
}
