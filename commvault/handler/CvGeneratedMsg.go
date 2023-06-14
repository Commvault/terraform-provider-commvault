package handler

type MsgCreateAccessPathForBucketOfCloudStorageRequest struct {
    MediaAgent     *MsgIdName `json:"mediaAgent,omitempty"`
}

type MsgIdName struct {
    Name     *string `json:"name,omitempty"`
    Id       *int `json:"id,omitempty"`
}

type MsgCreateAccessPathForBucketOfCloudStorageResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgDeleteAccessPathForBucketOfCloudStorageResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateBackupLocationRequest struct {
    MediaAgent           *MsgIdName `json:"mediaAgent,omitempty"`
    Credentials          *MsgUserNamePassword `json:"credentials,omitempty"`
    BackupLocation       *string `json:"backupLocation,omitempty"`
    SavedCredentials     *MsgIdName `json:"savedCredentials,omitempty"`
}

type MsgUserNamePassword struct {
    Password     *string `json:"password,omitempty"`
    Name         *string `json:"name,omitempty"`
}

type MsgCreateBackupLocationResponse struct {
    Name     *string `json:"name,omitempty"`
    Id       *int `json:"id,omitempty"`
}

type MsgGetBackupLocationDetailsResponse struct {
    MediaAgent           *MsgIdName `json:"mediaAgent,omitempty"`
    DiskAccessPaths[]    MsgAccessPathDetailsSet `json:"diskAccessPaths,omitempty"`
    Access               *string `json:"access,omitempty"`
    Configuration        *MsgDiskStorageConfiguration `json:"configuration,omitempty"`
    FreeSpace            *int `json:"freeSpace,omitempty"`
    Credentials          *MsgCredentialUserName `json:"credentials,omitempty"`
    BackupLocation       *string `json:"backupLocation,omitempty"`
    TotalCapacity        *int `json:"totalCapacity,omitempty"`
    Name                 *string `json:"name,omitempty"`
    SavedCredentials     *MsgIdName `json:"savedCredentials,omitempty"`
    Id                   *int `json:"id,omitempty"`
    Enabled              *bool `json:"enabled,omitempty"`
}

type MsgAccessPathDetailsSet struct {
    Path           *string `json:"path,omitempty"`
    Accessible     *bool `json:"accessible,omitempty"`
    MediaAgent     *MsgIdName `json:"mediaAgent,omitempty"`
    Access         *string `json:"access,omitempty"`
    Id             *int `json:"id,omitempty"`
    UserName       *string `json:"userName,omitempty"`
}

type MsgDiskStorageConfiguration struct {
    EnableBackupLocation                      *bool `json:"enableBackupLocation,omitempty"`
    DisableBackupLocationforFutureBackups     *bool `json:"disableBackupLocationforFutureBackups,omitempty"`
    PrepareForRetirement                      *bool `json:"prepareForRetirement,omitempty"`
}

type MsgCredentialUserName struct {
    Name     *string `json:"name,omitempty"`
}

type MsgDeleteBackupLocationResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgModifyBackupLocationRequest struct {
    MediaAgent           *MsgIdName `json:"mediaAgent,omitempty"`
    Access               *string `json:"access,omitempty"`
    Credentials          *MsgUserNamePassword `json:"credentials,omitempty"`
    Configuration        *MsgDiskStorageConfiguration `json:"configuration,omitempty"`
    BackupLocation       *string `json:"backupLocation,omitempty"`
    SavedCredentials     *MsgIdName `json:"savedCredentials,omitempty"`
    Enabled              *bool `json:"enabled,omitempty"`
}

type MsgModifyBackupLocationResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateCloudStorageAzureRequest struct {
    MediaAgent                  *MsgIdName `json:"mediaAgent,omitempty"`
    Name                        *string `json:"name,omitempty"`
    Container                   *string `json:"container,omitempty"`
    StorageClass                *string `json:"storageClass,omitempty"`
    ServiceHost                 *string `json:"serviceHost,omitempty"`
    AccountName                 *string `json:"accountName,omitempty"`
    Credentials                 *MsgIdName `json:"credentials,omitempty"`
    CloudType                   *string `json:"cloudType,omitempty"`
    Authentication              *string `json:"authentication,omitempty"`
    DeduplicationDBLocation[]   MsgDedupePathSet `json:"deduplicationDBLocation,omitempty"`
    UseDeduplication            *bool `json:"useDeduplication,omitempty"`
}

type MsgDedupePathSet struct {
    Path           *string `json:"path,omitempty"`
    MediaAgent     *MsgIdName `json:"mediaAgent,omitempty"`
}

