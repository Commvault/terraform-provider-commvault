package commvault

import (
    "fmt"
    "strconv"

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
            "schedule": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Describes the Schedule object for Data Classification Plan",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "policyid": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "pattern": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "weeklydays": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Days of the week for weekly frequency",
                                        Elem: &schema.Schema{
                                            Type:    schema.TypeString,
                                        },
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
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "schedule frequency type [MINUTES, DAILY, WEEKLY, MONTHLY, YEARLY]",
                                    },
                                    "weekofmonth": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Specific week of a month [FIRST, SECOND, THIRD, FOURTH, LAST]",
                                    },
                                    "starttime": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "start time of schedule in seconds for daily, weekly, monthly, yearly frequency",
                                    },
                                    "frequency": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Frequency of the schedule based on schedule frequency type eg. for Hours, value 2 is 2 hours, for Minutes, 30 is 30 minutes, for Daily, 2 is 2 days. for Monthly 2 is it repeats every 2 months",
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
                                    "startdate": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "start date of schedule in epoch format",
                                    },
                                    "exceptions": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Exceptions to when a schedule should not run, either in dates or week of month and day",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "onweekofthemonth": {
                                                    Type:        schema.TypeSet,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "On which week of month, for ex: FIRST, LAST",
                                                    Elem: &schema.Schema{
                                                        Type:    schema.TypeString,
                                                    },
                                                },
                                                "ondates": {
                                                    Type:        schema.TypeSet,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "list of dates in a month. For ex: 1, 20",
                                                    Elem: &schema.Schema{
                                                        Type:    schema.TypeInt,
                                                    },
                                                },
                                                "ondayoftheweek": {
                                                    Type:        schema.TypeSet,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "On which days, for ex: MONDAY, FRIDAY",
                                                    Elem: &schema.Schema{
                                                        Type:    schema.TypeString,
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "scheduleid": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
            },
            "contentindexing": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Create Content Indexing Policy Model for DC Plan",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "filefilters": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "includedoctypes": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "excludepaths": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Schema{
                                            Type:    schema.TypeString,
                                        },
                                    },
                                    "mindocsize": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "maxdocsize": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "copyprecedence": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Set Content Indexing Copy Precedence",
                        },
                        "searchtype": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Describes the Content Indexing Operation Type for Data Classification Plan [METADATA, METADATA_CONTENT, METADATA_CONTENT_PREVIEW]",
                        },
                        "contentlanguage": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "OCR Languages Supported By DC Plan [NONE, ENGLISH, HEBREW, SPANISH, FRENCH, ITALIAN, DANISH]",
                        },
                        "exactsearch": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Enable Exact Seach in Data Classification Plan",
                        },
                        "backupcopy": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "copyid": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "storagepoolid": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "extracttextfromimage": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Enable OCR in Data Classification Plan",
                        },
                    },
                },
            },
            "threatindicator": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Threat Indicator Model for Data Classification Plan",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "threatdetection": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "Threat Detection Model for Threat Indicator Policy",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "backupsize": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "flag to enable Backup Size detection",
                                    },
                                    "canaryfile": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "flag to enable Canary Files detection",
                                    },
                                    "fileextension": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "flag to enable File Extention detection",
                                    },
                                    "fileactivity": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "flag to enable File Activity detection",
                                    },
                                    "filetype": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "flag to enable File Type detection",
                                    },
                                },
                            },
                        },
                        "threatnexus": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "Threat Nexus Model for Threat Indicator Policy",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "networkanddatasecurity": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Network And Data Security Model for Threat Nexus Policy",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "netskope": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "flag to enable Netskope",
                                                },
                                                "darktrace": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "flag to enable Darktrace",
                                                },
                                                "crowdstrike": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "flag to enable Crowdstrike",
                                                },
                                            },
                                        },
                                    },
                                    "soar": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "SOAR Model for Threat Nexus Policy",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "paloalto_xsoar": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "flag to enable Palo Alto “ XSOAR",
                                                },
                                                "mssentinel": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "flag to enable MS Sentinel",
                                                },
                                                "splunk": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "flag to enable Splunk",
                                                },
                                            },
                                        },
                                    },
                                    "cspm_dspm": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "CSPM/DSPM Model for Threat Nexus Policy",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "acante": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "flag to enable Acante",
                                                },
                                                "dasera": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "flag to enable Dasera",
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "threatscan": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "Threat Scan Model for Threat Indicator Policy",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "yararules": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "yara rules configuration info",
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
                                    "yararule": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "flag to enable Yara Rule Analysis",
                                    },
                                    "iocrules": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "IOC rules configuration info",
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
                                    "filedataanalysis": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "flag to enable File Data Analysis",
                                    },
                                    "threatanalysis": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "flag to enable Threat Analysis",
                                    },
                                    "usecommvaulthashfeed": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "flag to enable usage of Commvault hash feed",
                                    },
                                    "usecustomhashfeed": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "flag to enable usage of custom hash feed",
                                    },
                                },
                            },
                        },
                        "accessnodesinfo": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "access nodes info configured within Threat Indicator Plan",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "accessnodesgroup": {
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
                                    "accessnodes": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "list of access nodes used for Threat Indicator Plan",
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
                        "correlationpolicyrules": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "Correlation policies for threat indicator plan",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "threatdetection": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Threat detection correlation rule",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "uselastjobstatus": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "enable threat detection based on last backup job",
                                                },
                                                "mindays": {
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "min days for threat detection rule",
                                                },
                                            },
                                        },
                                    },
                                    "threatnexus": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Threat nexus correlation rule",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "enablethreatnexus": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "enable threat nexus",
                                                },
                                                "mindays": {
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "min days for threat nexus rule",
                                                },
                                            },
                                        },
                                    },
                                    "threatscan": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Threat scan correlation rule",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "maxdays": {
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "max threshold days for threat scan to run",
                                                },
                                                "uselastjobstatus": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "consider if last threat scan detected anomolies or not",
                                                },
                                                "mindays": {
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "min days for threat scan to run",
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "enablesmartscanning": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Enable Smart Scanning for Threat detection plan",
                        },
                    },
                },
            },
            "application": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Content indexing target app [NOT_AVAILABLE, FILE_STORAGE_OPTIMIZATION, SENSITIVE_DATA_GOVERNANCE, CASE_MANAGER, CONTENT_INDEXING, THREAT_ANALYSIS, RISK_ANALYSIS, THREAT_INDICATOR]",
            },
            "indexserver": {
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
            "threatanalysis": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Enables\\Disables Threat Analysis support for DC Plan",
            },
            "name": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Name of Data Classification Plan",
            },
            "entitydetection": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Entity Extraction Model for Data Classification Plan",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "entities": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "Entity TagId for Extraction",
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
                        "classifiers": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "Classifier TagId For Extraction",
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
                        "copyprecendence": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
            },
            "contentanalyzer": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "Content Analyzer Id`s for Entity Detection",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: "",
                        },
                    },
                },
            },
        },
    }
}

