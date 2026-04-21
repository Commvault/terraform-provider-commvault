package handler

import "encoding/json"

// Request/Response types for Oracle Instance operations

// MsgCreateOracleInstanceRequest represents the request to add an Oracle instance
type MsgCreateOracleInstanceRequest struct {
	InstanceProperties *MsgOracleInstanceProperties `json:"instanceProperties,omitempty"`
}

// MsgOracleInstanceProperties represents Oracle instance properties
type MsgOracleInstanceProperties struct {
	Instance       *MsgOracleInstance        `json:"instance,omitempty"`
	OracleInstance *MsgOracleInstanceDetails `json:"oracleInstance,omitempty"`
	PlanEntity     *MsgOraclePlanEntity      `json:"planEntity,omitempty"`
}

// MsgOraclePlanEntity supports both `id` and `planId` forms used by Commvault APIs.
type MsgOraclePlanEntity struct {
	Name   *string `json:"name,omitempty"`
	Id     *int    `json:"id,omitempty"`
	PlanID *int    `json:"planId,omitempty"`
}

// MsgOracleInstance represents Oracle instance details
type MsgOracleInstance struct {
	ClientName    *string `json:"clientName,omitempty"`
	ClientId      *int    `json:"clientId,omitempty"`
	InstanceName  *string `json:"instanceName,omitempty"`
	AppName       *string `json:"appName,omitempty"`
	ApplicationId *int    `json:"applicationId,omitempty"`
	CommCellId    *int    `json:"commCellId,omitempty"`
}

// MsgCreateOracleInstanceResponse represents the response from creating an Oracle instance
type MsgCreateOracleInstanceResponse struct {
	ProcessingInstructionInfo *MsgProcessingInstructionInfo `json:"processinginstructioninfo,omitempty"`
	Response                  json.RawMessage               `json:"response,omitempty"`
	ErrorCode                 *int                          `json:"errorCode,omitempty"`
	ErrorMessage              *string                       `json:"errorMessage,omitempty"`
	ErrorString               *string                       `json:"errorString,omitempty"`
	Entity                    *MsgInstanceEntity            `json:"entity,omitempty"`
}

// MsgCreateResponseItem is one element inside the "response" array returned by POST /instance
type MsgCreateResponseItem struct {
	ErrorCode   *int               `json:"errorCode,omitempty"`
	ErrorString *string            `json:"errorString,omitempty"`
	Entity      *MsgInstanceEntity `json:"entity,omitempty"`
}

// MsgInstanceEntity holds the entity identifiers returned after creating an instance
type MsgInstanceEntity struct {
	InstanceId    *int    `json:"instanceId,omitempty"`
	ClientId      *int    `json:"clientId,omitempty"`
	InstanceName  *string `json:"instanceName,omitempty"`
	ClientName    *string `json:"clientName,omitempty"`
	ApplicationId *int    `json:"applicationId,omitempty"`
}

// MsgProcessingInstructionInfo represents processing instruction info
type MsgProcessingInstructionInfo struct {
	Attributes []MsgAttribute `json:"attributes,omitempty"`
}

// MsgAttribute represents a generic attribute
type MsgAttribute struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// MsgGenericResponse represents a generic response
type MsgGenericResponse struct {
	ErrorCode   *int       `json:"errorCode,omitempty"`
	ErrorString *string    `json:"errorString,omitempty"`
	Entity      *MsgIdName `json:"entity,omitempty"`
}

// MsgFetchOracleInstancesResponse represents the response from fetching Oracle instances
type MsgFetchOracleInstancesResponse struct {
	InstanceProperties []MsgOracleInstancePropertiesResp `json:"instanceProperties,omitempty"`
}

// MsgOracleInstancePropertiesResp represents Oracle instance properties in response
type MsgOracleInstancePropertiesResp struct {
	Instance *MsgOracleInstanceResp `json:"instance,omitempty"`
}

