package handler

import "encoding/xml"



type ApiCreatePlanReq struct {
	XMLName xml.Name `xml:"Api_CreatePlanReq"`
	Plan    struct {
		Inheritance struct {
			IsSealed string `xml:"isSealed,attr"`
		} `xml:"inheritance"`
		Storage struct {
			EnableBackupCopy string `xml:"enableBackupCopy,attr"`
			Copy             struct {
				Text           string `xml:",chardata"`
				CopyType       string `xml:"copyType,attr"`
				Active         string `xml:"active,attr"`
				IsDefault      string `xml:"isDefault,attr"`
				IsSnapCopy     string `xml:"isSnapCopy,attr"`
				CopyPrecedence string `xml:"copyPrecedence,attr"`
				DedupeFlags    struct {
					Text           string `xml:",chardata"`
					EnableDASHFull string `xml:"enableDASHFull,attr"`
				} `xml:"dedupeFlags"`
				RetentionRules struct {
					RetainArchiverDataForDays string `xml:"retainArchiverDataForDays,attr"`
					RetainBackupDataForDays   string `xml:"retainBackupDataForDays,attr"`
					RetainBackupDataForCycles string `xml:"retainBackupDataForCycles,attr"`
					RetentionFlags            struct {
						EnableDataAging string `xml:"enableDataAging,attr"`
					} `xml:"retentionFlags"`
				} `xml:"retentionRules"`
				StoragePolicyCopy struct {
					CopyName string `xml:"copyName,attr"`
				} `xml:"StoragePolicyCopy"`
				ExtendedFlags struct {
					UseGlobalStoragePolicy string `xml:"useGlobalStoragePolicy,attr"`
				} `xml:"extendedFlags"`
				UseGlobalPolicy struct {
					StoragePolicyName string `xml:"storagePolicyName,attr"`
					StoragePolicyId   string `xml:"storagePolicyId,attr"`
				} `xml:"useGlobalPolicy"`
			} `xml:"copy"`
		} `xml:"storage"`
		Laptop struct {
			Content struct {
				BackupContent []struct {
					Idatype         string `xml:"idatype,attr"`
					SubClientPolicy struct {
						BackupSetEntity struct {
							BackupsetName string `xml:"backupsetName,attr"`
						} `xml:"backupSetEntity"`
						SubClientList struct {
							ContentOperationType string `xml:"contentOperationType,attr"`
							FsSubClientProp      struct {
								UseVSSForSystemState            string `xml:"useVSSForSystemState,attr"`
								BackupSystemState               string `xml:"backupSystemState,attr"`
								UseVSS                          string `xml:"useVSS,attr"`
								KeepAtLeastPreviousVersions     string `xml:"keepAtLeastPreviousVersions,attr"`
								CatalogACL                      string `xml:"catalogACL,attr"`
								BackupSystemStateforFullBkpOnly string `xml:"backupSystemStateforFullBkpOnly,attr"`
								ScanOption                      string `xml:"scanOption,attr"`
							} `xml:"fsSubClientProp"`
							Content struct {
								Path string `xml:"path,attr"`
							} `xml:"content"`
							CommonProperties struct {
								NumberOfBackupStreams string `xml:"numberOfBackupStreams,attr"`
								ReadBuffersize        string `xml:"readBuffersize,attr"`
								StorageDevice         struct {
									SoftwareCompression  string `xml:"softwareCompression,attr"`
									DeDuplicationOptions struct {
										EnableDeduplication string `xml:"enableDeduplication,attr"`
										GenerateSignature   string `xml:"generateSignature,attr"`
									} `xml:"deDuplicationOptions"`
								} `xml:"storageDevice"`
							} `xml:"commonProperties"`
						} `xml:"subClientList"`
					} `xml:"subClientPolicy"`
				} `xml:"backupContent"`
			} `xml:"content"`
		} `xml:"laptop"`
		Database struct {
			SlaInMinutes string `xml:"slaInMinutes,attr"`
			RpoInMinutes string `xml:"rpoInMinutes,attr"`
			ScheduleLog  struct {
				Task struct {
					Description string `xml:"description,attr"`
					TaskType    string `xml:"taskType,attr"`
					PolicyType  string `xml:"policyType,attr"`
					TaskName    string `xml:"taskName,attr"`
					TaskFlags   struct {
						IsEdgeDrive string `xml:"isEdgeDrive,attr"`
					} `xml:"taskFlags"`
				} `xml:"task"`
				AppGroup struct {
					AppGroups []struct {
						AppGroupName string `xml:"appGroupName,attr"`
						AppGroupId   string `xml:"appGroupId,attr"`
					} `xml:"appGroups"`
				} `xml:"appGroup"`
				SubTasks struct {
					SubTask struct {
						SubTaskName   string `xml:"subTaskName,attr"`
						SubTaskType   string `xml:"subTaskType,attr"`
						Flags         string `xml:"flags,attr"`
						OperationType string `xml:"operationType,attr"`
					} `xml:"subTask"`
					Pattern struct {
						FreqType string `xml:"freq_type,attr"`
						TimeZone string `xml:"timeZone"`
					} `xml:"pattern"`
					Options struct {
						BackupOpts struct {
							TruncateLogsOnSource      string `xml:"truncateLogsOnSource,attr"`
							SybaseSkipFullafterLogBkp string `xml:"sybaseSkipFullafterLogBkp,attr"`
							BackupLevel               string `xml:"backupLevel,attr"`
							IncLevel                  string `xml:"incLevel,attr"`
							RunIncrementalBackup      string `xml:"runIncrementalBackup,attr"`
							Cumulative                string `xml:"cumulative,attr"`
							DoNotTruncateLog          string `xml:"doNotTruncateLog,attr"`
							DataOpt                   struct {
								Text                       string `xml:",chardata"`
								UseCatalogServer           string `xml:"useCatalogServer,attr"`
								FollowMountPoints          string `xml:"followMountPoints,attr"`
								EnforceTransactionLogUsage string `xml:"enforceTransactionLogUsage,attr"`
								SkipConsistencyCheck       string `xml:"skipConsistencyCheck,attr"`
								DaysBetweenSyntheticBackup string `xml:"daysBetweenSyntheticBackup,attr"`
								AutoCopy                   string `xml:"autoCopy,attr"`
							} `xml:"dataOpt"`
							OracleOptions struct {
								DeleteArchLogOptions struct {
									Text                     string `xml:",chardata"`
									BackupArchiveLogCriteria string `xml:"backupArchiveLogCriteria,attr"`
								} `xml:"deleteArchLogOptions"`
								BackupArchLogOptions struct {
									Text                     string `xml:",chardata"`
									BackupArchiveLogCriteria string `xml:"backupArchiveLogCriteria,attr"`
									StartLSN                 string `xml:"startLSN,attr"`
									EndLSN                   string `xml:"endLSN,attr"`
									BackupArchiveLog         string `xml:"backupArchiveLog,attr"`
								} `xml:"backupArchLogOptions"`
							} `xml:"oracleOptions"`
						} `xml:"backupOpts"`
						CommonOpts struct {
							Text                     string `xml:",chardata"`
							PerfJobOpts              string `xml:"perfJobOpts"`
							AutomaticSchedulePattern struct {
								Text                     string `xml:",chardata"`
								MaxBackupInterval        string `xml:"maxBackupInterval,attr"`
								SweepStartTime           string `xml:"sweepStartTime,attr"`
								UseStorageSpaceFromMA    string `xml:"useStorageSpaceFromMA,attr"`
								MinBackupIntervalMinutes string `xml:"minBackupIntervalMinutes,attr"`
								MaxBackupIntervalMinutes string `xml:"maxBackupIntervalMinutes,attr"`
								MinBackupInterval        string `xml:"minBackupInterval,attr"`
								LogFileNum               struct {
									Text      string `xml:",chardata"`
									Threshold string `xml:"threshold,attr"`
									Enabled   string `xml:"enabled,attr"`
								} `xml:"logFileNum"`
								DiskUsedPercent struct {
									Text      string `xml:",chardata"`
									Threshold string `xml:"threshold,attr"`
									Enabled   string `xml:"enabled,attr"`
								} `xml:"diskUsedPercent"`
								LogPercent struct {
									Text      string `xml:",chardata"`
									Threshold string `xml:"threshold,attr"`
									Enabled   string `xml:"enabled,attr"`
								} `xml:"logPercent"`
							} `xml:"automaticSchedulePattern"`
						} `xml:"commonOpts"`
					} `xml:"options"`
				} `xml:"subTasks"`
			} `xml:"scheduleLog"`
		} `xml:"database"`
		Options struct {
			Text  string `xml:",chardata"`
			Quota string `xml:"quota,attr"`
		} `xml:"options"`
		Definition struct {
			Text            string `xml:",chardata"`
			BasePlanSubtype string `xml:"basePlanSubtype,attr"`
			Possible        []struct {
				Text string `xml:",chardata"`
				Val  string `xml:"val,attr"`
			} `xml:"possible"`
			Required []struct {
				Text string `xml:",chardata"`
				Val  string `xml:"val,attr"`
			} `xml:"required"`
		} `xml:"definition"`
		Summary struct {
			Text         string `xml:",chardata"`
			Description  string `xml:"description,attr"`
			Type         string `xml:"type,attr"`
			SlaInterval  string `xml:"slaInterval,attr"`
			Subtype      string `xml:"subtype,attr"`
			Restrictions string `xml:"restrictions,attr"`
			RpoInMinutes string `xml:"rpoInMinutes,attr"`
			Addons       struct {
				Text       string `xml:",chardata"`
				Database   string `xml:"database,attr"`
				Filesystem string `xml:"filesystem,attr"`
				Snap       string `xml:"snap,attr"`
			} `xml:"addons"`
			Plan struct {
				Text     string `xml:",chardata"`
				PlanName string `xml:"planName,attr"`
			} `xml:"plan"`
			PlanOwner struct {
				Text     string `xml:",chardata"`
				UserGUID string `xml:"userGUID,attr"`
				UserName string `xml:"userName,attr"`
				UserId   string `xml:"userId,attr"`
			} `xml:"planOwner"`
		} `xml:"summary"`
		Schedule struct {
			Text string `xml:",chardata"`
			Task struct {
				Text      string `xml:",chardata"`
				TaskType  string `xml:"taskType,attr"`
				TaskFlags struct {
					Text          string `xml:",chardata"`
					IsEdgeDrive   string `xml:"isEdgeDrive,attr"`
					IsEZOperation string `xml:"isEZOperation,attr"`
					Disabled      string `xml:"disabled,attr"`
				} `xml:"taskFlags"`
			} `xml:"task"`
			SubTasks []struct {
				Text    string `xml:",chardata"`
				SubTask struct {
					Text          string `xml:",chardata"`
					SubTaskName   string `xml:"subTaskName,attr"`
					SubTaskType   string `xml:"subTaskType,attr"`
					Flags         string `xml:"flags,attr"`
					OperationType string `xml:"operationType,attr"`
					SubTaskId     string `xml:"subTaskId,attr"`
				} `xml:"subTask"`
				Pattern struct {
					Text                 string `xml:",chardata"`
					ActiveEndOccurence   string `xml:"active_end_occurence,attr"`
					FreqSubdayInterval   string `xml:"freq_subday_interval,attr"`
					FreqType             string `xml:"freq_type,attr"`
					Description          string `xml:"description,attr"`
					ActiveEndTime        string `xml:"active_end_time,attr"`
					ActiveStartTime      string `xml:"active_start_time,attr"`
					ActiveStartDate      string `xml:"active_start_date,attr"`
					FreqInterval         string `xml:"freq_interval,attr"`
					Name                 string `xml:"name,attr"`
					FreqRecurrenceFactor string `xml:"freq_recurrence_factor,attr"`
					TimeZone             struct {
						Text         string `xml:",chardata"`
						TimeZoneID   string `xml:"TimeZoneID,attr"`
						TimeZoneName string `xml:"TimeZoneName,attr"`
					} `xml:"timeZone"`
				} `xml:"pattern"`
				Options struct {
					Text       string `xml:",chardata"`
					BackupOpts struct {
						Text                      string `xml:",chardata"`
						SybaseSkipFullafterLogBkp string `xml:"sybaseSkipFullafterLogBkp,attr"`
						BackupLevel               string `xml:"backupLevel,attr"`
						IncLevel                  string `xml:"incLevel,attr"`
						RunIncrementalBackup      string `xml:"runIncrementalBackup,attr"`
						IsSpHasInLineCopy         string `xml:"isSpHasInLineCopy,attr"`
						RunSILOBackup             string `xml:"runSILOBackup,attr"`
						DoNotTruncateLog          string `xml:"doNotTruncateLog,attr"`
						TruncateLogsOnSource      string `xml:"truncateLogsOnSource,attr"`
						DataOpt                   struct {
							Text                       string `xml:",chardata"`
							DaysBetweenSyntheticBackup string `xml:"daysBetweenSyntheticBackup,attr"`
							UseCatalogServer           string `xml:"useCatalogServer,attr"`
							FollowMountPoints          string `xml:"followMountPoints,attr"`
							UseMultiStream             string `xml:"useMultiStream,attr"`
							UseMaximumStreams          string `xml:"useMaximumStreams,attr"`
							EnforceTransactionLogUsage string `xml:"enforceTransactionLogUsage,attr"`
							SkipConsistencyCheck       string `xml:"skipConsistencyCheck,attr"`
							MaxNumberOfStreams         string `xml:"maxNumberOfStreams,attr"`
							CreateNewIndex             string `xml:"createNewIndex,attr"`
							AutoCopy                   string `xml:"autoCopy,attr"`
						} `xml:"dataOpt"`
						VsaBackupOptions struct {
							Text                string `xml:",chardata"`
							BackupFailedVMsOnly string `xml:"backupFailedVMsOnly,attr"`
						} `xml:"vsaBackupOptions"`
					} `xml:"backupOpts"`
					CommonOpts struct {
						Text                     string `xml:",chardata"`
						AutomaticSchedulePattern struct {
							Text                          string `xml:",chardata"`
							MaxBackupInterval             string `xml:"maxBackupInterval,attr"`
							IgnoreOpWindowPastMaxInterval string `xml:"ignoreOpWindowPastMaxInterval,attr"`
							MinBackupIntervalMinutes      string `xml:"minBackupIntervalMinutes,attr"`
							MaxBackupIntervalMinutes      string `xml:"maxBackupIntervalMinutes,attr"`
							MinSyncInterval               string `xml:"minSyncInterval,attr"`
							MinBackupInterval             string `xml:"minBackupInterval,attr"`
							MinSyncIntervalMinutes        string `xml:"minSyncIntervalMinutes,attr"`
							StopIfOnBattery               struct {
								Text    string `xml:",chardata"`
								Enabled string `xml:"enabled,attr"`
							} `xml:"stopIfOnBattery"`
							AcPower struct {
								Text    string `xml:",chardata"`
								Enabled string `xml:"enabled,attr"`
							} `xml:"acPower"`
							SpecfificNetwork struct {
								Text      string `xml:",chardata"`
								Enabled   string `xml:"enabled,attr"`
								IpAddress struct {
									Text    string `xml:",chardata"`
									Subnet  string `xml:"subnet,attr"`
									Address string `xml:"address,attr"`
								} `xml:"ipAddress"`
							} `xml:"specfificNetwork"`
							StopSleepIfBackUp struct {
								Text    string `xml:",chardata"`
								Enabled string `xml:"enabled,attr"`
							} `xml:"stopSleepIfBackUp"`
							NewOrModifiedFile struct {
								Text    string `xml:",chardata"`
								Enabled string `xml:"enabled,attr"`
							} `xml:"newOrModifiedFile"`
							WiredNetworkConnection struct {
								Text    string `xml:",chardata"`
								Enabled string `xml:"enabled,attr"`
							} `xml:"wiredNetworkConnection"`
							MinNetworkBandwidth struct {
								Text    string `xml:",chardata"`
								Enabled string `xml:"enabled,attr"`
							} `xml:"minNetworkBandwidth"`
						} `xml:"automaticSchedulePattern"`
					} `xml:"commonOpts"`
				} `xml:"options"`
			} `xml:"subTasks"`
		} `xml:"schedule"`
		SnapInfo struct {
			Text               string `xml:",chardata"`
			SnapToTapeSchedule struct {
				Text    string `xml:",chardata"`
				SubTask struct {
					Text          string `xml:",chardata"`
					SubTaskOrder  string `xml:"subTaskOrder,attr"`
					SubTaskName   string `xml:"subTaskName,attr"`
					SubTaskType   string `xml:"subTaskType,attr"`
					Flags         string `xml:"flags,attr"`
					OperationType string `xml:"operationType,attr"`
				} `xml:"subTask"`
				Pattern struct {
					Text                 string `xml:",chardata"`
					FreqType             string `xml:"freq_type,attr"`
					ActiveStartTime      string `xml:"active_start_time,attr"`
					FreqInterval         string `xml:"freq_interval,attr"`
					Name                 string `xml:"name,attr"`
					FreqRecurrenceFactor string `xml:"freq_recurrence_factor,attr"`
					TimeZone             struct {
						Text         string `xml:",chardata"`
						TimeZoneName string `xml:"TimeZoneName,attr"`
					} `xml:"timeZone"`
				} `xml:"pattern"`
				Options struct {
					Text       string `xml:",chardata"`
					BackupOpts struct {
						Text                      string `xml:",chardata"`
						SybaseSkipFullafterLogBkp string `xml:"sybaseSkipFullafterLogBkp,attr"`
						BackupLevel               string `xml:"backupLevel,attr"`
						RunIncrementalBackup      string `xml:"runIncrementalBackup,attr"`
						IsSpHasInLineCopy         string `xml:"isSpHasInLineCopy,attr"`
						RunSILOBackup             string `xml:"runSILOBackup,attr"`
						DoNotTruncateLog          string `xml:"doNotTruncateLog,attr"`
						DataOpt                   struct {
							Text                       string `xml:",chardata"`
							DaysBetweenSyntheticBackup string `xml:"daysBetweenSyntheticBackup,attr"`
						} `xml:"dataOpt"`
						MediaOpt struct {
							Text                   string `xml:",chardata"`
							MarkMediaFullOnSuccess string `xml:"markMediaFullOnSuccess,attr"`
							StartNewMedia          string `xml:"startNewMedia,attr"`
							AuxcopyJobOption       struct {
								Text     string `xml:",chardata"`
								AutoCopy string `xml:"autoCopy,attr"`
							} `xml:"auxcopyJobOption"`
						} `xml:"mediaOpt"`
					} `xml:"backupOpts"`
					AdminOpts struct {
						Text             string `xml:",chardata"`
						SnapToTapeOption struct {
							Text         string `xml:",chardata"`
							AllowMaximum string `xml:"allowMaximum,attr"`
						} `xml:"snapToTapeOption"`
					} `xml:"adminOpts"`
				} `xml:"options"`
			} `xml:"snapToTapeSchedule"`
		} `xml:"snapInfo"`
	} `xml:"plan"`
}



