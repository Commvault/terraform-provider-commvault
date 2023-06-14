package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePlan_BackupDestination() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreatePlan_BackupDestination,
        Read:   resourceReadPlan_BackupDestination,
        Update: resourceUpdatePlan_BackupDestination,
        Delete: resourceDeletePlan_BackupDestination,

        Schema: map[string]*schema.Schema{
            "ismirrorcopy": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Is this a mirror copy? Only considered when isSnapCopy is true.",
            },
            "retentionperioddays": {
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: "Retention period in days. -1 can be specified for infinite retention. If this and snapRecoveryPoints both are not specified, this takes  precedence.",
            },
            "backupstocopy": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "All_JOBS means SYNCHRONOUS copy type, others are applicable for SELECTIVE copy Type only. [All_JOBS, ALL_FULLS, HOURLY_FULLS, DAILY_FULLS, WEEKLY_FULLS, MONTHLY_FULLS, QUARTERLY_FULLS, HALF_YEARLY_FULLS, YEARLY_FULLS, ADVANCED]",
            },
            "extendedretentionrules": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "If you want to update, specify the whole object. Extended retention rules should be bigger than retention period.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "thirdextendedretentionrule": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "isinfiniteretention": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "If this is set as true, no need to specify retentionPeriodDays.",
                                    },
                                    "retentionperioddays": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "If this is set, no need to specify isInfiniteRetention as false.",
                                    },
                                    "type": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "All_JOBS means SYNCHRONOUS copy type, others are applicable for SELECTIVE copy Type only. [All_JOBS, ALL_FULLS, HOURLY_FULLS, DAILY_FULLS, WEEKLY_FULLS, MONTHLY_FULLS, QUARTERLY_FULLS, HALF_YEARLY_FULLS, YEARLY_FULLS, ADVANCED]",
                                    },
                                },
                            },
                        },
                        "firstextendedretentionrule": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "isinfiniteretention": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "If this is set as true, no need to specify retentionPeriodDays.",
                                    },
                                    "retentionperioddays": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "If this is set, no need to specify isInfiniteRetention as false.",
                                    },
                                    "type": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "All_JOBS means SYNCHRONOUS copy type, others are applicable for SELECTIVE copy Type only. [All_JOBS, ALL_FULLS, HOURLY_FULLS, DAILY_FULLS, WEEKLY_FULLS, MONTHLY_FULLS, QUARTERLY_FULLS, HALF_YEARLY_FULLS, YEARLY_FULLS, ADVANCED]",
                                    },
                                },
                            },
                        },
                        "secondextendedretentionrule": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "isinfiniteretention": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "If this is set as true, no need to specify retentionPeriodDays.",
                                    },
                                    "retentionperioddays": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "If this is set, no need to specify isInfiniteRetention as false.",
                                    },
                                    "type": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
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
                Computed:    true,
                Description: "Which type of retention rule should be used for the given backup destination [RETENTION_PERIOD, SNAP_RECOVERY_POINTS]",
            },
            "snaprecoverypoints": {
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: "Number of snap recovery points for snap copy for retention. Can be specified instead of retention period in Days for snap copy.",
            },
            "sourcecopy": {
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
            "fullbackuptypestocopy": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Which type of backup type should be copied for the given backup destination when backup type is not all jobs. Default is LAST while adding new backup destination. [FIRST, LAST]",
            },
            "useextendedretentionrules": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Use extended retention rules",
            },
            "backupstarttime": {
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: "Backup start time in seconds. The time is provided in unix time format.",
            },
            "overrideretentionsettings": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Tells if this copy should use storage pool retention period days or the retention defined for this copy. Set as true to use retention defined on this copy.",
            },
            "optimizeforinstantclone": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Flag to specify if primary storage is copy data management enabled.",
            },
            "netappcloudtarget": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
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
                Computed:    true,
                Description: "Is this a snap copy? If isMirrorCopy is not set, then default is Vault/Replica.",
            },
            "name": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Name of backup destination",
            },
            "storagetype": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "[ALL, DISK, CLOUD, HYPERSCALE, TAPE]",
            },
            "region": {
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
            "storagepool": {
                Type:        schema.TypeList,
                Required:    true,
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
            "enabledataaging": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Tells if this copy has data aging enabled",
            },
        },
    }
}

