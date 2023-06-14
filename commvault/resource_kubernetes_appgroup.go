package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceKubernetes_Appgroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateKubernetes_Appgroup,
        Read:   resourceReadKubernetes_Appgroup,
        Update: resourceUpdateKubernetes_Appgroup,
        Delete: resourceDeleteKubernetes_Appgroup,

        Schema: map[string]*schema.Schema{
            "name": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Specify new name to rename an Application Group",
            },
            "filters": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "skipstatelessapps": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Specify whether to skip backup of stateless applications",
                        },
                        "labelselectors": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "List of label selectors to be added as content",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "selectorlevel": {
                                        Type:        schema.TypeString,
                                        Required:    true,
                                        Description: "Selector level of the label selector [Application, Volumes, Namespace]",
                                    },
                                    "selectorvalue": {
                                        Type:        schema.TypeString,
                                        Required:    true,
                                        Description: "Value of the label selector in key=value format",
                                    },
                                },
                            },
                        },
                        "applications": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "List of applications to be added as content",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "guid": {
                                        Type:        schema.TypeString,
                                        Required:    true,
                                        Description: "GUID value of the Kubernetes Application to be associated as content",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "Name of the application",
                                    },
                                    "type": {
                                        Type:        schema.TypeString,
                                        Required:    true,
                                        Description: "Type of the Kubernetes application [NAMESPACE, APPLICATION, PVC, LABELS]",
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "plan": {
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
            "content": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Item describing the content for Application Group",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "labelselectors": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "List of label selectors to be added as content",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "selectorlevel": {
                                        Type:        schema.TypeString,
                                        Required:    true,
                                        Description: "Selector level of the label selector [Application, Volumes, Namespace]",
                                    },
                                    "selectorvalue": {
                                        Type:        schema.TypeString,
                                        Required:    true,
                                        Description: "Value of the label selector in key=value format",
                                    },
                                },
                            },
                        },
                        "applications": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "List of applications to be added as content",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "guid": {
                                        Type:        schema.TypeString,
                                        Required:    true,
                                        Description: "GUID value of the Kubernetes Application to be associated as content",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "Name of the application",
                                    },
                                    "type": {
                                        Type:        schema.TypeString,
                                        Required:    true,
                                        Description: "Type of the Kubernetes application [NAMESPACE, APPLICATION, PVC, LABELS]",
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "cluster": {
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
            "activitycontrol": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enablebackup": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
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
            "options": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "backupstreams": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Define number of parallel data readers",
                        },
                        "cvnamespacescheduling": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Define setting to enable scheduling worker Pods to CV Namespace for CSI-Snapshot enabled backups",
                        },
                        "workerresources": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "memoryrequests": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Define requests.memory to set on the worker Pod",
                                    },
                                    "memorylimits": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Define limits.memory to set on the worker Pod",
                                    },
                                    "cpulimits": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Define limits.cpu to set on the worker Pod",
                                    },
                                    "cpurequests": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Define requests.cpu to set on the worker Pod",
                                    },
                                },
                            },
                        },
                        "snapfallbacktolivevolumebackup": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Define setting to enable fallback to live volume backup in case of snap failure",
                        },
                        "jobstarttime": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Define the backup job start time in epochs",
                        },
                    },
                },
            },
            "tags": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "",
                        },
                        "value": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "",
                        },
                    },
                },
            },
        },
    }
}