func resourceCreateplan_v2(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/DCPlan
    var response_id = strconv.Itoa(0)
    var t_schedule *handler.MsgDCPlanJobSchedule
    if val, ok := d.GetOk("schedule"); ok {
        t_schedule = build_plan_v2_msgdcplanjobschedule(d, val.([]interface{}))
    }
    var t_contentindexing *handler.MsgDCPlanCIpolicy
    if val, ok := d.GetOk("contentindexing"); ok {
        t_contentindexing = build_plan_v2_msgdcplancipolicy(d, val.([]interface{}))
    }
    var t_threatindicator *handler.MsgThreatIndicatorPolicy
    if val, ok := d.GetOk("threatindicator"); ok {
        t_threatindicator = build_plan_v2_msgthreatindicatorpolicy(d, val.([]interface{}))
    }
    var t_application *string
    if val, ok := d.GetOk("application"); ok {
        t_application = handler.ToStringValue(val, false)
    }
    var t_indexserver *handler.MsgIdName
    if val, ok := d.GetOk("indexserver"); ok {
        t_indexserver = build_plan_v2_msgidname(d, val.([]interface{}))
    }
    var t_threatanalysis *bool
    if val, ok := d.GetOk("threatanalysis"); ok {
        t_threatanalysis = handler.ToBooleanValue(val, false)
    }
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_entitydetection *handler.MsgDCPlanEEPolicy
    if val, ok := d.GetOk("entitydetection"); ok {
        t_entitydetection = build_plan_v2_msgdcplaneepolicy(d, val.([]interface{}))
    }
    var t_contentanalyzer []handler.MsgIdNameSet
    if val, ok := d.GetOk("contentanalyzer"); ok {
        t_contentanalyzer = build_plan_v2_msgidnameset_array(d, val.(*schema.Set).List())
    }
    var req = handler.MsgCreateDCPlanRequest{Schedule:t_schedule, ContentIndexing:t_contentindexing, ThreatIndicator:t_threatindicator, Application:t_application, IndexServer:t_indexserver, ThreatAnalysis:t_threatanalysis, Name:t_name, EntityDetection:t_entitydetection, ContentAnalyzer:t_contentanalyzer}
    _, err := handler.CvCreateDCPlan(req)
    if err != nil {
        return fmt.Errorf("operation [CreateDCPlan] failed, Error %s", err)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateDCPlan] failed")
    } else {
        d.SetId(response_id)
        return resourceReadplan_v2(d, m)
    }
}

