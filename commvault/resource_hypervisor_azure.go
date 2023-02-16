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
            "skipcredentialvalidation": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "if credential validation has to be skipped.",
            },
            "etcdprotection": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Flag to create an application group etcd (system generated) with pre-defined content",
            },
            "credentials": {
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
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "The name of the hypervisor group being created",
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
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: "Type of access node , Ex: 3 - access Node , 28 - Access Node Groups",
                        },
                    },
                },
            },
            "applicationpassword": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Application Password of Azure login Application",
            },
            "tenantid": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Tenant id of Azure login Application",
            },
            "workloadregion": {
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
            "subscriptionid": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "subscription id of Azure ",
            },
            "applicationid": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Application id of Azure login Application",
            },
            "usemanagedidentity": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "",
            },
            "activitycontrol": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "restoreactivitycontroloptions": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "delaytime": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
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
                                                "time": {
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "delay time in unix timestamp",
                                                },
                                                "value": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "actual delay time value in string format according to the timezone",
                                                },
                                            },
                                        },
                                    },
                                    "activitytype": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "denotes the activity type being considered [BACKUP, RESTORE, ONLINECI, ARCHIVEPRUNE]",
                                    },
                                    "enableafteradelay": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if the activity will be enabled after a delay time interval",
                                    },
                                    "enableactivitytype": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if the activity type is enabled",
                                    },
                                },
                            },
                        },
                        "backupactivitycontroloptions": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "delaytime": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
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
                                                "time": {
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "delay time in unix timestamp",
                                                },
                                                "value": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "actual delay time value in string format according to the timezone",
                                                },
                                            },
                                        },
                                    },
                                    "activitytype": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "denotes the activity type being considered [BACKUP, RESTORE, ONLINECI, ARCHIVEPRUNE]",
                                    },
                                    "enableafteradelay": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if the activity will be enabled after a delay time interval",
                                    },
                                    "enableactivitytype": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if the activity type is enabled",
                                    },
                                },
                            },
                        },
                        "enablebackup": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "true if Backup is enabled",
                        },
                        "enablerestore": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "true if Restore is enabled",
                        },
                    },
                },
            },
            "settings": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "metricsmonitoringpolicy": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "isenabled": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if Metrics Monioring policy is enabled",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Metrics Moitoring Policy Name",
                                    },
                                    "id": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Metrics Moitoring PolicyId",
                                    },
                                },
                            },
                        },
                        "applicationcredentials": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "password": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Sensitive:    true,
                                        Description: "password to access the network path",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "username to access the network path",
                                    },
                                },
                            },
                        },
                        "guestcredentials": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "password": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Sensitive:    true,
                                        Description: "password to access the network path",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "username to access the network path",
                                    },
                                },
                            },
                        },
                        "mountaccessnode": {
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
                        "regioninfo": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "displayname": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Display Name of Region",
                                    },
                                    "latitude": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Geolocation Latitude",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Region Name",
                                    },
                                    "id": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Region Id",
                                    },
                                    "longitude": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Geolocation Longitude",
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
                        "customattributes": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "type": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "client custom attribute type . Ex- 3 - For client 8- For clientGroup",
                                    },
                                    "value": {
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
            "security": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "clientowners": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Client owners for the Hypervisor",
                        },
                        "associatedusergroups": {
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
            "displayname": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "The name of the hypervisor that has to be changed",
            },
            "fbrunixmediaagent": {
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
            "password": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Application Password of Azure login Application",
            },
            "servername": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Client Name to Update",
            },
            "hypervisortype": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "[Azure_V2]",
            },
            "username": {
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
    if val, ok := d.GetOk("skipcredentialvalidation"); ok {
        t_skipcredentialvalidation = handler.ToBooleanValue(val, false)
    }
    var t_etcdprotection *bool
    if val, ok := d.GetOk("etcdprotection"); ok {
        t_etcdprotection = handler.ToBooleanValue(val, false)
    }
    var t_credentials *handler.MsgIdName
    if val, ok := d.GetOk("credentials"); ok {
        t_credentials = build_hypervisor_azure_msgidname(d, val.([]interface{}))
    }
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_accessnodes []handler.MsgaccessNodeModelSet
    if val, ok := d.GetOk("accessnodes"); ok {
        t_accessnodes = build_hypervisor_azure_msgaccessnodemodelset_array(d, val.(*schema.Set).List())
    }
    var t_applicationpassword *string
    if val, ok := d.GetOk("applicationpassword"); ok {
        t_applicationpassword = handler.ToStringValue(val, false)
    }
    var t_tenantid *string
    if val, ok := d.GetOk("tenantid"); ok {
        t_tenantid = handler.ToStringValue(val, false)
    }
    var t_hypervisortype *string
    var c_hypervisortype string = "AZURE_V2"
    t_hypervisortype = &c_hypervisortype
    var t_workloadregion *handler.MsgIdName
    if val, ok := d.GetOk("workloadregion"); ok {
        t_workloadregion = build_hypervisor_azure_msgidname(d, val.([]interface{}))
    }
    var t_subscriptionid *string
    if val, ok := d.GetOk("subscriptionid"); ok {
        t_subscriptionid = handler.ToStringValue(val, false)
    }
    var t_applicationid *string
    if val, ok := d.GetOk("applicationid"); ok {
        t_applicationid = handler.ToStringValue(val, false)
    }
    var t_usemanagedidentity *bool
    if val, ok := d.GetOk("usemanagedidentity"); ok {
        t_usemanagedidentity = handler.ToBooleanValue(val, false)
    }
    var req = handler.MsgCreateHypervisorAzureRequest{SkipCredentialValidation:t_skipcredentialvalidation, EtcdProtection:t_etcdprotection, Credentials:t_credentials, Name:t_name, AccessNodes:t_accessnodes, ApplicationPassword:t_applicationpassword, TenantId:t_tenantid, HypervisorType:t_hypervisortype, WorkloadRegion:t_workloadregion, SubscriptionId:t_subscriptionid, ApplicationId:t_applicationid, UseManagedIdentity:t_usemanagedidentity}
    resp, err := handler.CvCreateHypervisorAzure(req)
    if err != nil {
        return fmt.Errorf("operation [CreateHypervisorAzure] failed, Error %s", err)
    }
    if resp.Response != nil {
        if resp.Response.HypervisorId != nil {
            response_id = strconv.Itoa(*resp.Response.HypervisorId)
        }
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateHypervisorAzure] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateHypervisor_Azure(d, m)
    }
}

