package commvault

import (
	"fmt"
	"strconv"
	"strings"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePlan_Server() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreatePlan_Server,
		Read:   resourceReadPlan_Server,
		Update: resourceUpdatePlan_Server,
		Delete: resourceDeletePlan_Server,

		Schema: map[string]*schema.Schema{
			"settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enableadvancedview": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Setting to suggest plan has some advanced settings present. Setting is OEM specific and not applicable for all cases.",
						},
						"filesearch": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "This feature applies to file servers and virtualization. Enabling this feature allows you to search for backed-up files using the global search bar, and creates resource pools with required infrastructure entities.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Flag for enabling indexing",
									},
									"statusmessage": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Tells what is happening behind the scene, so that user can knows why indexing is not enabled or if its in progress",
									},
									"status": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Type of indexing status. [NOT_APPLICABLE, ENABLED, SETUP_IN_PROGRESS]",
									},
								},
							},
						},
					},
				},
			},
			"backupcontent": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "This feature applies only to file system agents",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"windowsincludedpaths": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "Paths to include for Windows",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"unixincludedpaths": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "Paths to include for UNIX",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"macexcludedpaths": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "Paths to exclude for Mac",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"macfiltertoexcludepaths": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "Paths that are exception to excluded paths for Mac",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"macincludedpaths": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "Paths to include for Mac",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"unixexcludedpaths": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "Paths to exclude for UNIX",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"unixnumberofdatareaders": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"count": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Number of data readers.",
									},
									"useoptimal": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Set optimal number of data readers. if it is set to true, count will be ignored.",
									},
								},
							},
						},
						"backupsystemstate": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Do you want to back up the system state? Applicable only for Windows",
						},
						"backupsystemstateonlywithfullbackup": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Do you want to back up system state only with full backup? Applicable only if the value of backupSystemState is true",
						},
						"windowsexcludedpaths": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "Paths to exclude for Windows",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"usevssforsystemstate": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Do you want to back up system state with VSS? Applicable only if the value of backupSystemState is true",
						},
						"windowsnumberofdatareaders": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"count": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Number of data readers.",
									},
									"useoptimal": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Set optimal number of data readers. if it is set to true, count will be ignored.",
									},
								},
							},
						},
						"macnumberofdatareaders": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"count": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Number of data readers.",
									},
									"useoptimal": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Set optimal number of data readers. if it is set to true, count will be ignored.",
									},
								},
							},
						},
						"windowsfiltertoexcludepaths": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "Paths that are exception to excluded paths for Windows",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"unixfiltertoexcludepaths": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "Paths that are exception to excluded paths for Unix",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"forceupdateproperties": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Do you want to sync properties on associated subclients even if properties are overriden at subclient level?",
						},
					},
				},
			},
			"filesystemaddon": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "flag to enable backup content association for applicable file system workload.",
			},
			"allowplanoverride": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Flag to enable overriding of plan. Plan cannot be overriden by default.",
			},
			"planname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the new plan",
			},
			"workload": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"workloadtypes": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "",
									},
								},
							},
						},
						"workloadgrouptypes": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"solutions": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"backupdestinationids": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "Primary Backup Destination Ids (which were created before plan creation). This is only considered when backupDestinations array object is not defined.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"backupdestinations": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Backup destinations for the plan. Specify where you want to store your backup data.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ismirrorcopy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Is this a mirror copy? Only considered when isSnapCopy is true.",
						},
						"retentionperioddays": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Retention period in days. -1 can be specified for infinite retention. If this and snapRecoveryPoints both are not specified, this takes  precedence.",
						},
						"backupstocopy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "All_JOBS means SYNCHRONOUS copy type, others are applicable for SELECTIVE copy Type only. [All_JOBS, ALL_FULLS, HOURLY_FULLS, DAILY_FULLS, WEEKLY_FULLS, MONTHLY_FULLS, QUARTERLY_FULLS, HALF_YEARLY_FULLS, YEARLY_FULLS, ADVANCED]",
						},
						"backupdestinationname": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Backup destination details. Enter the name during creation.",
						},
						"extendedretentionrules": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "If you want to update, specify the whole object. Extended retention rules should be bigger than retention period.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"thirdextendedretentionrule": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"isinfiniteretention": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "If this is set as true, no need to specify retentionPeriodDays.",
												},
												"retentionperioddays": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "If this is set, no need to specify isInfiniteRetention as false.",
												},
												"type": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "All_JOBS means SYNCHRONOUS copy type, others are applicable for SELECTIVE copy Type only. [All_JOBS, ALL_FULLS, HOURLY_FULLS, DAILY_FULLS, WEEKLY_FULLS, MONTHLY_FULLS, QUARTERLY_FULLS, HALF_YEARLY_FULLS, YEARLY_FULLS, ADVANCED]",
												},
											},
										},
									},
									"firstextendedretentionrule": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"isinfiniteretention": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "If this is set as true, no need to specify retentionPeriodDays.",
												},
												"retentionperioddays": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "If this is set, no need to specify isInfiniteRetention as false.",
												},
												"type": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "All_JOBS means SYNCHRONOUS copy type, others are applicable for SELECTIVE copy Type only. [All_JOBS, ALL_FULLS, HOURLY_FULLS, DAILY_FULLS, WEEKLY_FULLS, MONTHLY_FULLS, QUARTERLY_FULLS, HALF_YEARLY_FULLS, YEARLY_FULLS, ADVANCED]",
												},
											},
										},
									},
									"secondextendedretentionrule": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"isinfiniteretention": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "If this is set as true, no need to specify retentionPeriodDays.",
												},
												"retentionperioddays": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "If this is set, no need to specify isInfiniteRetention as false.",
												},
												"type": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "All_JOBS means SYNCHRONOUS copy type, others are applicable for SELECTIVE copy Type only. [All_JOBS, ALL_FULLS, HOURLY_FULLS, DAILY_FULLS, WEEKLY_FULLS, MONTHLY_FULLS, QUARTERLY_FULLS, HALF_YEARLY_FULLS, YEARLY_FULLS, ADVANCED]",
												},
											},
										},
									},
								},
							},
						},
						"retentionruletype": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Which type of retention rule should be used for the given backup destination [RETENTION_PERIOD, SNAP_RECOVERY_POINTS]",
						},
						"snaprecoverypoints": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of snap recovery points for snap copy for retention. Can be specified instead of retention period in Days for snap copy.",
						},
						"sourcecopy": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"fullbackuptypestocopy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Which type of backup type should be copied for the given backup destination when backup type is not all jobs. Default is LAST while adding new backup destination. [FIRST, LAST]",
						},
						"useextendedretentionrules": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Use extended retention rules",
						},
						"backupstarttime": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Backup start time in seconds. The time is provided in unix time format.",
						},
						"overrideretentionsettings": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Tells if this copy should use storage pool retention period days or the retention defined for this copy. Set as true to use retention defined on this copy.",
						},
						"optimizeforinstantclone": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Flag to specify if primary storage is copy data management enabled.",
						},
						"netappcloudtarget": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Only for snap copy. Enabling this changes SVM Mapping  to NetApp cloud targets only.",
						},
						"mappings": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vendor": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Snapshot vendors available for Snap Copy mappings [NETAPP, AMAZON, PURE]",
									},
									"targetvendor": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"source": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"sourcevendor": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
									"target": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"id": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "",
												},
											},
										},
									},
								},
							},
						},
						"issnapcopy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Is this a snap copy? If isMirrorCopy is not set, then default is Vault/Replica.",
						},
						"storagetype": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "[ALL, DISK, CLOUD, HYPERSCALE, TAPE]",
						},
						"region": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
						"storagepool": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "",
									},
									"id": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "",
									},
								},
							},
						},
					},
				},
			},
			"overriderestrictions": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "To allow the derived plans that use this plan as the base plan to override the settings, property allowPlanOverride must be true, and then select one of the options for Storage pool, RPO and backup Content.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rpo": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "[OPTIONAL, MUST, NOT_ALLOWED]",
						},
						"backupcontent": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "[OPTIONAL, MUST, NOT_ALLOWED]",
						},
						"storagepool": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "[OPTIONAL, MUST, NOT_ALLOWED]",
						},
					},
				},
			},
			"snapshotoptions": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "This feature applies only to File System Agents",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enablebackupcopy": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Flag to enable backup copy",
						},
						"backupcopyrpomins": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Backup copy RPO in minutes",
						},
					},
				},
			},
			"parentplan": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "",
						},
						"id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "",
						},
					},
				},
			},
			"regiontoconfigure": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "",
						},
						"id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "",
						},
					},
				},
			},
			"databaseoptions": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "This feature applies only to database agents",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logbackuprpomins": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Log backup RPO in minutes",
						},
						"runfullbackupevery": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Full backup frequency in days",
						},
						"commitfrequencyinhours": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Commit frequency in hours",
						},
						"usediskcacheforlogbackups": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Use disk cache for log backups",
						},
					},
				},
			},
			"overrideinheritsettings": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "This feature applies to derived plans when inherit mode is optional.Provides user to set entity preference between parent and derived plan.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rpo": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Flag to specify if parent or derived plan rpo should be used when inherit mode is optional. True - derived, False - Base.",
						},
						"backupcontent": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Flag to specify if parent or derived plan backupContent should be used when inherit mode is optional. True - derived, False - Base.",
						},
						"backupdestination": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Flag to specify if parent or derived plan backupDestination should be used when inherit mode is optional. True - derived, False - Base.",
						},
					},
				},
			},
			"rpo": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Recovery Point Objective (RPO) is the maximum amount of time that data can be lost during a service disruption. Your RPO determines the frequency of your backup jobs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fullbackupwindow": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "Backup window for full backup",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dayofweek": {
										Type:        schema.TypeSet,
										Optional:    true,
										Computed:    true,
										Description: "",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"starttime": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Time in seconds since the beginning of the day",
									},
									"endtime": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Time in seconds since the beginning of the day",
									},
								},
							},
						},
						"sla": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "A server meets SLA (Service Level Agreement) when all of its subclients have at least one successful backup during the number of days specified at the CommCell, Server Group or plan level.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"exclusionreason": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Reason for exclusion from SLA",
									},
									"usesystemdefaultsla": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Flag to set to use System Default Service Level Agreement",
									},
									"enableafterdelay": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Time provided in Unix format. Give 0 to reset any existing delay.",
									},
									"excludefromsla": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Flag to set to exclude plan from SLA",
									},
									"slaperiod": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "SLA Period in Days",
									},
								},
							},
						},
						"backupfrequency": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"schedules": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"schedulename": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Name of the schedule, for modify",
												},
												"scheduleoption": {
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Description: "Specific options to be set on schedules",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"daysbetweenautoconvert": {
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
																Description: "Number of days between auto conversion of backup level applicable for databases on incremental and differential schedules of server plan",
															},
															"commitfrequencyinhours": {
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
																Description: "Commit frequency in hours for disk cache backups from automatic schedules",
															},
															"jobrunningtimeinmins": {
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
																Description: "total job running time in minutes",
															},
															"o365itemselectionoption": {
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "item backup option for O365 V2 backup jobs [SELECT_ALL, SELECT_NEVER_PROCESSED, SELECT_MEETING_SLA, SELECT_NOT_MEETING_SLA_PROCESSED_ATLEAST_ONCE, SELECT_FAILED_LAST_ATTEMPT, SELECT_PROCESSED_ATLEAST_ONCE, SELECT_NOT_MEETING_SLA, SELECT_MEETING_SLA_NOT_RECENTLY_BACKED_UP]",
															},
															"usediskcacheforlogbackups": {
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "Used to enable disk caching feature on databases for automatic schedules on server plan",
															},
														},
													},
												},
												"vmoperationtype": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Type of DR operation (only applicable for Failover groups) [PLANNED_FAILOVER, TEST_BOOT]",
												},
												"fordatabasesonly": {
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Boolean to indicate if schedule is for database agents",
												},
												"schedulepattern": {
													Type:        schema.TypeList,
													Required:    true,
													Description: "Used to describe when the schedule runs",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"enddate": {
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
																Description: "Schedule end date in epoch format",
															},
															"maxbackupintervalinmins": {
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
																Description: "The number of mins to force a backup on automatic schedule.",
															},
															"timezone": {
																Type:        schema.TypeList,
																Optional:    true,
																Computed:    true,
																Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"name": {
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																			Description: "",
																		},
																		"id": {
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Computed:    true,
																			Description: "",
																		},
																	},
																},
															},
															"weekofmonth": {
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "Specific week of a month [FIRST, SECOND, THIRD, FOURTH, LAST]",
															},
															"daysbetweensyntheticfulls": {
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
																Description: "No of days between two synthetic full jobs",
															},
															"exceptions": {
																Type:        schema.TypeSet,
																Optional:    true,
																Computed:    true,
																Description: "Exceptions to when a schedule should not run, either in dates or week of month and days",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"onweekofthemonth": {
																			Type:        schema.TypeSet,
																			Optional:    true,
																			Computed:    true,
																			Description: "On which week of month, for ex: FIRST, LAST",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"ondates": {
																			Type:        schema.TypeSet,
																			Optional:    true,
																			Computed:    true,
																			Description: "list of dates in a month. For ex: 1, 20",
																			Elem: &schema.Schema{
																				Type: schema.TypeInt,
																			},
																		},
																		"ondayoftheweek": {
																			Type:        schema.TypeSet,
																			Optional:    true,
																			Computed:    true,
																			Description: "On which days, for ex: MONDAY, FRIDAY",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																	},
																},
															},
															"frequency": {
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
																Description: "Frequency of the schedule based on schedule frequency type eg. for Hours, value 2 is 2 hours, for Minutes, 30 is 30 minutes, for Daily, 2 is 2 days. for Monthly 2 is it repeats every 2 months",
															},
															"weeklydays": {
																Type:        schema.TypeSet,
																Optional:    true,
																Computed:    true,
																Description: "Days of the week for weekly frequency",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"repeatuntiltime": {
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
																Description: "Until what time to repeat the schedule in a day, requires repeatIntervalInMinutes",
															},
															"monthofyear": {
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "[JANUARY, FEBRUARY, MARCH, APRIL, MAY, JUNE, JULY, AUGUST, SEPTEMBER, OCTOBER, NOVEMBER, DECEMBER]",
															},
															"dayofweek": {
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "[SUNDAY, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, DAY, WEEKDAY, WEEKEND_DAYS]",
															},
															"dayofmonth": {
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
																Description: "Day on which to run the schedule, applicable for monthly, yearly",
															},
															"schedulefrequencytype": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "schedule frequency type [MINUTES, DAILY, WEEKLY, MONTHLY, YEARLY, AUTOMATIC]",
															},
															"starttime": {
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
																Description: "start time of schedule in seconds",
															},
															"nooftimes": {
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
																Description: "The number of times you want the schedule to run.",
															},
															"repeatintervalinminutes": {
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
																Description: "How often in minutes in a day the schedule runs, applicable for daily, weekly, monthly and yearly frequency types.",
															},
															"startdate": {
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
																Description: "start date of schedule in epoch format",
															},
														},
													},
												},
												"backuptype": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Schedule Backup level [FULL, INCREMENTAL, DIFFERENTIAL, SYNTHETICFULL, TRANSACTIONLOG]",
												},
											},
										},
									},
								},
							},
						},
						"backupwindow": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "Backup window for incremental backup",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dayofweek": {
										Type:        schema.TypeSet,
										Optional:    true,
										Computed:    true,
										Description: "",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"starttime": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Time in seconds since the beginning of the day",
									},
									"endtime": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Time in seconds since the beginning of the day",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceCreatePlan_Server(d *schema.ResourceData, m interface{}) error {
	//API: (POST) /V4/ServerPlan
	var response_id = strconv.Itoa(0)
	var t_settings *handler.MsgServerPlanSettings
	if val, ok := d.GetOk("settings"); ok {
		t_settings = build_plan_server_msgserverplansettings(d, val.([]interface{}))
	}
	var t_backupcontent *handler.MsgPlanContent
	if val, ok := d.GetOk("backupcontent"); ok {
		t_backupcontent = build_plan_server_msgplancontent(d, val.([]interface{}))
	}
	var t_filesystemaddon *bool
	if val, ok := d.GetOk("filesystemaddon"); ok {
		t_filesystemaddon = handler.ToBooleanValue(val, false)
	}
	var t_allowplanoverride *bool
	if val, ok := d.GetOk("allowplanoverride"); ok {
		t_allowplanoverride = handler.ToBooleanValue(val, false)
	}
	var t_planname *string
	if val, ok := d.GetOk("planname"); ok {
		t_planname = handler.ToStringValue(val, false)
	}
	var t_workload *handler.MsgPlanWorkloads
	if val, ok := d.GetOk("workload"); ok {
		t_workload = build_plan_server_msgplanworkloads(d, val.([]interface{}))
	}
	var t_backupdestinationids []int
	if val, ok := d.GetOk("backupdestinationids"); ok {
		t_backupdestinationids = handler.ToIntArray(val.(*schema.Set).List())
	}
	var t_backupdestinations []handler.MsgCreatePlanBackupDestinationSet
	if val, ok := d.GetOk("backupdestinations"); ok {
		t_backupdestinations = build_plan_server_msgcreateplanbackupdestinationset_array(d, val.(*schema.Set).List())
	}
	var t_overriderestrictions *handler.MsgPlanOverrideSettings
	if val, ok := d.GetOk("overriderestrictions"); ok {
		t_overriderestrictions = build_plan_server_msgplanoverridesettings(d, val.([]interface{}))
	}
	var t_snapshotoptions *handler.MsgCreatePlanSnapshotOptions
	if val, ok := d.GetOk("snapshotoptions"); ok {
		t_snapshotoptions = build_plan_server_msgcreateplansnapshotoptions(d, val.([]interface{}))
	}
	var t_parentplan *handler.MsgIdName
	if val, ok := d.GetOk("parentplan"); ok {
		t_parentplan = build_plan_server_msgidname(d, val.([]interface{}))
	}
	var req = handler.MsgCreateServerPlanRequest{Settings: t_settings, BackupContent: t_backupcontent, FilesystemAddon: t_filesystemaddon, AllowPlanOverride: t_allowplanoverride, PlanName: t_planname, Workload: t_workload, BackupDestinationIds: t_backupdestinationids, BackupDestinations: t_backupdestinations, OverrideRestrictions: t_overriderestrictions, SnapshotOptions: t_snapshotoptions, ParentPlan: t_parentplan}
	resp, err := handler.CvCreateServerPlan(req)
	if err != nil {
		return fmt.Errorf("operation [CreateServerPlan] failed, Error %s", err)
	}
	if resp.Plan != nil {
		if resp.Plan.Id != nil {
			response_id = strconv.Itoa(*resp.Plan.Id)
		}
	}
	if response_id == "0" {
		return fmt.Errorf("operation [CreateServerPlan] failed")
	} else {
		d.SetId(response_id)
		return resourceCreateUpdatePlan_Server(d, m)
	}
}