func resourceReadplan_v2(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceUpdateplan_v2(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceDeleteplan_v2(d *schema.ResourceData, m interface{}) error {
    return nil
}

func build_plan_v2_msgidnameset_array(d *schema.ResourceData, r []interface{}) []handler.MsgIdNameSet {
    if r != nil {
        tmp := make([]handler.MsgIdNameSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_id *int
            if val, ok := raw_a["id"]; ok {
                t_id = handler.ToIntValue(val, true)
            }
            tmp[a] = handler.MsgIdNameSet{Id:t_id}
        }
        return tmp
    } else {
        return nil
    }
}

func build_plan_v2_msgdcplaneepolicy(d *schema.ResourceData, r []interface{}) *handler.MsgDCPlanEEPolicy {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_entities []handler.MsgIdNameSet
        if val, ok := tmp["entities"]; ok {
            t_entities = build_plan_v2_msgidnameset_array(d, val.(*schema.Set).List())
        }
        var t_classifiers []handler.MsgIdNameSet
        if val, ok := tmp["classifiers"]; ok {
            t_classifiers = build_plan_v2_msgidnameset_array(d, val.(*schema.Set).List())
        }
        var t_copyprecendence *int
        if val, ok := tmp["copyprecendence"]; ok {
            t_copyprecendence = handler.ToIntValue(val, true)
        }
        return &handler.MsgDCPlanEEPolicy{Entities:t_entities, Classifiers:t_classifiers, CopyPrecendence:t_copyprecendence}
    } else {
        return nil
    }
}

func build_plan_v2_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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
        return &handler.MsgIdName{Name:t_name, Id:t_id}
    } else {
        return nil
    }
}

func build_plan_v2_msgthreatindicatorpolicy(d *schema.ResourceData, r []interface{}) *handler.MsgThreatIndicatorPolicy {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_threatdetection *handler.MsgThreatDetectionPolicy
        if val, ok := tmp["threatdetection"]; ok {
            t_threatdetection = build_plan_v2_msgthreatdetectionpolicy(d, val.([]interface{}))
        }
        var t_threatnexus *handler.MsgThreatNexusPolicy
        if val, ok := tmp["threatnexus"]; ok {
            t_threatnexus = build_plan_v2_msgthreatnexuspolicy(d, val.([]interface{}))
        }
        var t_threatscan *handler.MsgThreatScanPolicy
        if val, ok := tmp["threatscan"]; ok {
            t_threatscan = build_plan_v2_msgthreatscanpolicy(d, val.([]interface{}))
        }
        var t_accessnodesinfo *handler.MsgAccessNodesInfo
        if val, ok := tmp["accessnodesinfo"]; ok {
            t_accessnodesinfo = build_plan_v2_msgaccessnodesinfo(d, val.([]interface{}))
        }
        var t_correlationpolicyrules *handler.MsgCorrelationPolicyRule
        if val, ok := tmp["correlationpolicyrules"]; ok {
            t_correlationpolicyrules = build_plan_v2_msgcorrelationpolicyrule(d, val.([]interface{}))
        }
        var t_enablesmartscanning *bool
        if val, ok := tmp["enablesmartscanning"]; ok {
            t_enablesmartscanning = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgThreatIndicatorPolicy{ThreatDetection:t_threatdetection, ThreatNexus:t_threatnexus, ThreatScan:t_threatscan, AccessNodesInfo:t_accessnodesinfo, CorrelationPolicyRules:t_correlationpolicyrules, EnableSmartScanning:t_enablesmartscanning}
    } else {
        return nil
    }
}

