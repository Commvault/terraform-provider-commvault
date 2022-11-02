package handler

type MsgCreateUserRequest struct {
    Users[]   MsgCreateUserSet `json:"users,omitempty"`
}

type MsgCreateUserSet struct {
    Password                      *string `json:"password,omitempty"`
    Name                          *string `json:"name,omitempty"`
    FullName                      *string `json:"fullName,omitempty"`
    Company                       *MsgIdName `json:"company,omitempty"`
    UseSystemGeneratePassword     *bool `json:"useSystemGeneratePassword,omitempty"`
    InviteUser                    *bool `json:"inviteUser,omitempty"`
    Plan                          *MsgIdName `json:"plan,omitempty"`
    Email                         *string `json:"email,omitempty"`
}

type MsgIdName struct {
    Name     *string `json:"name,omitempty"`
    Id       *int `json:"id,omitempty"`
}

type MsgCreateUserResponse struct {
    Users[]   MsgIdNameGUIDSet `json:"users,omitempty"`
}

type MsgIdNameGUIDSet struct {
    GUID     *string `json:"GUID,omitempty"`
    Name     *string `json:"name,omitempty"`
    Id       *int `json:"id,omitempty"`
}

type MsgGetUserDetailsResponse struct {
    ServiceType              *string `json:"serviceType,omitempty"`
    GUID                     *string `json:"GUID,omitempty"`
    FullName                 *string `json:"fullName,omitempty"`
    Description              *string `json:"description,omitempty"`
    AssociatedUserGroups[]   MsgIdNameProviderSet `json:"associatedUserGroups,omitempty"`
    Enabled                  *bool `json:"enabled,omitempty"`
    LockInfo                 *MsgLockProperties `json:"lockInfo,omitempty"`
    AuthenticationMethod     *string `json:"authenticationMethod,omitempty"`
    Name                     *string `json:"name,omitempty"`
    LastLoggedIn             *int64 `json:"lastLoggedIn,omitempty"`
    Company                  *MsgIdName `json:"company,omitempty"`
    Id                       *int `json:"id,omitempty"`
    Plan                     *MsgIdName `json:"plan,omitempty"`
    Email                    *string `json:"email,omitempty"`
    UserPrincipalName        *string `json:"userPrincipalName,omitempty"`
}

type MsgIdNameProviderSet struct {
    Provider     *MsgIdName `json:"provider,omitempty"`
    Name         *string `json:"name,omitempty"`
    Id           *int `json:"id,omitempty"`
}

type MsgLockProperties struct {
    IsLocked      *bool `json:"isLocked,omitempty"`
    StartTime     *int `json:"startTime,omitempty"`
    EndTime       *int `json:"endTime,omitempty"`
}

type MsgDeleteUserResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgModifyUserRequest struct {
    NewName                *string `json:"newName,omitempty"`
    FullName               *string `json:"fullName,omitempty"`
    NewPassword            *string `json:"newPassword,omitempty"`
    Plan                   *MsgIdName `json:"plan,omitempty"`
    Email                  *string `json:"email,omitempty"`
    Enabled                *bool `json:"enabled,omitempty"`
    UserPrincipalName      *string `json:"userPrincipalName,omitempty"`
    ValidationPassword     *string `json:"validationPassword,omitempty"`
}

type MsgModifyUserResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateHypervisorAWSRequest struct {
    SkipCredentialValidation     *bool `json:"skipCredentialValidation,omitempty"`
    EtcdProtection               *bool `json:"etcdProtection,omitempty"`
    Credentials                  *MsgIdName `json:"credentials,omitempty"`
    Name                         *string `json:"name,omitempty"`
    AccessNodes[]                MsgIdNameSet `json:"accessNodes,omitempty"`
    PlanEntity                   *MsgIdName `json:"planEntity,omitempty"`
    SecretKey                    *string `json:"secretKey,omitempty"`
    AccessKey                    *string `json:"accessKey,omitempty"`
    Region                       *string `json:"Region,omitempty"`
    HypervisorType               *string `json:"hypervisorType,omitempty"`
    UseServiceAccount            *string `json:"useServiceAccount,omitempty"`
    UseIamRole                   *bool `json:"useIamRole,omitempty"`
    RoleARN                      *string `json:"RoleARN,omitempty"`
}

type MsgIdNameSet struct {
    Id     *int `json:"id,omitempty"`
}

type MsgCreateHypervisorAWSResponse struct {
    Response     *MsgCreateHypervisorResp `json:"response,omitempty"`
}

type MsgCreateHypervisorResp struct {
    WarningCode        *int `json:"warningCode,omitempty"`
    ErrorMessage       *string `json:"errorMessage,omitempty"`
    ErrorCode          *int `json:"errorCode,omitempty"`
    WarningMessage     *string `json:"warningMessage,omitempty"`
    HypervisorId       *int `json:"hypervisorId,omitempty"`
}

type MsgupdateHypervisorAWSRequest struct {
    ActivityControl              *MsgActivityControlOptions `json:"activityControl,omitempty"`
    Settings                     *MsghypervisorSettings `json:"settings,omitempty"`
    Security                     *MsgVMHypervisorSecurityProp `json:"security,omitempty"`
    NewName                      *string `json:"newName,omitempty"`
    SkipCredentialValidation     *bool `json:"skipCredentialValidation,omitempty"`
    AccessNode[]                 MsgaccessNodeModelSet `json:"accessNode,omitempty"`
    FbrUnixMediaAgent            *MsgIdName `json:"fbrUnixMediaAgent,omitempty"`
    SecretKey                    *string `json:"secretKey,omitempty"`
    AccessKey                    *string `json:"accessKey,omitempty"`
    Region                       *string `json:"Region,omitempty"`
    HypervisorType               *string `json:"hypervisorType,omitempty"`
    UseServiceAccount            *string `json:"useServiceAccount,omitempty"`
    UseIamRole                   *bool `json:"useIamRole,omitempty"`
    RoleARN                      *string `json:"RoleARN,omitempty"`
}

type MsgActivityControlOptions struct {
    RestoreActivityControlOptions     *MsgbackupActivityControlOptionsProp `json:"restoreActivityControlOptions,omitempty"`
    BackupActivityControlOptions      *MsgbackupActivityControlOptionsProp `json:"backupActivityControlOptions,omitempty"`
    EnableBackup                      *bool `json:"enableBackup,omitempty"`
    EnableRestore                     *bool `json:"enableRestore,omitempty"`
}

type MsgbackupActivityControlOptionsProp struct {
    TimeZone             *MsgIdName `json:"timeZone,omitempty"`
    EnableAfterDelay     *bool `json:"enableAfterDelay,omitempty"`
    DelayTime            *string `json:"delayTime,omitempty"`
}

type MsghypervisorSettings struct {
    MetricsMonitoringPolicy     *MsghypervisorMonitoringPolicy `json:"metricsMonitoringPolicy,omitempty"`
    ApplicationCredentials      *MsgIdName `json:"applicationCredentials,omitempty"`
    GuestCredentials            *MsgIdName `json:"guestCredentials,omitempty"`
    MountAccessNode             *MsgIdName `json:"mountAccessNode,omitempty"`
    RegionInfo                  *MsgRegionInfo `json:"regionInfo,omitempty"`
    TimeZone                    *MsgIdName `json:"timeZone,omitempty"`
    CustomAttributes            *MsghypervisorCustomAttribute `json:"customAttributes,omitempty"`
}

