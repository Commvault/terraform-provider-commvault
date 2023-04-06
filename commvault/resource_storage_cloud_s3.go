package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceStorage_Cloud_S3() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateStorage_Cloud_S3,
        Read:   resourceReadStorage_Cloud_S3,
        Update: resourceUpdateStorage_Cloud_S3,
        Delete: resourceDeleteStorage_Cloud_S3,

        Schema: map[string]*schema.Schema{
            "mediaagent": {
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
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Name of the cloud storage library",
            },
            "bucket": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Name of bucket",
            },
            "storageclass": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Appropriate storage class for your account [Standard, Reduced Redundancy Storage, Standard - Infrequent access, One zone - Infrequent access, Intelligent tiering, S3 Glacier, Standard/Glacier (Combined Storage Tiers), Standard-IA/Glacier (Combined Storage Tiers), One Zone-IA/Glacier (Combined Storage Tiers), Intelligent-Tiering/Glacier (Combined Storage Tiers), S3 Glacier Deep Archive, Standard/Deep Archive (Combined Storage Tiers), Standard-IA/Deep Archive (Combined Storage Tiers), One Zone-IA/Deep Archive (Combined Storage Tiers), Intelligent-Tiering/Deep Archive (Combined Storage Tiers)]",
            },
            "servicehost": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "IP address or fully qualified domain name or URL for the cloud library based on cloud vendor",
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
            "arnrole": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Needed for AWS STS assume role and AWS STS assume role with IAM role policy",
            },
            "authentication": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Authentication type for the cloud storage server [Access and secret keys, AWS IAM role policy, AWS STS assume role, AWS STS assume role with IAM role policy]",
            },
            "deduplicationdblocation": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "A list of dedupe locations can be provided for the storage pool being created. This provides an efficient way to save/store data by eliminating duplicate blocks of data during backups",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "path": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "",
                        },
                        "mediaagent": {
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
            "usededuplication": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Enables or disables deduplication on the storage",
            },
            "security": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "role": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
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
                        "user": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
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
                        "usergroup": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
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
                },
            },
            "encryption": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Different ways in which data can be encrypted.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "cipher": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "The different types of encryption keys that can be used for encrypting the data. The values are case sensitive [BlowFish, AES, DES3, GOST, Serpent, Twofish]",
                        },
                        "keylength": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Different keylengths are present for different kinds of ciphers. Blowfish,Twofish,AES and Serpent all accept both 128 and 256. DES3 accepts only 192. GOST accepts only 256. ",
                        },
                        "encrypt": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "keyprovider": {
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
        },
    }
}