func resourceCreatePlan_BackupDestination(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/Plan/BackupDestinations
    var response_id = strconv.Itoa(0)
    t_backupdestinations := build_plan_backupdestination_msgcreatebackupdestinationset_array(d)
    var req = handler.MsgCreateBackupDestinationWithoutPlanInfoRequest{BackupDestinations:t_backupdestinations}
    resp, err := handler.CvCreateBackupDestinationWithoutPlanInfo(req)
    if err != nil {
        return fmt.Errorf("operation [CreateBackupDestinationWithoutPlanInfo] failed, Error %s", err)
    }
    if resp.PlanBackupDestination != nil && len(resp.PlanBackupDestination) > 0 {
        if resp.PlanBackupDestination[0].Id != nil {
            response_id = strconv.Itoa(*resp.PlanBackupDestination[0].Id)
        }
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateBackupDestinationWithoutPlanInfo] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdatePlan_BackupDestination(d, m)
    }
}

func resourceReadPlan_BackupDestination(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/Plan/BackupDestination/{BackupDestinationId}
    resp, err := handler.CvGetBackupDestinationDetailsWithoutPlanInfo(d.Id())
    if err != nil {
        return fmt.Errorf("operation [GetBackupDestinationDetailsWithoutPlanInfo] failed, Error %s", err)
    }
    if resp.IsMirrorCopy != nil {
        d.Set("ismirrorcopy", strconv.FormatBool(*resp.IsMirrorCopy))
    }
    if resp.RetentionPeriodDays != nil {
        d.Set("retentionperioddays", resp.RetentionPeriodDays)
    }
    if resp.BackupsToCopy != nil {
        d.Set("backupstocopy", resp.BackupsToCopy)
    }
    if rtn, ok := serialize_plan_backupdestination_msgextendedretentionrules(d, resp.ExtendedRetentionRules); ok {
        d.Set("extendedretentionrules", rtn)
    } else {
        d.Set("extendedretentionrules", make([]map[string]interface{}, 0))
    }
    if resp.RetentionRuleType != nil {
        d.Set("retentionruletype", resp.RetentionRuleType)
    }
    if resp.SnapRecoveryPoints != nil {
        d.Set("snaprecoverypoints", resp.SnapRecoveryPoints)
    }
    if rtn, ok := serialize_plan_backupdestination_msgidname(d, resp.SourceCopy); ok {
        d.Set("sourcecopy", rtn)
    } else {
        d.Set("sourcecopy", make([]map[string]interface{}, 0))
    }
    if resp.FullBackupTypesToCopy != nil {
        d.Set("fullbackuptypestocopy", resp.FullBackupTypesToCopy)
    }
    if resp.UseExtendedRetentionRules != nil {
        d.Set("useextendedretentionrules", strconv.FormatBool(*resp.UseExtendedRetentionRules))
    }
    if resp.BackupStartTime != nil {
        d.Set("backupstarttime", resp.BackupStartTime)
    }
    if resp.OverrideRetentionSettings != nil {
        d.Set("overrideretentionsettings", strconv.FormatBool(*resp.OverrideRetentionSettings))
    }
    if resp.NetAppCloudTarget != nil {
        d.Set("netappcloudtarget", strconv.FormatBool(*resp.NetAppCloudTarget))
    }
    if rtn, ok := serialize_plan_backupdestination_msgsnapshotcopymappingset_array(d, resp.Mappings); ok {
        d.Set("mappings", rtn)
    } else {
        d.Set("mappings", make([]map[string]interface{}, 0))
    }
    if resp.IsSnapCopy != nil {
        d.Set("issnapcopy", strconv.FormatBool(*resp.IsSnapCopy))
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    if resp.StorageType != nil {
        d.Set("storagetype", resp.StorageType)
    }
    if resp.EnableDataAging != nil {
        d.Set("enabledataaging", strconv.FormatBool(*resp.EnableDataAging))
    }
    if rtn, ok := serialize_plan_backupdestination_msgidnamedisplayname(d, resp.Region); ok {
        d.Set("region", rtn)
    } else {
        d.Set("region", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_plan_backupdestination_msgstoragepool(d, resp.StoragePool); ok {
        d.Set("storagepool", rtn)
    } else {
        d.Set("storagepool", make([]map[string]interface{}, 0))
    }
    return nil
}

func resourceUpdatePlan_BackupDestination(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Plan/BackupDestination/{BackupDestinationId}
    var t_retentionperioddays *int
    if d.HasChange("retentionperioddays") {
        val := d.Get("retentionperioddays")
        t_retentionperioddays = handler.ToIntValue(val, false)
    }
    var t_backupstocopy *string
    if d.HasChange("backupstocopy") {
        val := d.Get("backupstocopy")
        t_backupstocopy = handler.ToStringValue(val, false)
    }
    var t_extendedretentionrules *handler.MsgExtendedRetentionRules
    if d.HasChange("extendedretentionrules") {
        val := d.Get("extendedretentionrules")
        t_extendedretentionrules = build_plan_backupdestination_msgextendedretentionrules(d, val.([]interface{}))
    }
    var t_retentionruletype *string
    if d.HasChange("retentionruletype") {
        val := d.Get("retentionruletype")
        t_retentionruletype = handler.ToStringValue(val, false)
    }
    var t_snaprecoverypoints *int
    if d.HasChange("snaprecoverypoints") {
        val := d.Get("snaprecoverypoints")
        t_snaprecoverypoints = handler.ToIntValue(val, false)
    }
    var t_sourcecopy *handler.MsgIdName
    if d.HasChange("sourcecopy") {
        val := d.Get("sourcecopy")
        t_sourcecopy = build_plan_backupdestination_msgidname(d, val.([]interface{}))
    }
    var t_useextendedretentionrules *bool
    if d.HasChange("useextendedretentionrules") {
        val := d.Get("useextendedretentionrules")
        t_useextendedretentionrules = handler.ToBooleanValue(val, false)
    }
    var t_backupstarttime *int
    if d.HasChange("backupstarttime") {
        val := d.Get("backupstarttime")
        t_backupstarttime = handler.ToIntValue(val, false)
    }
    var t_overrideretentionsettings *bool
    if d.HasChange("overrideretentionsettings") {
        val := d.Get("overrideretentionsettings")
        t_overrideretentionsettings = handler.ToBooleanValue(val, false)
    }
    var t_mappings []handler.MsgSnapshotCopyMappingSet
    if d.HasChange("mappings") {
        val := d.Get("mappings")
        t_mappings = build_plan_backupdestination_msgsnapshotcopymappingset_array(d, val.(*schema.Set).List())
    }
    var t_name *string
    if d.HasChange("name") {
        val := d.Get("name")
        t_name = handler.ToStringValue(val, false)
    }
    var t_enabledataaging *bool
    if d.HasChange("enabledataaging") {
        val := d.Get("enabledataaging")
        t_enabledataaging = handler.ToBooleanValue(val, false)
    }
    var t_region *handler.MsgIdName
    if d.HasChange("region") {
        val := d.Get("region")
        t_region = build_plan_backupdestination_msgidname(d, val.([]interface{}))
    }
    var req = handler.MsgModifyBackupDestinationWithoutPlanInfoRequest{RetentionPeriodDays:t_retentionperioddays, BackupsToCopy:t_backupstocopy, ExtendedRetentionRules:t_extendedretentionrules, RetentionRuleType:t_retentionruletype, SnapRecoveryPoints:t_snaprecoverypoints, SourceCopy:t_sourcecopy, UseExtendedRetentionRules:t_useextendedretentionrules, BackupStartTime:t_backupstarttime, OverrideRetentionSettings:t_overrideretentionsettings, Mappings:t_mappings, Name:t_name, EnableDataAging:t_enabledataaging, Region:t_region}
    _, err := handler.CvModifyBackupDestinationWithoutPlanInfo(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [ModifyBackupDestinationWithoutPlanInfo] failed, Error %s", err)
    }
    return resourceReadPlan_BackupDestination(d, m)
}

func resourceCreateUpdatePlan_BackupDestination(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Plan/BackupDestination/{BackupDestinationId}
    var execUpdate bool = false
    var t_enabledataaging *bool
    if val, ok := d.GetOk("enabledataaging"); ok {
        t_enabledataaging = handler.ToBooleanValue(val, false)
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgModifyBackupDestinationWithoutPlanInfoRequest{EnableDataAging:t_enabledataaging}
        _, err := handler.CvModifyBackupDestinationWithoutPlanInfo(req, d.Id())
        if err != nil {
            return fmt.Errorf("operation [ModifyBackupDestinationWithoutPlanInfo] failed, Error %s", err)
        }
    }
    return resourceReadPlan_BackupDestination(d, m)
}

func resourceDeletePlan_BackupDestination(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/Plan/BackupDestination/{BackupDestinationId}
    _, err := handler.CvDeleteBackupDestinationWithoutPlanInfo(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteBackupDestinationWithoutPlanInfo] failed, Error %s", err)
    }
    return nil
}

func build_plan_backupdestination_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_plan_backupdestination_msgsnapshotcopymappingset_array(d *schema.ResourceData, r []interface{}) []handler.MsgSnapshotCopyMappingSet {
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
                t_targetvendor = build_plan_backupdestination_msgidname(d, val.([]interface{}))
            }
            var t_source *handler.MsgIdName
            if val, ok := raw_a["source"]; ok {
                t_source = build_plan_backupdestination_msgidname(d, val.([]interface{}))
            }
            var t_sourcevendor *handler.MsgIdName
            if val, ok := raw_a["sourcevendor"]; ok {
                t_sourcevendor = build_plan_backupdestination_msgidname(d, val.([]interface{}))
            }
            var t_target *handler.MsgIdName
            if val, ok := raw_a["target"]; ok {
                t_target = build_plan_backupdestination_msgidname(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgSnapshotCopyMappingSet{Vendor:t_vendor, TargetVendor:t_targetvendor, Source:t_source, SourceVendor:t_sourcevendor, Target:t_target}
        }
        return tmp
    } else {
        return nil
    }
}

func build_plan_backupdestination_msgextendedretentionrules(d *schema.ResourceData, r []interface{}) *handler.MsgExtendedRetentionRules {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_thirdextendedretentionrule *handler.MsgPlanRetentionRule
        if val, ok := tmp["thirdextendedretentionrule"]; ok {
            t_thirdextendedretentionrule = build_plan_backupdestination_msgplanretentionrule(d, val.([]interface{}))
        }
        var t_firstextendedretentionrule *handler.MsgPlanRetentionRule
        if val, ok := tmp["firstextendedretentionrule"]; ok {
            t_firstextendedretentionrule = build_plan_backupdestination_msgplanretentionrule(d, val.([]interface{}))
        }
        var t_secondextendedretentionrule *handler.MsgPlanRetentionRule
        if val, ok := tmp["secondextendedretentionrule"]; ok {
            t_secondextendedretentionrule = build_plan_backupdestination_msgplanretentionrule(d, val.([]interface{}))
        }
        return &handler.MsgExtendedRetentionRules{ThirdExtendedRetentionRule:t_thirdextendedretentionrule, FirstExtendedRetentionRule:t_firstextendedretentionrule, SecondExtendedRetentionRule:t_secondextendedretentionrule}
    } else {
        return nil
    }
}

