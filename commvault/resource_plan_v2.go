package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceplan_v2() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateplan_v2,
        Read:   resourceReadplan_v2,
        Update: resourceUpdateplan_v2,
        Delete: resourceDeleteplan_v2,

        Schema: map[string]*schema.Schema{
            "settings": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enableadvancedview": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Setting to suggest plan has some advanced settings present. Setting is OEM specific and not applicable for all cases.",
                        },
                        "filesearch": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "This feature applies to file servers and virtualization. Enabling this feature allows you to search for backed-up files using the global search bar, and creates resource pools with required infrastructure entities.",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "enabled": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Flag for enabling indexing",
                                    },
                                    "statusmessage": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Tells what is happening behind the scene, so that user can knows why indexing is not enabled or if its in progress",
                                    },
                                    "status": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Type of indexing status.",
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "backupcontent": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "This feature applies only to file system agents",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "windowsincludedpaths": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "Paths to include for Windows",
                            Elem: &schema.Schema{
                                Type:    schema.TypeString,
                            },
                        },
                        "backupsystemstate": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Do you want to back up the system state? Applicable only for Windows",
                        },
                        "backupsystemstateonlywithfullbackup": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Do you want to back up system state only with full backup? Applicable only if the value of backupSystemState is true",
                        },
                        "windowsexcludedpaths": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "Paths to exclude for Windows",
                            Elem: &schema.Schema{
                                Type:    schema.TypeString,
                            },
                        },
                        "unixincludedpaths": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "Paths to include for UNIX",
                            Elem: &schema.Schema{
                                Type:    schema.TypeString,
                            },
                        },
                        "usevssforsystemstate": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Do you want to back up system state with VSS? Applicable only if the value of backupSystemState is true",
                        },
                        "macexcludedpaths": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "Paths to exclude for Mac",
                            Elem: &schema.Schema{
                                Type:    schema.TypeString,
                            },
                        },
                        "macfiltertoexcludepaths": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "Paths that are exception to excluded paths for Mac",
                            Elem: &schema.Schema{
                                Type:    schema.TypeString,
                            },
                        },
                        "windowsfiltertoexcludepaths": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "Paths that are exception to excluded paths for Windows",
                            Elem: &schema.Schema{
                                Type:    schema.TypeString,
                            },
                        },
                        "unixfiltertoexcludepaths": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "Paths that are exception to excluded paths for Unix",
                            Elem: &schema.Schema{
                                Type:    schema.TypeString,
                            },
                        },
                        "macincludedpaths": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "Paths to include for Mac",
                            Elem: &schema.Schema{
                                Type:    schema.TypeString,
                            },
                        },
                        "unixexcludedpaths": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "Paths to exclude for UNIX",
                            Elem: &schema.Schema{
                                Type:    schema.TypeString,
                            },
                        },
                    },
                },
            },
            "databaseoptions": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "This feature applies only to database agents",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "logbackuprpomins": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Log backup RPO in minutes",
                        },
                        "commitfrequencyinhours": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Commit frequency in hours",
                        },
                        "usediskcacheforlogbackups": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Use disk cache for log backups",
                        },
                    },
                },
            },
            "filesystemaddon": &schema.Schema{
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: "flag to enable backup content association for applicable file system workload.",
            },
            "allowplanoverride": &schema.Schema{
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: "Flag to enable overriding of plan. Plan cannot be overriden by default.",
            },
            "planname": &schema.Schema{
                Type:        schema.TypeString,
                Required:    true,
                Description: "Name of the new plan",
            },
            "workload": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "workloadtypes": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "id": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "workloadgrouptypes": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Schema{
                                Type:    schema.TypeString,
                            },
                        },
                        "solutions": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "id": &schema.Schema{
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
            "rpo": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Recovery Point Objective (RPO) is the maximum amount of time that data can be lost during a service disruption. Your RPO determines the frequency of your backup jobs.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "fullbackupwindow": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "Backup window for full backup",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "dayofweek": &schema.Schema{
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Description: "",
                                        Elem: &schema.Schema{
                                            Type:    schema.TypeString,
                                        },
                                    },
                                    "starttime": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Time in seconds since the beginning of the day",
                                    },
                                    "endtime": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Time in seconds since the beginning of the day",
                                    },
                                },
                            },
                        },
                        "backupfrequency": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "schedules": &schema.Schema{
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "schedulename": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Name of the schedule, for modify",
                                                },
                                                "policyid": &schema.Schema{
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Schedule policy Id to which the schedule belongs",
                                                },
                                                "vmoperationtype": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Type of DR operation (only applicable for Failover groups)",
                                                },
                                                "fordatabasesonly": &schema.Schema{
                                                    Type:        schema.TypeBool,
                                                    Required:    true,
                                                    Description: "Boolean to indicate if schedule is for database agents",
                                                },
                                                "scheduleoperation": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Required:    true,
                                                    Description: "Operation being performed on schedule",
                                                },
                                                "schedulepattern": &schema.Schema{
                                                    Type:        schema.TypeList,
                                                    Required:    true,
                                                    Description: "Used to describe when the schedule runs",
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "enddate": &schema.Schema{
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "Schedule end date in epoch format",
                                                            },
                                                            "timezone": &schema.Schema{
                                                                Type:        schema.TypeList,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "",
                                                                Elem: &schema.Resource{
                                                                    Schema: map[string]*schema.Schema{
                                                                        "name": &schema.Schema{
                                                                            Type:        schema.TypeString,
                                                                            Optional:    true,
                                                                            Computed:    true,
                                                                            Description: "",
                                                                        },
                                                                        "id": &schema.Schema{
                                                                            Type:        schema.TypeInt,
                                                                            Optional:    true,
                                                                            Computed:    true,
                                                                            Description: "",
                                                                        },
                                                                    },
                                                                },
                                                            },
                                                            "weekofmonth": &schema.Schema{
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "Specific week of a month",
                                                            },
                                                            "daysbetweensyntheticfulls": &schema.Schema{
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "No of days between two synthetic full jobs",
                                                            },
                                                            "exceptions": &schema.Schema{
                                                                Type:        schema.TypeSet,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "Exceptions to when a schedule should not run, either in dates or week of month and days",
                                                                Elem: &schema.Resource{
                                                                    Schema: map[string]*schema.Schema{
                                                                        "onweekofthemonth": &schema.Schema{
                                                                            Type:        schema.TypeSet,
                                                                            Optional:    true,
                                                                            Description: "On which week of month, for ex: FIRST, LAST",
                                                                            Elem: &schema.Schema{
                                                                                Type:    schema.TypeString,
                                                                            },
                                                                        },
                                                                        "ondates": &schema.Schema{
                                                                            Type:        schema.TypeSet,
                                                                            Optional:    true,
                                                                            Description: "list of dates in a month. For ex: 1, 20",
                                                                            Elem: &schema.Schema{
                                                                                Type:    schema.TypeInt,
                                                                            },
                                                                        },
                                                                        "ondayoftheweek": &schema.Schema{
                                                                            Type:        schema.TypeSet,
                                                                            Optional:    true,
                                                                            Description: "On which days, for ex: MONDAY, FRIDAY",
                                                                            Elem: &schema.Schema{
                                                                                Type:    schema.TypeString,
                                                                            },
                                                                        },
                                                                    },
                                                                },
                                                            },
                                                            "frequency": &schema.Schema{
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "Frequency of the schedule based on schedule frequency type eg. for Hours, value 2 is 2 hours, for Minutes, 30 is 30 minutes, for Daily, 2 is 2 days. for Monthly 2 is it repeats every 2 months",
                                                            },
                                                            "weeklydays": &schema.Schema{
                                                                Type:        schema.TypeSet,
                                                                Optional:    true,
                                                                Description: "Days of the week for weekly frequency",
                                                                Elem: &schema.Schema{
                                                                    Type:    schema.TypeString,
                                                                },
                                                            },
                                                            "repeatuntiltime": &schema.Schema{
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "Until what time to repeat the schedule in a day, requires repeatIntervalInMinutes",
                                                            },
                                                            "monthofyear": &schema.Schema{
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "",
                                                            },
                                                            "dayofweek": &schema.Schema{
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "",
                                                            },
                                                            "dayofmonth": &schema.Schema{
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "Day on which to run the schedule, applicable for monthly, yearly",
                                                            },
                                                            "schedulefrequencytype": &schema.Schema{
                                                                Type:        schema.TypeString,
                                                                Required:    true,
                                                                Description: "schedule frequency type",
                                                            },
                                                            "starttime": &schema.Schema{
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "start time of schedule in seconds",
                                                            },
                                                            "nooftimes": &schema.Schema{
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "The number of times you want the schedule to run.",
                                                            },
                                                            "repeatintervalinminutes": &schema.Schema{
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "How often in minutes in a day the schedule runs, applicable for daily, weekly, monthly and yearly frequency types.",
                                                            },
                                                            "startdate": &schema.Schema{
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "start date of schedule in epoch format",
                                                            },
                                                        },
                                                    },
                                                },
                                                "scheduleid": &schema.Schema{
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Id of the schedule if available, required for modifying, deleting schedule",
                                                },
                                                "backuptype": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Required:    true,
                                                    Description: "Schedule Backup level",
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "backupwindow": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "Backup window for incremental backup",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "dayofweek": &schema.Schema{
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Description: "",
                                        Elem: &schema.Schema{
                                            Type:    schema.TypeString,
                                        },
                                    },
                                    "starttime": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Time in seconds since the beginning of the day",
                                    },
                                    "endtime": &schema.Schema{
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
            "backupdestinations": &schema.Schema{
                Type:        schema.TypeSet,
                Required:    true,
                Description: "Backup destinations for the plan. Specify where you want to store your backup data.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "ismirrorcopy": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Is this a mirror copy? Only considered when isSnapCopy is true.",
                        },
                        "retentionperioddays": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Retention period in days. -1 can be specified for infinite retention. If this and snapRecoveryPoints both are not specified, this takes  precedence.",
                        },
                        "backupstocopy": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "backupdestinationname": &schema.Schema{
                            Type:        schema.TypeString,
                            Required:    true,
                            Description: "Backup destination details. Enter the name during creation.",
                        },
                        "extendedretentionrules": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "If you want to update, specify the whole object. Extended retention rules should be bigger than retention period.",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "thirdextendedretentionrule": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "isinfiniteretention": &schema.Schema{
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "retentionperioddays": &schema.Schema{
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Default value is 30 days. Infinite retention takes precedence over retentionPeriodDays.",
                                                },
                                                "type": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                            },
                                        },
                                    },
                                    "firstextendedretentionrule": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "isinfiniteretention": &schema.Schema{
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "retentionperioddays": &schema.Schema{
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Default value is 30 days. Infinite retention takes precedence over retentionPeriodDays.",
                                                },
                                                "type": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                            },
                                        },
                                    },
                                    "secondextendedretentionrule": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "isinfiniteretention": &schema.Schema{
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "retentionperioddays": &schema.Schema{
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Default value is 30 days. Infinite retention takes precedence over retentionPeriodDays.",
                                                },
                                                "type": &schema.Schema{
                                                    Type:        schema.TypeString,
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
                        "retentionruletype": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Which type of retention rule should be used for the given backup destination",
                        },
                        "snaprecoverypoints": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Number of snap recovery points for snap copy for retention. Can be specified instead of retention period in Days for snap copy.",
                        },
                        "sourcecopy": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "id": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "fullbackuptypestocopy": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Which type of backup type should be copied for the given backup destination when backup type is not all jobs. Default is LAST while adding new backup destination.",
                        },
                        "useextendedretentionrules": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Use extended retention rules",
                        },
                        "backupstarttime": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Backup start time in seconds. The time is provided in unix time format.",
                        },
                        "overrideretentionsettings": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Tells if this copy should use storage pool retention period days or the retention defined for this copy. Set as true to use retention defined on this copy.",
                        },
                        "optimizeforinstantclone": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Flag to specify if primary storage is copy data management enabled.",
                        },
                        "netappcloudtarget": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Only for snap copy. Enabling this changes SVM Mapping  to NetApp cloud targets only.",
                        },
                        "mappings": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "vendor": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Snapshot vendors available for Snap Copy mappings",
                                    },
                                    "targetvendor": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "id": &schema.Schema{
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                            },
                                        },
                                    },
                                    "source": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "id": &schema.Schema{
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                            },
                                        },
                                    },
                                    "sourcevendor": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "id": &schema.Schema{
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                            },
                                        },
                                    },
                                    "target": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "id": &schema.Schema{
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
                        "issnapcopy": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Is this a snap copy? If isMirrorCopy is not set, then default is Vault/Replica.",
                        },
                        "storagetype": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "region": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "id": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "storagepool": &schema.Schema{
                            Type:        schema.TypeList,
                            Required:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "id": &schema.Schema{
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
            "overriderestrictions": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rpo": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "backupcontent": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "storagepool": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
            },
            "snapshotoptions": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "This feature applies only to File System Agents",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "retentionperioddays": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Retention period in days. -1 can be specified for infinite retention. If this and snapRecoveryPoints both are not specified, this takes precedence.",
                        },
                        "snaprecoverypoints": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Number of snap recovery points for default snap copy for retention. Can be specified instead of retention period in Days for default snap copy.",
                        },
                        "enablebackupcopy": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Flag to enable backup copy",
                        },
                        "backupcopyrpomins": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Backup copy RPO in minutes",
                        },
                    },
                },
            },
            "parentplan": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "id": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
            },
            "additionalproperties": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rpo": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "The least RPO in minutes for the plan",
                        },
                        "addons": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "filesystem": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "File system options should be shown with this plan",
                                    },
                                    "indexcopy": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Index copy options should be shown with this plan",
                                    },
                                    "database": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Database options should be shown with this plan",
                                    },
                                    "snapstatus": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "snap": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Snap options should be shown with this plan",
                                    },
                                },
                            },
                        },
                        "status": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
            },
            "regiontoconfigure": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "id": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
            },
            "newname": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "New plan name to update",
            },
            "overrideinheritsettings": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "This feature applies to derived plans when inherit mode is optional.Provides user to set entity preference between parent and derived plan.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rpo": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Flag to specify if parent or derived plan rpo should be used when inherit mode is optional. True - derived, False - Base.",
                        },
                        "backupcontent": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Flag to specify if parent or derived plan backupContent should be used when inherit mode is optional. True - derived, False - Base.",
                        },
                        "backupdestination": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Flag to specify if parent or derived plan backupDestination should be used when inherit mode is optional. True - derived, False - Base.",
                        },
                    },
                },
            },
        },
    }
}