func resourceReadHypervisor_Azure(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/Hypervisor/{hypervisorId}
    resp, err := handler.CvGetHypervisors(d.Id())
    if err != nil {
        return fmt.Errorf("operation [GetHypervisors] failed, Error %s", err)
    }
    if rtn, ok := serialize_hypervisor_azure_msgactivitycontroloptions(d, resp.ActivityControl); ok {
        d.Set("activitycontrol", rtn)
    } else {
        d.Set("activitycontrol", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_hypervisor_azure_msghypervisorsettings(d, resp.Settings); ok {
        d.Set("settings", rtn)
    } else {
        d.Set("settings", make([]map[string]interface{}, 0))
    }
    if rtn, ok := handler.GetAccessNodes(d, resp.AccessNodeList); ok {
        d.Set("accessnodes", rtn)
    } else {
        d.Set("accessnodes", make([]map[string]interface{}, 0))
    }
    if resp.DisplayName != nil {
        d.Set("displayname", resp.DisplayName)
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
        val := d.Get("activitycontrol")
        t_activitycontrol = build_hypervisor_azure_msgactivitycontroloptions(d, val.([]interface{}))
    }
    var t_settings *handler.MsghypervisorSettings
    if d.HasChange("settings") {
        val := d.Get("settings")
        t_settings = build_hypervisor_azure_msghypervisorsettings(d, val.([]interface{}))
    }
    var t_security *handler.MsgVMHypervisorSecurityProp
    if d.HasChange("security") {
        val := d.Get("security")
        t_security = build_hypervisor_azure_msgvmhypervisorsecurityprop(d, val.([]interface{}))
    }
    var t_newname *string
    if d.HasChange("displayname") {
        val := d.Get("displayname")
        t_newname = handler.ToStringValue(val, false)
    }
    var t_skipcredentialvalidation *bool
    if d.HasChange("skipcredentialvalidation") {
        val := d.Get("skipcredentialvalidation")
        t_skipcredentialvalidation = handler.ToBooleanValue(val, false)
    }
    var t_credentials *handler.MsgIdName
    if d.HasChange("credentials") {
        val := d.Get("credentials")
        t_credentials = build_hypervisor_azure_msgidname(d, val.([]interface{}))
    }
    var t_accessnodes []handler.MsgaccessNodeModelSet
    if d.HasChange("accessnodes") {
        val := d.Get("accessnodes")
        t_accessnodes = build_hypervisor_azure_msgaccessnodemodelset_array(d, val.(*schema.Set).List())
    }
    var t_fbrunixmediaagent *handler.MsgIdName
    if d.HasChange("fbrunixmediaagent") {
        val := d.Get("fbrunixmediaagent")
        t_fbrunixmediaagent = build_hypervisor_azure_msgidname(d, val.([]interface{}))
    }
    var t_password *string
    if d.HasChange("password") {
        val := d.Get("password")
        t_password = handler.ToStringValue(val, false)
    }
    var t_tenantid *string
    if d.HasChange("tenantid") {
        val := d.Get("tenantid")
        t_tenantid = handler.ToStringValue(val, false)
    }
    var t_servername *string
    if d.HasChange("servername") {
        val := d.Get("servername")
        t_servername = handler.ToStringValue(val, false)
    }
    var t_hypervisortype *string
    if d.HasChange("hypervisortype") {
        val := d.Get("hypervisortype")
        t_hypervisortype = handler.ToStringValue(val, false)
    }
    var t_subscriptionid *string
    if d.HasChange("subscriptionid") {
        val := d.Get("subscriptionid")
        t_subscriptionid = handler.ToStringValue(val, false)
    }
    var t_username *string
    if d.HasChange("username") {
        val := d.Get("username")
        t_username = handler.ToStringValue(val, false)
    }
    var t_usemanagedidentity *bool
    if d.HasChange("usemanagedidentity") {
        val := d.Get("usemanagedidentity")
        t_usemanagedidentity = handler.ToBooleanValue(val, false)
    }
    var req = handler.MsgupdateHypervisorAzureRequest{ActivityControl:t_activitycontrol, Settings:t_settings, Security:t_security, NewName:t_newname, SkipCredentialValidation:t_skipcredentialvalidation, Credentials:t_credentials, AccessNodes:t_accessnodes, FbrUnixMediaAgent:t_fbrunixmediaagent, Password:t_password, TenantId:t_tenantid, ServerName:t_servername, HypervisorType:t_hypervisortype, SubscriptionId:t_subscriptionid, UserName:t_username, UseManagedIdentity:t_usemanagedidentity}
    _, err := handler.CvupdateHypervisorAzure(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [updateHypervisorAzure] failed, Error %s", err)
    }
    return resourceReadHypervisor_Azure(d, m)
}

func resourceCreateUpdateHypervisor_Azure(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Hypervisor/{hypervisorId}
    var execUpdate bool = false
    var t_activitycontrol *handler.MsgActivityControlOptions
    if val, ok := d.GetOk("activitycontrol"); ok {
        t_activitycontrol = build_hypervisor_azure_msgactivitycontroloptions(d, val.([]interface{}))
        execUpdate = true
    }
    var t_settings *handler.MsghypervisorSettings
    if val, ok := d.GetOk("settings"); ok {
        t_settings = build_hypervisor_azure_msghypervisorsettings(d, val.([]interface{}))
        execUpdate = true
    }
    var t_security *handler.MsgVMHypervisorSecurityProp
    if val, ok := d.GetOk("security"); ok {
        t_security = build_hypervisor_azure_msgvmhypervisorsecurityprop(d, val.([]interface{}))
        execUpdate = true
    }
    var t_newname *string
    if val, ok := d.GetOk("displayname"); ok {
        t_newname = handler.ToStringValue(val, false)
        execUpdate = true
    }
    var t_fbrunixmediaagent *handler.MsgIdName
    if val, ok := d.GetOk("fbrunixmediaagent"); ok {
        t_fbrunixmediaagent = build_hypervisor_azure_msgidname(d, val.([]interface{}))
        execUpdate = true
    }
    var t_password *string
    if val, ok := d.GetOk("password"); ok {
        t_password = handler.ToStringValue(val, false)
        execUpdate = true
    }
    var t_servername *string
    if val, ok := d.GetOk("servername"); ok {
        t_servername = handler.ToStringValue(val, false)
        execUpdate = true
    }
    var t_username *string
    if val, ok := d.GetOk("username"); ok {
        t_username = handler.ToStringValue(val, false)
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgupdateHypervisorAzureRequest{ActivityControl:t_activitycontrol, Settings:t_settings, Security:t_security, NewName:t_newname, FbrUnixMediaAgent:t_fbrunixmediaagent, Password:t_password, ServerName:t_servername, UserName:t_username}
        _, err := handler.CvupdateHypervisorAzure(req, d.Id())
        if err != nil {
            return fmt.Errorf("operation [updateHypervisorAzure] failed, Error %s", err)
        }
    }
    return resourceReadHypervisor_Azure(d, m)
}

func resourceDeleteHypervisor_Azure(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/Hypervisor/{hypervisorId}
    _, err := handler.CvDeleteHypervisor(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteHypervisor] failed, Error %s", err)
    }
    return nil
}

func build_hypervisor_azure_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_hypervisor_azure_msgvmhypervisorsecurityprop(d *schema.ResourceData, r []interface{}) *handler.MsgVMHypervisorSecurityProp {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_clientowners *string
        if val, ok := tmp["clientowners"]; ok {
            t_clientowners = handler.ToStringValue(val, true)
        }
        var t_associatedusergroups []handler.MsgIdNameSet
        if val, ok := tmp["associatedusergroups"]; ok {
            t_associatedusergroups = build_hypervisor_azure_msgidnameset_array(d, val.(*schema.Set).List())
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

func build_hypervisor_azure_msghypervisorsettings(d *schema.ResourceData, r []interface{}) *handler.MsghypervisorSettings {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_metricsmonitoringpolicy *handler.MsghypervisorMonitoringPolicy
        if val, ok := tmp["metricsmonitoringpolicy"]; ok {
            t_metricsmonitoringpolicy = build_hypervisor_azure_msghypervisormonitoringpolicy(d, val.([]interface{}))
        }
        var t_applicationcredentials *handler.MsgUserNamePassword
        if val, ok := tmp["applicationcredentials"]; ok {
            t_applicationcredentials = build_hypervisor_azure_msgusernamepassword(d, val.([]interface{}))
        }
        var t_guestcredentials *handler.MsgUserNamePassword
        if val, ok := tmp["guestcredentials"]; ok {
            t_guestcredentials = build_hypervisor_azure_msgusernamepassword(d, val.([]interface{}))
        }
        var t_mountaccessnode *handler.MsgIdName
        if val, ok := tmp["mountaccessnode"]; ok {
            t_mountaccessnode = build_hypervisor_azure_msgidname(d, val.([]interface{}))
        }
        var t_regioninfo *handler.MsgRegionInfo
        if val, ok := tmp["regioninfo"]; ok {
            t_regioninfo = build_hypervisor_azure_msgregioninfo(d, val.([]interface{}))
        }
        var t_timezone *handler.MsgIdName
        if val, ok := tmp["timezone"]; ok {
            t_timezone = build_hypervisor_azure_msgidname(d, val.([]interface{}))
        }
        var t_customattributes *handler.MsghypervisorCustomAttribute
        if val, ok := tmp["customattributes"]; ok {
            t_customattributes = build_hypervisor_azure_msghypervisorcustomattribute(d, val.([]interface{}))
        }
        return &handler.MsghypervisorSettings{MetricsMonitoringPolicy:t_metricsmonitoringpolicy, ApplicationCredentials:t_applicationcredentials, GuestCredentials:t_guestcredentials, MountAccessNode:t_mountaccessnode, RegionInfo:t_regioninfo, TimeZone:t_timezone, CustomAttributes:t_customattributes}
    } else {
        return nil
    }
}

func build_hypervisor_azure_msghypervisorcustomattribute(d *schema.ResourceData, r []interface{}) *handler.MsghypervisorCustomAttribute {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_type *int
        if val, ok := tmp["type"]; ok {
            t_type = handler.ToIntValue(val, true)
        }
        var t_value *string
        if val, ok := tmp["value"]; ok {
            t_value = handler.ToStringValue(val, true)
        }
        return &handler.MsghypervisorCustomAttribute{Type:t_type, Value:t_value}
    } else {
        return nil
    }
}

func build_hypervisor_azure_msgregioninfo(d *schema.ResourceData, r []interface{}) *handler.MsgRegionInfo {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_displayname *string
        if val, ok := tmp["displayname"]; ok {
            t_displayname = handler.ToStringValue(val, true)
        }
        var t_latitude *string
        if val, ok := tmp["latitude"]; ok {
            t_latitude = handler.ToStringValue(val, true)
        }
        var t_name *string
        if val, ok := tmp["name"]; ok {
            t_name = handler.ToStringValue(val, true)
        }
        var t_id *int
        if val, ok := tmp["id"]; ok {
            t_id = handler.ToIntValue(val, true)
        }
        var t_longitude *string
        if val, ok := tmp["longitude"]; ok {
            t_longitude = handler.ToStringValue(val, true)
        }
        return &handler.MsgRegionInfo{DisplayName:t_displayname, Latitude:t_latitude, Name:t_name, Id:t_id, Longitude:t_longitude}
    } else {
        return nil
    }
}

