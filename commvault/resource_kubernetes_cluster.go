package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceKubernetes_Cluster() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateKubernetes_Cluster,
        Read:   resourceReadKubernetes_Cluster,
        Update: resourceUpdateKubernetes_Cluster,
        Delete: resourceDeleteKubernetes_Cluster,

        Schema: map[string]*schema.Schema{
            "apiserver": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "API Server Endpoint of the cluster",
            },
            "serviceaccount": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Name of the Service Account to authenticate with the cluster",
            },
            "servicetoken": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Secret token to authenticate with the cluster",
            },
            "accessnodes": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: "",
                        },
                        "type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "",
                        },
                    },
                },
            },
            "servicetype": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "The Service Type of the Kubernetes cluster [ONPREM, AKS]",
            },
            "etcdprotection": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "ETCD Protection options for a cluster",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
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
                        "enabled": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Denote if etcd protection is enabled",
                        },
                    },
                },
            },
            "name": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Name of the Kubernetes Cluster",
            },
            "activitycontrol": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Request definition changing activity control options for cluster",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enablebackupafteradelay": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Enabling backup after a delay. Provide UTC Time in Unix format",
                        },
                        "enablebackup": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Enable or disable backup for cluster",
                        },
                        "enablerestoreafteradelay": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Enabling restore after a delay. Provide UTC Time in Unix format",
                        },
                        "enablerestore": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Enable or disable restore for cluster",
                        },
                    },
                },
            },
            "options": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Request definition for cluster advanced options",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "imageregistry": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "Request definition changing image registry options for cluster",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "registryurl": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Specify image registry URL for internal image repository",
                                    },
                                    "imagepullsecret": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Specify image pull secret to authenticate with the image repository",
                                    },
                                },
                            },
                        },
                    },
                },
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
            "tags": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "Modify or add tags on the cluster",
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