func build_plan_backupdestination_msgplanretentionrule(d *schema.ResourceData, r []interface{}) *handler.MsgPlanRetentionRule {
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
        return &handler.MsgPlanRetentionRule{IsInfiniteRetention:t_isinfiniteretention, RetentionPeriodDays:t_retentionperioddays, Type:t_type}
    } else {
        return nil
    }
}

func build_plan_backupdestination_msgcreatebackupdestinationset_array(d *schema.ResourceData) []handler.MsgCreateBackupDestinationSet {
    tmp := make([]handler.MsgCreateBackupDestinationSet, 1)
    var t_ismirrorcopy *bool
    if val, ok := d.GetOk("ismirrorcopy"); ok {
        t_ismirrorcopy = handler.ToBooleanValue(val, false)
    }
    var t_retentionperioddays *int
    if val, ok := d.GetOk("retentionperioddays"); ok {
        t_retentionperioddays = handler.ToIntValue(val, false)
    }
    var t_backupstocopy *string
    if val, ok := d.GetOk("backupstocopy"); ok {
        t_backupstocopy = handler.ToStringValue(val, false)
    }
    var t_extendedretentionrules *handler.MsgExtendedRetentionRules
    if val, ok := d.GetOk("extendedretentionrules"); ok {
        t_extendedretentionrules = build_plan_backupdestination_msgextendedretentionrules(d, val.([]interface{}))
    }
    var t_retentionruletype *string
    if val, ok := d.GetOk("retentionruletype"); ok {
        t_retentionruletype = handler.ToStringValue(val, false)
    }
    var t_snaprecoverypoints *int
    if val, ok := d.GetOk("snaprecoverypoints"); ok {
        t_snaprecoverypoints = handler.ToIntValue(val, false)
    }
    var t_sourcecopy *handler.MsgIdName
    if val, ok := d.GetOk("sourcecopy"); ok {
        t_sourcecopy = build_plan_backupdestination_msgidname(d, val.([]interface{}))
    }
    var t_fullbackuptypestocopy *string
    if val, ok := d.GetOk("fullbackuptypestocopy"); ok {
        t_fullbackuptypestocopy = handler.ToStringValue(val, false)
    }
    var t_useextendedretentionrules *bool
    if val, ok := d.GetOk("useextendedretentionrules"); ok {
        t_useextendedretentionrules = handler.ToBooleanValue(val, false)
    }
    var t_backupstarttime *int
    if val, ok := d.GetOk("backupstarttime"); ok {
        t_backupstarttime = handler.ToIntValue(val, false)
    }
    var t_overrideretentionsettings *bool
    if val, ok := d.GetOk("overrideretentionsettings"); ok {
        t_overrideretentionsettings = handler.ToBooleanValue(val, false)
    }
    var t_optimizeforinstantclone *bool
    if val, ok := d.GetOk("optimizeforinstantclone"); ok {
        t_optimizeforinstantclone = handler.ToBooleanValue(val, false)
    }
    var t_netappcloudtarget *bool
    if val, ok := d.GetOk("netappcloudtarget"); ok {
        t_netappcloudtarget = handler.ToBooleanValue(val, false)
    }
    var t_mappings []handler.MsgSnapshotCopyMappingSet
    if val, ok := d.GetOk("mappings"); ok {
        t_mappings = build_plan_backupdestination_msgsnapshotcopymappingset_array(d, val.(*schema.Set).List())
    }
    var t_issnapcopy *bool
    if val, ok := d.GetOk("issnapcopy"); ok {
        t_issnapcopy = handler.ToBooleanValue(val, false)
    }
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_storagetype *string
    if val, ok := d.GetOk("storagetype"); ok {
        t_storagetype = handler.ToStringValue(val, false)
    }
    var t_region *handler.MsgIdName
    if val, ok := d.GetOk("region"); ok {
        t_region = build_plan_backupdestination_msgidname(d, val.([]interface{}))
    }
    var t_storagepool *handler.MsgIdName
    if val, ok := d.GetOk("storagepool"); ok {
        t_storagepool = build_plan_backupdestination_msgidname(d, val.([]interface{}))
    }
    tmp[0] = handler.MsgCreateBackupDestinationSet{IsMirrorCopy:t_ismirrorcopy, RetentionPeriodDays:t_retentionperioddays, BackupsToCopy:t_backupstocopy, ExtendedRetentionRules:t_extendedretentionrules, RetentionRuleType:t_retentionruletype, SnapRecoveryPoints:t_snaprecoverypoints, SourceCopy:t_sourcecopy, FullBackupTypesToCopy:t_fullbackuptypestocopy, UseExtendedRetentionRules:t_useextendedretentionrules, BackupStartTime:t_backupstarttime, OverrideRetentionSettings:t_overrideretentionsettings, OptimizeForInstantClone:t_optimizeforinstantclone, NetAppCloudTarget:t_netappcloudtarget, Mappings:t_mappings, IsSnapCopy:t_issnapcopy, Name:t_name, StorageType:t_storagetype, Region:t_region, StoragePool:t_storagepool}
    return tmp
}