func resourceCreateplan_v2(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/ServerPlan
    var response_id = strconv.Itoa(0)
    var t_settings *handler.MsgServerPlanSettings
    if v, ok := d.GetOk("settings"); ok {
        val := v.([]interface{})
        t_settings = build_plan_v2_msgserverplansettings(d, val)
    }
    var t_backupcontent *handler.MsgPlanContent
    if v, ok := d.GetOk("backupcontent"); ok {
        val := v.([]interface{})
        t_backupcontent = build_plan_v2_msgplancontent(d, val)
    }
    var t_databaseoptions *handler.MsgServerPlanDatabaseOptions
    if v, ok := d.GetOk("databaseoptions"); ok {
        val := v.([]interface{})
        t_databaseoptions = build_plan_v2_msgserverplandatabaseoptions(d, val)
    }
    var t_filesystemaddon *bool
    if v, ok := d.GetOkExists("filesystemaddon"); ok {
        val := v.(bool)
        t_filesystemaddon = new(bool)
        t_filesystemaddon = &val
    }
    var t_allowplanoverride *bool
    if v, ok := d.GetOkExists("allowplanoverride"); ok {
        val := v.(bool)
        t_allowplanoverride = new(bool)
        t_allowplanoverride = &val
    }
    var t_planname *string
    if v, ok := d.GetOk("planname"); ok {
        val := v.(string)
        t_planname = new(string)
        t_planname = &val
    }
    var t_workload *handler.MsgPlanWorkloads
    if v, ok := d.GetOk("workload"); ok {
        val := v.([]interface{})
        t_workload = build_plan_v2_msgplanworkloads(d, val)
    }
    var t_rpo *handler.MsgServerBackupPlanRPO
    if v, ok := d.GetOk("rpo"); ok {
        val := v.([]interface{})
        t_rpo = build_plan_v2_msgserverbackupplanrpo(d, val)
    }
    var t_backupdestinations []handler.MsgCreatePlanBackupDestinationSet
    if v, ok := d.GetOk("backupdestinations"); ok {
        val := v.(*schema.Set)
        t_backupdestinations = build_plan_v2_msgcreateplanbackupdestinationset_array(d, val.List())
    }
    var t_overriderestrictions *handler.MsgPlanOverrideSettings
    if v, ok := d.GetOk("overriderestrictions"); ok {
        val := v.([]interface{})
        t_overriderestrictions = build_plan_v2_msgplanoverridesettings(d, val)
    }
    var t_snapshotoptions *handler.MsgCreatePlanSnapshotOptions
    if v, ok := d.GetOk("snapshotoptions"); ok {
        val := v.([]interface{})
        t_snapshotoptions = build_plan_v2_msgcreateplansnapshotoptions(d, val)
    }
    var t_parentplan *handler.MsgIdName
    if v, ok := d.GetOk("parentplan"); ok {
        val := v.([]interface{})
        t_parentplan = build_plan_v2_msgidname(d, val)
    }
    var t_additionalproperties *handler.MsgPlanAdditionalProperties
    if v, ok := d.GetOk("additionalproperties"); ok {
        val := v.([]interface{})
        t_additionalproperties = build_plan_v2_msgplanadditionalproperties(d, val)
    }
    var req = handler.MsgCreateServerPlanRequest{Settings:t_settings, BackupContent:t_backupcontent, DatabaseOptions:t_databaseoptions, FilesystemAddon:t_filesystemaddon, AllowPlanOverride:t_allowplanoverride, PlanName:t_planname, Workload:t_workload, Rpo:t_rpo, BackupDestinations:t_backupdestinations, OverrideRestrictions:t_overriderestrictions, SnapshotOptions:t_snapshotoptions, ParentPlan:t_parentplan, AdditionalProperties:t_additionalproperties}
    resp, err := handler.CvCreateServerPlan(req)
    if err != nil {
        return fmt.Errorf("Operation [CreateServerPlan] failed, Error %s", err)
    }
    if resp.Plan != nil {
        if resp.Plan.Id != nil {
            response_id = strconv.Itoa(*resp.Plan.Id)
        }
    }
    if response_id == "0" {
        return fmt.Errorf("Operation [CreateServerPlan] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateplan_v2(d, m)
    }
}

func resourceReadplan_v2(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/ServerPlan/{planId}
    resp, err := handler.CvGetPlanById(d.Id())
    if err != nil {
        return fmt.Errorf("Operation [GetPlanById] failed, Error %s", err)
    }
    if resp.Settings != nil {
        d.Set("settings", serialize_plan_v2_msgserverplansettings(resp.Settings))
    } else {
        d.Set("settings", make([]map[string]interface{}, 0))
    }
    if resp.BackupContent != nil {
        d.Set("backupcontent", serialize_plan_v2_msgplancontent(resp.BackupContent))
    } else {
        d.Set("backupcontent", make([]map[string]interface{}, 0))
    }
    if resp.DatabaseOptions != nil {
        d.Set("databaseoptions", serialize_plan_v2_msgserverplandatabaseoptionsinfo(resp.DatabaseOptions))
    } else {
        d.Set("databaseoptions", make([]map[string]interface{}, 0))
    }
    if resp.AllowPlanOverride != nil {
        d.Set("allowplanoverride", resp.AllowPlanOverride)
    }
    if resp.Workload != nil {
        d.Set("workload", serialize_plan_v2_msgplanworkloads(resp.Workload))
    } else {
        d.Set("workload", make([]map[string]interface{}, 0))
    }
    if resp.Rpo != nil {
        d.Set("rpo", serialize_plan_v2_msgserverplanrpo(resp.Rpo))
    } else {
        d.Set("rpo", make([]map[string]interface{}, 0))
    }
    if resp.BackupDestinations != nil {
        d.Set("backupdestinations", serialize_plan_v2_msgplanbackupdestinationset_array(resp.BackupDestinations))
    } else {
        d.Set("backupdestinations", make([]map[string]interface{}, 0))
    }
    if resp.OverrideRestrictions != nil {
        d.Set("overriderestrictions", serialize_plan_v2_msgplanoverridesettings(resp.OverrideRestrictions))
    } else {
        d.Set("overriderestrictions", make([]map[string]interface{}, 0))
    }
    if resp.SnapshotOptions != nil {
        d.Set("snapshotoptions", serialize_plan_v2_msgplansnapshotoptions(resp.SnapshotOptions))
    } else {
        d.Set("snapshotoptions", make([]map[string]interface{}, 0))
    }
    if resp.AdditionalProperties != nil {
        d.Set("additionalproperties", serialize_plan_v2_msgplanadditionalproperties(resp.AdditionalProperties))
    } else {
        d.Set("additionalproperties", make([]map[string]interface{}, 0))
    }
    return nil
}

func resourceUpdateplan_v2(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/ServerPlan/{planId}
    var t_rpo *handler.MsgServerPlanUpdateRPO
    if d.HasChange("rpo") {
        val := d.Get("rpo").([]interface{})
        t_rpo = build_plan_v2_msgserverplanupdaterpo(d, val)
    }
    var t_regiontoconfigure *handler.MsgIdName
    if d.HasChange("regiontoconfigure") {
        val := d.Get("regiontoconfigure").([]interface{})
        t_regiontoconfigure = build_plan_v2_msgidname(d, val)
    }
    var t_settings *handler.MsgServerPlanSettings
    if d.HasChange("settings") {
        val := d.Get("settings").([]interface{})
        t_settings = build_plan_v2_msgserverplansettings(d, val)
    }
    var t_backupcontent *handler.MsgPlanContent
    if d.HasChange("backupcontent") {
        val := d.Get("backupcontent").([]interface{})
        t_backupcontent = build_plan_v2_msgplancontent(d, val)
    }
    var t_databaseoptions *handler.MsgServerPlanDatabaseOptionsInfo
    if d.HasChange("databaseoptions") {
        val := d.Get("databaseoptions").([]interface{})
        t_databaseoptions = build_plan_v2_msgserverplandatabaseoptionsinfo(d, val)
    }
    var t_newname *string
    if d.HasChange("newname") {
        val := d.Get("newname").(string)
        t_newname = new(string)
        t_newname = &val
    }
    var t_overrideinheritsettings *handler.MsgPlanOverrideInheritSetting
    if d.HasChange("overrideinheritsettings") {
        val := d.Get("overrideinheritsettings").([]interface{})
        t_overrideinheritsettings = build_plan_v2_msgplanoverrideinheritsetting(d, val)
    }
    var t_filesystemaddon *bool
    if d.HasChange("filesystemaddon") {
        val := d.Get("filesystemaddon").(bool)
        t_filesystemaddon = new(bool)
        t_filesystemaddon = &val
    }
    var t_allowplanoverride *bool
    if d.HasChange("allowplanoverride") {
        val := d.Get("allowplanoverride").(bool)
        t_allowplanoverride = new(bool)
        t_allowplanoverride = &val
    }
    var t_workload *handler.MsgPlanWorkloads
    if d.HasChange("workload") {
        val := d.Get("workload").([]interface{})
        t_workload = build_plan_v2_msgplanworkloads(d, val)
    }
    var t_overriderestrictions *handler.MsgPlanOverrideSettings
    if d.HasChange("overriderestrictions") {
        val := d.Get("overriderestrictions").([]interface{})
        t_overriderestrictions = build_plan_v2_msgplanoverridesettings(d, val)
    }
    var t_snapshotoptions *handler.MsgPlanSnapshotOptions
    if d.HasChange("snapshotoptions") {
        val := d.Get("snapshotoptions").([]interface{})
        t_snapshotoptions = build_plan_v2_msgplansnapshotoptions(d, val)
    }
    var req = handler.MsgModifyPlanRequest{Rpo:t_rpo, RegionToConfigure:t_regiontoconfigure, Settings:t_settings, BackupContent:t_backupcontent, DatabaseOptions:t_databaseoptions, NewName:t_newname, OverrideInheritSettings:t_overrideinheritsettings, FilesystemAddon:t_filesystemaddon, AllowPlanOverride:t_allowplanoverride, Workload:t_workload, OverrideRestrictions:t_overriderestrictions, SnapshotOptions:t_snapshotoptions}
    _, err := handler.CvModifyPlan(req, d.Id())
    if err != nil {
        return fmt.Errorf("Operation [ModifyPlan] failed, Error %s", err)
    }
    return resourceReadplan_v2(d, m)
}

func resourceCreateUpdateplan_v2(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/ServerPlan/{planId}
    var execUpdate bool = false
    var t_regiontoconfigure *handler.MsgIdName
    if v, ok := d.GetOk("regiontoconfigure"); ok {
        val := v.([]interface{})
        t_regiontoconfigure = build_plan_v2_msgidname(d, val)
        execUpdate = true
    }
    var t_newname *string
    if v, ok := d.GetOk("newname"); ok {
        val := v.(string)
        t_newname = new(string)
        t_newname = &val
        execUpdate = true
    }
    var t_overrideinheritsettings *handler.MsgPlanOverrideInheritSetting
    if v, ok := d.GetOk("overrideinheritsettings"); ok {
        val := v.([]interface{})
        t_overrideinheritsettings = build_plan_v2_msgplanoverrideinheritsetting(d, val)
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgModifyPlanRequest{RegionToConfigure:t_regiontoconfigure, NewName:t_newname, OverrideInheritSettings:t_overrideinheritsettings}
        _, err := handler.CvModifyPlan(req, d.Id())
        if err != nil {
            return fmt.Errorf("Operation [ModifyPlan] failed, Error %s", err)
        }
    }
    return resourceReadplan_v2(d, m)
}

func resourceDeleteplan_v2(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/ServerPlan/{planId}
    _, err := handler.CvDeletePlan(d.Id())
    if err != nil {
        return fmt.Errorf("Operation [DeletePlan] failed, Error %s", err)
    }
    return nil
}

func build_plan_v2_msgplanoverrideinheritsetting(d *schema.ResourceData, r []interface{}) *handler.MsgPlanOverrideInheritSetting {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_rpo *bool
        if val, ok := tmp["rpo"].(bool); ok {
            t_rpo = new(bool)
            t_rpo = &val
        }
        var t_backupcontent *bool
        if val, ok := tmp["backupcontent"].(bool); ok {
            t_backupcontent = new(bool)
            t_backupcontent = &val
        }
        var t_backupdestination *bool
        if val, ok := tmp["backupdestination"].(bool); ok {
            t_backupdestination = new(bool)
            t_backupdestination = &val
        }
        return &handler.MsgPlanOverrideInheritSetting{Rpo:t_rpo, BackupContent:t_backupcontent, BackupDestination:t_backupdestination}
    } else {
        return nil
    }
}

