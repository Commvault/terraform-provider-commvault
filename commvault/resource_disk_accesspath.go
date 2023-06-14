package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDisk_AccessPath() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateDisk_AccessPath,
        Read:   resourceReadDisk_AccessPath,
        Update: resourceUpdateDisk_AccessPath,
        Delete: resourceDeleteDisk_AccessPath,

        Schema: map[string]*schema.Schema{
            "storagepoolid": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: "Id of the disk storage pool whose details have to be fetched to add a new access path",
            },
            "backuplocationid": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: "Id of the backup location whose details have to be fetched to add a new access path",
            },
            "mediaagent": {
                Type:        schema.TypeList,
                Required:    true,
                Description: "Can add a list of media agents to be added as the disk access path.",
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
            "access": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "The access type for the access path can be either read (writing to path not allowed) or read and write (writing to path allowed). [READ_AND_WRITE, READ]",
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

func resourceCreateDisk_AccessPath(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}/AccessPath
    var response_id = strconv.Itoa(0)
    var t_mediaagents []handler.MsgIdNameSet
    if val, ok := d.GetOk("mediaagent"); ok {
        t_mediaagents = build_disk_accesspath_msgidnameset_array(d, val.([]interface{}))
    }
    var req = handler.MsgAddMediaAgentRequest{MediaAgents:t_mediaagents}
    resp, err := handler.CvAddMediaAgent(req, strconv.Itoa(d.Get("storagepoolid").(int)), strconv.Itoa(d.Get("backuplocationid").(int)))
    if err != nil {
        return fmt.Errorf("operation [AddMediaAgent] failed, Error %s", err)
    }
    response_id = handler.RetrieveBackupLocationAccessPathId(req, resp, d, m)
    if response_id == "0" {
        return fmt.Errorf("operation [AddMediaAgent] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateDisk_AccessPath(d, m)
    }
}

func resourceReadDisk_AccessPath(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceUpdateDisk_AccessPath(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}/AccessPath/{accessPathId}
    var t_access *string
    if d.HasChange("access") {
        val := d.Get("access")
        t_access = handler.ToStringValue(val, false)
    }
    var t_enabled *bool
    if d.HasChange("enabled") {
        val := d.Get("enabled")
        t_enabled = handler.ToBooleanValue(val, false)
    }
    var req = handler.MsgModifyDiskAccessPathRequest{Access:t_access, Enabled:t_enabled}
    _, err := handler.CvModifyDiskAccessPath(req, strconv.Itoa(d.Get("storagepoolid").(int)), strconv.Itoa(d.Get("backuplocationid").(int)), d.Id())
    if err != nil {
        return fmt.Errorf("operation [ModifyDiskAccessPath] failed, Error %s", err)
    }
    return resourceReadDisk_AccessPath(d, m)
}

func resourceCreateUpdateDisk_AccessPath(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}/AccessPath/{accessPathId}
    var execUpdate bool = false
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
        var req = handler.MsgModifyDiskAccessPathRequest{Access:t_access, Enabled:t_enabled}
        _, err := handler.CvModifyDiskAccessPath(req, strconv.Itoa(d.Get("storagepoolid").(int)), strconv.Itoa(d.Get("backuplocationid").(int)), d.Id())
        if err != nil {
            return fmt.Errorf("operation [ModifyDiskAccessPath] failed, Error %s", err)
        }
    }
    return resourceReadDisk_AccessPath(d, m)
}

func resourceDeleteDisk_AccessPath(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}/AccessPath/{accessPathId}
    _, err := handler.CvDeleteDiskAccessPath(strconv.Itoa(d.Get("storagepoolid").(int)), strconv.Itoa(d.Get("backuplocationid").(int)), d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteDiskAccessPath] failed, Error %s", err)
    }
    return nil
}

func build_disk_accesspath_msgidnameset_array(d *schema.ResourceData, r []interface{}) []handler.MsgIdNameSet {
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
