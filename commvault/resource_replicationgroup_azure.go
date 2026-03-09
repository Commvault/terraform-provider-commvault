package commvault

import (
    "fmt"
    "strconv"
    "strings"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceReplicationGroup_Azure() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateReplicationGroup_Azure,
        Read:   resourceReadReplicationGroup_Azure,
        Update: resourceUpdateReplicationGroup_Azure,
        Delete: resourceDeleteReplicationGroup_Azure,

        Schema: map[string]*schema.Schema{
            "recoverytarget": {
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
            "frequencyinminutes": {
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: "The field denotes the frequency of replication.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "The name of the replication group being created",
            },
            "storage": {
                Type:        schema.TypeSet,
                Required:    true,
                Description: "The primary and an optional secondary storage that will be used for storing the source VM data for replication. The secondary storage if provided, will be the default source for replication.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "isdedupe": {
                            Type:        schema.TypeString,
                            Required:    true,
                            Description: "Denotes if this is a deduplication copy",
                        },
                        "storagepool": {
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
                    },
                },
            },
            "sourcehypervisor": {
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
            "vms": {
                Type:        schema.TypeSet,
                Required:    true,
                Description: "A list of name and GUID of all the virtual machines that have to be replicated",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "guid": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "",
                        },
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "",
                        },
                    },
                },
            },
            "overridereplicationoptions": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "The replication options per instance, if provided, will override the replication options from the target. For the instances not in this list, the options are applied from the target.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "resourcegroup": {
                            Type:        schema.TypeString,
                            Required:    true,
                            Description: "The resource group to be used for the destination VM",
                        },
                        "sourcevm": {
                            Type:        schema.TypeList,
                            Required:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "guid": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "vmsizeid": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "The id of the vm size to be applied to the destination VM. Default value is Auto",
                        },
                        "publicipaddress": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Public IP address of the network",
                        },
                        "disktypeid": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "The id of the disk type to be used for the destination VM. Default value is Auto",
                        },
                        "vmdisplayname": {
                            Type:        schema.TypeString,
                            Required:    true,
                            Description: "Display name of destination VM",
                        },
                        "securitygroup": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "The id and the name of the security group",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "security group name",
                                    },
                                    "id": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "Security group id",
                                    },
                                },
                            },
                        },
                        "storageaccount": {
                            Type:        schema.TypeString,
                            Required:    true,
                            Description: "The name of the storage account to be used on the destination VM",
                        },
                        "privateipaddress": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Private IP address of the network",
                        },
                        "virtualnetwork": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "The network to be used on the destination VM. Default value is Auto",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "subnetid": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "The ID of the subnet",
                                    },
                                    "network": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "The name of the network",
                                    },
                                },
                            },
                        },
                        "createpublicip": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Creates a public IP on the network",
                        },
                        "publicipaddressid": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Public IP address id of the network",
                        },
                        "restoreasmanagedvm": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Creates the destination as a managed VM",
                        },
                        "region": {
                            Type:        schema.TypeString,
                            Required:    true,
                            Description: "The name of the region where the destination VM will reside",
                        },
                    },
                },
            },
            "advancedoptions": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Advanced options for Azure replication group",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "deployvmonlyduringfailover": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Select this option to deploy a VM only when a failover operation is requested",
                        },
                        "unconditionaloverwrite": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "This will replace the instance at the destination if the instance with the same name already exists.",
                        },
                    },
                },
            },
            "destvendor": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "[Azure]",
            },
            "enable": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Used to enable a replication group",
            },
            "disable": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Used to disable a replication group",
            },
        },
    }
}

