package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type MsgCloudInstanceType struct {
	InstanceType json.RawMessage `json:"instanceType"`
}

type MsgCloudInstanceState struct {
	Instance          *MsgS3Entity          `json:"instance"`
	CloudAppsInstance *MsgCloudInstanceType `json:"cloudAppsInstance"`
	Status            *int                  `json:"status"`
	IsDeleted         *bool                 `json:"isDeleted"`
}

func CvGetCloudInstanceState(id string) (*MsgCloudInstanceState, string, error) {
	if err := rdsPathID(id); err != nil {
		return nil, "", err
	}
	var response struct {
		InstanceProperties *[]MsgCloudInstanceState `json:"instanceProperties"`
	}
	if err := rdsRead("/Instance/"+id, &response); err != nil {
		return nil, "", err
	}
	if response.InstanceProperties == nil || len(*response.InstanceProperties) != 1 {
		return nil, "", fmt.Errorf("expected one cloud instance for %s", id)
	}
	state := &(*response.InstanceProperties)[0]
	if state.CloudAppsInstance == nil {
		return nil, "", fmt.Errorf("instance %s has no cloud workload identity", id)
	}
	workload, err := cloudNamedEnum(state.CloudAppsInstance.InstanceType, map[string]string{"4": "AMAZON_RDS", "5": "AMAZON_S3"})
	if err != nil {
		return nil, "", err
	}
	if workload == nil {
		return nil, "", fmt.Errorf("instance %s has no instanceType", id)
	}
	return state, *workload, nil
}

func CvDeconfigureCloudInstance(id string, req MsgDeleteS3ObjectStorageRequest) error {
	if err := rdsPathID(id); err != nil {
		return err
	}
	_, err := s3Operation(http.MethodPost, "/Instance/"+id+"/Deconfigure", req)
	return err
}

func CvGetCloudAccessNodeName(id int) (*string, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid access-node ID %d", id)
	}
	var response struct {
		ClientProperties []struct {
			Client *struct {
				ClientEntity *MsgS3Entity `json:"clientEntity"`
			} `json:"client"`
		} `json:"clientProperties"`
	}
	if err := rdsRead("/Client/"+strconv.Itoa(id), &response); err != nil {
		return nil, err
	}
	if len(response.ClientProperties) != 1 || response.ClientProperties[0].Client == nil {
		return nil, fmt.Errorf("missing client properties for access node %d", id)
	}
	entity := response.ClientProperties[0].Client.ClientEntity
	if entity == nil || entity.ClientId == nil || *entity.ClientId != id || entity.ClientName == nil || *entity.ClientName == "" {
		return nil, fmt.Errorf("client properties do not identify access node %d", id)
	}
	return entity.ClientName, nil
}
