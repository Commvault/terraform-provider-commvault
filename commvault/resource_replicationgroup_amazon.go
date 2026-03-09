package commvault

import (
    "fmt"
    "strconv"
    "strings"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceReplicationGroup_Amazon() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateReplicationGroup_Amazon,
        Read:   resourceReadReplicationGroup_Amazon,
        Update: resourceUpdateReplicationGroup_Amazon,
        Delete: resourceDeleteReplicationGroup_Amazon,

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
                        "volumetype": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "ID of the volume type that will be used on the destination instance. Default value is Auto",
                        },
                        "guestcredentials": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "Faster performance can be achieved during replication by using the VSA access node running on the Amazon instance. Applicable only if the source VM is Windows and the source hypervisor is not Amazon vendor.",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "password": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "The password of the user in Base64 format",
                                    },
                                    "computername": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "The name of the computer",
                                    },
                                    "username": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "The user to be used to access the computer",
                                    },
                                },
                            },
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
                        "instancename": {
                            Type:        schema.TypeString,
                            Required:    true,
                            Description: "Display name for the destination instance",
                        },
                        "regionname": {
                            Type:        schema.TypeString,
                            Required:    true,
                            Description: "The name of the AWS region where the destination instance will reside",
                        },
                        "instancetype": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "The available CPU cores and memory to be used on the destination instance. Default value is Auto",
                        },
                        "securitygroups": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "The security group to be applied to the destination instance that acts as a virtual firewall. Default value is Auto",
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
                        "encryptionkey": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "The ID of the encryption key that will be used to encrypt the data of the desination instance. Default value is Auto",
                        },
                        "availabilityzone": {
                            Type:        schema.TypeString,
                            Required:    true,
                            Description: "The name of AWS zone where the destination instance will reside",
                        },
                        "network": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: "Network info",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "subnetid": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "The ID of the subnet",
                                    },
                                    "vpc": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "The ID of the VPC",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "The name of the network",
                                    },
                                    "privateipaddress": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: "The private IP address of the network",
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "advancedoptions": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Advanced options for Amazon replication group",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "validatedestinationvm": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Validates that the destination VM is bootable by powering it on and then powering off",
                        },
                        "continueonfailure": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "If true, the operation will continue for the remaining VMs even if the failover operation fails for the current VM",
                        },
                        "transportmode": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "transport mode based on environment. Values are case sensitive [AUTO, SAN, HOT_ADD, NAS, NBD_SSL, NBD]",
                        },
                        "deployvmwhenfailover": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "If true, the destination VM will be created during failover operation",
                        },
                        "failoverdelay": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "The delay (in minutes) between performing operations on entities of different priorities",
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
                Description: "[Amazon]",
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