func serialize_plan_backupdestination_msgstoragepool(d *schema.ResourceData, data *handler.MsgStoragePool) ([]map[string]interface{}, bool) {
    //MsgIdName
    //MsgStoragePool
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

func serialize_plan_backupdestination_msgidnamedisplayname(d *schema.ResourceData, data *handler.MsgIdNameDisplayName) ([]map[string]interface{}, bool) {
    //MsgIdName
    //MsgIdNameDisplayName
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

func serialize_plan_backupdestination_msgsnapshotcopymappingset_array(d *schema.ResourceData, data []handler.MsgSnapshotCopyMappingSet) ([]map[string]interface{}, bool) {
    //MsgSnapshotCopyMappingSet
    //MsgSnapshotCopyMappingSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].Vendor != nil {
            tmp["vendor"] = data[i].Vendor
            added = true
        }
        if rtn, ok := serialize_plan_backupdestination_msgidname(d, data[i].TargetVendor); ok {
            tmp["targetvendor"] = rtn
            added = true
        }
        if rtn, ok := serialize_plan_backupdestination_msgidname(d, data[i].Source); ok {
            tmp["source"] = rtn
            added = true
        }
        if rtn, ok := serialize_plan_backupdestination_msgidname(d, data[i].SourceVendor); ok {
            tmp["sourcevendor"] = rtn
            added = true
        }
        if rtn, ok := serialize_plan_backupdestination_msgidname(d, data[i].Target); ok {
            tmp["target"] = rtn
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_plan_backupdestination_msgidname(d *schema.ResourceData, data *handler.MsgIdName) ([]map[string]interface{}, bool) {
    //MsgSnapshotCopyMappingSet -> MsgIdName
    //MsgSnapshotCopyMappingSet -> MsgIdName
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

func serialize_plan_backupdestination_msgextendedretentionrules(d *schema.ResourceData, data *handler.MsgExtendedRetentionRules) ([]map[string]interface{}, bool) {
    //MsgExtendedRetentionRules
    //MsgExtendedRetentionRules
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_plan_backupdestination_msgplanretentionrule(d, data.ThirdExtendedRetentionRule); ok {
        val[0]["thirdextendedretentionrule"] = rtn
        added = true
    }
    if rtn, ok := serialize_plan_backupdestination_msgplanretentionrule(d, data.FirstExtendedRetentionRule); ok {
        val[0]["firstextendedretentionrule"] = rtn
        added = true
    }
    if rtn, ok := serialize_plan_backupdestination_msgplanretentionrule(d, data.SecondExtendedRetentionRule); ok {
        val[0]["secondextendedretentionrule"] = rtn
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_plan_backupdestination_msgplanretentionrule(d *schema.ResourceData, data *handler.MsgPlanRetentionRule) ([]map[string]interface{}, bool) {
    //MsgExtendedRetentionRules -> MsgPlanRetentionRule
    //MsgExtendedRetentionRules -> MsgPlanRetentionRule
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.IsInfiniteRetention != nil {
        val[0]["isinfiniteretention"] = strconv.FormatBool(*data.IsInfiniteRetention)
        added = true
    }
    if data.RetentionPeriodDays != nil {
        val[0]["retentionperioddays"] = data.RetentionPeriodDays
        added = true
    }
    if data.Type != nil {
        val[0]["type"] = data.Type
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}