type ApiCreatePlanResp struct {
	XMLName xml.Name `xml:"Api_CreatePlanResp"`
	Text    string   `xml:",chardata"`
	Plan    struct {
		Text           string `xml:",chardata"`
		DefinesStorage struct {
			Text           string `xml:",chardata"`
			DefinesEntity  string `xml:"definesEntity,attr"`
			OverrideEntity string `xml:"overrideEntity,attr"`
		} `xml:"definesStorage"`
		SecurityAssociations struct {
			Text         string `xml:",chardata"`
			Associations struct {
				Text        string `xml:",chardata"`
				UserOrGroup struct {
					Text          string `xml:",chardata"`
					UserGroupName string `xml:"userGroupName,attr"`
					UserGroupId   string `xml:"userGroupId,attr"`
					Type          string `xml:"_type_,attr"`
				} `xml:"userOrGroup"`
				Properties struct {
					Text                 string `xml:",chardata"`
					IsCreatorAssociation string `xml:"isCreatorAssociation,attr"`
					Role                 struct {
						Text     string `xml:",chardata"`
						Type     string `xml:"_type_,attr"`
						RoleId   string `xml:"roleId,attr"`
						RoleName string `xml:"roleName,attr"`
					} `xml:"role"`
				} `xml:"properties"`
			} `xml:"associations"`
			OwnerAssociations struct {
				Text            string `xml:",chardata"`
				InheritedOwners []struct {
					Text       string `xml:",chardata"`
					Permission struct {
						Text           string `xml:",chardata"`
						PermissionId   string `xml:"permissionId,attr"`
						Type           string `xml:"_type_,attr"`
						PermissionName string `xml:"permissionName,attr"`
					} `xml:"permission"`
					Parents struct {
						Text         string `xml:",chardata"`
						CommCellName string `xml:"commCellName,attr"`
						CommCellId   string `xml:"commCellId,attr"`
						Type         string `xml:"_type_,attr"`
					} `xml:"parents"`
				} `xml:"inheritedOwners"`
			} `xml:"ownerAssociations"`
		} `xml:"securityAssociations"`
		FeatureInfo struct {
			Text          string `xml:",chardata"`
			EdgedriveInfo struct {
				Text                         string `xml:",chardata"`
				EnableNotificationsForShares string `xml:"enableNotificationsForShares,attr"`
				AuditDriveActivities         string `xml:"auditDriveActivities,attr"`
				EdgeDriveAssociations        string `xml:"edgeDriveAssociations"`
			} `xml:"edgedriveInfo"`
			DefinesEdgeDriveInfo struct {
				Text           string `xml:",chardata"`
				DefinesEntity  string `xml:"definesEntity,attr"`
				OverrideEntity string `xml:"overrideEntity,attr"`
			} `xml:"definesEdgeDriveInfo"`
		} `xml:"featureInfo"`
		Inheritance struct {
			Text             string `xml:",chardata"`
			IsSealed         string `xml:"isSealed,attr"`
			EnforcedEntities []struct {
				Text string `xml:",chardata"`
				Val  string `xml:"val,attr"`
			} `xml:"enforcedEntities"`
		} `xml:"inheritance"`
		Storage struct {
			Text          string `xml:",chardata"`
			StoragePolicy struct {
				Text            string `xml:",chardata"`
				StoragePolicyId string `xml:"storagePolicyId,attr"`
			} `xml:"storagePolicy"`
		} `xml:"storage"`
		DefinesSchedule struct {
			Text           string `xml:",chardata"`
			DefinesEntity  string `xml:"definesEntity,attr"`
			OverrideEntity string `xml:"overrideEntity,attr"`
		} `xml:"definesSchedule"`
		Laptop struct {
			Text     string `xml:",chardata"`
			Features struct {
				Text               string `xml:",chardata"`
				CategoryPermission string `xml:"categoryPermission"`
			} `xml:"features"`
			AccessPolicies struct {
				Text               string `xml:",chardata"`
				CategoryPermission string `xml:"categoryPermission"`
			} `xml:"accessPolicies"`
			DefinesAccessPolicies struct {
				Text           string `xml:",chardata"`
				DefinesEntity  string `xml:"definesEntity,attr"`
				OverrideEntity string `xml:"overrideEntity,attr"`
			} `xml:"definesAccessPolicies"`
			Content struct {
				Text          string `xml:",chardata"`
				BackupContent []struct {
					Text            string `xml:",chardata"`
					Idatype         string `xml:"idatype,attr"`
					SubClientPolicy struct {
						Text            string `xml:",chardata"`
						BackupSetEntity struct {
							Text        string `xml:",chardata"`
							BackupsetId string `xml:"backupsetId,attr"`
						} `xml:"backupSetEntity"`
					} `xml:"subClientPolicy"`
				} `xml:"backupContent"`
				DefinesSubclientContentMac struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesSubclientContentMac"`
				DefinesSubclientContentLin struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesSubclientContentLin"`
				DefinesSubclientWin struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesSubclientWin"`
				DefinesSubclientArcRulesWin struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesSubclientArcRulesWin"`
				DefinesSubclientRetentionWin struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesSubclientRetentionWin"`
				DefinesSubclientArcRulesLin struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesSubclientArcRulesLin"`
				DefinesSubclientRetentionLin struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesSubclientRetentionLin"`
				DefinesSubclientContentWin struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesSubclientContentWin"`
				DefinesSubclientRetentionMac struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesSubclientRetentionMac"`
				DefinesSubclientMac struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesSubclientMac"`
				DefinesSubclientLin struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesSubclientLin"`
			} `xml:"content"`
			Users           string `xml:"users"`
			DefinesFeatures struct {
				Text           string `xml:",chardata"`
				DefinesEntity  string `xml:"definesEntity,attr"`
				OverrideEntity string `xml:"overrideEntity,attr"`
			} `xml:"definesFeatures"`
		} `xml:"laptop"`
		OperationWindow struct {
			Text   string `xml:",chardata"`
			RuleId string `xml:"ruleId,attr"`
		} `xml:"operationWindow"`
		ReplicationTargets string `xml:"replicationTargets"`
		Database           struct {
			Text         string `xml:",chardata"`
			SlaInMinutes string `xml:"slaInMinutes,attr"`
			RpoInMinutes string `xml:"rpoInMinutes,attr"`
			ScheduleLog  struct {
				Text string `xml:",chardata"`
				Task struct {
					Text   string `xml:",chardata"`
					TaskId string `xml:"taskId,attr"`
				} `xml:"task"`
			} `xml:"scheduleLog"`
			StorageLog         string `xml:"storageLog"`
			DefinesScheduleLog struct {
				Text           string `xml:",chardata"`
				DefinesEntity  string `xml:"definesEntity,attr"`
				OverrideEntity string `xml:"overrideEntity,attr"`
			} `xml:"definesScheduleLog"`
			DefinesStorageLog struct {
				Text           string `xml:",chardata"`
				DefinesEntity  string `xml:"definesEntity,attr"`
				OverrideEntity string `xml:"overrideEntity,attr"`
			} `xml:"definesStorageLog"`
		} `xml:"database"`
		Options struct {
			Text            string `xml:",chardata"`
			ForcedArchiving string `xml:"forcedArchiving,attr"`
			Quota           string `xml:"quota,attr"`
			EnableIndexing  string `xml:"enableIndexing,attr"`
			EncryptionInfo  struct {
				Text               string `xml:",chardata"`
				EncryptKeyLength   string `xml:"encryptKeyLength,attr"`
				DirectMediaAccess  string `xml:"directMediaAccess,attr"`
				CipherType         string `xml:"cipherType,attr"`
				EncryptionSettings string `xml:"encryptionSettings,attr"`
			} `xml:"encryptionInfo"`
		} `xml:"options"`
		Definition struct {
			Text     string `xml:",chardata"`
			Possible []struct {
				Text string `xml:",chardata"`
				Val  string `xml:"val,attr"`
			} `xml:"possible"`
			Required []struct {
				Text string `xml:",chardata"`
				Val  string `xml:"val,attr"`
			} `xml:"required"`
		} `xml:"definition"`
		DefinesEEPolicy struct {
			Text           string `xml:",chardata"`
			DefinesEntity  string `xml:"definesEntity,attr"`
			OverrideEntity string `xml:"overrideEntity,attr"`
		} `xml:"definesEEPolicy"`
		Summary struct {
			Text           string `xml:",chardata"`
			Type           string `xml:"type,attr"`
			SlaInterval    string `xml:"slaInterval,attr"`
			NumDevices     string `xml:"numDevices,attr"`
			Subtype        string `xml:"subtype,attr"`
			IsElastic      string `xml:"isElastic,attr"`
			SlaInMinutes   string `xml:"slaInMinutes,attr"`
			Restrictions   string `xml:"restrictions,attr"`
			PlanStatusFlag string `xml:"planStatusFlag,attr"`
			RpoInMinutes   string `xml:"rpoInMinutes,attr"`
			NumUsers       string `xml:"numUsers,attr"`
			Addons         struct {
				Text       string `xml:",chardata"`
				Database   string `xml:"database,attr"`
				Filesystem string `xml:"filesystem,attr"`
				Snap       string `xml:"snap,attr"`
			} `xml:"addons"`
			Permissions []struct {
				Text         string `xml:",chardata"`
				PermissionId string `xml:"permissionId,attr"`
			} `xml:"permissions"`
			Plan struct {
				Text        string `xml:",chardata"`
				PlanSubtype string `xml:"planSubtype,attr"`
				Type        string `xml:"_type_,attr"`
				PlanType    string `xml:"planType,attr"`
				PlanName    string `xml:"planName,attr"`
				PlanId      string `xml:"planId,attr"`
			} `xml:"plan"`
			PlanOwner struct {
				Text     string `xml:",chardata"`
				Type     string `xml:"_type_,attr"`
				UserName string `xml:"userName,attr"`
				UserId   string `xml:"userId,attr"`
			} `xml:"planOwner"`
		} `xml:"summary"`
		FullOperationWindow struct {
			Text   string `xml:",chardata"`
			RuleId string `xml:"ruleId,attr"`
		} `xml:"fullOperationWindow"`
		Alerts   string `xml:"alerts"`
		Schedule struct {
			Text string `xml:",chardata"`
			Task struct {
				Text   string `xml:",chardata"`
				TaskId string `xml:"taskId,attr"`
			} `xml:"task"`
		} `xml:"schedule"`
		IndexCopy struct {
			Text     string `xml:",chardata"`
			Schedule struct {
				Text string `xml:",chardata"`
				Task string `xml:"task"`
			} `xml:"schedule"`
		} `xml:"indexCopy"`
		Office365Info struct {
			Text         string `xml:",chardata"`
			O365OneDrive struct {
				Text            string `xml:",chardata"`
				DefinesCABackup struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesCABackup"`
				DefinesCARetention struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesCARetention"`
				DefinesCACleanup struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesCACleanup"`
			} `xml:"o365OneDrive"`
			O365Exchange struct {
				Text               string `xml:",chardata"`
				DefinesMBRetention struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesMBRetention"`
				DefinesMBCleanup struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesMBCleanup"`
				DefinesMBArchiving struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesMBArchiving"`
				DefinesMBJournal struct {
					Text           string `xml:",chardata"`
					DefinesEntity  string `xml:"definesEntity,attr"`
					OverrideEntity string `xml:"overrideEntity,attr"`
				} `xml:"definesMBJournal"`
			} `xml:"o365Exchange"`
		} `xml:"office365Info"`
		AssociatedEntitiesCount struct {
			Text           string `xml:",chardata"`
			SnapSubclients string `xml:"snapSubclients,attr"`
		} `xml:"associatedEntitiesCount"`
		EDiscoveryInfo struct {
			Text                        string `xml:",chardata"`
			DefinesContentAnalyzerCloud struct {
				Text           string `xml:",chardata"`
				DefinesEntity  string `xml:"definesEntity,attr"`
				OverrideEntity string `xml:"overrideEntity,attr"`
			} `xml:"definesContentAnalyzerCloud"`
			DefinesAnalyticsIndexServer struct {
				Text           string `xml:",chardata"`
				DefinesEntity  string `xml:"definesEntity,attr"`
				OverrideEntity string `xml:"overrideEntity,attr"`
			} `xml:"definesAnalyticsIndexServer"`
		} `xml:"eDiscoveryInfo"`
		Exchange struct {
			Text               string `xml:",chardata"`
			DefinesMBRetention struct {
				Text           string `xml:",chardata"`
				DefinesEntity  string `xml:"definesEntity,attr"`
				OverrideEntity string `xml:"overrideEntity,attr"`
			} `xml:"definesMBRetention"`
			DefinesMBCleanup struct {
				Text           string `xml:",chardata"`
				DefinesEntity  string `xml:"definesEntity,attr"`
				OverrideEntity string `xml:"overrideEntity,attr"`
			} `xml:"definesMBCleanup"`
			DefinesMBArchiving struct {
				Text           string `xml:",chardata"`
				DefinesEntity  string `xml:"definesEntity,attr"`
				OverrideEntity string `xml:"overrideEntity,attr"`
			} `xml:"definesMBArchiving"`
			DefinesMBJournal struct {
				Text           string `xml:",chardata"`
				DefinesEntity  string `xml:"definesEntity,attr"`
				OverrideEntity string `xml:"overrideEntity,attr"`
			} `xml:"definesMBJournal"`
		} `xml:"exchange"`
		DefinesCIPolicy struct {
			Text           string `xml:",chardata"`
			DefinesEntity  string `xml:"definesEntity,attr"`
			OverrideEntity string `xml:"overrideEntity,attr"`
		} `xml:"definesCIPolicy"`
	} `xml:"plan"`
	Errors struct {
		Text   string `xml:",chardata"`
		Entity struct {
			Text string `xml:",chardata"`
			Type string `xml:"_type_,attr"`
		} `xml:"entity"`
		Status struct {
			Text         string `xml:",chardata"`
			ErrorMessage string `xml:"errorMessage,attr"`
			ErrorCode    string `xml:"errorCode,attr"`
		} `xml:"status"`
	} `xml:"errors"`
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