func resourceCreateReplicationGroup_Amazon(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/ReplicationGroup
    var response_id = strconv.Itoa(0)
    var t_recoverytarget *handler.MsgIdName
    if val, ok := d.GetOk("recoverytarget"); ok {
        t_recoverytarget = build_replicationgroup_amazon_msgidname(d, val.([]interface{}))
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
        t_storage = build_replicationgroup_amazon_msgstoragecopycreateset_array(d, val.(*schema.Set).List())
    }
    var t_sourcehypervisor *handler.MsgIdName
    if val, ok := d.GetOk("sourcehypervisor"); ok {
        t_sourcehypervisor = build_replicationgroup_amazon_msgidname(d, val.([]interface{}))
    }
    var t_vms []handler.MsgNameGUIDSet
    if val, ok := d.GetOk("vms"); ok {
        t_vms = build_replicationgroup_amazon_msgnameguidset_array(d, val.(*schema.Set).List())
    }
    var t_overridereplicationoptions []handler.MsgOverrideReplicationOptionsAmazonCreateSet
    if val, ok := d.GetOk("overridereplicationoptions"); ok {
        t_overridereplicationoptions = build_replicationgroup_amazon_msgoverridereplicationoptionsamazoncreateset_array(d, val.(*schema.Set).List())
    }
    var t_advancedoptions *handler.MsgReplicationGroupAdvOptionsAmazon
    if val, ok := d.GetOk("advancedoptions"); ok {
        t_advancedoptions = build_replicationgroup_amazon_msgreplicationgroupadvoptionsamazon(d, val.([]interface{}))
    }
    var t_destvendor *string
    if val, ok := d.GetOk("destvendor"); ok {
        t_destvendor = handler.ToStringValue(val, false)
    }
    var req = handler.MsgCreateReplicationGroupAmazonRequest{RecoveryTarget:t_recoverytarget, FrequencyInMinutes:t_frequencyinminutes, Name:t_name, Storage:t_storage, SourceHypervisor:t_sourcehypervisor, Vms:t_vms, OverrideReplicationOptions:t_overridereplicationoptions, AdvancedOptions:t_advancedoptions, DestVendor:t_destvendor}
    resp, err := handler.CvCreateReplicationGroupAmazon(req)
    if err != nil {
        return fmt.Errorf("operation [CreateReplicationGroupAmazon] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateReplicationGroupAmazon] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateReplicationGroup_Amazon(d, m)
    }
}

func resourceReadReplicationGroup_Amazon(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/ReplicationGroup/{replicationGroupId}
    resp, err := handler.CvgetReplicationGroupDetailsAmazon(d.Id())
    if err != nil {
        if strings.Contains(err.Error(), "status: 404") {
            handler.LogEntry("debug", "entity not present, removing from state")
            d.SetId("")
            return nil
        }
        return fmt.Errorf("operation [getReplicationGroupDetailsAmazon] failed, Error %s", err)
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    if rtn, ok := serialize_replicationgroup_amazon_msgreplicationgroupstoragelist(d, resp.Storage); ok {
        d.Set("storage", rtn)
    } else {
        d.Set("storage", make([]map[string]interface{}, 0))
    }
    return nil
}

func resourceUpdateReplicationGroup_Amazon(d *schema.ResourceData, m interface{}) error {
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
        t_advancedoptions = build_replicationgroup_amazon_msgreplicationgroupadvancedoptions(d, val.([]interface{}))
    }
    var req = handler.MsgModifyReplicationGroupRequest{NewName:t_newname, Enable:t_enable, Disable:t_disable, AdvancedOptions:t_advancedoptions}
    _, err := handler.CvModifyReplicationGroup(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [ModifyReplicationGroup] failed, Error %s", err)
    }
    return resourceReadReplicationGroup_Amazon(d, m)
}

func resourceCreateUpdateReplicationGroup_Amazon(d *schema.ResourceData, m interface{}) error {
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
    return resourceReadReplicationGroup_Amazon(d, m)
}

func resourceDeleteReplicationGroup_Amazon(d *schema.ResourceData, m interface{}) error {
    return nil
}

func build_replicationgroup_amazon_msgreplicationgroupadvancedoptions(d *schema.ResourceData, r []interface{}) *handler.MsgReplicationGroupAdvancedOptions {
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
            t_script = build_replicationgroup_amazon_msgreplicationgroupscript(d, val.([]interface{}))
        }
        return &handler.MsgReplicationGroupAdvancedOptions{DelayBetweenPriorityMachines:t_delaybetweenprioritymachines, ContinueOnFailure:t_continueonfailure, Script:t_script}
    } else {
        return nil
    }
}

func build_replicationgroup_amazon_msgreplicationgroupscript(d *schema.ResourceData, r []interface{}) *handler.MsgReplicationGroupScript {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_prefailover *handler.MsgDROperationScript
        if val, ok := tmp["prefailover"]; ok {
            t_prefailover = build_replicationgroup_amazon_msgdroperationscript(d, val.([]interface{}))
        }
        var t_postfailover *handler.MsgDROperationScript
        if val, ok := tmp["postfailover"]; ok {
            t_postfailover = build_replicationgroup_amazon_msgdroperationscript(d, val.([]interface{}))
        }
        var t_prefailback *handler.MsgDROperationScript
        if val, ok := tmp["prefailback"]; ok {
            t_prefailback = build_replicationgroup_amazon_msgdroperationscript(d, val.([]interface{}))
        }
        var t_postfailback *handler.MsgDROperationScript
        if val, ok := tmp["postfailback"]; ok {
            t_postfailback = build_replicationgroup_amazon_msgdroperationscript(d, val.([]interface{}))
        }
        return &handler.MsgReplicationGroupScript{PreFailover:t_prefailover, PostFailover:t_postfailover, PreFailback:t_prefailback, PostFailback:t_postfailback}
    } else {
        return nil
    }
}