type MsghypervisorMonitoringPolicy struct {
    IsEnabled     *bool `json:"isEnabled,omitempty"`
    Name          *string `json:"name,omitempty"`
    Id            *int `json:"id,omitempty"`
}

type MsgRegionInfo struct {
    DisplayName     *string `json:"displayName,omitempty"`
    Latitude        *string `json:"latitude,omitempty"`
    Name            *string `json:"name,omitempty"`
    Id              *int `json:"id,omitempty"`
    Longitude       *string `json:"longitude,omitempty"`
}

type MsghypervisorCustomAttribute struct {
    Type      *int `json:"type,omitempty"`
    Value     *string `json:"value,omitempty"`
}

type MsgVMHypervisorSecurityProp struct {
    ClientOwners             *string `json:"clientOwners,omitempty"`
    AssociatedUserGroups[]   MsgIdNameSet `json:"associatedUserGroups,omitempty"`
}

type MsgaccessNodeModelSet struct {
    Id     *int `json:"id,omitempty"`
}

type MsgupdateHypervisorAWSResponse struct {
    ActivityControl                *MsgActivityControlOptions `json:"activityControl,omitempty"`
    General                        *MsghypervisorGeneralProperties `json:"general,omitempty"`
    Settings                       *MsghypervisorSettings `json:"settings,omitempty"`
    AccessNodeList                 *MsgaccessNodeListModel `json:"accessNodeList,omitempty"`
    Instance                       *MsgIdName `json:"instance,omitempty"`
    HypervisorCommonProperties     *MsghypervisorCommonProps `json:"hypervisorCommonProperties,omitempty"`
    DisplayName                    *string `json:"displayName,omitempty"`
    AccountDetails                 *MsghypervisorAccountDetails `json:"accountDetails,omitempty"`
    Name                           *string `json:"name,omitempty"`
    Id                             *int `json:"id,omitempty"`
    ManageSnapshot                 *MsgIdName `json:"manageSnapshot,omitempty"`
}

type MsghypervisorGeneralProperties struct {
    VmBackupInfo     *MsgvmBackupInfo `json:"vmBackupInfo,omitempty"`
    Vendor           *string `json:"vendor,omitempty"`
    Version          *string `json:"version,omitempty"`
}

type MsgvmBackupInfo struct {
    VmPendingCount               *int `json:"vmPendingCount,omitempty"`
    VmProtectedCount             *int `json:"vmProtectedCount,omitempty"`
    VmNotProtectedCount          *int `json:"vmNotProtectedCount,omitempty"`
    VmNeverBackedUpCount         *int `json:"vmNeverBackedUpCount,omitempty"`
    VmBackedUpWithErrorCount     *int `json:"vmBackedUpWithErrorCount,omitempty"`
    VmTotalCount                 *int `json:"vmTotalCount,omitempty"`
}

type MsgaccessNodeListModel struct {
    AccessNodeMessage     *string `json:"accessNodeMessage,omitempty"`
    AccessNode[]          MsgaccessNodeModelSet `json:"accessNode,omitempty"`
}

type MsghypervisorCommonProps struct {
    IsDeconfigured     *bool `json:"isDeconfigured,omitempty"`
    RetirePhase        *string `json:"retirePhase,omitempty"`
    Company            *MsgIdName `json:"company,omitempty"`
    IsIndexingV2       *bool `json:"isIndexingV2,omitempty"`
}

type MsghypervisorAccountDetails struct {
    HostName     *string `json:"hostName,omitempty"`
}

type MsgGetHypervisorsResponse struct {
    ActivityControl                *MsgActivityControlOptions `json:"activityControl,omitempty"`
    General                        *MsghypervisorGeneralProperties `json:"general,omitempty"`
    Settings                       *MsghypervisorSettings `json:"settings,omitempty"`
    AccessNodeList                 *MsgaccessNodeListModel `json:"accessNodeList,omitempty"`
    Instance                       *MsgIdName `json:"instance,omitempty"`
    HypervisorCommonProperties     *MsghypervisorCommonProps `json:"hypervisorCommonProperties,omitempty"`
    DisplayName                    *string `json:"displayName,omitempty"`
    AccountDetails                 *MsghypervisorAccountDetails `json:"accountDetails,omitempty"`
    Name                           *string `json:"name,omitempty"`
    Id                             *int `json:"id,omitempty"`
    ManageSnapshot                 *MsgIdName `json:"manageSnapshot,omitempty"`
}

type MsgDeleteHypervisorResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateHypervisorAzureRequest struct {
    SkipCredentialValidation     *bool `json:"skipCredentialValidation,omitempty"`
    EtcdProtection               *bool `json:"etcdProtection,omitempty"`
    Credentials                  *MsgIdName `json:"credentials,omitempty"`
    Name                         *string `json:"name,omitempty"`
    AccessNodes[]                MsgIdNameSet `json:"accessNodes,omitempty"`
    PlanEntity                   *MsgIdName `json:"planEntity,omitempty"`
    ApplicationPassword          *string `json:"ApplicationPassword,omitempty"`
    TenantId                     *string `json:"tenantId,omitempty"`
    HypervisorType               *string `json:"hypervisorType,omitempty"`
    WorkloadRegion               *MsgIdName `json:"workloadRegion,omitempty"`
    SubscriptionId               *string `json:"subscriptionId,omitempty"`
    ApplicationId                *string `json:"ApplicationId,omitempty"`
}

type MsgCreateHypervisorAzureResponse struct {
    Response     *MsgCreateHypervisorResp `json:"response,omitempty"`
}

type MsgupdateHypervisorAzureRequest struct {
    ActivityControl              *MsgActivityControlOptions `json:"activityControl,omitempty"`
    Settings                     *MsghypervisorSettings `json:"settings,omitempty"`
    Security                     *MsgVMHypervisorSecurityProp `json:"security,omitempty"`
    NewName                      *string `json:"newName,omitempty"`
    SkipCredentialValidation     *bool `json:"skipCredentialValidation,omitempty"`
    AccessNode[]                 MsgaccessNodeModelSet `json:"accessNode,omitempty"`
    FbrUnixMediaAgent            *MsgIdName `json:"fbrUnixMediaAgent,omitempty"`
    Password                     *string `json:"password,omitempty"`
    TenantId                     *string `json:"tenantId,omitempty"`
    ServerName                   *string `json:"serverName,omitempty"`
    HypervisorType               *string `json:"hypervisorType,omitempty"`
    SubscriptionId               *string `json:"subscriptionId,omitempty"`
    UserName                     *string `json:"userName,omitempty"`
}

type MsgupdateHypervisorAzureResponse struct {
    ActivityControl                *MsgActivityControlOptions `json:"activityControl,omitempty"`
    General                        *MsghypervisorGeneralProperties `json:"general,omitempty"`
    Settings                       *MsghypervisorSettings `json:"settings,omitempty"`
    AccessNodeList                 *MsgaccessNodeListModel `json:"accessNodeList,omitempty"`
    Instance                       *MsgIdName `json:"instance,omitempty"`
    HypervisorCommonProperties     *MsghypervisorCommonProps `json:"hypervisorCommonProperties,omitempty"`
    DisplayName                    *string `json:"displayName,omitempty"`
    AccountDetails                 *MsghypervisorAccountDetails `json:"accountDetails,omitempty"`
    Name                           *string `json:"name,omitempty"`
    Id                             *int `json:"id,omitempty"`
    ManageSnapshot                 *MsgIdName `json:"manageSnapshot,omitempty"`
}

