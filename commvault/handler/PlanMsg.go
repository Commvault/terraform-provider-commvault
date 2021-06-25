package handler

import (
	"encoding/xml"
)

type ApiCreatePlanReq struct {
	PlanName           string              `json:"planName"`
	BackupDestinations []BackupDestination `json:"backupDestinations"`
}

type BackupDestination struct {
	BackupDestinationName string `json:"backupDestinationName"`
	RetentionPeriodDays   int    `json:"retentionPeriodDays"`
	StoragePool           struct {
		Name string `json:"name"`
	} `json:"storagePool"`
}

type ApiCreatePlanResp struct {
	Plan struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"plan"`
}

type V2CreatePlanResp struct {
	Plan struct {
		Summary struct {
			SLAInMinutes   int `json:"slaInMinutes"`
			Restrictions   int `json:"restrictions"`
			Type           int `json:"type"`
			PlanStatusFlag int `json:"planStatusFlag"`
			NumDevices     int `json:"numDevices"`
			Subtype        int `json:"subtype"`
			NumUsers       int `json:"numUsers"`
			Permissions    []struct {
				PermissionID int `json:"permissionId"`
			} `json:"permissions"`
			PlanOwner struct {
				Type     int    `json:"_type_"`
				UserName string `json:"userName"`
				UserID   int    `json:"userId"`
			} `json:"planOwner"`
			Plan struct {
				Type     int    `json:"_type_"`
				PlanName string `json:"planName"`
				PlanID   int    `json:"planId"`
			} `json:"plan"`
		} `json:"summary"`
		DefinesStorage struct {
			DefinesEntity  bool `json:"definesEntity"`
			OverrideEntity int  `json:"overrideEntity"`
		} `json:"definesStorage"`
		SecurityAssociations struct {
			Associations []struct {
				UserOrGroup []struct {
					UserID   int    `json:"userId"`
					Type     int    `json:"_type_"`
					UserName string `json:"userName"`
				} `json:"userOrGroup"`
				Properties struct {
					IsCreatorAssociation bool `json:"isCreatorAssociation"`
					Role                 struct {
						Type     int    `json:"_type_"`
						RoleID   int    `json:"roleId"`
						RoleName string `json:"roleName"`
					} `json:"role"`
				} `json:"properties"`
			} `json:"associations"`
			OwnerAssociations struct {
			} `json:"ownerAssociations"`
		} `json:"securityAssociations"`
		FeatureInfo struct {
			EdgedriveInfo struct {
			} `json:"edgedriveInfo"`
			DefinesEdgeDriveInfo struct {
				DefinesEntity  bool `json:"definesEntity"`
				OverrideEntity int  `json:"overrideEntity"`
			} `json:"definesEdgeDriveInfo"`
		} `json:"featureInfo"`
		Inheritance struct {
			IsSealed bool `json:"isSealed"`
		} `json:"inheritance"`
		Storage struct {
			StoragePolicy struct {
				StoragePolicyName string `json:"storagePolicyName"`
				StoragePolicyID   int    `json:"storagePolicyId"`
			} `json:"storagePolicy"`
			Copy []struct {
				CopyType       int `json:"copyType"`
				Active         int `json:"active"`
				IsDefault      int `json:"isDefault"`
				CopyPrecedence int `json:"copyPrecedence"`
				RetentionRules struct {
					RetainBackupDataForCycles int `json:"retainBackupDataForCycles"`
					Jobs                      int `json:"jobs"`
					RetainArchiverDataForDays int `json:"retainArchiverDataForDays"`
					RetainBackupDataForDays   int `json:"retainBackupDataForDays"`
				} `json:"retentionRules"`
				StoragePolicyCopy struct {
					CopyID   int    `json:"copyId"`
					CopyName string `json:"copyName"`
				} `json:"StoragePolicyCopy"`
				DrivePool struct {
					DrivePoolName string `json:"drivePoolName"`
					DrivePoolID   int    `json:"drivePoolId"`
				} `json:"drivePool"`
				Library struct {
					LibraryName string `json:"libraryName"`
					LibraryID   int    `json:"libraryId"`
				} `json:"library"`
				UseGlobalPolicy struct {
					StoragePolicyName string `json:"storagePolicyName"`
					StoragePolicyID   int    `json:"storagePolicyId"`
				} `json:"useGlobalPolicy"`
			} `json:"copy"`
		} `json:"storage"`
		DefinesSchedule struct {
			DefinesEntity  bool `json:"definesEntity"`
			OverrideEntity int  `json:"overrideEntity"`
		} `json:"definesSchedule"`
		Laptop struct {
			Features struct {
				CategoryPermission struct {
				} `json:"categoryPermission"`
			} `json:"features"`
			AccessPolicies struct {
				CategoryPermission struct {
				} `json:"categoryPermission"`
			} `json:"accessPolicies"`
			DefinesAccessPolicies struct {
				DefinesEntity  bool `json:"definesEntity"`
				OverrideEntity int  `json:"overrideEntity"`
			} `json:"definesAccessPolicies"`
			Content struct {
				DefinesSubclientWin struct {
					DefinesEntity  bool `json:"definesEntity"`
					OverrideEntity int  `json:"overrideEntity"`
				} `json:"definesSubclientWin"`
				DefinesSubclientMac struct {
					DefinesEntity  bool `json:"definesEntity"`
					OverrideEntity int  `json:"overrideEntity"`
				} `json:"definesSubclientMac"`
				DefinesSubclientLin struct {
					DefinesEntity  bool `json:"definesEntity"`
					OverrideEntity int  `json:"overrideEntity"`
				} `json:"definesSubclientLin"`
			} `json:"content"`
			Users struct {
			} `json:"users"`
			DefinesFeatures struct {
				DefinesEntity  bool `json:"definesEntity"`
				OverrideEntity int  `json:"overrideEntity"`
			} `json:"definesFeatures"`
		} `json:"laptop"`
		Alerts struct {
		} `json:"alerts"`
		ReplicationTargets struct {
		} `json:"replicationTargets"`
		Schedule struct {
			Task struct {
				Description          string `json:"description"`
				OwnerID              int    `json:"ownerId"`
				RunUserID            int    `json:"runUserId"`
				TaskType             int    `json:"taskType"`
				OwnerName            string `json:"ownerName"`
				AlertID              int    `json:"alertId"`
				GUID                 string `json:"GUID"`
				PolicyType           int    `json:"policyType"`
				AssociatedObjects    int    `json:"associatedObjects"`
				TaskName             string `json:"taskName"`
				TaskID               int    `json:"taskId"`
				SecurityAssociations struct {
					OwnerAssociations struct {
					} `json:"ownerAssociations"`
				} `json:"securityAssociations"`
				OriginalCC struct {
					CommCellID int `json:"commCellId"`
				} `json:"originalCC"`
				TaskSecurity struct {
					AssociatedUserGroups []struct {
						UserGroupID   int    `json:"userGroupId"`
						Type          int    `json:"_type_"`
						UserGroupName string `json:"userGroupName"`
					} `json:"associatedUserGroups"`
					OwnerCapabilities struct {
					} `json:"ownerCapabilities"`
				} `json:"taskSecurity"`
				CreateAs struct {
					User struct {
						User struct {
							UserName string `json:"userName"`
							UserID   int    `json:"userId"`
						} `json:"user"`
					} `json:"user"`
				} `json:"createAs"`
				TaskFlags struct {
					IsEdgeDrive   bool `json:"isEdgeDrive"`
					IsEZOperation bool `json:"isEZOperation"`
					ForDDB        bool `json:"forDDB"`
					Uninstalled   bool `json:"uninstalled"`
					IsSystem      bool `json:"isSystem"`
					Disabled      bool `json:"disabled"`
				} `json:"taskFlags"`
				Task struct {
					TaskName string `json:"taskName"`
					TaskID   int    `json:"taskId"`
				} `json:"task"`
			} `json:"task"`
			AppGroup struct {
			} `json:"appGroup"`
			SubTasks []struct {
				SubTask struct {
					SubTaskOrder  int    `json:"subTaskOrder"`
					SubTaskName   string `json:"subTaskName"`
					SubTaskType   int    `json:"subTaskType"`
					Flags         int    `json:"flags"`
					OperationType int    `json:"operationType"`
					SubTaskID     int    `json:"subTaskId"`
					SubTask       struct {
						SubtaskID   int    `json:"subtaskId"`
						SubtaskName string `json:"subtaskName"`
					} `json:"subTask"`
				} `json:"subTask"`
				Options struct {
					BackupOpts struct {
						BkpLatestVersion     bool `json:"bkpLatestVersion"`
						BackupLevel          int  `json:"backupLevel"`
						IncLevel             int  `json:"incLevel"`
						RunIncrementalBackup bool `json:"runIncrementalBackup"`
						DoNotTruncateLog     bool `json:"doNotTruncateLog"`
						CdrOptions           struct {
							Incremental          bool `json:"incremental"`
							DataVerificationOnly bool `json:"dataVerificationOnly"`
							Full                 bool `json:"full"`
						} `json:"cdrOptions"`
						DataOpt struct {
							StopWinService             bool `json:"stopWinService"`
							StopDhcpService            bool `json:"stopDhcpService"`
							UseCatalogServer           bool `json:"useCatalogServer"`
							OptimizedBackup            bool `json:"optimizedBackup"`
							FollowMountPoints          bool `json:"followMountPoints"`
							BkpFilesProctedByFS        bool `json:"bkpFilesProctedByFS"`
							Granularrecovery           bool `json:"granularrecovery"`
							VerifySynthFull            bool `json:"verifySynthFull"`
							DaysBetweenSyntheticBackup int  `json:"daysBetweenSyntheticBackup"`
						} `json:"dataOpt"`
						NasOptions struct {
							SnapShotType int  `json:"snapShotType"`
							BackupQuotas bool `json:"backupQuotas"`
						} `json:"nasOptions"`
						VaultTrackerOpt struct {
							MediaStatus struct {
								Bad                bool `json:"bad"`
								OverwriteProtected bool `json:"overwriteProtected"`
								Full               bool `json:"full"`
							} `json:"mediaStatus"`
						} `json:"vaultTrackerOpt"`
						MediaOpt struct {
							NumberofDays                     int  `json:"numberofDays"`
							RetentionJobType                 int  `json:"retentionJobType"`
							WaitForInlineBackupResources     bool `json:"waitForInlineBackupResources"`
							AllowOtherSchedulesToUseMediaSet bool `json:"allowOtherSchedulesToUseMediaSet"`
						} `json:"mediaOpt"`
					} `json:"backupOpts"`
					CommonOpts struct {
						JobRetryOpts struct {
							RunningTime struct {
								TotalRunningTime int `json:"totalRunningTime"`
							} `json:"runningTime"`
						} `json:"jobRetryOpts"`
					} `json:"commonOpts"`
				} `json:"options"`
				Pattern struct {
					ActiveEndOccurence   int    `json:"active_end_occurence"`
					FreqSubdayInterval   int    `json:"freq_subday_interval"`
					FreqType             int    `json:"freq_type"`
					PatternID            int    `json:"patternId"`
					Flags                int    `json:"flags"`
					Description          string `json:"description"`
					ActiveEndTime        int    `json:"active_end_time"`
					ActiveEndDate        int    `json:"active_end_date"`
					SkipOccurence        int    `json:"skipOccurence"`
					SkipDayNumber        int    `json:"skipDayNumber"`
					ActiveStartTime      int    `json:"active_start_time"`
					FreqRestartInterval  int    `json:"freq_restart_interval"`
					ActiveStartDate      int    `json:"active_start_date"`
					FreqInterval         int    `json:"freq_interval"`
					FreqRelativeInterval int    `json:"freq_relative_interval"`
					Name                 string `json:"name"`
					FreqRecurrenceFactor int    `json:"freq_recurrence_factor"`
					DaysToRun            struct {
						Week   int  `json:"week"`
						Friday bool `json:"Friday"`
						Day    int  `json:"day"`
					} `json:"daysToRun"`
					RepeatPattern []struct {
						Exception   bool   `json:"exception"`
						OnDayNumber int    `json:"onDayNumber"`
						OnDay       int    `json:"onDay"`
						Description string `json:"description"`
						Occurrence  int    `json:"occurrence"`
						RepeatOn    int    `json:"repeatOn"`
					} `json:"repeatPattern"`
					Calendar struct {
						CalendarName string `json:"calendarName"`
						CalendarID   int    `json:"calendarId"`
					} `json:"calendar"`
					TimeZone struct {
						TimeZoneID int `json:"TimeZoneID"`
					} `json:"timeZone"`
				} `json:"pattern,omitempty"`
			} `json:"subTasks"`
		} `json:"schedule"`
		Database struct {
			SLAInMinutes int `json:"slaInMinutes"`
			ScheduleLog  struct {
				Task struct {
				} `json:"task"`
			} `json:"scheduleLog"`
			StorageLog struct {
				StoragePolicy struct {
				} `json:"storagePolicy"`
			} `json:"storageLog"`
			DefinesScheduleLog struct {
				DefinesEntity  bool `json:"definesEntity"`
				OverrideEntity int  `json:"overrideEntity"`
			} `json:"definesScheduleLog"`
			DefinesStorageLog struct {
				DefinesEntity  bool `json:"definesEntity"`
				OverrideEntity int  `json:"overrideEntity"`
			} `json:"definesStorageLog"`
		} `json:"database"`
		DefinesAlerts struct {
			DefinesEntity  bool `json:"definesEntity"`
			OverrideEntity int  `json:"overrideEntity"`
		} `json:"definesAlerts"`
		EDiscoveryInfo struct {
			DefinesContentAnalyzerCloud struct {
				DefinesEntity  bool `json:"definesEntity"`
				OverrideEntity int  `json:"overrideEntity"`
			} `json:"definesContentAnalyzerCloud"`
			DefinesAnalyticsEngineCloud struct {
				DefinesEntity  bool `json:"definesEntity"`
				OverrideEntity int  `json:"overrideEntity"`
			} `json:"definesAnalyticsEngineCloud"`
		} `json:"eDiscoveryInfo"`
		Options struct {
			ForcedArchiving bool `json:"forcedArchiving"`
			Quota           int  `json:"quota"`
		} `json:"options"`
		Exchange struct {
			DefinesMBRetention struct {
				DefinesEntity  bool `json:"definesEntity"`
				OverrideEntity int  `json:"overrideEntity"`
			} `json:"definesMBRetention"`
			DefinesMBCleanup struct {
				DefinesEntity  bool `json:"definesEntity"`
				OverrideEntity int  `json:"overrideEntity"`
			} `json:"definesMBCleanup"`
			DefinesMBArchiving struct {
				DefinesEntity  bool `json:"definesEntity"`
				OverrideEntity int  `json:"overrideEntity"`
			} `json:"definesMBArchiving"`
			DefinesMBJournal struct {
				DefinesEntity  bool `json:"definesEntity"`
				OverrideEntity int  `json:"overrideEntity"`
			} `json:"definesMBJournal"`
		} `json:"exchange"`
		Definition struct {
			Possible []int `json:"possible"`
			Required []int `json:"required"`
		} `json:"definition"`
	} `json:"plan"`
	Errors []struct {
		Entity struct {
		} `json:"entity"`
		Status struct {
			ErrorMessage string `json:"errorMessage"`
			ErrorCode    int    `json:"errorCode"`
		} `json:"status"`
	} `json:"errors"`
}

type ApiUpdatePlanReq struct {
	XMLName  xml.Name `xml:"Api_UpdatePlanReq"`
	Database struct {
		SlaInMinutes string `xml:"slaInMinutes,attr"`
		RpoInMinutes string `xml:"rpoInMinutes,attr"`
	} `xml:"database"`
	Summary struct {
		Addons struct {
			Database string `xml:"database,attr"`
		} `xml:"addons"`
		Plan struct {
			PlanId string `xml:"planId,attr"`
		} `xml:"plan"`
	} `xml:"summary"`
}

type GenericResp struct {
	ErrorString string `json:"errorString"`
	ErrorCode   int    `json:"errorCode"`
}

type AppAssociateVMToPlanRequest struct {
	XMLName xml.Name `xml:"App_AssociateVMToPlanRequest"`
	VmInfo  struct {
		Plan struct {
			PlanSubtype string `xml:"planSubtype,attr"`
			PlanType    string `xml:"planType,attr"`
			PlanSummary string `xml:"planSummary,attr"`
			PlanName    string `xml:"planName,attr"`
			PlanId      string `xml:"planId,attr"`
		} `xml:"plan"`
		VmClients struct {
			ClientId   string `xml:"clientId,attr"`
			ClientName string `xml:"clientName,attr"`
			ClientGUID string `xml:"clientGUID,attr"`
		} `xml:"vmClients"`
	} `xml:"vmInfo"`
}

type AppAssociateVMToPlanResponse struct {
	XMLName  xml.Name `xml:"App_AssociateVMToPlanResponse"`
	Response struct {
		ErrorString string `xml:"errorString,attr"`
		ErrorCode   string `xml:"errorCode,attr"`
	} `xml:"response"`
}

type GetStoragePoolListResp struct {
	StoragePoolList []struct {
		NumberOfNodes     int    `json:"numberOfNodes"`
		TotalFreeSpace    int64  `json:"totalFreeSpace"`
		StoragePoolType   int    `json:"storagePoolType"`
		TotalCapacity     int64  `json:"totalCapacity"`
		Reserved1         int    `json:"reserved1"`
		Status            string `json:"status"`
		StatusCode        int    `json:"statusCode"`
		StoragePoolEntity struct {
			StoragePoolName string `json:"storagePoolName"`
			Type            int    `json:"_type_"`
			StoragePoolID   int    `json:"storagePoolId"`
			EntityInfo      struct {
				CompanyID       int    `json:"companyId"`
				CompanyName     string `json:"companyName"`
				MultiCommcellID int    `json:"multiCommcellId"`
			} `json:"entityInfo"`
		} `json:"storagePoolEntity"`
		StoragePool struct {
			Type            int    `json:"_type_"`
			ClientGroupID   int    `json:"clientGroupId"`
			ClientGroupName string `json:"clientGroupName"`
			EntityInfo      struct {
				CompanyID       int    `json:"companyId"`
				CompanyName     string `json:"companyName"`
				MultiCommcellID int    `json:"multiCommcellId"`
			} `json:"entityInfo"`
		} `json:"storagePool"`
	} `json:"storagePoolList"`
}