func build_plan_v2_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_name *string
        if val, ok := tmp["name"].(string); ok {
            t_name = new(string)
            t_name = &val
        }
        var t_id *int
        if val, ok := tmp["id"].(int); ok {
            if !handler.IsEmptyInt(val) {
                t_id = new(int)
                t_id = &val
            }
        }
        return &handler.MsgIdName{Name:t_name, Id:t_id}
    } else {
        return nil
    }
}

func build_plan_v2_msgplansnapshotoptions(d *schema.ResourceData, r []interface{}) *handler.MsgPlanSnapshotOptions {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_backupcopyfrequency *handler.MsgBackupFrequencyPattern
        if val, ok := tmp["backupcopyfrequency"].([]interface{}); ok {
            t_backupcopyfrequency = build_plan_v2_msgbackupfrequencypattern(d, val)
        }
        var t_enablebackupcopy *bool
        if val, ok := tmp["enablebackupcopy"].(bool); ok {
            t_enablebackupcopy = new(bool)
            t_enablebackupcopy = &val
        }
        var t_backupcopyrpomins *int
        if val, ok := tmp["backupcopyrpomins"].(int); ok {
            t_backupcopyrpomins = new(int)
            t_backupcopyrpomins = &val
        }
        return &handler.MsgPlanSnapshotOptions{BackupCopyFrequency:t_backupcopyfrequency, EnableBackupCopy:t_enablebackupcopy, BackupCopyRPOMins:t_backupcopyrpomins}
    } else {
        return nil
    }
}

func build_plan_v2_msgbackupfrequencypattern(d *schema.ResourceData, r []interface{}) *handler.MsgBackupFrequencyPattern {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_weeklydays []string
        if val, ok := tmp["weeklydays"].(*schema.Set); ok {
            t_weeklydays = handler.ToStringArray(val.List())
        }
        var t_monthofyear *string
        if val, ok := tmp["monthofyear"].(string); ok {
            t_monthofyear = new(string)
            t_monthofyear = &val
        }
        var t_dayofweek *string
        if val, ok := tmp["dayofweek"].(string); ok {
            t_dayofweek = new(string)
            t_dayofweek = &val
        }
        var t_dayofmonth *int
        if val, ok := tmp["dayofmonth"].(int); ok {
            t_dayofmonth = new(int)
            t_dayofmonth = &val
        }
        var t_schedulefrequencytype *string
        if val, ok := tmp["schedulefrequencytype"].(string); ok {
            t_schedulefrequencytype = new(string)
            t_schedulefrequencytype = &val
        }
        var t_weekofmonth *string
        if val, ok := tmp["weekofmonth"].(string); ok {
            t_weekofmonth = new(string)
            t_weekofmonth = &val
        }
        var t_starttime *int
        if val, ok := tmp["starttime"].(int); ok {
            t_starttime = new(int)
            t_starttime = &val
        }
        var t_frequency *int
        if val, ok := tmp["frequency"].(int); ok {
            t_frequency = new(int)
            t_frequency = &val
        }
        return &handler.MsgBackupFrequencyPattern{WeeklyDays:t_weeklydays, MonthOfYear:t_monthofyear, DayOfWeek:t_dayofweek, DayOfMonth:t_dayofmonth, ScheduleFrequencyType:t_schedulefrequencytype, WeekOfMonth:t_weekofmonth, StartTime:t_starttime, Frequency:t_frequency}
    } else {
        return nil
    }
}

func build_plan_v2_msgplanoverridesettings(d *schema.ResourceData, r []interface{}) *handler.MsgPlanOverrideSettings {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_rpo *string
        if val, ok := tmp["rpo"].(string); ok {
            t_rpo = new(string)
            t_rpo = &val
        }
        var t_backupcontent *string
        if val, ok := tmp["backupcontent"].(string); ok {
            t_backupcontent = new(string)
            t_backupcontent = &val
        }
        var t_storagepool *string
        if val, ok := tmp["storagepool"].(string); ok {
            t_storagepool = new(string)
            t_storagepool = &val
        }
        return &handler.MsgPlanOverrideSettings{RPO:t_rpo, BackupContent:t_backupcontent, StoragePool:t_storagepool}
    } else {
        return nil
    }
}