func resourceCreateReplicationGroup_Azure(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/ReplicationGroup
    var response_id = strconv.Itoa(0)
    var t_recoverytarget *handler.MsgIdName
    if val, ok := d.GetOk("recoverytarget"); ok {
        t_recoverytarget = build_replicationgroup_azure_msgidname(d, val.([]interface{}))
    }
    var t_frequencyinminutes *int
    if val, ok := d.GetOk("frequencyinminutes"); ok {
        t_frequencyinminutes = handler.ToIntValue(val, false)
    }
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_storage []handler.MsgStorageCopyCreateSet
    if val, ok := d.GetOk("storage"); ok {
        t_storage = build_replicationgroup_azure_msgstoragecopycreateset_array(d, val.(*schema.Set).List())
    }
    var t_sourcehypervisor *handler.MsgIdName
    if val, ok := d.GetOk("sourcehypervisor"); ok {
        t_sourcehypervisor = build_replicationgroup_azure_msgidname(d, val.([]interface{}))
    }
    var t_vms []handler.MsgNameGUIDSet
    if val, ok := d.GetOk("vms"); ok {
        t_vms = build_replicationgroup_azure_msgnameguidset_array(d, val.(*schema.Set).List())
    }
    var t_overridereplicationoptions []handler.MsgOverrideReplicationOptionsAzureCreateSet
    if val, ok := d.GetOk("overridereplicationoptions"); ok {
        t_overridereplicationoptions = build_replicationgroup_azure_msgoverridereplicationoptionsazurecreateset_array(d, val.(*schema.Set).List())
    }
    var t_advancedoptions *handler.MsgReplicationGroupAdvOptionsAzure
    if val, ok := d.GetOk("advancedoptions"); ok {
        t_advancedoptions = build_replicationgroup_azure_msgreplicationgroupadvoptionsazure(d, val.([]interface{}))
    }
    var t_destvendor *string
    if val, ok := d.GetOk("destvendor"); ok {
        t_destvendor = handler.ToStringValue(val, false)
    }
    var req = handler.MsgCreateReplicationGroupAzureRequest{RecoveryTarget:t_recoverytarget, FrequencyInMinutes:t_frequencyinminutes, Name:t_name, Storage:t_storage, SourceHypervisor:t_sourcehypervisor, Vms:t_vms, OverrideReplicationOptions:t_overridereplicationoptions, AdvancedOptions:t_advancedoptions, DestVendor:t_destvendor}
    resp, err := handler.CvCreateReplicationGroupAzure(req)
    if err != nil {
        return fmt.Errorf("operation [CreateReplicationGroupAzure] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateReplicationGroupAzure] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateReplicationGroup_Azure(d, m)
    }
}

func resourceReadReplicationGroup_Azure(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/ReplicationGroup/{replicationGroupId}
    resp, err := handler.CvgetReplicationGroupDetailsAzure(d.Id())
    if err != nil {
        if strings.Contains(err.Error(), "status: 404") {
            handler.LogEntry("debug", "entity not present, removing from state")
            d.SetId("")
            return nil
        }
        return fmt.Errorf("operation [getReplicationGroupDetailsAzure] failed, Error %s", err)
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    if rtn, ok := serialize_replicationgroup_azure_msgreplicationgroupstoragelist(d, resp.Storage); ok {
        d.Set("storage", rtn)
    } else {
        d.Set("storage", make([]map[string]interface{}, 0))
    }
    return nil
}

func resourceUpdateReplicationGroup_Azure(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/ReplicationGroup/{replicationGroupId}
    var t_newname *string
    if d.HasChange("name") {
        val := d.Get("name")
        t_newname = handler.ToStringValue(val, false)
    }
    var t_enable *bool
    if d.HasChange("enable") {
        val := d.Get("enable")
        t_enable = handler.ToBooleanValue(val, false)
    }
    var t_disable *bool
    if d.HasChange("disable") {
        val := d.Get("disable")
        t_disable = handler.ToBooleanValue(val, false)
    }
    var t_advancedoptions *handler.MsgReplicationGroupAdvancedOptions
    if d.HasChange("advancedoptions") {
        val := d.Get("advancedoptions")
        t_advancedoptions = build_replicationgroup_azure_msgreplicationgroupadvancedoptions(d, val.([]interface{}))
    }
    var req = handler.MsgModifyReplicationGroupRequest{NewName:t_newname, Enable:t_enable, Disable:t_disable, AdvancedOptions:t_advancedoptions}
    _, err := handler.CvModifyReplicationGroup(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [ModifyReplicationGroup] failed, Error %s", err)
    }
    return resourceReadReplicationGroup_Azure(d, m)
}

func resourceCreateUpdateReplicationGroup_Azure(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/ReplicationGroup/{replicationGroupId}
    var execUpdate bool = false
    var t_enable *bool
    if val, ok := d.GetOk("enable"); ok {
        t_enable = handler.ToBooleanValue(val, false)
        execUpdate = true
    }
    var t_disable *bool
    if val, ok := d.GetOk("disable"); ok {
        t_disable = handler.ToBooleanValue(val, false)
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgModifyReplicationGroupRequest{Enable:t_enable, Disable:t_disable}
        _, err := handler.CvModifyReplicationGroup(req, d.Id())
        if err != nil {
            return fmt.Errorf("operation [ModifyReplicationGroup] failed, Error %s", err)
        }
    }
    return resourceReadReplicationGroup_Azure(d, m)
}

func resourceDeleteReplicationGroup_Azure(d *schema.ResourceData, m interface{}) error {
    return nil
}

func build_replicationgroup_azure_msgreplicationgroupadvancedoptions(d *schema.ResourceData, r []interface{}) *handler.MsgReplicationGroupAdvancedOptions {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_delaybetweenprioritymachines *int
        if val, ok := tmp["delaybetweenprioritymachines"]; ok {
            t_delaybetweenprioritymachines = handler.ToIntValue(val, true)
        }
        var t_continueonfailure *bool
        if val, ok := tmp["continueonfailure"]; ok {
            t_continueonfailure = handler.ToBooleanValue(val, true)
        }
        var t_script *handler.MsgReplicationGroupScript
        if val, ok := tmp["script"]; ok {
            t_script = build_replicationgroup_azure_msgreplicationgroupscript(d, val.([]interface{}))
        }
        return &handler.MsgReplicationGroupAdvancedOptions{DelayBetweenPriorityMachines:t_delaybetweenprioritymachines, ContinueOnFailure:t_continueonfailure, Script:t_script}
    } else {
        return nil
    }
}

func build_replicationgroup_azure_msgreplicationgroupscript(d *schema.ResourceData, r []interface{}) *handler.MsgReplicationGroupScript {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_prefailover *handler.MsgDROperationScript
        if val, ok := tmp["prefailover"]; ok {
            t_prefailover = build_replicationgroup_azure_msgdroperationscript(d, val.([]interface{}))
        }
        var t_postfailover *handler.MsgDROperationScript
        if val, ok := tmp["postfailover"]; ok {
            t_postfailover = build_replicationgroup_azure_msgdroperationscript(d, val.([]interface{}))
        }
        var t_prefailback *handler.MsgDROperationScript
        if val, ok := tmp["prefailback"]; ok {
            t_prefailback = build_replicationgroup_azure_msgdroperationscript(d, val.([]interface{}))
        }
        var t_postfailback *handler.MsgDROperationScript
        if val, ok := tmp["postfailback"]; ok {
            t_postfailback = build_replicationgroup_azure_msgdroperationscript(d, val.([]interface{}))
        }
        return &handler.MsgReplicationGroupScript{PreFailover:t_prefailover, PostFailover:t_postfailover, PreFailback:t_prefailback, PostFailback:t_postfailback}
    } else {
        return nil
    }
}