func resourceReadPlan_Server(d *schema.ResourceData, m interface{}) error {
	//API: (GET) /V4/ServerPlan/{planId}
	resp, err := handler.CvGetPlanById(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			handler.LogEntry("debug", "entity not present, removig from state")
			d.SetId("")
			return nil
		}
		return fmt.Errorf("operation [GetPlanById] failed, Error %s", err)
	}
	if rtn, ok := serialize_plan_server_msgserverplansettings(d, resp.Settings); ok {
		d.Set("settings", rtn)
	} else {
		d.Set("settings", make([]map[string]interface{}, 0))
	}
	if rtn, ok := serialize_plan_server_msgplancontent(d, resp.BackupContent); ok {
		d.Set("backupcontent", rtn)
	} else {
		d.Set("backupcontent", make([]map[string]interface{}, 0))
	}
	if rtn, ok := statecopy_plan_server_databaseoptions(d); ok {
		d.Set("databaseoptions", rtn)
	} else {
		d.Set("databaseoptions", make([]map[string]interface{}, 0))
	}
	if resp.AllowPlanOverride != nil {
		d.Set("allowplanoverride", strconv.FormatBool(*resp.AllowPlanOverride))
	}
	if rtn, ok := serialize_plan_server_msgplanworkloads(d, resp.Workload); ok {
		d.Set("workload", rtn)
	} else {
		d.Set("workload", make([]map[string]interface{}, 0))
	}
	if resp.BackupDestinationIds != nil {
		d.Set("backupdestinationids", resp.BackupDestinationIds)
	}
	if rtn, ok := serialize_plan_server_msgserverplanrpo(d, resp.Rpo); ok {
		d.Set("rpo", rtn)
	} else {
		d.Set("rpo", make([]map[string]interface{}, 0))
	}
	if rtn, ok := statecopy_plan_server_backupdestinations(d); ok {
		d.Set("backupdestinations", rtn)
	} else {
		d.Set("backupdestinations", make([]map[string]interface{}, 0))
	}
	if rtn, ok := serialize_plan_server_msgplanoverridesettings(d, resp.OverrideRestrictions); ok {
		d.Set("overriderestrictions", rtn)
	} else {
		d.Set("overriderestrictions", make([]map[string]interface{}, 0))
	}
	if rtn, ok := serialize_plan_server_msgplansnapshotoptions(d, resp.SnapshotOptions); ok {
		d.Set("snapshotoptions", rtn)
	} else {
		d.Set("snapshotoptions", make([]map[string]interface{}, 0))
	}
	if resp.Plan.Name != nil {
		d.Set("planname", resp.Plan.Name)
	}
	return nil
}