func build_replicationgroup_amazon_msgdroperationscript(d *schema.ResourceData, r []interface{}) *handler.MsgDROperationScript {
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
            t_guestcredentials = build_replicationgroup_amazon_msgguestcredentialinfo(d, val.([]interface{}))
        }
        var t_scriptcredentials *handler.MsgguestCredentialInfo
        if val, ok := tmp["scriptcredentials"]; ok {
            t_scriptcredentials = build_replicationgroup_amazon_msgguestcredentialinfo(d, val.([]interface{}))
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

func build_replicationgroup_amazon_msgguestcredentialinfo(d *schema.ResourceData, r []interface{}) *handler.MsgguestCredentialInfo {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_credentials *handler.MsgUserNamePassword
        if val, ok := tmp["credentials"]; ok {
            t_credentials = build_replicationgroup_amazon_msgusernamepassword(d, val.([]interface{}))
        }
        var t_savedcredentials *handler.MsgIdName
        if val, ok := tmp["savedcredentials"]; ok {
            t_savedcredentials = build_replicationgroup_amazon_msgidname(d, val.([]interface{}))
        }
        return &handler.MsgguestCredentialInfo{Credentials:t_credentials, SavedCredentials:t_savedcredentials}
    } else {
        return nil
    }
}

func build_replicationgroup_amazon_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_replicationgroup_amazon_msgusernamepassword(d *schema.ResourceData, r []interface{}) *handler.MsgUserNamePassword {
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

func build_replicationgroup_amazon_msgreplicationgroupadvoptionsamazon(d *schema.ResourceData, r []interface{}) *handler.MsgReplicationGroupAdvOptionsAmazon {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_validatedestinationvm *bool
        if val, ok := tmp["validatedestinationvm"]; ok {
            t_validatedestinationvm = handler.ToBooleanValue(val, true)
        }
        var t_continueonfailure *bool
        if val, ok := tmp["continueonfailure"]; ok {
            t_continueonfailure = handler.ToBooleanValue(val, true)
        }
        var t_transportmode *string
        if val, ok := tmp["transportmode"]; ok {
            t_transportmode = handler.ToStringValue(val, true)
        }
        var t_deployvmwhenfailover *bool
        if val, ok := tmp["deployvmwhenfailover"]; ok {
            t_deployvmwhenfailover = handler.ToBooleanValue(val, true)
        }
        var t_failoverdelay *int
        if val, ok := tmp["failoverdelay"]; ok {
            t_failoverdelay = handler.ToIntValue(val, true)
        }
        var t_unconditionaloverwrite *bool
        if val, ok := tmp["unconditionaloverwrite"]; ok {
            t_unconditionaloverwrite = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgReplicationGroupAdvOptionsAmazon{ValidateDestinationVM:t_validatedestinationvm, ContinueOnFailure:t_continueonfailure, TransportMode:t_transportmode, DeployVmWhenFailover:t_deployvmwhenfailover, FailoverDelay:t_failoverdelay, UnconditionalOverwrite:t_unconditionaloverwrite}
    } else {
        return nil
    }
}

func build_replicationgroup_amazon_msgoverridereplicationoptionsamazoncreateset_array(d *schema.ResourceData, r []interface{}) []handler.MsgOverrideReplicationOptionsAmazonCreateSet {
    if r != nil {
        tmp := make([]handler.MsgOverrideReplicationOptionsAmazonCreateSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_volumetype *string
            if val, ok := raw_a["volumetype"]; ok {
                t_volumetype = handler.ToStringValue(val, true)
            }
            var t_guestcredentials *handler.MsgGuestCredentialsCreate
            if val, ok := raw_a["guestcredentials"]; ok {
                t_guestcredentials = build_replicationgroup_amazon_msgguestcredentialscreate(d, val.([]interface{}))
            }
            var t_sourcevm *handler.MsgNameGUID
            if val, ok := raw_a["sourcevm"]; ok {
                t_sourcevm = build_replicationgroup_amazon_msgnameguid(d, val.([]interface{}))
            }
            var t_instancename *string
            if val, ok := raw_a["instancename"]; ok {
                t_instancename = handler.ToStringValue(val, true)
            }
            var t_regionname *string
            if val, ok := raw_a["regionname"]; ok {
                t_regionname = handler.ToStringValue(val, true)
            }
            var t_instancetype *string
            if val, ok := raw_a["instancetype"]; ok {
                t_instancetype = handler.ToStringValue(val, true)
            }
            var t_securitygroups []handler.MsgSecurityGroupSet
            if val, ok := raw_a["securitygroups"]; ok {
                t_securitygroups = build_replicationgroup_amazon_msgsecuritygroupset_array(d, val.(*schema.Set).List())
            }
            var t_encryptionkey *string
            if val, ok := raw_a["encryptionkey"]; ok {
                t_encryptionkey = handler.ToStringValue(val, true)
            }
            var t_availabilityzone *string
            if val, ok := raw_a["availabilityzone"]; ok {
                t_availabilityzone = handler.ToStringValue(val, true)
            }
            var t_network *handler.MsgNetworkVPCSubnet
            if val, ok := raw_a["network"]; ok {
                t_network = build_replicationgroup_amazon_msgnetworkvpcsubnet(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgOverrideReplicationOptionsAmazonCreateSet{VolumeType:t_volumetype, GuestCredentials:t_guestcredentials, SourceVM:t_sourcevm, InstanceName:t_instancename, RegionName:t_regionname, InstanceType:t_instancetype, SecurityGroups:t_securitygroups, EncryptionKey:t_encryptionkey, AvailabilityZone:t_availabilityzone, Network:t_network}
        }
        return tmp
    } else {
        return nil
    }
}

func build_replicationgroup_amazon_msgnetworkvpcsubnet(d *schema.ResourceData, r []interface{}) *handler.MsgNetworkVPCSubnet {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_subnetid *string
        if val, ok := tmp["subnetid"]; ok {
            t_subnetid = handler.ToStringValue(val, true)
        }
        var t_vpc *string
        if val, ok := tmp["vpc"]; ok {
            t_vpc = handler.ToStringValue(val, true)
        }
        var t_name *string
        if val, ok := tmp["name"]; ok {
            t_name = handler.ToStringValue(val, true)
        }
        var t_privateipaddress *string
        if val, ok := tmp["privateipaddress"]; ok {
            t_privateipaddress = handler.ToStringValue(val, true)
        }
        return &handler.MsgNetworkVPCSubnet{SubnetId:t_subnetid, Vpc:t_vpc, Name:t_name, PrivateIPaddress:t_privateipaddress}
    } else {
        return nil
    }
}

func build_replicationgroup_amazon_msgsecuritygroupset_array(d *schema.ResourceData, r []interface{}) []handler.MsgSecurityGroupSet {
    if r != nil {
        tmp := make([]handler.MsgSecurityGroupSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_name *string
            if val, ok := raw_a["name"]; ok {
                t_name = handler.ToStringValue(val, true)
            }
            var t_id *string
            if val, ok := raw_a["id"]; ok {
                t_id = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgSecurityGroupSet{Name:t_name, Id:t_id}
        }
        return tmp
    } else {
        return nil
    }
}

func build_replicationgroup_amazon_msgnameguid(d *schema.ResourceData, r []interface{}) *handler.MsgNameGUID {
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

func build_replicationgroup_amazon_msgguestcredentialscreate(d *schema.ResourceData, r []interface{}) *handler.MsgGuestCredentialsCreate {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_password *string
        if val, ok := tmp["password"]; ok {
            t_password = handler.ToStringValue(val, true)
        }
        var t_computername *string
        if val, ok := tmp["computername"]; ok {
            t_computername = handler.ToStringValue(val, true)
        }
        var t_username *string
        if val, ok := tmp["username"]; ok {
            t_username = handler.ToStringValue(val, true)
        }
        return &handler.MsgGuestCredentialsCreate{Password:t_password, ComputerName:t_computername, UserName:t_username}
    } else {
        return nil
    }
}

func build_replicationgroup_amazon_msgnameguidset_array(d *schema.ResourceData, r []interface{}) []handler.MsgNameGUIDSet {
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

func build_replicationgroup_amazon_msgstoragecopycreateset_array(d *schema.ResourceData, r []interface{}) []handler.MsgStorageCopyCreateSet {
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
                t_storagepool = build_replicationgroup_amazon_msgidname(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgStorageCopyCreateSet{IsDedupe:t_isdedupe, StoragePool:t_storagepool}
        }
        return tmp
    } else {
        return nil
    }
}

func serialize_replicationgroup_amazon_msgreplicationgroupstoragelist(d *schema.ResourceData, data *handler.MsgReplicationGroupStorageList) ([]map[string]interface{}, bool) {
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
