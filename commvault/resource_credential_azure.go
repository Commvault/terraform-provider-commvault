package commvault

import (
    "fmt"
    "strconv"
    "strings"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCredential_Azure() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateCredential_Azure,
        Read:   resourceReadCredential_Azure,
        Update: resourceUpdateCredential_Azure,
        Delete: resourceDeleteCredential_Azure,

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
            "accesskeyid": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Access key ID of Credential, applicable only if authType is Access Secret Key and must be in base64 encoded format.",
            },
            "accountname": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Account name of Credential, applicable only if authType is Access Key",
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

func resourceCreateCredential_Azure(d *schema.ResourceData, m interface{}) error {
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
    var t_accesskeyid *string
    if val, ok := d.GetOk("accesskeyid"); ok {
        t_accesskeyid = handler.ToStringValue(val, false)
    }
    var t_accountname *string
    if val, ok := d.GetOk("accountname"); ok {
        t_accountname = handler.ToStringValue(val, false)
    }
    var t_authtype *string
    var c_authtype string = "MICROSOFT_AZURE"
    t_authtype = &c_authtype
    var t_description *string
    if val, ok := d.GetOk("description"); ok {
        t_description = handler.ToStringValue(val, false)
    }
    var req = handler.MsgCreateCredentialAzureRequest{VendorType:t_vendortype, AccountType:t_accounttype, Name:t_name, AccessKeyId:t_accesskeyid, AccountName:t_accountname, AuthType:t_authtype, Description:t_description}
    h_err := handler.ConfigureCredential_Azure(&req, d, m)
    if h_err != nil {
        return fmt.Errorf("operation [CreateCredentialAzure] failed, Error %s", h_err)
    }
    resp, err := handler.CvCreateCredentialAzure(req)
    if err != nil {
        return fmt.Errorf("operation [CreateCredentialAzure] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateCredentialAzure] failed")
    } else {
        d.SetId(response_id)
        return resourceReadCredential_Azure(d, m)
    }
}

func resourceReadCredential_Azure(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V5/Credential/{credentialId}
    resp, err := handler.CvGetCredentialDetailsAzure(d.Id())
    if err != nil {
        if strings.Contains(err.Error(), "status: 404") {
            handler.LogEntry("debug", "entity not present, removing from state")
            d.SetId("")
            return nil
        }
        return fmt.Errorf("operation [GetCredentialDetailsAzure] failed, Error %s", err)
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
    if resp.AccessKeyId != nil {
        d.Set("accesskeyid", resp.AccessKeyId)
    }
    if resp.AccountName != nil {
        d.Set("accountname", resp.AccountName)
    }
    if resp.Description != nil {
        d.Set("description", resp.Description)
    }
    return nil
}

func resourceUpdateCredential_Azure(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V5/Credential/{credentialId}
    var t_accesskeyid *string
    if val, ok := d.GetOk("accesskeyid"); ok {
        t_accesskeyid = handler.ToStringValue(val, false)
    }
    var t_newname *string
    if val, ok := d.GetOk("name"); ok {
        t_newname = handler.ToStringValue(val, false)
    }
    var t_accountname *string
    if val, ok := d.GetOk("accountname"); ok {
        t_accountname = handler.ToStringValue(val, false)
    }
    var t_description *string
    if val, ok := d.GetOk("description"); ok {
        t_description = handler.ToStringValue(val, false)
    }
    var req = handler.MsgUpdateCredentialAzureRequest{AccessKeyId:t_accesskeyid, NewName:t_newname, AccountName:t_accountname, Description:t_description}
    h_err := handler.UpdateCredential_Azure(&req, d, m)
    if h_err != nil {
        return fmt.Errorf("operation [UpdateCredentialAzure] failed, Error %s", h_err)
    }
    _, err := handler.CvUpdateCredentialAzure(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [UpdateCredentialAzure] failed, Error %s", err)
    }
    return resourceReadCredential_Azure(d, m)
}

func resourceDeleteCredential_Azure(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V5/Credential/{credentialId}
    _, err := handler.CvDeleteCredential(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteCredential] failed, Error %s", err)
    }
    return nil
}
