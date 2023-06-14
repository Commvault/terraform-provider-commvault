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
            "access": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "The access type for the access path can be either read (writing to path not allowed) or read and write (writing to path allowed). [READ_AND_WRITE, READ]",
            },
            "configuration": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "While adding network access path, please add credentials or saved credentials. If both are provided, credentials will be selected.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enablebackuplocation": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Used to enable or disable backup location",
                        },
                        "disablebackuplocationforfuturebackups": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Used to determine if backup location has to be disabled or enabled for future backups",
                        },
                        "prepareforretirement": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Used to determine if the backup location has to be prepared for retirement",
                        },
                    },
                },
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
    resp, err := handler.CvGetBackupLocationDetails(strconv.Itoa(d.Get("storagepoolid").(int)), d.Id())
    if err != nil {
        return fmt.Errorf("operation [GetBackupLocationDetails] failed, Error %s", err)
    }
    if rtn, ok := serialize_storage_disk_backup_location_msgidname(d, resp.MediaAgent); ok {
        d.Set("mediaagent", rtn)
    } else {
        d.Set("mediaagent", make([]map[string]interface{}, 0))
    }
    if resp.Access != nil {
        d.Set("access", resp.Access)
    }
    if rtn, ok := serialize_storage_disk_backup_location_msgdiskstorageconfiguration(d, resp.Configuration); ok {
        d.Set("configuration", rtn)
    } else {
        d.Set("configuration", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_storage_disk_backup_location_msgcredentialusername(d, resp.Credentials); ok {
        d.Set("credentials", rtn)
    } else {
        d.Set("credentials", make([]map[string]interface{}, 0))
    }
    if resp.BackupLocation != nil {
        d.Set("backuplocation", resp.BackupLocation)
    }
    if rtn, ok := statecopy_storage_disk_backup_location_savedcredentials(d); ok {
        d.Set("savedcredentials", rtn)
    } else {
        d.Set("savedcredentials", make([]map[string]interface{}, 0))
    }
    if resp.Enabled != nil {
        d.Set("enabled", strconv.FormatBool(*resp.Enabled))
    }
    return nil
}

func resourceUpdateStorage_Disk_Backup_Location(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}
    var t_mediaagent *handler.MsgIdName
    if val, ok := d.GetOk("mediaagent"); ok {
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
    var t_configuration *handler.MsgDiskStorageConfiguration
    if d.HasChange("configuration") {
        val := d.Get("configuration")
        t_configuration = build_storage_disk_backup_location_msgdiskstorageconfiguration(d, val.([]interface{}))
    }
    var t_backuplocation *string
    if d.HasChange("backuplocation") {
        val := d.Get("backuplocation")
        t_backuplocation = handler.ToStringValue(val, false)
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
    var req = handler.MsgModifyBackupLocationRequest{MediaAgent:t_mediaagent, Access:t_access, Credentials:t_credentials, Configuration:t_configuration, BackupLocation:t_backuplocation, SavedCredentials:t_savedcredentials, Enabled:t_enabled}
    _, err := handler.CvModifyBackupLocation(req, strconv.Itoa(d.Get("storagepoolid").(int)), d.Id())
    if err != nil {
        return fmt.Errorf("operation [ModifyBackupLocation] failed, Error %s", err)
    }
    return resourceReadStorage_Disk_Backup_Location(d, m)
}

func resourceCreateUpdateStorage_Disk_Backup_Location(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}
    var execUpdate bool = false
    var t_mediaagent *handler.MsgIdName
    if val, ok := d.GetOk("mediaagent"); ok {
        t_mediaagent = build_storage_disk_backup_location_msgidname(d, val.([]interface{}))
    }
    var t_access *string
    if val, ok := d.GetOk("access"); ok {
        t_access = handler.ToStringValue(val, false)
        execUpdate = true
    }
    var t_configuration *handler.MsgDiskStorageConfiguration
    if val, ok := d.GetOk("configuration"); ok {
        t_configuration = build_storage_disk_backup_location_msgdiskstorageconfiguration(d, val.([]interface{}))
        execUpdate = true
    }
    var t_enabled *bool
    if val, ok := d.GetOk("enabled"); ok {
        t_enabled = handler.ToBooleanValue(val, false)
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgModifyBackupLocationRequest{MediaAgent:t_mediaagent, Access:t_access, Configuration:t_configuration, Enabled:t_enabled}
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

func build_storage_disk_backup_location_msgdiskstorageconfiguration(d *schema.ResourceData, r []interface{}) *handler.MsgDiskStorageConfiguration {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_enablebackuplocation *bool
        if val, ok := tmp["enablebackuplocation"]; ok {
            t_enablebackuplocation = handler.ToBooleanValue(val, true)
        }
        var t_disablebackuplocationforfuturebackups *bool
        if val, ok := tmp["disablebackuplocationforfuturebackups"]; ok {
            t_disablebackuplocationforfuturebackups = handler.ToBooleanValue(val, true)
        }
        var t_prepareforretirement *bool
        if val, ok := tmp["prepareforretirement"]; ok {
            t_prepareforretirement = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgDiskStorageConfiguration{EnableBackupLocation:t_enablebackuplocation, DisableBackupLocationforFutureBackups:t_disablebackuplocationforfuturebackups, PrepareForRetirement:t_prepareforretirement}
    } else {
        return nil
    }
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

func statecopy_storage_disk_backup_location_savedcredentials(d *schema.ResourceData) ([]interface{}, bool) {
    //STATE COPY
    var_a := d.Get("savedcredentials").([]interface{})
    if len(var_a) > 0 {
        return var_a, true
    }
    return nil, false
}

func serialize_storage_disk_backup_location_msgcredentialusername(d *schema.ResourceData, data *handler.MsgCredentialUserName) ([]map[string]interface{}, bool) {
    //MsgUserNamePassword
    //MsgCredentialUserName
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
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_storage_disk_backup_location_msgdiskstorageconfiguration(d *schema.ResourceData, data *handler.MsgDiskStorageConfiguration) ([]map[string]interface{}, bool) {
    //MsgDiskStorageConfiguration
    //MsgDiskStorageConfiguration
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.EnableBackupLocation != nil {
        val[0]["enablebackuplocation"] = strconv.FormatBool(*data.EnableBackupLocation)
        added = true
    }
    if data.DisableBackupLocationforFutureBackups != nil {
        val[0]["disablebackuplocationforfuturebackups"] = strconv.FormatBool(*data.DisableBackupLocationforFutureBackups)
        added = true
    }
    if data.PrepareForRetirement != nil {
        val[0]["prepareforretirement"] = strconv.FormatBool(*data.PrepareForRetirement)
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_storage_disk_backup_location_msgidname(d *schema.ResourceData, data *handler.MsgIdName) ([]map[string]interface{}, bool) {
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