// MsgOracleInstanceResp represents Oracle instance response details
type MsgOracleInstanceResp struct {
	ClientId      *int    `json:"clientId,omitempty"`
	ClientName    *string `json:"clientName,omitempty"`
	InstanceId    *int    `json:"instanceId,omitempty"`
	InstanceName  *string `json:"instanceName,omitempty"`
	ApplicationId *int    `json:"applicationId,omitempty"`
	AppName       *string `json:"appName,omitempty"`
}

// MsgGetOracleInstancePropertiesResponse represents the response from getting instance properties
type MsgGetOracleInstancePropertiesResponse struct {
	InstanceProperties json.RawMessage `json:"instanceProperties,omitempty"`
}

// MsgOracleInstanceFullProperties represents full Oracle instance properties
type MsgOracleInstanceFullProperties struct {
	Instance       *MsgOracleInstanceResp    `json:"instance,omitempty"`
	OracleInstance *MsgOracleInstanceDetails `json:"oracleInstance,omitempty"`
	PlanEntity     *MsgOraclePlanEntity      `json:"planEntity,omitempty"`
}

// MsgOracleInstanceDetails represents Oracle-specific instance details
type MsgOracleInstanceDetails struct {
	OracleHome                 *string           `json:"oracleHome,omitempty"`
	OracleUser                 *MsgOracleUser    `json:"oracleUser,omitempty"`
	OracleWalletAuthentication *bool             `json:"oracleWalletAuthentication,omitempty"`
	SqlConnect                 *MsgOracleConnect `json:"sqlConnect,omitempty"`
	UseCatalogConnect          *bool             `json:"useCatalogConnect,omitempty"`
	CatalogConnect             *MsgOracleConnect `json:"catalogConnect,omitempty"`
	BlockSize                  *int              `json:"blockSize,omitempty"`
	CrossCheckTimeout          *int              `json:"crossCheckTimeout,omitempty"`
	TNSAdminPath               *string           `json:"TNSAdminPath,omitempty"`
	ArchiveLogDest             *string           `json:"archiveLogDest,omitempty"`
	// Credential-based authentication (preferred over inline credentials)
	OsUserCredInfo         *MsgCredentialInfo `json:"osUserCredInfo,omitempty"`
	DbConnectCredInfo      *MsgCredentialInfo `json:"dbConnectCredInfo,omitempty"`
	CatalogConnectCredInfo *MsgCredentialInfo `json:"catalogConnectCredInfo,omitempty"`
	// Storage configuration
	OracleStorageDevice *MsgOracleStorageDevice `json:"oracleStorageDevice,omitempty"`
}

// MsgOracleUser represents Oracle OS user
type MsgOracleUser struct {
	UserName   *string `json:"userName,omitempty"`
	DomainName *string `json:"domainName,omitempty"`
}

// MsgOracleConnect represents Oracle connection credentials (deprecated - use credential IDs instead)
type MsgOracleConnect struct {
	UserName        *string `json:"userName,omitempty"`
	Password        *string `json:"password,omitempty"`
	DomainName      *string `json:"domainName,omitempty"`
	ConfirmPassword *string `json:"confirmPassword,omitempty"`
}

// MsgCredentialInfo represents a reference to a stored credential
type MsgCredentialInfo struct {
	CredentialId   *int    `json:"credentialId,omitempty"`
	CredentialName *string `json:"credentialName,omitempty"`
}

// MsgOracleStorageDevice represents Oracle storage device configuration
type MsgOracleStorageDevice struct {
	LogBackupStoragePolicy   *MsgIdName `json:"logBackupStoragePolicy,omitempty"`
	CommandLineStoragePolicy *MsgIdName `json:"commandLineStoragePolicy,omitempty"`
	NetworkAgents            *int       `json:"networkAgents,omitempty"`
	SoftwareCompression      *int       `json:"softwareCompression,omitempty"`
	ThrottleNetworkBandwidth *int       `json:"throttleNetworkBandwidth,omitempty"`
}

