package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type MsgRDSInstance struct {
	InstanceId    *int    `json:"instanceId,omitempty"`
	InstanceName  *string `json:"instanceName,omitempty"`
	ApplicationId *int    `json:"applicationId,omitempty"`
	ClientName    *string `json:"clientName,omitempty"`
	ClientId      *int    `json:"clientId,omitempty"`
}

type MsgRDSPlan struct {
	PlanName *string `json:"planName,omitempty"`
	PlanId   *int    `json:"planId,omitempty"`
	Name     *string `json:"name,omitempty"`
}

type MsgRDSAccessNodeClient struct {
	ClientId   *int    `json:"clientId,omitempty"`
	ClientName *string `json:"clientName,omitempty"`
	Type       *string `json:"_type_,omitempty"`
}

func (m *MsgRDSAccessNodeClient) UnmarshalJSON(data []byte) error {
	type alias MsgRDSAccessNodeClient
	var response struct {
		*alias
		Type json.RawMessage `json:"_type_"`
	}
	response.alias = (*alias)(m)
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}
	var err error
	m.Type, err = cloudNamedEnum(response.Type, map[string]string{"3": "CLIENT_ENTITY"})
	return err
}

type MsgRDSMemberServer struct {
	Client *MsgRDSAccessNodeClient `json:"client,omitempty"`
}

type MsgRDSAccessNodes struct {
	MemberServers []MsgRDSMemberServer `json:"memberServers"`
}

type MsgRDSAmazonInstanceInfo struct {
	UseIamRole *bool `json:"useIamRole,omitempty"`
}

type MsgRDSGeneralCloudProperties struct {
	NumberOfBackupStreams              *int                      `json:"numberOfBackupStreams,omitempty"`
	RegionEndPoints                    *string                   `json:"regionEndPoints,omitempty"`
	PointInTimeRestoreEnabled          *bool                     `json:"pointInTimeRestoreEnabled,omitempty"`
	PointInTimeRestoreRetentionPeriod  *int                      `json:"pointInTimeRestoreRetentionPeriod,omitempty"`
	PointInTimeRestoreBackupPeriod     *string                   `json:"pointInTimeRestoreBackupPeriod,omitempty"`
	PointInTimeRestoreApplyImmediately *bool                     `json:"pointInTimeRestoreApplyImmediately,omitempty"`
	AccessNodes                        *MsgRDSAccessNodes        `json:"accessNodes,omitempty"`
	Credentials                        *MsgCredentialInfo        `json:"credentials,omitempty"`
	AmazonInstanceInfo                 *MsgRDSAmazonInstanceInfo `json:"amazonInstanceInfo,omitempty"`
}

type MsgRDSCloudAppsInstance struct {
	InstanceType           *string                       `json:"instanceType,omitempty"`
	GeneralCloudProperties *MsgRDSGeneralCloudProperties `json:"generalCloudProperties,omitempty"`
}

func (m *MsgRDSCloudAppsInstance) UnmarshalJSON(data []byte) error {
	type alias MsgRDSCloudAppsInstance
	var response struct {
		*alias
		InstanceType json.RawMessage `json:"instanceType"`
	}
	response.alias = (*alias)(m)
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}
	var err error
	m.InstanceType, err = cloudNamedEnum(response.InstanceType, map[string]string{"4": "AMAZON_RDS"})
	if err != nil {
		return fmt.Errorf("RDS instanceType: %w", err)
	}
	return nil
}

type MsgRDSInstanceProperties struct {
	Instance            *MsgRDSInstance          `json:"instance,omitempty"`
	CloudAppsInstance   *MsgRDSCloudAppsInstance `json:"cloudAppsInstance,omitempty"`
	PlanEntity          *MsgRDSPlan              `json:"planEntity,omitempty"`
	UseResourcePoolInfo *bool                    `json:"useResourcePoolInfo,omitempty"`
}

type MsgCreateRDSInstanceRequest struct {
	InstanceProperties *MsgRDSInstanceProperties `json:"instanceProperties,omitempty"`
}

type MsgRDSAssociationEntity struct {
	InstanceId    *int    `json:"instanceId,omitempty"`
	ClientId      *int    `json:"clientId,omitempty"`
	SubclientId   *int    `json:"subclientId,omitempty"`
	ApplicationId *int    `json:"applicationId,omitempty"`
	ClientName    *string `json:"clientName,omitempty"`
	InstanceName  *string `json:"instanceName,omitempty"`
	DisplayName   *string `json:"displayName,omitempty"`
	AppName       *string `json:"appName,omitempty"`
	Type          *string `json:"_type_,omitempty"`
	BackupsetName *string `json:"backupsetName,omitempty"`
	SubclientName *string `json:"subclientName,omitempty"`
}

type MsgRDSAssociation struct {
	Entity []MsgRDSAssociationEntity `json:"entity,omitempty"`
}