func build_plan_v2_msgcorrelationpolicyrule(d *schema.ResourceData, r []interface{}) *handler.MsgCorrelationPolicyRule {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_threatdetection *handler.MsgThreatDetectionCorrelationRule
        if val, ok := tmp["threatdetection"]; ok {
            t_threatdetection = build_plan_v2_msgthreatdetectioncorrelationrule(d, val.([]interface{}))
        }
        var t_threatnexus *handler.MsgThreatNexusCorrelationRule
        if val, ok := tmp["threatnexus"]; ok {
            t_threatnexus = build_plan_v2_msgthreatnexuscorrelationrule(d, val.([]interface{}))
        }
        var t_threatscan *handler.MsgThreatScanCorrelationRule
        if val, ok := tmp["threatscan"]; ok {
            t_threatscan = build_plan_v2_msgthreatscancorrelationrule(d, val.([]interface{}))
        }
        return &handler.MsgCorrelationPolicyRule{ThreatDetection:t_threatdetection, ThreatNexus:t_threatnexus, ThreatScan:t_threatscan}
    } else {
        return nil
    }
}

func build_plan_v2_msgthreatscancorrelationrule(d *schema.ResourceData, r []interface{}) *handler.MsgThreatScanCorrelationRule {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_maxdays *int
        if val, ok := tmp["maxdays"]; ok {
            t_maxdays = handler.ToIntValue(val, true)
        }
        var t_uselastjobstatus *bool
        if val, ok := tmp["uselastjobstatus"]; ok {
            t_uselastjobstatus = handler.ToBooleanValue(val, true)
        }
        var t_mindays *int
        if val, ok := tmp["mindays"]; ok {
            t_mindays = handler.ToIntValue(val, true)
        }
        return &handler.MsgThreatScanCorrelationRule{MaxDays:t_maxdays, UseLastJobStatus:t_uselastjobstatus, MinDays:t_mindays}
    } else {
        return nil
    }
}

func build_plan_v2_msgthreatnexuscorrelationrule(d *schema.ResourceData, r []interface{}) *handler.MsgThreatNexusCorrelationRule {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_enablethreatnexus *bool
        if val, ok := tmp["enablethreatnexus"]; ok {
            t_enablethreatnexus = handler.ToBooleanValue(val, true)
        }
        var t_mindays *int
        if val, ok := tmp["mindays"]; ok {
            t_mindays = handler.ToIntValue(val, true)
        }
        return &handler.MsgThreatNexusCorrelationRule{EnableThreatNexus:t_enablethreatnexus, MinDays:t_mindays}
    } else {
        return nil
    }
}

func build_plan_v2_msgthreatdetectioncorrelationrule(d *schema.ResourceData, r []interface{}) *handler.MsgThreatDetectionCorrelationRule {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_uselastjobstatus *bool
        if val, ok := tmp["uselastjobstatus"]; ok {
            t_uselastjobstatus = handler.ToBooleanValue(val, true)
        }
        var t_mindays *int
        if val, ok := tmp["mindays"]; ok {
            t_mindays = handler.ToIntValue(val, true)
        }
        return &handler.MsgThreatDetectionCorrelationRule{UseLastJobStatus:t_uselastjobstatus, MinDays:t_mindays}
    } else {
        return nil
    }
}

