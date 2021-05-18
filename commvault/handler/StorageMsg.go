package handler

type CreateStorageReq struct {
	StoragePolicyName     string `json:"storagePolicyName"`
	CopyName              string `json:"copyName"`
	Type                  string `json:"type"`
	NumberOfCopies        int    `json:"numberOfCopies"`
	StoragePolicyCopyInfo struct {
		CopyType           string `json:"copyType"`
		IsDefault          string `json:"isDefault"`
		Active             string `json:"active"`
		StoragePolicyFlags struct {
			BlockLevelDedup           string `json:"blockLevelDedup"`
			EnableGlobalDeduplication string `json:"enableGlobalDeduplication"`
			GlobalStoragePolicy       string `json:"globalStoragePolicy"`
		} `json:"storagePolicyFlags"`
		CopyFlags struct {
			PreserveEncryptionModeAsInSource string `json:"preserveEncryptionModeAsInSource"`
		} `json:"copyFlags"`
		ExtendedFlags struct {
			GlobalStoragePolicy string `json:"globalStoragePolicy"`
		} `json:"extendedFlags"`
		Library struct {
			LibraryID   int    `json:"libraryId"`
			LibraryName string `json:"libraryName"`
		} `json:"library"`
		MediaAgent struct {
			MediaAgentID   int    `json:"mediaAgentId"`
			MediaAgentName string `json:"mediaAgentName"`
		} `json:"mediaAgent"`
		RetentionRules struct {
			RetentionFlags struct {
				EnableDataAging string `json:"enableDataAging"`
			} `json:"retentionFlags"`
			RetainBackupDataForDays   int `json:"retainBackupDataForDays"`
			RetainBackupDataForCycles int `json:"retainBackupDataForCycles"`
			RetainArchiverDataForDays int `json:"retainArchiverDataForDays"`
		} `json:"retentionRules"`
		IsFromGui                bool `json:"isFromGui"`
		NumberOfStreamsToCombine int  `json:"numberOfStreamsToCombine"`
		DedupeFlags              struct {
			EnableDeduplication  string `json:"enableDeduplication"`
			EnableDASHFull       string `json:"enableDASHFull"`
			HostGlobalDedupStore string `json:"hostGlobalDedupStore"`
		} `json:"dedupeFlags"`
		DDBPartitionInfo struct {
			MaInfoList    []MaInfoList `json:"maInfoList"`
			SidbStoreInfo struct {
				NumSIDBStore int `json:"numSIDBStore"`
			} `json:"sidbStoreInfo"`
		} `json:"DDBPartitionInfo"`
	} `json:"storagePolicyCopyInfo"`
	ClientGroup struct {
		ClientGroupID   int    `json:"clientGroupId"`
		ClientGroupName string `json:"clientGroupName"`
	} `json:"clientGroup"`
	Storage []Storage `json:"storage"`
}

type Storage struct {
	MediaAgent struct {
		MediaAgentID   int    `json:"mediaAgentId"`
		MediaAgentName string `json:"mediaAgentName"`
	} `json:"mediaAgent"`
	Path        string `json:"path"`
	Credentials struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	} `json:"credentials"`
	SavedCredential struct {
		CredentialID   int    `json:"credentialId"`
		CredentialName string `json:"credentialName"`
	} `json:"savedCredential"`
	DeviceType          int `json:"deviceType"`
	MetallicStorageInfo struct {
		StorageClass []string `json:"storageClass"`
		Replication  []string `json:"replication"`
	} `json:"metallicStorageInfo"`
}

type SubStoreList struct {
	AccessPath struct {
		Path string `json:"path"`
	} `json:"accessPath"`
	DiskFreeThresholdMB         int `json:"diskFreeThresholdMB"`
	DiskFreeWarningThreshholdMB int `json:"diskFreeWarningThreshholdMB"`
}

type MaInfoList struct {
	MediaAgent struct {
		MediaAgentName string `json:"mediaAgentName"`
	} `json:"mediaAgent"`
	SubStoreList []SubStoreList `json:"subStoreList"`
}

type CreateStorageResp struct {
	ResponseType     int `json:"responseType"`
	ArchiveGroupCopy struct {
		CopyID            int    `json:"copyId"`
		Type              int    `json:"_type_"`
		CopyName          string `json:"copyName"`
		StoragePolicyName string `json:"storagePolicyName"`
		StoragePolicyID   int    `json:"storagePolicyId"`
	} `json:"archiveGroupCopy"`
	Error struct {
		ErrorMessage string `json:"errorMessage"`
		ErrorCode    int    `json:"errorCode"`
	} `json:"error"`
}
