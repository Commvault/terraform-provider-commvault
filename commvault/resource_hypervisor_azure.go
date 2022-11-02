package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceHypervisor_Azure() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateHypervisor_Azure,
        Read:   resourceReadHypervisor_Azure,
        Update: resourceUpdateHypervisor_Azure,
        Delete: resourceDeleteHypervisor_Azure,

        Schema: map[string]*schema.Schema{
            "skipcredentialvalidation": &schema.Schema{
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: "if credential validation has to be skipped.",
            },
            "etcdprotection": &schema.Schema{
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: "Flag to create an application group etcd (system generated) with pre-defined content",
            },
            "credentials": &schema.Schema{
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
            "name": &schema.Schema{
                Type:        schema.TypeString,
                Required:    true,
                Description: "The name of the hypervisor group being created",
            },
            "accessnodes": &schema.Schema{
                Type:        schema.TypeSet,
                Required:    true,
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
            "planentity": &schema.Schema{
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
            "applicationpassword": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Application Password of Azure login Application",
            },
            "tenantid": &schema.Schema{
                Type:        schema.TypeString,
                Required:    true,
                Description: "Tenant id of Azure login Application",
            },
            "workloadregion": &schema.Schema{
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
            "subscriptionid": &schema.Schema{
                Type:        schema.TypeString,
                Required:    true,
                Description: "subscription id of Azure ",
            },
            "applicationid": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Application id of Azure login Application",
            },
            "activitycontrol": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "restoreactivitycontroloptions": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
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
                                    "enableafterdelay": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if the activity will be enabled after a delay time interval",
                                    },
                                    "delaytime": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Delayed by n Hrs",
                                    },
                                },
                            },
                        },
                        "backupactivitycontroloptions": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
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
                                    "enableafterdelay": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if the activity will be enabled after a delay time interval",
                                    },
                                    "delaytime": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Delayed by n Hrs",
                                    },
                                },
                            },
                        },
                        "enablebackup": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "true if Backup is enabled",
                        },
                        "enablerestore": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "true if Restore is enabled",
                        },
                    },
                },
            },
            "settings": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "metricsmonitoringpolicy": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "isenabled": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if Metrics Monioring policy is enabled",
                                    },
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Metrics Moitoring Policy Name",
                                    },
                                    "id": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Metrics Moitoring PolicyId",
                                    },
                                },
                            },
                        },
                        "applicationcredentials": &schema.Schema{
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
                        "guestcredentials": &schema.Schema{
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
                        "mountaccessnode": &schema.Schema{
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
                        "regioninfo": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "displayname": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Display Name of Region",
                                    },
                                    "latitude": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Geolocation Latitude",
                                    },
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Region Name",
                                    },
                                    "id": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Region Id",
                                    },
                                    "longitude": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Geolocation Longitude",
                                    },
                                },
                            },
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
                        "customattributes": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "type": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "client custom attribute type . Ex- 3 - For client 8- For clientGroup",
                                    },
                                    "value": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "client/Client Group custom attribute value",
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "security": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "clientowners": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Client owners for the Hypervisor",
                        },
                        "associatedusergroups": &schema.Schema{
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
            "newname": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "The name of the hypervisor that has to be changed",
            },
            "accessnode": &schema.Schema{
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
            "fbrunixmediaagent": &schema.Schema{
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
            "password": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Application Password of Azure login Application",
            },
            "servername": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Client Name to Update",
            },
            "hypervisortype": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "",
            },
            "username": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Application id of Azure login Application",
            },
        },
    }
}