func resourceUpdatePlan_Server(d *schema.ResourceData, m interface{}) error {
	//API: (PUT) /V4/ServerPlan/{planId}
	var t_regiontoconfigure *handler.MsgIdName
	if d.HasChange("regiontoconfigure") {
		val := d.Get("regiontoconfigure")
		t_regiontoconfigure = build_plan_server_msgidname(d, val.([]interface{}))
	}
	var t_settings *handler.MsgServerPlanSettings
	if d.HasChange("settings") {
		val := d.Get("settings")
		t_settings = build_plan_server_msgserverplansettings(d, val.([]interface{}))
	}
	var t_backupcontent *handler.MsgPlanContent
	if d.HasChange("backupcontent") {
		val := d.Get("backupcontent")
		t_backupcontent = build_plan_server_msgplancontent(d, val.([]interface{}))
	}
	var t_databaseoptions *handler.MsgServerPlanDatabaseOptionsInfo
	if d.HasChange("databaseoptions") {
		val := d.Get("databaseoptions")
		t_databaseoptions = build_plan_server_msgserverplandatabaseoptionsinfo(d, val.([]interface{}))
	}
	var t_overrideinheritsettings *handler.MsgPlanOverrideInheritSetting
	if d.HasChange("overrideinheritsettings") {
		val := d.Get("overrideinheritsettings")
		t_overrideinheritsettings = build_plan_server_msgplanoverrideinheritsetting(d, val.([]interface{}))
	}
	var t_filesystemaddon *bool
	if d.HasChange("filesystemaddon") {
		val := d.Get("filesystemaddon")
		t_filesystemaddon = handler.ToBooleanValue(val, false)
	}
	var t_allowplanoverride *bool
	if d.HasChange("allowplanoverride") {
		val := d.Get("allowplanoverride")
		t_allowplanoverride = handler.ToBooleanValue(val, false)
	}
	var t_workload *handler.MsgPlanWorkloads
	if d.HasChange("workload") {
		val := d.Get("workload")
		t_workload = build_plan_server_msgplanworkloads(d, val.([]interface{}))
	}
	var t_backupdestinationids []int
	if d.HasChange("backupdestinationids") {
		val := d.Get("backupdestinationids")
		t_backupdestinationids = handler.ToIntArray(val.(*schema.Set).List())
	}
	var t_rpo *handler.MsgServerPlanUpdateRPO
	if d.HasChange("rpo") {
		val := d.Get("rpo")
		t_rpo = build_plan_server_msgserverplanupdaterpo(d, val.([]interface{}))
	}
	var t_newname *string
	if d.HasChange("planname") {
		val := d.Get("planname")
		t_newname = handler.ToStringValue(val, false)
	}
	var t_overriderestrictions *handler.MsgPlanOverrideSettings
	if d.HasChange("overriderestrictions") {
		val := d.Get("overriderestrictions")
		t_overriderestrictions = build_plan_server_msgplanoverridesettings(d, val.([]interface{}))
	}
	var t_snapshotoptions *handler.MsgPlanSnapshotOptions
	if d.HasChange("snapshotoptions") {
		val := d.Get("snapshotoptions")
		t_snapshotoptions = build_plan_server_msgplansnapshotoptions(d, val.([]interface{}))
	}
	var req = handler.MsgModifyPlanRequest{RegionToConfigure: t_regiontoconfigure, Settings: t_settings, BackupContent: t_backupcontent, DatabaseOptions: t_databaseoptions, OverrideInheritSettings: t_overrideinheritsettings, FilesystemAddon: t_filesystemaddon, AllowPlanOverride: t_allowplanoverride, Workload: t_workload, BackupDestinationIds: t_backupdestinationids, Rpo: t_rpo, NewName: t_newname, OverrideRestrictions: t_overriderestrictions, SnapshotOptions: t_snapshotoptions}
	_, err := handler.CvModifyPlan(req, d.Id())
	if err != nil {
		return fmt.Errorf("operation [ModifyPlan] failed, Error %s", err)
	}
	return resourceReadPlan_Server(d, m)
}

func resourceCreateUpdatePlan_Server(d *schema.ResourceData, m interface{}) error {
	//API: (PUT) /V4/ServerPlan/{planId}
	var execUpdate bool = false
	var t_regiontoconfigure *handler.MsgIdName
	if val, ok := d.GetOk("regiontoconfigure"); ok {
		t_regiontoconfigure = build_plan_server_msgidname(d, val.([]interface{}))
		execUpdate = true
	}
	var t_databaseoptions *handler.MsgServerPlanDatabaseOptionsInfo
	if val, ok := d.GetOk("databaseoptions"); ok {
		t_databaseoptions = build_plan_server_msgserverplandatabaseoptionsinfo(d, val.([]interface{}))
		execUpdate = true
	}
	var t_overrideinheritsettings *handler.MsgPlanOverrideInheritSetting
	if val, ok := d.GetOk("overrideinheritsettings"); ok {
		t_overrideinheritsettings = build_plan_server_msgplanoverrideinheritsetting(d, val.([]interface{}))
		execUpdate = true
	}
	var t_rpo *handler.MsgServerPlanUpdateRPO
	if val, ok := d.GetOk("rpo"); ok {
		t_rpo = build_plan_server_msgserverplanupdaterpo(d, val.([]interface{}))
		execUpdate = true
	}
	if execUpdate {
		var req = handler.MsgModifyPlanRequest{RegionToConfigure: t_regiontoconfigure, DatabaseOptions: t_databaseoptions, OverrideInheritSettings: t_overrideinheritsettings, Rpo: t_rpo}
		_, err := handler.CvModifyPlan(req, d.Id())
		if err != nil {
			return fmt.Errorf("operation [ModifyPlan] failed, Error %s", err)
		}
	}
	return resourceReadPlan_Server(d, m)
}

func resourceDeletePlan_Server(d *schema.ResourceData, m interface{}) error {
	//API: (DELETE) /V4/ServerPlan/{planId}
	_, err := handler.CvDeletePlan(d.Id())
	if err != nil {
		return fmt.Errorf("operation [DeletePlan] failed, Error %s", err)
	}
	return nil
}

func build_plan_server_msgserverplanupdaterpo(d *schema.ResourceData, r []interface{}) *handler.MsgServerPlanUpdateRPO {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_fullbackupwindow []handler.MsgDayAndTimeSet
		if val, ok := tmp["fullbackupwindow"]; ok {
			t_fullbackupwindow = build_plan_server_msgdayandtimeset_array(d, val.(*schema.Set).List())
		}
		var t_sla *handler.MsgSLAUpdateOptions
		if val, ok := tmp["sla"]; ok {
			t_sla = build_plan_server_msgslaupdateoptions(d, val.([]interface{}))
		}
		var t_backupfrequency *handler.MsgPlanSchedules
		if val, ok := tmp["backupfrequency"]; ok {
			t_backupfrequency = build_plan_server_msgplanschedules(d, val.([]interface{}))
		}
		var t_backupwindow []handler.MsgDayAndTimeSet
		if val, ok := tmp["backupwindow"]; ok {
			t_backupwindow = build_plan_server_msgdayandtimeset_array(d, val.(*schema.Set).List())
		}
		return &handler.MsgServerPlanUpdateRPO{FullBackupWindow: t_fullbackupwindow, SLA: t_sla, BackupFrequency: t_backupfrequency, BackupWindow: t_backupwindow}
	} else {
		return nil
	}
}

func build_plan_server_msgdayandtimeset_array(d *schema.ResourceData, r []interface{}) []handler.MsgDayAndTimeSet {
	if r != nil {
		tmp := make([]handler.MsgDayAndTimeSet, len(r))
		for a, iter_a := range r {
			raw_a := iter_a.(map[string]interface{})
			var t_dayofweek []string
			if val, ok := raw_a["dayofweek"]; ok {
				t_dayofweek = handler.ToStringArray(val.(*schema.Set).List())
			}
			var t_starttime *int64
			if val, ok := raw_a["starttime"]; ok {
				t_starttime = handler.ToLongValue(val, false)
			}
			var t_endtime *int64
			if val, ok := raw_a["endtime"]; ok {
				t_endtime = handler.ToLongValue(val, true)
			}
			tmp[a] = handler.MsgDayAndTimeSet{DayOfWeek: t_dayofweek, StartTime: t_starttime, EndTime: t_endtime}
		}
		return tmp
	} else {
		return nil
	}
}

func build_plan_server_msgplanschedules(d *schema.ResourceData, r []interface{}) *handler.MsgPlanSchedules {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_schedules []handler.MsgPlanScheduleSet
		if val, ok := tmp["schedules"]; ok {
			t_schedules = build_plan_server_msgplanscheduleset_array(d, val.([]interface{}))
		}
		var t_operationtype *string
		if len(t_schedules) > 0 {
			var c_operationtype string = "OVERWRITE"
			t_operationtype = &c_operationtype
		}
		return &handler.MsgPlanSchedules{Schedules: t_schedules, OperationType: t_operationtype}
	} else {
		return nil
	}
}

func build_plan_server_msgplanscheduleset_array(d *schema.ResourceData, r []interface{}) []handler.MsgPlanScheduleSet {
	if r != nil {
		tmp := make([]handler.MsgPlanScheduleSet, len(r))
		for a, iter_a := range r {
			raw_a := iter_a.(map[string]interface{})
			var t_schedulename *string
			if val, ok := raw_a["schedulename"]; ok {
				t_schedulename = handler.ToStringValue(val, true)
			}
			var t_scheduleoption *handler.MsgScheduleOption
			if val, ok := raw_a["scheduleoption"]; ok {
				t_scheduleoption = build_plan_server_msgscheduleoption(d, val.([]interface{}))
			}
			var t_vmoperationtype *string
			if val, ok := raw_a["vmoperationtype"]; ok {
				t_vmoperationtype = handler.ToStringValue(val, true)
			}
			var t_fordatabasesonly *bool
			if val, ok := raw_a["fordatabasesonly"]; ok {
				t_fordatabasesonly = handler.ToBooleanValue(val, true)
			}
			var t_schedulepattern *handler.MsgSchedulePattern
			if val, ok := raw_a["schedulepattern"]; ok {
				t_schedulepattern = build_plan_server_msgschedulepattern(d, val.([]interface{}))
			}
			var t_backuptype *string
			if val, ok := raw_a["backuptype"]; ok {
				t_backuptype = handler.ToStringValue(val, true)
			}
			tmp[a] = handler.MsgPlanScheduleSet{ScheduleName: t_schedulename, ScheduleOption: t_scheduleoption, VmOperationType: t_vmoperationtype, ForDatabasesOnly: t_fordatabasesonly, SchedulePattern: t_schedulepattern, BackupType: t_backuptype}
		}
		return tmp
	} else {
		return nil
	}
}