type MsgCreateCloudStorageAzureResponse struct {
    Name             *string `json:"name,omitempty"`
    Id               *int `json:"id,omitempty"`
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgGetCloudStorageByIdResponse struct {
    Name                    *string `json:"name,omitempty"`
    Id                      *int `json:"id,omitempty"`
    Bucket[]                MsgIdNameStatusSet `json:"bucket,omitempty"`
    General                 *MsgCloudStorageGeneralInfo `json:"general,omitempty"`
    Security[]              MsgSecurityAssocSet `json:"security,omitempty"`
    Encryption              *MsgEncryption `json:"encryption,omitempty"`
    AssociatedPlans[]       MsgIdNameSet `json:"associatedPlans,omitempty"`
    CacheConfigurations     *MsgCacheConfigurations `json:"cacheConfigurations,omitempty"`
}

type MsgIdNameStatusSet struct {
    Name       *string `json:"name,omitempty"`
    Id         *int `json:"id,omitempty"`
    Status     *string `json:"status,omitempty"`
}

type MsgCloudStorageGeneralInfo struct {
    VendorType               *string `json:"vendorType,omitempty"`
    DeduplicationSavings     *string `json:"deduplicationSavings,omitempty"`
    FreeSpace                *int `json:"freeSpace,omitempty"`
    TotalCapacity            *int `json:"totalCapacity,omitempty"`
    Type                     *string `json:"type,omitempty"`
    SizeOnDisk               *int `json:"sizeOnDisk,omitempty"`
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

type MsgEncryption struct {
    Cipher          *string `json:"cipher,omitempty"`
    KeyLength       *int `json:"keyLength,omitempty"`
    Encrypt         *bool `json:"encrypt,omitempty"`
    KeyProvider     *MsgIdName `json:"keyProvider,omitempty"`
}

type MsgIdNameSet struct {
    Id     *int `json:"id,omitempty"`
}

type MsgCacheConfigurations struct {
    Enable            *bool `json:"enable,omitempty"`
    MetadataCache[]   MsgMetadataCacheSet `json:"metadataCache,omitempty"`
}

type MsgMetadataCacheSet struct {
    Path           *string `json:"path,omitempty"`
    MediaAgent     *MsgIdName `json:"mediaAgent,omitempty"`
}

type MsgDeleteCloudStorageByIdResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgModifyCloudStorageByIdRequest struct {
    Security[]     MsgUpdateSecurityAssocSet `json:"security,omitempty"`
    NewName        *string `json:"newName,omitempty"`
    Encryption     *MsgEncryption `json:"encryption,omitempty"`
}

type MsgUpdateSecurityAssocSet struct {
    Role          *MsgIdName `json:"role,omitempty"`
    User          *MsgIdName `json:"user,omitempty"`
    UserGroup     *MsgIdName `json:"userGroup,omitempty"`
}

type MsgModifyCloudStorageByIdResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateCloudStorageS3Request struct {
    MediaAgent                  *MsgIdName `json:"mediaAgent,omitempty"`
    Name                        *string `json:"name,omitempty"`
    Bucket                      *string `json:"bucket,omitempty"`
    StorageClass                *string `json:"storageClass,omitempty"`
    ServiceHost                 *string `json:"serviceHost,omitempty"`
    Credentials                 *MsgIdName `json:"credentials,omitempty"`
    CloudType                   *string `json:"cloudType,omitempty"`
    ArnRole                     *string `json:"arnRole,omitempty"`
    Authentication              *string `json:"authentication,omitempty"`
    DeduplicationDBLocation[]   MsgDedupePathSet `json:"deduplicationDBLocation,omitempty"`
    UseDeduplication            *bool `json:"useDeduplication,omitempty"`
}

type MsgCreateCloudStorageS3Response struct {
    Name             *string `json:"name,omitempty"`
    Id               *int `json:"id,omitempty"`
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateCredentialAzureRequest struct {
    VendorType      *string `json:"vendorType,omitempty"`
    AccountType     *string `json:"accountType,omitempty"`
    Name            *string `json:"name,omitempty"`
    AccessKeyId     *string `json:"accessKeyId,omitempty"`
    AccountName     *string `json:"accountName,omitempty"`
    AuthType        *string `json:"authType,omitempty"`
    Security        *MsgCredentialSecurity `json:"security,omitempty"`
    Description     *string `json:"description,omitempty"`
}

type MsgCredentialSecurity struct {
    Owner            *MsgCredentialOwner `json:"owner,omitempty"`
    Associations[]   MsgCredentialSecurityAssociationsSet `json:"associations,omitempty"`
}

type MsgCredentialOwner struct {
    User          *MsgIdName `json:"user,omitempty"`
    UserGroup     *MsgIdName `json:"userGroup,omitempty"`
}

type MsgCredentialSecurityAssociationsSet struct {
    IsCreatorAssociation     *bool `json:"isCreatorAssociation,omitempty"`
    Permissions              *MsgPermissionResp `json:"permissions,omitempty"`
    User                     *MsgIdName `json:"user,omitempty"`
    UserGroup                *MsgIdName `json:"userGroup,omitempty"`
}

type MsgPermissionResp struct {
    PermissionId       *int `json:"permissionId,omitempty"`
    Exclude            *bool `json:"exclude,omitempty"`
    Type               *string `json:"type,omitempty"`
    CategoryName       *string `json:"categoryName,omitempty"`
    CategoryId         *int `json:"categoryId,omitempty"`
    PermissionName     *string `json:"permissionName,omitempty"`
}

type MsgCreateCredentialAzureResponse struct {
    Id       *int `json:"id,omitempty"`
    Type     *string `json:"type,omitempty"`
}

type MsgGetCredentailDetailsAzureRequest struct {
    VendorType      *string `json:"vendorType,omitempty"`
    AccountType     *string `json:"accountType,omitempty"`
    Name            *string `json:"name,omitempty"`
    AccessKeyId     *string `json:"accessKeyId,omitempty"`
    AccountName     *string `json:"accountName,omitempty"`
    AuthType        *string `json:"authType,omitempty"`
    Security        *MsgCredentialSecurity `json:"security,omitempty"`
    Description     *string `json:"description,omitempty"`
}

type MsgUpdateCredentialAzureRequest struct {
    AccessKeyId     *string `json:"accessKeyId,omitempty"`
    Security        *MsgCredentialSecurity `json:"security,omitempty"`
    NewName         *string `json:"newName,omitempty"`
    AccountName     *string `json:"accountName,omitempty"`
    Description     *string `json:"description,omitempty"`
}

type MsgUpdateCredentialAzureResponse struct {
}

type MsgDeleteCredentialResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateCredentialAWSRequest struct {
    VendorType          *string `json:"vendorType,omitempty"`
    AccountType         *string `json:"accountType,omitempty"`
    Name                *string `json:"name,omitempty"`
    AccessKeyId         *string `json:"accessKeyId,omitempty"`
    SecretAccessKey     *string `json:"secretAccessKey,omitempty"`
    AuthType            *string `json:"authType,omitempty"`
    Security            *MsgCredentialSecurity `json:"security,omitempty"`
    Description         *string `json:"description,omitempty"`
}

type MsgCreateCredentialAWSResponse struct {
    Id       *int `json:"id,omitempty"`
    Type     *string `json:"type,omitempty"`
}

type MsgGetCredentailDetailsAWSRequest struct {
    VendorType          *string `json:"vendorType,omitempty"`
    AccountType         *string `json:"accountType,omitempty"`
    Name                *string `json:"name,omitempty"`
    AccessKeyId         *string `json:"accessKeyId,omitempty"`
    SecretAccessKey     *string `json:"secretAccessKey,omitempty"`
    AuthType            *string `json:"authType,omitempty"`
    Security            *MsgCredentialSecurity `json:"security,omitempty"`
    Description         *string `json:"description,omitempty"`
}

type MsgUpdateCredentialAWSRequest struct {
    AccessKeyId         *string `json:"accessKeyId,omitempty"`
    SecretAccessKey     *string `json:"secretAccessKey,omitempty"`
    Security            *MsgUpdateCredentialSecurity `json:"security,omitempty"`
    NewName             *string `json:"newName,omitempty"`
    Description         *string `json:"description,omitempty"`
}

type MsgUpdateCredentialSecurity struct {
    Owner            *MsgCredentialOwner `json:"owner,omitempty"`
    Associations[]   MsgCredentialSecurityAssociationsSet `json:"associations,omitempty"`
}

type MsgUpdateCredentialAWSResponse struct {
}

type MsgCreateCredentialAWSWithRoleArnRequest struct {
    VendorType      *string `json:"vendorType,omitempty"`
    AccountType     *string `json:"accountType,omitempty"`
    Name            *string `json:"name,omitempty"`
    Password        *string `json:"password,omitempty"`
    RoleArn         *string `json:"roleArn,omitempty"`
    AuthType        *string `json:"authType,omitempty"`
    Security        *MsgCredentialSecurity `json:"security,omitempty"`
    Description     *string `json:"description,omitempty"`
}

type MsgCreateCredentialAWSWithRoleArnResponse struct {
    Id       *int `json:"id,omitempty"`
    Type     *string `json:"type,omitempty"`
}

type MsgGetCredentailDetailsAWSWithRoleArnRequest struct {
    VendorType      *string `json:"vendorType,omitempty"`
    AccountType     *string `json:"accountType,omitempty"`
    Name            *string `json:"name,omitempty"`
    Password        *string `json:"password,omitempty"`
    RoleArn         *string `json:"roleArn,omitempty"`
    AuthType        *string `json:"authType,omitempty"`
    Security        *MsgCredentialSecurity `json:"security,omitempty"`
    Description     *string `json:"description,omitempty"`
}

type MsgUpdateCredentialAWSWithRoleArnRequest struct {
    Password        *string `json:"password,omitempty"`
    Security        *MsgUpdateCredentialSecurity `json:"security,omitempty"`
    NewName         *string `json:"newName,omitempty"`
    RoleArn         *string `json:"roleArn,omitempty"`
    Description     *string `json:"description,omitempty"`
}

type MsgUpdateCredentialAWSWithRoleArnResponse struct {
}

type MsgCreateCredentialAzureWithTenantIdRequest struct {
    VendorType            *string `json:"vendorType,omitempty"`
    AccountType           *string `json:"accountType,omitempty"`
    Name                  *string `json:"name,omitempty"`
    Environment           *string `json:"environment,omitempty"`
    Endpoints             *MsgAzureEndpoints `json:"endpoints,omitempty"`
    TenantId              *string `json:"tenantId,omitempty"`
    AuthType              *string `json:"authType,omitempty"`
    ApplicationId         *string `json:"applicationId,omitempty"`
    ApplicationSecret     *string `json:"applicationSecret,omitempty"`
    Security              *MsgCredentialSecurity `json:"security,omitempty"`
    Description           *string `json:"description,omitempty"`
}

type MsgAzureEndpoints struct {
    Storage             *string `json:"storage,omitempty"`
    ActiveDirectory     *string `json:"activeDirectory,omitempty"`
    ResourceManager     *string `json:"resourceManager,omitempty"`
}

type MsgCreateCredentialAzureWithTenantIdResponse struct {
    Id       *int `json:"id,omitempty"`
    Type     *string `json:"type,omitempty"`
}

type MsgGetCredentailDetailsAzureWithTenantIdRequest struct {
    VendorType            *string `json:"vendorType,omitempty"`
    AccountType           *string `json:"accountType,omitempty"`
    Name                  *string `json:"name,omitempty"`
    Environment           *string `json:"environment,omitempty"`
    Endpoints             *MsgAzureEndpoints `json:"endpoints,omitempty"`
    TenantId              *string `json:"tenantId,omitempty"`
    AuthType              *string `json:"authType,omitempty"`
    ApplicationId         *string `json:"applicationId,omitempty"`
    ApplicationSecret     *string `json:"applicationSecret,omitempty"`
    Security              *MsgCredentialSecurity `json:"security,omitempty"`
    Description           *string `json:"description,omitempty"`
}

type MsgUpdateCredentialAzureWithTenantIdRequest struct {
    Environment              *string `json:"environment,omitempty"`
    Endpoints                *MsgAzureEndpoints `json:"endpoints,omitempty"`
    Security                 *MsgCredentialSecurity `json:"security,omitempty"`
    NewName                  *string `json:"newName,omitempty"`
    NewApplicationSecret     *string `json:"newApplicationSecret,omitempty"`
    TenantId                 *string `json:"tenantId,omitempty"`
    Description              *string `json:"description,omitempty"`
    AuthType                 *string `json:"authType,omitempty"`
    ApplicationId            *string `json:"applicationId,omitempty"`
}

type MsgUpdateCredentialAzureWithTenantIdResponse struct {
}

type MsgGetPlanByIdResponse struct {
    Settings                  *MsgServerPlanSettings `json:"settings,omitempty"`
    BackupContent             *MsgPlanContent `json:"backupContent,omitempty"`
    DatabaseOptions           *MsgServerPlanDatabaseOptionsInfo `json:"databaseOptions,omitempty"`
    AllowPlanOverride         *bool `json:"allowPlanOverride,omitempty"`
    Workload                  *MsgPlanWorkloads `json:"workload,omitempty"`
    BackupDestinationIds[]    int `json:"backupDestinationIds,omitempty"`
    InheritSettings           *MsgServerPlanInheritSettings `json:"inheritSettings,omitempty"`
    Rpo                       *MsgServerPlanRPO `json:"rpo,omitempty"`
    AssociatedEntities[]      MsgIdNameCountSet `json:"associatedEntities,omitempty"`
    ParentInheritSettings     *MsgServerPlanInheritSettings `json:"parentInheritSettings,omitempty"`
    BackupDestinations[]      MsgPlanBackupDestinationSet `json:"backupDestinations,omitempty"`
    Permissions[]             MsgIdNameSet `json:"permissions,omitempty"`
    OverrideRestrictions      *MsgPlanOverrideSettings `json:"overrideRestrictions,omitempty"`
    SnapshotOptions           *MsgPlanSnapshotOptions `json:"snapshotOptions,omitempty"`
    AdditionalProperties      *MsgPlanAdditionalProperties `json:"additionalProperties,omitempty"`
    Plan                      *MsgIdName `json:"plan,omitempty"`
    RegionsConfigured         *bool `json:"regionsConfigured,omitempty"`
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
    UnixIncludedPaths[]                     string `json:"unixIncludedPaths,omitempty"`
    MacExcludedPaths[]                      string `json:"macExcludedPaths,omitempty"`
    MacFilterToExcludePaths[]               string `json:"macFilterToExcludePaths,omitempty"`
    MacIncludedPaths[]                      string `json:"macIncludedPaths,omitempty"`
    UnixExcludedPaths[]                     string `json:"unixExcludedPaths,omitempty"`
    UnixNumberOfDataReaders                 *MsgPlanContentDataReaders `json:"unixNumberOfDataReaders,omitempty"`
    BackupSystemState                       *bool `json:"backupSystemState,omitempty"`
    BackupSystemStateOnlyWithFullBackup     *bool `json:"backupSystemStateOnlyWithFullBackup,omitempty"`
    WindowsExcludedPaths[]                  string `json:"windowsExcludedPaths,omitempty"`
    UseVSSForSystemState                    *bool `json:"useVSSForSystemState,omitempty"`
    WindowsNumberOfDataReaders              *MsgPlanContentDataReaders `json:"windowsNumberOfDataReaders,omitempty"`
    MacNumberOfDataReaders                  *MsgPlanContentDataReaders `json:"macNumberOfDataReaders,omitempty"`
    WindowsFilterToExcludePaths[]           string `json:"windowsFilterToExcludePaths,omitempty"`
    UnixFilterToExcludePaths[]              string `json:"unixFilterToExcludePaths,omitempty"`
    ForceUpdateProperties                   *bool `json:"forceUpdateProperties,omitempty"`
}

type MsgPlanContentDataReaders struct {
    Count          *int `json:"count,omitempty"`
    UseOptimal     *bool `json:"useOptimal,omitempty"`
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
    Schedules[]       MsgPlanScheduleSet `json:"schedules,omitempty"`
    OperationType     *string `json:"operationType,omitempty"`
}

type MsgPlanScheduleSet struct {
    ScheduleName         *string `json:"scheduleName,omitempty"`
    ScheduleOption       *MsgScheduleOption `json:"scheduleOption,omitempty"`
    VmOperationType      *string `json:"vmOperationType,omitempty"`
    ForDatabasesOnly     *bool `json:"forDatabasesOnly,omitempty"`
    SchedulePattern      *MsgSchedulePattern `json:"schedulePattern,omitempty"`
    BackupType           *string `json:"backupType,omitempty"`
}

type MsgScheduleOption struct {
    DaysBetweenAutoConvert        *int `json:"daysBetweenAutoConvert,omitempty"`
    CommitFrequencyInHours        *int `json:"commitFrequencyInHours,omitempty"`
    JobRunningTimeInMins          *int `json:"jobRunningTimeInMins,omitempty"`
    O365ItemSelectionOption       *string `json:"o365ItemSelectionOption,omitempty"`
    UseDiskCacheForLogBackups     *bool `json:"useDiskCacheForLogBackups,omitempty"`
}

type MsgSchedulePattern struct {
    EndDate                       *int `json:"endDate,omitempty"`
    MaxBackupIntervalInMins       *int `json:"maxBackupIntervalInMins,omitempty"`
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
    IsSourceBackupCopy            *bool `json:"isSourceBackupCopy,omitempty"`
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
    Region                        *MsgIdName `json:"region,omitempty"`
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
    RegionToConfigure           *MsgIdName `json:"regionToConfigure,omitempty"`
    Settings                    *MsgServerPlanSettings `json:"settings,omitempty"`
    BackupContent               *MsgPlanContent `json:"backupContent,omitempty"`
    DatabaseOptions             *MsgServerPlanDatabaseOptionsInfo `json:"databaseOptions,omitempty"`
    OverrideInheritSettings     *MsgPlanOverrideInheritSetting `json:"overrideInheritSettings,omitempty"`
    FilesystemAddon             *bool `json:"filesystemAddon,omitempty"`
    AllowPlanOverride           *bool `json:"allowPlanOverride,omitempty"`
    Workload                    *MsgPlanWorkloads `json:"workload,omitempty"`
    BackupDestinationIds[]      int `json:"backupDestinationIds,omitempty"`
    Rpo                         *MsgServerPlanUpdateRPO `json:"rpo,omitempty"`
    NewName                     *string `json:"newName,omitempty"`
    OverrideRestrictions        *MsgPlanOverrideSettings `json:"overrideRestrictions,omitempty"`
    SnapshotOptions             *MsgPlanSnapshotOptions `json:"snapshotOptions,omitempty"`
}

type MsgPlanOverrideInheritSetting struct {
    Rpo                   *bool `json:"rpo,omitempty"`
    BackupContent         *bool `json:"backupContent,omitempty"`
    BackupDestination     *bool `json:"backupDestination,omitempty"`
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

type MsgModifyPlanResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateServerPlanRequest struct {
    Settings                 *MsgServerPlanSettings `json:"settings,omitempty"`
    BackupContent            *MsgPlanContent `json:"backupContent,omitempty"`
    FilesystemAddon          *bool `json:"filesystemAddon,omitempty"`
    AllowPlanOverride        *bool `json:"allowPlanOverride,omitempty"`
    PlanName                 *string `json:"planName,omitempty"`
    Workload                 *MsgPlanWorkloads `json:"workload,omitempty"`
    BackupDestinationIds[]   int `json:"backupDestinationIds,omitempty"`
    BackupDestinations[]     MsgCreatePlanBackupDestinationSet `json:"backupDestinations,omitempty"`
    OverrideRestrictions     *MsgPlanOverrideSettings `json:"overrideRestrictions,omitempty"`
    SnapshotOptions          *MsgCreatePlanSnapshotOptions `json:"snapshotOptions,omitempty"`
    ParentPlan               *MsgIdName `json:"parentPlan,omitempty"`
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
    EnableBackupCopy      *bool `json:"enableBackupCopy,omitempty"`
    BackupCopyRPOMins     *int `json:"backupCopyRPOMins,omitempty"`
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

type MsgCreateUserResponse struct {
    Users[]   MsgIdNameGUIDSet `json:"users,omitempty"`
}

type MsgIdNameGUIDSet struct {
    GUID     *string `json:"GUID,omitempty"`
    Name     *string `json:"name,omitempty"`
    Id       *int `json:"id,omitempty"`
}

type MsgGetUserDetailsResponse struct {
    ServiceType                    *string `json:"serviceType,omitempty"`
    GUID                           *string `json:"GUID,omitempty"`
    FullName                       *string `json:"fullName,omitempty"`
    Description                    *string `json:"description,omitempty"`
    AssociatedUserGroups[]         MsgIdNameProviderSet `json:"associatedUserGroups,omitempty"`
    Enabled                        *bool `json:"enabled,omitempty"`
    LockInfo                       *MsgLockProperties `json:"lockInfo,omitempty"`
    AuthenticationMethod           *string `json:"authenticationMethod,omitempty"`
    Name                           *string `json:"name,omitempty"`
    LastLoggedIn                   *int64 `json:"lastLoggedIn,omitempty"`
    Company                        *MsgIdName `json:"company,omitempty"`
    AuthenticationMethodServer     *MsgIdNameCompany `json:"authenticationMethodServer,omitempty"`
    Id                             *int `json:"id,omitempty"`
    Plan                           *MsgIdName `json:"plan,omitempty"`
    Email                          *string `json:"email,omitempty"`
    UserPrincipalName              *string `json:"userPrincipalName,omitempty"`
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

type MsgIdNameCompany struct {
    Name        *string `json:"name,omitempty"`
    Company     *MsgIdName `json:"company,omitempty"`
    Id          *int `json:"id,omitempty"`
}

type MsgDeleteUserResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgModifyUserRequest struct {
    NewName                  *string `json:"newName,omitempty"`
    AuthenticationMethod     *string `json:"authenticationMethod,omitempty"`
    FullName                 *string `json:"fullName,omitempty"`
    NewPassword              *string `json:"newPassword,omitempty"`
    Plan                     *MsgIdName `json:"plan,omitempty"`
    Email                    *string `json:"email,omitempty"`
    Enabled                  *bool `json:"enabled,omitempty"`
    UserPrincipalName        *string `json:"userPrincipalName,omitempty"`
    ValidationPassword       *string `json:"validationPassword,omitempty"`
}

type MsgModifyUserResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateBucketforCloudStorageAzureRequest struct {
    Container          *string `json:"container,omitempty"`
    MediaAgent         *MsgIdName `json:"mediaAgent,omitempty"`
    StorageClass       *string `json:"storageClass,omitempty"`
    ServiceHost        *string `json:"serviceHost,omitempty"`
    AccountName        *string `json:"accountName,omitempty"`
    Credentials        *MsgIdName `json:"credentials,omitempty"`
    CloudType          *string `json:"cloudType,omitempty"`
    Authentication     *string `json:"authentication,omitempty"`
    Password           *string `json:"password,omitempty"`
    Port               *int `json:"port,omitempty"`
    ProxyAddress       *string `json:"proxyAddress,omitempty"`
    Username           *string `json:"username,omitempty"`
}

type MsgCreateBucketforCloudStorageAzureResponse struct {
    Name             *string `json:"name,omitempty"`
    Id               *int `json:"id,omitempty"`
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgGetBucketDetailsOfCloudStorageAzureRequest struct {
    Configuration        *MsgCloudBucketConfiguration `json:"configuration,omitempty"`
    Name                 *string `json:"name,omitempty"`
    CloudAccessPaths[]   MsgCloudAccessPathsRespSet `json:"cloudAccessPaths,omitempty"`
    Id                   *int `json:"id,omitempty"`
    Container            *string `json:"container,omitempty"`
    MediaAgent           *MsgIdName `json:"mediaAgent,omitempty"`
    StorageClass         *string `json:"storageClass,omitempty"`
    ServiceHost          *string `json:"serviceHost,omitempty"`
    AccountName          *string `json:"accountName,omitempty"`
    Credentials          *MsgIdName `json:"credentials,omitempty"`
    CloudType            *string `json:"cloudType,omitempty"`
    Authentication       *string `json:"authentication,omitempty"`
    Password             *string `json:"password,omitempty"`
    Port                 *int `json:"port,omitempty"`
    ProxyAddress         *string `json:"proxyAddress,omitempty"`
    Username             *string `json:"username,omitempty"`
    Access               *string `json:"access,omitempty"`
    Enable               *bool `json:"enable,omitempty"`
}

type MsgCloudBucketConfiguration struct {
    Enable                                    *bool `json:"enable,omitempty"`
    DisableBackupLocationForFutureBackups     *bool `json:"disableBackupLocationForFutureBackups,omitempty"`
    PrepareForRetirement                      *bool `json:"prepareForRetirement,omitempty"`
    StorageAcceleratorCredentials             *MsgIdName `json:"storageAcceleratorCredentials,omitempty"`
}

type MsgCloudAccessPathsRespSet struct {
    Bucket           *string `json:"bucket,omitempty"`
    Accessible       *string `json:"accessible,omitempty"`
    MediaAgent       *MsgIdNameDisplayName `json:"mediaAgent,omitempty"`
    Access           *string `json:"access,omitempty"`
    AccessPathId     *int `json:"accessPathId,omitempty"`
    Username         *string `json:"username,omitempty"`
}

type MsgIdNameDisplayName struct {
    DisplayName     *string `json:"displayName,omitempty"`
    Name            *string `json:"name,omitempty"`
    Id              *int `json:"id,omitempty"`
}

type MsgModifyBucketOfCloudStorageAzureRequest struct {
    Container          *string `json:"container,omitempty"`
    MediaAgent         *MsgIdName `json:"mediaAgent,omitempty"`
    StorageClass       *string `json:"storageClass,omitempty"`
    ServiceHost        *string `json:"serviceHost,omitempty"`
    AccountName        *string `json:"accountName,omitempty"`
    Credentials        *MsgIdName `json:"credentials,omitempty"`
    CloudType          *string `json:"cloudType,omitempty"`
    Authentication     *string `json:"authentication,omitempty"`
    Password           *string `json:"password,omitempty"`
    Port               *int `json:"port,omitempty"`
    ProxyAddress       *string `json:"proxyAddress,omitempty"`
    Username           *string `json:"username,omitempty"`
    Access             *string `json:"access,omitempty"`
    Enable             *bool `json:"enable,omitempty"`
    Configuration      *MsgCloudBucketConfiguration `json:"configuration,omitempty"`
}

type MsgModifyBucketOfCloudStorageAzureResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgDeleteBucketOfCloudStorageResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateBucketforCloudStorageS3Request struct {
    Bucket             *string `json:"bucket,omitempty"`
    MediaAgent         *MsgIdName `json:"mediaAgent,omitempty"`
    StorageClass       *string `json:"storageClass,omitempty"`
    ServiceHost        *string `json:"serviceHost,omitempty"`
    Credentials        *MsgIdName `json:"credentials,omitempty"`
    CloudType          *string `json:"cloudType,omitempty"`
    ArnRole            *string `json:"arnRole,omitempty"`
    Authentication     *string `json:"authentication,omitempty"`
    Password           *string `json:"password,omitempty"`
    Port               *int `json:"port,omitempty"`
    ProxyAddress       *string `json:"proxyAddress,omitempty"`
    Username           *string `json:"username,omitempty"`
}

type MsgCreateBucketforCloudStorageS3Response struct {
    Name             *string `json:"name,omitempty"`
    Id               *int `json:"id,omitempty"`
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgGetBucketDetailsOfCloudStorageS3Request struct {
    Configuration        *MsgCloudBucketConfiguration `json:"configuration,omitempty"`
    Name                 *string `json:"name,omitempty"`
    CloudAccessPaths[]   MsgCloudAccessPathsRespSet `json:"cloudAccessPaths,omitempty"`
    Id                   *int `json:"id,omitempty"`
    Bucket               *string `json:"bucket,omitempty"`
    MediaAgent           *MsgIdName `json:"mediaAgent,omitempty"`
    StorageClass         *string `json:"storageClass,omitempty"`
    ServiceHost          *string `json:"serviceHost,omitempty"`
    Credentials          *MsgIdName `json:"credentials,omitempty"`
    CloudType            *string `json:"cloudType,omitempty"`
    ArnRole              *string `json:"arnRole,omitempty"`
    Authentication       *string `json:"authentication,omitempty"`
    Password             *string `json:"password,omitempty"`
    Port                 *int `json:"port,omitempty"`
    ProxyAddress         *string `json:"proxyAddress,omitempty"`
    Username             *string `json:"username,omitempty"`
    Access               *string `json:"access,omitempty"`
    Enable               *bool `json:"enable,omitempty"`
}

type MsgModifyBucketOfCloudStorageS3Request struct {
    Bucket             *string `json:"bucket,omitempty"`
    MediaAgent         *MsgIdName `json:"mediaAgent,omitempty"`
    StorageClass       *string `json:"storageClass,omitempty"`
    ServiceHost        *string `json:"serviceHost,omitempty"`
    Credentials        *MsgIdName `json:"credentials,omitempty"`
    CloudType          *string `json:"cloudType,omitempty"`
    ArnRole            *string `json:"arnRole,omitempty"`
    Authentication     *string `json:"authentication,omitempty"`
    Password           *string `json:"password,omitempty"`
    Port               *int `json:"port,omitempty"`
    ProxyAddress       *string `json:"proxyAddress,omitempty"`
    Username           *string `json:"username,omitempty"`
    Access             *string `json:"access,omitempty"`
    Enable             *bool `json:"enable,omitempty"`
    Configuration      *MsgCloudBucketConfiguration `json:"configuration,omitempty"`
}

type MsgModifyBucketOfCloudStorageS3Response struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateHypervisorAWSRequest struct {
    SkipCredentialValidation     *bool `json:"skipCredentialValidation,omitempty"`
    EtcdProtection               *MsgEtcdProtectionItem `json:"etcdProtection,omitempty"`
    Credentials                  *MsgIdName `json:"credentials,omitempty"`
    Name                         *string `json:"name,omitempty"`
    AccessNodes[]                MsgaccessNodeModelSet `json:"accessNodes,omitempty"`
    SecretKey                    *string `json:"secretKey,omitempty"`
    AccessKey                    *string `json:"accessKey,omitempty"`
    Region                       *string `json:"Region,omitempty"`
    HypervisorType               *string `json:"hypervisorType,omitempty"`
    UseServiceAccount            *string `json:"useServiceAccount,omitempty"`
    UseIamRole                   *bool `json:"useIamRole,omitempty"`
    RoleARN                      *string `json:"RoleARN,omitempty"`
    EnableAWSAdminAccount        *bool `json:"enableAWSAdminAccount,omitempty"`
}

type MsgEtcdProtectionItem struct {
    Plan        *MsgIdName `json:"plan,omitempty"`
    Enabled     *bool `json:"enabled,omitempty"`
}

type MsgaccessNodeModelSet struct {
    Id       *int `json:"id,omitempty"`
    Type     *int `json:"type,omitempty"`
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
    Credentials                  *MsgIdName `json:"credentials,omitempty"`
    AccessNodes[]                MsgaccessNodeModelSet `json:"accessNodes,omitempty"`
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
    RestoreActivityControlOptions     *MsgActivityControlOptionsProp `json:"restoreActivityControlOptions,omitempty"`
    BackupActivityControlOptions      *MsgActivityControlOptionsProp `json:"backupActivityControlOptions,omitempty"`
    EnableBackup                      *bool `json:"enableBackup,omitempty"`
    EnableRestore                     *bool `json:"enableRestore,omitempty"`
}

type MsgActivityControlOptionsProp struct {
    DelayTime              *MsgActivityControlTileDelayTime `json:"delayTime,omitempty"`
    ActivityType           *string `json:"activityType,omitempty"`
    EnableAfterADelay      *bool `json:"enableAfterADelay,omitempty"`
    EnableActivityType     *bool `json:"enableActivityType,omitempty"`
}

type MsgActivityControlTileDelayTime struct {
    TimeZone     *MsgIdName `json:"timeZone,omitempty"`
    Time         *int `json:"time,omitempty"`
    Value        *string `json:"value,omitempty"`
}

type MsghypervisorSettings struct {
    MetricsMonitoringPolicy     *MsghypervisorMonitoringPolicy `json:"metricsMonitoringPolicy,omitempty"`
    ApplicationCredentials      *MsgUserNamePassword `json:"applicationCredentials,omitempty"`
    GuestCredentials            *MsgUserNamePassword `json:"guestCredentials,omitempty"`
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
    ResourcePoolName      *string `json:"resourcePoolName,omitempty"`
    AccessNode[]          MsgaccessNodeModelSet `json:"accessNode,omitempty"`
}

type MsghypervisorCommonProps struct {
    IsRegionBasedBackup     *bool `json:"isRegionBasedBackup,omitempty"`
    IsDeconfigured          *bool `json:"isDeconfigured,omitempty"`
    RetirePhase             *string `json:"retirePhase,omitempty"`
    IsSnapBackupEnabled     *bool `json:"isSnapBackupEnabled,omitempty"`
    Company                 *MsgIdName `json:"company,omitempty"`
    IsIndexingV2            *bool `json:"isIndexingV2,omitempty"`
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
    EtcdProtection               *MsgEtcdProtectionItem `json:"etcdProtection,omitempty"`
    Credentials                  *MsgIdName `json:"credentials,omitempty"`
    Name                         *string `json:"name,omitempty"`
    AccessNodes[]                MsgaccessNodeModelSet `json:"accessNodes,omitempty"`
    ApplicationPassword          *string `json:"ApplicationPassword,omitempty"`
    TenantId                     *string `json:"tenantId,omitempty"`
    HypervisorType               *string `json:"hypervisorType,omitempty"`
    WorkloadRegion               *MsgIdName `json:"workloadRegion,omitempty"`
    SubscriptionId               *string `json:"subscriptionId,omitempty"`
    ApplicationId                *string `json:"ApplicationId,omitempty"`
    UseManagedIdentity           *bool `json:"useManagedIdentity,omitempty"`
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
    Credentials                  *MsgIdName `json:"credentials,omitempty"`
    AccessNodes[]                MsgaccessNodeModelSet `json:"accessNodes,omitempty"`
    FbrUnixMediaAgent            *MsgIdName `json:"fbrUnixMediaAgent,omitempty"`
    Password                     *string `json:"password,omitempty"`
    TenantId                     *string `json:"tenantId,omitempty"`
    ServerName                   *string `json:"serverName,omitempty"`
    HypervisorType               *string `json:"hypervisorType,omitempty"`
    SubscriptionId               *string `json:"subscriptionId,omitempty"`
    UserName                     *string `json:"userName,omitempty"`
    UseManagedIdentity           *bool `json:"useManagedIdentity,omitempty"`
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

type MsgGetKubernetesClusterDetailsResponse struct {
    DisplayName          *string `json:"displayName,omitempty"`
    Name                 *string `json:"name,omitempty"`
    Id                   *int `json:"id,omitempty"`
    ApplicationCount     *MsgKubernetesApplicationCountInfo `json:"applicationCount,omitempty"`
    ActivityControl      *MsgClusterActivityControlOptions `json:"activityControl,omitempty"`
    CommonProperties     *MsgClusterDetailsCommonProps `json:"commonProperties,omitempty"`
    Instance             *MsgIdName `json:"instance,omitempty"`
    EtcdProtection       *MsgGetEtcdProtectionItem `json:"etcdProtection,omitempty"`
    Options              *MsgEditClusterAdvancedOptionsInfo `json:"options,omitempty"`
    ApiServer            *string `json:"apiServer,omitempty"`
    ClientGroups[]       MsgIdNameSet `json:"clientGroups,omitempty"`
    Region               *MsgIdName `json:"region,omitempty"`
    Version              *string `json:"version,omitempty"`
    AccessNode[]         MsgIdNameSet `json:"accessNode,omitempty"`
    Tags[]               MsgIdNameValueSet `json:"tags,omitempty"`
}

type MsgKubernetesApplicationCountInfo struct {
    ProtectedCount             *int `json:"protectedCount,omitempty"`
    BackupUpWithErrorCount     *int `json:"backupUpWithErrorCount,omitempty"`
    PendingCount               *int `json:"pendingCount,omitempty"`
    NotProtectedCount          *int `json:"notProtectedCount,omitempty"`
    NeverBackedUpCount         *int `json:"neverBackedUpCount,omitempty"`
    TotalCount                 *int `json:"totalCount,omitempty"`
}

type MsgClusterActivityControlOptions struct {
    EnableBackup      *bool `json:"enableBackup,omitempty"`
    EnableRestore     *bool `json:"enableRestore,omitempty"`
}

type MsgClusterDetailsCommonProps struct {
    Security                 *MsgClusterDetailsCommonClientSecurity `json:"security,omitempty"`
    AssociatedUserGroups[]   MsgIdNameSet `json:"associatedUserGroups,omitempty"`
}

type MsgClusterDetailsCommonClientSecurity struct {
    Owners[]   string `json:"owners,omitempty"`
}

type MsgGetEtcdProtectionItem struct {
    Name        *string `json:"name,omitempty"`
    Id          *int `json:"id,omitempty"`
    Plan        *MsgIdName `json:"plan,omitempty"`
    Enabled     *bool `json:"enabled,omitempty"`
}

type MsgEditClusterAdvancedOptionsInfo struct {
    ImageRegistry     *MsgClusterImageRegistryOptions `json:"imageRegistry,omitempty"`
}

type MsgClusterImageRegistryOptions struct {
    RegistryUrl         *string `json:"registryUrl,omitempty"`
    ImagePullSecret     *string `json:"imagePullSecret,omitempty"`
}

type MsgIdNameValueSet struct {
    Name      *string `json:"name,omitempty"`
    Id        *int `json:"id,omitempty"`
    Value     *string `json:"value,omitempty"`
}

type MsgUpdateKubernetesPropertiesRequest struct {
    ApiServer           *string `json:"apiServer,omitempty"`
    ServiceAccount      *string `json:"serviceAccount,omitempty"`
    ServiceToken        *string `json:"serviceToken,omitempty"`
    AccessNodes[]       MsgIdNameTypeSet `json:"accessNodes,omitempty"`
    ServiceType         *string `json:"serviceType,omitempty"`
    ActivityControl     *MsgEditClusterActivityControlOptions `json:"activityControl,omitempty"`
    EtcdProtection      *MsgGetEtcdProtectionItem `json:"etcdProtection,omitempty"`
    Name                *string `json:"name,omitempty"`
    Options             *MsgEditClusterAdvancedOptionsInfo `json:"options,omitempty"`
    Region              *MsgIdName `json:"region,omitempty"`
    Tags[]              MsgNameValueSet `json:"tags,omitempty"`
}

type MsgIdNameTypeSet struct {
    Id       *int `json:"id,omitempty"`
    Type     *string `json:"type,omitempty"`
}

type MsgEditClusterActivityControlOptions struct {
    EnableBackupAfterADelay      *int `json:"enableBackupAfterADelay,omitempty"`
    EnableBackup                 *bool `json:"enableBackup,omitempty"`
    EnableRestoreAfterADelay     *int `json:"enableRestoreAfterADelay,omitempty"`
    EnableRestore                *bool `json:"enableRestore,omitempty"`
}

type MsgNameValueSet struct {
    Name      *string `json:"name,omitempty"`
    Value     *string `json:"value,omitempty"`
}

type MsgUpdateKubernetesPropertiesResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgRetireKubernetesClusterResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateKubernetesClusterOpRequest struct {
    ApiServer          *string `json:"apiServer,omitempty"`
    ServiceAccount     *string `json:"serviceAccount,omitempty"`
    ServiceToken       *string `json:"serviceToken,omitempty"`
    AccessNodes[]      MsgIdNameTypeSet `json:"accessNodes,omitempty"`
    ServiceType        *string `json:"serviceType,omitempty"`
    EtcdProtection     *MsgEtcdProtectionItem `json:"etcdProtection,omitempty"`
    Name               *string `json:"name,omitempty"`
}

type MsgCreateKubernetesClusterOpResponse struct {
    Name     *string `json:"name,omitempty"`
    Id       *int `json:"id,omitempty"`
    Etcd     *MsgIdName `json:"etcd,omitempty"`
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
    EnableLocalAuthentication                 *string `json:"enableLocalAuthentication,omitempty"`
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
    DoNotInheritRestrictConsoleTypes          *bool `json:"doNotInheritRestrictConsoleTypes,omitempty"`
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

type MsgDeleteUserGroupResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgModifyUserGroupRequest struct {
    EnableLocalAuthentication            *string `json:"enableLocalAuthentication,omitempty"`
    EnableTwoFactorAuthentication        *string `json:"enableTwoFactorAuthentication,omitempty"`
    LaptopAdmins                         *bool `json:"laptopAdmins,omitempty"`
    AllowMultipleCompanyMembers          *bool `json:"allowMultipleCompanyMembers,omitempty"`
    EnforceFSQuota                       *bool `json:"enforceFSQuota,omitempty"`
    QuotaLimitInGB                       *int `json:"quotaLimitInGB,omitempty"`
    ExternalUserGroupsOperationType      *string `json:"externalUserGroupsOperationType,omitempty"`
    NewDescription                       *string `json:"newDescription,omitempty"`
    Enabled                              *bool `json:"enabled,omitempty"`
    Users[]                              MsgIdNameSet `json:"users,omitempty"`
    UserOperationType                    *string `json:"userOperationType,omitempty"`
    RestrictConsoleTypes                 *MsgRestrictConsoleTypes `json:"restrictConsoleTypes,omitempty"`
    NewName                              *string `json:"newName,omitempty"`
    AzureGUID                            *string `json:"azureGUID,omitempty"`
    DoNotInheritRestrictConsoleTypes     *bool `json:"doNotInheritRestrictConsoleTypes,omitempty"`
    ConsoleTypeOperationType             *string `json:"consoleTypeOperationType,omitempty"`
    PlanOperationType                    *string `json:"planOperationType,omitempty"`
    AssociatedExternalGroups[]           MsgIdNameSet `json:"associatedExternalGroups,omitempty"`
}

type MsgRestrictConsoleTypes struct {
    ConsoleType[]   string `json:"consoleType,omitempty"`
}

type MsgModifyUserGroupResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgGetRoleDetailsResponse struct {
    Security[]       MsgSecurityAssocSet `json:"security,omitempty"`
    Permissions[]    MsgPermissionCategoryTypeSet `json:"permissions,omitempty"`
    Name             *string `json:"name,omitempty"`
    Id               *int `json:"id,omitempty"`
    VisibleToAll     *bool `json:"visibleToAll,omitempty"`
    Status           *string `json:"status,omitempty"`
}

type MsgPermissionCategoryTypeSet struct {
    Permission     *MsgIdName `json:"permission,omitempty"`
    Category       *MsgIdName `json:"category,omitempty"`
    Type           *string `json:"type,omitempty"`
}

type MsgDeleteRolesResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgModifyRoleRequest struct {
    PermissionOperationType     *string `json:"permissionOperationType,omitempty"`
    Security[]                  MsgUpdateSecurityAssocSet `json:"security,omitempty"`
    NewName                     *string `json:"newName,omitempty"`
    PermissionList[]            MsgPermissionsSet `json:"permissionList,omitempty"`
    VisibleToAll                *bool `json:"visibleToAll,omitempty"`
    Enabled                     *bool `json:"enabled,omitempty"`
}

type MsgPermissionsSet struct {
    Permission     *MsgIdName `json:"permission,omitempty"`
    Category       *MsgIdName `json:"category,omitempty"`
}

type MsgModifyRoleResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateNewRoleRequest struct {
    Name               *string `json:"name,omitempty"`
    PermissionList[]   MsgPermissionsSet `json:"permissionList,omitempty"`
    VisibleToAll       *bool `json:"visibleToAll,omitempty"`
    Enabled            *bool `json:"enabled,omitempty"`
}

type MsgCreateNewRoleResponse struct {
    Name     *string `json:"name,omitempty"`
    Id       *int `json:"id,omitempty"`
}

type MsgGetDiskStorageDetailsResponse struct {
    General                *MsgDiskStorageGeneralInfo `json:"general,omitempty"`
    Security[]             MsgSecurityAssocSet `json:"security,omitempty"`
    AssociatedPlanList[]   MsgIdNameSet `json:"associatedPlanList,omitempty"`
    Encryption             *MsgEncryption `json:"encryption,omitempty"`
    Name                   *string `json:"name,omitempty"`
    Id                     *int `json:"id,omitempty"`
    BackupLocations[]      MsgIdNameStatusSet `json:"backupLocations,omitempty"`
}

type MsgDiskStorageGeneralInfo struct {
    FreeSpace                *int `json:"freeSpace,omitempty"`
    SizeOnDisk               *int `json:"sizeOnDisk,omitempty"`
    DedupeSavingsPercent     *int `json:"dedupeSavingsPercent,omitempty"`
    Capacity                 *int `json:"capacity,omitempty"`
}

type MsgDeleteDiskStorageResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgModifyDiskStorageRequest struct {
    Security[]         MsgUpdateSecurityAssocSet `json:"security,omitempty"`
    NewName            *string `json:"newName,omitempty"`
    DataEncryption     *MsgEncryption `json:"dataEncryption,omitempty"`
}

type MsgModifyDiskStorageResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateDiskStorageRequest struct {
    Name                       *string `json:"name,omitempty"`
    EnableDeduplication        *bool `json:"enableDeduplication,omitempty"`
    DeduplicationDBStorage[]   MsgDedupePathSet `json:"deduplicationDBStorage,omitempty"`
    Storage[]                  MsgPathSet `json:"storage,omitempty"`
}

type MsgPathSet struct {
    MediaAgent           *MsgIdName `json:"mediaAgent,omitempty"`
    Credentials          *MsgUserNamePassword `json:"credentials,omitempty"`
    BackupLocation       *string `json:"backupLocation,omitempty"`
    SavedCredentials     *MsgIdName `json:"savedCredentials,omitempty"`
}

type MsgCreateDiskStorageResponse struct {
    Name     *string `json:"name,omitempty"`
    Id       *int `json:"id,omitempty"`
}

type MsgCreateBackupDestinationWithoutPlanInfoRequest struct {
    BackupDestinations[]   MsgCreateBackupDestinationSet `json:"backupDestinations,omitempty"`
}

type MsgCreateBackupDestinationSet struct {
    IsMirrorCopy                  *bool `json:"isMirrorCopy,omitempty"`
    RetentionPeriodDays           *int `json:"retentionPeriodDays,omitempty"`
    BackupsToCopy                 *string `json:"backupsToCopy,omitempty"`
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
    Name                          *string `json:"name,omitempty"`
    StorageType                   *string `json:"storageType,omitempty"`
    Region                        *MsgIdName `json:"region,omitempty"`
    StoragePool                   *MsgIdName `json:"storagePool,omitempty"`
}

type MsgCreateBackupDestinationWithoutPlanInfoResponse struct {
    PlanBackupDestination[]   MsgPlanBackupDestinationBasicRespSet `json:"planBackupDestination,omitempty"`
    FailedDestinations[]      string `json:"failedDestinations,omitempty"`
    Error                     *MsgGenericResp `json:"error,omitempty"`
}

type MsgPlanBackupDestinationBasicRespSet struct {
    IsDefault     *bool `json:"isDefault,omitempty"`
    Name          *string `json:"name,omitempty"`
    Id            *int `json:"id,omitempty"`
}

type MsgGenericResp struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgGetBackupDestinationDetailsWithoutPlanInfoResponse struct {
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
    IsSnapCopy                    *bool `json:"isSnapCopy,omitempty"`
    Name                          *string `json:"name,omitempty"`
    CopyType                      *string `json:"copyType,omitempty"`
    StorageType                   *string `json:"storageType,omitempty"`
    EnableDataAging               *bool `json:"enableDataAging,omitempty"`
    Id                            *int `json:"id,omitempty"`
    Region                        *MsgIdNameDisplayName `json:"region,omitempty"`
    StoragePool                   *MsgStoragePool `json:"storagePool,omitempty"`
}

type MsgDeleteBackupDestinationWithoutPlanInfoResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgModifyBackupDestinationWithoutPlanInfoRequest struct {
    RetentionPeriodDays           *int `json:"retentionPeriodDays,omitempty"`
    BackupsToCopy                 *string `json:"backupsToCopy,omitempty"`
    ExtendedRetentionRules        *MsgExtendedRetentionRules `json:"extendedRetentionRules,omitempty"`
    RetentionRuleType             *string `json:"retentionRuleType,omitempty"`
    SnapRecoveryPoints            *int `json:"snapRecoveryPoints,omitempty"`
    SourceCopy                    *MsgIdName `json:"sourceCopy,omitempty"`
    UseExtendedRetentionRules     *bool `json:"useExtendedRetentionRules,omitempty"`
    BackupStartTime               *int `json:"backupStartTime,omitempty"`
    OverrideRetentionSettings     *bool `json:"overrideRetentionSettings,omitempty"`
    Mappings[]                    MsgSnapshotCopyMappingSet `json:"mappings,omitempty"`
    Name                          *string `json:"name,omitempty"`
    EnableDataAging               *bool `json:"enableDataAging,omitempty"`
    Region                        *MsgIdName `json:"region,omitempty"`
}

type MsgModifyBackupDestinationWithoutPlanInfoResponse struct {
    PlanBackupDestination[]   MsgIdNameSet `json:"planBackupDestination,omitempty"`
    FailedDestinations[]      string `json:"failedDestinations,omitempty"`
    Error                     *MsgGenericResp `json:"error,omitempty"`
}

type MsgAddMediaAgentRequest struct {
    MediaAgents[]   MsgIdNameSet `json:"mediaAgents,omitempty"`
}

type MsgAddMediaAgentResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgDeleteDiskAccessPathResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgModifyDiskAccessPathRequest struct {
    Access      *string `json:"access,omitempty"`
    Enabled     *bool `json:"enabled,omitempty"`
}

type MsgModifyDiskAccessPathResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgGetApplicationGroupDetailsResponse struct {
    Name                          *string `json:"name,omitempty"`
    Id                            *int `json:"id,omitempty"`
    ApplicationCount              *MsgKubernetesApplicationCountInfo `json:"applicationCount,omitempty"`
    ActivityControl               *MsgApplicationGroupActivityControl `json:"activityControl,omitempty"`
    LastBackup                    *MsgLastBackupJobInfo `json:"lastBackup,omitempty"`
    IsDefaultApplicationGroup     *bool `json:"isDefaultApplicationGroup,omitempty"`
    Timezone                      *MsgIdName `json:"timezone,omitempty"`
    NextBackupTime                *int `json:"nextBackupTime,omitempty"`
    AccessNodes[]                 MsgIdNameSet `json:"accessNodes,omitempty"`
    Options                       *MsgApplicationGroupGetOptions `json:"options,omitempty"`
    IsETCDApplicationGroup        *bool `json:"isETCDApplicationGroup,omitempty"`
    Filters                       *MsgKubernetesApplicationGroupFilterItem `json:"filters,omitempty"`
    Plan                          *MsgIdName `json:"plan,omitempty"`
    Content                       *MsgKubernetesApplicationGroupContentItem `json:"content,omitempty"`
    Tags[]                        MsgIdNameValueSet `json:"tags,omitempty"`
    Cluster                       *MsgIdNameDisplayName `json:"cluster,omitempty"`
    Instance                      *MsgIdName `json:"instance,omitempty"`
    Backupset                     *MsgIdName `json:"backupset,omitempty"`
}

type MsgApplicationGroupActivityControl struct {
    EnableBackup     *bool `json:"enableBackup,omitempty"`
}

type MsgLastBackupJobInfo struct {
    JobId             *int `json:"jobId,omitempty"`
    FailureReason     *string `json:"failureReason,omitempty"`
    Time              *int64 `json:"time,omitempty"`
    Status            *string `json:"status,omitempty"`
}

type MsgApplicationGroupGetOptions struct {
    BackupStreams                      *int `json:"backupStreams,omitempty"`
    CvNamespaceScheduling              *bool `json:"cvNamespaceScheduling,omitempty"`
    WorkerResources                    *MsgApplicationGroupWorkerResourcesOptions `json:"workerResources,omitempty"`
    SnapFallbackToLiveVolumeBackup     *bool `json:"snapFallbackToLiveVolumeBackup,omitempty"`
    JobStartTime                       *int `json:"jobStartTime,omitempty"`
}

type MsgApplicationGroupWorkerResourcesOptions struct {
    MemoryRequests     *string `json:"memoryRequests,omitempty"`
    MemoryLimits       *string `json:"memoryLimits,omitempty"`
    CpuLimits          *string `json:"cpuLimits,omitempty"`
    CpuRequests        *string `json:"cpuRequests,omitempty"`
}

type MsgKubernetesApplicationGroupFilterItem struct {
    SkipStatelessApps     *bool `json:"skipStatelessApps,omitempty"`
    LabelSelectors[]      MsgKubernetesContentSelectorsSet `json:"labelSelectors,omitempty"`
    Applications[]        MsgKubernetesContentApplicationsSet `json:"applications,omitempty"`
}

type MsgKubernetesContentSelectorsSet struct {
    SelectorLevel     *string `json:"selectorLevel,omitempty"`
    SelectorValue     *string `json:"selectorValue,omitempty"`
}

type MsgKubernetesContentApplicationsSet struct {
    GUID     *string `json:"GUID,omitempty"`
    Name     *string `json:"name,omitempty"`
    Type     *string `json:"type,omitempty"`
}

type MsgKubernetesApplicationGroupContentItem struct {
    LabelSelectors[]   MsgKubernetesContentSelectorsSet `json:"labelSelectors,omitempty"`
    Applications[]     MsgKubernetesContentApplicationsSet `json:"applications,omitempty"`
}

type MsgDeleteKubernetesAppGroupResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgUpdateKubernetesAppGroupOpRequest struct {
    Name                *string `json:"name,omitempty"`
    Filters             *MsgKubernetesApplicationGroupFilters `json:"filters,omitempty"`
    Plan                *MsgIdName `json:"plan,omitempty"`
    Content             *MsgKubernetesApplicationGroupContent `json:"content,omitempty"`
    ActivityControl     *MsgApplicationGroupActivityControl `json:"activityControl,omitempty"`
    Timezone            *MsgIdName `json:"timezone,omitempty"`
    Options             *MsgApplicationGroupGetOptions `json:"options,omitempty"`
    Tags[]              MsgNameValueSet `json:"tags,omitempty"`
}

type MsgKubernetesApplicationGroupFilters struct {
    Overwrite             *bool `json:"overwrite,omitempty"`
    SkipStatelessApps     *bool `json:"skipStatelessApps,omitempty"`
    LabelSelectors[]      MsgKubernetesContentSelectorsSet `json:"labelSelectors,omitempty"`
    Applications[]        MsgKubernetesContentApplicationsSet `json:"applications,omitempty"`
}

type MsgKubernetesApplicationGroupContent struct {
    Overwrite          *bool `json:"overwrite,omitempty"`
    LabelSelectors[]   MsgKubernetesContentSelectorsSet `json:"labelSelectors,omitempty"`
    Applications[]     MsgKubernetesContentApplicationsSet `json:"applications,omitempty"`
}

type MsgUpdateKubernetesAppGroupOpResponse struct {
    ErrorMessage     *string `json:"errorMessage,omitempty"`
    ErrorCode        *int `json:"errorCode,omitempty"`
}

type MsgCreateKubernetesApplicationGroupRequest struct {
    Name        *string `json:"name,omitempty"`
    Filters     *MsgKubernetesApplicationGroupFilters `json:"filters,omitempty"`
    Plan        *MsgIdName `json:"plan,omitempty"`
    Content     *MsgKubernetesApplicationGroupContent `json:"content,omitempty"`
    Cluster     *MsgIdName `json:"cluster,omitempty"`
}

type MsgCreateKubernetesApplicationGroupResponse struct {
    Name     *string `json:"name,omitempty"`
    Id       *int `json:"id,omitempty"`
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
    LastBackup                *MsgLastBackupJobInfo `json:"lastBackup,omitempty"`
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
    Plan                     *MsgIdName `json:"plan,omitempty"`
    LastBackupTime           *int `json:"lastBackupTime,omitempty"`
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
    IsVMGroupDiskFiltersIncluded              *bool `json:"isVMGroupDiskFiltersIncluded,omitempty"`
    DatastoreFreespaceCheck                   *bool `json:"datastoreFreespaceCheck,omitempty"`
    AllowEmptySubclient                       *bool `json:"allowEmptySubclient,omitempty"`
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

type MsgvmContentSet struct {
    RuleGroups[]   MsgRuleGroupContentSet `json:"ruleGroups,omitempty"`
    Overwrite      *bool `json:"overwrite,omitempty"`
}

type MsgRuleGroupContentSet struct {
    MatchRule     *string `json:"matchRule,omitempty"`
    Rules[]       MsgRuleContentSet `json:"rules,omitempty"`
}

type MsgRuleContentSet struct {
    Condition     *string `json:"condition,omitempty"`
    Name          *string `json:"name,omitempty"`
    Type          *string `json:"type,omitempty"`
}

type MsgVMGroupCommonProperties struct {
    DataBackupStoragePolicy      *MsgIdName `json:"dataBackupStoragePolicy,omitempty"`
    IsDeletedHypervisor          *bool `json:"isDeletedHypervisor,omitempty"`
    Instance                     *MsgIdName `json:"instance,omitempty"`
    IsETCDSubclient              *bool `json:"isETCDSubclient,omitempty"`
    Hypervisor                   *MsgIdName `json:"hypervisor,omitempty"`
    Backupset                    *MsgIdName `json:"backupset,omitempty"`
    IsIndexingV2                 *bool `json:"isIndexingV2,omitempty"`
    ShowFullBackupLevel          *bool `json:"showFullBackupLevel,omitempty"`
    HypervisorType               *string `json:"hypervisorType,omitempty"`
    IndexingInfo                 *MsgindexingInfo `json:"indexingInfo,omitempty"`
    IsHypervisorDeconfigured     *bool `json:"isHypervisorDeconfigured,omitempty"`
    IDataAgent                   *MsgIdName `json:"iDataAgent,omitempty"`
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
    Copy                        *MsgAppValidationSourceCopy `json:"copy,omitempty"`
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

type MsgAppValidationSourceCopy struct {
    CopyPrecedence        *int `json:"copyPrecedence,omitempty"`
    IsSnapCopy            *bool `json:"isSnapCopy,omitempty"`
    Name                  *string `json:"name,omitempty"`
    Id                    *int `json:"id,omitempty"`
    IsActive              *bool `json:"isActive,omitempty"`
    BackupDestination     *MsgIdName `json:"backupDestination,omitempty"`
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
    TimeZone                  *MsgIdName `json:"timeZone,omitempty"`
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
    RuleGroups[]   MsgRuleGroupContentSet `json:"ruleGroups,omitempty"`
    Overwrite      *bool `json:"overwrite,omitempty"`
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
    LastBackup                *MsgLastBackupJobInfo `json:"lastBackup,omitempty"`
    CommonProperties          *MsgVMGroupCommonProperties `json:"commonProperties,omitempty"`
    SnapshotManagement        *MsgsnapCopyInfo `json:"snapshotManagement,omitempty"`
    Name                      *string `json:"name,omitempty"`
    Id                        *int `json:"id,omitempty"`
    ApplicationValidation     *MsgvmAppValidation `json:"applicationValidation,omitempty"`
    Status                    *int `json:"status,omitempty"`
    MeditechSystems           *MsgmeditechPropResp `json:"meditechSystems,omitempty"`
}

type MsgCreateVMGroupRequest struct {
    Meditech       *MsgmeditechPropResp `json:"Meditech,omitempty"`
    Hypervisor     *MsgIdName `json:"Hypervisor,omitempty"`
    Name           *string `json:"name,omitempty"`
    Storage        *MsgIdName `json:"storage,omitempty"`
    Plan           *MsgIdName `json:"plan,omitempty"`
    Content        *MsgvmContent `json:"content,omitempty"`
}

type MsgCreateVMGroupResponse struct {
    SubclientId        *int `json:"subclientId,omitempty"`
    WarningCode        *int `json:"warningCode,omitempty"`
    ErrorMessage       *string `json:"errorMessage,omitempty"`
    ErrorCode          *int `json:"errorCode,omitempty"`
    WarningMessage     *string `json:"warningMessage,omitempty"`
}
