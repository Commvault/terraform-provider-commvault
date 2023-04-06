package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCredential_AWS() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateCredential_AWS,
        Read:   resourceReadCredential_AWS,
        Update: resourceUpdateCredential_AWS,
        Delete: resourceDeleteCredential_AWS,

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
                Description: "Access Key ID of Credential",
            },
            "secretaccesskey": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Secret Access Key of Credential and must be in base64 encoded format.",
            },
            "security": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Security association of a list of users and user groups",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "owner": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "Owner of a credential can be a user or user group",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "user": {
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
                                    "usergroup": {
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
                                },
                            },
                        },
                        "associations": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "iscreatorassociation": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "To check if the user/user group associated is the owner.",
                                    },
                                    "permissions": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "List of permissions associated with the entity. Either categoryId and categoryName or permissionId and permissionName will be returned. If categoryId or categoryName is returned, all the corresponding permissions in the category are associated with the entity.",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "permissionid": {
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "exclude": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Flag to specify if this is included permission or excluded permission.",
                                                },
                                                "type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Returns the type of association. [ALL_CATEGORIES, CATEGORY_ENTITY, PERMISSION_ENTITY]",
                                                },
                                                "categoryname": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "categoryid": {
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "permissionname": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                            },
                                        },
                                    },
                                    "user": {
                                        Type:        schema.TypeList,
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
                                    "usergroup": {
                                        Type:        schema.TypeList,
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
                    },
                },
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

func resourceCreateCredential_AWS(d *schema.ResourceData, m interface{}) error {
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
    var t_secretaccesskey *string
    if val, ok := d.GetOk("secretaccesskey"); ok {
        t_secretaccesskey = handler.ToStringValue(val, false)
    }
    var t_authtype *string
    var c_authtype string = "AMAZON_S3"
    t_authtype = &c_authtype
    var t_security *handler.MsgCredentialSecurity
    if val, ok := d.GetOk("security"); ok {
        t_security = build_credential_aws_msgcredentialsecurity(d, val.([]interface{}))
    }
    var t_description *string
    if val, ok := d.GetOk("description"); ok {
        t_description = handler.ToStringValue(val, false)
    }
    var req = handler.MsgCreateCredentialAWSRequest{VendorType:t_vendortype, AccountType:t_accounttype, Name:t_name, AccessKeyId:t_accesskeyid, SecretAccessKey:t_secretaccesskey, AuthType:t_authtype, Security:t_security, Description:t_description}
    h_err := handler.ConfigureCredential_AWS(&req, d, m)
    if h_err != nil {
        return fmt.Errorf("operation [CreateCredentialAWS] failed, Error %s", h_err)
    }
    resp, err := handler.CvCreateCredentialAWS(req)
    if err != nil {
        return fmt.Errorf("operation [CreateCredentialAWS] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateCredentialAWS] failed")
    } else {
        d.SetId(response_id)
        return resourceReadCredential_AWS(d, m)
    }
}

func resourceReadCredential_AWS(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V5/Credential/{credentialId}
    resp, err := handler.CvGetCredentailDetailsAWS(d.Id())
    if err != nil {
        return fmt.Errorf("operation [GetCredentailDetailsAWS] failed, Error %s", err)
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
    if resp.SecretAccessKey != nil {
        d.Set("secretaccesskey", resp.SecretAccessKey)
    }
    if rtn, ok := serialize_credential_aws_msgcredentialsecurity(d, resp.Security); ok {
        d.Set("security", rtn)
    } else {
        d.Set("security", make([]map[string]interface{}, 0))
    }
    if resp.Description != nil {
        d.Set("description", resp.Description)
    }
    return nil
}

func resourceUpdateCredential_AWS(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V5/Credential/{credentialId}
    var t_accesskeyid *string
    if val, ok := d.GetOk("accesskeyid"); ok {
        t_accesskeyid = handler.ToStringValue(val, false)
    }
    var t_secretaccesskey *string
    if val, ok := d.GetOk("secretaccesskey"); ok {
        t_secretaccesskey = handler.ToStringValue(val, false)
    }
    var t_security *handler.MsgUpdateCredentialSecurity
    if val, ok := d.GetOk("security"); ok {
        t_security = build_credential_aws_msgupdatecredentialsecurity(d, val.([]interface{}))
    }
    var t_newname *string
    if val, ok := d.GetOk("name"); ok {
        t_newname = handler.ToStringValue(val, false)
    }
    var t_description *string
    if val, ok := d.GetOk("description"); ok {
        t_description = handler.ToStringValue(val, false)
    }
    var req = handler.MsgUpdateCredentialAWSRequest{AccessKeyId:t_accesskeyid, SecretAccessKey:t_secretaccesskey, Security:t_security, NewName:t_newname, Description:t_description}
    h_err := handler.UpdateCredential_AWS(&req, d, m)
    if h_err != nil {
        return fmt.Errorf("operation [UpdateCredentialAWS] failed, Error %s", h_err)
    }
    _, err := handler.CvUpdateCredentialAWS(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [UpdateCredentialAWS] failed, Error %s", err)
    }
    return resourceReadCredential_AWS(d, m)
}

func resourceDeleteCredential_AWS(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V5/Credential/{credentialId}
    _, err := handler.CvDeleteCredential(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteCredential] failed, Error %s", err)
    }
    return nil
}