func build_plan_server_msgschedulepattern(d *schema.ResourceData, r []interface{}) *handler.MsgSchedulePattern {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_enddate *int
		if val, ok := tmp["enddate"]; ok {
			t_enddate = handler.ToIntValue(val, true)
		}
		var t_maxbackupintervalinmins *int
		if val, ok := tmp["maxbackupintervalinmins"]; ok {
			t_maxbackupintervalinmins = handler.ToIntValue(val, true)
		}
		var t_timezone *handler.MsgIdName
		if val, ok := tmp["timezone"]; ok {
			t_timezone = build_plan_server_msgidname(d, val.([]interface{}))
		}
		var t_weekofmonth *string
		if val, ok := tmp["weekofmonth"]; ok {
			t_weekofmonth = handler.ToStringValue(val, true)
		}
		var t_daysbetweensyntheticfulls *int
		if val, ok := tmp["daysbetweensyntheticfulls"]; ok {
			t_daysbetweensyntheticfulls = handler.ToIntValue(val, true)
		}
		var t_exceptions []handler.MsgScheduleRunExceptionSet
		if val, ok := tmp["exceptions"]; ok {
			t_exceptions = build_plan_server_msgschedulerunexceptionset_array(d, val.(*schema.Set).List())
		}
		var t_frequency *int
		if val, ok := tmp["frequency"]; ok {
			t_frequency = handler.ToIntValue(val, true)
		}
		var t_weeklydays []string
		if val, ok := tmp["weeklydays"]; ok {
			t_weeklydays = handler.ToStringArray(val.(*schema.Set).List())
		}
		var t_repeatuntiltime *int
		if val, ok := tmp["repeatuntiltime"]; ok {
			t_repeatuntiltime = handler.ToIntValue(val, true)
		}
		var t_monthofyear *string
		if val, ok := tmp["monthofyear"]; ok {
			t_monthofyear = handler.ToStringValue(val, true)
		}
		var t_dayofweek *string
		if val, ok := tmp["dayofweek"]; ok {
			t_dayofweek = handler.ToStringValue(val, true)
		}
		var t_dayofmonth *int
		if val, ok := tmp["dayofmonth"]; ok {
			t_dayofmonth = handler.ToIntValue(val, true)
		}
		var t_schedulefrequencytype *string
		if val, ok := tmp["schedulefrequencytype"]; ok {
			t_schedulefrequencytype = handler.ToStringValue(val, true)
		}
		var t_starttime *int
		if val, ok := tmp["starttime"]; ok {
			t_starttime = handler.ToIntValue(val, true)
		}
		var t_nooftimes *int
		if val, ok := tmp["nooftimes"]; ok {
			t_nooftimes = handler.ToIntValue(val, true)
		}
		var t_repeatintervalinminutes *int
		if val, ok := tmp["repeatintervalinminutes"]; ok {
			t_repeatintervalinminutes = handler.ToIntValue(val, true)
		}
		var t_startdate *int
		if val, ok := tmp["startdate"]; ok {
			t_startdate = handler.ToIntValue(val, true)
		}
		return &handler.MsgSchedulePattern{EndDate: t_enddate, MaxBackupIntervalInMins: t_maxbackupintervalinmins, Timezone: t_timezone, WeekOfMonth: t_weekofmonth, DaysBetweenSyntheticFulls: t_daysbetweensyntheticfulls, Exceptions: t_exceptions, Frequency: t_frequency, WeeklyDays: t_weeklydays, RepeatUntilTime: t_repeatuntiltime, MonthOfYear: t_monthofyear, DayOfWeek: t_dayofweek, DayOfMonth: t_dayofmonth, ScheduleFrequencyType: t_schedulefrequencytype, StartTime: t_starttime, NoOfTimes: t_nooftimes, RepeatIntervalInMinutes: t_repeatintervalinminutes, StartDate: t_startdate}
	} else {
		return nil
	}
}

func build_plan_server_msgschedulerunexceptionset_array(d *schema.ResourceData, r []interface{}) []handler.MsgScheduleRunExceptionSet {
	if r != nil {
		tmp := make([]handler.MsgScheduleRunExceptionSet, len(r))
		for a, iter_a := range r {
			raw_a := iter_a.(map[string]interface{})
			var t_onweekofthemonth []string
			if val, ok := raw_a["onweekofthemonth"]; ok {
				t_onweekofthemonth = handler.ToStringArray(val.(*schema.Set).List())
			}
			var t_ondates []int
			if val, ok := raw_a["ondates"]; ok {
				t_ondates = handler.ToIntArray(val.(*schema.Set).List())
			}
			var t_ondayoftheweek []string
			if val, ok := raw_a["ondayoftheweek"]; ok {
				t_ondayoftheweek = handler.ToStringArray(val.(*schema.Set).List())
			}
			tmp[a] = handler.MsgScheduleRunExceptionSet{OnWeekOfTheMonth: t_onweekofthemonth, OnDates: t_ondates, OnDayOfTheWeek: t_ondayoftheweek}
		}
		return tmp
	} else {
		return nil
	}
}

func build_plan_server_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_name *string
		if val, ok := tmp["name"]; ok {
			t_name = handler.ToStringValue(val, true)
		}
		var t_id *int
		if val, ok := tmp["id"]; ok {
			t_id = handler.ToIntValue(val, true)
		}
		return &handler.MsgIdName{Name: t_name, Id: t_id}
	} else {
		return nil
	}
}

func build_plan_server_msgscheduleoption(d *schema.ResourceData, r []interface{}) *handler.MsgScheduleOption {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_daysbetweenautoconvert *int
		if val, ok := tmp["daysbetweenautoconvert"]; ok {
			t_daysbetweenautoconvert = handler.ToIntValue(val, true)
		}
		var t_commitfrequencyinhours *int
		if val, ok := tmp["commitfrequencyinhours"]; ok {
			t_commitfrequencyinhours = handler.ToIntValue(val, true)
		}
		var t_jobrunningtimeinmins *int
		if val, ok := tmp["jobrunningtimeinmins"]; ok {
			t_jobrunningtimeinmins = handler.ToIntValue(val, true)
		}
		var t_o365itemselectionoption *string
		if val, ok := tmp["o365itemselectionoption"]; ok {
			t_o365itemselectionoption = handler.ToStringValue(val, true)
		}
		var t_usediskcacheforlogbackups *bool
		if val, ok := tmp["usediskcacheforlogbackups"]; ok {
			t_usediskcacheforlogbackups = handler.ToBooleanValue(val, true)
		}
		return &handler.MsgScheduleOption{DaysBetweenAutoConvert: t_daysbetweenautoconvert, CommitFrequencyInHours: t_commitfrequencyinhours, JobRunningTimeInMins: t_jobrunningtimeinmins, O365ItemSelectionOption: t_o365itemselectionoption, UseDiskCacheForLogBackups: t_usediskcacheforlogbackups}
	} else {
		return nil
	}
}

func build_plan_server_msgslaupdateoptions(d *schema.ResourceData, r []interface{}) *handler.MsgSLAUpdateOptions {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_exclusionreason *string
		if val, ok := tmp["exclusionreason"]; ok {
			t_exclusionreason = handler.ToStringValue(val, true)
		}
		var t_usesystemdefaultsla *bool
		if val, ok := tmp["usesystemdefaultsla"]; ok {
			t_usesystemdefaultsla = handler.ToBooleanValue(val, true)
		}
		var t_enableafterdelay *int
		if val, ok := tmp["enableafterdelay"]; ok {
			t_enableafterdelay = handler.ToIntValue(val, true)
		}
		var t_excludefromsla *bool
		if val, ok := tmp["excludefromsla"]; ok {
			t_excludefromsla = handler.ToBooleanValue(val, true)
		}
		var t_slaperiod *int
		if val, ok := tmp["slaperiod"]; ok {
			t_slaperiod = handler.ToIntValue(val, true)
		}
		return &handler.MsgSLAUpdateOptions{ExclusionReason: t_exclusionreason, UseSystemDefaultSLA: t_usesystemdefaultsla, EnableAfterDelay: t_enableafterdelay, ExcludeFromSLA: t_excludefromsla, SLAPeriod: t_slaperiod}
	} else {
		return nil
	}
}

func build_plan_server_msgplanoverrideinheritsetting(d *schema.ResourceData, r []interface{}) *handler.MsgPlanOverrideInheritSetting {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_rpo *bool
		if val, ok := tmp["rpo"]; ok {
			t_rpo = handler.ToBooleanValue(val, true)
		}
		var t_backupcontent *bool
		if val, ok := tmp["backupcontent"]; ok {
			t_backupcontent = handler.ToBooleanValue(val, true)
		}
		var t_backupdestination *bool
		if val, ok := tmp["backupdestination"]; ok {
			t_backupdestination = handler.ToBooleanValue(val, true)
		}
		return &handler.MsgPlanOverrideInheritSetting{Rpo: t_rpo, BackupContent: t_backupcontent, BackupDestination: t_backupdestination}
	} else {
		return nil
	}
}

func build_plan_server_msgserverplandatabaseoptionsinfo(d *schema.ResourceData, r []interface{}) *handler.MsgServerPlanDatabaseOptionsInfo {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_logbackuprpomins *int
		if val, ok := tmp["logbackuprpomins"]; ok {
			t_logbackuprpomins = handler.ToIntValue(val, true)
		}
		var t_runfullbackupevery *int
		if val, ok := tmp["runfullbackupevery"]; ok {
			t_runfullbackupevery = handler.ToIntValue(val, true)
		}
		var t_commitfrequencyinhours *int
		if val, ok := tmp["commitfrequencyinhours"]; ok {
			t_commitfrequencyinhours = handler.ToIntValue(val, true)
		}
		var t_usediskcacheforlogbackups *bool
		if val, ok := tmp["usediskcacheforlogbackups"]; ok {
			t_usediskcacheforlogbackups = handler.ToBooleanValue(val, true)
		}
		return &handler.MsgServerPlanDatabaseOptionsInfo{LogBackupRPOMins: t_logbackuprpomins, RunFullBackupEvery: t_runfullbackupevery, CommitFrequencyInHours: t_commitfrequencyinhours, UseDiskCacheForLogBackups: t_usediskcacheforlogbackups}
	} else {
		return nil
	}
}

func build_plan_server_msgplansnapshotoptions(d *schema.ResourceData, r []interface{}) *handler.MsgPlanSnapshotOptions {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_backupcopyfrequency *handler.MsgBackupFrequencyPattern
		if val, ok := tmp["backupcopyfrequency"]; ok {
			t_backupcopyfrequency = build_plan_server_msgbackupfrequencypattern(d, val.([]interface{}))
		}
		var t_enablebackupcopy *bool
		if val, ok := tmp["enablebackupcopy"]; ok {
			t_enablebackupcopy = handler.ToBooleanValue(val, true)
		}
		var t_backupcopyrpomins *int
		if val, ok := tmp["backupcopyrpomins"]; ok {
			t_backupcopyrpomins = handler.ToIntValue(val, true)
		}
		return &handler.MsgPlanSnapshotOptions{BackupCopyFrequency: t_backupcopyfrequency, EnableBackupCopy: t_enablebackupcopy, BackupCopyRPOMins: t_backupcopyrpomins}
	} else {
		return nil
	}
}

