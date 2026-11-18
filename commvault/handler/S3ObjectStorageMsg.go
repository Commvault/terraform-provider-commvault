package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type MsgS3Entity struct {
	ClientId      *int             `json:"clientId,omitempty"`
	ClientName    *string          `json:"clientName,omitempty"`
	InstanceId    *int             `json:"instanceId,omitempty"`
	InstanceName  *string          `json:"instanceName,omitempty"`
	SubclientId   *int             `json:"subclientId,omitempty"`
	SubclientName *string          `json:"subclientName,omitempty"`
	ApplicationId *int             `json:"applicationId,omitempty"`
	CommCellId    *int             `json:"commCellId,omitempty"`
	AppName       *string          `json:"appName,omitempty"`
	Type          *int             `json:"_type_,omitempty"`
	EntityInfo    *MsgS3EntityInfo `json:"entityInfo,omitempty"`
}

type MsgS3EntityInfo struct {
	CompanyName     *string `json:"companyName,omitempty"`
	MultiCommcellId *int    `json:"multiCommcellId,omitempty"`
}

type MsgS3Plan struct {
	PlanId   *int    `json:"planId,omitempty"`
	PlanName *string `json:"planName,omitempty"`
	Id       *int    `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
}

type MsgS3MemberServer struct {
	Client *MsgS3Entity `json:"client,omitempty"`
}

type MsgS3GeneralCloudProperties struct {
	ObjectStorageIndexingSchemaV2 *bool               `json:"objectStorageIndexingSchemaV2,omitempty"`
	ObjectStorageMultinode        *int                `json:"objectStorageMultinode,omitempty"`
	NumberOfBackupStreams         *int                `json:"numberOfBackupStreams,omitempty"`
	MemberServers                 []MsgS3MemberServer `json:"memberServers"`
	Credentials                   *MsgCredentialInfo  `json:"credentials,omitempty"`
}

type MsgS3RegionInfo struct {
	RegionName *string `json:"regionName,omitempty"`
}

type MsgS3ObjectStorageInstance struct {
	IsCustomerManagedAccessNode *bool            `json:"isCustomerManagedAccessNode,omitempty"`
	IsCloudEncryptionKeySet     *bool            `json:"isCloudEncryptionKeySet,omitempty"`
	RegionInfo                  *MsgS3RegionInfo `json:"regionInfo,omitempty"`
}

type MsgS3Connection struct {
	HostURL    *string `json:"hostURL,omitempty"`
	UseIamRole *bool   `json:"useIamRole,omitempty"`
}

type MsgS3CloudAppsInstance struct {
	InstanceTypeDisplayName *string                      `json:"instanceTypeDisplayName,omitempty"`
	InstanceType            *string                      `json:"instanceType,omitempty"`
	GeneralCloudProperties  *MsgS3GeneralCloudProperties `json:"generalCloudProperties,omitempty"`
	ObjectStorageInstance   *MsgS3ObjectStorageInstance  `json:"objectStorageInstance,omitempty"`
	S3Instance              *MsgS3Connection             `json:"s3Instance,omitempty"`
	CredentialType          *string                      `json:"credentialType,omitempty"`
	ValidateRegionInfo      *bool                        `json:"validateRegionInfo,omitempty"`
}

// GET endpoints can return numeric enum values while the write API uses names.
func (m *MsgS3CloudAppsInstance) UnmarshalJSON(data []byte) error {
	type alias MsgS3CloudAppsInstance
	var response struct {
		*alias
		InstanceType   json.RawMessage `json:"instanceType"`
		CredentialType json.RawMessage `json:"credentialType"`
	}
	response.alias = (*alias)(m)
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}
	var err error
	m.InstanceType, err = cloudNamedEnum(response.InstanceType, map[string]string{"5": "AMAZON_S3"})
	if err != nil {
		return fmt.Errorf("cloudAppsInstance.instanceType: %w", err)
	}
	m.CredentialType, err = cloudNamedEnum(response.CredentialType, map[string]string{"102": "AMAZON_S3"})
	if err != nil {
		return fmt.Errorf("cloudAppsInstance.credentialType: %w", err)
	}
	return nil
}

func cloudNamedEnum(value json.RawMessage, numbers map[string]string) (*string, error) {
	value = bytes.TrimSpace(value)
	if len(value) == 0 || bytes.Equal(value, []byte("null")) {
		return nil, nil
	}
	if value[0] == '"' {
		var name string
		if err := json.Unmarshal(value, &name); err != nil {
			return nil, err
		}
		return &name, nil
	}
	if name, ok := numbers[string(value)]; ok {
		return &name, nil
	}
	return nil, fmt.Errorf("unsupported numeric enum value %s", value)
}

type MsgS3InstanceProperties struct {
	Instance            *MsgS3Entity            `json:"instance,omitempty"`
	CloudAppsInstance   *MsgS3CloudAppsInstance `json:"cloudAppsInstance,omitempty"`
	PlanEntity          *MsgS3Plan              `json:"planEntity,omitempty"`
	UseResourcePoolInfo *bool                   `json:"useResourcePoolInfo,omitempty"`
}

type MsgS3CloudConnector struct {
	InstanceType *string                  `json:"instanceType,omitempty"`
	Instance     *MsgS3InstanceProperties `json:"instance,omitempty"`
}

type MsgS3ClientInfo struct {
	ClientType                *int                 `json:"clientType,omitempty"`
	CloudClonnectorProperties *MsgS3CloudConnector `json:"cloudClonnectorProperties,omitempty"`
	Plan                      *MsgS3Plan           `json:"plan,omitempty"`
}

type MsgCreateS3ObjectStorageRequest struct {
	ClientInfo *MsgS3ClientInfo `json:"clientInfo,omitempty"`
	Entity     *MsgS3Entity     `json:"entity,omitempty"`
}

type MsgS3Association struct {
	Entity []MsgS3Entity `json:"entity,omitempty"`
}

type MsgModifyS3ObjectStorageRequest struct {
	NewName            *string                  `json:"newName,omitempty"`
	InstanceProperties *MsgS3InstanceProperties `json:"instanceProperties,omitempty"`
	Association        *MsgS3Association        `json:"association,omitempty"`
}

type MsgDeleteS3ObjectStorageRequest struct {
	Instance *MsgS3Entity `json:"instance,omitempty"`
}

type MsgS3OperationResult struct {
	ErrorCode      *int         `json:"errorCode,omitempty"`
	ErrorMessage   *string      `json:"errorMessage,omitempty"`
	ErrorString    *string      `json:"errorString,omitempty"`
	WarningCode    *int         `json:"warningCode,omitempty"`
	WarningMessage *string      `json:"warningMessage,omitempty"`
	Entity         *MsgS3Entity `json:"entity,omitempty"`
}

type MsgS3OperationResponse struct {
	ProcessingInstructionInfo *MsgProcessingInstructionInfo `json:"processinginstructioninfo,omitempty"`
	Response                  json.RawMessage               `json:"response,omitempty"`
	ErrorCode                 *int                          `json:"errorCode,omitempty"`
	ErrorMessage              *string                       `json:"errorMessage,omitempty"`
	ErrorString               *string                       `json:"errorString,omitempty"`
}

type MsgGetS3ObjectStorageResponse struct {
	InstanceProperties json.RawMessage `json:"instanceProperties,omitempty"`
	ErrorCode          *int            `json:"errorCode,omitempty"`
	ErrorMessage       *string         `json:"errorMessage,omitempty"`
}

type MsgS3Content struct {
	Path        *string `json:"path,omitempty"`
	ExcludePath *string `json:"excludePath,omitempty"`
	IncludePath *string `json:"includePath,omitempty"`
}

type MsgS3TagRule struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type MsgS3TagFilter struct {
	RuleItemType *int           `json:"ruleItemType,omitempty"`
	Rules        []MsgS3TagRule `json:"rules,omitempty"`
}

type MsgS3TagFilters struct {
	TagFilters []MsgS3TagFilter `json:"tagFilters"`
}

type MsgS3ObjectStorageSubclient struct {
	IsEventScanEnabled     *bool                     `json:"isEventScanEnabled,omitempty"`
	EventScanProperties    *MsgS3EventScanProperties `json:"eventScanProperties,omitempty"`
	BackupContentType      *int                      `json:"backupContentType,omitempty"`
	ObjectStorageTagFilter *MsgS3TagFilters          `json:"objectStorageTagFilter,omitempty"`
}

type MsgS3EventScanProperties struct {
	CloudStagingPath *string `json:"cloudStagingPath,omitempty"`
}

type MsgS3CloudAppsSubclient struct {
	InstanceType           *string                      `json:"instanceType,omitempty"`
	CollectACL             *bool                        `json:"collectACL,omitempty"`
	ObjectStorageSubclient *MsgS3ObjectStorageSubclient `json:"objectStorageSubclient,omitempty"`
}

func (m *MsgS3CloudAppsSubclient) UnmarshalJSON(data []byte) error {
	type alias MsgS3CloudAppsSubclient
	var response struct {
		*alias
		InstanceType json.RawMessage `json:"instanceType"`
	}
	response.alias = (*alias)(m)
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}
	var err error
	m.InstanceType, err = cloudNamedEnum(response.InstanceType, map[string]string{"5": "AMAZON_S3"})
	if err != nil {
		return fmt.Errorf("cloudAppsSubClientProp.instanceType: %w", err)
	}
	return nil
}

type MsgS3CommonProperties struct {
	NumberOfBackupStreams *int `json:"numberOfBackupStreams,omitempty"`
	JobStartTime          *int `json:"jobStartTime,omitempty"`
}

type MsgS3SubclientProperties struct {
	SubClientEntity              *MsgS3Entity             `json:"subClientEntity,omitempty"`
	PlanEntity                   *MsgS3Plan               `json:"planEntity,omitempty"`
	CloudAppsSubClientProp       *MsgS3CloudAppsSubclient `json:"cloudAppsSubClientProp,omitempty"`
	CommonProperties             *MsgS3CommonProperties   `json:"commonProperties,omitempty"`
	Content                      *[]MsgS3Content          `json:"content,omitempty"`
	FsContentOperationType       *string                  `json:"fsContentOperationType,omitempty"`
	FsExcludeFilterOperationType *string                  `json:"fsExcludeFilterOperationType,omitempty"`
	FsIncludeFilterOperationType *string                  `json:"fsIncludeFilterOperationType,omitempty"`
}

type MsgS3SubclientRequest struct {
	NewName             *string                   `json:"newName,omitempty"`
	SubClientProperties *MsgS3SubclientProperties `json:"subClientProperties,omitempty"`
}

type MsgGetS3SubclientResponse struct {
	SubClientProperties json.RawMessage `json:"subClientProperties,omitempty"`
	ErrorCode           *int            `json:"errorCode,omitempty"`
	ErrorMessage        *string         `json:"errorMessage,omitempty"`
}

type MsgS3AccessNode struct {
	Entity *MsgS3Entity `json:"entity,omitempty"`
}

type MsgGetS3AccessNodesResponse struct {
	AccessNodes  []MsgS3AccessNode `json:"accessNodes,omitempty"`
	ErrorCode    *int              `json:"errorCode,omitempty"`
	ErrorMessage *string           `json:"errorMessage,omitempty"`
}