func build_plan_v2_msgplanworkloads(d *schema.ResourceData, r []interface{}) *handler.MsgPlanWorkloads {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_workloadtypes []handler.MsgIdNameSet
        if val, ok := tmp["workloadtypes"].(*schema.Set); ok {
            t_workloadtypes = build_plan_v2_msgidnameset_array(d, val.List())
        }
        var t_workloadgrouptypes []string
        if val, ok := tmp["workloadgrouptypes"].(*schema.Set); ok {
            t_workloadgrouptypes = handler.ToStringArray(val.List())
        }
        var t_solutions []handler.MsgIdNameSet
        if val, ok := tmp["solutions"].(*schema.Set); ok {
            t_solutions = build_plan_v2_msgidnameset_array(d, val.List())
        }
        return &handler.MsgPlanWorkloads{WorkloadTypes:t_workloadtypes, WorkloadGroupTypes:t_workloadgrouptypes, Solutions:t_solutions}
    } else {
        return nil
    }
}

func build_plan_v2_msgidnameset_array(d *schema.ResourceData, r []interface{}) []handler.MsgIdNameSet {
    if r != nil {
        tmp := make([]handler.MsgIdNameSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_id *int
            if val, ok := raw_a["id"].(int); ok {
                if !handler.IsEmptyInt(val) {
                    t_id = new(int)
                    t_id = &val
                }
            }
            tmp[a] = handler.MsgIdNameSet{Id:t_id}
        }
        return tmp
    } else {
        return nil
    }
}

func build_plan_v2_msgserverplandatabaseoptionsinfo(d *schema.ResourceData, r []interface{}) *handler.MsgServerPlanDatabaseOptionsInfo {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_logbackuprpomins *int
        if val, ok := tmp["logbackuprpomins"].(int); ok {
            t_logbackuprpomins = new(int)
            t_logbackuprpomins = &val
        }
        var t_runfullbackupevery *int
        if val, ok := tmp["runfullbackupevery"].(int); ok {
            t_runfullbackupevery = new(int)
            t_runfullbackupevery = &val
        }
        var t_commitfrequencyinhours *int
        if val, ok := tmp["commitfrequencyinhours"].(int); ok {
            t_commitfrequencyinhours = new(int)
            t_commitfrequencyinhours = &val
        }
        var t_usediskcacheforlogbackups *bool
        if val, ok := tmp["usediskcacheforlogbackups"].(bool); ok {
            t_usediskcacheforlogbackups = new(bool)
            t_usediskcacheforlogbackups = &val
        }
        return &handler.MsgServerPlanDatabaseOptionsInfo{LogBackupRPOMins:t_logbackuprpomins, RunFullBackupEvery:t_runfullbackupevery, CommitFrequencyInHours:t_commitfrequencyinhours, UseDiskCacheForLogBackups:t_usediskcacheforlogbackups}
    } else {
        return nil
    }
}

func build_plan_v2_msgplancontent(d *schema.ResourceData, r []interface{}) *handler.MsgPlanContent {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_windowsincludedpaths []string
        if val, ok := tmp["windowsincludedpaths"].(*schema.Set); ok {
            t_windowsincludedpaths = handler.ToStringArray(val.List())
        }
        var t_backupsystemstate *bool
        if val, ok := tmp["backupsystemstate"].(bool); ok {
            t_backupsystemstate = new(bool)
            t_backupsystemstate = &val
        }
        var t_backupsystemstateonlywithfullbackup *bool
        if val, ok := tmp["backupsystemstateonlywithfullbackup"].(bool); ok {
            t_backupsystemstateonlywithfullbackup = new(bool)
            t_backupsystemstateonlywithfullbackup = &val
        }
        var t_windowsexcludedpaths []string
        if val, ok := tmp["windowsexcludedpaths"].(*schema.Set); ok {
            t_windowsexcludedpaths = handler.ToStringArray(val.List())
        }
        var t_unixincludedpaths []string
        if val, ok := tmp["unixincludedpaths"].(*schema.Set); ok {
            t_unixincludedpaths = handler.ToStringArray(val.List())
        }
        var t_usevssforsystemstate *bool
        if val, ok := tmp["usevssforsystemstate"].(bool); ok {
            t_usevssforsystemstate = new(bool)
            t_usevssforsystemstate = &val
        }
        var t_macexcludedpaths []string
        if val, ok := tmp["macexcludedpaths"].(*schema.Set); ok {
            t_macexcludedpaths = handler.ToStringArray(val.List())
        }
        var t_macfiltertoexcludepaths []string
        if val, ok := tmp["macfiltertoexcludepaths"].(*schema.Set); ok {
            t_macfiltertoexcludepaths = handler.ToStringArray(val.List())
        }
        var t_windowsfiltertoexcludepaths []string
        if val, ok := tmp["windowsfiltertoexcludepaths"].(*schema.Set); ok {
            t_windowsfiltertoexcludepaths = handler.ToStringArray(val.List())
        }
        var t_unixfiltertoexcludepaths []string
        if val, ok := tmp["unixfiltertoexcludepaths"].(*schema.Set); ok {
            t_unixfiltertoexcludepaths = handler.ToStringArray(val.List())
        }
        var t_macincludedpaths []string
        if val, ok := tmp["macincludedpaths"].(*schema.Set); ok {
            t_macincludedpaths = handler.ToStringArray(val.List())
        }
        var t_unixexcludedpaths []string
        if val, ok := tmp["unixexcludedpaths"].(*schema.Set); ok {
            t_unixexcludedpaths = handler.ToStringArray(val.List())
        }
        return &handler.MsgPlanContent{WindowsIncludedPaths:t_windowsincludedpaths, BackupSystemState:t_backupsystemstate, BackupSystemStateOnlyWithFullBackup:t_backupsystemstateonlywithfullbackup, WindowsExcludedPaths:t_windowsexcludedpaths, UnixIncludedPaths:t_unixincludedpaths, UseVSSForSystemState:t_usevssforsystemstate, MacExcludedPaths:t_macexcludedpaths, MacFilterToExcludePaths:t_macfiltertoexcludepaths, WindowsFilterToExcludePaths:t_windowsfiltertoexcludepaths, UnixFilterToExcludePaths:t_unixfiltertoexcludepaths, MacIncludedPaths:t_macincludedpaths, UnixExcludedPaths:t_unixexcludedpaths}
    } else {
        return nil
    }
}

func build_plan_v2_msgserverplansettings(d *schema.ResourceData, r []interface{}) *handler.MsgServerPlanSettings {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_enableadvancedview *bool
        if val, ok := tmp["enableadvancedview"].(bool); ok {
            t_enableadvancedview = new(bool)
            t_enableadvancedview = &val
        }
        var t_filesearch *handler.MsgPlanFileSearch
        if val, ok := tmp["filesearch"].([]interface{}); ok {
            t_filesearch = build_plan_v2_msgplanfilesearch(d, val)
        }
        return &handler.MsgServerPlanSettings{EnableAdvancedView:t_enableadvancedview, FileSearch:t_filesearch}
    } else {
        return nil
    }
}

func build_plan_v2_msgplanfilesearch(d *schema.ResourceData, r []interface{}) *handler.MsgPlanFileSearch {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_enabled *bool
        if val, ok := tmp["enabled"].(bool); ok {
            t_enabled = new(bool)
            t_enabled = &val
        }
        var t_statusmessage *string
        if val, ok := tmp["statusmessage"].(string); ok {
            t_statusmessage = new(string)
            t_statusmessage = &val
        }
        var t_status *string
        if val, ok := tmp["status"].(string); ok {
            t_status = new(string)
            t_status = &val
        }
        return &handler.MsgPlanFileSearch{Enabled:t_enabled, StatusMessage:t_statusmessage, Status:t_status}
    } else {
        return nil
    }
}

func build_plan_v2_msgserverplanupdaterpo(d *schema.ResourceData, r []interface{}) *handler.MsgServerPlanUpdateRPO {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_fullbackupwindow []handler.MsgDayAndTimeSet
        if val, ok := tmp["fullbackupwindow"].(*schema.Set); ok {
            t_fullbackupwindow = build_plan_v2_msgdayandtimeset_array(d, val.List())
        }
        var t_sla *handler.MsgSLAUpdateOptions
        if val, ok := tmp["sla"].([]interface{}); ok {
            t_sla = build_plan_v2_msgslaupdateoptions(d, val)
        }
        var t_backupfrequency *handler.MsgPlanSchedules
        if val, ok := tmp["backupfrequency"].([]interface{}); ok {
            t_backupfrequency = build_plan_v2_msgplanschedules(d, val)
        }
        var t_backupwindow []handler.MsgDayAndTimeSet
        if val, ok := tmp["backupwindow"].(*schema.Set); ok {
            t_backupwindow = build_plan_v2_msgdayandtimeset_array(d, val.List())
        }
        return &handler.MsgServerPlanUpdateRPO{FullBackupWindow:t_fullbackupwindow, SLA:t_sla, BackupFrequency:t_backupfrequency, BackupWindow:t_backupwindow}
    } else {
        return nil
    }
}

func build_plan_v2_msgdayandtimeset_array(d *schema.ResourceData, r []interface{}) []handler.MsgDayAndTimeSet {
    if r != nil {
        tmp := make([]handler.MsgDayAndTimeSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_dayofweek []string
            if val, ok := raw_a["dayofweek"].(*schema.Set); ok {
                t_dayofweek = handler.ToStringArray(val.List())
            }
            var t_starttime *int64
            if val, ok := raw_a["starttime"].(int64); ok {
                t_starttime = new(int64)
                t_starttime = &val
            }
            var t_endtime *int64
            if val, ok := raw_a["endtime"].(int64); ok {
                t_endtime = new(int64)
                t_endtime = &val
            }
            tmp[a] = handler.MsgDayAndTimeSet{DayOfWeek:t_dayofweek, StartTime:t_starttime, EndTime:t_endtime}
        }
        return tmp
    } else {
        return nil
    }
}

func build_plan_v2_msgplanschedules(d *schema.ResourceData, r []interface{}) *handler.MsgPlanSchedules {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_schedules []handler.MsgPlanScheduleSet
        if val, ok := tmp["schedules"].(*schema.Set); ok {
            t_schedules = build_plan_v2_msgplanscheduleset_array(d, val.List())
        }
        return &handler.MsgPlanSchedules{Schedules:t_schedules}
    } else {
        return nil
    }
}