func build_replicationgroup_azure_msgdroperationscript(d *schema.ResourceData, r []interface{}) *handler.MsgDROperationScript {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_scriptcontent *string
        if val, ok := tmp["scriptcontent"]; ok {
            t_scriptcontent = handler.ToStringValue(val, true)
        }
        var t_path *string
        if val, ok := tmp["path"]; ok {
            t_path = handler.ToStringValue(val, true)
        }
        var t_guestcredentials *handler.MsgguestCredentialInfo
        if val, ok := tmp["guestcredentials"]; ok {
            t_guestcredentials = build_replicationgroup_azure_msgguestcredentialinfo(d, val.([]interface{}))
        }
        var t_scriptcredentials *handler.MsgguestCredentialInfo
        if val, ok := tmp["scriptcredentials"]; ok {
            t_scriptcredentials = build_replicationgroup_azure_msgguestcredentialinfo(d, val.([]interface{}))
        }
        var t_ostype *string
        if val, ok := tmp["ostype"]; ok {
            t_ostype = handler.ToStringValue(val, true)
        }
        var t_guid *string
        if val, ok := tmp["guid"]; ok {
            t_guid = handler.ToStringValue(val, true)
        }
        var t_scriptname *string
        if val, ok := tmp["scriptname"]; ok {
            t_scriptname = handler.ToStringValue(val, true)
        }
        var t_reset *bool
        if val, ok := tmp["reset"]; ok {
            t_reset = handler.ToBooleanValue(val, true)
        }
        var t_type *string
        if val, ok := tmp["type"]; ok {
            t_type = handler.ToStringValue(val, true)
        }
        return &handler.MsgDROperationScript{ScriptContent:t_scriptcontent, Path:t_path, GuestCredentials:t_guestcredentials, ScriptCredentials:t_scriptcredentials, OsType:t_ostype, Guid:t_guid, ScriptName:t_scriptname, Reset:t_reset, Type:t_type}
    } else {
        return nil
    }
}