func build_plan_server_msgbackupfrequencypattern(d *schema.ResourceData, r []interface{}) *handler.MsgBackupFrequencyPattern {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_weeklydays []string
		if val, ok := tmp["weeklydays"]; ok {
			t_weeklydays = handler.ToStringArray(val.(*schema.Set).List())
		}
		var t_monthofyear *string
		if val, ok := tmp["monthofyear"]; ok {
			t_monthofyear = handler.ToStringValue(val, true)
		}
		var t_dayofweek *string
		if val, ok := tmp["dayofweek"]; ok {
			t_dayofweek = handler.ToStringValue(val, true)
		}
		var t_dayofmonth *int
		if val, ok := tmp["dayofmonth"]; ok {
			t_dayofmonth = handler.ToIntValue(val, true)
		}
		var t_schedulefrequencytype *string
		if val, ok := tmp["schedulefrequencytype"]; ok {
			t_schedulefrequencytype = handler.ToStringValue(val, true)
		}
		var t_weekofmonth *string
		if val, ok := tmp["weekofmonth"]; ok {
			t_weekofmonth = handler.ToStringValue(val, true)
		}
		var t_starttime *int
		if val, ok := tmp["starttime"]; ok {
			t_starttime = handler.ToIntValue(val, true)
		}
		var t_frequency *int
		if val, ok := tmp["frequency"]; ok {
			t_frequency = handler.ToIntValue(val, true)
		}
		return &handler.MsgBackupFrequencyPattern{WeeklyDays: t_weeklydays, MonthOfYear: t_monthofyear, DayOfWeek: t_dayofweek, DayOfMonth: t_dayofmonth, ScheduleFrequencyType: t_schedulefrequencytype, WeekOfMonth: t_weekofmonth, StartTime: t_starttime, Frequency: t_frequency}
	} else {
		return nil
	}
}

func build_plan_server_msgplanoverridesettings(d *schema.ResourceData, r []interface{}) *handler.MsgPlanOverrideSettings {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_rpo *string
		if val, ok := tmp["rpo"]; ok {
			t_rpo = handler.ToStringValue(val, true)
		}
		var t_backupcontent *string
		if val, ok := tmp["backupcontent"]; ok {
			t_backupcontent = handler.ToStringValue(val, true)
		}
		var t_storagepool *string
		if val, ok := tmp["storagepool"]; ok {
			t_storagepool = handler.ToStringValue(val, true)
		}
		return &handler.MsgPlanOverrideSettings{RPO: t_rpo, BackupContent: t_backupcontent, StoragePool: t_storagepool}
	} else {
		return nil
	}
}

func build_plan_server_msgplanworkloads(d *schema.ResourceData, r []interface{}) *handler.MsgPlanWorkloads {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_workloadtypes []handler.MsgIdNameSet
		if val, ok := tmp["workloadtypes"]; ok {
			t_workloadtypes = build_plan_server_msgidnameset_array(d, val.(*schema.Set).List())
		}
		var t_workloadgrouptypes []string
		if val, ok := tmp["workloadgrouptypes"]; ok {
			t_workloadgrouptypes = handler.ToStringArray(val.(*schema.Set).List())
		}
		var t_solutions []handler.MsgIdNameSet
		if val, ok := tmp["solutions"]; ok {
			t_solutions = build_plan_server_msgidnameset_array(d, val.(*schema.Set).List())
		}
		return &handler.MsgPlanWorkloads{WorkloadTypes: t_workloadtypes, WorkloadGroupTypes: t_workloadgrouptypes, Solutions: t_solutions}
	} else {
		return nil
	}
}

func build_plan_server_msgidnameset_array(d *schema.ResourceData, r []interface{}) []handler.MsgIdNameSet {
	if r != nil {
		tmp := make([]handler.MsgIdNameSet, len(r))
		for a, iter_a := range r {
			raw_a := iter_a.(map[string]interface{})
			var t_id *int
			if val, ok := raw_a["id"]; ok {
				t_id = handler.ToIntValue(val, true)
			}
			tmp[a] = handler.MsgIdNameSet{Id: t_id}
		}
		return tmp
	} else {
		return nil
	}
}

func build_plan_server_msgplancontent(d *schema.ResourceData, r []interface{}) *handler.MsgPlanContent {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_windowsincludedpaths []string
		if val, ok := tmp["windowsincludedpaths"]; ok {
			t_windowsincludedpaths = handler.ToStringArray(val.(*schema.Set).List())
		}
		var t_unixincludedpaths []string
		if val, ok := tmp["unixincludedpaths"]; ok {
			t_unixincludedpaths = handler.ToStringArray(val.(*schema.Set).List())
		}
		var t_macexcludedpaths []string
		if val, ok := tmp["macexcludedpaths"]; ok {
			t_macexcludedpaths = handler.ToStringArray(val.(*schema.Set).List())
		}
		var t_macfiltertoexcludepaths []string
		if val, ok := tmp["macfiltertoexcludepaths"]; ok {
			t_macfiltertoexcludepaths = handler.ToStringArray(val.(*schema.Set).List())
		}
		var t_macincludedpaths []string
		if val, ok := tmp["macincludedpaths"]; ok {
			t_macincludedpaths = handler.ToStringArray(val.(*schema.Set).List())
		}
		var t_unixexcludedpaths []string
		if val, ok := tmp["unixexcludedpaths"]; ok {
			t_unixexcludedpaths = handler.ToStringArray(val.(*schema.Set).List())
		}
		var t_unixnumberofdatareaders *handler.MsgPlanContentDataReaders
		if val, ok := tmp["unixnumberofdatareaders"]; ok {
			t_unixnumberofdatareaders = build_plan_server_msgplancontentdatareaders(d, val.([]interface{}))
		}
		var t_backupsystemstate *bool
		if val, ok := tmp["backupsystemstate"]; ok {
			t_backupsystemstate = handler.ToBooleanValue(val, true)
		}
		var t_backupsystemstateonlywithfullbackup *bool
		if val, ok := tmp["backupsystemstateonlywithfullbackup"]; ok {
			t_backupsystemstateonlywithfullbackup = handler.ToBooleanValue(val, true)
		}
		var t_windowsexcludedpaths []string
		if val, ok := tmp["windowsexcludedpaths"]; ok {
			t_windowsexcludedpaths = handler.ToStringArray(val.(*schema.Set).List())
		}
		var t_usevssforsystemstate *bool
		if val, ok := tmp["usevssforsystemstate"]; ok {
			t_usevssforsystemstate = handler.ToBooleanValue(val, true)
		}
		var t_windowsnumberofdatareaders *handler.MsgPlanContentDataReaders
		if val, ok := tmp["windowsnumberofdatareaders"]; ok {
			t_windowsnumberofdatareaders = build_plan_server_msgplancontentdatareaders(d, val.([]interface{}))
		}
		var t_macnumberofdatareaders *handler.MsgPlanContentDataReaders
		if val, ok := tmp["macnumberofdatareaders"]; ok {
			t_macnumberofdatareaders = build_plan_server_msgplancontentdatareaders(d, val.([]interface{}))
		}
		var t_windowsfiltertoexcludepaths []string
		if val, ok := tmp["windowsfiltertoexcludepaths"]; ok {
			t_windowsfiltertoexcludepaths = handler.ToStringArray(val.(*schema.Set).List())
		}
		var t_unixfiltertoexcludepaths []string
		if val, ok := tmp["unixfiltertoexcludepaths"]; ok {
			t_unixfiltertoexcludepaths = handler.ToStringArray(val.(*schema.Set).List())
		}
		var t_forceupdateproperties *bool
		if val, ok := tmp["forceupdateproperties"]; ok {
			t_forceupdateproperties = handler.ToBooleanValue(val, true)
		}
		return &handler.MsgPlanContent{WindowsIncludedPaths: t_windowsincludedpaths, UnixIncludedPaths: t_unixincludedpaths, MacExcludedPaths: t_macexcludedpaths, MacFilterToExcludePaths: t_macfiltertoexcludepaths, MacIncludedPaths: t_macincludedpaths, UnixExcludedPaths: t_unixexcludedpaths, UnixNumberOfDataReaders: t_unixnumberofdatareaders, BackupSystemState: t_backupsystemstate, BackupSystemStateOnlyWithFullBackup: t_backupsystemstateonlywithfullbackup, WindowsExcludedPaths: t_windowsexcludedpaths, UseVSSForSystemState: t_usevssforsystemstate, WindowsNumberOfDataReaders: t_windowsnumberofdatareaders, MacNumberOfDataReaders: t_macnumberofdatareaders, WindowsFilterToExcludePaths: t_windowsfiltertoexcludepaths, UnixFilterToExcludePaths: t_unixfiltertoexcludepaths, ForceUpdateProperties: t_forceupdateproperties}
	} else {
		return nil
	}
}

func build_plan_server_msgplancontentdatareaders(d *schema.ResourceData, r []interface{}) *handler.MsgPlanContentDataReaders {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_count *int
		if val, ok := tmp["count"]; ok {
			t_count = handler.ToIntValue(val, true)
		}
		var t_useoptimal *bool
		if val, ok := tmp["useoptimal"]; ok {
			t_useoptimal = handler.ToBooleanValue(val, true)
		}
		return &handler.MsgPlanContentDataReaders{Count: t_count, UseOptimal: t_useoptimal}
	} else {
		return nil
	}
}

func build_plan_server_msgserverplansettings(d *schema.ResourceData, r []interface{}) *handler.MsgServerPlanSettings {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_enableadvancedview *bool
		if val, ok := tmp["enableadvancedview"]; ok {
			t_enableadvancedview = handler.ToBooleanValue(val, true)
		}
		var t_filesearch *handler.MsgPlanFileSearch
		if val, ok := tmp["filesearch"]; ok {
			t_filesearch = build_plan_server_msgplanfilesearch(d, val.([]interface{}))
		}
		return &handler.MsgServerPlanSettings{EnableAdvancedView: t_enableadvancedview, FileSearch: t_filesearch}
	} else {
		return nil
	}
}

func build_plan_server_msgplanfilesearch(d *schema.ResourceData, r []interface{}) *handler.MsgPlanFileSearch {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_enabled *bool
		if val, ok := tmp["enabled"]; ok {
			t_enabled = handler.ToBooleanValue(val, true)
		}
		var t_statusmessage *string
		if val, ok := tmp["statusmessage"]; ok {
			t_statusmessage = handler.ToStringValue(val, true)
		}
		var t_status *string
		if val, ok := tmp["status"]; ok {
			t_status = handler.ToStringValue(val, true)
		}
		return &handler.MsgPlanFileSearch{Enabled: t_enabled, StatusMessage: t_statusmessage, Status: t_status}
	} else {
		return nil
	}
}

func build_plan_server_msgcreateplansnapshotoptions(d *schema.ResourceData, r []interface{}) *handler.MsgCreatePlanSnapshotOptions {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_enablebackupcopy *bool
		if val, ok := tmp["enablebackupcopy"]; ok {
			t_enablebackupcopy = handler.ToBooleanValue(val, true)
		}
		var t_backupcopyrpomins *int
		if val, ok := tmp["backupcopyrpomins"]; ok {
			t_backupcopyrpomins = handler.ToIntValue(val, true)
		}
		return &handler.MsgCreatePlanSnapshotOptions{EnableBackupCopy: t_enablebackupcopy, BackupCopyRPOMins: t_backupcopyrpomins}
	} else {
		return nil
	}
}