func resourceCreateHypervisor_Azure(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/Hypervisor
    var response_id = strconv.Itoa(0)
    var t_skipcredentialvalidation *bool
    if v, ok := d.GetOkExists("skipcredentialvalidation"); ok {
        val := v.(bool)
        t_skipcredentialvalidation = new(bool)
        t_skipcredentialvalidation = &val
    }
    var t_etcdprotection *bool
    if v, ok := d.GetOkExists("etcdprotection"); ok {
        val := v.(bool)
        t_etcdprotection = new(bool)
        t_etcdprotection = &val
    }
    var t_credentials *handler.MsgIdName
    if v, ok := d.GetOk("credentials"); ok {
        val := v.([]interface{})
        t_credentials = build_hypervisor_azure_msgidname(d, val)
    }
    var t_name *string
    if v, ok := d.GetOk("name"); ok {
        val := v.(string)
        t_name = new(string)
        t_name = &val
    }
    var t_accessnodes []handler.MsgIdNameSet
    if v, ok := d.GetOk("accessnodes"); ok {
        val := v.(*schema.Set)
        t_accessnodes = build_hypervisor_azure_msgidnameset_array(d, val.List())
    }
    var t_planentity *handler.MsgIdName
    if v, ok := d.GetOk("planentity"); ok {
        val := v.([]interface{})
        t_planentity = build_hypervisor_azure_msgidname(d, val)
    }
    var t_applicationpassword *string
    if v, ok := d.GetOk("applicationpassword"); ok {
        val := v.(string)
        t_applicationpassword = new(string)
        t_applicationpassword = &val
    }
    var t_tenantid *string
    if v, ok := d.GetOk("tenantid"); ok {
        val := v.(string)
        t_tenantid = new(string)
        t_tenantid = &val
    }
    var t_hypervisortype *string
    var c_hypervisortype string = "AZURE_V2"
    t_hypervisortype = new(string)
    t_hypervisortype = &c_hypervisortype
    var t_workloadregion *handler.MsgIdName
    if v, ok := d.GetOk("workloadregion"); ok {
        val := v.([]interface{})
        t_workloadregion = build_hypervisor_azure_msgidname(d, val)
    }
    var t_subscriptionid *string
    if v, ok := d.GetOk("subscriptionid"); ok {
        val := v.(string)
        t_subscriptionid = new(string)
        t_subscriptionid = &val
    }
    var t_applicationid *string
    if v, ok := d.GetOk("applicationid"); ok {
        val := v.(string)
        t_applicationid = new(string)
        t_applicationid = &val
    }
    var req = handler.MsgCreateHypervisorAzureRequest{SkipCredentialValidation:t_skipcredentialvalidation, EtcdProtection:t_etcdprotection, Credentials:t_credentials, Name:t_name, AccessNodes:t_accessnodes, PlanEntity:t_planentity, ApplicationPassword:t_applicationpassword, TenantId:t_tenantid, HypervisorType:t_hypervisortype, WorkloadRegion:t_workloadregion, SubscriptionId:t_subscriptionid, ApplicationId:t_applicationid}
    resp, err := handler.CvCreateHypervisorAzure(req)
    if err != nil {
        return fmt.Errorf("Operation [CreateHypervisorAzure] failed, Error %s", err)
    }
    if resp.Response != nil {
        if resp.Response.HypervisorId != nil {
            response_id = strconv.Itoa(*resp.Response.HypervisorId)
        }
    }
    if response_id == "0" {
        return fmt.Errorf("Operation [CreateHypervisorAzure] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateHypervisor_Azure(d, m)
    }
}

func resourceReadHypervisor_Azure(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/Hypervisor/{hypervisorId}
    resp, err := handler.CvGetHypervisors(d.Id())
    if err != nil {
        return fmt.Errorf("Operation [GetHypervisors] failed, Error %s", err)
    }
    if resp.ActivityControl != nil {
        d.Set("activitycontrol", serialize_hypervisor_azure_msgactivitycontroloptions(resp.ActivityControl))
    } else {
        d.Set("activitycontrol", make([]map[string]interface{}, 0))
    }
    if resp.Settings != nil {
        d.Set("settings", serialize_hypervisor_azure_msghypervisorsettings(resp.Settings))
    } else {
        d.Set("settings", make([]map[string]interface{}, 0))
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    return nil
}

func resourceUpdateHypervisor_Azure(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Hypervisor/{hypervisorId}
    var t_activitycontrol *handler.MsgActivityControlOptions
    if d.HasChange("activitycontrol") {
        val := d.Get("activitycontrol").([]interface{})
        t_activitycontrol = build_hypervisor_azure_msgactivitycontroloptions(d, val)
    }
    var t_settings *handler.MsghypervisorSettings
    if d.HasChange("settings") {
        val := d.Get("settings").([]interface{})
        t_settings = build_hypervisor_azure_msghypervisorsettings(d, val)
    }
    var t_security *handler.MsgVMHypervisorSecurityProp
    if d.HasChange("security") {
        val := d.Get("security").([]interface{})
        t_security = build_hypervisor_azure_msgvmhypervisorsecurityprop(d, val)
    }
    var t_newname *string
    if d.HasChange("newname") {
        val := d.Get("newname").(string)
        t_newname = new(string)
        t_newname = &val
    }
    var t_skipcredentialvalidation *bool
    if d.HasChange("skipcredentialvalidation") {
        val := d.Get("skipcredentialvalidation").(bool)
        t_skipcredentialvalidation = new(bool)
        t_skipcredentialvalidation = &val
    }
    var t_accessnode []handler.MsgaccessNodeModelSet
    if d.HasChange("accessnode") {
        val := d.Get("accessnode").(*schema.Set)
        t_accessnode = build_hypervisor_azure_msgaccessnodemodelset_array(d, val.List())
    }
    var t_fbrunixmediaagent *handler.MsgIdName
    if d.HasChange("fbrunixmediaagent") {
        val := d.Get("fbrunixmediaagent").([]interface{})
        t_fbrunixmediaagent = build_hypervisor_azure_msgidname(d, val)
    }
    var t_password *string
    if d.HasChange("password") {
        val := d.Get("password").(string)
        t_password = new(string)
        t_password = &val
    }
    var t_tenantid *string
    if d.HasChange("tenantid") {
        val := d.Get("tenantid").(string)
        t_tenantid = new(string)
        t_tenantid = &val
    }
    var t_servername *string
    if d.HasChange("servername") {
        val := d.Get("servername").(string)
        t_servername = new(string)
        t_servername = &val
    }
    var t_hypervisortype *string
    if d.HasChange("hypervisortype") {
        val := d.Get("hypervisortype").(string)
        t_hypervisortype = new(string)
        t_hypervisortype = &val
    }
    var t_subscriptionid *string
    if d.HasChange("subscriptionid") {
        val := d.Get("subscriptionid").(string)
        t_subscriptionid = new(string)
        t_subscriptionid = &val
    }
    var t_username *string
    if d.HasChange("username") {
        val := d.Get("username").(string)
        t_username = new(string)
        t_username = &val
    }
    var req = handler.MsgupdateHypervisorAzureRequest{ActivityControl:t_activitycontrol, Settings:t_settings, Security:t_security, NewName:t_newname, SkipCredentialValidation:t_skipcredentialvalidation, AccessNode:t_accessnode, FbrUnixMediaAgent:t_fbrunixmediaagent, Password:t_password, TenantId:t_tenantid, ServerName:t_servername, HypervisorType:t_hypervisortype, SubscriptionId:t_subscriptionid, UserName:t_username}
    _, err := handler.CvupdateHypervisorAzure(req, d.Id())
    if err != nil {
        return fmt.Errorf("Operation [updateHypervisorAzure] failed, Error %s", err)
    }
    return resourceReadHypervisor_Azure(d, m)
}

func resourceCreateUpdateHypervisor_Azure(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Hypervisor/{hypervisorId}
    var execUpdate bool = false
    var t_activitycontrol *handler.MsgActivityControlOptions
    if v, ok := d.GetOk("activitycontrol"); ok {
        val := v.([]interface{})
        t_activitycontrol = build_hypervisor_azure_msgactivitycontroloptions(d, val)
        execUpdate = true
    }
    var t_settings *handler.MsghypervisorSettings
    if v, ok := d.GetOk("settings"); ok {
        val := v.([]interface{})
        t_settings = build_hypervisor_azure_msghypervisorsettings(d, val)
        execUpdate = true
    }
    var t_security *handler.MsgVMHypervisorSecurityProp
    if v, ok := d.GetOk("security"); ok {
        val := v.([]interface{})
        t_security = build_hypervisor_azure_msgvmhypervisorsecurityprop(d, val)
        execUpdate = true
    }
    var t_newname *string
    if v, ok := d.GetOk("newname"); ok {
        val := v.(string)
        t_newname = new(string)
        t_newname = &val
        execUpdate = true
    }
    var t_accessnode []handler.MsgaccessNodeModelSet
    if v, ok := d.GetOk("accessnode"); ok {
        val := v.(*schema.Set)
        t_accessnode = build_hypervisor_azure_msgaccessnodemodelset_array(d, val.List())
        execUpdate = true
    }
    var t_fbrunixmediaagent *handler.MsgIdName
    if v, ok := d.GetOk("fbrunixmediaagent"); ok {
        val := v.([]interface{})
        t_fbrunixmediaagent = build_hypervisor_azure_msgidname(d, val)
        execUpdate = true
    }
    var t_password *string
    if v, ok := d.GetOk("password"); ok {
        val := v.(string)
        t_password = new(string)
        t_password = &val
        execUpdate = true
    }
    var t_servername *string
    if v, ok := d.GetOk("servername"); ok {
        val := v.(string)
        t_servername = new(string)
        t_servername = &val
        execUpdate = true
    }
    var t_username *string
    if v, ok := d.GetOk("username"); ok {
        val := v.(string)
        t_username = new(string)
        t_username = &val
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgupdateHypervisorAzureRequest{ActivityControl:t_activitycontrol, Settings:t_settings, Security:t_security, NewName:t_newname, AccessNode:t_accessnode, FbrUnixMediaAgent:t_fbrunixmediaagent, Password:t_password, ServerName:t_servername, UserName:t_username}
        _, err := handler.CvupdateHypervisorAzure(req, d.Id())
        if err != nil {
            return fmt.Errorf("Operation [updateHypervisorAzure] failed, Error %s", err)
        }
    }
    return resourceReadHypervisor_Azure(d, m)
}

func resourceDeleteHypervisor_Azure(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/Hypervisor/{hypervisorId}
    _, err := handler.CvDeleteHypervisor(d.Id())
    if err != nil {
        return fmt.Errorf("Operation [DeleteHypervisor] failed, Error %s", err)
    }
    return nil
}

func build_hypervisor_azure_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_hypervisor_azure_msgaccessnodemodelset_array(d *schema.ResourceData, r []interface{}) []handler.MsgaccessNodeModelSet {
    if r != nil {
        tmp := make([]handler.MsgaccessNodeModelSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_id *int
            if val, ok := raw_a["id"].(int); ok {
                t_id = new(int)
                t_id = &val
            }
            tmp[a] = handler.MsgaccessNodeModelSet{Id:t_id}
        }
        return tmp
    } else {
        return nil
    }
}

func build_hypervisor_azure_msgvmhypervisorsecurityprop(d *schema.ResourceData, r []interface{}) *handler.MsgVMHypervisorSecurityProp {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_clientowners *string
        if val, ok := tmp["clientowners"].(string); ok {
            t_clientowners = new(string)
            t_clientowners = &val
        }
        var t_associatedusergroups []handler.MsgIdNameSet
        if val, ok := tmp["associatedusergroups"].(*schema.Set); ok {
            t_associatedusergroups = build_hypervisor_azure_msgidnameset_array(d, val.List())
        }
        return &handler.MsgVMHypervisorSecurityProp{ClientOwners:t_clientowners, AssociatedUserGroups:t_associatedusergroups}
    } else {
        return nil
    }
}

func build_hypervisor_azure_msgidnameset_array(d *schema.ResourceData, r []interface{}) []handler.MsgIdNameSet {
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

func build_hypervisor_azure_msghypervisorsettings(d *schema.ResourceData, r []interface{}) *handler.MsghypervisorSettings {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_metricsmonitoringpolicy *handler.MsghypervisorMonitoringPolicy
        if val, ok := tmp["metricsmonitoringpolicy"].([]interface{}); ok {
            t_metricsmonitoringpolicy = build_hypervisor_azure_msghypervisormonitoringpolicy(d, val)
        }
        var t_applicationcredentials *handler.MsgIdName
        if val, ok := tmp["applicationcredentials"].([]interface{}); ok {
            t_applicationcredentials = build_hypervisor_azure_msgidname(d, val)
        }
        var t_guestcredentials *handler.MsgIdName
        if val, ok := tmp["guestcredentials"].([]interface{}); ok {
            t_guestcredentials = build_hypervisor_azure_msgidname(d, val)
        }
        var t_mountaccessnode *handler.MsgIdName
        if val, ok := tmp["mountaccessnode"].([]interface{}); ok {
            t_mountaccessnode = build_hypervisor_azure_msgidname(d, val)
        }
        var t_regioninfo *handler.MsgRegionInfo
        if val, ok := tmp["regioninfo"].([]interface{}); ok {
            t_regioninfo = build_hypervisor_azure_msgregioninfo(d, val)
        }
        var t_timezone *handler.MsgIdName
        if val, ok := tmp["timezone"].([]interface{}); ok {
            t_timezone = build_hypervisor_azure_msgidname(d, val)
        }
        var t_customattributes *handler.MsghypervisorCustomAttribute
        if val, ok := tmp["customattributes"].([]interface{}); ok {
            t_customattributes = build_hypervisor_azure_msghypervisorcustomattribute(d, val)
        }
        return &handler.MsghypervisorSettings{MetricsMonitoringPolicy:t_metricsmonitoringpolicy, ApplicationCredentials:t_applicationcredentials, GuestCredentials:t_guestcredentials, MountAccessNode:t_mountaccessnode, RegionInfo:t_regioninfo, TimeZone:t_timezone, CustomAttributes:t_customattributes}
    } else {
        return nil
    }
}

func build_hypervisor_azure_msghypervisorcustomattribute(d *schema.ResourceData, r []interface{}) *handler.MsghypervisorCustomAttribute {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_type *int
        if val, ok := tmp["type"].(int); ok {
            t_type = new(int)
            t_type = &val
        }
        var t_value *string
        if val, ok := tmp["value"].(string); ok {
            t_value = new(string)
            t_value = &val
        }
        return &handler.MsghypervisorCustomAttribute{Type:t_type, Value:t_value}
    } else {
        return nil
    }
}