type MsgGetVMGroupResponse struct {
    Summary                   *MsgvmGroupDetailsSummary `json:"summary,omitempty"`
    ActivityControl           *MsgActivityControlOptions `json:"activityControl,omitempty"`
    Settings                  *MsgvmGroupSettings `json:"settings,omitempty"`
    AccessNodeList            *MsgaccessNodeListModel `json:"accessNodeList,omitempty"`
    DiskFilters[]             MsgvmDiskFilterPropSet `json:"diskFilters,omitempty"`
    VmBackupInfo              *MsgvmBackupInfo `json:"vmBackupInfo,omitempty"`
    SecurityAssociations[]    MsgSecurityAssocSet `json:"securityAssociations,omitempty"`
    Filters[]                 MsgvmContentSet `json:"filters,omitempty"`
    Content[]                 MsgvmContentSet `json:"content,omitempty"`
    CommonProperties          *MsgVMGroupCommonProperties `json:"commonProperties,omitempty"`
    SnapshotManagement        *MsgsnapCopyInfo `json:"snapshotManagement,omitempty"`
    Name                      *string `json:"name,omitempty"`
    Id                        *int `json:"id,omitempty"`
    ApplicationValidation     *MsgvmAppValidation `json:"applicationValidation,omitempty"`
    Status                    *int `json:"status,omitempty"`
    MeditechSystems           *MsgmeditechPropResp `json:"meditechSystems,omitempty"`
}

type MsgvmGroupDetailsSummary struct {
    LastBackupSize           *int64 `json:"lastBackupSize,omitempty"`
    IsDefaultVMGroup         *bool `json:"isDefaultVMGroup,omitempty"`
    HypervisorName           *string `json:"hypervisorName,omitempty"`
    BackupActivityStatus     *string `json:"backupActivityStatus,omitempty"`
    NextBackupTime           *int `json:"nextBackupTime,omitempty"`
    ReplicationGroup         *MsgIdName `json:"replicationGroup,omitempty"`
    TimeZone                 *MsgIdName `json:"timeZone,omitempty"`
    Region                   *MsgRegionInfo `json:"region,omitempty"`
    Plan                     *MsgPlanIdNameType `json:"plan,omitempty"`
    LastBackupTime           *int `json:"lastBackupTime,omitempty"`
}

type MsgPlanIdNameType struct {
    Name        *string `json:"name,omitempty"`
    SubType     *string `json:"subType,omitempty"`
    Id          *int `json:"id,omitempty"`
}

type MsgvmGroupSettings struct {
    AutoDetectVMOwner                         *bool `json:"autoDetectVMOwner,omitempty"`
    CollectFileDetailsforGranularRecovery     *bool `json:"collectFileDetailsforGranularRecovery,omitempty"`
    NoOfReaders                               *int `json:"noOfReaders,omitempty"`
    UseChangedBlockTrackingOnVM               *bool `json:"useChangedBlockTrackingOnVM,omitempty"`
    JobStartTime                              *int `json:"jobStartTime,omitempty"`
    UseVMCheckpointSetting                    *bool `json:"useVMCheckpointSetting,omitempty"`
    CustomSnapshotResourceGroup               *string `json:"customSnapshotResourceGroup,omitempty"`
    RegionalSnapshot                          *bool `json:"regionalSnapshot,omitempty"`
    GuestCredentials                          *MsgguestCredentialInfo `json:"guestCredentials,omitempty"`
    VmBackupType                              *string `json:"vmBackupType,omitempty"`
    DatastoreFreespaceCheck                   *bool `json:"datastoreFreespaceCheck,omitempty"`
    DatastoreFreespaceRequired                *int `json:"datastoreFreespaceRequired,omitempty"`
    CustomSnapshotTags[]                      MsgresourceTagSet `json:"customSnapshotTags,omitempty"`
    IsApplicationAware                        *bool `json:"isApplicationAware,omitempty"`
    TransportMode                             *string `json:"transportMode,omitempty"`
    CollectFileDetailsFromSnapshotCopy        *bool `json:"collectFileDetailsFromSnapshotCopy,omitempty"`
    CrossAccount                              *MsgAmazonCrossAccount `json:"crossAccount,omitempty"`
}

type MsgguestCredentialInfo struct {
    Credentials          *MsgUserNamePassword `json:"credentials,omitempty"`
    SavedCredentials     *MsgIdName `json:"savedCredentials,omitempty"`
}

type MsgUserNamePassword struct {
    Password     *string `json:"password,omitempty"`
    Name         *string `json:"name,omitempty"`
}

type MsgresourceTagSet struct {
    Name      *string `json:"name,omitempty"`
    Value     *string `json:"value,omitempty"`
}

type MsgAmazonCrossAccount struct {
    ShareOnly              *bool `json:"shareOnly,omitempty"`
    FullCopy               *bool `json:"fullCopy,omitempty"`
    DestinationAccount     *MsgIdName `json:"destinationAccount,omitempty"`
}

type MsgvmDiskFilterPropSet struct {
    Rules[]       MsgvmDiskFilterSet `json:"rules,omitempty"`
    Overwrite     *bool `json:"overwrite,omitempty"`
}

type MsgvmDiskFilterSet struct {
    Condition      *string `json:"condition,omitempty"`
    VmName         *string `json:"vmName,omitempty"`
    Name           *string `json:"name,omitempty"`
    FilterType     *string `json:"filterType,omitempty"`
    Overwrite      *bool `json:"overwrite,omitempty"`
    Value          *string `json:"value,omitempty"`
    VmGuid         *string `json:"vmGuid,omitempty"`
}

type MsgSecurityAssocSet struct {
    Role                     *MsgIdName `json:"role,omitempty"`
    IsCreatorAssociation     *bool `json:"isCreatorAssociation,omitempty"`
    ExternalUserGroup        *MsgexternalUserGroup `json:"externalUserGroup,omitempty"`
    PermissionList[]         MsgPermissionRespSet `json:"permissionList,omitempty"`
    User                     *MsgIdName `json:"user,omitempty"`
    UserGroup                *MsgIdName `json:"userGroup,omitempty"`
}

type MsgexternalUserGroup struct {
    ProviderId       *int `json:"providerId,omitempty"`
    Name             *string `json:"name,omitempty"`
    Id               *int `json:"id,omitempty"`
    ProviderName     *string `json:"providerName,omitempty"`
}

type MsgPermissionRespSet struct {
    PermissionId       *int `json:"permissionId,omitempty"`
    Exclude            *bool `json:"exclude,omitempty"`
    Type               *string `json:"type,omitempty"`
    CategoryName       *string `json:"categoryName,omitempty"`
    CategoryId         *int `json:"categoryId,omitempty"`
    PermissionName     *string `json:"permissionName,omitempty"`
}