func build_plan_server_msgcreateplanbackupdestinationset_array(d *schema.ResourceData, r []interface{}) []handler.MsgCreatePlanBackupDestinationSet {
	if r != nil {
		tmp := make([]handler.MsgCreatePlanBackupDestinationSet, len(r))
		for a, iter_a := range r {
			raw_a := iter_a.(map[string]interface{})
			var t_ismirrorcopy *bool
			if val, ok := raw_a["ismirrorcopy"]; ok {
				t_ismirrorcopy = handler.ToBooleanValue(val, true)
			}
			var t_retentionperioddays *int
			if val, ok := raw_a["retentionperioddays"]; ok {
				t_retentionperioddays = handler.ToIntValue(val, true)
			}
			var t_backupstocopy *string
			if val, ok := raw_a["backupstocopy"]; ok {
				t_backupstocopy = handler.ToStringValue(val, true)
			}
			var t_backupdestinationname *string
			if val, ok := raw_a["backupdestinationname"]; ok {
				t_backupdestinationname = handler.ToStringValue(val, true)
			}
			var t_extendedretentionrules *handler.MsgExtendedRetentionRules
			if val, ok := raw_a["extendedretentionrules"]; ok {
				t_extendedretentionrules = build_plan_server_msgextendedretentionrules(d, val.([]interface{}))
			}
			var t_retentionruletype *string
			if val, ok := raw_a["retentionruletype"]; ok {
				t_retentionruletype = handler.ToStringValue(val, true)
			}
			var t_snaprecoverypoints *int
			if val, ok := raw_a["snaprecoverypoints"]; ok {
				t_snaprecoverypoints = handler.ToIntValue(val, true)
			}
			var t_sourcecopy *handler.MsgIdName
			if val, ok := raw_a["sourcecopy"]; ok {
				t_sourcecopy = build_plan_server_msgidname(d, val.([]interface{}))
			}
			var t_fullbackuptypestocopy *string
			if val, ok := raw_a["fullbackuptypestocopy"]; ok {
				t_fullbackuptypestocopy = handler.ToStringValue(val, true)
			}
			var t_useextendedretentionrules *bool
			if val, ok := raw_a["useextendedretentionrules"]; ok {
				t_useextendedretentionrules = handler.ToBooleanValue(val, true)
			}
			var t_backupstarttime *int
			if val, ok := raw_a["backupstarttime"]; ok {
				t_backupstarttime = handler.ToIntValue(val, true)
			}
			var t_overrideretentionsettings *bool
			if val, ok := raw_a["overrideretentionsettings"]; ok {
				t_overrideretentionsettings = handler.ToBooleanValue(val, true)
			}
			var t_optimizeforinstantclone *bool
			if val, ok := raw_a["optimizeforinstantclone"]; ok {
				t_optimizeforinstantclone = handler.ToBooleanValue(val, true)
			}
			var t_netappcloudtarget *bool
			if val, ok := raw_a["netappcloudtarget"]; ok {
				t_netappcloudtarget = handler.ToBooleanValue(val, true)
			}
			var t_mappings []handler.MsgSnapshotCopyMappingSet
			if val, ok := raw_a["mappings"]; ok {
				t_mappings = build_plan_server_msgsnapshotcopymappingset_array(d, val.(*schema.Set).List())
			}
			var t_issnapcopy *bool
			if val, ok := raw_a["issnapcopy"]; ok {
				t_issnapcopy = handler.ToBooleanValue(val, true)
			}
			var t_storagetype *string
			if val, ok := raw_a["storagetype"]; ok {
				t_storagetype = handler.ToStringValue(val, true)
			}
			var t_region *handler.MsgIdName
			if val, ok := raw_a["region"]; ok {
				t_region = build_plan_server_msgidname(d, val.([]interface{}))
			}
			var t_storagepool *handler.MsgIdName
			if val, ok := raw_a["storagepool"]; ok {
				t_storagepool = build_plan_server_msgidname(d, val.([]interface{}))
			}
			tmp[a] = handler.MsgCreatePlanBackupDestinationSet{IsMirrorCopy: t_ismirrorcopy, RetentionPeriodDays: t_retentionperioddays, BackupsToCopy: t_backupstocopy, BackupDestinationName: t_backupdestinationname, ExtendedRetentionRules: t_extendedretentionrules, RetentionRuleType: t_retentionruletype, SnapRecoveryPoints: t_snaprecoverypoints, SourceCopy: t_sourcecopy, FullBackupTypesToCopy: t_fullbackuptypestocopy, UseExtendedRetentionRules: t_useextendedretentionrules, BackupStartTime: t_backupstarttime, OverrideRetentionSettings: t_overrideretentionsettings, OptimizeForInstantClone: t_optimizeforinstantclone, NetAppCloudTarget: t_netappcloudtarget, Mappings: t_mappings, IsSnapCopy: t_issnapcopy, StorageType: t_storagetype, Region: t_region, StoragePool: t_storagepool}
		}
		return tmp
	} else {
		return nil
	}
}

func build_plan_server_msgsnapshotcopymappingset_array(d *schema.ResourceData, r []interface{}) []handler.MsgSnapshotCopyMappingSet {
	if r != nil {
		tmp := make([]handler.MsgSnapshotCopyMappingSet, len(r))
		for a, iter_a := range r {
			raw_a := iter_a.(map[string]interface{})
			var t_vendor *string
			if val, ok := raw_a["vendor"]; ok {
				t_vendor = handler.ToStringValue(val, true)
			}
			var t_targetvendor *handler.MsgIdName
			if val, ok := raw_a["targetvendor"]; ok {
				t_targetvendor = build_plan_server_msgidname(d, val.([]interface{}))
			}
			var t_source *handler.MsgIdName
			if val, ok := raw_a["source"]; ok {
				t_source = build_plan_server_msgidname(d, val.([]interface{}))
			}
			var t_sourcevendor *handler.MsgIdName
			if val, ok := raw_a["sourcevendor"]; ok {
				t_sourcevendor = build_plan_server_msgidname(d, val.([]interface{}))
			}
			var t_target *handler.MsgIdName
			if val, ok := raw_a["target"]; ok {
				t_target = build_plan_server_msgidname(d, val.([]interface{}))
			}
			tmp[a] = handler.MsgSnapshotCopyMappingSet{Vendor: t_vendor, TargetVendor: t_targetvendor, Source: t_source, SourceVendor: t_sourcevendor, Target: t_target}
		}
		return tmp
	} else {
		return nil
	}
}

func build_plan_server_msgextendedretentionrules(d *schema.ResourceData, r []interface{}) *handler.MsgExtendedRetentionRules {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_thirdextendedretentionrule *handler.MsgPlanRetentionRule
		if val, ok := tmp["thirdextendedretentionrule"]; ok {
			t_thirdextendedretentionrule = build_plan_server_msgplanretentionrule(d, val.([]interface{}))
		}
		var t_firstextendedretentionrule *handler.MsgPlanRetentionRule
		if val, ok := tmp["firstextendedretentionrule"]; ok {
			t_firstextendedretentionrule = build_plan_server_msgplanretentionrule(d, val.([]interface{}))
		}
		var t_secondextendedretentionrule *handler.MsgPlanRetentionRule
		if val, ok := tmp["secondextendedretentionrule"]; ok {
			t_secondextendedretentionrule = build_plan_server_msgplanretentionrule(d, val.([]interface{}))
		}
		return &handler.MsgExtendedRetentionRules{ThirdExtendedRetentionRule: t_thirdextendedretentionrule, FirstExtendedRetentionRule: t_firstextendedretentionrule, SecondExtendedRetentionRule: t_secondextendedretentionrule}
	} else {
		return nil
	}
}

func build_plan_server_msgplanretentionrule(d *schema.ResourceData, r []interface{}) *handler.MsgPlanRetentionRule {
	if len(r) > 0 && r[0] != nil {
		tmp := r[0].(map[string]interface{})
		var t_isinfiniteretention *bool
		if val, ok := tmp["isinfiniteretention"]; ok {
			t_isinfiniteretention = handler.ToBooleanValue(val, true)
		}
		var t_retentionperioddays *int
		if val, ok := tmp["retentionperioddays"]; ok {
			t_retentionperioddays = handler.ToIntValue(val, true)
		}
		var t_type *string
		if val, ok := tmp["type"]; ok {
			t_type = handler.ToStringValue(val, true)
		}
		return &handler.MsgPlanRetentionRule{IsInfiniteRetention: t_isinfiniteretention, RetentionPeriodDays: t_retentionperioddays, Type: t_type}
	} else {
		return nil
	}
}

func serialize_plan_server_msgplansnapshotoptions(d *schema.ResourceData, data *handler.MsgPlanSnapshotOptions) ([]map[string]interface{}, bool) {
	//MsgCreatePlanSnapshotOptions
	//MsgPlanSnapshotOptions
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 1)
	val[0] = make(map[string]interface{})
	added := false
	if data.EnableBackupCopy != nil {
		val[0]["enablebackupcopy"] = strconv.FormatBool(*data.EnableBackupCopy)
		added = true
	}
	if data.BackupCopyRPOMins != nil {
		val[0]["backupcopyrpomins"] = data.BackupCopyRPOMins
		added = true
	}
	if added {
		return val, true
	} else {
		return nil, false
	}
}

func serialize_plan_server_msgplanoverridesettings(d *schema.ResourceData, data *handler.MsgPlanOverrideSettings) ([]map[string]interface{}, bool) {
	//MsgPlanOverrideSettings
	//MsgPlanOverrideSettings
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 1)
	val[0] = make(map[string]interface{})
	added := false
	if data.RPO != nil {
		val[0]["rpo"] = data.RPO
		added = true
	}
	if data.BackupContent != nil {
		val[0]["backupcontent"] = data.BackupContent
		added = true
	}
	if data.StoragePool != nil {
		val[0]["storagepool"] = data.StoragePool
		added = true
	}
	if added {
		return val, true
	} else {
		return nil, false
	}
}

func statecopy_plan_server_backupdestinations(d *schema.ResourceData) ([]interface{}, bool) {
	//STATE COPY
	var_a := d.Get("backupdestinations")
	if var_a != nil {
		return var_a.(*schema.Set).List(), true
	}
	return nil, false
}

func serialize_plan_server_msgserverplanrpo(d *schema.ResourceData, data *handler.MsgServerPlanRPO) ([]map[string]interface{}, bool) {
	//MsgServerPlanUpdateRPO
	//MsgServerPlanRPO
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 1)
	val[0] = make(map[string]interface{})
	added := false
	if rtn, ok := serialize_plan_server_msgdayandtimeset_array(d, data.FullBackupWindow); ok {
		val[0]["fullbackupwindow"] = rtn
		added = true
	}
	if rtn, ok := serialize_plan_server_msgslaoptions(d, data.SLA); ok {
		val[0]["sla"] = rtn
		added = true
	}
	if rtn, ok := serialize_plan_server_msgplanschedules(d, data.BackupFrequency); ok {
		val[0]["backupfrequency"] = rtn
		added = true
	}
	if rtn, ok := serialize_plan_server_msgdayandtimeset_array(d, data.BackupWindow); ok {
		val[0]["backupwindow"] = rtn
		added = true
	}
	if added {
		return val, true
	} else {
		return nil, false
	}
}

