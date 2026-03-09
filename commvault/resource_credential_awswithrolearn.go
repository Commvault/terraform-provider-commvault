package commvault

import (
    "fmt"
    "strconv"
    "strings"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCredential_AWSWithRoleArn() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateCredential_AWSWithRoleArn,
        Read:   resourceReadCredential_AWSWithRoleArn,
        Update: resourceUpdateCredential_AWSWithRoleArn,
        Delete: resourceDeleteCredential_AWSWithRoleArn,

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
            "password": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Role ARN of credential",
            },
            "rolearn": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Role ARN of credential",
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

func resourceCreateCredential_AWSWithRoleArn(d *schema.ResourceData, m interface{}) error {
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
    var t_password *string
    if val, ok := d.GetOk("password"); ok {
        t_password = handler.ToStringValue(val, false)
    }
    var t_rolearn *string
    if val, ok := d.GetOk("rolearn"); ok {
        t_rolearn = handler.ToStringValue(val, false)
    }
    var t_authtype *string
    var c_authtype string = "AMAZON_STS_IAM_ROLE"
    t_authtype = &c_authtype
    var t_description *string
    if val, ok := d.GetOk("description"); ok {
        t_description = handler.ToStringValue(val, false)
    }
    var req = handler.MsgCreateCredentialAWSWithRoleArnRequest{VendorType:t_vendortype, AccountType:t_accounttype, Name:t_name, Password:t_password, RoleArn:t_rolearn, AuthType:t_authtype, Description:t_description}
    h_err := handler.ConfigureCredential_AWSWithRoleArn(&req, d, m)
    if h_err != nil {
        return fmt.Errorf("operation [CreateCredentialAWSWithRoleArn] failed, Error %s", h_err)
    }
    resp, err := handler.CvCreateCredentialAWSWithRoleArn(req)
    if err != nil {
        return fmt.Errorf("operation [CreateCredentialAWSWithRoleArn] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateCredentialAWSWithRoleArn] failed")
    } else {
        d.SetId(response_id)
        return resourceReadCredential_AWSWithRoleArn(d, m)
    }
}

func resourceReadCredential_AWSWithRoleArn(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V5/Credential/{credentialId}
    resp, err := handler.CvGetCredentialDetailsAWSWithRoleArn(d.Id())
    if err != nil {
        if strings.Contains(err.Error(), "status: 404") {
            handler.LogEntry("debug", "entity not present, removing from state")
            d.SetId("")
            return nil
        }
        return fmt.Errorf("operation [GetCredentialDetailsAWSWithRoleArn] failed, Error %s", err)
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
    if resp.Password != nil {
        d.Set("password", resp.Password)
    }
    if resp.RoleArn != nil {
        d.Set("rolearn", resp.RoleArn)
    }
    if resp.Description != nil {
        d.Set("description", resp.Description)
    }
    return nil
}

func resourceUpdateCredential_AWSWithRoleArn(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V5/Credential/{credentialId}
    var t_password *string
    if val, ok := d.GetOk("password"); ok {
        t_password = handler.ToStringValue(val, false)
    }
    var t_newname *string
    if val, ok := d.GetOk("name"); ok {
        t_newname = handler.ToStringValue(val, false)
    }
    var t_rolearn *string
    if val, ok := d.GetOk("rolearn"); ok {
        t_rolearn = handler.ToStringValue(val, false)
    }
    var t_description *string
    if val, ok := d.GetOk("description"); ok {
        t_description = handler.ToStringValue(val, false)
    }
    var req = handler.MsgUpdateCredentialAWSWithRoleArnRequest{Password:t_password, NewName:t_newname, RoleArn:t_rolearn, Description:t_description}
    h_err := handler.UpdateCredential_AWSWithRoleArn(&req, d, m)
    if h_err != nil {
        return fmt.Errorf("operation [UpdateCredentialAWSWithRoleArn] failed, Error %s", h_err)
    }
    _, err := handler.CvUpdateCredentialAWSWithRoleArn(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [UpdateCredentialAWSWithRoleArn] failed, Error %s", err)
    }
    return resourceReadCredential_AWSWithRoleArn(d, m)
}

func resourceDeleteCredential_AWSWithRoleArn(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V5/Credential/{credentialId}
    _, err := handler.CvDeleteCredential(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteCredential] failed, Error %s", err)
    }
    return nil
}