func build_hypervisor_azure_msgusernamepassword(d *schema.ResourceData, r []interface{}) *handler.MsgUserNamePassword {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_password *string
        if val, ok := tmp["password"]; ok {
            t_password = handler.ToStringValue(val, true)
        }
        var t_name *string
        if val, ok := tmp["name"]; ok {
            t_name = handler.ToStringValue(val, true)
        }
        return &handler.MsgUserNamePassword{Password:t_password, Name:t_name}
    } else {
        return nil
    }
}

func build_hypervisor_azure_msghypervisormonitoringpolicy(d *schema.ResourceData, r []interface{}) *handler.MsghypervisorMonitoringPolicy {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_isenabled *bool
        if val, ok := tmp["isenabled"]; ok {
            t_isenabled = handler.ToBooleanValue(val, true)
        }
        var t_name *string
        if val, ok := tmp["name"]; ok {
            t_name = handler.ToStringValue(val, true)
        }
        var t_id *int
        if val, ok := tmp["id"]; ok {
            t_id = handler.ToIntValue(val, true)
        }
        return &handler.MsghypervisorMonitoringPolicy{IsEnabled:t_isenabled, Name:t_name, Id:t_id}
    } else {
        return nil
    }
}

func build_hypervisor_azure_msgactivitycontroloptions(d *schema.ResourceData, r []interface{}) *handler.MsgActivityControlOptions {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_restoreactivitycontroloptions *handler.MsgActivityControlOptionsProp
        if val, ok := tmp["restoreactivitycontroloptions"]; ok {
            t_restoreactivitycontroloptions = build_hypervisor_azure_msgactivitycontroloptionsprop(d, val.([]interface{}))
        }
        var t_backupactivitycontroloptions *handler.MsgActivityControlOptionsProp
        if val, ok := tmp["backupactivitycontroloptions"]; ok {
            t_backupactivitycontroloptions = build_hypervisor_azure_msgactivitycontroloptionsprop(d, val.([]interface{}))
        }
        var t_enablebackup *bool
        if val, ok := tmp["enablebackup"]; ok {
            t_enablebackup = handler.ToBooleanValue(val, true)
        }
        var t_enablerestore *bool
        if val, ok := tmp["enablerestore"]; ok {
            t_enablerestore = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgActivityControlOptions{RestoreActivityControlOptions:t_restoreactivitycontroloptions, BackupActivityControlOptions:t_backupactivitycontroloptions, EnableBackup:t_enablebackup, EnableRestore:t_enablerestore}
    } else {
        return nil
    }
}