func build_plan_v2_msgplanscheduleset_array(d *schema.ResourceData, r []interface{}) []handler.MsgPlanScheduleSet {
    if r != nil {
        tmp := make([]handler.MsgPlanScheduleSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_schedulename *string
            if val, ok := raw_a["schedulename"].(string); ok {
                t_schedulename = new(string)
                t_schedulename = &val
            }
            var t_policyid *int
            if val, ok := raw_a["policyid"].(int); ok {
                t_policyid = new(int)
                t_policyid = &val
            }
            var t_vmoperationtype *string
            if val, ok := raw_a["vmoperationtype"].(string); ok {
                t_vmoperationtype = new(string)
                t_vmoperationtype = &val
            }
            var t_fordatabasesonly *bool
            if val, ok := raw_a["fordatabasesonly"].(bool); ok {
                t_fordatabasesonly = new(bool)
                t_fordatabasesonly = &val
            }
            var t_scheduleoperation *string
            if val, ok := raw_a["scheduleoperation"].(string); ok {
                t_scheduleoperation = new(string)
                t_scheduleoperation = &val
            }
            var t_schedulepattern *handler.MsgSchedulePattern
            if val, ok := raw_a["schedulepattern"].([]interface{}); ok {
                t_schedulepattern = build_plan_v2_msgschedulepattern(d, val)
            }
            var t_scheduleid *int
            if val, ok := raw_a["scheduleid"].(int); ok {
                t_scheduleid = new(int)
                t_scheduleid = &val
            }
            var t_backuptype *string
            if val, ok := raw_a["backuptype"].(string); ok {
                t_backuptype = new(string)
                t_backuptype = &val
            }
            tmp[a] = handler.MsgPlanScheduleSet{ScheduleName:t_schedulename, PolicyId:t_policyid, VmOperationType:t_vmoperationtype, ForDatabasesOnly:t_fordatabasesonly, ScheduleOperation:t_scheduleoperation, SchedulePattern:t_schedulepattern, ScheduleId:t_scheduleid, BackupType:t_backuptype}
        }
        return tmp
    } else {
        return nil
    }
}

func build_plan_v2_msgschedulepattern(d *schema.ResourceData, r []interface{}) *handler.MsgSchedulePattern {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_enddate *int
        if val, ok := tmp["enddate"].(int); ok {
            t_enddate = new(int)
            t_enddate = &val
        }
        var t_timezone *handler.MsgIdName
        if val, ok := tmp["timezone"].([]interface{}); ok {
            t_timezone = build_plan_v2_msgidname(d, val)
        }
        var t_weekofmonth *string
        if val, ok := tmp["weekofmonth"].(string); ok {
            t_weekofmonth = new(string)
            t_weekofmonth = &val
        }
        var t_daysbetweensyntheticfulls *int
        if val, ok := tmp["daysbetweensyntheticfulls"].(int); ok {
            t_daysbetweensyntheticfulls = new(int)
            t_daysbetweensyntheticfulls = &val
        }
        var t_exceptions []handler.MsgScheduleRunExceptionSet
        if val, ok := tmp["exceptions"].(*schema.Set); ok {
            t_exceptions = build_plan_v2_msgschedulerunexceptionset_array(d, val.List())
        }
        var t_frequency *int
        if val, ok := tmp["frequency"].(int); ok {
            t_frequency = new(int)
            t_frequency = &val
        }
        var t_weeklydays []string
        if val, ok := tmp["weeklydays"].(*schema.Set); ok {
            t_weeklydays = handler.ToStringArray(val.List())
        }
        var t_repeatuntiltime *int
        if val, ok := tmp["repeatuntiltime"].(int); ok {
            t_repeatuntiltime = new(int)
            t_repeatuntiltime = &val
        }
        var t_monthofyear *string
        if val, ok := tmp["monthofyear"].(string); ok {
            t_monthofyear = new(string)
            t_monthofyear = &val
        }
        var t_dayofweek *string
        if val, ok := tmp["dayofweek"].(string); ok {
            t_dayofweek = new(string)
            t_dayofweek = &val
        }
        var t_dayofmonth *int
        if val, ok := tmp["dayofmonth"].(int); ok {
            t_dayofmonth = new(int)
            t_dayofmonth = &val
        }
        var t_schedulefrequencytype *string
        if val, ok := tmp["schedulefrequencytype"].(string); ok {
            t_schedulefrequencytype = new(string)
            t_schedulefrequencytype = &val
        }
        var t_starttime *int
        if val, ok := tmp["starttime"].(int); ok {
            t_starttime = new(int)
            t_starttime = &val
        }
        var t_nooftimes *int
        if val, ok := tmp["nooftimes"].(int); ok {
            t_nooftimes = new(int)
            t_nooftimes = &val
        }
        var t_repeatintervalinminutes *int
        if val, ok := tmp["repeatintervalinminutes"].(int); ok {
            t_repeatintervalinminutes = new(int)
            t_repeatintervalinminutes = &val
        }
        var t_startdate *int
        if val, ok := tmp["startdate"].(int); ok {
            t_startdate = new(int)
            t_startdate = &val
        }
        return &handler.MsgSchedulePattern{EndDate:t_enddate, Timezone:t_timezone, WeekOfMonth:t_weekofmonth, DaysBetweenSyntheticFulls:t_daysbetweensyntheticfulls, Exceptions:t_exceptions, Frequency:t_frequency, WeeklyDays:t_weeklydays, RepeatUntilTime:t_repeatuntiltime, MonthOfYear:t_monthofyear, DayOfWeek:t_dayofweek, DayOfMonth:t_dayofmonth, ScheduleFrequencyType:t_schedulefrequencytype, StartTime:t_starttime, NoOfTimes:t_nooftimes, RepeatIntervalInMinutes:t_repeatintervalinminutes, StartDate:t_startdate}
    } else {
        return nil
    }
}

func build_plan_v2_msgschedulerunexceptionset_array(d *schema.ResourceData, r []interface{}) []handler.MsgScheduleRunExceptionSet {
    if r != nil {
        tmp := make([]handler.MsgScheduleRunExceptionSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_onweekofthemonth []string
            if val, ok := raw_a["onweekofthemonth"].(*schema.Set); ok {
                t_onweekofthemonth = handler.ToStringArray(val.List())
            }
            var t_ondates []int
            if val, ok := raw_a["ondates"].(*schema.Set); ok {
                t_ondates = handler.ToIntArray(val.List())
            }
            var t_ondayoftheweek []string
            if val, ok := raw_a["ondayoftheweek"].(*schema.Set); ok {
                t_ondayoftheweek = handler.ToStringArray(val.List())
            }
            tmp[a] = handler.MsgScheduleRunExceptionSet{OnWeekOfTheMonth:t_onweekofthemonth, OnDates:t_ondates, OnDayOfTheWeek:t_ondayoftheweek}
        }
        return tmp
    } else {
        return nil
    }
}

func build_plan_v2_msgslaupdateoptions(d *schema.ResourceData, r []interface{}) *handler.MsgSLAUpdateOptions {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_exclusionreason *string
        if val, ok := tmp["exclusionreason"].(string); ok {
            t_exclusionreason = new(string)
            t_exclusionreason = &val
        }
        var t_usesystemdefaultsla *bool
        if val, ok := tmp["usesystemdefaultsla"].(bool); ok {
            t_usesystemdefaultsla = new(bool)
            t_usesystemdefaultsla = &val
        }
        var t_enableafterdelay *int
        if val, ok := tmp["enableafterdelay"].(int); ok {
            t_enableafterdelay = new(int)
            t_enableafterdelay = &val
        }
        var t_excludefromsla *bool
        if val, ok := tmp["excludefromsla"].(bool); ok {
            t_excludefromsla = new(bool)
            t_excludefromsla = &val
        }
        var t_slaperiod *int
        if val, ok := tmp["slaperiod"].(int); ok {
            t_slaperiod = new(int)
            t_slaperiod = &val
        }
        return &handler.MsgSLAUpdateOptions{ExclusionReason:t_exclusionreason, UseSystemDefaultSLA:t_usesystemdefaultsla, EnableAfterDelay:t_enableafterdelay, ExcludeFromSLA:t_excludefromsla, SLAPeriod:t_slaperiod}
    } else {
        return nil
    }
}

func build_plan_v2_msgplanadditionalproperties(d *schema.ResourceData, r []interface{}) *handler.MsgPlanAdditionalProperties {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_rpo *int
        if val, ok := tmp["rpo"].(int); ok {
            t_rpo = new(int)
            t_rpo = &val
        }
        var t_addons *handler.MsgPlanAddons
        if val, ok := tmp["addons"].([]interface{}); ok {
            t_addons = build_plan_v2_msgplanaddons(d, val)
        }
        var t_status *string
        if val, ok := tmp["status"].(string); ok {
            t_status = new(string)
            t_status = &val
        }
        return &handler.MsgPlanAdditionalProperties{RPO:t_rpo, Addons:t_addons, Status:t_status}
    } else {
        return nil
    }
}

func build_plan_v2_msgplanaddons(d *schema.ResourceData, r []interface{}) *handler.MsgPlanAddons {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_filesystem *bool
        if val, ok := tmp["filesystem"].(bool); ok {
            t_filesystem = new(bool)
            t_filesystem = &val
        }
        var t_indexcopy *bool
        if val, ok := tmp["indexcopy"].(bool); ok {
            t_indexcopy = new(bool)
            t_indexcopy = &val
        }
        var t_database *bool
        if val, ok := tmp["database"].(bool); ok {
            t_database = new(bool)
            t_database = &val
        }
        var t_snapstatus *string
        if val, ok := tmp["snapstatus"].(string); ok {
            t_snapstatus = new(string)
            t_snapstatus = &val
        }
        var t_snap *bool
        if val, ok := tmp["snap"].(bool); ok {
            t_snap = new(bool)
            t_snap = &val
        }
        return &handler.MsgPlanAddons{FileSystem:t_filesystem, IndexCopy:t_indexcopy, Database:t_database, SnapStatus:t_snapstatus, Snap:t_snap}
    } else {
        return nil
    }
}