// MsgModifyOracleInstanceRequest represents the request to modify an Oracle instance
type MsgModifyOracleInstanceRequest struct {
	InstanceProperties *MsgOracleInstanceFullProperties `json:"instanceProperties,omitempty"`
}

// MsgModifyOracleInstanceResponse represents the response from modifying an Oracle instance
type MsgModifyOracleInstanceResponse struct {
	ProcessingInstructionInfo *MsgProcessingInstructionInfo `json:"processinginstructioninfo,omitempty"`
	Response                  *MsgGenericResponse           `json:"response,omitempty"`
}

// MsgDeleteOracleInstanceResponse represents the response from deleting an Oracle instance
type MsgDeleteOracleInstanceResponse struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
	ErrorCode    *int    `json:"errorCode,omitempty"`
}

// MsgOracleInstanceDiscoveryResponse represents the response from instance discovery
type MsgOracleInstanceDiscoveryResponse struct {
	Error *MsgOracleDiscoveryError `json:"error,omitempty"`
}

// MsgOracleDiscoveryError represents discovery error details
type MsgOracleDiscoveryError struct {
	ErrorCode   *int    `json:"errorCode,omitempty"`
	ErrorString *string `json:"errorString,omitempty"`
}

// MsgGetOracleBackupPiecesResponse represents the response from getting backup pieces
type MsgGetOracleBackupPiecesResponse struct {
	BackupPieces []MsgOracleBackupPiece `json:"backupPieces,omitempty"`
}

// MsgOracleBackupPiece represents an Oracle backup piece
type MsgOracleBackupPiece struct {
	BackupPieceName *string `json:"backupPieceName,omitempty"`
	Tag             *string `json:"tag,omitempty"`
	StartTime       *int    `json:"startTime,omitempty"`
	EndTime         *int    `json:"endTime,omitempty"`
	BackupType      *string `json:"backupType,omitempty"`
	Size            *int64  `json:"size,omitempty"`
}

// MsgBrowseOracleDBRequest represents the request to browse Oracle database
type MsgBrowseOracleDBRequest struct {
	Path        *string    `json:"path,omitempty"`
	PointInTime *int       `json:"pointInTime,omitempty"`
	Entity      *MsgIdName `json:"entity,omitempty"`
}

// MsgBrowseOracleDBResponse represents the response from browsing Oracle database
type MsgBrowseOracleDBResponse struct {
	OracleContent []MsgOracleContent `json:"oracleContent,omitempty"`
}

// MsgOracleContent represents Oracle content from browse
type MsgOracleContent struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Size *int64  `json:"size,omitempty"`
}

// Request/Response types for Oracle Subclient operations

// MsgCreateOracleSubclientRequest represents the request to create an Oracle subclient
type MsgCreateOracleSubclientRequest struct {
	SubClientProperties *MsgOracleSubclientProperties `json:"subClientProperties,omitempty"`
}

// MsgOracleSubclientProperties represents Oracle subclient properties
type MsgOracleSubclientProperties struct {
	SubClientEntity      *MsgOracleSubclientEntity     `json:"subClientEntity,omitempty"`
	ContentOperationType *string                       `json:"contentOperationType,omitempty"`
	CommonProperties     *MsgSubclientCommonProperties `json:"commonProperties,omitempty"`
	OracleSubclientProp  *MsgOracleSubclientDetails    `json:"oracleSubclientProp,omitempty"`
	Content              []json.RawMessage             `json:"content,omitempty"`
	PlanEntity           *MsgOraclePlanEntity          `json:"planEntity,omitempty"`
}