func build_hypervisor_azure_msgregioninfo(d *schema.ResourceData, r []interface{}) *handler.MsgRegionInfo {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_displayname *string
        if val, ok := tmp["displayname"].(string); ok {
            t_displayname = new(string)
            t_displayname = &val
        }
        var t_latitude *string
        if val, ok := tmp["latitude"].(string); ok {
            t_latitude = new(string)
            t_latitude = &val
        }
        var t_name *string
        if val, ok := tmp["name"].(string); ok {
            t_name = new(string)
            t_name = &val
        }
        var t_id *int
        if val, ok := tmp["id"].(int); ok {
            t_id = new(int)
            t_id = &val
        }
        var t_longitude *string
        if val, ok := tmp["longitude"].(string); ok {
            t_longitude = new(string)
            t_longitude = &val
        }
        return &handler.MsgRegionInfo{DisplayName:t_displayname, Latitude:t_latitude, Name:t_name, Id:t_id, Longitude:t_longitude}
    } else {
        return nil
    }
}

func build_hypervisor_azure_msghypervisormonitoringpolicy(d *schema.ResourceData, r []interface{}) *handler.MsghypervisorMonitoringPolicy {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_isenabled *bool
        if val, ok := tmp["isenabled"].(bool); ok {
            t_isenabled = new(bool)
            t_isenabled = &val
        }
        var t_name *string
        if val, ok := tmp["name"].(string); ok {
            t_name = new(string)
            t_name = &val
        }
        var t_id *int
        if val, ok := tmp["id"].(int); ok {
            t_id = new(int)
            t_id = &val
        }
        return &handler.MsghypervisorMonitoringPolicy{IsEnabled:t_isenabled, Name:t_name, Id:t_id}
    } else {
        return nil
    }
}