type MsgvmContentSet struct {
    RuleGroups[]        MsgRuleGroupContentSet `json:"ruleGroups,omitempty"`
    VirtualMachines[]   MsgVirtualMachinecontentSet `json:"virtualMachines,omitempty"`
    Overwrite           *bool `json:"overwrite,omitempty"`
}

type MsgRuleGroupContentSet struct {
    MatchRule     *string `json:"matchRule,omitempty"`
    Rules[]       MsgRuleContentSet `json:"rules,omitempty"`
}

type MsgRuleContentSet struct {
    Condition       *string `json:"condition,omitempty"`
    DisplayName     *string `json:"displayName,omitempty"`
    Name            *string `json:"name,omitempty"`
    GUID            *string `json:"GUID,omitempty"`
    Type            *string `json:"type,omitempty"`
    Value           *string `json:"value,omitempty"`
}

type MsgVirtualMachinecontentSet struct {
    GuestCredentialAssocId     *int `json:"guestCredentialAssocId,omitempty"`
    GuestCredentials           *MsgVMGuestCredentials `json:"guestCredentials,omitempty"`
    Name                       *string `json:"name,omitempty"`
    GUID                       *string `json:"GUID,omitempty"`
    Type                       *string `json:"type,omitempty"`
    ExistingCredential         *MsgVMExistingCredential `json:"existingCredential,omitempty"`
}

type MsgVMGuestCredentials struct {
    Password     *string `json:"password,omitempty"`
    UserName     *string `json:"userName,omitempty"`
}

type MsgVMExistingCredential struct {
    CredentialId       *int `json:"credentialId,omitempty"`
    CredentialName     *string `json:"credentialName,omitempty"`
}

type MsgVMGroupCommonProperties struct {
    DataBackupStoragePolicy      *MsgIdName `json:"dataBackupStoragePolicy,omitempty"`
    IsDeletedHypervisor          *bool `json:"isDeletedHypervisor,omitempty"`
    Instance                     *MsgIdName `json:"instance,omitempty"`
    IsETCDSubclient              *bool `json:"isETCDSubclient,omitempty"`
    Hypervisor                   *MsghypervisorvmGroupCommonProp `json:"hypervisor,omitempty"`
    Backupset                    *MsgIdName `json:"backupset,omitempty"`
    IsIndexingV2                 *bool `json:"isIndexingV2,omitempty"`
    ShowFullBackupLevel          *bool `json:"showFullBackupLevel,omitempty"`
    HypervisorType               *string `json:"hypervisorType,omitempty"`
    IndexingInfo                 *MsgindexingInfo `json:"indexingInfo,omitempty"`
    IsHypervisorDeconfigured     *bool `json:"isHypervisorDeconfigured,omitempty"`
    IDataAgent                   *MsgIdName `json:"iDataAgent,omitempty"`
}

type MsghypervisorvmGroupCommonProp struct {
    ActivityControl     *MsgActivityControlOptions `json:"activityControl,omitempty"`
    Name                *string `json:"name,omitempty"`
    Id                  *int `json:"id,omitempty"`
}

type MsgindexingInfo struct {
    Message     *string `json:"message,omitempty"`
    Status      *string `json:"status,omitempty"`
}

type MsgsnapCopyInfo struct {
    UseSeparateProxyForSnapToTape     *bool `json:"useSeparateProxyForSnapToTape,omitempty"`
    SnapEngine                        *MsgIdName `json:"snapEngine,omitempty"`
    IsIndependentDisksEnabled         *bool `json:"isIndependentDisksEnabled,omitempty"`
    BackupCopyInterface               *string `json:"backupCopyInterface,omitempty"`
    EnableHardwareSnapshot            *bool `json:"enableHardwareSnapshot,omitempty"`
    SnapMountProxy                    *MsgIdName `json:"snapMountProxy,omitempty"`
    VmApplicationUserName             *string `json:"vmApplicationUserName,omitempty"`
    SnapMountESXHost                  *string `json:"snapMountESXHost,omitempty"`
    IsRawDeviceMapsEnabled            *bool `json:"isRawDeviceMapsEnabled,omitempty"`
}

type MsgvmAppValidation struct {
    RecoveryTarget              *MsgIdName `json:"recoveryTarget,omitempty"`
    Schedule                    *MsgValidationScheduleObject `json:"schedule,omitempty"`
    MaximumNoOfThreads          *int `json:"maximumNoOfThreads,omitempty"`
    GuestCredentials            *MsgguestCredentialInfo `json:"guestCredentials,omitempty"`
    KeepValidatedVMsRunning     *bool `json:"keepValidatedVMsRunning,omitempty"`
    ValidateVMBackups           *bool `json:"validateVMBackups,omitempty"`
    UseSourceVmESXToMount       *bool `json:"useSourceVmESXToMount,omitempty"`
    CustomValidationScript      *MsgappValidationScript `json:"customValidationScript,omitempty"`
    Copy                        *MsgPlanSourceCopy `json:"copy,omitempty"`
}

type MsgValidationScheduleObject struct {
    IsScheduleEnabled     *bool `json:"isScheduleEnabled,omitempty"`
    Description           *string `json:"description,omitempty"`
    Id                    *int `json:"id,omitempty"`
    TaskId                *int `json:"taskId,omitempty"`
}

type MsgappValidationScript struct {
    Windows     *MsgValidationScript `json:"windows,omitempty"`
    Unix        *MsgValidationScript `json:"unix,omitempty"`
}

type MsgValidationScript struct {
    Path                    *string `json:"path,omitempty"`
    UNCCredentials          *MsgUserNamePassword `json:"UNCCredentials,omitempty"`
    UNCSavedCredentials     *MsgIdName `json:"UNCSavedCredentials,omitempty"`
    Arguments               *string `json:"arguments,omitempty"`
    IsDisabled              *bool `json:"isDisabled,omitempty"`
    IsLocal                 *bool `json:"isLocal,omitempty"`
}

type MsgPlanSourceCopy struct {
    IsMirrorCopy           *bool `json:"isMirrorCopy,omitempty"`
    SnapCopyType           *string `json:"snapCopyType,omitempty"`
    IsDefault              *bool `json:"isDefault,omitempty"`
    CopyPrecedence         *int `json:"copyPrecedence,omitempty"`
    IsSnapCopy             *bool `json:"isSnapCopy,omitempty"`
    CopyType               *string `json:"copyType,omitempty"`
    DefaultReplicaCopy     *bool `json:"defaultReplicaCopy,omitempty"`
    IsActive               *bool `json:"isActive,omitempty"`
    ArrayReplicaCopy       *bool `json:"arrayReplicaCopy,omitempty"`
    BackupDestination      *MsgIdName `json:"backupDestination,omitempty"`
}

type MsgmeditechPropResp struct {
    SystemName       *string `json:"systemName,omitempty"`
    ListenerIP       *string `json:"listenerIP,omitempty"`
    UserAccount      *MsgUserNamePassword `json:"userAccount,omitempty"`
    ListenerPort     *int `json:"listenerPort,omitempty"`
    MBFtimeout       *int `json:"MBFtimeout,omitempty"`
}

type MsgDeleteVMGroupResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgUpdateVMGroupRequest struct {
    ActivityControl           *MsgActivityControlOptions `json:"activityControl,omitempty"`
    Settings                  *MsgvmGroupSettings `json:"settings,omitempty"`
    DiskFilters               *MsgvmDiskFilterProp `json:"diskFilters,omitempty"`
    SecurityAssociations[]    MsgSecurityAssocSet `json:"securityAssociations,omitempty"`
    Storage                   *MsgIdName `json:"storage,omitempty"`
    Filters                   *MsgvmContent `json:"filters,omitempty"`
    AccessNode[]              MsgIdNameSet `json:"accessNode,omitempty"`
    Content                   *MsgvmContent `json:"content,omitempty"`
    SnapshotManagement        *MsgsnapCopyInfo `json:"snapshotManagement,omitempty"`
    EnableFileIndexing        *bool `json:"enableFileIndexing,omitempty"`
    NewName                   *string `json:"newName,omitempty"`
    ApplicationValidation     *MsgvmAppValidation `json:"applicationValidation,omitempty"`
    Plan                      *MsgIdName `json:"plan,omitempty"`
    MeditechSystems           *MsgmeditechPropResp `json:"meditechSystems,omitempty"`
}

type MsgvmDiskFilterProp struct {
    Rules[]       MsgvmDiskFilterSet `json:"rules,omitempty"`
    Overwrite     *bool `json:"overwrite,omitempty"`
}

type MsgvmContent struct {
    RuleGroups[]        MsgRuleGroupContentSet `json:"ruleGroups,omitempty"`
    VirtualMachines[]   MsgVirtualMachinecontentSet `json:"virtualMachines,omitempty"`
    Overwrite           *bool `json:"overwrite,omitempty"`
}

type MsgUpdateVMGroupResponse struct {
    Summary                   *MsgvmGroupDetailsSummary `json:"summary,omitempty"`
    ActivityControl           *MsgActivityControlOptions `json:"activityControl,omitempty"`
    Settings                  *MsgvmGroupSettings `json:"settings,omitempty"`
    AccessNodeList            *MsgaccessNodeListModel `json:"accessNodeList,omitempty"`
    DiskFilters[]             MsgvmDiskFilterPropSet `json:"diskFilters,omitempty"`
    VmBackupInfo              *MsgvmBackupInfo `json:"vmBackupInfo,omitempty"`
    SecurityAssociations[]    MsgSecurityAssocSet `json:"securityAssociations,omitempty"`
    Filters[]                 MsgvmContentSet `json:"filters,omitempty"`
    Content[]                 MsgvmContentSet `json:"content,omitempty"`
    CommonProperties          *MsgVMGroupCommonProperties `json:"commonProperties,omitempty"`
    SnapshotManagement        *MsgsnapCopyInfo `json:"snapshotManagement,omitempty"`
    Name                      *string `json:"name,omitempty"`
    Id                        *int `json:"id,omitempty"`
    ApplicationValidation     *MsgvmAppValidation `json:"applicationValidation,omitempty"`
    Status                    *int `json:"status,omitempty"`
    MeditechSystems           *MsgmeditechPropResp `json:"meditechSystems,omitempty"`
}

type MsgCreateVMGroupRequest struct {
    Meditech              *MsgmeditechPropResp `json:"Meditech,omitempty"`
    Hypervisor            *MsgIdName `json:"Hypervisor,omitempty"`
    Name                  *string `json:"name,omitempty"`
    Storage               *MsgIdName `json:"storage,omitempty"`
    EnableIntellisnap     *bool `json:"enableIntellisnap,omitempty"`
    Plan                  *MsgIdName `json:"plan,omitempty"`
    Content               *MsgvmContent `json:"content,omitempty"`
}

type MsgCreateVMGroupResponse struct {
    SubclientId        *int `json:"subclientId,omitempty"`
    WarningCode        *int `json:"warningCode,omitempty"`
    ErrorMessage       *string `json:"errorMessage,omitempty"`
    ErrorCode          *int `json:"errorCode,omitempty"`
    WarningMessage     *string `json:"warningMessage,omitempty"`
}

type MsgGetPlanByIdResponse struct {
    Settings                 *MsgServerPlanSettings `json:"settings,omitempty"`
    BackupContent            *MsgPlanContent `json:"backupContent,omitempty"`
    DatabaseOptions          *MsgServerPlanDatabaseOptionsInfo `json:"databaseOptions,omitempty"`
    AllowPlanOverride        *bool `json:"allowPlanOverride,omitempty"`
    Workload                 *MsgPlanWorkloads `json:"workload,omitempty"`
    InheritSettings          *MsgServerPlanInheritSettings `json:"inheritSettings,omitempty"`
    Rpo                      *MsgServerPlanRPO `json:"rpo,omitempty"`
    AssociatedEntities[]     MsgIdNameCountSet `json:"associatedEntities,omitempty"`
    BackupDestinations[]     MsgPlanBackupDestinationSet `json:"backupDestinations,omitempty"`
    Permissions[]            MsgIdNameSet `json:"permissions,omitempty"`
    OverrideRestrictions     *MsgPlanOverrideSettings `json:"overrideRestrictions,omitempty"`
    SnapshotOptions          *MsgPlanSnapshotOptions `json:"snapshotOptions,omitempty"`
    AdditionalProperties     *MsgPlanAdditionalProperties `json:"additionalProperties,omitempty"`
    RegionsConfigured        *bool `json:"regionsConfigured,omitempty"`
}

type MsgServerPlanSettings struct {
    EnableAdvancedView     *bool `json:"enableAdvancedView,omitempty"`
    FileSearch             *MsgPlanFileSearch `json:"fileSearch,omitempty"`
}

type MsgPlanFileSearch struct {
    Enabled           *bool `json:"enabled,omitempty"`
    StatusMessage     *string `json:"statusMessage,omitempty"`
    Status            *string `json:"status,omitempty"`
}

type MsgPlanContent struct {
    WindowsIncludedPaths[]                  string `json:"windowsIncludedPaths,omitempty"`
    BackupSystemState                       *bool `json:"backupSystemState,omitempty"`
    BackupSystemStateOnlyWithFullBackup     *bool `json:"backupSystemStateOnlyWithFullBackup,omitempty"`
    WindowsExcludedPaths[]                  string `json:"windowsExcludedPaths,omitempty"`
    UnixIncludedPaths[]                     string `json:"unixIncludedPaths,omitempty"`
    UseVSSForSystemState                    *bool `json:"useVSSForSystemState,omitempty"`
    MacExcludedPaths[]                      string `json:"macExcludedPaths,omitempty"`
    MacFilterToExcludePaths[]               string `json:"macFilterToExcludePaths,omitempty"`
    WindowsFilterToExcludePaths[]           string `json:"windowsFilterToExcludePaths,omitempty"`
    UnixFilterToExcludePaths[]              string `json:"unixFilterToExcludePaths,omitempty"`
    MacIncludedPaths[]                      string `json:"macIncludedPaths,omitempty"`
    UnixExcludedPaths[]                     string `json:"unixExcludedPaths,omitempty"`
}

type MsgServerPlanDatabaseOptionsInfo struct {
    LogBackupRPOMins              *int `json:"logBackupRPOMins,omitempty"`
    RunFullBackupEvery            *int `json:"runFullBackupEvery,omitempty"`
    CommitFrequencyInHours        *int `json:"commitFrequencyInHours,omitempty"`
    UseDiskCacheForLogBackups     *bool `json:"useDiskCacheForLogBackups,omitempty"`
}