// MsgOracleSubclientEntity represents Oracle subclient entity
type MsgOracleSubclientEntity struct {
	SubclientId   *int                 `json:"subclientId,omitempty"`
	SubclientName *string              `json:"subclientName,omitempty"`
	ClientName    *string              `json:"clientName,omitempty"`
	InstanceName  *string              `json:"instanceName,omitempty"`
	AppName       *string              `json:"appName,omitempty"`
	ApplicationId *int                 `json:"applicationId,omitempty"`
	BackupsetId   *int                 `json:"backupsetId,omitempty"`
	InstanceId    *int                 `json:"instanceId,omitempty"`
	ClientId      *int                 `json:"clientId,omitempty"`
	EntityInfo    *MsgOracleEntityInfo `json:"entityInfo,omitempty"`
}

type MsgOracleEntityInfo struct {
	CompanyID       *int `json:"companyId,omitempty"`
	MultiCommcellID *int `json:"multiCommcellId,omitempty"`
}

// MsgSubclientCommonProperties represents common subclient properties
type MsgSubclientCommonProperties struct {
	StorageDevice *MsgStorageDevice `json:"storageDevice,omitempty"`
	EnableBackup  *bool             `json:"enableBackup,omitempty"`
	Description   *string           `json:"description,omitempty"`
}

// MsgStorageDevice represents storage device configuration
type MsgStorageDevice struct {
	DataBackupStoragePolicy *MsgIdName `json:"dataBackupStoragePolicy,omitempty"`
	LogBackupStoragePolicy  *MsgIdName `json:"logBackupStoragePolicy,omitempty"`
}

// MsgCreateOracleSubclientResponse represents the response from creating an Oracle subclient
type MsgCreateOracleSubclientResponse struct {
	ProcessingInstructionInfo *MsgProcessingInstructionInfo `json:"processinginstructioninfo,omitempty"`
	Response                  *MsgGenericResponse           `json:"response,omitempty"`
}

// MsgFetchOracleSubclientsResponse represents the response from fetching Oracle subclients
type MsgFetchOracleSubclientsResponse struct {
	FilterQueryCount    *int                           `json:"filterQueryCount,omitempty"`
	SubClientProperties []MsgOracleSubclientProperties `json:"subClientProperties,omitempty"`
}

// MsgGetOracleSubclientPropertiesResponse represents the response from getting subclient properties
type MsgGetOracleSubclientPropertiesResponse struct {
	SubClientProperties json.RawMessage `json:"subClientProperties,omitempty"`
}

// MsgOracleSubclientFullProperties represents full Oracle subclient properties
type MsgOracleSubclientFullProperties struct {
	SubClientEntity      *MsgOracleSubclientEntity     `json:"subClientEntity,omitempty"`
	CommonProperties     *MsgSubclientCommonProperties `json:"commonProperties,omitempty"`
	OracleSubclientProp  *MsgOracleSubclientDetails    `json:"oracleSubclientProp,omitempty"`
	Content              []json.RawMessage             `json:"content,omitempty"`
	ContentOperationType *string                       `json:"contentOperationType,omitempty"`
	PlanEntity           *MsgOraclePlanEntity          `json:"planEntity,omitempty"`
}

// MsgOracleSubclientDetails represents Oracle-specific subclient details
type MsgOracleSubclientDetails struct {
	BackupMode                  *int    `json:"backupMode,omitempty"`
	BackupArchiveLog            *bool   `json:"backupArchiveLog,omitempty"`
	BackupSPFile                *bool   `json:"backupSPFile,omitempty"`
	BackupControlFile           *bool   `json:"backupControlFile,omitempty"`
	DeleteArchiveLogAfterBackup *bool   `json:"deleteArchiveLogAfterBackup,omitempty"`
	SelectiveOnlineFull         *bool   `json:"selectiveOnlineFull,omitempty"`
	Data                        *bool   `json:"data,omitempty"`
	LightsOutScript             *string `json:"lightsOutScript,omitempty"`
	EnableTableBrowse           *bool   `json:"enableTableBrowse,omitempty"`
	ArchiveDelete               *bool   `json:"archiveDelete,omitempty"`
	DataThresholdStreams        *int    `json:"dataThresholdStreams,omitempty"`
	ArchiveDeleteAll            *bool   `json:"archiveDeleteAll,omitempty"`
}

