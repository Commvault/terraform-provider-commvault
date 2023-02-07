package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceStorage_Disk() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateStorage_Disk,
        Read:   resourceReadStorage_Disk,
        Update: resourceUpdateStorage_Disk,
        Delete: resourceDeleteStorage_Disk,

        Schema: map[string]*schema.Schema{
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Name of the Disk Storage to be created.",
            },
            "enablededuplication": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "enables or disables deduplication",
            },
            "deduplicationdbstorage": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "A list of dedupe locations can be provided for the storage pool being created. This provides an efficient way to save/store data by eliminating duplicate blocks of data during backups.",
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
            "storage": {
                Type:        schema.TypeSet,
                Required:    true,
                Description: "A list of backup locations can be provided for the storage pool being created.",
                Elem: &schema.Resource{
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
                        "credentials": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "password": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Sensitive:    true,
                                        Description: "password to access the network path",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "username to access the network path",
                                    },
                                },
                            },
                        },
                        "backuplocation": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "",
                        },
                        "savedcredentials": {
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
                        "user": {
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
                        "usergroup": {
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
            "dataencryption": {
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

func resourceCreateStorage_Disk(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/Storage/Disk
    var response_id = strconv.Itoa(0)
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_enablededuplication *bool
    if val, ok := d.GetOk("enablededuplication"); ok {
        t_enablededuplication = handler.ToBooleanValue(val, false)
    }
    var t_deduplicationdbstorage []handler.MsgDedupePathSet
    if val, ok := d.GetOk("deduplicationdbstorage"); ok {
        t_deduplicationdbstorage = build_storage_disk_msgdedupepathset_array(d, val.(*schema.Set).List())
    }
    var t_storage []handler.MsgPathSet
    if val, ok := d.GetOk("storage"); ok {
        t_storage = build_storage_disk_msgpathset_array(d, val.(*schema.Set).List())
    }
    var req = handler.MsgCreateDiskStorageRequest{Name:t_name, EnableDeduplication:t_enablededuplication, DeduplicationDBStorage:t_deduplicationdbstorage, Storage:t_storage}
    resp, err := handler.CvCreateDiskStorage(req)
    if err != nil {
        return fmt.Errorf("operation [CreateDiskStorage] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateDiskStorage] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateStorage_Disk(d, m)
    }
}

func resourceReadStorage_Disk(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/Storage/Disk/{storagePoolId}
    resp, err := handler.CvGetDiskStorageDetails(d.Id())
    if err != nil {
        return fmt.Errorf("operation [GetDiskStorageDetails] failed, Error %s", err)
    }
    if rtn, ok := serialize_storage_disk_msgsecurityassoc(d, resp.Security); ok {
        d.Set("security", rtn)
    } else {
        d.Set("security", make([]map[string]interface{}, 0))
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    return nil
}

func resourceUpdateStorage_Disk(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Storage/Disk/{storagePoolId}
    var t_security []handler.MsgUpdateSecurityAssocSet
    if d.HasChange("security") {
        val := d.Get("security")
        t_security = build_storage_disk_msgupdatesecurityassocset_array(d, val.(*schema.Set).List())
    }
    var t_newname *string
    if d.HasChange("name") {
        val := d.Get("name")
        t_newname = handler.ToStringValue(val, false)
    }
    var t_dataencryption *handler.MsgEncryption
    if d.HasChange("dataencryption") {
        val := d.Get("dataencryption")
        t_dataencryption = build_storage_disk_msgencryption(d, val.([]interface{}))
    }
    var req = handler.MsgModifyDiskStorageRequest{Security:t_security, NewName:t_newname, DataEncryption:t_dataencryption}
    _, err := handler.CvModifyDiskStorage(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [ModifyDiskStorage] failed, Error %s", err)
    }
    return resourceReadStorage_Disk(d, m)
}

func resourceCreateUpdateStorage_Disk(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Storage/Disk/{storagePoolId}
    var execUpdate bool = false
    var t_security []handler.MsgUpdateSecurityAssocSet
    if val, ok := d.GetOk("security"); ok {
        t_security = build_storage_disk_msgupdatesecurityassocset_array(d, val.(*schema.Set).List())
        execUpdate = true
    }
    var t_dataencryption *handler.MsgEncryption
    if val, ok := d.GetOk("dataencryption"); ok {
        t_dataencryption = build_storage_disk_msgencryption(d, val.([]interface{}))
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgModifyDiskStorageRequest{Security:t_security, DataEncryption:t_dataencryption}
        _, err := handler.CvModifyDiskStorage(req, d.Id())
        if err != nil {
            return fmt.Errorf("operation [ModifyDiskStorage] failed, Error %s", err)
        }
    }
    return resourceReadStorage_Disk(d, m)
}

func resourceDeleteStorage_Disk(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/Storage/Disk/{storagePoolId}
    _, err := handler.CvDeleteDiskStorage(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteDiskStorage] failed, Error %s", err)
    }
    return nil
}

func build_storage_disk_msgencryption(d *schema.ResourceData, r []interface{}) *handler.MsgEncryption {
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
            t_keyprovider = build_storage_disk_msgidname(d, val.([]interface{}))
        }
        return &handler.MsgEncryption{Cipher:t_cipher, KeyLength:t_keylength, Encrypt:t_encrypt, KeyProvider:t_keyprovider}
    } else {
        return nil
    }
}

func build_storage_disk_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_storage_disk_msgupdatesecurityassocset_array(d *schema.ResourceData, r []interface{}) []handler.MsgUpdateSecurityAssocSet {
    if r != nil {
        tmp := make([]handler.MsgUpdateSecurityAssocSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_role *handler.MsgIdName
            if val, ok := raw_a["role"]; ok {
                t_role = build_storage_disk_msgidname(d, val.([]interface{}))
            }
            var t_user *handler.MsgIdName
            if val, ok := raw_a["user"]; ok {
                t_user = build_storage_disk_msgidname(d, val.([]interface{}))
            }
            var t_usergroup *handler.MsgIdName
            if val, ok := raw_a["usergroup"]; ok {
                t_usergroup = build_storage_disk_msgidname(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgUpdateSecurityAssocSet{Role:t_role, User:t_user, UserGroup:t_usergroup}
        }
        return tmp
    } else {
        return nil
    }
}

func build_storage_disk_msgpathset_array(d *schema.ResourceData, r []interface{}) []handler.MsgPathSet {
    if r != nil {
        tmp := make([]handler.MsgPathSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_mediaagent *handler.MsgIdName
            if val, ok := raw_a["mediaagent"]; ok {
                t_mediaagent = build_storage_disk_msgidname(d, val.([]interface{}))
            }
            var t_credentials *handler.MsgUserNamePassword
            if val, ok := raw_a["credentials"]; ok {
                t_credentials = build_storage_disk_msgusernamepassword(d, val.([]interface{}))
            }
            var t_backuplocation *string
            if val, ok := raw_a["backuplocation"]; ok {
                t_backuplocation = handler.ToStringValue(val, true)
            }
            var t_savedcredentials *handler.MsgIdName
            if val, ok := raw_a["savedcredentials"]; ok {
                t_savedcredentials = build_storage_disk_msgidname(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgPathSet{MediaAgent:t_mediaagent, Credentials:t_credentials, BackupLocation:t_backuplocation, SavedCredentials:t_savedcredentials}
        }
        return tmp
    } else {
        return nil
    }
}

func build_storage_disk_msgusernamepassword(d *schema.ResourceData, r []interface{}) *handler.MsgUserNamePassword {
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

func build_storage_disk_msgdedupepathset_array(d *schema.ResourceData, r []interface{}) []handler.MsgDedupePathSet {
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
                t_mediaagent = build_storage_disk_msgidname(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgDedupePathSet{Path:t_path, MediaAgent:t_mediaagent}
        }
        return tmp
    } else {
        return nil
    }
}

func serialize_storage_disk_msgsecurityassoc(d *schema.ResourceData, data *handler.MsgSecurityAssoc) ([]map[string]interface{}, bool) {
    //MsgUpdateSecurityAssocSet
    //MsgSecurityAssoc
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if rtn, ok := serialize_storage_disk_msgidname(d, data.Role); ok {
        val[0]["role"] = rtn
        added = true
    }
    if rtn, ok := serialize_storage_disk_msgidname(d, data.User); ok {
        val[0]["user"] = rtn
        added = true
    }
    if rtn, ok := serialize_storage_disk_msgidname(d, data.UserGroup); ok {
        val[0]["usergroup"] = rtn
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_storage_disk_msgidname(d *schema.ResourceData, data *handler.MsgIdName) ([]map[string]interface{}, bool) {
    //MsgUpdateSecurityAssocSet -> MsgIdName
    //MsgSecurityAssoc -> MsgIdName
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