func resourceCreateKubernetes_Appgroup(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V5/Kubernetes/ApplicationGroup
    var response_id = strconv.Itoa(0)
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_filters *handler.MsgKubernetesApplicationGroupFilters
    if val, ok := d.GetOk("filters"); ok {
        t_filters = build_kubernetes_appgroup_msgkubernetesapplicationgroupfilters(d, val.([]interface{}))
    }
    var t_plan *handler.MsgIdName
    if val, ok := d.GetOk("plan"); ok {
        t_plan = build_kubernetes_appgroup_msgidname(d, val.([]interface{}))
    }
    var t_content *handler.MsgKubernetesApplicationGroupContent
    if val, ok := d.GetOk("content"); ok {
        t_content = build_kubernetes_appgroup_msgkubernetesapplicationgroupcontent(d, val.([]interface{}))
    }
    var t_cluster *handler.MsgIdName
    if val, ok := d.GetOk("cluster"); ok {
        t_cluster = build_kubernetes_appgroup_msgidname(d, val.([]interface{}))
    }
    var req = handler.MsgCreateKubernetesApplicationGroupRequest{Name:t_name, Filters:t_filters, Plan:t_plan, Content:t_content, Cluster:t_cluster}
    resp, err := handler.CvCreateKubernetesApplicationGroup(req)
    if err != nil {
        return fmt.Errorf("operation [CreateKubernetesApplicationGroup] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateKubernetesApplicationGroup] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateKubernetes_Appgroup(d, m)
    }
}

func resourceReadKubernetes_Appgroup(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V5/Kubernetes/ApplicationGroup/{applicationGroupId}
    resp, err := handler.CvGetApplicationGroupDetails(d.Id())
    if err != nil {
        return fmt.Errorf("operation [GetApplicationGroupDetails] failed, Error %s", err)
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    if rtn, ok := serialize_kubernetes_appgroup_msgapplicationgroupactivitycontrol(d, resp.ActivityControl); ok {
        d.Set("activitycontrol", rtn)
    } else {
        d.Set("activitycontrol", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_kubernetes_appgroup_msgidname(d, resp.Timezone); ok {
        d.Set("timezone", rtn)
    } else {
        d.Set("timezone", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_kubernetes_appgroup_msgapplicationgroupgetoptions(d, resp.Options); ok {
        d.Set("options", rtn)
    } else {
        d.Set("options", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_kubernetes_appgroup_msgkubernetesapplicationgroupfilteritem(d, resp.Filters); ok {
        d.Set("filters", rtn)
    } else {
        d.Set("filters", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_kubernetes_appgroup_msgidname(d, resp.Plan); ok {
        d.Set("plan", rtn)
    } else {
        d.Set("plan", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_kubernetes_appgroup_msgkubernetesapplicationgroupcontentitem(d, resp.Content); ok {
        d.Set("content", rtn)
    } else {
        d.Set("content", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_kubernetes_appgroup_msgidnamevalueset_array(d, resp.Tags); ok {
        d.Set("tags", rtn)
    } else {
        d.Set("tags", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_kubernetes_appgroup_msgidnamedisplayname(d, resp.Cluster); ok {
        d.Set("cluster", rtn)
    } else {
        d.Set("cluster", make([]map[string]interface{}, 0))
    }
    return nil
}

func resourceUpdateKubernetes_Appgroup(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V5/Kubernetes/ApplicationGroup/{applicationGroupId}
    var t_name *string
    if d.HasChange("name") {
        val := d.Get("name")
        t_name = handler.ToStringValue(val, false)
    }
    var t_filters *handler.MsgKubernetesApplicationGroupFilters
    if d.HasChange("filters") {
        val := d.Get("filters")
        t_filters = build_kubernetes_appgroup_msgkubernetesapplicationgroupfilters(d, val.([]interface{}))
    }
    var t_plan *handler.MsgIdName
    if d.HasChange("plan") {
        val := d.Get("plan")
        t_plan = build_kubernetes_appgroup_msgidname(d, val.([]interface{}))
    }
    var t_content *handler.MsgKubernetesApplicationGroupContent
    if d.HasChange("content") {
        val := d.Get("content")
        t_content = build_kubernetes_appgroup_msgkubernetesapplicationgroupcontent(d, val.([]interface{}))
    }
    var t_activitycontrol *handler.MsgApplicationGroupActivityControl
    if d.HasChange("activitycontrol") {
        val := d.Get("activitycontrol")
        t_activitycontrol = build_kubernetes_appgroup_msgapplicationgroupactivitycontrol(d, val.([]interface{}))
    }
    var t_timezone *handler.MsgIdName
    if d.HasChange("timezone") {
        val := d.Get("timezone")
        t_timezone = build_kubernetes_appgroup_msgidname(d, val.([]interface{}))
    }
    var t_options *handler.MsgApplicationGroupGetOptions
    if d.HasChange("options") {
        val := d.Get("options")
        t_options = build_kubernetes_appgroup_msgapplicationgroupgetoptions(d, val.([]interface{}))
    }
    var t_tags []handler.MsgNameValueSet
    if d.HasChange("tags") {
        val := d.Get("tags")
        t_tags = build_kubernetes_appgroup_msgnamevalueset_array(d, val.(*schema.Set).List())
    }
    var req = handler.MsgUpdateKubernetesAppGroupOpRequest{Name:t_name, Filters:t_filters, Plan:t_plan, Content:t_content, ActivityControl:t_activitycontrol, Timezone:t_timezone, Options:t_options, Tags:t_tags}
    _, err := handler.CvUpdateKubernetesAppGroupOp(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [UpdateKubernetesAppGroupOp] failed, Error %s", err)
    }
    return resourceReadKubernetes_Appgroup(d, m)
}

func resourceCreateUpdateKubernetes_Appgroup(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V5/Kubernetes/ApplicationGroup/{applicationGroupId}
    var execUpdate bool = false
    var t_activitycontrol *handler.MsgApplicationGroupActivityControl
    if val, ok := d.GetOk("activitycontrol"); ok {
        t_activitycontrol = build_kubernetes_appgroup_msgapplicationgroupactivitycontrol(d, val.([]interface{}))
        execUpdate = true
    }
    var t_timezone *handler.MsgIdName
    if val, ok := d.GetOk("timezone"); ok {
        t_timezone = build_kubernetes_appgroup_msgidname(d, val.([]interface{}))
        execUpdate = true
    }
    var t_options *handler.MsgApplicationGroupGetOptions
    if val, ok := d.GetOk("options"); ok {
        t_options = build_kubernetes_appgroup_msgapplicationgroupgetoptions(d, val.([]interface{}))
        execUpdate = true
    }
    var t_tags []handler.MsgNameValueSet
    if val, ok := d.GetOk("tags"); ok {
        t_tags = build_kubernetes_appgroup_msgnamevalueset_array(d, val.(*schema.Set).List())
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgUpdateKubernetesAppGroupOpRequest{ActivityControl:t_activitycontrol, Timezone:t_timezone, Options:t_options, Tags:t_tags}
        _, err := handler.CvUpdateKubernetesAppGroupOp(req, d.Id())
        if err != nil {
            return fmt.Errorf("operation [UpdateKubernetesAppGroupOp] failed, Error %s", err)
        }
    }
    return resourceReadKubernetes_Appgroup(d, m)
}

func resourceDeleteKubernetes_Appgroup(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V5/Kubernetes/ApplicationGroup/{applicationGroupId}
    _, err := handler.CvDeleteKubernetesAppGroup(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteKubernetesAppGroup] failed, Error %s", err)
    }
    return nil
}

func build_kubernetes_appgroup_msgnamevalueset_array(d *schema.ResourceData, r []interface{}) []handler.MsgNameValueSet {
    if r != nil {
        tmp := make([]handler.MsgNameValueSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_name *string
            if val, ok := raw_a["name"]; ok {
                t_name = handler.ToStringValue(val, true)
            }
            var t_value *string
            if val, ok := raw_a["value"]; ok {
                t_value = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgNameValueSet{Name:t_name, Value:t_value}
        }
        return tmp
    } else {
        return nil
    }
}

func build_kubernetes_appgroup_msgapplicationgroupgetoptions(d *schema.ResourceData, r []interface{}) *handler.MsgApplicationGroupGetOptions {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_backupstreams *int
        if val, ok := tmp["backupstreams"]; ok {
            t_backupstreams = handler.ToIntValue(val, true)
        }
        var t_cvnamespacescheduling *bool
        if val, ok := tmp["cvnamespacescheduling"]; ok {
            t_cvnamespacescheduling = handler.ToBooleanValue(val, true)
        }
        var t_workerresources *handler.MsgApplicationGroupWorkerResourcesOptions
        if val, ok := tmp["workerresources"]; ok {
            t_workerresources = build_kubernetes_appgroup_msgapplicationgroupworkerresourcesoptions(d, val.([]interface{}))
        }
        var t_snapfallbacktolivevolumebackup *bool
        if val, ok := tmp["snapfallbacktolivevolumebackup"]; ok {
            t_snapfallbacktolivevolumebackup = handler.ToBooleanValue(val, true)
        }
        var t_jobstarttime *int
        if val, ok := tmp["jobstarttime"]; ok {
            t_jobstarttime = handler.ToIntValue(val, true)
        }
        return &handler.MsgApplicationGroupGetOptions{BackupStreams:t_backupstreams, CvNamespaceScheduling:t_cvnamespacescheduling, WorkerResources:t_workerresources, SnapFallbackToLiveVolumeBackup:t_snapfallbacktolivevolumebackup, JobStartTime:t_jobstarttime}
    } else {
        return nil
    }
}

func build_kubernetes_appgroup_msgapplicationgroupworkerresourcesoptions(d *schema.ResourceData, r []interface{}) *handler.MsgApplicationGroupWorkerResourcesOptions {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_memoryrequests *string
        if val, ok := tmp["memoryrequests"]; ok {
            t_memoryrequests = handler.ToStringValue(val, true)
        }
        var t_memorylimits *string
        if val, ok := tmp["memorylimits"]; ok {
            t_memorylimits = handler.ToStringValue(val, true)
        }
        var t_cpulimits *string
        if val, ok := tmp["cpulimits"]; ok {
            t_cpulimits = handler.ToStringValue(val, true)
        }
        var t_cpurequests *string
        if val, ok := tmp["cpurequests"]; ok {
            t_cpurequests = handler.ToStringValue(val, true)
        }
        return &handler.MsgApplicationGroupWorkerResourcesOptions{MemoryRequests:t_memoryrequests, MemoryLimits:t_memorylimits, CpuLimits:t_cpulimits, CpuRequests:t_cpurequests}
    } else {
        return nil
    }
}

func build_kubernetes_appgroup_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_kubernetes_appgroup_msgapplicationgroupactivitycontrol(d *schema.ResourceData, r []interface{}) *handler.MsgApplicationGroupActivityControl {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_enablebackup *bool
        if val, ok := tmp["enablebackup"]; ok {
            t_enablebackup = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgApplicationGroupActivityControl{EnableBackup:t_enablebackup}
    } else {
        return nil
    }
}

func build_kubernetes_appgroup_msgkubernetesapplicationgroupcontent(d *schema.ResourceData, r []interface{}) *handler.MsgKubernetesApplicationGroupContent {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_overwrite *bool
        var c_overwrite bool = true
        t_overwrite = &c_overwrite
        var t_labelselectors []handler.MsgKubernetesContentSelectorsSet
        if val, ok := tmp["labelselectors"]; ok {
            t_labelselectors = build_kubernetes_appgroup_msgkubernetescontentselectorsset_array(d, val.(*schema.Set).List())
        }
        var t_applications []handler.MsgKubernetesContentApplicationsSet
        if val, ok := tmp["applications"]; ok {
            t_applications = build_kubernetes_appgroup_msgkubernetescontentapplicationsset_array(d, val.(*schema.Set).List())
        }
        return &handler.MsgKubernetesApplicationGroupContent{Overwrite:t_overwrite, LabelSelectors:t_labelselectors, Applications:t_applications}
    } else {
        return nil
    }
}

func build_kubernetes_appgroup_msgkubernetescontentapplicationsset_array(d *schema.ResourceData, r []interface{}) []handler.MsgKubernetesContentApplicationsSet {
    if r != nil {
        tmp := make([]handler.MsgKubernetesContentApplicationsSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_guid *string
            if val, ok := raw_a["guid"]; ok {
                t_guid = handler.ToStringValue(val, true)
            }
            var t_name *string
            if val, ok := raw_a["name"]; ok {
                t_name = handler.ToStringValue(val, true)
            }
            var t_type *string
            if val, ok := raw_a["type"]; ok {
                t_type = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgKubernetesContentApplicationsSet{GUID:t_guid, Name:t_name, Type:t_type}
        }
        return tmp
    } else {
        return nil
    }
}

func build_kubernetes_appgroup_msgkubernetescontentselectorsset_array(d *schema.ResourceData, r []interface{}) []handler.MsgKubernetesContentSelectorsSet {
    if r != nil {
        tmp := make([]handler.MsgKubernetesContentSelectorsSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_selectorlevel *string
            if val, ok := raw_a["selectorlevel"]; ok {
                t_selectorlevel = handler.ToStringValue(val, true)
            }
            var t_selectorvalue *string
            if val, ok := raw_a["selectorvalue"]; ok {
                t_selectorvalue = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgKubernetesContentSelectorsSet{SelectorLevel:t_selectorlevel, SelectorValue:t_selectorvalue}
        }
        return tmp
    } else {
        return nil
    }
}

func build_kubernetes_appgroup_msgkubernetesapplicationgroupfilters(d *schema.ResourceData, r []interface{}) *handler.MsgKubernetesApplicationGroupFilters {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_overwrite *bool
        var c_overwrite bool = true
        t_overwrite = &c_overwrite
        var t_skipstatelessapps *bool
        if val, ok := tmp["skipstatelessapps"]; ok {
            t_skipstatelessapps = handler.ToBooleanValue(val, true)
        }
        var t_labelselectors []handler.MsgKubernetesContentSelectorsSet
        if val, ok := tmp["labelselectors"]; ok {
            t_labelselectors = build_kubernetes_appgroup_msgkubernetescontentselectorsset_array(d, val.(*schema.Set).List())
        }
        var t_applications []handler.MsgKubernetesContentApplicationsSet
        if val, ok := tmp["applications"]; ok {
            t_applications = build_kubernetes_appgroup_msgkubernetescontentapplicationsset_array(d, val.(*schema.Set).List())
        }
        return &handler.MsgKubernetesApplicationGroupFilters{Overwrite:t_overwrite, SkipStatelessApps:t_skipstatelessapps, LabelSelectors:t_labelselectors, Applications:t_applications}
    } else {
        return nil
    }
}

func serialize_kubernetes_appgroup_msgidnamedisplayname(d *schema.ResourceData, data *handler.MsgIdNameDisplayName) ([]map[string]interface{}, bool) {
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

func serialize_kubernetes_appgroup_msgidnamevalueset_array(d *schema.ResourceData, data []handler.MsgIdNameValueSet) ([]map[string]interface{}, bool) {
    //MsgNameValueSet
    //MsgIdNameValueSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].Name != nil {
            tmp["name"] = data[i].Name
            added = true
        }
        if data[i].Value != nil {
            tmp["value"] = data[i].Value
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_kubernetes_appgroup_msgkubernetesapplicationgroupcontentitem(d *schema.ResourceData, data *handler.MsgKubernetesApplicationGroupContentItem) ([]map[string]interface{}, bool) {
    //MsgKubernetesApplicationGroupContent
    //MsgKubernetesApplicationGroupContentItem
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_kubernetes_appgroup_msgkubernetescontentselectorsset_array(d, data.LabelSelectors); ok {
        val[0]["labelselectors"] = rtn
        added = true
    }
    if rtn, ok := serialize_kubernetes_appgroup_msgkubernetescontentapplicationsset_array(d, data.Applications); ok {
        val[0]["applications"] = rtn
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_kubernetes_appgroup_msgkubernetescontentapplicationsset_array(d *schema.ResourceData, data []handler.MsgKubernetesContentApplicationsSet) ([]map[string]interface{}, bool) {
    //MsgKubernetesApplicationGroupContent -> MsgKubernetesContentApplicationsSet
    //MsgKubernetesApplicationGroupContentItem -> MsgKubernetesContentApplicationsSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].GUID != nil {
            tmp["guid"] = data[i].GUID
            added = true
        }
        if data[i].Name != nil {
            tmp["name"] = data[i].Name
            added = true
        }
        if data[i].Type != nil {
            tmp["type"] = data[i].Type
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_kubernetes_appgroup_msgkubernetescontentselectorsset_array(d *schema.ResourceData, data []handler.MsgKubernetesContentSelectorsSet) ([]map[string]interface{}, bool) {
    //MsgKubernetesApplicationGroupContent -> MsgKubernetesContentSelectorsSet
    //MsgKubernetesApplicationGroupContentItem -> MsgKubernetesContentSelectorsSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].SelectorLevel != nil {
            tmp["selectorlevel"] = data[i].SelectorLevel
            added = true
        }
        if data[i].SelectorValue != nil {
            tmp["selectorvalue"] = data[i].SelectorValue
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_kubernetes_appgroup_msgidname(d *schema.ResourceData, data *handler.MsgIdName) ([]map[string]interface{}, bool) {
    //MsgIdName
    //MsgIdName
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

func serialize_kubernetes_appgroup_msgkubernetesapplicationgroupfilteritem(d *schema.ResourceData, data *handler.MsgKubernetesApplicationGroupFilterItem) ([]map[string]interface{}, bool) {
    //MsgKubernetesApplicationGroupFilters
    //MsgKubernetesApplicationGroupFilterItem
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.SkipStatelessApps != nil {
        val[0]["skipstatelessapps"] = strconv.FormatBool(*data.SkipStatelessApps)
        added = true
    }
    if rtn, ok := serialize_kubernetes_appgroup_msgkubernetescontentselectorsset_array(d, data.LabelSelectors); ok {
        val[0]["labelselectors"] = rtn
        added = true
    }
    if rtn, ok := serialize_kubernetes_appgroup_msgkubernetescontentapplicationsset_array(d, data.Applications); ok {
        val[0]["applications"] = rtn
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_kubernetes_appgroup_msgapplicationgroupgetoptions(d *schema.ResourceData, data *handler.MsgApplicationGroupGetOptions) ([]map[string]interface{}, bool) {
    //MsgApplicationGroupGetOptions
    //MsgApplicationGroupGetOptions
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.BackupStreams != nil {
        val[0]["backupstreams"] = data.BackupStreams
        added = true
    }
    if data.CvNamespaceScheduling != nil {
        val[0]["cvnamespacescheduling"] = strconv.FormatBool(*data.CvNamespaceScheduling)
        added = true
    }
    if rtn, ok := serialize_kubernetes_appgroup_msgapplicationgroupworkerresourcesoptions(d, data.WorkerResources); ok {
        val[0]["workerresources"] = rtn
        added = true
    }
    if data.SnapFallbackToLiveVolumeBackup != nil {
        val[0]["snapfallbacktolivevolumebackup"] = strconv.FormatBool(*data.SnapFallbackToLiveVolumeBackup)
        added = true
    }
    if data.JobStartTime != nil {
        val[0]["jobstarttime"] = data.JobStartTime
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_kubernetes_appgroup_msgapplicationgroupworkerresourcesoptions(d *schema.ResourceData, data *handler.MsgApplicationGroupWorkerResourcesOptions) ([]map[string]interface{}, bool) {
    //MsgApplicationGroupGetOptions -> MsgApplicationGroupWorkerResourcesOptions
    //MsgApplicationGroupGetOptions -> MsgApplicationGroupWorkerResourcesOptions
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.MemoryRequests != nil {
        val[0]["memoryrequests"] = data.MemoryRequests
        added = true
    }
    if data.MemoryLimits != nil {
        val[0]["memorylimits"] = data.MemoryLimits
        added = true
    }
    if data.CpuLimits != nil {
        val[0]["cpulimits"] = data.CpuLimits
        added = true
    }
    if data.CpuRequests != nil {
        val[0]["cpurequests"] = data.CpuRequests
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_kubernetes_appgroup_msgapplicationgroupactivitycontrol(d *schema.ResourceData, data *handler.MsgApplicationGroupActivityControl) ([]map[string]interface{}, bool) {
    //MsgApplicationGroupActivityControl
    //MsgApplicationGroupActivityControl
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.EnableBackup != nil {
        val[0]["enablebackup"] = strconv.FormatBool(*data.EnableBackup)
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}