func build_plan_v2_msgcreateplansnapshotoptions(d *schema.ResourceData, r []interface{}) *handler.MsgCreatePlanSnapshotOptions {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_retentionperioddays *int
        if val, ok := tmp["retentionperioddays"].(int); ok {
            t_retentionperioddays = new(int)
            t_retentionperioddays = &val
        }
        var t_snaprecoverypoints *int
        if val, ok := tmp["snaprecoverypoints"].(int); ok {
            t_snaprecoverypoints = new(int)
            t_snaprecoverypoints = &val
        }
        var t_enablebackupcopy *bool
        if val, ok := tmp["enablebackupcopy"].(bool); ok {
            t_enablebackupcopy = new(bool)
            t_enablebackupcopy = &val
        }
        var t_backupcopyrpomins *int
        if val, ok := tmp["backupcopyrpomins"].(int); ok {
            t_backupcopyrpomins = new(int)
            t_backupcopyrpomins = &val
        }
        return &handler.MsgCreatePlanSnapshotOptions{RetentionPeriodDays:t_retentionperioddays, SnapRecoveryPoints:t_snaprecoverypoints, EnableBackupCopy:t_enablebackupcopy, BackupCopyRPOMins:t_backupcopyrpomins}
    } else {
        return nil
    }
}

func build_plan_v2_msgcreateplanbackupdestinationset_array(d *schema.ResourceData, r []interface{}) []handler.MsgCreatePlanBackupDestinationSet {
    if r != nil {
        tmp := make([]handler.MsgCreatePlanBackupDestinationSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_ismirrorcopy *bool
            if val, ok := raw_a["ismirrorcopy"].(bool); ok {
                t_ismirrorcopy = new(bool)
                t_ismirrorcopy = &val
            }
            var t_retentionperioddays *int
            if val, ok := raw_a["retentionperioddays"].(int); ok {
                t_retentionperioddays = new(int)
                t_retentionperioddays = &val
            }
            var t_backupstocopy *string
            if val, ok := raw_a["backupstocopy"].(string); ok {
                t_backupstocopy = new(string)
                t_backupstocopy = &val
            }
            var t_backupdestinationname *string
            if val, ok := raw_a["backupdestinationname"].(string); ok {
                t_backupdestinationname = new(string)
                t_backupdestinationname = &val
            }
            var t_extendedretentionrules *handler.MsgExtendedRetentionRules
            if val, ok := raw_a["extendedretentionrules"].([]interface{}); ok {
                t_extendedretentionrules = build_plan_v2_msgextendedretentionrules(d, val)
            }
            var t_retentionruletype *string
            if val, ok := raw_a["retentionruletype"].(string); ok {
                t_retentionruletype = new(string)
                t_retentionruletype = &val
            }
            var t_snaprecoverypoints *int
            if val, ok := raw_a["snaprecoverypoints"].(int); ok {
                t_snaprecoverypoints = new(int)
                t_snaprecoverypoints = &val
            }
            var t_sourcecopy *handler.MsgIdName
            if val, ok := raw_a["sourcecopy"].([]interface{}); ok {
                t_sourcecopy = build_plan_v2_msgidname(d, val)
            }
            var t_fullbackuptypestocopy *string
            if val, ok := raw_a["fullbackuptypestocopy"].(string); ok {
                t_fullbackuptypestocopy = new(string)
                t_fullbackuptypestocopy = &val
            }
            var t_useextendedretentionrules *bool
            if val, ok := raw_a["useextendedretentionrules"].(bool); ok {
                t_useextendedretentionrules = new(bool)
                t_useextendedretentionrules = &val
            }
            var t_backupstarttime *int
            if val, ok := raw_a["backupstarttime"].(int); ok {
                t_backupstarttime = new(int)
                t_backupstarttime = &val
            }
            var t_overrideretentionsettings *bool
            if val, ok := raw_a["overrideretentionsettings"].(bool); ok {
                t_overrideretentionsettings = new(bool)
                t_overrideretentionsettings = &val
            }
            var t_optimizeforinstantclone *bool
            if val, ok := raw_a["optimizeforinstantclone"].(bool); ok {
                t_optimizeforinstantclone = new(bool)
                t_optimizeforinstantclone = &val
            }
            var t_netappcloudtarget *bool
            if val, ok := raw_a["netappcloudtarget"].(bool); ok {
                t_netappcloudtarget = new(bool)
                t_netappcloudtarget = &val
            }
            var t_mappings []handler.MsgSnapshotCopyMappingSet
            if val, ok := raw_a["mappings"].(*schema.Set); ok {
                t_mappings = build_plan_v2_msgsnapshotcopymappingset_array(d, val.List())
            }
            var t_issnapcopy *bool
            if val, ok := raw_a["issnapcopy"].(bool); ok {
                t_issnapcopy = new(bool)
                t_issnapcopy = &val
            }
            var t_storagetype *string
            if val, ok := raw_a["storagetype"].(string); ok {
                t_storagetype = new(string)
                t_storagetype = &val
            }
            var t_region *handler.MsgIdName
            if val, ok := raw_a["region"].([]interface{}); ok {
                t_region = build_plan_v2_msgidname(d, val)
            }
            var t_storagepool *handler.MsgIdName
            if val, ok := raw_a["storagepool"].([]interface{}); ok {
                t_storagepool = build_plan_v2_msgidname(d, val)
            }
            tmp[a] = handler.MsgCreatePlanBackupDestinationSet{IsMirrorCopy:t_ismirrorcopy, RetentionPeriodDays:t_retentionperioddays, BackupsToCopy:t_backupstocopy, BackupDestinationName:t_backupdestinationname, ExtendedRetentionRules:t_extendedretentionrules, RetentionRuleType:t_retentionruletype, SnapRecoveryPoints:t_snaprecoverypoints, SourceCopy:t_sourcecopy, FullBackupTypesToCopy:t_fullbackuptypestocopy, UseExtendedRetentionRules:t_useextendedretentionrules, BackupStartTime:t_backupstarttime, OverrideRetentionSettings:t_overrideretentionsettings, OptimizeForInstantClone:t_optimizeforinstantclone, NetAppCloudTarget:t_netappcloudtarget, Mappings:t_mappings, IsSnapCopy:t_issnapcopy, StorageType:t_storagetype, Region:t_region, StoragePool:t_storagepool}
        }
        return tmp
    } else {
        return nil
    }
}

func build_plan_v2_msgsnapshotcopymappingset_array(d *schema.ResourceData, r []interface{}) []handler.MsgSnapshotCopyMappingSet {
    if r != nil {
        tmp := make([]handler.MsgSnapshotCopyMappingSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_vendor *string
            if val, ok := raw_a["vendor"].(string); ok {
                t_vendor = new(string)
                t_vendor = &val
            }
            var t_targetvendor *handler.MsgIdName
            if val, ok := raw_a["targetvendor"].([]interface{}); ok {
                t_targetvendor = build_plan_v2_msgidname(d, val)
            }
            var t_source *handler.MsgIdName
            if val, ok := raw_a["source"].([]interface{}); ok {
                t_source = build_plan_v2_msgidname(d, val)
            }
            var t_sourcevendor *handler.MsgIdName
            if val, ok := raw_a["sourcevendor"].([]interface{}); ok {
                t_sourcevendor = build_plan_v2_msgidname(d, val)
            }
            var t_target *handler.MsgIdName
            if val, ok := raw_a["target"].([]interface{}); ok {
                t_target = build_plan_v2_msgidname(d, val)
            }
            tmp[a] = handler.MsgSnapshotCopyMappingSet{Vendor:t_vendor, TargetVendor:t_targetvendor, Source:t_source, SourceVendor:t_sourcevendor, Target:t_target}
        }
        return tmp
    } else {
        return nil
    }
}

func build_plan_v2_msgextendedretentionrules(d *schema.ResourceData, r []interface{}) *handler.MsgExtendedRetentionRules {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_thirdextendedretentionrule *handler.MsgPlanRetentionRule
        if val, ok := tmp["thirdextendedretentionrule"].([]interface{}); ok {
            t_thirdextendedretentionrule = build_plan_v2_msgplanretentionrule(d, val)
        }
        var t_firstextendedretentionrule *handler.MsgPlanRetentionRule
        if val, ok := tmp["firstextendedretentionrule"].([]interface{}); ok {
            t_firstextendedretentionrule = build_plan_v2_msgplanretentionrule(d, val)
        }
        var t_secondextendedretentionrule *handler.MsgPlanRetentionRule
        if val, ok := tmp["secondextendedretentionrule"].([]interface{}); ok {
            t_secondextendedretentionrule = build_plan_v2_msgplanretentionrule(d, val)
        }
        return &handler.MsgExtendedRetentionRules{ThirdExtendedRetentionRule:t_thirdextendedretentionrule, FirstExtendedRetentionRule:t_firstextendedretentionrule, SecondExtendedRetentionRule:t_secondextendedretentionrule}
    } else {
        return nil
    }
}

func build_plan_v2_msgplanretentionrule(d *schema.ResourceData, r []interface{}) *handler.MsgPlanRetentionRule {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_isinfiniteretention *bool
        if val, ok := tmp["isinfiniteretention"].(bool); ok {
            t_isinfiniteretention = new(bool)
            t_isinfiniteretention = &val
        }
        var t_retentionperioddays *int
        if val, ok := tmp["retentionperioddays"].(int); ok {
            t_retentionperioddays = new(int)
            t_retentionperioddays = &val
        }
        var t_type *string
        if val, ok := tmp["type"].(string); ok {
            t_type = new(string)
            t_type = &val
        }
        return &handler.MsgPlanRetentionRule{IsInfiniteRetention:t_isinfiniteretention, RetentionPeriodDays:t_retentionperioddays, Type:t_type}
    } else {
        return nil
    }
}

func build_plan_v2_msgserverbackupplanrpo(d *schema.ResourceData, r []interface{}) *handler.MsgServerBackupPlanRPO {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_fullbackupwindow []handler.MsgDayAndTimeSet
        if val, ok := tmp["fullbackupwindow"].(*schema.Set); ok {
            t_fullbackupwindow = build_plan_v2_msgdayandtimeset_array(d, val.List())
        }
        var t_backupfrequency *handler.MsgPlanSchedules
        if val, ok := tmp["backupfrequency"].([]interface{}); ok {
            t_backupfrequency = build_plan_v2_msgplanschedules(d, val)
        }
        var t_backupwindow []handler.MsgDayAndTimeSet
        if val, ok := tmp["backupwindow"].(*schema.Set); ok {
            t_backupwindow = build_plan_v2_msgdayandtimeset_array(d, val.List())
        }
        return &handler.MsgServerBackupPlanRPO{FullBackupWindow:t_fullbackupwindow, BackupFrequency:t_backupfrequency, BackupWindow:t_backupwindow}
    } else {
        return nil
    }
}

