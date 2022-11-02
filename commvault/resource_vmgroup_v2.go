package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceVMGroup_V2() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateVMGroup_V2,
        Read:   resourceReadVMGroup_V2,
        Update: resourceUpdateVMGroup_V2,
        Delete: resourceDeleteVMGroup_V2,

        Schema: map[string]*schema.Schema{
            "meditech": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "systemname": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Meditech system name",
                        },
                        "listenerip": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Meditech Listener IP of FQDN name",
                        },
                        "useraccount": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "password": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "password to access the network path",
                                    },
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "username to access the network path",
                                    },
                                },
                            },
                        },
                        "listenerport": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Meditech Listener Port",
                        },
                        "mbftimeout": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "MBF timeout (in seconds)",
                        },
                    },
                },
            },
            "hypervisor": &schema.Schema{
                Type:        schema.TypeList,
                Required:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "id": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
            },
            "name": &schema.Schema{
                Type:        schema.TypeString,
                Required:    true,
                Description: "subclient name ",
            },
            "storage": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "id": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
            },
            "enableintellisnap": &schema.Schema{
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: "True if Intellisnap has to be  enabled",
            },
            "plan": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "id": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
            },
            "content": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rulegroups": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "matchrule": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Enum which specifies the whether to match all rules or any of the rules",
                                    },
                                    "rules": &schema.Schema{
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "condition": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Operation type for VM rules/filters",
                                                },
                                                "displayname": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "The display name of the entity to be added",
                                                },
                                                "name": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "name of the VM to be added as content",
                                                },
                                                "guid": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "GUID of the entity to be added as content",
                                                },
                                                "type": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "value": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "value for the few type of VM Content like powerstate",
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "virtualmachines": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "guestcredentialassocid": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Credential association ID given to link entity with credential id.",
                                    },
                                    "guestcredentials": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "password": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "username": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                            },
                                        },
                                    },
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "name of the VM to be added as content",
                                    },
                                    "guid": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "GUID of the VM to be added as content",
                                    },
                                    "type": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "existingcredential": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "credentialid": &schema.Schema{
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "credentialname": &schema.Schema{
                                                    Type:        schema.TypeString,
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
                        "overwrite": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if content in vmgroup has to be overwritten, by default it will append the content",
                        },
                    },
                },
            },
            "activitycontrol": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "restoreactivitycontroloptions": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "timezone": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "id": &schema.Schema{
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                            },
                                        },
                                    },
                                    "enableafterdelay": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if the activity will be enabled after a delay time interval",
                                    },
                                    "delaytime": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Delayed by n Hrs",
                                    },
                                },
                            },
                        },
                        "backupactivitycontroloptions": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "timezone": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "id": &schema.Schema{
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                            },
                                        },
                                    },
                                    "enableafterdelay": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if the activity will be enabled after a delay time interval",
                                    },
                                    "delaytime": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Delayed by n Hrs",
                                    },
                                },
                            },
                        },
                        "enablebackup": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "true if Backup is enabled",
                        },
                        "enablerestore": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "true if Restore is enabled",
                        },
                    },
                },
            },
            "settings": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "autodetectvmowner": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if auto detect VM Owner enabled",
                        },
                        "collectfiledetailsforgranularrecovery": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if metadata collection is enabled. Only applicable for Indexing v1",
                        },
                        "noofreaders": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Number of readers for backup",
                        },
                        "usechangedblocktrackingonvm": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if Changed Block Tracking is enabled",
                        },
                        "jobstarttime": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Start Time for the VM Group Job",
                        },
                        "usevmcheckpointsetting": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if use VM CheckPoint setting is enabled",
                        },
                        "customsnapshotresourcegroup": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Custom snapshot resource group name for Azure",
                        },
                        "regionalsnapshot": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "True when snapshot storage location is regional",
                        },
                        "guestcredentials": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "credentials": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "password": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "password to access the network path",
                                                },
                                                "name": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "username to access the network path",
                                                },
                                            },
                                        },
                                    },
                                    "savedcredentials": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": &schema.Schema{
                                                    Type:        schema.TypeString,
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
                        "vmbackuptype": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "datastorefreespacecheck": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if Datastore Free space check is enabled",
                        },
                        "datastorefreespacerequired": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "precentage of datastore free space check value",
                        },
                        "customsnapshottags": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "represents custom tags to be set on snapshots",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "represents name of the tag",
                                    },
                                    "value": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "represents value of the tag",
                                    },
                                },
                            },
                        },
                        "isapplicationaware": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Is the VM App Aware",
                        },
                        "transportmode": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "transport mode based on environment. Values are case sensitive",
                        },
                        "collectfiledetailsfromsnapshotcopy": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if metadata collection is enabled for intellisnap jobs. Only applicable for Indexing v1",
                        },
                        "crossaccount": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "shareonly": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if replicate and copy or sharing of amazon snapshot to different amazon account in same or different geographic location is enabled",
                                    },
                                    "fullcopy": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if full copy of amazon snapshot to different amazon account is enabled",
                                    },
                                    "destinationaccount": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "id": &schema.Schema{
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
            "diskfilters": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rules": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "condition": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Operation type for VM rules/filters",
                                    },
                                    "vmname": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "VM Name of the Virtual Machine whose disk has to be filtered . This is optional. if not given, all disks of name and type from all Vms added in content will be filtered",
                                    },
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "The string to be filtered",
                                    },
                                    "filtertype": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "overwrite": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if disk filter in vmgroup has to be overwritten, by default it will append the content",
                                    },
                                    "value": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "The value string to be filtered, in case of disk tag , value of tag to be filtered",
                                    },
                                    "vmguid": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "VM Guid of the Virtual Machine whose disk has to be filtered . This is optional. if not given, all disks of name and type from all Vms added in content will be filtered",
                                    },
                                },
                            },
                        },
                        "overwrite": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if content in vmgroup has to be overwritten, by default it will append the content",
                        },
                    },
                },
            },
            "securityassociations": &schema.Schema{
                Type:        schema.TypeSet,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "role": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "id": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "iscreatorassociation": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "externalusergroup": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "External User Group Entity",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "providerid": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Provider id",
                                    },
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "External Group Name",
                                    },
                                    "id": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "User Group Id",
                                    },
                                    "providername": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Provider Name",
                                    },
                                },
                            },
                        },
                        "permissionlist": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "permissionid": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "exclude": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Flag to specify if this is included permission or excluded permission.",
                                    },
                                    "type": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Returns the type of association.",
                                    },
                                    "categoryname": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "categoryid": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "permissionname": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "user": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "id": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "usergroup": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "id": &schema.Schema{
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
            "filters": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rulegroups": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "matchrule": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Enum which specifies the whether to match all rules or any of the rules",
                                    },
                                    "rules": &schema.Schema{
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "condition": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Operation type for VM rules/filters",
                                                },
                                                "displayname": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "The display name of the entity to be added",
                                                },
                                                "name": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "name of the VM to be added as content",
                                                },
                                                "guid": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "GUID of the entity to be added as content",
                                                },
                                                "type": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "value": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "value for the few type of VM Content like powerstate",
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "virtualmachines": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "guestcredentialassocid": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Credential association ID given to link entity with credential id.",
                                    },
                                    "guestcredentials": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "password": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "username": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                            },
                                        },
                                    },
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "name of the VM to be added as content",
                                    },
                                    "guid": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "GUID of the VM to be added as content",
                                    },
                                    "type": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "existingcredential": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "credentialid": &schema.Schema{
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "credentialname": &schema.Schema{
                                                    Type:        schema.TypeString,
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
                        "overwrite": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if content in vmgroup has to be overwritten, by default it will append the content",
                        },
                    },
                },
            },
            "accessnode": &schema.Schema{
                Type:        schema.TypeSet,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "id": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                    },
                },
            },
            "snapshotmanagement": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "useseparateproxyforsnaptotape": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if separate proxy client is used for snap to tape",
                        },
                        "snapengine": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "id": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "isindependentdisksenabled": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if independent disk option is enabled",
                        },
                        "backupcopyinterface": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "enablehardwaresnapshot": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if hardware snapshot is enabled",
                        },
                        "snapmountproxy": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "id": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "vmapplicationusername": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Virtual machine application user name",
                        },
                        "snapmountesxhost": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Name of ESX Host",
                        },
                        "israwdevicemapsenabled": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if raw device maps option is enabled",
                        },
                    },
                },
            },
            "enablefileindexing": &schema.Schema{
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: "True if file indexing needs to be enabled",
            },
            "newname": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "subclient name ",
            },
            "applicationvalidation": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "recoverytarget": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "id": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                },
                            },
                        },
                        "schedule": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "Schedule for application validation for VM Group",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "isscheduleenabled": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "True if application validation schedule is enabled",
                                    },
                                    "description": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Description for validation schedule",
                                    },
                                    "id": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "taskid": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Job Id for the application validation task. 0 if schedule is disabled",
                                    },
                                },
                            },
                        },
                        "maximumnoofthreads": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Number of backup Validation Threads",
                        },
                        "guestcredentials": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "credentials": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "password": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "password to access the network path",
                                                },
                                                "name": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "username to access the network path",
                                                },
                                            },
                                        },
                                    },
                                    "savedcredentials": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": &schema.Schema{
                                                    Type:        schema.TypeString,
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
                        "keepvalidatedvmsrunning": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "If true then validated VMs will be available until expiration time set on the recovery target",
                        },
                        "validatevmbackups": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "True if VM Backup validation is enabled",
                        },
                        "usesourcevmesxtomount": &schema.Schema{
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Computed:    true,
                            Description: "Use Source VM ESX To Mount",
                        },
                        "customvalidationscript": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "Custom validation script to be used during VM backup validation",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "windows": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "path": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Path for the validation script",
                                                },
                                                "unccredentials": &schema.Schema{
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "password": &schema.Schema{
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "password to access the network path",
                                                            },
                                                            "name": &schema.Schema{
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "username to access the network path",
                                                            },
                                                        },
                                                    },
                                                },
                                                "uncsavedcredentials": &schema.Schema{
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "name": &schema.Schema{
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "",
                                                            },
                                                            "id": &schema.Schema{
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "",
                                                            },
                                                        },
                                                    },
                                                },
                                                "arguments": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Arguments for the script",
                                                },
                                                "isdisabled": &schema.Schema{
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Is the script disabled",
                                                },
                                                "islocal": &schema.Schema{
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "True if the script is local",
                                                },
                                            },
                                        },
                                    },
                                    "unix": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "path": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Path for the validation script",
                                                },
                                                "unccredentials": &schema.Schema{
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "password": &schema.Schema{
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "password to access the network path",
                                                            },
                                                            "name": &schema.Schema{
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "username to access the network path",
                                                            },
                                                        },
                                                    },
                                                },
                                                "uncsavedcredentials": &schema.Schema{
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "name": &schema.Schema{
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "",
                                                            },
                                                            "id": &schema.Schema{
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "",
                                                            },
                                                        },
                                                    },
                                                },
                                                "arguments": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Arguments for the script",
                                                },
                                                "isdisabled": &schema.Schema{
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Is the script disabled",
                                                },
                                                "islocal": &schema.Schema{
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "True if the script is local",
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "copy": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "ismirrorcopy": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Is this a mirror copy?",
                                    },
                                    "snapcopytype": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "isdefault": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Is this a default backup destination?",
                                    },
                                    "copyprecedence": &schema.Schema{
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Order of backup destinaion copy created in storage policy",
                                    },
                                    "issnapcopy": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Is this a snap copy?",
                                    },
                                    "copytype": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "defaultreplicacopy": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Is this a default replica copy?",
                                    },
                                    "isactive": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Is this an active backup destination?",
                                    },
                                    "arrayreplicacopy": &schema.Schema{
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Is this an array replica copy?",
                                    },
                                    "backupdestination": &schema.Schema{
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": &schema.Schema{
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "id": &schema.Schema{
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
            "meditechsystems": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "systemname": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Meditech system name",
                        },
                        "listenerip": &schema.Schema{
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Meditech Listener IP of FQDN name",
                        },
                        "useraccount": &schema.Schema{
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "password": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "password to access the network path",
                                    },
                                    "name": &schema.Schema{
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "username to access the network path",
                                    },
                                },
                            },
                        },
                        "listenerport": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Meditech Listener Port",
                        },
                        "mbftimeout": &schema.Schema{
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "MBF timeout (in seconds)",
                        },
                    },
                },
            },
        },
    }
}