type MsgPlanWorkloads struct {
    WorkloadTypes[]        MsgIdNameSet `json:"workloadTypes,omitempty"`
    WorkloadGroupTypes[]   string `json:"workloadGroupTypes,omitempty"`
    Solutions[]            MsgIdNameSet `json:"solutions,omitempty"`
}

type MsgServerPlanInheritSettings struct {
    RPO               *MsgPlanOverridenOptions `json:"RPO,omitempty"`
    BackupContent     *MsgPlanOverridenOptions `json:"backupContent,omitempty"`
    BasePlan          *MsgIdName `json:"basePlan,omitempty"`
    StoragePool       *MsgPlanOverridenOptions `json:"storagePool,omitempty"`
}

type MsgPlanOverridenOptions struct {
    OverrideBase     *string `json:"overrideBase,omitempty"`
    Overridden       *bool `json:"overridden,omitempty"`
}

type MsgServerPlanRPO struct {
    FullBackupWindow[]   MsgDayAndTimeSet `json:"fullBackupWindow,omitempty"`
    SLA                  *MsgSLAOptions `json:"SLA,omitempty"`
    BackupFrequency      *MsgPlanSchedules `json:"backupFrequency,omitempty"`
    BackupWindow[]       MsgDayAndTimeSet `json:"backupWindow,omitempty"`
}

type MsgDayAndTimeSet struct {
    DayOfWeek[]   string `json:"dayOfWeek,omitempty"`
    StartTime     *int64 `json:"startTime,omitempty"`
    EndTime       *int64 `json:"endTime,omitempty"`
}

type MsgSLAOptions struct {
    ExclusionReason         *string `json:"exclusionReason,omitempty"`
    InheritedSLAPeriod      *int `json:"inheritedSLAPeriod,omitempty"`
    UseSystemDefaultSLA     *bool `json:"useSystemDefaultSLA,omitempty"`
    EnableAfterDelay        *int `json:"enableAfterDelay,omitempty"`
    InheritedFrom           *string `json:"inheritedFrom,omitempty"`
    ExcludeFromSLA          *bool `json:"excludeFromSLA,omitempty"`
    SLAPeriod               *int `json:"SLAPeriod,omitempty"`
}

type MsgPlanSchedules struct {
    Schedules[]   MsgPlanScheduleSet `json:"schedules,omitempty"`
}

type MsgPlanScheduleSet struct {
    ScheduleName          *string `json:"scheduleName,omitempty"`
    PolicyId              *int `json:"policyId,omitempty"`
    VmOperationType       *string `json:"vmOperationType,omitempty"`
    ForDatabasesOnly      *bool `json:"forDatabasesOnly,omitempty"`
    ScheduleOperation     *string `json:"scheduleOperation,omitempty"`
    SchedulePattern       *MsgSchedulePattern `json:"schedulePattern,omitempty"`
    ScheduleId            *int `json:"scheduleId,omitempty"`
    BackupType            *string `json:"backupType,omitempty"`
}

type MsgSchedulePattern struct {
    EndDate                       *int `json:"endDate,omitempty"`
    Timezone                      *MsgIdName `json:"timezone,omitempty"`
    WeekOfMonth                   *string `json:"weekOfMonth,omitempty"`
    DaysBetweenSyntheticFulls     *int `json:"daysBetweenSyntheticFulls,omitempty"`
    Exceptions[]                  MsgScheduleRunExceptionSet `json:"exceptions,omitempty"`
    Frequency                     *int `json:"frequency,omitempty"`
    WeeklyDays[]                  string `json:"weeklyDays,omitempty"`
    RepeatUntilTime               *int `json:"repeatUntilTime,omitempty"`
    MonthOfYear                   *string `json:"monthOfYear,omitempty"`
    DayOfWeek                     *string `json:"dayOfWeek,omitempty"`
    DayOfMonth                    *int `json:"dayOfMonth,omitempty"`
    ScheduleFrequencyType         *string `json:"scheduleFrequencyType,omitempty"`
    StartTime                     *int `json:"startTime,omitempty"`
    NoOfTimes                     *int `json:"noOfTimes,omitempty"`
    RepeatIntervalInMinutes       *int `json:"repeatIntervalInMinutes,omitempty"`
    StartDate                     *int `json:"startDate,omitempty"`
}

type MsgScheduleRunExceptionSet struct {
    OnWeekOfTheMonth[]   string `json:"onWeekOfTheMonth,omitempty"`
    OnDates[]            int `json:"onDates,omitempty"`
    OnDayOfTheWeek[]     string `json:"onDayOfTheWeek,omitempty"`
}

type MsgIdNameCountSet struct {
    Name      *string `json:"name,omitempty"`
    Count     *int `json:"count,omitempty"`
    Id        *int `json:"id,omitempty"`
}

type MsgPlanBackupDestinationSet struct {
    IsMirrorCopy                  *bool `json:"isMirrorCopy,omitempty"`
    CopyPrecedence                *int `json:"copyPrecedence,omitempty"`
    RetentionPeriodDays           *int `json:"retentionPeriodDays,omitempty"`
    CopyTypeName                  *string `json:"copyTypeName,omitempty"`
    BackupsToCopy                 *string `json:"backupsToCopy,omitempty"`
    ExtendedRetentionRules        *MsgExtendedRetentionRules `json:"extendedRetentionRules,omitempty"`
    RetentionRuleType             *string `json:"retentionRuleType,omitempty"`
    SnapRecoveryPoints            *int `json:"snapRecoveryPoints,omitempty"`
    SourceCopy                    *MsgIdName `json:"sourceCopy,omitempty"`
    FullBackupTypesToCopy         *string `json:"fullBackupTypesToCopy,omitempty"`
    UseExtendedRetentionRules     *bool `json:"useExtendedRetentionRules,omitempty"`
    BackupStartTime               *int `json:"backupStartTime,omitempty"`
    OverrideRetentionSettings     *bool `json:"overrideRetentionSettings,omitempty"`
    NetAppCloudTarget             *bool `json:"netAppCloudTarget,omitempty"`
    IsDefault                     *bool `json:"isDefault,omitempty"`
    Mappings[]                    MsgSnapshotCopyMappingSet `json:"mappings,omitempty"`
    PlanBackupDestination         *MsgIdName `json:"planBackupDestination,omitempty"`
    IsSnapCopy                    *bool `json:"isSnapCopy,omitempty"`
    CopyType                      *string `json:"copyType,omitempty"`
    StorageType                   *string `json:"storageType,omitempty"`
    EnableDataAging               *bool `json:"enableDataAging,omitempty"`
    Region                        *MsgIdNameDisplayName `json:"region,omitempty"`
    StoragePool                   *MsgStoragePool `json:"storagePool,omitempty"`
}

type MsgExtendedRetentionRules struct {
    ThirdExtendedRetentionRule      *MsgPlanRetentionRule `json:"thirdExtendedRetentionRule,omitempty"`
    FirstExtendedRetentionRule      *MsgPlanRetentionRule `json:"firstExtendedRetentionRule,omitempty"`
    SecondExtendedRetentionRule     *MsgPlanRetentionRule `json:"secondExtendedRetentionRule,omitempty"`
}