// MsgModifyOracleSubclientRequest represents the request to modify an Oracle subclient
type MsgModifyOracleSubclientRequest struct {
	SubClientProperties *MsgOracleSubclientFullProperties `json:"subClientProperties,omitempty"`
}

// MsgModifyOracleSubclientResponse represents the response from modifying an Oracle subclient
type MsgModifyOracleSubclientResponse struct {
	ProcessingInstructionInfo *MsgProcessingInstructionInfo `json:"processinginstructioninfo,omitempty"`
}

// MsgDeleteOracleSubclientResponse represents the response from deleting an Oracle subclient
type MsgDeleteOracleSubclientResponse struct {
	ProcessingInstructionInfo *MsgProcessingInstructionInfo `json:"processinginstructioninfo,omitempty"`
	ErrorMessage              *string                       `json:"errorMessage,omitempty"`
	ErrorCode                 *int                          `json:"errorCode,omitempty"`
}

// Request/Response types for Oracle Backup/Restore operations

// MsgOracleBackupRequest represents the request to trigger an Oracle backup
type MsgOracleBackupRequest struct {
	TaskInfo *MsgOracleTaskInfo `json:"taskInfo,omitempty"`
}

// MsgOracleTaskInfo represents Oracle task info
type MsgOracleTaskInfo struct {
	Associations []MsgOracleAssociation `json:"associations,omitempty"`
	Task         *MsgOracleTask         `json:"task,omitempty"`
	SubTasks     []MsgOracleSubTask     `json:"subTasks,omitempty"`
}

// MsgOracleAssociation represents Oracle association
type MsgOracleAssociation struct {
	ClientName    *string `json:"clientName,omitempty"`
	ClientId      *int    `json:"clientId,omitempty"`
	InstanceName  *string `json:"instanceName,omitempty"`
	InstanceId    *int    `json:"instanceId,omitempty"`
	SubclientName *string `json:"subclientName,omitempty"`
	SubclientId   *int    `json:"subclientId,omitempty"`
	AppName       *string `json:"appName,omitempty"`
	ApplicationId *int    `json:"applicationId,omitempty"`
}

// MsgOracleTask represents Oracle task
type MsgOracleTask struct {
	TaskType      *int                `json:"taskType,omitempty"`
	InitiatedFrom *int                `json:"initiatedFrom,omitempty"`
	PurityType    *int                `json:"purityType,omitempty"`
	TaskFlags     *MsgOracleTaskFlags `json:"taskFlags,omitempty"`
}

// MsgOracleTaskFlags represents Oracle task flags
type MsgOracleTaskFlags struct {
	Disabled *bool `json:"disabled,omitempty"`
}

// MsgOracleSubTask represents Oracle subtask
type MsgOracleSubTask struct {
	SubTaskType   *int                     `json:"subTaskType,omitempty"`
	OperationType *int                     `json:"operationType,omitempty"`
	Options       *MsgOracleSubTaskOptions `json:"options,omitempty"`
}

// MsgOracleSubTaskOptions represents Oracle subtask options
type MsgOracleSubTaskOptions struct {
	BackupOpts  *MsgOracleBackupOptions  `json:"backupOpts,omitempty"`
	RestoreOpts *MsgOracleRestoreOptions `json:"restoreOpts,omitempty"`
}

// MsgOracleBackupOptions represents Oracle backup options
type MsgOracleBackupOptions struct {
	BackupLevel     *int  `json:"backupLevel,omitempty"`
	IncludeDataFile *bool `json:"includeDataFile,omitempty"`
	RunIncrBackup   *bool `json:"runIncrBackup,omitempty"`
}