func build_replicationgroup_azure_msgguestcredentialinfo(d *schema.ResourceData, r []interface{}) *handler.MsgguestCredentialInfo {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_credentials *handler.MsgUserNamePassword
        if val, ok := tmp["credentials"]; ok {
            t_credentials = build_replicationgroup_azure_msgusernamepassword(d, val.([]interface{}))
        }
        var t_savedcredentials *handler.MsgIdName
        if val, ok := tmp["savedcredentials"]; ok {
            t_savedcredentials = build_replicationgroup_azure_msgidname(d, val.([]interface{}))
        }
        return &handler.MsgguestCredentialInfo{Credentials:t_credentials, SavedCredentials:t_savedcredentials}
    } else {
        return nil
    }
}

func build_replicationgroup_azure_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_replicationgroup_azure_msgusernamepassword(d *schema.ResourceData, r []interface{}) *handler.MsgUserNamePassword {
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

func build_replicationgroup_azure_msgreplicationgroupadvoptionsazure(d *schema.ResourceData, r []interface{}) *handler.MsgReplicationGroupAdvOptionsAzure {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_deployvmonlyduringfailover *bool
        if val, ok := tmp["deployvmonlyduringfailover"]; ok {
            t_deployvmonlyduringfailover = handler.ToBooleanValue(val, true)
        }
        var t_unconditionaloverwrite *bool
        if val, ok := tmp["unconditionaloverwrite"]; ok {
            t_unconditionaloverwrite = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgReplicationGroupAdvOptionsAzure{DeployVmOnlyDuringfailover:t_deployvmonlyduringfailover, UnconditionalOverwrite:t_unconditionaloverwrite}
    } else {
        return nil
    }
}

func build_replicationgroup_azure_msgoverridereplicationoptionsazurecreateset_array(d *schema.ResourceData, r []interface{}) []handler.MsgOverrideReplicationOptionsAzureCreateSet {
    if r != nil {
        tmp := make([]handler.MsgOverrideReplicationOptionsAzureCreateSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_resourcegroup *string
            if val, ok := raw_a["resourcegroup"]; ok {
                t_resourcegroup = handler.ToStringValue(val, true)
            }
            var t_sourcevm *handler.MsgNameGUID
            if val, ok := raw_a["sourcevm"]; ok {
                t_sourcevm = build_replicationgroup_azure_msgnameguid(d, val.([]interface{}))
            }
            var t_vmsizeid *string
            if val, ok := raw_a["vmsizeid"]; ok {
                t_vmsizeid = handler.ToStringValue(val, true)
            }
            var t_publicipaddress *string
            if val, ok := raw_a["publicipaddress"]; ok {
                t_publicipaddress = handler.ToStringValue(val, true)
            }
            var t_disktypeid *string
            if val, ok := raw_a["disktypeid"]; ok {
                t_disktypeid = handler.ToStringValue(val, true)
            }
            var t_vmdisplayname *string
            if val, ok := raw_a["vmdisplayname"]; ok {
                t_vmdisplayname = handler.ToStringValue(val, true)
            }
            var t_securitygroup *handler.MsgSecurityGroup
            if val, ok := raw_a["securitygroup"]; ok {
                t_securitygroup = build_replicationgroup_azure_msgsecuritygroup(d, val.([]interface{}))
            }
            var t_storageaccount *string
            if val, ok := raw_a["storageaccount"]; ok {
                t_storageaccount = handler.ToStringValue(val, true)
            }
            var t_privateipaddress *string
            if val, ok := raw_a["privateipaddress"]; ok {
                t_privateipaddress = handler.ToStringValue(val, true)
            }
            var t_virtualnetwork []handler.MsgNetworkSubnetSet
            if val, ok := raw_a["virtualnetwork"]; ok {
                t_virtualnetwork = build_replicationgroup_azure_msgnetworksubnetset_array(d, val.(*schema.Set).List())
            }
            var t_createpublicip *bool
            if val, ok := raw_a["createpublicip"]; ok {
                t_createpublicip = handler.ToBooleanValue(val, true)
            }
            var t_publicipaddressid *string
            if val, ok := raw_a["publicipaddressid"]; ok {
                t_publicipaddressid = handler.ToStringValue(val, true)
            }
            var t_restoreasmanagedvm *bool
            if val, ok := raw_a["restoreasmanagedvm"]; ok {
                t_restoreasmanagedvm = handler.ToBooleanValue(val, true)
            }
            var t_region *string
            if val, ok := raw_a["region"]; ok {
                t_region = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgOverrideReplicationOptionsAzureCreateSet{ResourceGroup:t_resourcegroup, SourceVM:t_sourcevm, VmSizeId:t_vmsizeid, PublicIPaddress:t_publicipaddress, DiskTypeId:t_disktypeid, VmDisplayName:t_vmdisplayname, SecurityGroup:t_securitygroup, StorageAccount:t_storageaccount, PrivateIPaddress:t_privateipaddress, VirtualNetwork:t_virtualnetwork, CreatePublicIP:t_createpublicip, PublicIPaddressID:t_publicipaddressid, RestoreAsManagedVM:t_restoreasmanagedvm, Region:t_region}
        }
        return tmp
    } else {
        return nil
    }
}

func build_replicationgroup_azure_msgnetworksubnetset_array(d *schema.ResourceData, r []interface{}) []handler.MsgNetworkSubnetSet {
    if r != nil {
        tmp := make([]handler.MsgNetworkSubnetSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_subnetid *string
            if val, ok := raw_a["subnetid"]; ok {
                t_subnetid = handler.ToStringValue(val, true)
            }
            var t_network *string
            if val, ok := raw_a["network"]; ok {
                t_network = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgNetworkSubnetSet{SubnetId:t_subnetid, Network:t_network}
        }
        return tmp
    } else {
        return nil
    }
}

func build_replicationgroup_azure_msgsecuritygroup(d *schema.ResourceData, r []interface{}) *handler.MsgSecurityGroup {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_name *string
        if val, ok := tmp["name"]; ok {
            t_name = handler.ToStringValue(val, true)
        }
        var t_id *string
        if val, ok := tmp["id"]; ok {
            t_id = handler.ToStringValue(val, true)
        }
        return &handler.MsgSecurityGroup{Name:t_name, Id:t_id}
    } else {
        return nil
    }
}

func build_replicationgroup_azure_msgnameguid(d *schema.ResourceData, r []interface{}) *handler.MsgNameGUID {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_guid *string
        if val, ok := tmp["guid"]; ok {
            t_guid = handler.ToStringValue(val, true)
        }
        var t_name *string
        if val, ok := tmp["name"]; ok {
            t_name = handler.ToStringValue(val, true)
        }
        return &handler.MsgNameGUID{GUID:t_guid, Name:t_name}
    } else {
        return nil
    }
}

func build_replicationgroup_azure_msgnameguidset_array(d *schema.ResourceData, r []interface{}) []handler.MsgNameGUIDSet {
    if r != nil {
        tmp := make([]handler.MsgNameGUIDSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_guid *string
            if val, ok := raw_a["guid"]; ok {
                t_guid = handler.ToStringValue(val, true)
            }
            var t_name *string
            if val, ok := raw_a["name"]; ok {
                t_name = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgNameGUIDSet{GUID:t_guid, Name:t_name}
        }
        return tmp
    } else {
        return nil
    }
}

func build_replicationgroup_azure_msgstoragecopycreateset_array(d *schema.ResourceData, r []interface{}) []handler.MsgStorageCopyCreateSet {
    if r != nil {
        tmp := make([]handler.MsgStorageCopyCreateSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_isdedupe *bool
            if val, ok := raw_a["isdedupe"]; ok {
                t_isdedupe = handler.ToBooleanValue(val, true)
            }
            var t_storagepool *handler.MsgIdName
            if val, ok := raw_a["storagepool"]; ok {
                t_storagepool = build_replicationgroup_azure_msgidname(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgStorageCopyCreateSet{IsDedupe:t_isdedupe, StoragePool:t_storagepool}
        }
        return tmp
    } else {
        return nil
    }
}

func serialize_replicationgroup_azure_msgreplicationgroupstoragelist(d *schema.ResourceData, data *handler.MsgReplicationGroupStorageList) ([]map[string]interface{}, bool) {
    //MsgStorageCopyCreateSet
    //MsgReplicationGroupStorageList
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if added {
        return val, true
    } else {
        return nil, false
    }
}