type MsgPlanRetentionRule struct {
    IsInfiniteRetention     *bool `json:"isInfiniteRetention,omitempty"`
    RetentionPeriodDays     *int `json:"retentionPeriodDays,omitempty"`
    Type                    *string `json:"type,omitempty"`
}

type MsgSnapshotCopyMappingSet struct {
    Vendor           *string `json:"vendor,omitempty"`
    TargetVendor     *MsgIdName `json:"targetVendor,omitempty"`
    Source           *MsgIdName `json:"source,omitempty"`
    SourceVendor     *MsgIdName `json:"sourceVendor,omitempty"`
    Target           *MsgIdName `json:"target,omitempty"`
}

type MsgIdNameDisplayName struct {
    DisplayName     *string `json:"displayName,omitempty"`
    Name            *string `json:"name,omitempty"`
    Id              *int `json:"id,omitempty"`
}

type MsgStoragePool struct {
    RetentionPeriodDays     *int `json:"retentionPeriodDays,omitempty"`
    WormStoragePoolFlag     *int `json:"wormStoragePoolFlag,omitempty"`
    Name                    *string `json:"name,omitempty"`
    Id                      *int `json:"id,omitempty"`
    Type                    *string `json:"type,omitempty"`
}

type MsgPlanOverrideSettings struct {
    RPO               *string `json:"RPO,omitempty"`
    BackupContent     *string `json:"backupContent,omitempty"`
    StoragePool       *string `json:"storagePool,omitempty"`
}

type MsgPlanSnapshotOptions struct {
    BackupCopyFrequency     *MsgBackupFrequencyPattern `json:"backupCopyFrequency,omitempty"`
    EnableBackupCopy        *bool `json:"enableBackupCopy,omitempty"`
    BackupCopyRPOMins       *int `json:"backupCopyRPOMins,omitempty"`
}

type MsgBackupFrequencyPattern struct {
    WeeklyDays[]              string `json:"weeklyDays,omitempty"`
    MonthOfYear               *string `json:"monthOfYear,omitempty"`
    DayOfWeek                 *string `json:"dayOfWeek,omitempty"`
    DayOfMonth                *int `json:"dayOfMonth,omitempty"`
    ScheduleFrequencyType     *string `json:"scheduleFrequencyType,omitempty"`
    WeekOfMonth               *string `json:"weekOfMonth,omitempty"`
    StartTime                 *int `json:"startTime,omitempty"`
    Frequency                 *int `json:"frequency,omitempty"`
}

type MsgPlanAdditionalProperties struct {
    RPO        *int `json:"RPO,omitempty"`
    Addons     *MsgPlanAddons `json:"addons,omitempty"`
    Status     *string `json:"status,omitempty"`
}

type MsgPlanAddons struct {
    FileSystem     *bool `json:"fileSystem,omitempty"`
    IndexCopy      *bool `json:"indexCopy,omitempty"`
    Database       *bool `json:"database,omitempty"`
    SnapStatus     *string `json:"snapStatus,omitempty"`
    Snap           *bool `json:"snap,omitempty"`
}

type MsgDeletePlanResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgModifyPlanRequest struct {
    Rpo                         *MsgServerPlanUpdateRPO `json:"rpo,omitempty"`
    RegionToConfigure           *MsgIdName `json:"regionToConfigure,omitempty"`
    Settings                    *MsgServerPlanSettings `json:"settings,omitempty"`
    BackupContent               *MsgPlanContent `json:"backupContent,omitempty"`
    DatabaseOptions             *MsgServerPlanDatabaseOptionsInfo `json:"databaseOptions,omitempty"`
    NewName                     *string `json:"newName,omitempty"`
    OverrideInheritSettings     *MsgPlanOverrideInheritSetting `json:"overrideInheritSettings,omitempty"`
    FilesystemAddon             *bool `json:"filesystemAddon,omitempty"`
    AllowPlanOverride           *bool `json:"allowPlanOverride,omitempty"`
    Workload                    *MsgPlanWorkloads `json:"workload,omitempty"`
    OverrideRestrictions        *MsgPlanOverrideSettings `json:"overrideRestrictions,omitempty"`
    SnapshotOptions             *MsgPlanSnapshotOptions `json:"snapshotOptions,omitempty"`
}

type MsgServerPlanUpdateRPO struct {
    FullBackupWindow[]   MsgDayAndTimeSet `json:"fullBackupWindow,omitempty"`
    SLA                  *MsgSLAUpdateOptions `json:"SLA,omitempty"`
    BackupFrequency      *MsgPlanSchedules `json:"backupFrequency,omitempty"`
    BackupWindow[]       MsgDayAndTimeSet `json:"backupWindow,omitempty"`
}

type MsgSLAUpdateOptions struct {
    ExclusionReason         *string `json:"exclusionReason,omitempty"`
    UseSystemDefaultSLA     *bool `json:"useSystemDefaultSLA,omitempty"`
    EnableAfterDelay        *int `json:"enableAfterDelay,omitempty"`
    ExcludeFromSLA          *bool `json:"excludeFromSLA,omitempty"`
    SLAPeriod               *int `json:"SLAPeriod,omitempty"`
}

type MsgPlanOverrideInheritSetting struct {
    Rpo                   *bool `json:"rpo,omitempty"`
    BackupContent         *bool `json:"backupContent,omitempty"`
    BackupDestination     *bool `json:"backupDestination,omitempty"`
}

type MsgModifyPlanResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateServerPlanRequest struct {
    Settings                 *MsgServerPlanSettings `json:"settings,omitempty"`
    BackupContent            *MsgPlanContent `json:"backupContent,omitempty"`
    DatabaseOptions          *MsgServerPlanDatabaseOptions `json:"databaseOptions,omitempty"`
    FilesystemAddon          *bool `json:"filesystemAddon,omitempty"`
    AllowPlanOverride        *bool `json:"allowPlanOverride,omitempty"`
    PlanName                 *string `json:"planName,omitempty"`
    Workload                 *MsgPlanWorkloads `json:"workload,omitempty"`
    Rpo                      *MsgServerBackupPlanRPO `json:"rpo,omitempty"`
    BackupDestinations[]     MsgCreatePlanBackupDestinationSet `json:"backupDestinations,omitempty"`
    OverrideRestrictions     *MsgPlanOverrideSettings `json:"overrideRestrictions,omitempty"`
    SnapshotOptions          *MsgCreatePlanSnapshotOptions `json:"snapshotOptions,omitempty"`
    ParentPlan               *MsgIdName `json:"parentPlan,omitempty"`
    AdditionalProperties     *MsgPlanAdditionalProperties `json:"additionalProperties,omitempty"`
}

type MsgServerPlanDatabaseOptions struct {
    LogBackupRPOMins              *int `json:"logBackupRPOMins,omitempty"`
    CommitFrequencyInHours        *int `json:"commitFrequencyInHours,omitempty"`
    UseDiskCacheForLogBackups     *bool `json:"useDiskCacheForLogBackups,omitempty"`
}

type MsgServerBackupPlanRPO struct {
    FullBackupWindow[]   MsgDayAndTimeSet `json:"fullBackupWindow,omitempty"`
    BackupFrequency      *MsgPlanSchedules `json:"backupFrequency,omitempty"`
    BackupWindow[]       MsgDayAndTimeSet `json:"backupWindow,omitempty"`
}

