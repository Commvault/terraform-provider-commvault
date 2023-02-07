package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceStorage_Disk_Backup_Location() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateStorage_Disk_Backup_Location,
        Read:   resourceReadStorage_Disk_Backup_Location,
        Update: resourceUpdateStorage_Disk_Backup_Location,
        Delete: resourceDeleteStorage_Disk_Backup_Location,

        Schema: map[string]*schema.Schema{
            "storagepoolid": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: "Id of the disk storage to update",
            },
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
            "credentials": {
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
            "backuplocation": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "The mount path on the media agent where the data is to be backed up.",
            },
            "savedcredentials": {
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
            "path": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Can be used to change the disk access path.",
            },
            "access": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "The access type for the disk access path can be either read (writing to path not allowed) or read and write (writing to path allowed). [READ_AND_WRITE, READ]",
            },
            "enabled": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "",
            },
        },
    }
}

func resourceCreateStorage_Disk_Backup_Location(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/Storage/Disk/{storagePoolId}/BackupLocation
    var response_id = strconv.Itoa(0)
    var t_mediaagent *handler.MsgIdName
    if val, ok := d.GetOk("mediaagent"); ok {
        t_mediaagent = build_storage_disk_backup_location_msgidname(d, val.([]interface{}))
    }
    var t_credentials *handler.MsgUserNamePassword
    if val, ok := d.GetOk("credentials"); ok {
        t_credentials = build_storage_disk_backup_location_msgusernamepassword(d, val.([]interface{}))
    }
    var t_backuplocation *string
    if val, ok := d.GetOk("backuplocation"); ok {
        t_backuplocation = handler.ToStringValue(val, false)
    }
    var t_savedcredentials *handler.MsgIdName
    if val, ok := d.GetOk("savedcredentials"); ok {
        t_savedcredentials = build_storage_disk_backup_location_msgidname(d, val.([]interface{}))
    }
    var req = handler.MsgCreateBackupLocationRequest{MediaAgent:t_mediaagent, Credentials:t_credentials, BackupLocation:t_backuplocation, SavedCredentials:t_savedcredentials}
    resp, err := handler.CvCreateBackupLocation(req, strconv.Itoa(d.Get("storagepoolid").(int)))
    if err != nil {
        return fmt.Errorf("operation [CreateBackupLocation] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateBackupLocation] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateStorage_Disk_Backup_Location(d, m)
    }
}

func resourceReadStorage_Disk_Backup_Location(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}
    _, err := handler.CvGetBackupLocationDetails(strconv.Itoa(d.Get("storagepoolid").(int)), d.Id())
    if err != nil {
        return fmt.Errorf("operation [GetBackupLocationDetails] failed, Error %s", err)
    }
    return nil
}

func resourceUpdateStorage_Disk_Backup_Location(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}
    var t_path *string
    if d.HasChange("path") {
        val := d.Get("path")
        t_path = handler.ToStringValue(val, false)
    }
    var t_mediaagent *handler.MsgIdName
    if d.HasChange("mediaagent") {
        val := d.Get("mediaagent")
        t_mediaagent = build_storage_disk_backup_location_msgidname(d, val.([]interface{}))
    }
    var t_access *string
    if d.HasChange("access") {
        val := d.Get("access")
        t_access = handler.ToStringValue(val, false)
    }
    var t_credentials *handler.MsgUserNamePassword
    if d.HasChange("credentials") {
        val := d.Get("credentials")
        t_credentials = build_storage_disk_backup_location_msgusernamepassword(d, val.([]interface{}))
    }
    var t_savedcredentials *handler.MsgIdName
    if d.HasChange("savedcredentials") {
        val := d.Get("savedcredentials")
        t_savedcredentials = build_storage_disk_backup_location_msgidname(d, val.([]interface{}))
    }
    var t_enabled *bool
    if d.HasChange("enabled") {
        val := d.Get("enabled")
        t_enabled = handler.ToBooleanValue(val, false)
    }
    var req = handler.MsgModifyBackupLocationRequest{Path:t_path, MediaAgent:t_mediaagent, Access:t_access, Credentials:t_credentials, SavedCredentials:t_savedcredentials, Enabled:t_enabled}
    _, err := handler.CvModifyBackupLocation(req, strconv.Itoa(d.Get("storagepoolid").(int)), d.Id())
    if err != nil {
        return fmt.Errorf("operation [ModifyBackupLocation] failed, Error %s", err)
    }
    return resourceReadStorage_Disk_Backup_Location(d, m)
}

func resourceCreateUpdateStorage_Disk_Backup_Location(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}
    var execUpdate bool = false
    var t_path *string
    if val, ok := d.GetOk("path"); ok {
        t_path = handler.ToStringValue(val, false)
        execUpdate = true
    }
    var t_access *string
    if val, ok := d.GetOk("access"); ok {
        t_access = handler.ToStringValue(val, false)
        execUpdate = true
    }
    var t_enabled *bool
    if val, ok := d.GetOk("enabled"); ok {
        t_enabled = handler.ToBooleanValue(val, false)
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgModifyBackupLocationRequest{Path:t_path, Access:t_access, Enabled:t_enabled}
        _, err := handler.CvModifyBackupLocation(req, strconv.Itoa(d.Get("storagepoolid").(int)), d.Id())
        if err != nil {
            return fmt.Errorf("operation [ModifyBackupLocation] failed, Error %s", err)
        }
    }
    return resourceReadStorage_Disk_Backup_Location(d, m)
}

func resourceDeleteStorage_Disk_Backup_Location(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}
    _, err := handler.CvDeleteBackupLocation(strconv.Itoa(d.Get("storagepoolid").(int)), d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteBackupLocation] failed, Error %s", err)
    }
    return nil
}

func build_storage_disk_backup_location_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_storage_disk_backup_location_msgusernamepassword(d *schema.ResourceData, r []interface{}) *handler.MsgUserNamePassword {
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