func build_plan_v2_msgserverplandatabaseoptions(d *schema.ResourceData, r []interface{}) *handler.MsgServerPlanDatabaseOptions {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_logbackuprpomins *int
        if val, ok := tmp["logbackuprpomins"].(int); ok {
            t_logbackuprpomins = new(int)
            t_logbackuprpomins = &val
        }
        var t_commitfrequencyinhours *int
        if val, ok := tmp["commitfrequencyinhours"].(int); ok {
            t_commitfrequencyinhours = new(int)
            t_commitfrequencyinhours = &val
        }
        var t_usediskcacheforlogbackups *bool
        if val, ok := tmp["usediskcacheforlogbackups"].(bool); ok {
            t_usediskcacheforlogbackups = new(bool)
            t_usediskcacheforlogbackups = &val
        }
        return &handler.MsgServerPlanDatabaseOptions{LogBackupRPOMins:t_logbackuprpomins, CommitFrequencyInHours:t_commitfrequencyinhours, UseDiskCacheForLogBackups:t_usediskcacheforlogbackups}
    } else {
        return nil
    }
}

func serialize_plan_v2_msgplanadditionalproperties(data *handler.MsgPlanAdditionalProperties) map[string]interface{} {
    val := make(map[string]interface{})
    if data.RPO != nil {
        val["rpo"] = data.RPO
    }
    if data.Addons != nil {
        val["addons"] = serialize_plan_v2_msgplanaddons(data.Addons)
    }
    if data.Status != nil {
        val["status"] = data.Status
    }
    return val
}

func serialize_plan_v2_msgplanaddons(data *handler.MsgPlanAddons) map[string]interface{} {
    val := make(map[string]interface{})
    if data.FileSystem != nil {
        val["filesystem"] = data.FileSystem
    }
    if data.IndexCopy != nil {
        val["indexcopy"] = data.IndexCopy
    }
    if data.Database != nil {
        val["database"] = data.Database
    }
    if data.SnapStatus != nil {
        val["snapstatus"] = data.SnapStatus
    }
    if data.Snap != nil {
        val["snap"] = data.Snap
    }
    return val
}

func serialize_plan_v2_msgplansnapshotoptions(data *handler.MsgPlanSnapshotOptions) map[string]interface{} {
    val := make(map[string]interface{})
    if data.BackupCopyFrequency != nil {
        val["backupcopyfrequency"] = serialize_plan_v2_msgbackupfrequencypattern(data.BackupCopyFrequency)
    }
    if data.EnableBackupCopy != nil {
        val["enablebackupcopy"] = data.EnableBackupCopy
    }
    if data.BackupCopyRPOMins != nil {
        val["backupcopyrpomins"] = data.BackupCopyRPOMins
    }
    return val
}

func serialize_plan_v2_msgbackupfrequencypattern(data *handler.MsgBackupFrequencyPattern) map[string]interface{} {
    val := make(map[string]interface{})
    if data.WeeklyDays != nil {
        val["weeklydays"] = data.WeeklyDays
    }
    if data.MonthOfYear != nil {
        val["monthofyear"] = data.MonthOfYear
    }
    if data.DayOfWeek != nil {
        val["dayofweek"] = data.DayOfWeek
    }
    if data.DayOfMonth != nil {
        val["dayofmonth"] = data.DayOfMonth
    }
    if data.ScheduleFrequencyType != nil {
        val["schedulefrequencytype"] = data.ScheduleFrequencyType
    }
    if data.WeekOfMonth != nil {
        val["weekofmonth"] = data.WeekOfMonth
    }
    if data.StartTime != nil {
        val["starttime"] = data.StartTime
    }
    if data.Frequency != nil {
        val["frequency"] = data.Frequency
    }
    return val
}

func serialize_plan_v2_msgplanoverridesettings(data *handler.MsgPlanOverrideSettings) map[string]interface{} {
    val := make(map[string]interface{})
    if data.RPO != nil {
        val["rpo"] = data.RPO
    }
    if data.BackupContent != nil {
        val["backupcontent"] = data.BackupContent
    }
    if data.StoragePool != nil {
        val["storagepool"] = data.StoragePool
    }
    return val
}

func serialize_plan_v2_msgplanbackupdestinationset_array(data []handler.MsgPlanBackupDestinationSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].IsMirrorCopy != nil {
            val[i]["ismirrorcopy"] = data[i].IsMirrorCopy
        }
        if data[i].CopyPrecedence != nil {
            val[i]["copyprecedence"] = data[i].CopyPrecedence
        }
        if data[i].RetentionPeriodDays != nil {
            val[i]["retentionperioddays"] = data[i].RetentionPeriodDays
        }
        if data[i].CopyTypeName != nil {
            val[i]["copytypename"] = data[i].CopyTypeName
        }
        if data[i].BackupsToCopy != nil {
            val[i]["backupstocopy"] = data[i].BackupsToCopy
        }
        if data[i].ExtendedRetentionRules != nil {
            val[i]["extendedretentionrules"] = serialize_plan_v2_msgextendedretentionrules(data[i].ExtendedRetentionRules)
        }
        if data[i].RetentionRuleType != nil {
            val[i]["retentionruletype"] = data[i].RetentionRuleType
        }
        if data[i].SnapRecoveryPoints != nil {
            val[i]["snaprecoverypoints"] = data[i].SnapRecoveryPoints
        }
        if data[i].SourceCopy != nil {
            val[i]["sourcecopy"] = serialize_plan_v2_msgidname(data[i].SourceCopy)
        }
        if data[i].FullBackupTypesToCopy != nil {
            val[i]["fullbackuptypestocopy"] = data[i].FullBackupTypesToCopy
        }
        if data[i].UseExtendedRetentionRules != nil {
            val[i]["useextendedretentionrules"] = data[i].UseExtendedRetentionRules
        }
        if data[i].BackupStartTime != nil {
            val[i]["backupstarttime"] = data[i].BackupStartTime
        }
        if data[i].OverrideRetentionSettings != nil {
            val[i]["overrideretentionsettings"] = data[i].OverrideRetentionSettings
        }
        if data[i].NetAppCloudTarget != nil {
            val[i]["netappcloudtarget"] = data[i].NetAppCloudTarget
        }
        if data[i].IsDefault != nil {
            val[i]["isdefault"] = data[i].IsDefault
        }
        if data[i].Mappings != nil {
            val[i]["mappings"] = serialize_plan_v2_msgsnapshotcopymappingset_array(data[i].Mappings)
        }
        if data[i].PlanBackupDestination != nil {
            val[i]["planbackupdestination"] = serialize_plan_v2_msgidname(data[i].PlanBackupDestination)
        }
        if data[i].IsSnapCopy != nil {
            val[i]["issnapcopy"] = data[i].IsSnapCopy
        }
        if data[i].CopyType != nil {
            val[i]["copytype"] = data[i].CopyType
        }
        if data[i].StorageType != nil {
            val[i]["storagetype"] = data[i].StorageType
        }
        if data[i].EnableDataAging != nil {
            val[i]["enabledataaging"] = data[i].EnableDataAging
        }
        if data[i].Region != nil {
            val[i]["region"] = serialize_plan_v2_msgidnamedisplayname(data[i].Region)
        }
        if data[i].StoragePool != nil {
            val[i]["storagepool"] = serialize_plan_v2_msgstoragepool(data[i].StoragePool)
        }
    }
    return val
}

func serialize_plan_v2_msgstoragepool(data *handler.MsgStoragePool) map[string]interface{} {
    val := make(map[string]interface{})
    if data.RetentionPeriodDays != nil {
        val["retentionperioddays"] = data.RetentionPeriodDays
    }
    if data.WormStoragePoolFlag != nil {
        val["wormstoragepoolflag"] = data.WormStoragePoolFlag
    }
    if data.Name != nil {
        val["name"] = data.Name
    }
    if data.Id != nil {
        val["id"] = data.Id
    }
    if data.Type != nil {
        val["type"] = data.Type
    }
    return val
}

func serialize_plan_v2_msgidnamedisplayname(data *handler.MsgIdNameDisplayName) map[string]interface{} {
    val := make(map[string]interface{})
    if data.DisplayName != nil {
        val["displayname"] = data.DisplayName
    }
    if data.Name != nil {
        val["name"] = data.Name
    }
    if data.Id != nil {
        val["id"] = data.Id
    }
    return val
}

func serialize_plan_v2_msgidname(data *handler.MsgIdName) map[string]interface{} {
    val := make(map[string]interface{})
    if data.Name != nil {
        val["name"] = data.Name
    }
    if data.Id != nil {
        val["id"] = data.Id
    }
    return val
}

func serialize_plan_v2_msgsnapshotcopymappingset_array(data []handler.MsgSnapshotCopyMappingSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].Vendor != nil {
            val[i]["vendor"] = data[i].Vendor
        }
        if data[i].TargetVendor != nil {
            val[i]["targetvendor"] = serialize_plan_v2_msgidname(data[i].TargetVendor)
        }
        if data[i].Source != nil {
            val[i]["source"] = serialize_plan_v2_msgidname(data[i].Source)
        }
        if data[i].SourceVendor != nil {
            val[i]["sourcevendor"] = serialize_plan_v2_msgidname(data[i].SourceVendor)
        }
        if data[i].Target != nil {
            val[i]["target"] = serialize_plan_v2_msgidname(data[i].Target)
        }
    }
    return val
}

func serialize_plan_v2_msgextendedretentionrules(data *handler.MsgExtendedRetentionRules) map[string]interface{} {
    val := make(map[string]interface{})
    if data.ThirdExtendedRetentionRule != nil {
        val["thirdextendedretentionrule"] = serialize_plan_v2_msgplanretentionrule(data.ThirdExtendedRetentionRule)
    }
    if data.FirstExtendedRetentionRule != nil {
        val["firstextendedretentionrule"] = serialize_plan_v2_msgplanretentionrule(data.FirstExtendedRetentionRule)
    }
    if data.SecondExtendedRetentionRule != nil {
        val["secondextendedretentionrule"] = serialize_plan_v2_msgplanretentionrule(data.SecondExtendedRetentionRule)
    }
    return val
}