func build_credential_aws_msgupdatecredentialsecurity(d *schema.ResourceData, r []interface{}) *handler.MsgUpdateCredentialSecurity {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_owner *handler.MsgCredentialOwner
        if val, ok := tmp["owner"]; ok {
            t_owner = build_credential_aws_msgcredentialowner(d, val.([]interface{}))
        }
        var t_associations []handler.MsgCredentialSecurityAssociationsSet
        if val, ok := tmp["associations"]; ok {
            t_associations = build_credential_aws_msgcredentialsecurityassociationsset_array(d, val.(*schema.Set).List())
        }
        return &handler.MsgUpdateCredentialSecurity{Owner:t_owner, Associations:t_associations}
    } else {
        return nil
    }
}

func build_credential_aws_msgcredentialsecurityassociationsset_array(d *schema.ResourceData, r []interface{}) []handler.MsgCredentialSecurityAssociationsSet {
    if r != nil {
        tmp := make([]handler.MsgCredentialSecurityAssociationsSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_iscreatorassociation *bool
            if val, ok := raw_a["iscreatorassociation"]; ok {
                t_iscreatorassociation = handler.ToBooleanValue(val, true)
            }
            var t_permissions *handler.MsgPermissionResp
            if val, ok := raw_a["permissions"]; ok {
                t_permissions = build_credential_aws_msgpermissionresp(d, val.([]interface{}))
            }
            var t_user *handler.MsgIdName
            if val, ok := raw_a["user"]; ok {
                t_user = build_credential_aws_msgidname(d, val.([]interface{}))
            }
            var t_usergroup *handler.MsgIdName
            if val, ok := raw_a["usergroup"]; ok {
                t_usergroup = build_credential_aws_msgidname(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgCredentialSecurityAssociationsSet{IsCreatorAssociation:t_iscreatorassociation, Permissions:t_permissions, User:t_user, UserGroup:t_usergroup}
        }
        return tmp
    } else {
        return nil
    }
}

func build_credential_aws_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_id *int
        if val, ok := tmp["id"]; ok {
            t_id = handler.ToIntValue(val, true)
        }
        return &handler.MsgIdName{Id:t_id}
    } else {
        return nil
    }
}

func build_credential_aws_msgpermissionresp(d *schema.ResourceData, r []interface{}) *handler.MsgPermissionResp {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_permissionid *int
        if val, ok := tmp["permissionid"]; ok {
            t_permissionid = handler.ToIntValue(val, true)
        }
        var t_exclude *bool
        if val, ok := tmp["exclude"]; ok {
            t_exclude = handler.ToBooleanValue(val, true)
        }
        var t_type *string
        if val, ok := tmp["type"]; ok {
            t_type = handler.ToStringValue(val, true)
        }
        var t_categoryname *string
        if val, ok := tmp["categoryname"]; ok {
            t_categoryname = handler.ToStringValue(val, true)
        }
        var t_categoryid *int
        if val, ok := tmp["categoryid"]; ok {
            t_categoryid = handler.ToIntValue(val, true)
        }
        var t_permissionname *string
        if val, ok := tmp["permissionname"]; ok {
            t_permissionname = handler.ToStringValue(val, true)
        }
        return &handler.MsgPermissionResp{PermissionId:t_permissionid, Exclude:t_exclude, Type:t_type, CategoryName:t_categoryname, CategoryId:t_categoryid, PermissionName:t_permissionname}
    } else {
        return nil
    }
}