func build_hypervisor_azure_msgactivitycontroloptionsprop(d *schema.ResourceData, r []interface{}) *handler.MsgActivityControlOptionsProp {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_delaytime *handler.MsgActivityControlTileDelayTime
        if val, ok := tmp["delaytime"]; ok {
            t_delaytime = build_hypervisor_azure_msgactivitycontroltiledelaytime(d, val.([]interface{}))
        }
        var t_activitytype *string
        if val, ok := tmp["activitytype"]; ok {
            t_activitytype = handler.ToStringValue(val, true)
        }
        var t_enableafteradelay *bool
        if val, ok := tmp["enableafteradelay"]; ok {
            t_enableafteradelay = handler.ToBooleanValue(val, true)
        }
        var t_enableactivitytype *bool
        if val, ok := tmp["enableactivitytype"]; ok {
            t_enableactivitytype = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgActivityControlOptionsProp{DelayTime:t_delaytime, ActivityType:t_activitytype, EnableAfterADelay:t_enableafteradelay, EnableActivityType:t_enableactivitytype}
    } else {
        return nil
    }
}

func build_hypervisor_azure_msgactivitycontroltiledelaytime(d *schema.ResourceData, r []interface{}) *handler.MsgActivityControlTileDelayTime {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_timezone *handler.MsgIdName
        if val, ok := tmp["timezone"]; ok {
            t_timezone = build_hypervisor_azure_msgidname(d, val.([]interface{}))
        }
        var t_time *int
        if val, ok := tmp["time"]; ok {
            t_time = handler.ToIntValue(val, true)
        }
        var t_value *string
        if val, ok := tmp["value"]; ok {
            t_value = handler.ToStringValue(val, true)
        }
        return &handler.MsgActivityControlTileDelayTime{TimeZone:t_timezone, Time:t_time, Value:t_value}
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
            if val, ok := raw_a["id"]; ok {
                t_id = handler.ToIntValue(val, true)
            }
            var t_type *int
            if val, ok := raw_a["type"]; ok {
                t_type = handler.ToIntValue(val, true)
            }
            tmp[a] = handler.MsgaccessNodeModelSet{Id:t_id, Type:t_type}
        }
        return tmp
    } else {
        return nil
    }
}