func resourceCreateStorage_Cloud_S3(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/Storage/Cloud
    var response_id = strconv.Itoa(0)
    var t_mediaagent *handler.MsgIdName
    if val, ok := d.GetOk("mediaagent"); ok {
        t_mediaagent = build_storage_cloud_s3_msgidname(d, val.([]interface{}))
    }
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_bucket *string
    if val, ok := d.GetOk("bucket"); ok {
        t_bucket = handler.ToStringValue(val, false)
    }
    var t_storageclass *string
    if val, ok := d.GetOk("storageclass"); ok {
        t_storageclass = handler.ToStringValue(val, false)
    }
    var t_servicehost *string
    if val, ok := d.GetOk("servicehost"); ok {
        t_servicehost = handler.ToStringValue(val, false)
    }
    var t_credentials *handler.MsgIdName
    if val, ok := d.GetOk("credentials"); ok {
        t_credentials = build_storage_cloud_s3_msgidname(d, val.([]interface{}))
    }
    var t_cloudtype *string
    var c_cloudtype string = "Amazon S3"
    t_cloudtype = &c_cloudtype
    var t_arnrole *string
    if val, ok := d.GetOk("arnrole"); ok {
        t_arnrole = handler.ToStringValue(val, false)
    }
    var t_authentication *string
    if val, ok := d.GetOk("authentication"); ok {
        t_authentication = handler.ToStringValue(val, false)
    }
    var t_deduplicationdblocation []handler.MsgDedupePathSet
    if val, ok := d.GetOk("deduplicationdblocation"); ok {
        t_deduplicationdblocation = build_storage_cloud_s3_msgdedupepathset_array(d, val.(*schema.Set).List())
    }
    var t_usededuplication *bool
    if val, ok := d.GetOk("usededuplication"); ok {
        t_usededuplication = handler.ToBooleanValue(val, false)
    }
    var req = handler.MsgCreateCloudStorageS3Request{MediaAgent:t_mediaagent, Name:t_name, Bucket:t_bucket, StorageClass:t_storageclass, ServiceHost:t_servicehost, Credentials:t_credentials, CloudType:t_cloudtype, ArnRole:t_arnrole, Authentication:t_authentication, DeduplicationDBLocation:t_deduplicationdblocation, UseDeduplication:t_usededuplication}
    resp, err := handler.CvCreateCloudStorageS3(req)
    if err != nil {
        return fmt.Errorf("operation [CreateCloudStorageS3] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateCloudStorageS3] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateStorage_Cloud_S3(d, m)
    }
}

func resourceReadStorage_Cloud_S3(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/Storage/Cloud/{cloudStorageId}
    resp, err := handler.CvGetCloudStorageById(d.Id())
    if err != nil {
        return fmt.Errorf("operation [GetCloudStorageById] failed, Error %s", err)
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    if rtn, ok := serialize_storage_cloud_s3_msgidnamestatusset_array(d, resp.Bucket); ok {
        d.Set("bucket", rtn)
    } else {
        d.Set("bucket", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_storage_cloud_s3_msgsecurityassocset_array(d, resp.Security); ok {
        d.Set("security", rtn)
    } else {
        d.Set("security", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_storage_cloud_s3_msgencryption(d, resp.Encryption); ok {
        d.Set("encryption", rtn)
    } else {
        d.Set("encryption", make([]map[string]interface{}, 0))
    }
    return nil
}

func resourceUpdateStorage_Cloud_S3(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Storage/Cloud/{cloudStorageId}
    var t_security []handler.MsgUpdateSecurityAssocSet
    if d.HasChange("security") {
        val := d.Get("security")
        t_security = build_storage_cloud_s3_msgupdatesecurityassocset_array(d, val.(*schema.Set).List())
    }
    var t_newname *string
    if d.HasChange("name") {
        val := d.Get("name")
        t_newname = handler.ToStringValue(val, false)
    }
    var t_encryption *handler.MsgEncryption
    if d.HasChange("encryption") {
        val := d.Get("encryption")
        t_encryption = build_storage_cloud_s3_msgencryption(d, val.([]interface{}))
    }
    var req = handler.MsgModifyCloudStorageByIdRequest{Security:t_security, NewName:t_newname, Encryption:t_encryption}
    _, err := handler.CvModifyCloudStorageById(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [ModifyCloudStorageById] failed, Error %s", err)
    }
    return resourceReadStorage_Cloud_S3(d, m)
}

func resourceCreateUpdateStorage_Cloud_S3(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Storage/Cloud/{cloudStorageId}
    var execUpdate bool = false
    var t_security []handler.MsgUpdateSecurityAssocSet
    if val, ok := d.GetOk("security"); ok {
        t_security = build_storage_cloud_s3_msgupdatesecurityassocset_array(d, val.(*schema.Set).List())
        execUpdate = true
    }
    var t_encryption *handler.MsgEncryption
    if val, ok := d.GetOk("encryption"); ok {
        t_encryption = build_storage_cloud_s3_msgencryption(d, val.([]interface{}))
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgModifyCloudStorageByIdRequest{Security:t_security, Encryption:t_encryption}
        _, err := handler.CvModifyCloudStorageById(req, d.Id())
        if err != nil {
            return fmt.Errorf("operation [ModifyCloudStorageById] failed, Error %s", err)
        }
    }
    return resourceReadStorage_Cloud_S3(d, m)
}

func resourceDeleteStorage_Cloud_S3(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/Storage/Cloud/{cloudStorageId}
    _, err := handler.CvDeleteCloudStorageById(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteCloudStorageById] failed, Error %s", err)
    }
    return nil
}

func build_storage_cloud_s3_msgencryption(d *schema.ResourceData, r []interface{}) *handler.MsgEncryption {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_cipher *string
        if val, ok := tmp["cipher"]; ok {
            t_cipher = handler.ToStringValue(val, true)
        }
        var t_keylength *int
        if val, ok := tmp["keylength"]; ok {
            t_keylength = handler.ToIntValue(val, true)
        }
        var t_encrypt *bool
        if val, ok := tmp["encrypt"]; ok {
            t_encrypt = handler.ToBooleanValue(val, true)
        }
        var t_keyprovider *handler.MsgIdName
        if val, ok := tmp["keyprovider"]; ok {
            t_keyprovider = build_storage_cloud_s3_msgidname(d, val.([]interface{}))
        }
        return &handler.MsgEncryption{Cipher:t_cipher, KeyLength:t_keylength, Encrypt:t_encrypt, KeyProvider:t_keyprovider}
    } else {
        return nil
    }
}

func build_storage_cloud_s3_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_storage_cloud_s3_msgupdatesecurityassocset_array(d *schema.ResourceData, r []interface{}) []handler.MsgUpdateSecurityAssocSet {
    if r != nil {
        tmp := make([]handler.MsgUpdateSecurityAssocSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_role *handler.MsgIdName
            if val, ok := raw_a["role"]; ok {
                t_role = build_storage_cloud_s3_msgidname(d, val.([]interface{}))
            }
            var t_user *handler.MsgIdName
            if val, ok := raw_a["user"]; ok {
                t_user = build_storage_cloud_s3_msgidname(d, val.([]interface{}))
            }
            var t_usergroup *handler.MsgIdName
            if val, ok := raw_a["usergroup"]; ok {
                t_usergroup = build_storage_cloud_s3_msgidname(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgUpdateSecurityAssocSet{Role:t_role, User:t_user, UserGroup:t_usergroup}
        }
        return tmp
    } else {
        return nil
    }
}

func build_storage_cloud_s3_msgdedupepathset_array(d *schema.ResourceData, r []interface{}) []handler.MsgDedupePathSet {
    if r != nil {
        tmp := make([]handler.MsgDedupePathSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_path *string
            if val, ok := raw_a["path"]; ok {
                t_path = handler.ToStringValue(val, true)
            }
            var t_mediaagent *handler.MsgIdName
            if val, ok := raw_a["mediaagent"]; ok {
                t_mediaagent = build_storage_cloud_s3_msgidname(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgDedupePathSet{Path:t_path, MediaAgent:t_mediaagent}
        }
        return tmp
    } else {
        return nil
    }
}

func serialize_storage_cloud_s3_msgencryption(d *schema.ResourceData, data *handler.MsgEncryption) ([]map[string]interface{}, bool) {
    //MsgEncryption
    //MsgEncryption
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.Cipher != nil {
        val[0]["cipher"] = data.Cipher
        added = true
    }
    if data.KeyLength != nil {
        val[0]["keylength"] = data.KeyLength
        added = true
    }
    if data.Encrypt != nil {
        val[0]["encrypt"] = strconv.FormatBool(*data.Encrypt)
        added = true
    }
    if rtn, ok := serialize_storage_cloud_s3_msgidname(d, data.KeyProvider); ok {
        val[0]["keyprovider"] = rtn
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_storage_cloud_s3_msgidname(d *schema.ResourceData, data *handler.MsgIdName) ([]map[string]interface{}, bool) {
    //MsgEncryption -> MsgIdName
    //MsgEncryption -> MsgIdName
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

func serialize_storage_cloud_s3_msgsecurityassocset_array(d *schema.ResourceData, data []handler.MsgSecurityAssocSet) ([]map[string]interface{}, bool) {
    //MsgUpdateSecurityAssocSet
    //MsgSecurityAssocSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if rtn, ok := serialize_storage_cloud_s3_msgidname(d, data[i].Role); ok {
            tmp["role"] = rtn
            added = true
        }
        if rtn, ok := serialize_storage_cloud_s3_msgidname(d, data[i].User); ok {
            tmp["user"] = rtn
            added = true
        }
        if rtn, ok := serialize_storage_cloud_s3_msgidname(d, data[i].UserGroup); ok {
            tmp["usergroup"] = rtn
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_storage_cloud_s3_msgidnamestatusset_array(d *schema.ResourceData, data []handler.MsgIdNameStatusSet) ([]map[string]interface{}, bool) {
    //Msgnull
    //MsgIdNameStatusSet
    //no child properties in schema
    return nil, false
}