func serialize_plan_server_msgdayandtimeset_array(d *schema.ResourceData, data []handler.MsgDayAndTimeSet) ([]map[string]interface{}, bool) {
	//MsgServerPlanUpdateRPO -> MsgDayAndTimeSet
	//MsgServerPlanRPO -> MsgDayAndTimeSet
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 0)
	for i := range data {
		tmp := make(map[string]interface{})
		added := false
		if data[i].DayOfWeek != nil {
			tmp["dayofweek"] = data[i].DayOfWeek
			added = true
		}
		if data[i].StartTime != nil {
			tmp["starttime"] = data[i].StartTime
			added = true
		}
		if data[i].EndTime != nil {
			tmp["endtime"] = data[i].EndTime
			added = true
		}
		if added {
			val = append(val, tmp)
		}
	}
	return val, true
}

func serialize_plan_server_msgplanschedules(d *schema.ResourceData, data *handler.MsgPlanSchedules) ([]map[string]interface{}, bool) {
	//MsgServerPlanUpdateRPO -> MsgPlanSchedules
	//MsgServerPlanRPO -> MsgPlanSchedules
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 1)
	val[0] = make(map[string]interface{})
	added := false
	if rtn, ok := serialize_plan_server_msgplanscheduleset_array(d, data.Schedules); ok {
		val[0]["schedules"] = rtn
		added = true
	}
	if added {
		return val, true
	} else {
		return nil, false
	}
}

func serialize_plan_server_msgplanscheduleset_array(d *schema.ResourceData, data []handler.MsgPlanScheduleSet) ([]map[string]interface{}, bool) {
	//MsgServerPlanUpdateRPO -> MsgPlanSchedules -> MsgPlanScheduleSet
	//MsgServerPlanRPO -> MsgPlanSchedules -> MsgPlanScheduleSet
	if data == nil {
		return nil, false
	}
	data = handler.SortPlanSchedules(d, data)
	val := make([]map[string]interface{}, 0)
	for i := range data {
		tmp := make(map[string]interface{})
		added := false
		if data[i].ScheduleName != nil {
			tmp["schedulename"] = data[i].ScheduleName
			added = true
		}
		if rtn, ok := serialize_plan_server_msgscheduleoption(d, data[i].ScheduleOption); ok {
			tmp["scheduleoption"] = rtn
			added = true
		}
		if data[i].VmOperationType != nil {
			tmp["vmoperationtype"] = data[i].VmOperationType
			added = true
		}
		if data[i].ForDatabasesOnly != nil {
			tmp["fordatabasesonly"] = strconv.FormatBool(*data[i].ForDatabasesOnly)
			added = true
		}
		if rtn, ok := serialize_plan_server_msgschedulepattern(d, data[i].SchedulePattern); ok {
			tmp["schedulepattern"] = rtn
			added = true
		}
		if data[i].BackupType != nil {
			tmp["backuptype"] = data[i].BackupType
			added = true
		}
		if added {
			val = append(val, tmp)
		}
	}
	return val, true
}

func serialize_plan_server_msgschedulepattern(d *schema.ResourceData, data *handler.MsgSchedulePattern) ([]map[string]interface{}, bool) {
	//MsgServerPlanUpdateRPO -> MsgPlanSchedules -> MsgPlanScheduleSet -> MsgSchedulePattern
	//MsgServerPlanRPO -> MsgPlanSchedules -> MsgPlanScheduleSet -> MsgSchedulePattern
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 1)
	val[0] = make(map[string]interface{})
	added := false
	if data.EndDate != nil {
		val[0]["enddate"] = data.EndDate
		added = true
	}
	if data.MaxBackupIntervalInMins != nil {
		val[0]["maxbackupintervalinmins"] = data.MaxBackupIntervalInMins
		added = true
	}
	if rtn, ok := serialize_plan_server_msgidname(d, data.Timezone); ok {
		val[0]["timezone"] = rtn
		added = true
	}
	if data.WeekOfMonth != nil {
		val[0]["weekofmonth"] = data.WeekOfMonth
		added = true
	}
	if data.DaysBetweenSyntheticFulls != nil {
		val[0]["daysbetweensyntheticfulls"] = data.DaysBetweenSyntheticFulls
		added = true
	}
	if rtn, ok := serialize_plan_server_msgschedulerunexceptionset_array(d, data.Exceptions); ok {
		val[0]["exceptions"] = rtn
		added = true
	}
	if data.Frequency != nil {
		val[0]["frequency"] = data.Frequency
		added = true
	}
	if data.WeeklyDays != nil {
		val[0]["weeklydays"] = data.WeeklyDays
		added = true
	}
	if data.RepeatUntilTime != nil {
		val[0]["repeatuntiltime"] = data.RepeatUntilTime
		added = true
	}
	if data.MonthOfYear != nil {
		val[0]["monthofyear"] = data.MonthOfYear
		added = true
	}
	if data.DayOfWeek != nil {
		val[0]["dayofweek"] = data.DayOfWeek
		added = true
	}
	if data.DayOfMonth != nil {
		val[0]["dayofmonth"] = data.DayOfMonth
		added = true
	}
	if data.ScheduleFrequencyType != nil {
		val[0]["schedulefrequencytype"] = data.ScheduleFrequencyType
		added = true
	}
	if data.StartTime != nil {
		val[0]["starttime"] = data.StartTime
		added = true
	}
	if data.NoOfTimes != nil {
		val[0]["nooftimes"] = data.NoOfTimes
		added = true
	}
	if data.RepeatIntervalInMinutes != nil {
		val[0]["repeatintervalinminutes"] = data.RepeatIntervalInMinutes
		added = true
	}
	if data.StartDate != nil {
		val[0]["startdate"] = data.StartDate
		added = true
	}
	if added {
		return val, true
	} else {
		return nil, false
	}
}

func serialize_plan_server_msgschedulerunexceptionset_array(d *schema.ResourceData, data []handler.MsgScheduleRunExceptionSet) ([]map[string]interface{}, bool) {
	//MsgServerPlanUpdateRPO -> MsgPlanSchedules -> MsgPlanScheduleSet -> MsgSchedulePattern -> MsgScheduleRunExceptionSet
	//MsgServerPlanRPO -> MsgPlanSchedules -> MsgPlanScheduleSet -> MsgSchedulePattern -> MsgScheduleRunExceptionSet
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 0)
	for i := range data {
		tmp := make(map[string]interface{})
		added := false
		if data[i].OnWeekOfTheMonth != nil {
			tmp["onweekofthemonth"] = data[i].OnWeekOfTheMonth
			added = true
		}
		if data[i].OnDates != nil {
			tmp["ondates"] = data[i].OnDates
			added = true
		}
		if data[i].OnDayOfTheWeek != nil {
			tmp["ondayoftheweek"] = data[i].OnDayOfTheWeek
			added = true
		}
		if added {
			val = append(val, tmp)
		}
	}
	return val, true
}

func serialize_plan_server_msgidname(d *schema.ResourceData, data *handler.MsgIdName) ([]map[string]interface{}, bool) {
	//MsgServerPlanUpdateRPO -> MsgPlanSchedules -> MsgPlanScheduleSet -> MsgSchedulePattern -> MsgIdName
	//MsgServerPlanRPO -> MsgPlanSchedules -> MsgPlanScheduleSet -> MsgSchedulePattern -> MsgIdName
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 1)
	val[0] = make(map[string]interface{})
	added := false
	if data.Name != nil {
		val[0]["name"] = data.Name
		added = true
	}
	if data.Id != nil {
		val[0]["id"] = data.Id
		added = true
	}
	if added {
		return val, true
	} else {
		return nil, false
	}
}

func serialize_plan_server_msgscheduleoption(d *schema.ResourceData, data *handler.MsgScheduleOption) ([]map[string]interface{}, bool) {
	//MsgServerPlanUpdateRPO -> MsgPlanSchedules -> MsgPlanScheduleSet -> MsgScheduleOption
	//MsgServerPlanRPO -> MsgPlanSchedules -> MsgPlanScheduleSet -> MsgScheduleOption
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 1)
	val[0] = make(map[string]interface{})
	added := false
	if data.DaysBetweenAutoConvert != nil {
		val[0]["daysbetweenautoconvert"] = data.DaysBetweenAutoConvert
		added = true
	}
	if data.CommitFrequencyInHours != nil {
		val[0]["commitfrequencyinhours"] = data.CommitFrequencyInHours
		added = true
	}
	if data.JobRunningTimeInMins != nil {
		val[0]["jobrunningtimeinmins"] = data.JobRunningTimeInMins
		added = true
	}
	if data.O365ItemSelectionOption != nil {
		val[0]["o365itemselectionoption"] = data.O365ItemSelectionOption
		added = true
	}
	if data.UseDiskCacheForLogBackups != nil {
		val[0]["usediskcacheforlogbackups"] = strconv.FormatBool(*data.UseDiskCacheForLogBackups)
		added = true
	}
	if added {
		return val, true
	} else {
		return nil, false
	}
}

func serialize_plan_server_msgslaoptions(d *schema.ResourceData, data *handler.MsgSLAOptions) ([]map[string]interface{}, bool) {
	//MsgServerPlanUpdateRPO -> MsgSLAUpdateOptions
	//MsgServerPlanRPO -> MsgSLAOptions
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 1)
	val[0] = make(map[string]interface{})
	added := false
	if data.ExclusionReason != nil {
		val[0]["exclusionreason"] = data.ExclusionReason
		added = true
	}
	if data.UseSystemDefaultSLA != nil {
		val[0]["usesystemdefaultsla"] = strconv.FormatBool(*data.UseSystemDefaultSLA)
		added = true
	}
	if data.EnableAfterDelay != nil {
		val[0]["enableafterdelay"] = data.EnableAfterDelay
		added = true
	}
	if tmp, ok := statecopy_plan_server_rpo_sla_excludefromsla(d); ok {
		val[0]["excludefromsla"] = tmp
		added = true
	}
	if data.SLAPeriod != nil {
		val[0]["slaperiod"] = data.SLAPeriod
		added = true
	}
	if added {
		return val, true
	} else {
		return nil, false
	}
}

func statecopy_plan_server_rpo_sla_excludefromsla(d *schema.ResourceData) (*string, bool) {
	//STATE COPY
	var_a := d.Get("rpo").([]interface{})
	if len(var_a) > 0 {
		tmp_a := var_a[0].(map[string]interface{})
		if var_b, ok := tmp_a["sla"].([]interface{}); ok {
			if len(var_b) > 0 {
				tmp_b := var_b[0].(map[string]interface{})
				if var_c, ok := tmp_b["excludefromsla"].(string); ok {
					return &var_c, true
				}
			}
		}
	}
	return nil, false
}