func build_hypervisor_azure_msgactivitycontroloptions(d *schema.ResourceData, r []interface{}) *handler.MsgActivityControlOptions {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_restoreactivitycontroloptions *handler.MsgbackupActivityControlOptionsProp
        if val, ok := tmp["restoreactivitycontroloptions"].([]interface{}); ok {
            t_restoreactivitycontroloptions = build_hypervisor_azure_msgbackupactivitycontroloptionsprop(d, val)
        }
        var t_backupactivitycontroloptions *handler.MsgbackupActivityControlOptionsProp
        if val, ok := tmp["backupactivitycontroloptions"].([]interface{}); ok {
            t_backupactivitycontroloptions = build_hypervisor_azure_msgbackupactivitycontroloptionsprop(d, val)
        }
        var t_enablebackup *bool
        if val, ok := tmp["enablebackup"].(bool); ok {
            t_enablebackup = new(bool)
            t_enablebackup = &val
        }
        var t_enablerestore *bool
        if val, ok := tmp["enablerestore"].(bool); ok {
            t_enablerestore = new(bool)
            t_enablerestore = &val
        }
        return &handler.MsgActivityControlOptions{RestoreActivityControlOptions:t_restoreactivitycontroloptions, BackupActivityControlOptions:t_backupactivitycontroloptions, EnableBackup:t_enablebackup, EnableRestore:t_enablerestore}
    } else {
        return nil
    }
}