func build_plan_v2_msgaccessnodesinfo(d *schema.ResourceData, r []interface{}) *handler.MsgAccessNodesInfo {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_accessnodesgroup *handler.MsgIdName
        if val, ok := tmp["accessnodesgroup"]; ok {
            t_accessnodesgroup = build_plan_v2_msgidname(d, val.([]interface{}))
        }
        var t_accessnodes []handler.MsgIdNameSet
        if val, ok := tmp["accessnodes"]; ok {
            t_accessnodes = build_plan_v2_msgidnameset_array(d, val.(*schema.Set).List())
        }
        return &handler.MsgAccessNodesInfo{AccessNodesGroup:t_accessnodesgroup, AccessNodes:t_accessnodes}
    } else {
        return nil
    }
}

func build_plan_v2_msgthreatscanpolicy(d *schema.ResourceData, r []interface{}) *handler.MsgThreatScanPolicy {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_yararules []handler.MsgIdNameSet
        if val, ok := tmp["yararules"]; ok {
            t_yararules = build_plan_v2_msgidnameset_array(d, val.(*schema.Set).List())
        }
        var t_yararule *bool
        if val, ok := tmp["yararule"]; ok {
            t_yararule = handler.ToBooleanValue(val, true)
        }
        var t_iocrules []handler.MsgIdNameSet
        if val, ok := tmp["iocrules"]; ok {
            t_iocrules = build_plan_v2_msgidnameset_array(d, val.(*schema.Set).List())
        }
        var t_filedataanalysis *bool
        if val, ok := tmp["filedataanalysis"]; ok {
            t_filedataanalysis = handler.ToBooleanValue(val, true)
        }
        var t_threatanalysis *bool
        if val, ok := tmp["threatanalysis"]; ok {
            t_threatanalysis = handler.ToBooleanValue(val, true)
        }
        var t_usecommvaulthashfeed *bool
        if val, ok := tmp["usecommvaulthashfeed"]; ok {
            t_usecommvaulthashfeed = handler.ToBooleanValue(val, true)
        }
        var t_usecustomhashfeed *bool
        if val, ok := tmp["usecustomhashfeed"]; ok {
            t_usecustomhashfeed = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgThreatScanPolicy{YaraRules:t_yararules, YaraRule:t_yararule, IocRules:t_iocrules, FileDataAnalysis:t_filedataanalysis, ThreatAnalysis:t_threatanalysis, UseCommvaultHashFeed:t_usecommvaulthashfeed, UseCustomHashFeed:t_usecustomhashfeed}
    } else {
        return nil
    }
}

func build_plan_v2_msgthreatnexuspolicy(d *schema.ResourceData, r []interface{}) *handler.MsgThreatNexusPolicy {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_networkanddatasecurity *handler.MsgNetworkAndDataSecurity
        if val, ok := tmp["networkanddatasecurity"]; ok {
            t_networkanddatasecurity = build_plan_v2_msgnetworkanddatasecurity(d, val.([]interface{}))
        }
        var t_soar *handler.MsgSOAR
        if val, ok := tmp["soar"]; ok {
            t_soar = build_plan_v2_msgsoar(d, val.([]interface{}))
        }
        var t_cspm_dspm *handler.MsgCSPM_DSPM
        if val, ok := tmp["cspm_dspm"]; ok {
            t_cspm_dspm = build_plan_v2_msgcspm_dspm(d, val.([]interface{}))
        }
        return &handler.MsgThreatNexusPolicy{NetworkAndDataSecurity:t_networkanddatasecurity, Soar:t_soar, Cspm_dspm:t_cspm_dspm}
    } else {
        return nil
    }
}

func build_plan_v2_msgcspm_dspm(d *schema.ResourceData, r []interface{}) *handler.MsgCSPM_DSPM {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_acante *bool
        if val, ok := tmp["acante"]; ok {
            t_acante = handler.ToBooleanValue(val, true)
        }
        var t_dasera *bool
        if val, ok := tmp["dasera"]; ok {
            t_dasera = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgCSPM_DSPM{Acante:t_acante, Dasera:t_dasera}
    } else {
        return nil
    }
}