type MsgModifyRDSInstanceRequest struct {
	NewName            *string                   `json:"newName,omitempty"`
	InstanceProperties *MsgRDSInstanceProperties `json:"instanceProperties,omitempty"`
	Association        *MsgRDSAssociation        `json:"association,omitempty"`
}

type MsgRDSSubclientEntity struct {
	InstanceId    *int    `json:"instanceId,omitempty"`
	ClientId      *int    `json:"clientId,omitempty"`
	ApplicationId *int    `json:"applicationId,omitempty"`
	ClientName    *string `json:"clientName,omitempty"`
	InstanceName  *string `json:"instanceName,omitempty"`
	AppName       *string `json:"appName,omitempty"`
	DisplayName   *string `json:"displayName,omitempty"`
	SubclientId   *int    `json:"subclientId,omitempty"`
	SubclientName *string `json:"subclientName,omitempty"`
}

type MsgRDSCloudAppsSubclientProperties struct {
	InstanceType *string          `json:"instanceType,omitempty"`
	RDSSubclient *MsgRDSSubclient `json:"rdsSubclient,omitempty"`
}

type MsgRDSSubclient struct {
	AutoUpdateList                  *bool                      `json:"autoUpdateList,omitempty"`
	EnableExternalSnapshotDiscovery *bool                      `json:"enableExternalSnapshotDiscovery,omitempty"`
	EnableSnapshotBackupCopy        *bool                      `json:"enableSnapshotBackupCopy,omitempty"`
	WaitForSnapshot                 *bool                      `json:"waitForSnapshot,omitempty"`
	IsCrossAccountEnabled           *bool                      `json:"isCrossAccountEnabled,omitempty"`
	IsCrossAccountCopyEnabled       *bool                      `json:"isCrossAccountCopyEnabled,omitempty"`
	DestinationAccount              *MsgRDSCrossAccountDetails `json:"destinationAccount,omitempty"`
}

type MsgRDSCrossAccountDetails struct {
	AmazonClient *MsgRDSAccountClient `json:"amazonClient,omitempty"`
}

type MsgRDSAccountClient struct {
	ClientId   *int    `json:"clientId,omitempty"`
	ClientName *string `json:"clientName,omitempty"`
}

type MsgRDSRegionMapping struct {
	SourceRegion      *string                `json:"sourceRegion,omitempty"`
	DestinationRegion *string                `json:"destinationRegion,omitempty"`
	SubClient         *MsgRDSSubclientEntity `json:"subClient,omitempty"`
}

type MsgRDSRegionMapRequest struct {
	RegionMap []MsgRDSRegionMapping `json:"regionMap"`
}

type MsgRDSReadRegionMapping struct {
	SourceRegionName      *string `json:"srcRegionName,omitempty"`
	DestinationRegionName *string `json:"destRegionName,omitempty"`
}

type MsgRDSRegionMapResponse struct {
	RegionMap *[]MsgRDSReadRegionMapping `json:"regionMap"`
	Status    *int                       `json:"status"`
}

func (m *MsgRDSCloudAppsSubclientProperties) UnmarshalJSON(data []byte) error {
	type alias MsgRDSCloudAppsSubclientProperties
	var response struct {
		*alias
		InstanceType json.RawMessage `json:"instanceType"`
	}
	response.alias = (*alias)(m)
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}
	var err error
	m.InstanceType, err = cloudNamedEnum(response.InstanceType, map[string]string{"4": "AMAZON_RDS"})
	return err
}

// Creation uses named content types; modification uses numeric content types.
type MsgRDSCreateContent struct {
	DisplayName      *string `json:"displayName,omitempty"`
	Name             *string `json:"name,omitempty"`
	Path             *string `json:"path,omitempty"`
	Type             *string `json:"type,omitempty"`
	AllOrAnyChildren *bool   `json:"allOrAnyChildren,omitempty"`
	Negation         *bool   `json:"negation,omitempty"`
}

type MsgRDSModifyContent struct {
	DisplayName      *string `json:"displayName,omitempty"`
	Name             *string `json:"name,omitempty"`
	Path             *string `json:"path,omitempty"`
	Type             *int    `json:"type,omitempty"`
	Value            *string `json:"value,omitempty"`
	AllOrAnyChildren *bool   `json:"allOrAnyChildren,omitempty"`
	Negation         *bool   `json:"negation,omitempty"`
}

type MsgRDSCreateCloudDbContent struct {
	Children []MsgRDSCreateContent `json:"children,omitempty"`
}

type MsgRDSModifyCloudDbContent struct {
	Children []MsgRDSModifyContent `json:"children"`
}