func resourceCreateVMGroup_V2(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/VMGroup
    var response_id = strconv.Itoa(0)
    var t_meditech *handler.MsgmeditechPropResp
    if v, ok := d.GetOk("meditech"); ok {
        val := v.([]interface{})
        t_meditech = build_vmgroup_v2_msgmeditechpropresp(d, val)
    }
    var t_hypervisor *handler.MsgIdName
    if v, ok := d.GetOk("hypervisor"); ok {
        val := v.([]interface{})
        t_hypervisor = build_vmgroup_v2_msgidname(d, val)
    }
    var t_name *string
    if v, ok := d.GetOk("name"); ok {
        val := v.(string)
        t_name = new(string)
        t_name = &val
    }
    var t_storage *handler.MsgIdName
    if v, ok := d.GetOk("storage"); ok {
        val := v.([]interface{})
        t_storage = build_vmgroup_v2_msgidname(d, val)
    }
    var t_enableintellisnap *bool
    if v, ok := d.GetOkExists("enableintellisnap"); ok {
        val := v.(bool)
        t_enableintellisnap = new(bool)
        t_enableintellisnap = &val
    }
    var t_plan *handler.MsgIdName
    if v, ok := d.GetOk("plan"); ok {
        val := v.([]interface{})
        t_plan = build_vmgroup_v2_msgidname(d, val)
    }
    var t_content *handler.MsgvmContent
    if v, ok := d.GetOk("content"); ok {
        val := v.([]interface{})
        t_content = build_vmgroup_v2_msgvmcontent(d, val)
    }
    var req = handler.MsgCreateVMGroupRequest{Meditech:t_meditech, Hypervisor:t_hypervisor, Name:t_name, Storage:t_storage, EnableIntellisnap:t_enableintellisnap, Plan:t_plan, Content:t_content}
    resp, err := handler.CvCreateVMGroup(req)
    if err != nil {
        return fmt.Errorf("Operation [CreateVMGroup] failed, Error %s", err)
    }
    if resp.SubclientId != nil {
        response_id = strconv.Itoa(*resp.SubclientId)
    }
    if response_id == "0" {
        return fmt.Errorf("Operation [CreateVMGroup] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateVMGroup_V2(d, m)
    }
}

func resourceReadVMGroup_V2(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/VmGroup/{VmGroupId}
    resp, err := handler.CvGetVMGroup(d.Id())
    if err != nil {
        return fmt.Errorf("Operation [GetVMGroup] failed, Error %s", err)
    }
    if resp.ActivityControl != nil {
        d.Set("activitycontrol", serialize_vmgroup_v2_msgactivitycontroloptions(resp.ActivityControl))
    } else {
        d.Set("activitycontrol", make([]map[string]interface{}, 0))
    }
    if resp.Settings != nil {
        d.Set("settings", serialize_vmgroup_v2_msgvmgroupsettings(resp.Settings))
    } else {
        d.Set("settings", make([]map[string]interface{}, 0))
    }
    if resp.DiskFilters != nil {
        d.Set("diskfilters", serialize_vmgroup_v2_msgvmdiskfilterpropset_array(resp.DiskFilters))
    } else {
        d.Set("diskfilters", make([]map[string]interface{}, 0))
    }
    if resp.SecurityAssociations != nil {
        d.Set("securityassociations", serialize_vmgroup_v2_msgsecurityassocset_array(resp.SecurityAssociations))
    } else {
        d.Set("securityassociations", make([]map[string]interface{}, 0))
    }
    if resp.Filters != nil {
        d.Set("filters", serialize_vmgroup_v2_msgvmcontentset_array(resp.Filters))
    } else {
        d.Set("filters", make([]map[string]interface{}, 0))
    }
    if resp.Content != nil {
        d.Set("content", serialize_vmgroup_v2_msgvmcontentset_array(resp.Content))
    } else {
        d.Set("content", make([]map[string]interface{}, 0))
    }
    if resp.SnapshotManagement != nil {
        d.Set("snapshotmanagement", serialize_vmgroup_v2_msgsnapcopyinfo(resp.SnapshotManagement))
    } else {
        d.Set("snapshotmanagement", make([]map[string]interface{}, 0))
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    if resp.ApplicationValidation != nil {
        d.Set("applicationvalidation", serialize_vmgroup_v2_msgvmappvalidation(resp.ApplicationValidation))
    } else {
        d.Set("applicationvalidation", make([]map[string]interface{}, 0))
    }
    if resp.MeditechSystems != nil {
        d.Set("meditechsystems", serialize_vmgroup_v2_msgmeditechpropresp(resp.MeditechSystems))
    } else {
        d.Set("meditechsystems", make([]map[string]interface{}, 0))
    }
    return nil
}

func resourceUpdateVMGroup_V2(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/VmGroup/{VmGroupId}
    var t_activitycontrol *handler.MsgActivityControlOptions
    if d.HasChange("activitycontrol") {
        val := d.Get("activitycontrol").([]interface{})
        t_activitycontrol = build_vmgroup_v2_msgactivitycontroloptions(d, val)
    }
    var t_settings *handler.MsgvmGroupSettings
    if d.HasChange("settings") {
        val := d.Get("settings").([]interface{})
        t_settings = build_vmgroup_v2_msgvmgroupsettings(d, val)
    }
    var t_diskfilters *handler.MsgvmDiskFilterProp
    if d.HasChange("diskfilters") {
        val := d.Get("diskfilters").([]interface{})
        t_diskfilters = build_vmgroup_v2_msgvmdiskfilterprop(d, val)
    }
    var t_securityassociations []handler.MsgSecurityAssocSet
    if d.HasChange("securityassociations") {
        val := d.Get("securityassociations").(*schema.Set)
        t_securityassociations = build_vmgroup_v2_msgsecurityassocset_array(d, val.List())
    }
    var t_storage *handler.MsgIdName
    if d.HasChange("storage") {
        val := d.Get("storage").([]interface{})
        t_storage = build_vmgroup_v2_msgidname(d, val)
    }
    var t_filters *handler.MsgvmContent
    if d.HasChange("filters") {
        val := d.Get("filters").([]interface{})
        t_filters = build_vmgroup_v2_msgvmcontent(d, val)
    }
    var t_accessnode []handler.MsgIdNameSet
    if d.HasChange("accessnode") {
        val := d.Get("accessnode").(*schema.Set)
        t_accessnode = build_vmgroup_v2_msgidnameset_array(d, val.List())
    }
    var t_content *handler.MsgvmContent
    if d.HasChange("content") {
        val := d.Get("content").([]interface{})
        t_content = build_vmgroup_v2_msgvmcontent(d, val)
    }
    var t_snapshotmanagement *handler.MsgsnapCopyInfo
    if d.HasChange("snapshotmanagement") {
        val := d.Get("snapshotmanagement").([]interface{})
        t_snapshotmanagement = build_vmgroup_v2_msgsnapcopyinfo(d, val)
    }
    var t_enablefileindexing *bool
    if d.HasChange("enablefileindexing") {
        val := d.Get("enablefileindexing").(bool)
        t_enablefileindexing = new(bool)
        t_enablefileindexing = &val
    }
    var t_newname *string
    if d.HasChange("newname") {
        val := d.Get("newname").(string)
        t_newname = new(string)
        t_newname = &val
    }
    var t_applicationvalidation *handler.MsgvmAppValidation
    if d.HasChange("applicationvalidation") {
        val := d.Get("applicationvalidation").([]interface{})
        t_applicationvalidation = build_vmgroup_v2_msgvmappvalidation(d, val)
    }
    var t_plan *handler.MsgIdName
    if d.HasChange("plan") {
        val := d.Get("plan").([]interface{})
        t_plan = build_vmgroup_v2_msgidname(d, val)
    }
    var t_meditechsystems *handler.MsgmeditechPropResp
    if d.HasChange("meditechsystems") {
        val := d.Get("meditechsystems").([]interface{})
        t_meditechsystems = build_vmgroup_v2_msgmeditechpropresp(d, val)
    }
    var req = handler.MsgUpdateVMGroupRequest{ActivityControl:t_activitycontrol, Settings:t_settings, DiskFilters:t_diskfilters, SecurityAssociations:t_securityassociations, Storage:t_storage, Filters:t_filters, AccessNode:t_accessnode, Content:t_content, SnapshotManagement:t_snapshotmanagement, EnableFileIndexing:t_enablefileindexing, NewName:t_newname, ApplicationValidation:t_applicationvalidation, Plan:t_plan, MeditechSystems:t_meditechsystems}
    _, err := handler.CvUpdateVMGroup(req, d.Id())
    if err != nil {
        return fmt.Errorf("Operation [UpdateVMGroup] failed, Error %s", err)
    }
    return resourceReadVMGroup_V2(d, m)
}

func resourceCreateUpdateVMGroup_V2(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/VmGroup/{VmGroupId}
    var execUpdate bool = false
    var t_activitycontrol *handler.MsgActivityControlOptions
    if v, ok := d.GetOk("activitycontrol"); ok {
        val := v.([]interface{})
        t_activitycontrol = build_vmgroup_v2_msgactivitycontroloptions(d, val)
        execUpdate = true
    }
    var t_settings *handler.MsgvmGroupSettings
    if v, ok := d.GetOk("settings"); ok {
        val := v.([]interface{})
        t_settings = build_vmgroup_v2_msgvmgroupsettings(d, val)
        execUpdate = true
    }
    var t_diskfilters *handler.MsgvmDiskFilterProp
    if v, ok := d.GetOk("diskfilters"); ok {
        val := v.([]interface{})
        t_diskfilters = build_vmgroup_v2_msgvmdiskfilterprop(d, val)
        execUpdate = true
    }
    var t_securityassociations []handler.MsgSecurityAssocSet
    if v, ok := d.GetOk("securityassociations"); ok {
        val := v.(*schema.Set)
        t_securityassociations = build_vmgroup_v2_msgsecurityassocset_array(d, val.List())
        execUpdate = true
    }
    var t_filters *handler.MsgvmContent
    if v, ok := d.GetOk("filters"); ok {
        val := v.([]interface{})
        t_filters = build_vmgroup_v2_msgvmcontent(d, val)
        execUpdate = true
    }
    var t_accessnode []handler.MsgIdNameSet
    if v, ok := d.GetOk("accessnode"); ok {
        val := v.(*schema.Set)
        t_accessnode = build_vmgroup_v2_msgidnameset_array(d, val.List())
        execUpdate = true
    }
    var t_snapshotmanagement *handler.MsgsnapCopyInfo
    if v, ok := d.GetOk("snapshotmanagement"); ok {
        val := v.([]interface{})
        t_snapshotmanagement = build_vmgroup_v2_msgsnapcopyinfo(d, val)
        execUpdate = true
    }
    var t_enablefileindexing *bool
    if v, ok := d.GetOkExists("enablefileindexing"); ok {
        val := v.(bool)
        t_enablefileindexing = new(bool)
        t_enablefileindexing = &val
        execUpdate = true
    }
    var t_newname *string
    if v, ok := d.GetOk("newname"); ok {
        val := v.(string)
        t_newname = new(string)
        t_newname = &val
        execUpdate = true
    }
    var t_applicationvalidation *handler.MsgvmAppValidation
    if v, ok := d.GetOk("applicationvalidation"); ok {
        val := v.([]interface{})
        t_applicationvalidation = build_vmgroup_v2_msgvmappvalidation(d, val)
        execUpdate = true
    }
    var t_meditechsystems *handler.MsgmeditechPropResp
    if v, ok := d.GetOk("meditechsystems"); ok {
        val := v.([]interface{})
        t_meditechsystems = build_vmgroup_v2_msgmeditechpropresp(d, val)
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgUpdateVMGroupRequest{ActivityControl:t_activitycontrol, Settings:t_settings, DiskFilters:t_diskfilters, SecurityAssociations:t_securityassociations, Filters:t_filters, AccessNode:t_accessnode, SnapshotManagement:t_snapshotmanagement, EnableFileIndexing:t_enablefileindexing, NewName:t_newname, ApplicationValidation:t_applicationvalidation, MeditechSystems:t_meditechsystems}
        _, err := handler.CvUpdateVMGroup(req, d.Id())
        if err != nil {
            return fmt.Errorf("Operation [UpdateVMGroup] failed, Error %s", err)
        }
    }
    return resourceReadVMGroup_V2(d, m)
}

func resourceDeleteVMGroup_V2(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/VmGroup/{VmGroupId}
    _, err := handler.CvDeleteVMGroup(d.Id())
    if err != nil {
        return fmt.Errorf("Operation [DeleteVMGroup] failed, Error %s", err)
    }
    return nil
}

func build_vmgroup_v2_msgmeditechpropresp(d *schema.ResourceData, r []interface{}) *handler.MsgmeditechPropResp {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_systemname *string
        if val, ok := tmp["systemname"].(string); ok {
            t_systemname = new(string)
            t_systemname = &val
        }
        var t_listenerip *string
        if val, ok := tmp["listenerip"].(string); ok {
            t_listenerip = new(string)
            t_listenerip = &val
        }
        var t_useraccount *handler.MsgUserNamePassword
        if val, ok := tmp["useraccount"].([]interface{}); ok {
            t_useraccount = build_vmgroup_v2_msgusernamepassword(d, val)
        }
        var t_listenerport *int
        if val, ok := tmp["listenerport"].(int); ok {
            t_listenerport = new(int)
            t_listenerport = &val
        }
        var t_mbftimeout *int
        if val, ok := tmp["mbftimeout"].(int); ok {
            t_mbftimeout = new(int)
            t_mbftimeout = &val
        }
        return &handler.MsgmeditechPropResp{SystemName:t_systemname, ListenerIP:t_listenerip, UserAccount:t_useraccount, ListenerPort:t_listenerport, MBFtimeout:t_mbftimeout}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgusernamepassword(d *schema.ResourceData, r []interface{}) *handler.MsgUserNamePassword {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_password *string
        if val, ok := tmp["password"].(string); ok {
            t_password = new(string)
            t_password = &val
        }
        var t_name *string
        if val, ok := tmp["name"].(string); ok {
            t_name = new(string)
            t_name = &val
        }
        return &handler.MsgUserNamePassword{Password:t_password, Name:t_name}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvmappvalidation(d *schema.ResourceData, r []interface{}) *handler.MsgvmAppValidation {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_recoverytarget *handler.MsgIdName
        if val, ok := tmp["recoverytarget"].([]interface{}); ok {
            t_recoverytarget = build_vmgroup_v2_msgidname(d, val)
        }
        var t_schedule *handler.MsgValidationScheduleObject
        if val, ok := tmp["schedule"].([]interface{}); ok {
            t_schedule = build_vmgroup_v2_msgvalidationscheduleobject(d, val)
        }
        var t_maximumnoofthreads *int
        if val, ok := tmp["maximumnoofthreads"].(int); ok {
            t_maximumnoofthreads = new(int)
            t_maximumnoofthreads = &val
        }
        var t_guestcredentials *handler.MsgguestCredentialInfo
        if val, ok := tmp["guestcredentials"].([]interface{}); ok {
            t_guestcredentials = build_vmgroup_v2_msgguestcredentialinfo(d, val)
        }
        var t_keepvalidatedvmsrunning *bool
        if val, ok := tmp["keepvalidatedvmsrunning"].(bool); ok {
            t_keepvalidatedvmsrunning = new(bool)
            t_keepvalidatedvmsrunning = &val
        }
        var t_validatevmbackups *bool
        if val, ok := tmp["validatevmbackups"].(bool); ok {
            t_validatevmbackups = new(bool)
            t_validatevmbackups = &val
        }
        var t_usesourcevmesxtomount *bool
        if val, ok := tmp["usesourcevmesxtomount"].(bool); ok {
            t_usesourcevmesxtomount = new(bool)
            t_usesourcevmesxtomount = &val
        }
        var t_customvalidationscript *handler.MsgappValidationScript
        if val, ok := tmp["customvalidationscript"].([]interface{}); ok {
            t_customvalidationscript = build_vmgroup_v2_msgappvalidationscript(d, val)
        }
        var t_copy *handler.MsgPlanSourceCopy
        if val, ok := tmp["copy"].([]interface{}); ok {
            t_copy = build_vmgroup_v2_msgplansourcecopy(d, val)
        }
        return &handler.MsgvmAppValidation{RecoveryTarget:t_recoverytarget, Schedule:t_schedule, MaximumNoOfThreads:t_maximumnoofthreads, GuestCredentials:t_guestcredentials, KeepValidatedVMsRunning:t_keepvalidatedvmsrunning, ValidateVMBackups:t_validatevmbackups, UseSourceVmESXToMount:t_usesourcevmesxtomount, CustomValidationScript:t_customvalidationscript, Copy:t_copy}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgplansourcecopy(d *schema.ResourceData, r []interface{}) *handler.MsgPlanSourceCopy {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_ismirrorcopy *bool
        if val, ok := tmp["ismirrorcopy"].(bool); ok {
            t_ismirrorcopy = new(bool)
            t_ismirrorcopy = &val
        }
        var t_snapcopytype *string
        if val, ok := tmp["snapcopytype"].(string); ok {
            t_snapcopytype = new(string)
            t_snapcopytype = &val
        }
        var t_isdefault *bool
        if val, ok := tmp["isdefault"].(bool); ok {
            t_isdefault = new(bool)
            t_isdefault = &val
        }
        var t_copyprecedence *int
        if val, ok := tmp["copyprecedence"].(int); ok {
            t_copyprecedence = new(int)
            t_copyprecedence = &val
        }
        var t_issnapcopy *bool
        if val, ok := tmp["issnapcopy"].(bool); ok {
            t_issnapcopy = new(bool)
            t_issnapcopy = &val
        }
        var t_copytype *string
        if val, ok := tmp["copytype"].(string); ok {
            t_copytype = new(string)
            t_copytype = &val
        }
        var t_defaultreplicacopy *bool
        if val, ok := tmp["defaultreplicacopy"].(bool); ok {
            t_defaultreplicacopy = new(bool)
            t_defaultreplicacopy = &val
        }
        var t_isactive *bool
        if val, ok := tmp["isactive"].(bool); ok {
            t_isactive = new(bool)
            t_isactive = &val
        }
        var t_arrayreplicacopy *bool
        if val, ok := tmp["arrayreplicacopy"].(bool); ok {
            t_arrayreplicacopy = new(bool)
            t_arrayreplicacopy = &val
        }
        var t_backupdestination *handler.MsgIdName
        if val, ok := tmp["backupdestination"].([]interface{}); ok {
            t_backupdestination = build_vmgroup_v2_msgidname(d, val)
        }
        return &handler.MsgPlanSourceCopy{IsMirrorCopy:t_ismirrorcopy, SnapCopyType:t_snapcopytype, IsDefault:t_isdefault, CopyPrecedence:t_copyprecedence, IsSnapCopy:t_issnapcopy, CopyType:t_copytype, DefaultReplicaCopy:t_defaultreplicacopy, IsActive:t_isactive, ArrayReplicaCopy:t_arrayreplicacopy, BackupDestination:t_backupdestination}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_name *string
        if val, ok := tmp["name"].(string); ok {
            t_name = new(string)
            t_name = &val
        }
        var t_id *int
        if val, ok := tmp["id"].(int); ok {
            if !handler.IsEmptyInt(val) {
                t_id = new(int)
                t_id = &val
            }
        }
        return &handler.MsgIdName{Name:t_name, Id:t_id}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgappvalidationscript(d *schema.ResourceData, r []interface{}) *handler.MsgappValidationScript {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_windows *handler.MsgValidationScript
        if val, ok := tmp["windows"].([]interface{}); ok {
            t_windows = build_vmgroup_v2_msgvalidationscript(d, val)
        }
        var t_unix *handler.MsgValidationScript
        if val, ok := tmp["unix"].([]interface{}); ok {
            t_unix = build_vmgroup_v2_msgvalidationscript(d, val)
        }
        return &handler.MsgappValidationScript{Windows:t_windows, Unix:t_unix}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvalidationscript(d *schema.ResourceData, r []interface{}) *handler.MsgValidationScript {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_path *string
        if val, ok := tmp["path"].(string); ok {
            t_path = new(string)
            t_path = &val
        }
        var t_unccredentials *handler.MsgUserNamePassword
        if val, ok := tmp["unccredentials"].([]interface{}); ok {
            t_unccredentials = build_vmgroup_v2_msgusernamepassword(d, val)
        }
        var t_uncsavedcredentials *handler.MsgIdName
        if val, ok := tmp["uncsavedcredentials"].([]interface{}); ok {
            t_uncsavedcredentials = build_vmgroup_v2_msgidname(d, val)
        }
        var t_arguments *string
        if val, ok := tmp["arguments"].(string); ok {
            t_arguments = new(string)
            t_arguments = &val
        }
        var t_isdisabled *bool
        if val, ok := tmp["isdisabled"].(bool); ok {
            t_isdisabled = new(bool)
            t_isdisabled = &val
        }
        var t_islocal *bool
        if val, ok := tmp["islocal"].(bool); ok {
            t_islocal = new(bool)
            t_islocal = &val
        }
        return &handler.MsgValidationScript{Path:t_path, UNCCredentials:t_unccredentials, UNCSavedCredentials:t_uncsavedcredentials, Arguments:t_arguments, IsDisabled:t_isdisabled, IsLocal:t_islocal}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgguestcredentialinfo(d *schema.ResourceData, r []interface{}) *handler.MsgguestCredentialInfo {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_credentials *handler.MsgUserNamePassword
        if val, ok := tmp["credentials"].([]interface{}); ok {
            t_credentials = build_vmgroup_v2_msgusernamepassword(d, val)
        }
        var t_savedcredentials *handler.MsgIdName
        if val, ok := tmp["savedcredentials"].([]interface{}); ok {
            t_savedcredentials = build_vmgroup_v2_msgidname(d, val)
        }
        return &handler.MsgguestCredentialInfo{Credentials:t_credentials, SavedCredentials:t_savedcredentials}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvalidationscheduleobject(d *schema.ResourceData, r []interface{}) *handler.MsgValidationScheduleObject {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_isscheduleenabled *bool
        if val, ok := tmp["isscheduleenabled"].(bool); ok {
            t_isscheduleenabled = new(bool)
            t_isscheduleenabled = &val
        }
        var t_description *string
        if val, ok := tmp["description"].(string); ok {
            t_description = new(string)
            t_description = &val
        }
        var t_id *int
        if val, ok := tmp["id"].(int); ok {
            t_id = new(int)
            t_id = &val
        }
        var t_taskid *int
        if val, ok := tmp["taskid"].(int); ok {
            t_taskid = new(int)
            t_taskid = &val
        }
        return &handler.MsgValidationScheduleObject{IsScheduleEnabled:t_isscheduleenabled, Description:t_description, Id:t_id, TaskId:t_taskid}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgsnapcopyinfo(d *schema.ResourceData, r []interface{}) *handler.MsgsnapCopyInfo {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_useseparateproxyforsnaptotape *bool
        if val, ok := tmp["useseparateproxyforsnaptotape"].(bool); ok {
            t_useseparateproxyforsnaptotape = new(bool)
            t_useseparateproxyforsnaptotape = &val
        }
        var t_snapengine *handler.MsgIdName
        if val, ok := tmp["snapengine"].([]interface{}); ok {
            t_snapengine = build_vmgroup_v2_msgidname(d, val)
        }
        var t_isindependentdisksenabled *bool
        if val, ok := tmp["isindependentdisksenabled"].(bool); ok {
            t_isindependentdisksenabled = new(bool)
            t_isindependentdisksenabled = &val
        }
        var t_backupcopyinterface *string
        if val, ok := tmp["backupcopyinterface"].(string); ok {
            t_backupcopyinterface = new(string)
            t_backupcopyinterface = &val
        }
        var t_enablehardwaresnapshot *bool
        if val, ok := tmp["enablehardwaresnapshot"].(bool); ok {
            t_enablehardwaresnapshot = new(bool)
            t_enablehardwaresnapshot = &val
        }
        var t_snapmountproxy *handler.MsgIdName
        if val, ok := tmp["snapmountproxy"].([]interface{}); ok {
            t_snapmountproxy = build_vmgroup_v2_msgidname(d, val)
        }
        var t_vmapplicationusername *string
        if val, ok := tmp["vmapplicationusername"].(string); ok {
            t_vmapplicationusername = new(string)
            t_vmapplicationusername = &val
        }
        var t_snapmountesxhost *string
        if val, ok := tmp["snapmountesxhost"].(string); ok {
            t_snapmountesxhost = new(string)
            t_snapmountesxhost = &val
        }
        var t_israwdevicemapsenabled *bool
        if val, ok := tmp["israwdevicemapsenabled"].(bool); ok {
            t_israwdevicemapsenabled = new(bool)
            t_israwdevicemapsenabled = &val
        }
        return &handler.MsgsnapCopyInfo{UseSeparateProxyForSnapToTape:t_useseparateproxyforsnaptotape, SnapEngine:t_snapengine, IsIndependentDisksEnabled:t_isindependentdisksenabled, BackupCopyInterface:t_backupcopyinterface, EnableHardwareSnapshot:t_enablehardwaresnapshot, SnapMountProxy:t_snapmountproxy, VmApplicationUserName:t_vmapplicationusername, SnapMountESXHost:t_snapmountesxhost, IsRawDeviceMapsEnabled:t_israwdevicemapsenabled}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgidnameset_array(d *schema.ResourceData, r []interface{}) []handler.MsgIdNameSet {
    if r != nil {
        tmp := make([]handler.MsgIdNameSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_id *int
            if val, ok := raw_a["id"].(int); ok {
                if !handler.IsEmptyInt(val) {
                    t_id = new(int)
                    t_id = &val
                }
            }
            tmp[a] = handler.MsgIdNameSet{Id:t_id}
        }
        return tmp
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvmcontent(d *schema.ResourceData, r []interface{}) *handler.MsgvmContent {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_rulegroups []handler.MsgRuleGroupContentSet
        if val, ok := tmp["rulegroups"].(*schema.Set); ok {
            t_rulegroups = build_vmgroup_v2_msgrulegroupcontentset_array(d, val.List())
        }
        var t_virtualmachines []handler.MsgVirtualMachinecontentSet
        if val, ok := tmp["virtualmachines"].(*schema.Set); ok {
            t_virtualmachines = build_vmgroup_v2_msgvirtualmachinecontentset_array(d, val.List())
        }
        var t_overwrite *bool
        if val, ok := tmp["overwrite"].(bool); ok {
            t_overwrite = new(bool)
            t_overwrite = &val
        }
        return &handler.MsgvmContent{RuleGroups:t_rulegroups, VirtualMachines:t_virtualmachines, Overwrite:t_overwrite}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvirtualmachinecontentset_array(d *schema.ResourceData, r []interface{}) []handler.MsgVirtualMachinecontentSet {
    if r != nil {
        tmp := make([]handler.MsgVirtualMachinecontentSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_guestcredentialassocid *int
            if val, ok := raw_a["guestcredentialassocid"].(int); ok {
                t_guestcredentialassocid = new(int)
                t_guestcredentialassocid = &val
            }
            var t_guestcredentials *handler.MsgVMGuestCredentials
            if val, ok := raw_a["guestcredentials"].([]interface{}); ok {
                t_guestcredentials = build_vmgroup_v2_msgvmguestcredentials(d, val)
            }
            var t_name *string
            if val, ok := raw_a["name"].(string); ok {
                t_name = new(string)
                t_name = &val
            }
            var t_guid *string
            if val, ok := raw_a["guid"].(string); ok {
                t_guid = new(string)
                t_guid = &val
            }
            var t_type *string
            if val, ok := raw_a["type"].(string); ok {
                t_type = new(string)
                t_type = &val
            }
            var t_existingcredential *handler.MsgVMExistingCredential
            if val, ok := raw_a["existingcredential"].([]interface{}); ok {
                t_existingcredential = build_vmgroup_v2_msgvmexistingcredential(d, val)
            }
            tmp[a] = handler.MsgVirtualMachinecontentSet{GuestCredentialAssocId:t_guestcredentialassocid, GuestCredentials:t_guestcredentials, Name:t_name, GUID:t_guid, Type:t_type, ExistingCredential:t_existingcredential}
        }
        return tmp
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvmexistingcredential(d *schema.ResourceData, r []interface{}) *handler.MsgVMExistingCredential {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_credentialid *int
        if val, ok := tmp["credentialid"].(int); ok {
            t_credentialid = new(int)
            t_credentialid = &val
        }
        var t_credentialname *string
        if val, ok := tmp["credentialname"].(string); ok {
            t_credentialname = new(string)
            t_credentialname = &val
        }
        return &handler.MsgVMExistingCredential{CredentialId:t_credentialid, CredentialName:t_credentialname}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvmguestcredentials(d *schema.ResourceData, r []interface{}) *handler.MsgVMGuestCredentials {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_password *string
        if val, ok := tmp["password"].(string); ok {
            t_password = new(string)
            t_password = &val
        }
        var t_username *string
        if val, ok := tmp["username"].(string); ok {
            t_username = new(string)
            t_username = &val
        }
        return &handler.MsgVMGuestCredentials{Password:t_password, UserName:t_username}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgrulegroupcontentset_array(d *schema.ResourceData, r []interface{}) []handler.MsgRuleGroupContentSet {
    if r != nil {
        tmp := make([]handler.MsgRuleGroupContentSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_matchrule *string
            if val, ok := raw_a["matchrule"].(string); ok {
                t_matchrule = new(string)
                t_matchrule = &val
            }
            var t_rules []handler.MsgRuleContentSet
            if val, ok := raw_a["rules"].(*schema.Set); ok {
                t_rules = build_vmgroup_v2_msgrulecontentset_array(d, val.List())
            }
            tmp[a] = handler.MsgRuleGroupContentSet{MatchRule:t_matchrule, Rules:t_rules}
        }
        return tmp
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgrulecontentset_array(d *schema.ResourceData, r []interface{}) []handler.MsgRuleContentSet {
    if r != nil {
        tmp := make([]handler.MsgRuleContentSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_condition *string
            if val, ok := raw_a["condition"].(string); ok {
                t_condition = new(string)
                t_condition = &val
            }
            var t_displayname *string
            if val, ok := raw_a["displayname"].(string); ok {
                t_displayname = new(string)
                t_displayname = &val
            }
            var t_name *string
            if val, ok := raw_a["name"].(string); ok {
                t_name = new(string)
                t_name = &val
            }
            var t_guid *string
            if val, ok := raw_a["guid"].(string); ok {
                t_guid = new(string)
                t_guid = &val
            }
            var t_type *string
            if val, ok := raw_a["type"].(string); ok {
                t_type = new(string)
                t_type = &val
            }
            var t_value *string
            if val, ok := raw_a["value"].(string); ok {
                t_value = new(string)
                t_value = &val
            }
            tmp[a] = handler.MsgRuleContentSet{Condition:t_condition, DisplayName:t_displayname, Name:t_name, GUID:t_guid, Type:t_type, Value:t_value}
        }
        return tmp
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgsecurityassocset_array(d *schema.ResourceData, r []interface{}) []handler.MsgSecurityAssocSet {
    if r != nil {
        tmp := make([]handler.MsgSecurityAssocSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_role *handler.MsgIdName
            if val, ok := raw_a["role"].([]interface{}); ok {
                t_role = build_vmgroup_v2_msgidname(d, val)
            }
            var t_iscreatorassociation *bool
            if val, ok := raw_a["iscreatorassociation"].(bool); ok {
                t_iscreatorassociation = new(bool)
                t_iscreatorassociation = &val
            }
            var t_externalusergroup *handler.MsgexternalUserGroup
            if val, ok := raw_a["externalusergroup"].([]interface{}); ok {
                t_externalusergroup = build_vmgroup_v2_msgexternalusergroup(d, val)
            }
            var t_permissionlist []handler.MsgPermissionRespSet
            if val, ok := raw_a["permissionlist"].(*schema.Set); ok {
                t_permissionlist = build_vmgroup_v2_msgpermissionrespset_array(d, val.List())
            }
            var t_user *handler.MsgIdName
            if val, ok := raw_a["user"].([]interface{}); ok {
                t_user = build_vmgroup_v2_msgidname(d, val)
            }
            var t_usergroup *handler.MsgIdName
            if val, ok := raw_a["usergroup"].([]interface{}); ok {
                t_usergroup = build_vmgroup_v2_msgidname(d, val)
            }
            tmp[a] = handler.MsgSecurityAssocSet{Role:t_role, IsCreatorAssociation:t_iscreatorassociation, ExternalUserGroup:t_externalusergroup, PermissionList:t_permissionlist, User:t_user, UserGroup:t_usergroup}
        }
        return tmp
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgpermissionrespset_array(d *schema.ResourceData, r []interface{}) []handler.MsgPermissionRespSet {
    if r != nil {
        tmp := make([]handler.MsgPermissionRespSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_permissionid *int
            if val, ok := raw_a["permissionid"].(int); ok {
                t_permissionid = new(int)
                t_permissionid = &val
            }
            var t_exclude *bool
            if val, ok := raw_a["exclude"].(bool); ok {
                t_exclude = new(bool)
                t_exclude = &val
            }
            var t_type *string
            if val, ok := raw_a["type"].(string); ok {
                t_type = new(string)
                t_type = &val
            }
            var t_categoryname *string
            if val, ok := raw_a["categoryname"].(string); ok {
                t_categoryname = new(string)
                t_categoryname = &val
            }
            var t_categoryid *int
            if val, ok := raw_a["categoryid"].(int); ok {
                t_categoryid = new(int)
                t_categoryid = &val
            }
            var t_permissionname *string
            if val, ok := raw_a["permissionname"].(string); ok {
                t_permissionname = new(string)
                t_permissionname = &val
            }
            tmp[a] = handler.MsgPermissionRespSet{PermissionId:t_permissionid, Exclude:t_exclude, Type:t_type, CategoryName:t_categoryname, CategoryId:t_categoryid, PermissionName:t_permissionname}
        }
        return tmp
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgexternalusergroup(d *schema.ResourceData, r []interface{}) *handler.MsgexternalUserGroup {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_providerid *int
        if val, ok := tmp["providerid"].(int); ok {
            t_providerid = new(int)
            t_providerid = &val
        }
        var t_name *string
        if val, ok := tmp["name"].(string); ok {
            t_name = new(string)
            t_name = &val
        }
        var t_id *int
        if val, ok := tmp["id"].(int); ok {
            t_id = new(int)
            t_id = &val
        }
        var t_providername *string
        if val, ok := tmp["providername"].(string); ok {
            t_providername = new(string)
            t_providername = &val
        }
        return &handler.MsgexternalUserGroup{ProviderId:t_providerid, Name:t_name, Id:t_id, ProviderName:t_providername}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvmdiskfilterprop(d *schema.ResourceData, r []interface{}) *handler.MsgvmDiskFilterProp {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_rules []handler.MsgvmDiskFilterSet
        if val, ok := tmp["rules"].(*schema.Set); ok {
            t_rules = build_vmgroup_v2_msgvmdiskfilterset_array(d, val.List())
        }
        var t_overwrite *bool
        if val, ok := tmp["overwrite"].(bool); ok {
            t_overwrite = new(bool)
            t_overwrite = &val
        }
        return &handler.MsgvmDiskFilterProp{Rules:t_rules, Overwrite:t_overwrite}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvmdiskfilterset_array(d *schema.ResourceData, r []interface{}) []handler.MsgvmDiskFilterSet {
    if r != nil {
        tmp := make([]handler.MsgvmDiskFilterSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_condition *string
            if val, ok := raw_a["condition"].(string); ok {
                t_condition = new(string)
                t_condition = &val
            }
            var t_vmname *string
            if val, ok := raw_a["vmname"].(string); ok {
                t_vmname = new(string)
                t_vmname = &val
            }
            var t_name *string
            if val, ok := raw_a["name"].(string); ok {
                t_name = new(string)
                t_name = &val
            }
            var t_filtertype *string
            if val, ok := raw_a["filtertype"].(string); ok {
                t_filtertype = new(string)
                t_filtertype = &val
            }
            var t_overwrite *bool
            if val, ok := raw_a["overwrite"].(bool); ok {
                t_overwrite = new(bool)
                t_overwrite = &val
            }
            var t_value *string
            if val, ok := raw_a["value"].(string); ok {
                t_value = new(string)
                t_value = &val
            }
            var t_vmguid *string
            if val, ok := raw_a["vmguid"].(string); ok {
                t_vmguid = new(string)
                t_vmguid = &val
            }
            tmp[a] = handler.MsgvmDiskFilterSet{Condition:t_condition, VmName:t_vmname, Name:t_name, FilterType:t_filtertype, Overwrite:t_overwrite, Value:t_value, VmGuid:t_vmguid}
        }
        return tmp
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgvmgroupsettings(d *schema.ResourceData, r []interface{}) *handler.MsgvmGroupSettings {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_autodetectvmowner *bool
        if val, ok := tmp["autodetectvmowner"].(bool); ok {
            t_autodetectvmowner = new(bool)
            t_autodetectvmowner = &val
        }
        var t_collectfiledetailsforgranularrecovery *bool
        if val, ok := tmp["collectfiledetailsforgranularrecovery"].(bool); ok {
            t_collectfiledetailsforgranularrecovery = new(bool)
            t_collectfiledetailsforgranularrecovery = &val
        }
        var t_noofreaders *int
        if val, ok := tmp["noofreaders"].(int); ok {
            t_noofreaders = new(int)
            t_noofreaders = &val
        }
        var t_usechangedblocktrackingonvm *bool
        if val, ok := tmp["usechangedblocktrackingonvm"].(bool); ok {
            t_usechangedblocktrackingonvm = new(bool)
            t_usechangedblocktrackingonvm = &val
        }
        var t_jobstarttime *int
        if val, ok := tmp["jobstarttime"].(int); ok {
            t_jobstarttime = new(int)
            t_jobstarttime = &val
        }
        var t_usevmcheckpointsetting *bool
        if val, ok := tmp["usevmcheckpointsetting"].(bool); ok {
            t_usevmcheckpointsetting = new(bool)
            t_usevmcheckpointsetting = &val
        }
        var t_customsnapshotresourcegroup *string
        if val, ok := tmp["customsnapshotresourcegroup"].(string); ok {
            t_customsnapshotresourcegroup = new(string)
            t_customsnapshotresourcegroup = &val
        }
        var t_regionalsnapshot *bool
        if val, ok := tmp["regionalsnapshot"].(bool); ok {
            t_regionalsnapshot = new(bool)
            t_regionalsnapshot = &val
        }
        var t_guestcredentials *handler.MsgguestCredentialInfo
        if val, ok := tmp["guestcredentials"].([]interface{}); ok {
            t_guestcredentials = build_vmgroup_v2_msgguestcredentialinfo(d, val)
        }
        var t_vmbackuptype *string
        if val, ok := tmp["vmbackuptype"].(string); ok {
            t_vmbackuptype = new(string)
            t_vmbackuptype = &val
        }
        var t_datastorefreespacecheck *bool
        if val, ok := tmp["datastorefreespacecheck"].(bool); ok {
            t_datastorefreespacecheck = new(bool)
            t_datastorefreespacecheck = &val
        }
        var t_datastorefreespacerequired *int
        if val, ok := tmp["datastorefreespacerequired"].(int); ok {
            t_datastorefreespacerequired = new(int)
            t_datastorefreespacerequired = &val
        }
        var t_customsnapshottags []handler.MsgresourceTagSet
        if val, ok := tmp["customsnapshottags"].(*schema.Set); ok {
            t_customsnapshottags = build_vmgroup_v2_msgresourcetagset_array(d, val.List())
        }
        var t_isapplicationaware *bool
        if val, ok := tmp["isapplicationaware"].(bool); ok {
            t_isapplicationaware = new(bool)
            t_isapplicationaware = &val
        }
        var t_transportmode *string
        if val, ok := tmp["transportmode"].(string); ok {
            t_transportmode = new(string)
            t_transportmode = &val
        }
        var t_collectfiledetailsfromsnapshotcopy *bool
        if val, ok := tmp["collectfiledetailsfromsnapshotcopy"].(bool); ok {
            t_collectfiledetailsfromsnapshotcopy = new(bool)
            t_collectfiledetailsfromsnapshotcopy = &val
        }
        var t_crossaccount *handler.MsgAmazonCrossAccount
        if val, ok := tmp["crossaccount"].([]interface{}); ok {
            t_crossaccount = build_vmgroup_v2_msgamazoncrossaccount(d, val)
        }
        return &handler.MsgvmGroupSettings{AutoDetectVMOwner:t_autodetectvmowner, CollectFileDetailsforGranularRecovery:t_collectfiledetailsforgranularrecovery, NoOfReaders:t_noofreaders, UseChangedBlockTrackingOnVM:t_usechangedblocktrackingonvm, JobStartTime:t_jobstarttime, UseVMCheckpointSetting:t_usevmcheckpointsetting, CustomSnapshotResourceGroup:t_customsnapshotresourcegroup, RegionalSnapshot:t_regionalsnapshot, GuestCredentials:t_guestcredentials, VmBackupType:t_vmbackuptype, DatastoreFreespaceCheck:t_datastorefreespacecheck, DatastoreFreespaceRequired:t_datastorefreespacerequired, CustomSnapshotTags:t_customsnapshottags, IsApplicationAware:t_isapplicationaware, TransportMode:t_transportmode, CollectFileDetailsFromSnapshotCopy:t_collectfiledetailsfromsnapshotcopy, CrossAccount:t_crossaccount}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgamazoncrossaccount(d *schema.ResourceData, r []interface{}) *handler.MsgAmazonCrossAccount {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_shareonly *bool
        if val, ok := tmp["shareonly"].(bool); ok {
            t_shareonly = new(bool)
            t_shareonly = &val
        }
        var t_fullcopy *bool
        if val, ok := tmp["fullcopy"].(bool); ok {
            t_fullcopy = new(bool)
            t_fullcopy = &val
        }
        var t_destinationaccount *handler.MsgIdName
        if val, ok := tmp["destinationaccount"].([]interface{}); ok {
            t_destinationaccount = build_vmgroup_v2_msgidname(d, val)
        }
        return &handler.MsgAmazonCrossAccount{ShareOnly:t_shareonly, FullCopy:t_fullcopy, DestinationAccount:t_destinationaccount}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgresourcetagset_array(d *schema.ResourceData, r []interface{}) []handler.MsgresourceTagSet {
    if r != nil {
        tmp := make([]handler.MsgresourceTagSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_name *string
            if val, ok := raw_a["name"].(string); ok {
                t_name = new(string)
                t_name = &val
            }
            var t_value *string
            if val, ok := raw_a["value"].(string); ok {
                t_value = new(string)
                t_value = &val
            }
            tmp[a] = handler.MsgresourceTagSet{Name:t_name, Value:t_value}
        }
        return tmp
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgactivitycontroloptions(d *schema.ResourceData, r []interface{}) *handler.MsgActivityControlOptions {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_restoreactivitycontroloptions *handler.MsgbackupActivityControlOptionsProp
        if val, ok := tmp["restoreactivitycontroloptions"].([]interface{}); ok {
            t_restoreactivitycontroloptions = build_vmgroup_v2_msgbackupactivitycontroloptionsprop(d, val)
        }
        var t_backupactivitycontroloptions *handler.MsgbackupActivityControlOptionsProp
        if val, ok := tmp["backupactivitycontroloptions"].([]interface{}); ok {
            t_backupactivitycontroloptions = build_vmgroup_v2_msgbackupactivitycontroloptionsprop(d, val)
        }
        var t_enablebackup *bool
        if val, ok := tmp["enablebackup"].(bool); ok {
            t_enablebackup = new(bool)
            t_enablebackup = &val
        }
        var t_enablerestore *bool
        if val, ok := tmp["enablerestore"].(bool); ok {
            t_enablerestore = new(bool)
            t_enablerestore = &val
        }
        return &handler.MsgActivityControlOptions{RestoreActivityControlOptions:t_restoreactivitycontroloptions, BackupActivityControlOptions:t_backupactivitycontroloptions, EnableBackup:t_enablebackup, EnableRestore:t_enablerestore}
    } else {
        return nil
    }
}

func build_vmgroup_v2_msgbackupactivitycontroloptionsprop(d *schema.ResourceData, r []interface{}) *handler.MsgbackupActivityControlOptionsProp {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_timezone *handler.MsgIdName
        if val, ok := tmp["timezone"].([]interface{}); ok {
            t_timezone = build_vmgroup_v2_msgidname(d, val)
        }
        var t_enableafterdelay *bool
        if val, ok := tmp["enableafterdelay"].(bool); ok {
            t_enableafterdelay = new(bool)
            t_enableafterdelay = &val
        }
        var t_delaytime *string
        if val, ok := tmp["delaytime"].(string); ok {
            t_delaytime = new(string)
            t_delaytime = &val
        }
        return &handler.MsgbackupActivityControlOptionsProp{TimeZone:t_timezone, EnableAfterDelay:t_enableafterdelay, DelayTime:t_delaytime}
    } else {
        return nil
    }
}

func serialize_vmgroup_v2_msgmeditechpropresp(data *handler.MsgmeditechPropResp) map[string]interface{} {
    val := make(map[string]interface{})
    if data.SystemName != nil {
        val["systemname"] = data.SystemName
    }
    if data.ListenerIP != nil {
        val["listenerip"] = data.ListenerIP
    }
    if data.UserAccount != nil {
        val["useraccount"] = serialize_vmgroup_v2_msgusernamepassword(data.UserAccount)
    }
    if data.ListenerPort != nil {
        val["listenerport"] = data.ListenerPort
    }
    if data.MBFtimeout != nil {
        val["mbftimeout"] = data.MBFtimeout
    }
    return val
}

func serialize_vmgroup_v2_msgusernamepassword(data *handler.MsgUserNamePassword) map[string]interface{} {
    val := make(map[string]interface{})
    if data.Password != nil {
        val["password"] = data.Password
    }
    if data.Name != nil {
        val["name"] = data.Name
    }
    return val
}

func serialize_vmgroup_v2_msgvmappvalidation(data *handler.MsgvmAppValidation) map[string]interface{} {
    val := make(map[string]interface{})
    if data.RecoveryTarget != nil {
        val["recoverytarget"] = serialize_vmgroup_v2_msgidname(data.RecoveryTarget)
    }
    if data.Schedule != nil {
        val["schedule"] = serialize_vmgroup_v2_msgvalidationscheduleobject(data.Schedule)
    }
    if data.MaximumNoOfThreads != nil {
        val["maximumnoofthreads"] = data.MaximumNoOfThreads
    }
    if data.GuestCredentials != nil {
        val["guestcredentials"] = serialize_vmgroup_v2_msgguestcredentialinfo(data.GuestCredentials)
    }
    if data.KeepValidatedVMsRunning != nil {
        val["keepvalidatedvmsrunning"] = data.KeepValidatedVMsRunning
    }
    if data.ValidateVMBackups != nil {
        val["validatevmbackups"] = data.ValidateVMBackups
    }
    if data.UseSourceVmESXToMount != nil {
        val["usesourcevmesxtomount"] = data.UseSourceVmESXToMount
    }
    if data.CustomValidationScript != nil {
        val["customvalidationscript"] = serialize_vmgroup_v2_msgappvalidationscript(data.CustomValidationScript)
    }
    if data.Copy != nil {
        val["copy"] = serialize_vmgroup_v2_msgplansourcecopy(data.Copy)
    }
    return val
}

func serialize_vmgroup_v2_msgplansourcecopy(data *handler.MsgPlanSourceCopy) map[string]interface{} {
    val := make(map[string]interface{})
    if data.IsMirrorCopy != nil {
        val["ismirrorcopy"] = data.IsMirrorCopy
    }
    if data.SnapCopyType != nil {
        val["snapcopytype"] = data.SnapCopyType
    }
    if data.IsDefault != nil {
        val["isdefault"] = data.IsDefault
    }
    if data.CopyPrecedence != nil {
        val["copyprecedence"] = data.CopyPrecedence
    }
    if data.IsSnapCopy != nil {
        val["issnapcopy"] = data.IsSnapCopy
    }
    if data.CopyType != nil {
        val["copytype"] = data.CopyType
    }
    if data.DefaultReplicaCopy != nil {
        val["defaultreplicacopy"] = data.DefaultReplicaCopy
    }
    if data.IsActive != nil {
        val["isactive"] = data.IsActive
    }
    if data.ArrayReplicaCopy != nil {
        val["arrayreplicacopy"] = data.ArrayReplicaCopy
    }
    if data.BackupDestination != nil {
        val["backupdestination"] = serialize_vmgroup_v2_msgidname(data.BackupDestination)
    }
    return val
}

func serialize_vmgroup_v2_msgidname(data *handler.MsgIdName) map[string]interface{} {
    val := make(map[string]interface{})
    if data.Name != nil {
        val["name"] = data.Name
    }
    if data.Id != nil {
        val["id"] = data.Id
    }
    return val
}

func serialize_vmgroup_v2_msgappvalidationscript(data *handler.MsgappValidationScript) map[string]interface{} {
    val := make(map[string]interface{})
    if data.Windows != nil {
        val["windows"] = serialize_vmgroup_v2_msgvalidationscript(data.Windows)
    }
    if data.Unix != nil {
        val["unix"] = serialize_vmgroup_v2_msgvalidationscript(data.Unix)
    }
    return val
}

func serialize_vmgroup_v2_msgvalidationscript(data *handler.MsgValidationScript) map[string]interface{} {
    val := make(map[string]interface{})
    if data.Path != nil {
        val["path"] = data.Path
    }
    if data.UNCCredentials != nil {
        val["unccredentials"] = serialize_vmgroup_v2_msgusernamepassword(data.UNCCredentials)
    }
    if data.UNCSavedCredentials != nil {
        val["uncsavedcredentials"] = serialize_vmgroup_v2_msgidname(data.UNCSavedCredentials)
    }
    if data.Arguments != nil {
        val["arguments"] = data.Arguments
    }
    if data.IsDisabled != nil {
        val["isdisabled"] = data.IsDisabled
    }
    if data.IsLocal != nil {
        val["islocal"] = data.IsLocal
    }
    return val
}

func serialize_vmgroup_v2_msgguestcredentialinfo(data *handler.MsgguestCredentialInfo) map[string]interface{} {
    val := make(map[string]interface{})
    if data.Credentials != nil {
        val["credentials"] = serialize_vmgroup_v2_msgusernamepassword(data.Credentials)
    }
    if data.SavedCredentials != nil {
        val["savedcredentials"] = serialize_vmgroup_v2_msgidname(data.SavedCredentials)
    }
    return val
}

func serialize_vmgroup_v2_msgvalidationscheduleobject(data *handler.MsgValidationScheduleObject) map[string]interface{} {
    val := make(map[string]interface{})
    if data.IsScheduleEnabled != nil {
        val["isscheduleenabled"] = data.IsScheduleEnabled
    }
    if data.Description != nil {
        val["description"] = data.Description
    }
    if data.Id != nil {
        val["id"] = data.Id
    }
    if data.TaskId != nil {
        val["taskid"] = data.TaskId
    }
    return val
}

func serialize_vmgroup_v2_msgsnapcopyinfo(data *handler.MsgsnapCopyInfo) map[string]interface{} {
    val := make(map[string]interface{})
    if data.UseSeparateProxyForSnapToTape != nil {
        val["useseparateproxyforsnaptotape"] = data.UseSeparateProxyForSnapToTape
    }
    if data.SnapEngine != nil {
        val["snapengine"] = serialize_vmgroup_v2_msgidname(data.SnapEngine)
    }
    if data.IsIndependentDisksEnabled != nil {
        val["isindependentdisksenabled"] = data.IsIndependentDisksEnabled
    }
    if data.BackupCopyInterface != nil {
        val["backupcopyinterface"] = data.BackupCopyInterface
    }
    if data.EnableHardwareSnapshot != nil {
        val["enablehardwaresnapshot"] = data.EnableHardwareSnapshot
    }
    if data.SnapMountProxy != nil {
        val["snapmountproxy"] = serialize_vmgroup_v2_msgidname(data.SnapMountProxy)
    }
    if data.VmApplicationUserName != nil {
        val["vmapplicationusername"] = data.VmApplicationUserName
    }
    if data.SnapMountESXHost != nil {
        val["snapmountesxhost"] = data.SnapMountESXHost
    }
    if data.IsRawDeviceMapsEnabled != nil {
        val["israwdevicemapsenabled"] = data.IsRawDeviceMapsEnabled
    }
    return val
}

func serialize_vmgroup_v2_msgvmcontentset_array(data []handler.MsgvmContentSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].RuleGroups != nil {
            val[i]["rulegroups"] = serialize_vmgroup_v2_msgrulegroupcontentset_array(data[i].RuleGroups)
        }
        if data[i].VirtualMachines != nil {
            val[i]["virtualmachines"] = serialize_vmgroup_v2_msgvirtualmachinecontentset_array(data[i].VirtualMachines)
        }
        if data[i].Overwrite != nil {
            val[i]["overwrite"] = data[i].Overwrite
        }
    }
    return val
}

func serialize_vmgroup_v2_msgvirtualmachinecontentset_array(data []handler.MsgVirtualMachinecontentSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].GuestCredentialAssocId != nil {
            val[i]["guestcredentialassocid"] = data[i].GuestCredentialAssocId
        }
        if data[i].GuestCredentials != nil {
            val[i]["guestcredentials"] = serialize_vmgroup_v2_msgvmguestcredentials(data[i].GuestCredentials)
        }
        if data[i].Name != nil {
            val[i]["name"] = data[i].Name
        }
        if data[i].GUID != nil {
            val[i]["guid"] = data[i].GUID
        }
        if data[i].Type != nil {
            val[i]["type"] = data[i].Type
        }
        if data[i].ExistingCredential != nil {
            val[i]["existingcredential"] = serialize_vmgroup_v2_msgvmexistingcredential(data[i].ExistingCredential)
        }
    }
    return val
}

func serialize_vmgroup_v2_msgvmexistingcredential(data *handler.MsgVMExistingCredential) map[string]interface{} {
    val := make(map[string]interface{})
    if data.CredentialId != nil {
        val["credentialid"] = data.CredentialId
    }
    if data.CredentialName != nil {
        val["credentialname"] = data.CredentialName
    }
    return val
}

func serialize_vmgroup_v2_msgvmguestcredentials(data *handler.MsgVMGuestCredentials) map[string]interface{} {
    val := make(map[string]interface{})
    if data.Password != nil {
        val["password"] = data.Password
    }
    if data.UserName != nil {
        val["username"] = data.UserName
    }
    return val
}

func serialize_vmgroup_v2_msgrulegroupcontentset_array(data []handler.MsgRuleGroupContentSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].MatchRule != nil {
            val[i]["matchrule"] = data[i].MatchRule
        }
        if data[i].Rules != nil {
            val[i]["rules"] = serialize_vmgroup_v2_msgrulecontentset_array(data[i].Rules)
        }
    }
    return val
}

func serialize_vmgroup_v2_msgrulecontentset_array(data []handler.MsgRuleContentSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].Condition != nil {
            val[i]["condition"] = data[i].Condition
        }
        if data[i].DisplayName != nil {
            val[i]["displayname"] = data[i].DisplayName
        }
        if data[i].Name != nil {
            val[i]["name"] = data[i].Name
        }
        if data[i].GUID != nil {
            val[i]["guid"] = data[i].GUID
        }
        if data[i].Type != nil {
            val[i]["type"] = data[i].Type
        }
        if data[i].Value != nil {
            val[i]["value"] = data[i].Value
        }
    }
    return val
}

func serialize_vmgroup_v2_msgsecurityassocset_array(data []handler.MsgSecurityAssocSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].Role != nil {
            val[i]["role"] = serialize_vmgroup_v2_msgidname(data[i].Role)
        }
        if data[i].IsCreatorAssociation != nil {
            val[i]["iscreatorassociation"] = data[i].IsCreatorAssociation
        }
        if data[i].ExternalUserGroup != nil {
            val[i]["externalusergroup"] = serialize_vmgroup_v2_msgexternalusergroup(data[i].ExternalUserGroup)
        }
        if data[i].PermissionList != nil {
            val[i]["permissionlist"] = serialize_vmgroup_v2_msgpermissionrespset_array(data[i].PermissionList)
        }
        if data[i].User != nil {
            val[i]["user"] = serialize_vmgroup_v2_msgidname(data[i].User)
        }
        if data[i].UserGroup != nil {
            val[i]["usergroup"] = serialize_vmgroup_v2_msgidname(data[i].UserGroup)
        }
    }
    return val
}

func serialize_vmgroup_v2_msgpermissionrespset_array(data []handler.MsgPermissionRespSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].PermissionId != nil {
            val[i]["permissionid"] = data[i].PermissionId
        }
        if data[i].Exclude != nil {
            val[i]["exclude"] = data[i].Exclude
        }
        if data[i].Type != nil {
            val[i]["type"] = data[i].Type
        }
        if data[i].CategoryName != nil {
            val[i]["categoryname"] = data[i].CategoryName
        }
        if data[i].CategoryId != nil {
            val[i]["categoryid"] = data[i].CategoryId
        }
        if data[i].PermissionName != nil {
            val[i]["permissionname"] = data[i].PermissionName
        }
    }
    return val
}

func serialize_vmgroup_v2_msgexternalusergroup(data *handler.MsgexternalUserGroup) map[string]interface{} {
    val := make(map[string]interface{})
    if data.ProviderId != nil {
        val["providerid"] = data.ProviderId
    }
    if data.Name != nil {
        val["name"] = data.Name
    }
    if data.Id != nil {
        val["id"] = data.Id
    }
    if data.ProviderName != nil {
        val["providername"] = data.ProviderName
    }
    return val
}

func serialize_vmgroup_v2_msgvmdiskfilterpropset_array(data []handler.MsgvmDiskFilterPropSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].Rules != nil {
            val[i]["rules"] = serialize_vmgroup_v2_msgvmdiskfilterset_array(data[i].Rules)
        }
        if data[i].Overwrite != nil {
            val[i]["overwrite"] = data[i].Overwrite
        }
    }
    return val
}

func serialize_vmgroup_v2_msgvmdiskfilterset_array(data []handler.MsgvmDiskFilterSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].Condition != nil {
            val[i]["condition"] = data[i].Condition
        }
        if data[i].VmName != nil {
            val[i]["vmname"] = data[i].VmName
        }
        if data[i].Name != nil {
            val[i]["name"] = data[i].Name
        }
        if data[i].FilterType != nil {
            val[i]["filtertype"] = data[i].FilterType
        }
        if data[i].Overwrite != nil {
            val[i]["overwrite"] = data[i].Overwrite
        }
        if data[i].Value != nil {
            val[i]["value"] = data[i].Value
        }
        if data[i].VmGuid != nil {
            val[i]["vmguid"] = data[i].VmGuid
        }
    }
    return val
}

func serialize_vmgroup_v2_msgvmgroupsettings(data *handler.MsgvmGroupSettings) map[string]interface{} {
    val := make(map[string]interface{})
    if data.AutoDetectVMOwner != nil {
        val["autodetectvmowner"] = data.AutoDetectVMOwner
    }
    if data.CollectFileDetailsforGranularRecovery != nil {
        val["collectfiledetailsforgranularrecovery"] = data.CollectFileDetailsforGranularRecovery
    }
    if data.NoOfReaders != nil {
        val["noofreaders"] = data.NoOfReaders
    }
    if data.UseChangedBlockTrackingOnVM != nil {
        val["usechangedblocktrackingonvm"] = data.UseChangedBlockTrackingOnVM
    }
    if data.JobStartTime != nil {
        val["jobstarttime"] = data.JobStartTime
    }
    if data.UseVMCheckpointSetting != nil {
        val["usevmcheckpointsetting"] = data.UseVMCheckpointSetting
    }
    if data.CustomSnapshotResourceGroup != nil {
        val["customsnapshotresourcegroup"] = data.CustomSnapshotResourceGroup
    }
    if data.RegionalSnapshot != nil {
        val["regionalsnapshot"] = data.RegionalSnapshot
    }
    if data.GuestCredentials != nil {
        val["guestcredentials"] = serialize_vmgroup_v2_msgguestcredentialinfo(data.GuestCredentials)
    }
    if data.VmBackupType != nil {
        val["vmbackuptype"] = data.VmBackupType
    }
    if data.DatastoreFreespaceCheck != nil {
        val["datastorefreespacecheck"] = data.DatastoreFreespaceCheck
    }
    if data.DatastoreFreespaceRequired != nil {
        val["datastorefreespacerequired"] = data.DatastoreFreespaceRequired
    }
    if data.CustomSnapshotTags != nil {
        val["customsnapshottags"] = serialize_vmgroup_v2_msgresourcetagset_array(data.CustomSnapshotTags)
    }
    if data.IsApplicationAware != nil {
        val["isapplicationaware"] = data.IsApplicationAware
    }
    if data.TransportMode != nil {
        val["transportmode"] = data.TransportMode
    }
    if data.CollectFileDetailsFromSnapshotCopy != nil {
        val["collectfiledetailsfromsnapshotcopy"] = data.CollectFileDetailsFromSnapshotCopy
    }
    if data.CrossAccount != nil {
        val["crossaccount"] = serialize_vmgroup_v2_msgamazoncrossaccount(data.CrossAccount)
    }
    return val
}

func serialize_vmgroup_v2_msgamazoncrossaccount(data *handler.MsgAmazonCrossAccount) map[string]interface{} {
    val := make(map[string]interface{})
    if data.ShareOnly != nil {
        val["shareonly"] = data.ShareOnly
    }
    if data.FullCopy != nil {
        val["fullcopy"] = data.FullCopy
    }
    if data.DestinationAccount != nil {
        val["destinationaccount"] = serialize_vmgroup_v2_msgidname(data.DestinationAccount)
    }
    return val
}

func serialize_vmgroup_v2_msgresourcetagset_array(data []handler.MsgresourceTagSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].Name != nil {
            val[i]["name"] = data[i].Name
        }
        if data[i].Value != nil {
            val[i]["value"] = data[i].Value
        }
    }
    return val
}

func serialize_vmgroup_v2_msgactivitycontroloptions(data *handler.MsgActivityControlOptions) map[string]interface{} {
    val := make(map[string]interface{})
    if data.RestoreActivityControlOptions != nil {
        val["restoreactivitycontroloptions"] = serialize_vmgroup_v2_msgbackupactivitycontroloptionsprop(data.RestoreActivityControlOptions)
    }
    if data.BackupActivityControlOptions != nil {
        val["backupactivitycontroloptions"] = serialize_vmgroup_v2_msgbackupactivitycontroloptionsprop(data.BackupActivityControlOptions)
    }
    if data.EnableBackup != nil {
        val["enablebackup"] = data.EnableBackup
    }
    if data.EnableRestore != nil {
        val["enablerestore"] = data.EnableRestore
    }
    return val
}

func serialize_vmgroup_v2_msgbackupactivitycontroloptionsprop(data *handler.MsgbackupActivityControlOptionsProp) map[string]interface{} {
    val := make(map[string]interface{})
    if data.TimeZone != nil {
        val["timezone"] = serialize_vmgroup_v2_msgidname(data.TimeZone)
    }
    if data.EnableAfterDelay != nil {
        val["enableafterdelay"] = data.EnableAfterDelay
    }
    if data.DelayTime != nil {
        val["delaytime"] = data.DelayTime
    }
    return val
}