func resourceCreateKubernetes_Cluster(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V5/Kubernetes/Cluster
    var response_id = strconv.Itoa(0)
    var t_apiserver *string
    if val, ok := d.GetOk("apiserver"); ok {
        t_apiserver = handler.ToStringValue(val, false)
    }
    var t_serviceaccount *string
    if val, ok := d.GetOk("serviceaccount"); ok {
        t_serviceaccount = handler.ToStringValue(val, false)
    }
    var t_servicetoken *string
    if val, ok := d.GetOk("servicetoken"); ok {
        t_servicetoken = handler.ToStringValue(val, false)
    }
    var t_accessnodes []handler.MsgIdNameTypeSet
    if val, ok := d.GetOk("accessnodes"); ok {
        t_accessnodes = build_kubernetes_cluster_msgidnametypeset_array(d, val.(*schema.Set).List())
    }
    var t_servicetype *string
    if val, ok := d.GetOk("servicetype"); ok {
        t_servicetype = handler.ToStringValue(val, false)
    }
    var t_etcdprotection *handler.MsgEtcdProtectionItem
    if val, ok := d.GetOk("etcdprotection"); ok {
        t_etcdprotection = build_kubernetes_cluster_msgetcdprotectionitem(d, val.([]interface{}))
    }
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var req = handler.MsgCreateKubernetesClusterOpRequest{ApiServer:t_apiserver, ServiceAccount:t_serviceaccount, ServiceToken:t_servicetoken, AccessNodes:t_accessnodes, ServiceType:t_servicetype, EtcdProtection:t_etcdprotection, Name:t_name}
    resp, err := handler.CvCreateKubernetesClusterOp(req)
    if err != nil {
        return fmt.Errorf("operation [CreateKubernetesClusterOp] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateKubernetesClusterOp] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateKubernetes_Cluster(d, m)
    }
}

func resourceReadKubernetes_Cluster(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V5/Kubernetes/Cluster/{clusterId}
    resp, err := handler.CvGetKubernetesClusterDetails(d.Id())
    if err != nil {
        return fmt.Errorf("operation [GetKubernetesClusterDetails] failed, Error %s", err)
    }
    if rtn, ok := serialize_kubernetes_cluster_msgclusteractivitycontroloptions(d, resp.ActivityControl); ok {
        d.Set("activitycontrol", rtn)
    } else {
        d.Set("activitycontrol", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_kubernetes_cluster_msggetetcdprotectionitem(d, resp.EtcdProtection); ok {
        d.Set("etcdprotection", rtn)
    } else {
        d.Set("etcdprotection", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_kubernetes_cluster_msgeditclusteradvancedoptionsinfo(d, resp.Options); ok {
        d.Set("options", rtn)
    } else {
        d.Set("options", make([]map[string]interface{}, 0))
    }
    if resp.ApiServer != nil {
        d.Set("apiserver", resp.ApiServer)
    }
    if rtn, ok := serialize_kubernetes_cluster_msgidname(d, resp.Region); ok {
        d.Set("region", rtn)
    } else {
        d.Set("region", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_kubernetes_cluster_msgidnamevalueset_array(d, resp.Tags); ok {
        d.Set("tags", rtn)
    } else {
        d.Set("tags", make([]map[string]interface{}, 0))
    }
    if resp.DisplayName != nil {
        d.Set("name", resp.DisplayName)
    }
    return nil
}

func resourceUpdateKubernetes_Cluster(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V5/Kubernetes/Cluster/{clusterId}
    var t_apiserver *string
    if d.HasChange("apiserver") {
        val := d.Get("apiserver")
        t_apiserver = handler.ToStringValue(val, false)
    }
    var t_serviceaccount *string
    if d.HasChange("serviceaccount") {
        val := d.Get("serviceaccount")
        t_serviceaccount = handler.ToStringValue(val, false)
    }
    var t_servicetoken *string
    if d.HasChange("servicetoken") {
        val := d.Get("servicetoken")
        t_servicetoken = handler.ToStringValue(val, false)
    }
    var t_accessnodes []handler.MsgIdNameTypeSet
    if d.HasChange("accessnodes") {
        val := d.Get("accessnodes")
        t_accessnodes = build_kubernetes_cluster_msgidnametypeset_array(d, val.(*schema.Set).List())
    }
    var t_servicetype *string
    if d.HasChange("servicetype") {
        val := d.Get("servicetype")
        t_servicetype = handler.ToStringValue(val, false)
    }
    var t_activitycontrol *handler.MsgEditClusterActivityControlOptions
    if d.HasChange("activitycontrol") {
        val := d.Get("activitycontrol")
        t_activitycontrol = build_kubernetes_cluster_msgeditclusteractivitycontroloptions(d, val.([]interface{}))
    }
    var t_etcdprotection *handler.MsgGetEtcdProtectionItem
    if d.HasChange("etcdprotection") {
        val := d.Get("etcdprotection")
        t_etcdprotection = build_kubernetes_cluster_msggetetcdprotectionitem(d, val.([]interface{}))
    }
    var t_name *string
    if d.HasChange("name") {
        val := d.Get("name")
        t_name = handler.ToStringValue(val, false)
    }
    var t_options *handler.MsgEditClusterAdvancedOptionsInfo
    if d.HasChange("options") {
        val := d.Get("options")
        t_options = build_kubernetes_cluster_msgeditclusteradvancedoptionsinfo(d, val.([]interface{}))
    }
    var t_region *handler.MsgIdName
    if d.HasChange("region") {
        val := d.Get("region")
        t_region = build_kubernetes_cluster_msgidname(d, val.([]interface{}))
    }
    var t_tags []handler.MsgNameValueSet
    if d.HasChange("tags") {
        val := d.Get("tags")
        t_tags = build_kubernetes_cluster_msgnamevalueset_array(d, val.(*schema.Set).List())
    }
    var req = handler.MsgUpdateKubernetesPropertiesRequest{ApiServer:t_apiserver, ServiceAccount:t_serviceaccount, ServiceToken:t_servicetoken, AccessNodes:t_accessnodes, ServiceType:t_servicetype, ActivityControl:t_activitycontrol, EtcdProtection:t_etcdprotection, Name:t_name, Options:t_options, Region:t_region, Tags:t_tags}
    _, err := handler.CvUpdateKubernetesProperties(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [UpdateKubernetesProperties] failed, Error %s", err)
    }
    return resourceReadKubernetes_Cluster(d, m)
}

func resourceCreateUpdateKubernetes_Cluster(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V5/Kubernetes/Cluster/{clusterId}
    var execUpdate bool = false
    var t_activitycontrol *handler.MsgEditClusterActivityControlOptions
    if val, ok := d.GetOk("activitycontrol"); ok {
        t_activitycontrol = build_kubernetes_cluster_msgeditclusteractivitycontroloptions(d, val.([]interface{}))
        execUpdate = true
    }
    var t_options *handler.MsgEditClusterAdvancedOptionsInfo
    if val, ok := d.GetOk("options"); ok {
        t_options = build_kubernetes_cluster_msgeditclusteradvancedoptionsinfo(d, val.([]interface{}))
        execUpdate = true
    }
    var t_region *handler.MsgIdName
    if val, ok := d.GetOk("region"); ok {
        t_region = build_kubernetes_cluster_msgidname(d, val.([]interface{}))
        execUpdate = true
    }
    var t_tags []handler.MsgNameValueSet
    if val, ok := d.GetOk("tags"); ok {
        t_tags = build_kubernetes_cluster_msgnamevalueset_array(d, val.(*schema.Set).List())
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgUpdateKubernetesPropertiesRequest{ActivityControl:t_activitycontrol, Options:t_options, Region:t_region, Tags:t_tags}
        _, err := handler.CvUpdateKubernetesProperties(req, d.Id())
        if err != nil {
            return fmt.Errorf("operation [UpdateKubernetesProperties] failed, Error %s", err)
        }
    }
    return resourceReadKubernetes_Cluster(d, m)
}

func resourceDeleteKubernetes_Cluster(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V5/Kubernetes/Cluster/{clusterId}/Retire
    _, err := handler.CvRetireKubernetesCluster(d.Id())
    if err != nil {
        return fmt.Errorf("operation [RetireKubernetesCluster] failed, Error %s", err)
    }
    return nil
}

func build_kubernetes_cluster_msgnamevalueset_array(d *schema.ResourceData, r []interface{}) []handler.MsgNameValueSet {
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

func build_kubernetes_cluster_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_kubernetes_cluster_msgeditclusteradvancedoptionsinfo(d *schema.ResourceData, r []interface{}) *handler.MsgEditClusterAdvancedOptionsInfo {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_imageregistry *handler.MsgClusterImageRegistryOptions
        if val, ok := tmp["imageregistry"]; ok {
            t_imageregistry = build_kubernetes_cluster_msgclusterimageregistryoptions(d, val.([]interface{}))
        }
        return &handler.MsgEditClusterAdvancedOptionsInfo{ImageRegistry:t_imageregistry}
    } else {
        return nil
    }
}

func build_kubernetes_cluster_msgclusterimageregistryoptions(d *schema.ResourceData, r []interface{}) *handler.MsgClusterImageRegistryOptions {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_registryurl *string
        if val, ok := tmp["registryurl"]; ok {
            t_registryurl = handler.ToStringValue(val, true)
        }
        var t_imagepullsecret *string
        if val, ok := tmp["imagepullsecret"]; ok {
            t_imagepullsecret = handler.ToStringValue(val, true)
        }
        return &handler.MsgClusterImageRegistryOptions{RegistryUrl:t_registryurl, ImagePullSecret:t_imagepullsecret}
    } else {
        return nil
    }
}

func build_kubernetes_cluster_msgeditclusteractivitycontroloptions(d *schema.ResourceData, r []interface{}) *handler.MsgEditClusterActivityControlOptions {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_enablebackupafteradelay *int
        if val, ok := tmp["enablebackupafteradelay"]; ok {
            t_enablebackupafteradelay = handler.ToIntValue(val, true)
        }
        var t_enablebackup *bool
        if val, ok := tmp["enablebackup"]; ok {
            t_enablebackup = handler.ToBooleanValue(val, true)
        }
        var t_enablerestoreafteradelay *int
        if val, ok := tmp["enablerestoreafteradelay"]; ok {
            t_enablerestoreafteradelay = handler.ToIntValue(val, true)
        }
        var t_enablerestore *bool
        if val, ok := tmp["enablerestore"]; ok {
            t_enablerestore = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgEditClusterActivityControlOptions{EnableBackupAfterADelay:t_enablebackupafteradelay, EnableBackup:t_enablebackup, EnableRestoreAfterADelay:t_enablerestoreafteradelay, EnableRestore:t_enablerestore}
    } else {
        return nil
    }
}

func build_kubernetes_cluster_msggetetcdprotectionitem(d *schema.ResourceData, r []interface{}) *handler.MsgGetEtcdProtectionItem {
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
        var t_plan *handler.MsgIdName
        if val, ok := tmp["plan"]; ok {
            t_plan = build_kubernetes_cluster_msgidname(d, val.([]interface{}))
        }
        var t_enabled *bool
        if val, ok := tmp["enabled"]; ok {
            t_enabled = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgGetEtcdProtectionItem{Name:t_name, Id:t_id, Plan:t_plan, Enabled:t_enabled}
    } else {
        return nil
    }
}

func build_kubernetes_cluster_msgidnametypeset_array(d *schema.ResourceData, r []interface{}) []handler.MsgIdNameTypeSet {
    if r != nil {
        tmp := make([]handler.MsgIdNameTypeSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_id *int
            if val, ok := raw_a["id"]; ok {
                t_id = handler.ToIntValue(val, true)
            }
            var t_type *string
            if val, ok := raw_a["type"]; ok {
                t_type = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgIdNameTypeSet{Id:t_id, Type:t_type}
        }
        return tmp
    } else {
        return nil
    }
}

func build_kubernetes_cluster_msgetcdprotectionitem(d *schema.ResourceData, r []interface{}) *handler.MsgEtcdProtectionItem {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_plan *handler.MsgIdName
        if val, ok := tmp["plan"]; ok {
            t_plan = build_kubernetes_cluster_msgidname(d, val.([]interface{}))
        }
        var t_enabled *bool
        if val, ok := tmp["enabled"]; ok {
            t_enabled = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgEtcdProtectionItem{Plan:t_plan, Enabled:t_enabled}
    } else {
        return nil
    }
}

func serialize_kubernetes_cluster_msgidnamevalueset_array(d *schema.ResourceData, data []handler.MsgIdNameValueSet) ([]map[string]interface{}, bool) {
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

func serialize_kubernetes_cluster_msgidname(d *schema.ResourceData, data *handler.MsgIdName) ([]map[string]interface{}, bool) {
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

func serialize_kubernetes_cluster_msgeditclusteradvancedoptionsinfo(d *schema.ResourceData, data *handler.MsgEditClusterAdvancedOptionsInfo) ([]map[string]interface{}, bool) {
    //MsgEditClusterAdvancedOptionsInfo
    //MsgEditClusterAdvancedOptionsInfo
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_kubernetes_cluster_msgclusterimageregistryoptions(d, data.ImageRegistry); ok {
        val[0]["imageregistry"] = rtn
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_kubernetes_cluster_msgclusterimageregistryoptions(d *schema.ResourceData, data *handler.MsgClusterImageRegistryOptions) ([]map[string]interface{}, bool) {
    //MsgEditClusterAdvancedOptionsInfo -> MsgClusterImageRegistryOptions
    //MsgEditClusterAdvancedOptionsInfo -> MsgClusterImageRegistryOptions
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.RegistryUrl != nil {
        val[0]["registryurl"] = data.RegistryUrl
        added = true
    }
    if data.ImagePullSecret != nil {
        val[0]["imagepullsecret"] = data.ImagePullSecret
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_kubernetes_cluster_msggetetcdprotectionitem(d *schema.ResourceData, data *handler.MsgGetEtcdProtectionItem) ([]map[string]interface{}, bool) {
    //MsgEtcdProtectionItem
    //MsgGetEtcdProtectionItem
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_kubernetes_cluster_msgidname(d, data.Plan); ok {
        val[0]["plan"] = rtn
        added = true
    }
    if data.Enabled != nil {
        val[0]["enabled"] = strconv.FormatBool(*data.Enabled)
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_kubernetes_cluster_msgclusteractivitycontroloptions(d *schema.ResourceData, data *handler.MsgClusterActivityControlOptions) ([]map[string]interface{}, bool) {
    //MsgEditClusterActivityControlOptions
    //MsgClusterActivityControlOptions
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
    if data.EnableRestore != nil {
        val[0]["enablerestore"] = strconv.FormatBool(*data.EnableRestore)
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}
