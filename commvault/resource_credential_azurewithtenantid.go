package commvault

import (
    "fmt"
    "strconv"
    "strings"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCredential_AzureWithTenantId() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateCredential_AzureWithTenantId,
        Read:   resourceReadCredential_AzureWithTenantId,
        Update: resourceUpdateCredential_AzureWithTenantId,
        Delete: resourceDeleteCredential_AzureWithTenantId,

        Schema: map[string]*schema.Schema{
            "vendortype": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Cloud vendor types appilcable only for Cloud Account type [ALICLOUD_OSS, AMAZON_GLACIER, AMAZON, ATT_SYNAPTIC, REVERA_VAULT, CEPH_OBJECT_GATEWAY_S3, CMCC_ONEST, CLOUDIAN_HYPERSTORE, DELL_EMC_ECS_S3, EMC_ATMOS, FUJITSU_STORAGE_ETERNUS, GOOGLE_CLOUD, HDS_HCP, HITACHI_VANTARA_HCP_S3, HUAWEI_OSS, IBM_CLOUD, IBM_CLOUD_S3, INSPUR_CLOUD, IRON_MOUNTAIN_CLOUD, KINGSOFT_KS3, MICROSOFT_AZURE_TYPE, NETAPP_STORAGEGRID, NUTANIX_BUCKETS, OPENSTACK, AMPLIDATA, RACKSPACE_CLOUD_FILES, S3_COMPATIBLE, SALESFORCE_CONNECTED_APP, SCALITY_RING, TELEFONICA_OPEN_CLOUD_OBJECT_STORAGE, VERIZON_CLOUD, WASABI_HOT_CLOUD_STORAGE]",
            },
            "accounttype": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "[WINDOWSACCOUNT, LINUXACCOUNT, STORAGE_ARRAY_ACCOUNT, CLOUD_ACCOUNT]",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Name of Credential",
            },
            "environment": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Azure cloud deployed region [AZURE_CLOUD, AZURE_USGOV, AZURE_GERMANCLOUD, AZURE_CHINACLOUD, AZURE_STACK]",
            },
            "endpoints": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Endpoints for Authentication, Storage and Management",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "storageendpoint": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "activedirectoryendpoint": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "resourcemanagerendpoint": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
            },
            "tenantid": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Unique Azure active directory ID",
            },
            "applicationid": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Unique Azure application ID",
            },
            "applicationsecret": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Application secret of Credential and must be in base64 encoded format.",
            },
            "description": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Description of Credential",
            },
        },
    }
}

func resourceCreateCredential_AzureWithTenantId(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/Credential
    var response_id = strconv.Itoa(0)
    var t_vendortype *string
    if val, ok := d.GetOk("vendortype"); ok {
        t_vendortype = handler.ToStringValue(val, false)
    }
    var t_accounttype *string
    if val, ok := d.GetOk("accounttype"); ok {
        t_accounttype = handler.ToStringValue(val, false)
    }
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_environment *string
    if val, ok := d.GetOk("environment"); ok {
        t_environment = handler.ToStringValue(val, false)
    }
    var t_endpoints *handler.MsgAzureEndpoints
    if val, ok := d.GetOk("endpoints"); ok {
        t_endpoints = build_credential_azurewithtenantid_msgazureendpoints(d, val.([]interface{}))
    }
    var t_tenantid *string
    if val, ok := d.GetOk("tenantid"); ok {
        t_tenantid = handler.ToStringValue(val, false)
    }
    var t_authtype *string
    var c_authtype string = "AZUREACCOUNT"
    t_authtype = &c_authtype
    var t_applicationid *string
    if val, ok := d.GetOk("applicationid"); ok {
        t_applicationid = handler.ToStringValue(val, false)
    }
    var t_applicationsecret *string
    if val, ok := d.GetOk("applicationsecret"); ok {
        t_applicationsecret = handler.ToStringValue(val, false)
    }
    var t_description *string
    if val, ok := d.GetOk("description"); ok {
        t_description = handler.ToStringValue(val, false)
    }
    var req = handler.MsgCreateCredentialAzureWithTenantIdRequest{VendorType:t_vendortype, AccountType:t_accounttype, Name:t_name, Environment:t_environment, Endpoints:t_endpoints, TenantId:t_tenantid, AuthType:t_authtype, ApplicationId:t_applicationid, ApplicationSecret:t_applicationsecret, Description:t_description}
    h_err := handler.ConfigureCredential_AzureWithTenantId(&req, d, m)
    if h_err != nil {
        return fmt.Errorf("operation [CreateCredentialAzureWithTenantId] failed, Error %s", h_err)
    }
    resp, err := handler.CvCreateCredentialAzureWithTenantId(req)
    if err != nil {
        return fmt.Errorf("operation [CreateCredentialAzureWithTenantId] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateCredentialAzureWithTenantId] failed")
    } else {
        d.SetId(response_id)
        return resourceReadCredential_AzureWithTenantId(d, m)
    }
}