type MsgRDSCreateSubclientProperties struct {
	CommonProperties       *MsgS3CommonProperties              `json:"commonProperties,omitempty"`
	SubclientEntity        *MsgRDSSubclientEntity              `json:"subClientEntity,omitempty"`
	PlanEntity             *MsgRDSPlan                         `json:"planEntity,omitempty"`
	ContentOperationType   *int                                `json:"contentOperationType,omitempty"`
	CloudAppsSubclientProp *MsgRDSCloudAppsSubclientProperties `json:"cloudAppsSubClientProp,omitempty"`
	CloudDbContent         *MsgRDSCreateCloudDbContent         `json:"cloudDbContent,omitempty"`
}

type MsgRDSModifySubclientProperties struct {
	CommonProperties       *MsgS3CommonProperties              `json:"commonProperties,omitempty"`
	PlanEntity             *MsgRDSPlan                         `json:"planEntity,omitempty"`
	CloudAppsSubclientProp *MsgRDSCloudAppsSubclientProperties `json:"cloudAppsSubClientProp,omitempty"`
	CloudDbContent         *MsgRDSModifyCloudDbContent         `json:"cloudDbContent,omitempty"`
}

type MsgGetRDSInstanceResponse struct {
	InstanceProperties *[]MsgRDSInstanceProperties `json:"instanceProperties"`
}

type MsgRDSReadContent struct {
	DisplayName      *string             `json:"displayName,omitempty"`
	Name             *string             `json:"name,omitempty"`
	Path             *string             `json:"path,omitempty"`
	Type             json.RawMessage     `json:"type,omitempty"`
	Value            *string             `json:"value,omitempty"`
	AllOrAnyChildren *bool               `json:"allOrAnyChildren,omitempty"`
	Negation         *bool               `json:"negation,omitempty"`
	Children         []MsgRDSReadContent `json:"children,omitempty"`
}

type MsgRDSReadSubclientProperties struct {
	CommonProperties       *MsgS3CommonProperties              `json:"commonProperties,omitempty"`
	SubclientEntity        *MsgRDSSubclientEntity              `json:"subClientEntity,omitempty"`
	PlanEntity             *MsgRDSPlan                         `json:"planEntity,omitempty"`
	CloudAppsSubclientProp *MsgRDSCloudAppsSubclientProperties `json:"cloudAppsSubClientProp,omitempty"`
	CloudDbContent         *MsgRDSReadContent                  `json:"cloudDbContent,omitempty"`
}

type MsgGetRDSInstanceGroupResponse struct {
	SubclientProperties *[]MsgRDSReadSubclientProperties `json:"subClientProperties"`
}

type MsgCreateRDSInstanceGroupRequest struct {
	SubclientProperties *MsgRDSCreateSubclientProperties `json:"subClientProperties,omitempty"`
}

type MsgModifyRDSInstanceGroupRequest struct {
	NewName             *string                          `json:"newName,omitempty"`
	SubclientProperties *MsgRDSModifySubclientProperties `json:"subClientProperties,omitempty"`
	Association         *MsgRDSAssociation               `json:"association,omitempty"`
}

type MsgRDSResultEntity struct {
	InstanceId   *int    `json:"instanceId,omitempty"`
	InstanceName *string `json:"instanceName,omitempty"`
	ClientId     *int    `json:"clientId,omitempty"`
	SubclientId  *int    `json:"subclientId,omitempty"`
	Type         *int    `json:"_type_,omitempty"`
}

type MsgRDSOperationResult struct {
	ErrorCode      *int                `json:"errorCode,omitempty"`
	ErrorMessage   *string             `json:"errorMessage,omitempty"`
	ErrorString    *string             `json:"errorString,omitempty"`
	WarningCode    *int                `json:"warningCode,omitempty"`
	WarningMessage *string             `json:"warningMessage,omitempty"`
	Entity         *MsgRDSResultEntity `json:"entity,omitempty"`
}

type MsgRDSOperationResponse struct {
	ProcessingInstructionInfo *MsgProcessingInstructionInfo `json:"processinginstructioninfo,omitempty"`
	Response                  []MsgRDSOperationResult       `json:"response,omitempty"`
	ErrorCode                 *int                          `json:"errorCode,omitempty"`
	ErrorMessage              *string                       `json:"errorMessage,omitempty"`
	Entity                    *MsgRDSResultEntity           `json:"entity,omitempty"`
}

func (m *MsgRDSOperationResponse) UnmarshalJSON(data []byte) error {
	type alias MsgRDSOperationResponse
	var envelope struct {
		*alias
		Response json.RawMessage `json:"response"`
	}
	*m = MsgRDSOperationResponse{}
	envelope.alias = (*alias)(m)
	if err := json.Unmarshal(data, &envelope); err != nil {
		return err
	}
	response := bytes.TrimSpace(envelope.Response)
	if len(response) == 0 || bytes.Equal(response, []byte("null")) {
		return nil
	}
	if response[0] == '{' {
		var result MsgRDSOperationResult
		if err := json.Unmarshal(response, &result); err != nil {
			return err
		}
		m.Response = []MsgRDSOperationResult{result}
		return nil
	}
	return json.Unmarshal(response, &m.Response)
}