func build_credential_aws_msgcredentialowner(d *schema.ResourceData, r []interface{}) *handler.MsgCredentialOwner {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_user *handler.MsgIdName
        if val, ok := tmp["user"]; ok {
            t_user = build_credential_aws_msgidname(d, val.([]interface{}))
        }
        var t_usergroup *handler.MsgIdName
        if val, ok := tmp["usergroup"]; ok {
            t_usergroup = build_credential_aws_msgidname(d, val.([]interface{}))
        }
        return &handler.MsgCredentialOwner{User:t_user, UserGroup:t_usergroup}
    } else {
        return nil
    }
}

func build_credential_aws_msgcredentialsecurity(d *schema.ResourceData, r []interface{}) *handler.MsgCredentialSecurity {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_owner *handler.MsgCredentialOwner
        if val, ok := tmp["owner"]; ok {
            t_owner = build_credential_aws_msgcredentialowner(d, val.([]interface{}))
        }
        var t_associations []handler.MsgCredentialSecurityAssociationsSet
        if val, ok := tmp["associations"]; ok {
            t_associations = build_credential_aws_msgcredentialsecurityassociationsset_array(d, val.(*schema.Set).List())
        }
        return &handler.MsgCredentialSecurity{Owner:t_owner, Associations:t_associations}
    } else {
        return nil
    }
}

func serialize_credential_aws_msgcredentialsecurity(d *schema.ResourceData, data *handler.MsgCredentialSecurity) ([]map[string]interface{}, bool) {
    //MsgCredentialSecurity
    //MsgCredentialSecurity
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_credential_aws_msgcredentialowner(d, data.Owner); ok {
        val[0]["owner"] = rtn
        added = true
    }
    if rtn, ok := serialize_credential_aws_msgcredentialsecurityassociationsset_array(d, data.Associations); ok {
        val[0]["associations"] = rtn
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_credential_aws_msgcredentialsecurityassociationsset_array(d *schema.ResourceData, data []handler.MsgCredentialSecurityAssociationsSet) ([]map[string]interface{}, bool) {
    //MsgCredentialSecurity -> MsgCredentialSecurityAssociationsSet
    //MsgCredentialSecurity -> MsgCredentialSecurityAssociationsSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].IsCreatorAssociation != nil {
            tmp["iscreatorassociation"] = strconv.FormatBool(*data[i].IsCreatorAssociation)
            added = true
        }
        if rtn, ok := serialize_credential_aws_msgpermissionresp(d, data[i].Permissions); ok {
            tmp["permissions"] = rtn
            added = true
        }
        if rtn, ok := serialize_credential_aws_msgidname(d, data[i].User); ok {
            tmp["user"] = rtn
            added = true
        }
        if rtn, ok := serialize_credential_aws_msgidname(d, data[i].UserGroup); ok {
            tmp["usergroup"] = rtn
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_credential_aws_msgidname(d *schema.ResourceData, data *handler.MsgIdName) ([]map[string]interface{}, bool) {
    //MsgCredentialSecurity -> MsgCredentialSecurityAssociationsSet -> MsgIdName
    //MsgCredentialSecurity -> MsgCredentialSecurityAssociationsSet -> MsgIdName
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
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

func serialize_credential_aws_msgpermissionresp(d *schema.ResourceData, data *handler.MsgPermissionResp) ([]map[string]interface{}, bool) {
    //MsgCredentialSecurity -> MsgCredentialSecurityAssociationsSet -> MsgPermissionResp
    //MsgCredentialSecurity -> MsgCredentialSecurityAssociationsSet -> MsgPermissionResp
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.PermissionId != nil {
        val[0]["permissionid"] = data.PermissionId
        added = true
    }
    if data.Exclude != nil {
        val[0]["exclude"] = strconv.FormatBool(*data.Exclude)
        added = true
    }
    if data.Type != nil {
        val[0]["type"] = data.Type
        added = true
    }
    if data.CategoryName != nil {
        val[0]["categoryname"] = data.CategoryName
        added = true
    }
    if data.CategoryId != nil {
        val[0]["categoryid"] = data.CategoryId
        added = true
    }
    if data.PermissionName != nil {
        val[0]["permissionname"] = data.PermissionName
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_credential_aws_msgcredentialowner(d *schema.ResourceData, data *handler.MsgCredentialOwner) ([]map[string]interface{}, bool) {
    //MsgCredentialSecurity -> MsgCredentialOwner
    //MsgCredentialSecurity -> MsgCredentialOwner
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_credential_aws_msgidname(d, data.User); ok {
        val[0]["user"] = rtn
        added = true
    }
    if rtn, ok := serialize_credential_aws_msgidname(d, data.UserGroup); ok {
        val[0]["usergroup"] = rtn
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}
