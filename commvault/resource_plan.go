package commvault

import (
	"fmt"
	"strconv"
	"strings"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePlan() *schema.Resource {
	return &schema.Resource{
		Create: resourcePlanCreate,
		Read:   resourcePlanRead,
		Update: resourcePlanUpdate,
		Delete: resourcePlanDelete,

		Schema: map[string]*schema.Schema{
			"plan_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the Plan name used for creation of the plan.",
			},
			"retention_period_days": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Specifies the number of days that the software retains the data.",
			},
			"backup_destination_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the destination name for the backup.",
			},
			"backup_destination_storage": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the backup destination storage used for the plan.",
			},
			"company_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Specifies the companyid to which the created plan needs to be associated with.",
			},
			"rpo_in_days": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "Specifies the rpo in Days for created plan",
			},
		},
	}
}

func resourcePlanCreate(d *schema.ResourceData, m interface{}) error {

	/* 	var createPlanRequest handler.ApiCreatePlanReq
	   	createPlanRequest.PlanName = d.Get("plan_name").(string)
	   	var backupDestination handler.BackupDestination
	   	backupDestination.BackupDestinationName = d.Get("backup_destination_name").(string)
	   	backupDestination.RetentionPeriodDays = d.Get("retention_period_days").(int)
	   	backupDestination.StoragePool.Name = d.Get("backup_destination_storage").(string)
	   	createPlanRequest.BackupDestinations = append(createPlanRequest.BackupDestinations, backupDestination) */

	storagePoolResp := handler.GetStoragePools()
	storagePoolID := 0
	dedupeSPFlag := 0

	for i := range storagePoolResp.StoragePoolList {
		if strings.ToLower(storagePoolResp.StoragePoolList[i].StoragePoolEntity.StoragePoolName) == strings.ToLower(d.Get("backup_destination_storage").(string)) {
			storagePoolID = storagePoolResp.StoragePoolList[i].StoragePoolEntity.StoragePoolID
			if storagePoolResp.StoragePoolList[i].StoragePoolType == 1 {
				dedupeSPFlag = 1
			}
		}
	}

	if storagePoolID == 0 {
		return fmt.Errorf("Unable to fetch storagepoolID for provided storagePoolName: " + d.Get("backup_destination_storage").(string))
	}

	rpoInMinutes := 24 * 60 * (d.Get("rpo_in_days").(int))

	V2PlanReqStr := `{
		"plan":{
		   "inheritance":{
			  "isSealed":true
		   },
		   "storage":{
			  "enableBackupCopy":true,
			  "copy":[
				 {
					"copyType":1,
					"active":1,
					"isDefault":1,
					"extendedFlags": {
						"useGlobalStoragePolicy": "SET_TRUE"
					  },
					"dedupeFlags":{
					   "enableDASHFull":1,
					   "useGlobalDedupStore":` + strconv.Itoa(dedupeSPFlag) + `,
					   "enableDeduplication":` + strconv.Itoa(dedupeSPFlag) + `,
					   "enableClientSideDedup":` + strconv.Itoa(dedupeSPFlag) + `
					},
					"storagePolicyFlags":{
					   "blockLevelDedup":` + strconv.Itoa(dedupeSPFlag) + `
					},
					"retentionRules":{
					   "retainArchiverDataForDays":-1,
					   "retainBackupDataForDays":` + strconv.Itoa(d.Get("retention_period_days").(int)) + `,
					   "retentionFlags":{
						  "enableDataAging":1
					   }
					},
					"StoragePolicyCopy":{
					   "copyName":"` + d.Get("backup_destination_name").(string) + `"
					},
					"useGlobalPolicy":{
					   "storagePolicyName":"` + d.Get("backup_destination_storage").(string) + `",
					   "storagePolicyId":` + strconv.Itoa(storagePoolID) + `
					}
				 }
			  ]
		   },
		   "laptop":{
			  "content":{
				 "backupContent":[
					{
					   "idatype":2,
					   "subClientPolicy":{
						  "backupSetEntity":{
							 "backupsetName":"Windows subclient policy"
						  },
						  "subClientList":[
							 {
								"contentOperationType":1,
								"fsSubClientProp":{
								   "useVSSForSystemState":true,
								   "backupSystemState":true,
								   "isTrueUpOptionEnabledForFS":true,
								   "runTrueUpJobAfterDaysForFS":30,
								   "useVSS":true,
								   "keepAtLeastPreviousVersions":5,
								   "catalogACL":false,
								   "backupSystemStateforFullBkpOnly":false,
								   "scanOption":2
								},
								"content":[
								   {
									  "path":"\\"
								   }
								],
								"commonProperties":{
								   "numberOfBackupStreams":0,
								   "readBuffersize":512,
								   "storageDevice":{
									  "softwareCompression":2,
									  "deDuplicationOptions":{
										 "enableDeduplication":true,
										 "generateSignature":1
									  }
								   }
								}
							 }
						  ]
					   }
					},
					{
					   "idatype":3,
					   "subClientPolicy":{
						  "backupSetEntity":{
							 "backupsetName":"Linux subclient policy"
						  },
						  "subClientList":[
							 {
								"contentOperationType":1,
								"fsSubClientProp":{
								   "isTrueUpOptionEnabledForFS":true,
								   "runTrueUpJobAfterDaysForFS":30,
								   "useVSS":true,
								   "keepAtLeastPreviousVersions":5,
								   "catalogACL":false,
								   "scanOption":1
								},
								"content":[
								   {
									  "path":"/"
								   }
								],
								"commonProperties":{
								   "numberOfBackupStreams":0,
								   "readBuffersize":512,
								   "storageDevice":{
									  "softwareCompression":2,
									  "deDuplicationOptions":{
										 "enableDeduplication":true,
										 "generateSignature":1
									  }
								   }
								}
							 }
						  ]
					   }
					},
					{
					   "idatype":4,
					   "subClientPolicy":{
						  "backupSetEntity":{
							 "backupsetName":"Mac subclient policy"
						  },
						  "subClientList":[
							 {
								"contentOperationType":1,
								"fsSubClientProp":{
								   "isTrueUpOptionEnabledForFS":true,
								   "runTrueUpJobAfterDaysForFS":30,
								   "useVSS":true,
								   "keepAtLeastPreviousVersions":5,
								   "catalogACL":false,
								   "scanOption":1
								},
								"content":[
								   {
									  "path":"/"
								   }
								],
								"commonProperties":{
								   "numberOfBackupStreams":0,
								   "readBuffersize":512,
								   "storageDevice":{
									  "softwareCompression":2,
									  "deDuplicationOptions":{
										 "enableDeduplication":true,
										 "generateSignature":1
									  }
								   }
								}
							 }
						  ]
					   }
					}
				 ]
			  }
		   },
		   "database":{
			  "slaInMinutes":240,
			  "rpoInMinutes":240,
			  "scheduleLog":{
				 "task":{
					"description":"Automatic schedules for the log backup",
					"taskType":4,
					"policyType":0,
					"taskName":"Log schedule policy for the database clients",
					"taskFlags":{
					   "isEdgeDrive":false
					}
				 },
				 "appGroup":{
					"appGroups":[
					   {
						  "appGroupName":"APPGRP_Sybase",
						  "appGroupId":13
					   },
					   {
						  "appGroupName":"APPGRP_POSTGRES",
						  "appGroupId":114
					   },
					   {
						  "appGroupName":"APPGRP_SAP_HANA",
						  "appGroupId":129
					   },
					   {
						  "appGroupName":"APPGRP_MySql",
						  "appGroupId":100
					   },
					   {
						  "appGroupName":"APPGRP_DB2",
						  "appGroupId":11
					   },
					   {
						  "appGroupName":"APPGRP_INFORMIX",
						  "appGroupId":111
					   },
					   {
						  "appGroupName":"APPGRP_SAP_FOR_ORACLE",
						  "appGroupId":101
					   },
					   {
						  "appGroupName":"APPGRP_SQL_POLICY",
						  "appGroupId":90
					   },
					   {
						  "appGroupName":"APPGRP_AppTypeIndexedBased",
						  "appGroupId":83
					   },
					   {
						  "appGroupName":"APPGRP_NotesDb",
						  "appGroupId":12
					   },
					   {
						  "appGroupName":"PACKAGEGRP_DB",
						  "appGroupId":10
					   },
					   {
						  "appGroupName":"APPGRP_ORACLE",
						  "appGroupId":10
					   },
					   {
						  "appGroupName":"APPGRP_XchangeDB",
						  "appGroupId":14
					   },
					   {
						  "appGroupName":"APPGRP_NotesDb_Transaction_Log",
						  "appGroupId":104
					   },
					   {
						  "appGroupName":"APPGRP_DISTRIBUTEDAPPS",
						  "appGroupId":134
					   }
					]
				 },
				 "subTasks":[
					{
					   "subTask":{
						  "subTaskName":"Incremental automatic schedule for logs",
						  "subTaskType":2,
						  "flags":65536,
						  "operationType":2
					   },
					   "pattern":{
						  "freq_type":1024,
						  "timeZone":{
						  }
					   },
					   "options":{
						  "backupOpts":{
							 "truncateLogsOnSource":false,
							 "sybaseSkipFullafterLogBkp":false,
							 "backupLevel":2,
							 "incLevel":1,
							 "runIncrementalBackup":true,
							 "cumulative":false,
							 "doNotTruncateLog":false,
							 "dataOpt":{
								"useCatalogServer":true,
								"followMountPoints":true,
								"enforceTransactionLogUsage":false,
								"skipConsistencyCheck":false,
								"daysBetweenSyntheticBackup":7,
								"autoCopy":false
							 },
							 "oracleOptions":{
								"deleteArchLogOptions":{
								   "backupArchiveLogCriteria":0
								},
								"backupArchLogOptions":{
								   "backupArchiveLogCriteria":7,
								   "startLSN":1,
								   "endLSN":1,
								   "backupArchiveLog":true
								}
							 },
							 "distAppsBackupOptions":{
								"runLogBkp":true,
								"runDataBkp":false
							 }
						  },
						  "commonOpts":{
							 "perfJobOpts":{
								
							 },
							 "automaticSchedulePattern":{
								"maxBackupInterval":4,
								"sweepStartTime":86400,
								"useStorageSpaceFromMA":false,
								"minBackupIntervalMinutes":15,
								"maxBackupIntervalMinutes":0,
								"minBackupInterval":0,
								"logFileNum":{
								   "threshold":50,
								   "enabled":true
								},
								"diskUsedPercent":{
								   "threshold":80,
								   "enabled":true
								},
								"logPercent":{
								   "threshold":80,
								   "enabled":true
								}
							 }
						  }
					   }
					}
				 ]
			  }
		   },
		   "options":{
			  "quota":0
		   },
		   "summary":{
			  "description":"` + d.Get("plan_name").(string) + `",
			  "type":2,
			  "slaInterval":0,
			  "subtype":33554437,
			  "restrictions":1,
			  "rpoInMinutes":` + strconv.Itoa(rpoInMinutes) + `,
			  "addons":{
				 "database":true,
				 "filesystem":true,
				 "snap":true
			  },
			  "plan":{
				 "planName":"` + d.Get("plan_name").(string) + `"
			  },
			  "planOwner":{
				 "userGUID":"383C5C5D-578D-41C4-9C9B-8DD871243EF2",
				 "userName":"admin",
				 "userId":1
			  }
		   },
		   "schedule":{
			  "task":{
				 "taskType":4,
				 "taskFlags":{
					"isEdgeDrive":false,
					"isEZOperation":false,
					"disabled":false
				 }
			  },
			  "subTasks":[
				 {
					"subTask":{
					   "subTaskName":"Incremental backup schedule",
					   "subTaskType":2,
					   "flags":65536,
					   "operationType":2,
					   "subTaskId":1
					},
					"pattern":{
					   "active_end_occurence":0,
					   "freq_subday_interval":0,
					   "freq_type":4,
					   "description":"Every day at 9:00 PM  ",
					   "active_end_time":86340,
					   "active_start_time":75600,
					   "active_start_date":0,
					   "freq_interval":1,
					   "name":"Incremental backup schedule",
					   "freq_recurrence_factor":1,
					   "timeZone":{
						  "TimeZoneID":1001,
						  "TimeZoneName":"Client Time Zone"
					   }
					},
					"options":{
					   "backupOpts":{
						  "sybaseSkipFullafterLogBkp":false,
						  "backupLevel":2,
						  "incLevel":1,
						  "runIncrementalBackup":true,
						  "autoConvertBackupLevel":true,
						  "isSpHasInLineCopy":false,
						  "runSILOBackup":false,
						  "doNotTruncateLog":false,
						  "dataOpt":{
							 "daysBetweenSyntheticBackup":0
						  }
					   }
					}
				 },
				 {
					"subTask":{
					   "subTaskName":"Synthetic Fulls",
					   "subTaskType":2,
					   "flags":0,
					   "operationType":2,
					   "subTaskId":1
					},
					"pattern":{
					   "freq_subday_interval":0,
					   "freq_type":1024,
					   "description":"Based on automatic schedule settings",
					   "active_end_time":0,
					   "name":"Synthetic Fulls",
					   "freq_recurrence_factor":0,
					   "timeZone":{
						  
					   }
					},
					"options":{
					   "backupOpts":{
						  "truncateLogsOnSource":false,
						  "backupLevel":4,
						  "incLevel":0,
						  "runIncrementalBackup":false,
						  "doNotTruncateLog":false,
						  "vsaBackupOptions":{
							 "backupFailedVMsOnly":false
						  },
						  "dataOpt":{
							 "useCatalogServer":true,
							 "followMountPoints":true,
							 "useMultiStream":true,
							 "useMaximumStreams":true,
							 "enforceTransactionLogUsage":false,
							 "skipConsistencyCheck":false,
							 "maxNumberOfStreams":1,
							 "createNewIndex":true,
							 "daysBetweenSyntheticBackup":90,
							 "autoCopy":true
						  }
					   },
					   "commonOpts":{
						  "automaticSchedulePattern":{
							 "maxBackupInterval":72,
							 "ignoreOpWindowPastMaxInterval":true,
							 "minBackupIntervalMinutes":15,
							 "maxBackupIntervalMinutes":0,
							 "minSyncInterval":0,
							 "minBackupInterval":0,
							 "minSyncIntervalMinutes":0,
							 "stopIfOnBattery":{
								"enabled":false
							 },
							 "acPower":{
								"enabled":false
							 },
							 "specfificNetwork":{
								"enabled":false,
								"ipAddress":{
								   "subnet":24,
								   "address":"0.0.0.0"
								}
							 },
							 "stopSleepIfBackUp":{
								"enabled":false
							 },
							 "newOrModifiedFile":{
								"enabled":false
							 },
							 "wiredNetworkConnection":{
								"enabled":false
							 },
							 "minNetworkBandwidth":{
								"enabled":false
							 }
						  }
					   }
					}
				 }
			  ]
		   },
		   "snapInfo":{
			  "backupCopyRPO":240,
			  "snapTask":{
				 "task":{
					"description":"Schedule policy created automatically for plan.",
					"taskType":4,
					"isEditing":false,
					"initiatedFrom":5,
					"policyType":3,
					"taskName":"Plan Backup copy",
					"securityAssociations":{
					   
					},
					"taskSecurity":{
					   
					},
					"alert":{
					   "alertName":""
					},
					"taskFlags":{
					   "isEdgeDrive":false,
					   "disabled":false
					}
				 },
				 "appGroup":{
					
				 },
				 "subTasks":[
					{
					   "subTaskOperation":1,
					   "subTask":{
						  "subTaskName":"snap to tape",
						  "subTaskType":1,
						  "flags":33554432,
						  "operationType":4028,
						  "subTaskId":1
					   },
					   "pattern":{
						  "active_end_occurence":0,
						  "freq_type":4096,
						  "description":"Continuous Schedule After 240 minutes",
						  "active_start_time":28800,
						  "active_start_date":1594823400,
						  "freq_interval":240,
						  "freq_recurrence_factor":1,
						  "timeZone":{
							 "TimeZoneName":"CommServe Time Zone"
						  }
					   },
					   "options":{
						  "backupOpts":{
							 "backupLevel":2,
							 "dataOpt":{
								"useCatalogServer":true,
								"enforceTransactionLogUsage":true,
								"autoCopy":false
							 },
							 "mediaOpt":{
								"markMediaFullOnSuccess":false,
								"startNewMedia":false,
								"auxcopyJobOption":{
								   "autoCopy":true
								}
							 }
						  },
						  "adminOpts":{
							 "contentIndexingOption":{
								"subClientBasedAnalytics":false
							 },
							 "snapToTapeOption":{
								"allowMaximum":true,
								"noofJobsToRun":1
							 }
						  },
						  "restoreOptions":{
							 "commonOptions":{
								"syncRestore":false
							 }
						  },
						  "commonOpts":{
							 "automaticSchedulePattern":{
								"useStorageSpaceFromMA":false
							 }
						  }
					   }
					}
				 ]
			  }
		   }
		}
	 }`

	/* 	  apiResp := handler.PlanCreate(createPlanRequest, d.Get("company_id").(int))
	  if apiResp.Plan.ID > 0 {
		  d.SetId(strconv.Itoa(apiResp.Plan.ID))
		  return resourcePlanRead(d, m)
	  }
	  return fmt.Errorf("error in creation of plan") */

	apiResp := handler.PlanCreate(V2PlanReqStr, d.Get("company_id").(int))
	if apiResp.Plan.Summary.Plan.PlanID > 0 {
		d.SetId(strconv.Itoa(apiResp.Plan.Summary.Plan.PlanID))
		return resourcePlanRead(d, m)
	}
	return fmt.Errorf("error in creation of plan")
}

func resourcePlanRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePlanUpdate(d *schema.ResourceData, m interface{}) error {
	rpoinmin := d.Get("rpoinminutes").(string)
	slainminutes := d.Get("slainminutes").(string)
	id := d.Id()
	handler.PlanUpdate(rpoinmin, slainminutes, id)
	return resourcePlanRead(d, m)
}

func resourcePlanDelete(d *schema.ResourceData, m interface{}) error {
	planID := d.Id()
	genericResp := handler.PlanDelete(planID)
	if genericResp.ErrorCode != 0 {
		return fmt.Errorf("Error in deletion of plan")
	}
	d.SetId("")
	return nil
}