func resourceReadCredential_AzureWithTenantId(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V5/Credential/{credentialId}
    resp, err := handler.CvGetCredentialDetailsAzureWithTenantId(d.Id())
    if err != nil {
        if strings.Contains(err.Error(), "status: 404") {
            handler.LogEntry("debug", "entity not present, removing from state")
            d.SetId("")
            return nil
        }
        return fmt.Errorf("operation [GetCredentialDetailsAzureWithTenantId] failed, Error %s", err)
    }
    if resp.VendorType != nil {
        d.Set("vendortype", resp.VendorType)
    }
    if resp.AccountType != nil {
        d.Set("accounttype", resp.AccountType)
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    if resp.Environment != nil {
        d.Set("environment", resp.Environment)
    }
    if rtn, ok := serialize_credential_azurewithtenantid_msgazureendpoints(d, resp.Endpoints); ok {
        d.Set("endpoints", rtn)
    } else {
        d.Set("endpoints", make([]map[string]interface{}, 0))
    }
    if resp.TenantId != nil {
        d.Set("tenantid", resp.TenantId)
    }
    if resp.ApplicationId != nil {
        d.Set("applicationid", resp.ApplicationId)
    }
    if resp.ApplicationSecret != nil {
        d.Set("applicationsecret", resp.ApplicationSecret)
    }
    if resp.Description != nil {
        d.Set("description", resp.Description)
    }
    return nil
}

func resourceUpdateCredential_AzureWithTenantId(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V5/Credential/{credentialId}
    var t_environment *string
    if val, ok := d.GetOk("environment"); ok {
        t_environment = handler.ToStringValue(val, false)
    }
    var t_endpoints *handler.MsgAzureEndpoints
    if val, ok := d.GetOk("endpoints"); ok {
        t_endpoints = build_credential_azurewithtenantid_msgazureendpoints(d, val.([]interface{}))
    }
    var t_newname *string
    if val, ok := d.GetOk("name"); ok {
        t_newname = handler.ToStringValue(val, false)
    }
    var t_newapplicationsecret *string
    if val, ok := d.GetOk("applicationsecret"); ok {
        t_newapplicationsecret = handler.ToStringValue(val, false)
    }
    var t_tenantid *string
    if val, ok := d.GetOk("tenantid"); ok {
        t_tenantid = handler.ToStringValue(val, false)
    }
    var t_description *string
    if val, ok := d.GetOk("description"); ok {
        t_description = handler.ToStringValue(val, false)
    }
    var t_authtype *string
    var c_authtype string = "AZUREACCOUNT"
    t_authtype = &c_authtype
    var t_applicationid *string
    if val, ok := d.GetOk("applicationid"); ok {
        t_applicationid = handler.ToStringValue(val, false)
    }
    var req = handler.MsgUpdateCredentialAzureWithTenantIdRequest{Environment:t_environment, Endpoints:t_endpoints, NewName:t_newname, NewApplicationSecret:t_newapplicationsecret, TenantId:t_tenantid, Description:t_description, AuthType:t_authtype, ApplicationId:t_applicationid}
    h_err := handler.UpdateCredential_AzureWithTenantId(&req, d, m)
    if h_err != nil {
        return fmt.Errorf("operation [UpdateCredentialAzureWithTenantId] failed, Error %s", h_err)
    }
    _, err := handler.CvUpdateCredentialAzureWithTenantId(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [UpdateCredentialAzureWithTenantId] failed, Error %s", err)
    }
    return resourceReadCredential_AzureWithTenantId(d, m)
}

func resourceDeleteCredential_AzureWithTenantId(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V5/Credential/{credentialId}
    _, err := handler.CvDeleteCredential(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteCredential] failed, Error %s", err)
    }
    return nil
}

func build_credential_azurewithtenantid_msgazureendpoints(d *schema.ResourceData, r []interface{}) *handler.MsgAzureEndpoints {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_storageendpoint *string
        if val, ok := tmp["storageendpoint"]; ok {
            t_storageendpoint = handler.ToStringValue(val, true)
        }
        var t_activedirectoryendpoint *string
        if val, ok := tmp["activedirectoryendpoint"]; ok {
            t_activedirectoryendpoint = handler.ToStringValue(val, true)
        }
        var t_resourcemanagerendpoint *string
        if val, ok := tmp["resourcemanagerendpoint"]; ok {
            t_resourcemanagerendpoint = handler.ToStringValue(val, true)
        }
        return &handler.MsgAzureEndpoints{StorageEndpoint:t_storageendpoint, ActiveDirectoryEndpoint:t_activedirectoryendpoint, ResourceManagerEndpoint:t_resourcemanagerendpoint}
    } else {
        return nil
    }
}

func serialize_credential_azurewithtenantid_msgazureendpoints(d *schema.ResourceData, data *handler.MsgAzureEndpoints) ([]map[string]interface{}, bool) {
    //MsgAzureEndpoints
    //MsgAzureEndpoints
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.StorageEndpoint != nil {
        val[0]["storageendpoint"] = data.StorageEndpoint
        added = true
    }
    if data.ActiveDirectoryEndpoint != nil {
        val[0]["activedirectoryendpoint"] = data.ActiveDirectoryEndpoint
        added = true
    }
    if data.ResourceManagerEndpoint != nil {
        val[0]["resourcemanagerendpoint"] = data.ResourceManagerEndpoint
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}