func serialize_plan_v2_msgplanretentionrule(data *handler.MsgPlanRetentionRule) map[string]interface{} {
    val := make(map[string]interface{})
    if data.IsInfiniteRetention != nil {
        val["isinfiniteretention"] = data.IsInfiniteRetention
    }
    if data.RetentionPeriodDays != nil {
        val["retentionperioddays"] = data.RetentionPeriodDays
    }
    if data.Type != nil {
        val["type"] = data.Type
    }
    return val
}

func serialize_plan_v2_msgserverplanrpo(data *handler.MsgServerPlanRPO) map[string]interface{} {
    val := make(map[string]interface{})
    if data.FullBackupWindow != nil {
        val["fullbackupwindow"] = serialize_plan_v2_msgdayandtimeset_array(data.FullBackupWindow)
    }
    if data.SLA != nil {
        val["sla"] = serialize_plan_v2_msgslaoptions(data.SLA)
    }
    if data.BackupFrequency != nil {
        val["backupfrequency"] = serialize_plan_v2_msgplanschedules(data.BackupFrequency)
    }
    if data.BackupWindow != nil {
        val["backupwindow"] = serialize_plan_v2_msgdayandtimeset_array(data.BackupWindow)
    }
    return val
}

func serialize_plan_v2_msgdayandtimeset_array(data []handler.MsgDayAndTimeSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].DayOfWeek != nil {
            val[i]["dayofweek"] = data[i].DayOfWeek
        }
        if data[i].StartTime != nil {
            val[i]["starttime"] = data[i].StartTime
        }
        if data[i].EndTime != nil {
            val[i]["endtime"] = data[i].EndTime
        }
    }
    return val
}

func serialize_plan_v2_msgplanschedules(data *handler.MsgPlanSchedules) map[string]interface{} {
    val := make(map[string]interface{})
    if data.Schedules != nil {
        val["schedules"] = serialize_plan_v2_msgplanscheduleset_array(data.Schedules)
    }
    return val
}

func serialize_plan_v2_msgplanscheduleset_array(data []handler.MsgPlanScheduleSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].ScheduleName != nil {
            val[i]["schedulename"] = data[i].ScheduleName
        }
        if data[i].PolicyId != nil {
            val[i]["policyid"] = data[i].PolicyId
        }
        if data[i].VmOperationType != nil {
            val[i]["vmoperationtype"] = data[i].VmOperationType
        }
        if data[i].ForDatabasesOnly != nil {
            val[i]["fordatabasesonly"] = data[i].ForDatabasesOnly
        }
        if data[i].ScheduleOperation != nil {
            val[i]["scheduleoperation"] = data[i].ScheduleOperation
        }
        if data[i].SchedulePattern != nil {
            val[i]["schedulepattern"] = serialize_plan_v2_msgschedulepattern(data[i].SchedulePattern)
        }
        if data[i].ScheduleId != nil {
            val[i]["scheduleid"] = data[i].ScheduleId
        }
        if data[i].BackupType != nil {
            val[i]["backuptype"] = data[i].BackupType
        }
    }
    return val
}

func serialize_plan_v2_msgschedulepattern(data *handler.MsgSchedulePattern) map[string]interface{} {
    val := make(map[string]interface{})
    if data.EndDate != nil {
        val["enddate"] = data.EndDate
    }
    if data.Timezone != nil {
        val["timezone"] = serialize_plan_v2_msgidname(data.Timezone)
    }
    if data.WeekOfMonth != nil {
        val["weekofmonth"] = data.WeekOfMonth
    }
    if data.DaysBetweenSyntheticFulls != nil {
        val["daysbetweensyntheticfulls"] = data.DaysBetweenSyntheticFulls
    }
    if data.Exceptions != nil {
        val["exceptions"] = serialize_plan_v2_msgschedulerunexceptionset_array(data.Exceptions)
    }
    if data.Frequency != nil {
        val["frequency"] = data.Frequency
    }
    if data.WeeklyDays != nil {
        val["weeklydays"] = data.WeeklyDays
    }
    if data.RepeatUntilTime != nil {
        val["repeatuntiltime"] = data.RepeatUntilTime
    }
    if data.MonthOfYear != nil {
        val["monthofyear"] = data.MonthOfYear
    }
    if data.DayOfWeek != nil {
        val["dayofweek"] = data.DayOfWeek
    }
    if data.DayOfMonth != nil {
        val["dayofmonth"] = data.DayOfMonth
    }
    if data.ScheduleFrequencyType != nil {
        val["schedulefrequencytype"] = data.ScheduleFrequencyType
    }
    if data.StartTime != nil {
        val["starttime"] = data.StartTime
    }
    if data.NoOfTimes != nil {
        val["nooftimes"] = data.NoOfTimes
    }
    if data.RepeatIntervalInMinutes != nil {
        val["repeatintervalinminutes"] = data.RepeatIntervalInMinutes
    }
    if data.StartDate != nil {
        val["startdate"] = data.StartDate
    }
    return val
}

func serialize_plan_v2_msgschedulerunexceptionset_array(data []handler.MsgScheduleRunExceptionSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].OnWeekOfTheMonth != nil {
            val[i]["onweekofthemonth"] = data[i].OnWeekOfTheMonth
        }
        if data[i].OnDates != nil {
            val[i]["ondates"] = data[i].OnDates
        }
        if data[i].OnDayOfTheWeek != nil {
            val[i]["ondayoftheweek"] = data[i].OnDayOfTheWeek
        }
    }
    return val
}

func serialize_plan_v2_msgslaoptions(data *handler.MsgSLAOptions) map[string]interface{} {
    val := make(map[string]interface{})
    if data.ExclusionReason != nil {
        val["exclusionreason"] = data.ExclusionReason
    }
    if data.InheritedSLAPeriod != nil {
        val["inheritedslaperiod"] = data.InheritedSLAPeriod
    }
    if data.UseSystemDefaultSLA != nil {
        val["usesystemdefaultsla"] = data.UseSystemDefaultSLA
    }
    if data.EnableAfterDelay != nil {
        val["enableafterdelay"] = data.EnableAfterDelay
    }
    if data.InheritedFrom != nil {
        val["inheritedfrom"] = data.InheritedFrom
    }
    if data.ExcludeFromSLA != nil {
        val["excludefromsla"] = data.ExcludeFromSLA
    }
    if data.SLAPeriod != nil {
        val["slaperiod"] = data.SLAPeriod
    }
    return val
}

func serialize_plan_v2_msgplanworkloads(data *handler.MsgPlanWorkloads) map[string]interface{} {
    val := make(map[string]interface{})
    if data.WorkloadTypes != nil {
        val["workloadtypes"] = serialize_plan_v2_msgidnameset_array(data.WorkloadTypes)
    }
    if data.WorkloadGroupTypes != nil {
        val["workloadgrouptypes"] = data.WorkloadGroupTypes
    }
    if data.Solutions != nil {
        val["solutions"] = serialize_plan_v2_msgidnameset_array(data.Solutions)
    }
    return val
}

func serialize_plan_v2_msgidnameset_array(data []handler.MsgIdNameSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].Id != nil {
            val[i]["id"] = data[i].Id
        }
    }
    return val
}

func serialize_plan_v2_msgserverplandatabaseoptionsinfo(data *handler.MsgServerPlanDatabaseOptionsInfo) map[string]interface{} {
    val := make(map[string]interface{})
    if data.LogBackupRPOMins != nil {
        val["logbackuprpomins"] = data.LogBackupRPOMins
    }
    if data.RunFullBackupEvery != nil {
        val["runfullbackupevery"] = data.RunFullBackupEvery
    }
    if data.CommitFrequencyInHours != nil {
        val["commitfrequencyinhours"] = data.CommitFrequencyInHours
    }
    if data.UseDiskCacheForLogBackups != nil {
        val["usediskcacheforlogbackups"] = data.UseDiskCacheForLogBackups
    }
    return val
}

func serialize_plan_v2_msgplancontent(data *handler.MsgPlanContent) map[string]interface{} {
    val := make(map[string]interface{})
    if data.WindowsIncludedPaths != nil {
        val["windowsincludedpaths"] = data.WindowsIncludedPaths
    }
    if data.BackupSystemState != nil {
        val["backupsystemstate"] = data.BackupSystemState
    }
    if data.BackupSystemStateOnlyWithFullBackup != nil {
        val["backupsystemstateonlywithfullbackup"] = data.BackupSystemStateOnlyWithFullBackup
    }
    if data.WindowsExcludedPaths != nil {
        val["windowsexcludedpaths"] = data.WindowsExcludedPaths
    }
    if data.UnixIncludedPaths != nil {
        val["unixincludedpaths"] = data.UnixIncludedPaths
    }
    if data.UseVSSForSystemState != nil {
        val["usevssforsystemstate"] = data.UseVSSForSystemState
    }
    if data.MacExcludedPaths != nil {
        val["macexcludedpaths"] = data.MacExcludedPaths
    }
    if data.MacFilterToExcludePaths != nil {
        val["macfiltertoexcludepaths"] = data.MacFilterToExcludePaths
    }
    if data.WindowsFilterToExcludePaths != nil {
        val["windowsfiltertoexcludepaths"] = data.WindowsFilterToExcludePaths
    }
    if data.UnixFilterToExcludePaths != nil {
        val["unixfiltertoexcludepaths"] = data.UnixFilterToExcludePaths
    }
    if data.MacIncludedPaths != nil {
        val["macincludedpaths"] = data.MacIncludedPaths
    }
    if data.UnixExcludedPaths != nil {
        val["unixexcludedpaths"] = data.UnixExcludedPaths
    }
    return val
}

func serialize_plan_v2_msgserverplansettings(data *handler.MsgServerPlanSettings) map[string]interface{} {
    val := make(map[string]interface{})
    if data.EnableAdvancedView != nil {
        val["enableadvancedview"] = data.EnableAdvancedView
    }
    if data.FileSearch != nil {
        val["filesearch"] = serialize_plan_v2_msgplanfilesearch(data.FileSearch)
    }
    return val
}

func serialize_plan_v2_msgplanfilesearch(data *handler.MsgPlanFileSearch) map[string]interface{} {
    val := make(map[string]interface{})
    if data.Enabled != nil {
        val["enabled"] = data.Enabled
    }
    if data.StatusMessage != nil {
        val["statusmessage"] = data.StatusMessage
    }
    if data.Status != nil {
        val["status"] = data.Status
    }
    return val
}