func build_plan_v2_msgsoar(d *schema.ResourceData, r []interface{}) *handler.MsgSOAR {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_paloalto_xsoar *bool
        if val, ok := tmp["paloalto_xsoar"]; ok {
            t_paloalto_xsoar = handler.ToBooleanValue(val, true)
        }
        var t_mssentinel *bool
        if val, ok := tmp["mssentinel"]; ok {
            t_mssentinel = handler.ToBooleanValue(val, true)
        }
        var t_splunk *bool
        if val, ok := tmp["splunk"]; ok {
            t_splunk = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgSOAR{PaloAlto_XSOAR:t_paloalto_xsoar, MsSentinel:t_mssentinel, Splunk:t_splunk}
    } else {
        return nil
    }
}

func build_plan_v2_msgnetworkanddatasecurity(d *schema.ResourceData, r []interface{}) *handler.MsgNetworkAndDataSecurity {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_netskope *bool
        if val, ok := tmp["netskope"]; ok {
            t_netskope = handler.ToBooleanValue(val, true)
        }
        var t_darktrace *bool
        if val, ok := tmp["darktrace"]; ok {
            t_darktrace = handler.ToBooleanValue(val, true)
        }
        var t_crowdstrike *bool
        if val, ok := tmp["crowdstrike"]; ok {
            t_crowdstrike = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgNetworkAndDataSecurity{Netskope:t_netskope, Darktrace:t_darktrace, Crowdstrike:t_crowdstrike}
    } else {
        return nil
    }
}

func build_plan_v2_msgthreatdetectionpolicy(d *schema.ResourceData, r []interface{}) *handler.MsgThreatDetectionPolicy {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_backupsize *bool
        if val, ok := tmp["backupsize"]; ok {
            t_backupsize = handler.ToBooleanValue(val, true)
        }
        var t_canaryfile *bool
        if val, ok := tmp["canaryfile"]; ok {
            t_canaryfile = handler.ToBooleanValue(val, true)
        }
        var t_fileextension *bool
        if val, ok := tmp["fileextension"]; ok {
            t_fileextension = handler.ToBooleanValue(val, true)
        }
        var t_fileactivity *bool
        if val, ok := tmp["fileactivity"]; ok {
            t_fileactivity = handler.ToBooleanValue(val, true)
        }
        var t_filetype *bool
        if val, ok := tmp["filetype"]; ok {
            t_filetype = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgThreatDetectionPolicy{BackupSize:t_backupsize, CanaryFile:t_canaryfile, FileExtension:t_fileextension, FileActivity:t_fileactivity, FileType:t_filetype}
    } else {
        return nil
    }
}

func build_plan_v2_msgdcplancipolicy(d *schema.ResourceData, r []interface{}) *handler.MsgDCPlanCIpolicy {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_filefilters *handler.MsgDCPlanCIFileFilters
        if val, ok := tmp["filefilters"]; ok {
            t_filefilters = build_plan_v2_msgdcplancifilefilters(d, val.([]interface{}))
        }
        var t_copyprecedence *int
        if val, ok := tmp["copyprecedence"]; ok {
            t_copyprecedence = handler.ToIntValue(val, true)
        }
        var t_searchtype *string
        if val, ok := tmp["searchtype"]; ok {
            t_searchtype = handler.ToStringValue(val, true)
        }
        var t_contentlanguage *string
        if val, ok := tmp["contentlanguage"]; ok {
            t_contentlanguage = handler.ToStringValue(val, true)
        }
        var t_exactsearch *bool
        if val, ok := tmp["exactsearch"]; ok {
            t_exactsearch = handler.ToBooleanValue(val, true)
        }
        var t_backupcopy []handler.MsgDCBackupCopySet
        if val, ok := tmp["backupcopy"]; ok {
            t_backupcopy = build_plan_v2_msgdcbackupcopyset_array(d, val.(*schema.Set).List())
        }
        var t_extracttextfromimage *bool
        if val, ok := tmp["extracttextfromimage"]; ok {
            t_extracttextfromimage = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgDCPlanCIpolicy{FileFilters:t_filefilters, CopyPrecedence:t_copyprecedence, SearchType:t_searchtype, ContentLanguage:t_contentlanguage, ExactSearch:t_exactsearch, BackupCopy:t_backupcopy, ExtractTextFromImage:t_extracttextfromimage}
    } else {
        return nil
    }
}