func build_hypervisor_azure_msgbackupactivitycontroloptionsprop(d *schema.ResourceData, r []interface{}) *handler.MsgbackupActivityControlOptionsProp {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_timezone *handler.MsgIdName
        if val, ok := tmp["timezone"].([]interface{}); ok {
            t_timezone = build_hypervisor_azure_msgidname(d, val)
        }
        var t_enableafterdelay *bool
        if val, ok := tmp["enableafterdelay"].(bool); ok {
            t_enableafterdelay = new(bool)
            t_enableafterdelay = &val
        }
        var t_delaytime *string
        if val, ok := tmp["delaytime"].(string); ok {
            t_delaytime = new(string)
            t_delaytime = &val
        }
        return &handler.MsgbackupActivityControlOptionsProp{TimeZone:t_timezone, EnableAfterDelay:t_enableafterdelay, DelayTime:t_delaytime}
    } else {
        return nil
    }
}

func serialize_hypervisor_azure_msghypervisorsettings(data *handler.MsghypervisorSettings) map[string]interface{} {
    val := make(map[string]interface{})
    if data.MetricsMonitoringPolicy != nil {
        val["metricsmonitoringpolicy"] = serialize_hypervisor_azure_msghypervisormonitoringpolicy(data.MetricsMonitoringPolicy)
    }
    if data.ApplicationCredentials != nil {
        val["applicationcredentials"] = serialize_hypervisor_azure_msgidname(data.ApplicationCredentials)
    }
    if data.GuestCredentials != nil {
        val["guestcredentials"] = serialize_hypervisor_azure_msgidname(data.GuestCredentials)
    }
    if data.MountAccessNode != nil {
        val["mountaccessnode"] = serialize_hypervisor_azure_msgidname(data.MountAccessNode)
    }
    if data.RegionInfo != nil {
        val["regioninfo"] = serialize_hypervisor_azure_msgregioninfo(data.RegionInfo)
    }
    if data.TimeZone != nil {
        val["timezone"] = serialize_hypervisor_azure_msgidname(data.TimeZone)
    }
    if data.CustomAttributes != nil {
        val["customattributes"] = serialize_hypervisor_azure_msghypervisorcustomattribute(data.CustomAttributes)
    }
    return val
}