// MsgOracleRestoreOptions represents Oracle restore options
type MsgOracleRestoreOptions struct {
	RestoreType         *int    `json:"restoreType,omitempty"`
	DestinationClient   *string `json:"destinationClient,omitempty"`
	DestinationInstance *string `json:"destinationInstance,omitempty"`
	PointInTime         *int    `json:"pointInTime,omitempty"`
	SCN                 *string `json:"scn,omitempty"`
}

// MsgOracleBackupResponse represents the response from Oracle backup
type MsgOracleBackupResponse struct {
	TaskId *int  `json:"taskId,omitempty"`
	JobIds []int `json:"jobIds,omitempty"`
}

// MsgOracleRestoreResponse represents the response from Oracle restore
type MsgOracleRestoreResponse struct {
	TaskId *int  `json:"taskId,omitempty"`
	JobIds []int `json:"jobIds,omitempty"`
}

// MsgFetchRMANLogsResponse represents the response from fetching RMAN logs
type MsgFetchRMANLogsResponse struct {
	ErrorString *string `json:"errorString,omitempty"`
	LogContent  *string `json:"logContent,omitempty"`
}

// MsgFetchOracleEntityIdResponse represents the response from fetching Oracle entity ID.
// Commvault returns -32000 for fields it cannot resolve, so we use plain int (not pointer)
// so callers can check for the sentinel value with id > 0.
type MsgFetchOracleEntityIdResponse struct {
	ClientId    int `json:"clientId"`
	InstanceId  int `json:"instanceId"`
	SubclientId int `json:"subclientId"`
}

// MsgInstallOracleAgentRequest represents the request to install Oracle agent
type MsgInstallOracleAgentRequest struct {
	TaskInfo *MsgInstallOracleAgentTaskInfo `json:"taskInfo,omitempty"`
}

// MsgInstallOracleAgentTaskInfo represents install agent task info
type MsgInstallOracleAgentTaskInfo struct {
	Task     *MsgInstallOracleAgentTask     `json:"task,omitempty"`
	SubTasks []MsgInstallOracleAgentSubTask `json:"subTasks,omitempty"`
}

// MsgInstallOracleAgentTask represents install agent task
type MsgInstallOracleAgentTask struct {
	TaskType *int `json:"taskType,omitempty"`
}

// MsgInstallOracleAgentSubTask represents install agent subtask
type MsgInstallOracleAgentSubTask struct {
	SubTask *MsgInstallOracleAgentSubTaskInfo `json:"subTask,omitempty"`
	Options *MsgInstallOracleAgentOptions     `json:"options,omitempty"`
}

// MsgInstallOracleAgentSubTaskInfo represents subtask info
type MsgInstallOracleAgentSubTaskInfo struct {
	SubTaskType   *string `json:"subTaskType,omitempty"`
	OperationType *string `json:"operationType,omitempty"`
}

// MsgInstallOracleAgentOptions represents install options
type MsgInstallOracleAgentOptions struct {
	AdminOpts *MsgInstallOracleAgentAdminOpts `json:"adminOpts,omitempty"`
}

// MsgInstallOracleAgentAdminOpts represents admin options
type MsgInstallOracleAgentAdminOpts struct {
	ClientInstallOption *MsgClientInstallOption `json:"clientInstallOption,omitempty"`
	UpdateOption        *MsgUpdateOption        `json:"updateOption,omitempty"`
}

// MsgClientInstallOption represents client install option
type MsgClientInstallOption struct {
	InstallerOption  *MsgInstallerOption  `json:"installerOption,omitempty"`
	ClientDetails    []MsgClientDetails   `json:"clientDetails,omitempty"`
	InstallOSType    *int                 `json:"installOSType,omitempty"`
	DiscoveryType    *int                 `json:"discoveryType,omitempty"`
	ClientAuthForJob *MsgClientAuthForJob `json:"clientAuthForJob,omitempty"`
}