type MsgCreatePlanBackupDestinationSet struct {
    IsMirrorCopy                  *bool `json:"isMirrorCopy,omitempty"`
    RetentionPeriodDays           *int `json:"retentionPeriodDays,omitempty"`
    BackupsToCopy                 *string `json:"backupsToCopy,omitempty"`
    BackupDestinationName         *string `json:"backupDestinationName,omitempty"`
    ExtendedRetentionRules        *MsgExtendedRetentionRules `json:"extendedRetentionRules,omitempty"`
    RetentionRuleType             *string `json:"retentionRuleType,omitempty"`
    SnapRecoveryPoints            *int `json:"snapRecoveryPoints,omitempty"`
    SourceCopy                    *MsgIdName `json:"sourceCopy,omitempty"`
    FullBackupTypesToCopy         *string `json:"fullBackupTypesToCopy,omitempty"`
    UseExtendedRetentionRules     *bool `json:"useExtendedRetentionRules,omitempty"`
    BackupStartTime               *int `json:"backupStartTime,omitempty"`
    OverrideRetentionSettings     *bool `json:"overrideRetentionSettings,omitempty"`
    OptimizeForInstantClone       *bool `json:"optimizeForInstantClone,omitempty"`
    NetAppCloudTarget             *bool `json:"netAppCloudTarget,omitempty"`
    Mappings[]                    MsgSnapshotCopyMappingSet `json:"mappings,omitempty"`
    IsSnapCopy                    *bool `json:"isSnapCopy,omitempty"`
    StorageType                   *string `json:"storageType,omitempty"`
    Region                        *MsgIdName `json:"region,omitempty"`
    StoragePool                   *MsgIdName `json:"storagePool,omitempty"`
}

type MsgCreatePlanSnapshotOptions struct {
    RetentionPeriodDays     *int `json:"retentionPeriodDays,omitempty"`
    SnapRecoveryPoints      *int `json:"snapRecoveryPoints,omitempty"`
    EnableBackupCopy        *bool `json:"enableBackupCopy,omitempty"`
    BackupCopyRPOMins       *int `json:"backupCopyRPOMins,omitempty"`
}

type MsgCreateServerPlanResponse struct {
    Plan       *MsgIdName `json:"plan,omitempty"`
    Errors[]   MsgPlanComponentErrorSet `json:"errors,omitempty"`
}

type MsgPlanComponentErrorSet struct {
    Component        *string `json:"component,omitempty"`
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateUserGroupRequest struct {
    Name               *string `json:"name,omitempty"`
    Description        *string `json:"description,omitempty"`
    EnforceFSQuota     *bool `json:"enforceFSQuota,omitempty"`
    QuotaLimitInGB     *int `json:"quotaLimitInGB,omitempty"`
}

type MsgCreateUserGroupResponse struct {
    Name     *string `json:"name,omitempty"`
    Id       *int `json:"id,omitempty"`
}

type MsgGetUserGroupDetailsResponse struct {
    ServiceType                               *string `json:"serviceType,omitempty"`
    RestrictedConsoleTypes[]                  MsgRestrictedConsoleTypesSet `json:"restrictedConsoleTypes,omitempty"`
    AssociatedLocalGroups[]                   MsgIdNameSet `json:"associatedLocalGroups,omitempty"`
    EnableTwoFactorAuthentication             *string `json:"enableTwoFactorAuthentication,omitempty"`
    LaptopAdmins                              *bool `json:"laptopAdmins,omitempty"`
    GUID                                      *string `json:"GUID,omitempty"`
    AllowMultipleCompanyMembers               *bool `json:"allowMultipleCompanyMembers,omitempty"`
    Description                               *string `json:"description,omitempty"`
    EnforceFSQuota                            *bool `json:"enforceFSQuota,omitempty"`
    QuotaLimitInGB                            *int `json:"quotaLimitInGB,omitempty"`
    EligibleToAllowMultipleCompanyMembers     *bool `json:"eligibleToAllowMultipleCompanyMembers,omitempty"`
    Enabled                                   *bool `json:"enabled,omitempty"`
    Users[]                                   MsgIdNameSet `json:"users,omitempty"`
    AssociatedEntities[]                      MsgAssocEntitySet `json:"associatedEntities,omitempty"`
    ShowAzureGuidOption                       *bool `json:"showAzureGuidOption,omitempty"`
    AzureGUID                                 *string `json:"azureGUID,omitempty"`
    Name                                      *string `json:"name,omitempty"`
    Company                                   *MsgIdName `json:"company,omitempty"`
    Id                                        *string `json:"id,omitempty"`
    Plan                                      *MsgIdName `json:"plan,omitempty"`
    Email                                     *string `json:"email,omitempty"`
    AssociatedExternalGroups[]                MsgIdNameSet `json:"associatedExternalGroups,omitempty"`
}

type MsgRestrictedConsoleTypesSet struct {
    ConsoleType     *string `json:"consoleType,omitempty"`
}

type MsgAssocEntitySet struct {
    Role               *MsgIdName `json:"role,omitempty"`
    Name               *string `json:"name,omitempty"`
    PermissionList     *MsgPermissionResp `json:"permissionList,omitempty"`
    Id                 *int `json:"id,omitempty"`
    Type               *string `json:"type,omitempty"`
}

type MsgPermissionResp struct {
    PermissionId       *int `json:"permissionId,omitempty"`
    Exclude            *bool `json:"exclude,omitempty"`
    Type               *string `json:"type,omitempty"`
    CategoryName       *string `json:"categoryName,omitempty"`
    CategoryId         *int `json:"categoryId,omitempty"`
    PermissionName     *string `json:"permissionName,omitempty"`
}

type MsgDeleteUserGroupResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgModifyUserGroupRequest struct {
    EnableTwoFactorAuthentication       *string `json:"enableTwoFactorAuthentication,omitempty"`
    LaptopAdmins                        *bool `json:"laptopAdmins,omitempty"`
    AllowMultipleCompanyMembers         *bool `json:"allowMultipleCompanyMembers,omitempty"`
    EnforceFSQuota                      *bool `json:"enforceFSQuota,omitempty"`
    QuotaLimitInGB                      *int `json:"quotaLimitInGB,omitempty"`
    ExternalUserGroupsOperationType     *string `json:"externalUserGroupsOperationType,omitempty"`
    NewDescription                      *string `json:"newDescription,omitempty"`
    Enabled                             *bool `json:"enabled,omitempty"`
    Users[]                             MsgIdNameSet `json:"users,omitempty"`
    UserOperationType                   *string `json:"userOperationType,omitempty"`
    RestrictConsoleTypes                *MsgRestrictConsoleTypes `json:"restrictConsoleTypes,omitempty"`
    NewName                             *string `json:"newName,omitempty"`
    AzureGUID                           *string `json:"azureGUID,omitempty"`
    ConsoleTypeOperationType            *string `json:"consoleTypeOperationType,omitempty"`
    PlanOperationType                   *string `json:"planOperationType,omitempty"`
    AssociatedExternalGroups[]          MsgIdNameSet `json:"associatedExternalGroups,omitempty"`
}

type MsgRestrictConsoleTypes struct {
    ConsoleType[]   string `json:"consoleType,omitempty"`
}

type MsgModifyUserGroupResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}