func serialize_hypervisor_azure_msghypervisorcustomattribute(data *handler.MsghypervisorCustomAttribute) map[string]interface{} {
    val := make(map[string]interface{})
    if data.Type != nil {
        val["type"] = data.Type
    }
    if data.Value != nil {
        val["value"] = data.Value
    }
    return val
}

func serialize_hypervisor_azure_msgidname(data *handler.MsgIdName) map[string]interface{} {
    val := make(map[string]interface{})
    if data.Name != nil {
        val["name"] = data.Name
    }
    if data.Id != nil {
        val["id"] = data.Id
    }
    return val
}

func serialize_hypervisor_azure_msgregioninfo(data *handler.MsgRegionInfo) map[string]interface{} {
    val := make(map[string]interface{})
    if data.DisplayName != nil {
        val["displayname"] = data.DisplayName
    }
    if data.Latitude != nil {
        val["latitude"] = data.Latitude
    }
    if data.Name != nil {
        val["name"] = data.Name
    }
    if data.Id != nil {
        val["id"] = data.Id
    }
    if data.Longitude != nil {
        val["longitude"] = data.Longitude
    }
    return val
}

func serialize_hypervisor_azure_msghypervisormonitoringpolicy(data *handler.MsghypervisorMonitoringPolicy) map[string]interface{} {
    val := make(map[string]interface{})
    if data.IsEnabled != nil {
        val["isenabled"] = data.IsEnabled
    }
    if data.Name != nil {
        val["name"] = data.Name
    }
    if data.Id != nil {
        val["id"] = data.Id
    }
    return val
}

func serialize_hypervisor_azure_msgactivitycontroloptions(data *handler.MsgActivityControlOptions) map[string]interface{} {
    val := make(map[string]interface{})
    if data.RestoreActivityControlOptions != nil {
        val["restoreactivitycontroloptions"] = serialize_hypervisor_azure_msgbackupactivitycontroloptionsprop(data.RestoreActivityControlOptions)
    }
    if data.BackupActivityControlOptions != nil {
        val["backupactivitycontroloptions"] = serialize_hypervisor_azure_msgbackupactivitycontroloptionsprop(data.BackupActivityControlOptions)
    }
    if data.EnableBackup != nil {
        val["enablebackup"] = data.EnableBackup
    }
    if data.EnableRestore != nil {
        val["enablerestore"] = data.EnableRestore
    }
    return val
}

func serialize_hypervisor_azure_msgbackupactivitycontroloptionsprop(data *handler.MsgbackupActivityControlOptionsProp) map[string]interface{} {
    val := make(map[string]interface{})
    if data.TimeZone != nil {
        val["timezone"] = serialize_hypervisor_azure_msgidname(data.TimeZone)
    }
    if data.EnableAfterDelay != nil {
        val["enableafterdelay"] = data.EnableAfterDelay
    }
    if data.DelayTime != nil {
        val["delaytime"] = data.DelayTime
    }
    return val
}