// MsgInstallerOption represents installer options
type MsgInstallerOption struct {
	User              *MsgInstallerUser          `json:"User,omitempty"`
	CommServeHostName *string                    `json:"CommServeHostName,omitempty"`
	InstallFlags      *MsgInstallFlags           `json:"installFlags,omitempty"`
	ClientComposition []MsgClientCompositionItem `json:"clientComposition,omitempty"`
}

// MsgInstallerUser represents installer user
type MsgInstallerUser struct {
	UserId   *int    `json:"userId,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

// MsgInstallFlags represents install flags
type MsgInstallFlags struct {
	AllowMultipleInstances          *bool               `json:"allowMultipleInstances,omitempty"`
	Install32Base                   *bool               `json:"install32Base,omitempty"`
	DisableOSFirewall               *bool               `json:"disableOSFirewall,omitempty"`
	AddToFirewallExclusion          *bool               `json:"addToFirewallExclusion,omitempty"`
	ForceReboot                     *bool               `json:"forceReboot,omitempty"`
	KillBrowserProcesses            *bool               `json:"killBrowserProcesses,omitempty"`
	IgnoreJobsRunning               *bool               `json:"ignoreJobsRunning,omitempty"`
	StopOracleServices              *bool               `json:"stopOracleServices,omitempty"`
	SkipClientsOfCS                 *bool               `json:"skipClientsOfCS,omitempty"`
	RestoreOnlyAgents               *bool               `json:"restoreOnlyAgents,omitempty"`
	OverrideClientInfo              *bool               `json:"overrideClientInfo,omitempty"`
	OverrideUnixGroupAndPermissions *bool               `json:"overrideUnixGroupAndPermissions,omitempty"`
	UnixGroup                       *string             `json:"unixGroup,omitempty"`
	UnixGroupAccess                 *int                `json:"unixGroupAccess,omitempty"`
	UnixOtherAccess                 *int                `json:"unixOtherAccess,omitempty"`
	FirewallInstall                 *MsgFirewallInstall `json:"firewallInstall,omitempty"`
}

// MsgFirewallInstall represents firewall install settings
type MsgFirewallInstall struct {
	EnableFirewallConfig *bool `json:"enableFirewallConfig,omitempty"`
	PortNumber           *int  `json:"portNumber,omitempty"`
}

// MsgClientCompositionItem represents client composition item
type MsgClientCompositionItem struct {
	Components *MsgComponentsInfo `json:"components,omitempty"`
}

// MsgComponentsInfo represents components info
type MsgComponentsInfo struct {
	ComponentInfo []MsgComponentInfoItem `json:"componentInfo,omitempty"`
}

// MsgComponentInfoItem represents component info item
type MsgComponentInfoItem struct {
	ComponentName *string `json:"componentName,omitempty"`
	OSType        *string `json:"osType,omitempty"`
}

// MsgClientDetails represents client details
type MsgClientDetails struct {
	ClientEntity *MsgClientEntity `json:"clientEntity,omitempty"`
}

// MsgClientEntity represents client entity
type MsgClientEntity struct {
	ClientName  *string `json:"clientName,omitempty"`
	HostName    *string `json:"hostName,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
}

// MsgClientAuthForJob represents client auth for job
type MsgClientAuthForJob struct {
	UserName *string `json:"userName,omitempty"`
	Password *string `json:"password,omitempty"`
}

// MsgUpdateOption represents update option
type MsgUpdateOption struct {
	Plan *MsgPlanRef `json:"plan,omitempty"`
}

// MsgPlanRef represents plan reference
type MsgPlanRef struct {
	PlanId *int `json:"planId,omitempty"`
}

// MsgInstallOracleAgentResponse represents the response from installing Oracle agent
type MsgInstallOracleAgentResponse struct {
	TaskId *int     `json:"taskId,omitempty"`
	JobIds []string `json:"jobIds,omitempty"`
}