func serialize_hypervisor_azure_msghypervisorsettings(d *schema.ResourceData, data *handler.MsghypervisorSettings) ([]map[string]interface{}, bool) {
    //MsghypervisorSettings
    //MsghypervisorSettings
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_hypervisor_azure_msghypervisormonitoringpolicy(d, data.MetricsMonitoringPolicy); ok {
        val[0]["metricsmonitoringpolicy"] = rtn
        added = true
    }
    if rtn, ok := statecopy_hypervisor_azure_settings_applicationcredentials(d); ok {
        val[0]["applicationcredentials"] = rtn
        added = true
    }
    if rtn, ok := statecopy_hypervisor_azure_settings_guestcredentials(d); ok {
        val[0]["guestcredentials"] = rtn
        added = true
    }
    if rtn, ok := serialize_hypervisor_azure_msgidname(d, data.MountAccessNode); ok {
        val[0]["mountaccessnode"] = rtn
        added = true
    }
    if rtn, ok := serialize_hypervisor_azure_msgregioninfo(d, data.RegionInfo); ok {
        val[0]["regioninfo"] = rtn
        added = true
    }
    if rtn, ok := serialize_hypervisor_azure_msgidname(d, data.TimeZone); ok {
        val[0]["timezone"] = rtn
        added = true
    }
    if rtn, ok := serialize_hypervisor_azure_msghypervisorcustomattribute(d, data.CustomAttributes); ok {
        val[0]["customattributes"] = rtn
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_hypervisor_azure_msghypervisorcustomattribute(d *schema.ResourceData, data *handler.MsghypervisorCustomAttribute) ([]map[string]interface{}, bool) {
    //MsghypervisorSettings -> MsghypervisorCustomAttribute
    //MsghypervisorSettings -> MsghypervisorCustomAttribute
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.Type != nil {
        val[0]["type"] = data.Type
        added = true
    }
    if data.Value != nil {
        val[0]["value"] = data.Value
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_hypervisor_azure_msgidname(d *schema.ResourceData, data *handler.MsgIdName) ([]map[string]interface{}, bool) {
    //MsghypervisorSettings -> MsgIdName
    //MsghypervisorSettings -> MsgIdName
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

func serialize_hypervisor_azure_msgregioninfo(d *schema.ResourceData, data *handler.MsgRegionInfo) ([]map[string]interface{}, bool) {
    //MsghypervisorSettings -> MsgRegionInfo
    //MsghypervisorSettings -> MsgRegionInfo
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.DisplayName != nil {
        val[0]["displayname"] = data.DisplayName
        added = true
    }
    if data.Latitude != nil {
        val[0]["latitude"] = data.Latitude
        added = true
    }
    if data.Name != nil {
        val[0]["name"] = data.Name
        added = true
    }
    if data.Id != nil {
        val[0]["id"] = data.Id
        added = true
    }
    if data.Longitude != nil {
        val[0]["longitude"] = data.Longitude
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func statecopy_hypervisor_azure_settings_guestcredentials(d *schema.ResourceData) ([]interface{}, bool) {
    //STATE COPY
    var_a := d.Get("settings").([]interface{})
    if len(var_a) > 0 {
        tmp_a := var_a[0].(map[string]interface{})
        if var_b, ok := tmp_a["guestcredentials"].([]interface{}); ok {
            if len(var_b) > 0 {
                return var_b, true
            }
        }
    }
    return nil, false
}

func statecopy_hypervisor_azure_settings_applicationcredentials(d *schema.ResourceData) ([]interface{}, bool) {
    //STATE COPY
    var_a := d.Get("settings").([]interface{})
    if len(var_a) > 0 {
        tmp_a := var_a[0].(map[string]interface{})
        if var_b, ok := tmp_a["applicationcredentials"].([]interface{}); ok {
            if len(var_b) > 0 {
                return var_b, true
            }
        }
    }
    return nil, false
}

func serialize_hypervisor_azure_msghypervisormonitoringpolicy(d *schema.ResourceData, data *handler.MsghypervisorMonitoringPolicy) ([]map[string]interface{}, bool) {
    //MsghypervisorSettings -> MsghypervisorMonitoringPolicy
    //MsghypervisorSettings -> MsghypervisorMonitoringPolicy
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.IsEnabled != nil {
        val[0]["isenabled"] = strconv.FormatBool(*data.IsEnabled)
        added = true
    }
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

func serialize_hypervisor_azure_msgactivitycontroloptions(d *schema.ResourceData, data *handler.MsgActivityControlOptions) ([]map[string]interface{}, bool) {
    //MsgActivityControlOptions
    //MsgActivityControlOptions
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_hypervisor_azure_msgactivitycontroloptionsprop(d, data.RestoreActivityControlOptions); ok {
        val[0]["restoreactivitycontroloptions"] = rtn
        added = true
    }
    if rtn, ok := serialize_hypervisor_azure_msgactivitycontroloptionsprop(d, data.BackupActivityControlOptions); ok {
        val[0]["backupactivitycontroloptions"] = rtn
        added = true
    }
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

func serialize_hypervisor_azure_msgactivitycontroloptionsprop(d *schema.ResourceData, data *handler.MsgActivityControlOptionsProp) ([]map[string]interface{}, bool) {
    //MsgActivityControlOptions -> MsgActivityControlOptionsProp
    //MsgActivityControlOptions -> MsgActivityControlOptionsProp
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_hypervisor_azure_msgactivitycontroltiledelaytime(d, data.DelayTime); ok {
        val[0]["delaytime"] = rtn
        added = true
    }
    if data.ActivityType != nil {
        val[0]["activitytype"] = data.ActivityType
        added = true
    }
    if data.EnableAfterADelay != nil {
        val[0]["enableafteradelay"] = strconv.FormatBool(*data.EnableAfterADelay)
        added = true
    }
    if data.EnableActivityType != nil {
        val[0]["enableactivitytype"] = strconv.FormatBool(*data.EnableActivityType)
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_hypervisor_azure_msgactivitycontroltiledelaytime(d *schema.ResourceData, data *handler.MsgActivityControlTileDelayTime) ([]map[string]interface{}, bool) {
    //MsgActivityControlOptions -> MsgActivityControlOptionsProp -> MsgActivityControlTileDelayTime
    //MsgActivityControlOptions -> MsgActivityControlOptionsProp -> MsgActivityControlTileDelayTime
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_hypervisor_azure_msgidname(d, data.TimeZone); ok {
        val[0]["timezone"] = rtn
        added = true
    }
    if data.Time != nil {
        val[0]["time"] = data.Time
        added = true
    }
    if data.Value != nil {
        val[0]["value"] = data.Value
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}