func serialize_plan_server_msgplanworkloads(d *schema.ResourceData, data *handler.MsgPlanWorkloads) ([]map[string]interface{}, bool) {
	//MsgPlanWorkloads
	//MsgPlanWorkloads
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 1)
	val[0] = make(map[string]interface{})
	added := false
	if rtn, ok := serialize_plan_server_msgidnameset_array(d, data.WorkloadTypes); ok {
		val[0]["workloadtypes"] = rtn
		added = true
	}
	if data.WorkloadGroupTypes != nil {
		val[0]["workloadgrouptypes"] = data.WorkloadGroupTypes
		added = true
	}
	if rtn, ok := serialize_plan_server_msgidnameset_array(d, data.Solutions); ok {
		val[0]["solutions"] = rtn
		added = true
	}
	if added {
		return val, true
	} else {
		return nil, false
	}
}

func serialize_plan_server_msgidnameset_array(d *schema.ResourceData, data []handler.MsgIdNameSet) ([]map[string]interface{}, bool) {
	//MsgPlanWorkloads -> MsgIdNameSet
	//MsgPlanWorkloads -> MsgIdNameSet
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 0)
	for i := range data {
		tmp := make(map[string]interface{})
		added := false
		if data[i].Id != nil {
			tmp["id"] = data[i].Id
			added = true
		}
		if added {
			val = append(val, tmp)
		}
	}
	return val, true
}

func statecopy_plan_server_databaseoptions(d *schema.ResourceData) ([]interface{}, bool) {
	//STATE COPY
	var_a := d.Get("databaseoptions").([]interface{})
	if len(var_a) > 0 {
		return var_a, true
	}
	return nil, false
}

func serialize_plan_server_msgplancontent(d *schema.ResourceData, data *handler.MsgPlanContent) ([]map[string]interface{}, bool) {
	//MsgPlanContent
	//MsgPlanContent
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 1)
	val[0] = make(map[string]interface{})
	added := false
	if tmp, ok := statecopy_plan_server_backupcontent_windowsincludedpaths(d); ok {
		val[0]["windowsincludedpaths"] = tmp
		added = true
	}
	if tmp, ok := statecopy_plan_server_backupcontent_unixincludedpaths(d); ok {
		val[0]["unixincludedpaths"] = tmp
		added = true
	}
	if tmp, ok := statecopy_plan_server_backupcontent_macexcludedpaths(d); ok {
		val[0]["macexcludedpaths"] = tmp
		added = true
	}
	if tmp, ok := statecopy_plan_server_backupcontent_macfiltertoexcludepaths(d); ok {
		val[0]["macfiltertoexcludepaths"] = tmp
		added = true
	}
	if tmp, ok := statecopy_plan_server_backupcontent_macincludedpaths(d); ok {
		val[0]["macincludedpaths"] = tmp
		added = true
	}
	if tmp, ok := statecopy_plan_server_backupcontent_unixexcludedpaths(d); ok {
		val[0]["unixexcludedpaths"] = tmp
		added = true
	}
	if rtn, ok := serialize_plan_server_msgplancontentdatareaders(d, data.UnixNumberOfDataReaders); ok {
		val[0]["unixnumberofdatareaders"] = rtn
		added = true
	}
	if data.BackupSystemState != nil {
		val[0]["backupsystemstate"] = strconv.FormatBool(*data.BackupSystemState)
		added = true
	}
	if data.BackupSystemStateOnlyWithFullBackup != nil {
		val[0]["backupsystemstateonlywithfullbackup"] = strconv.FormatBool(*data.BackupSystemStateOnlyWithFullBackup)
		added = true
	}
	if tmp, ok := statecopy_plan_server_backupcontent_windowsexcludedpaths(d); ok {
		val[0]["windowsexcludedpaths"] = tmp
		added = true
	}
	if data.UseVSSForSystemState != nil {
		val[0]["usevssforsystemstate"] = strconv.FormatBool(*data.UseVSSForSystemState)
		added = true
	}
	if rtn, ok := serialize_plan_server_msgplancontentdatareaders(d, data.WindowsNumberOfDataReaders); ok {
		val[0]["windowsnumberofdatareaders"] = rtn
		added = true
	}
	if rtn, ok := serialize_plan_server_msgplancontentdatareaders(d, data.MacNumberOfDataReaders); ok {
		val[0]["macnumberofdatareaders"] = rtn
		added = true
	}
	if tmp, ok := statecopy_plan_server_backupcontent_windowsfiltertoexcludepaths(d); ok {
		val[0]["windowsfiltertoexcludepaths"] = tmp
		added = true
	}
	if tmp, ok := statecopy_plan_server_backupcontent_unixfiltertoexcludepaths(d); ok {
		val[0]["unixfiltertoexcludepaths"] = tmp
		added = true
	}
	if data.ForceUpdateProperties != nil {
		val[0]["forceupdateproperties"] = strconv.FormatBool(*data.ForceUpdateProperties)
		added = true
	}
	if added {
		return val, true
	} else {
		return nil, false
	}
}

func statecopy_plan_server_backupcontent_unixfiltertoexcludepaths(d *schema.ResourceData) ([]string, bool) {
	//STATE COPY
	var_a := d.Get("backupcontent").([]interface{})
	if len(var_a) > 0 {
		tmp_a := var_a[0].(map[string]interface{})
		if var_b, ok := tmp_a["unixfiltertoexcludepaths"]; ok {
			return handler.ToStringArray(var_b.(*schema.Set).List()), true
		}
	}
	return nil, false
}

func statecopy_plan_server_backupcontent_windowsfiltertoexcludepaths(d *schema.ResourceData) ([]string, bool) {
	//STATE COPY
	var_a := d.Get("backupcontent").([]interface{})
	if len(var_a) > 0 {
		tmp_a := var_a[0].(map[string]interface{})
		if var_b, ok := tmp_a["windowsfiltertoexcludepaths"]; ok {
			return handler.ToStringArray(var_b.(*schema.Set).List()), true
		}
	}
	return nil, false
}

func serialize_plan_server_msgplancontentdatareaders(d *schema.ResourceData, data *handler.MsgPlanContentDataReaders) ([]map[string]interface{}, bool) {
	//MsgPlanContent -> MsgPlanContentDataReaders
	//MsgPlanContent -> MsgPlanContentDataReaders
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 1)
	val[0] = make(map[string]interface{})
	added := false
	if data.Count != nil {
		val[0]["count"] = data.Count
		added = true
	}
	if data.UseOptimal != nil {
		val[0]["useoptimal"] = strconv.FormatBool(*data.UseOptimal)
		added = true
	}
	if added {
		return val, true
	} else {
		return nil, false
	}
}

func statecopy_plan_server_backupcontent_windowsexcludedpaths(d *schema.ResourceData) ([]string, bool) {
	//STATE COPY
	var_a := d.Get("backupcontent").([]interface{})
	if len(var_a) > 0 {
		tmp_a := var_a[0].(map[string]interface{})
		if var_b, ok := tmp_a["windowsexcludedpaths"]; ok {
			return handler.ToStringArray(var_b.(*schema.Set).List()), true
		}
	}
	return nil, false
}

func statecopy_plan_server_backupcontent_unixexcludedpaths(d *schema.ResourceData) ([]string, bool) {
	//STATE COPY
	var_a := d.Get("backupcontent").([]interface{})
	if len(var_a) > 0 {
		tmp_a := var_a[0].(map[string]interface{})
		if var_b, ok := tmp_a["unixexcludedpaths"]; ok {
			return handler.ToStringArray(var_b.(*schema.Set).List()), true
		}
	}
	return nil, false
}

func statecopy_plan_server_backupcontent_macincludedpaths(d *schema.ResourceData) ([]string, bool) {
	//STATE COPY
	var_a := d.Get("backupcontent").([]interface{})
	if len(var_a) > 0 {
		tmp_a := var_a[0].(map[string]interface{})
		if var_b, ok := tmp_a["macincludedpaths"]; ok {
			return handler.ToStringArray(var_b.(*schema.Set).List()), true
		}
	}
	return nil, false
}

func statecopy_plan_server_backupcontent_macfiltertoexcludepaths(d *schema.ResourceData) ([]string, bool) {
	//STATE COPY
	var_a := d.Get("backupcontent").([]interface{})
	if len(var_a) > 0 {
		tmp_a := var_a[0].(map[string]interface{})
		if var_b, ok := tmp_a["macfiltertoexcludepaths"]; ok {
			return handler.ToStringArray(var_b.(*schema.Set).List()), true
		}
	}
	return nil, false
}

func statecopy_plan_server_backupcontent_macexcludedpaths(d *schema.ResourceData) ([]string, bool) {
	//STATE COPY
	var_a := d.Get("backupcontent").([]interface{})
	if len(var_a) > 0 {
		tmp_a := var_a[0].(map[string]interface{})
		if var_b, ok := tmp_a["macexcludedpaths"]; ok {
			return handler.ToStringArray(var_b.(*schema.Set).List()), true
		}
	}
	return nil, false
}

func statecopy_plan_server_backupcontent_unixincludedpaths(d *schema.ResourceData) ([]string, bool) {
	//STATE COPY
	var_a := d.Get("backupcontent").([]interface{})
	if len(var_a) > 0 {
		tmp_a := var_a[0].(map[string]interface{})
		if var_b, ok := tmp_a["unixincludedpaths"]; ok {
			return handler.ToStringArray(var_b.(*schema.Set).List()), true
		}
	}
	return nil, false
}

func statecopy_plan_server_backupcontent_windowsincludedpaths(d *schema.ResourceData) ([]string, bool) {
	//STATE COPY
	var_a := d.Get("backupcontent").([]interface{})
	if len(var_a) > 0 {
		tmp_a := var_a[0].(map[string]interface{})
		if var_b, ok := tmp_a["windowsincludedpaths"]; ok {
			return handler.ToStringArray(var_b.(*schema.Set).List()), true
		}
	}
	return nil, false
}

func serialize_plan_server_msgserverplansettings(d *schema.ResourceData, data *handler.MsgServerPlanSettings) ([]map[string]interface{}, bool) {
	//MsgServerPlanSettings
	//MsgServerPlanSettings
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 1)
	val[0] = make(map[string]interface{})
	added := false
	if data.EnableAdvancedView != nil {
		val[0]["enableadvancedview"] = strconv.FormatBool(*data.EnableAdvancedView)
		added = true
	}
	if rtn, ok := serialize_plan_server_msgplanfilesearch(d, data.FileSearch); ok {
		val[0]["filesearch"] = rtn
		added = true
	}
	if added {
		return val, true
	} else {
		return nil, false
	}
}

func serialize_plan_server_msgplanfilesearch(d *schema.ResourceData, data *handler.MsgPlanFileSearch) ([]map[string]interface{}, bool) {
	//MsgServerPlanSettings -> MsgPlanFileSearch
	//MsgServerPlanSettings -> MsgPlanFileSearch
	if data == nil {
		return nil, false
	}
	val := make([]map[string]interface{}, 1)
	val[0] = make(map[string]interface{})
	added := false
	if data.Enabled != nil {
		val[0]["enabled"] = strconv.FormatBool(*data.Enabled)
		added = true
	}
	if data.StatusMessage != nil {
		val[0]["statusmessage"] = data.StatusMessage
		added = true
	}
	if data.Status != nil {
		val[0]["status"] = data.Status
		added = true
	}
	if added {
		return val, true
	} else {
		return nil, false
	}
}