func build_plan_v2_msgdcbackupcopyset_array(d *schema.ResourceData, r []interface{}) []handler.MsgDCBackupCopySet {
    if r != nil {
        tmp := make([]handler.MsgDCBackupCopySet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_copyid *int
            if val, ok := raw_a["copyid"]; ok {
                t_copyid = handler.ToIntValue(val, true)
            }
            var t_storagepoolid *int
            if val, ok := raw_a["storagepoolid"]; ok {
                t_storagepoolid = handler.ToIntValue(val, true)
            }
            tmp[a] = handler.MsgDCBackupCopySet{CopyId:t_copyid, StoragePoolId:t_storagepoolid}
        }
        return tmp
    } else {
        return nil
    }
}

func build_plan_v2_msgdcplancifilefilters(d *schema.ResourceData, r []interface{}) *handler.MsgDCPlanCIFileFilters {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_includedoctypes *string
        if val, ok := tmp["includedoctypes"]; ok {
            t_includedoctypes = handler.ToStringValue(val, true)
        }
        var t_excludepaths []string
        if val, ok := tmp["excludepaths"]; ok {
            t_excludepaths = handler.ToStringArray(val.(*schema.Set).List())
        }
        var t_mindocsize *int
        if val, ok := tmp["mindocsize"]; ok {
            t_mindocsize = handler.ToIntValue(val, true)
        }
        var t_maxdocsize *int
        if val, ok := tmp["maxdocsize"]; ok {
            t_maxdocsize = handler.ToIntValue(val, true)
        }
        return &handler.MsgDCPlanCIFileFilters{IncludeDocTypes:t_includedoctypes, ExcludePaths:t_excludepaths, MinDocSize:t_mindocsize, MaxDocSize:t_maxdocsize}
    } else {
        return nil
    }
}

func build_plan_v2_msgdcplanjobschedule(d *schema.ResourceData, r []interface{}) *handler.MsgDCPlanJobSchedule {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_policyid *int
        if val, ok := tmp["policyid"]; ok {
            t_policyid = handler.ToIntValue(val, true)
        }
        var t_name *string
        if val, ok := tmp["name"]; ok {
            t_name = handler.ToStringValue(val, true)
        }
        var t_pattern *handler.MsgContentIndexingFrequencyPattern
        if val, ok := tmp["pattern"]; ok {
            t_pattern = build_plan_v2_msgcontentindexingfrequencypattern(d, val.([]interface{}))
        }
        var t_scheduleid *int
        if val, ok := tmp["scheduleid"]; ok {
            t_scheduleid = handler.ToIntValue(val, true)
        }
        return &handler.MsgDCPlanJobSchedule{PolicyId:t_policyid, Name:t_name, Pattern:t_pattern, ScheduleId:t_scheduleid}
    } else {
        return nil
    }
}

func build_plan_v2_msgcontentindexingfrequencypattern(d *schema.ResourceData, r []interface{}) *handler.MsgContentIndexingFrequencyPattern {
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
        var t_timezone *handler.MsgIdName
        if val, ok := tmp["timezone"]; ok {
            t_timezone = build_plan_v2_msgidname(d, val.([]interface{}))
        }
        var t_startdate *int
        if val, ok := tmp["startdate"]; ok {
            t_startdate = handler.ToIntValue(val, true)
        }
        var t_exceptions []handler.MsgScheduleRunExceptionSet
        if val, ok := tmp["exceptions"]; ok {
            t_exceptions = build_plan_v2_msgschedulerunexceptionset_array(d, val.(*schema.Set).List())
        }
        return &handler.MsgContentIndexingFrequencyPattern{WeeklyDays:t_weeklydays, MonthOfYear:t_monthofyear, DayOfWeek:t_dayofweek, DayOfMonth:t_dayofmonth, ScheduleFrequencyType:t_schedulefrequencytype, WeekOfMonth:t_weekofmonth, StartTime:t_starttime, Frequency:t_frequency, TimeZone:t_timezone, StartDate:t_startdate, Exceptions:t_exceptions}
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
            tmp[a] = handler.MsgScheduleRunExceptionSet{OnWeekOfTheMonth:t_onweekofthemonth, OnDates:t_ondates, OnDayOfTheWeek:t_ondayoftheweek}
        }
        return tmp
    } else {
        return nil
    }
}
