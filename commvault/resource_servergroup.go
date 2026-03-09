package commvault

import (
    "fmt"
    "strconv"
    "strings"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServerGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateServerGroup,
        Read:   resourceReadServerGroup,
        Update: resourceUpdateServerGroup,
        Delete: resourceDeleteServerGroup,

        Schema: map[string]*schema.Schema{
            "virtualassociation": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "clientassociation": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "associationrule": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "name of the VM to be added as content",
                                                },
                                                "guid": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "GUID of the VM to be added as content",
                                                },
                                                "type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "[NONE, SERVER, RES_POOL, VAPP, DATACENTER, FOLDER, CLUSTER, DATASTORE, DATASTORE_CLUSTER, VM, VM_NAME, VM_GUEST_OS, VM_GUEST_HOST_NAME, CLUSTER_SHARED_VOLUMES, LOCAL_DISK, CLUSTER_DISK, UNPROTECTED_VMS, ROOT, FILE_SERVER, SMB_SHARE, TYPES_FOLDER, VM_FOLDER, SERVER_FOLDER, TEMPLATE_FOLDER, STORAGE_REPOSITORY_FOLDER, VAPPFOLDER, DATACENTER_FOLDER, CLUSTER_FOLDER, VM_POWER_STATE, VM_NOTES, VM_CUSTOM_ATTRIBUTE, NETWORK, USER, VM_TEMPLATE, TAG, TAG_CATEGORY, SUBCLIENT, CLIENT_GROUP, PROTECTION_DOMAIN, CONSISTENCY_GROUP, INSTANCE_SIZE, ORGANIZATION, IMAGES, STORAGE_POLICY, DATABASE, TABLE, PROJECT, SELECTOR, MANAGED_BY, REPLICATION_MODE, METADATATAG, CATALOG, VAPPTEMPLATE, VOLUME, VM_BADGE, REGION]",
                                                },
                                            },
                                        },
                                    },
                                    "associatedclient": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "clientid": {
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
                        "virtualservers": {
                            Type:        schema.TypeSet,
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
                        "rules": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "guestcredentialassocid": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Credential association ID given to link entity with credential id.",
                                    },
                                    "guestcredentials": {
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
                                                    Description: "",
                                                },
                                                "username": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                            },
                                        },
                                    },
                                    "matchrule": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Enum which specifies the whether to match all rules or any of the rules [ALL, ANY]",
                                    },
                                    "description": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Description of the rule group",
                                    },
                                    "rules": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "condition": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Operation type for VM rules/filters [CONTAINS, DOES_NOT_CONTAIN, DOES_NOT_EQUAL, ENDS_WITH, EQUALS, STARTS_WITH]",
                                                },
                                                "name": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "name of the VM to be added as content",
                                                },
                                                "type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "[NONE, SERVER, RES_POOL, VAPP, DATACENTER, FOLDER, CLUSTER, DATASTORE, DATASTORE_CLUSTER, VM, VM_NAME, VM_GUEST_OS, VM_GUEST_HOST_NAME, CLUSTER_SHARED_VOLUMES, LOCAL_DISK, CLUSTER_DISK, UNPROTECTED_VMS, ROOT, FILE_SERVER, SMB_SHARE, TYPES_FOLDER, VM_FOLDER, SERVER_FOLDER, TEMPLATE_FOLDER, STORAGE_REPOSITORY_FOLDER, VAPPFOLDER, DATACENTER_FOLDER, CLUSTER_FOLDER, VM_POWER_STATE, VM_NOTES, VM_CUSTOM_ATTRIBUTE, NETWORK, USER, VM_TEMPLATE, TAG, TAG_CATEGORY, SUBCLIENT, CLIENT_GROUP, PROTECTION_DOMAIN, CONSISTENCY_GROUP, INSTANCE_SIZE, ORGANIZATION, IMAGES, STORAGE_POLICY, DATABASE, TABLE, PROJECT, SELECTOR, MANAGED_BY, REPLICATION_MODE, METADATATAG, CATALOG, VAPPTEMPLATE, VOLUME, VM_BADGE, REGION]",
                                                },
                                                "value": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "value for the few type of VM Content like powerstate",
                                                },
                                            },
                                        },
                                    },
                                    "existingcredential": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "credentialid": {
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
                                                "credentialname": {
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
                        "virtualinstance": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "[NONE, VMW, MICROSOFT, XEN_SERVER, AMAZON, AZURE, REDHAT, AZURE_V2, SCVMM, NUTANIX, ORACLE_VM, DOCKER, OPENSTACK, ORACLE_CLOUD, FUSIONCOMPUTE, VCLOUD, GOOGLE_CLOUD, AZURE_STACK, ALIBABA_CLOUD, ORACLE_CLOUD_INFRASTRUCTURE, KUBERNETES, REDHAT_OPENSHIFT, MONGODB_ATLAS, PROXMOX, KUBERNETES_AKS, AZURE_STACK_HCI, KUBERNETES_EKS, MORPHEUS, KUBERNETES_GKE, KUBERNETES_OKE, NUTANIX_PRISM_CENTRAL, APACHE_CLOUDSTACK, VMWARE_CLOUD_FOUNDATION]",
                        },
                    },
                },
            },
            "automaticassociation": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "servergrouprule": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "match": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "[AND, OR, NOT]",
                                    },
                                    "rulegroup": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "List of rule groups",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "match": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "[AND, OR, NOT]",
                                                },
                                                "rules": {
                                                    Type:        schema.TypeSet,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "List of rules",
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "rulesecvalue": {
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "Secondary value used for BETWEEN and NOT_BETWEEN matchCondition. For ruleName which have enum values (like OS_TYPE), this is used to store displayName.",
                                                            },
                                                            "rulevalue": {
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "Primary value for the rule",
                                                            },
                                                            "matchcondition": {
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "[IS_TRUE, IS_FALSE, CONTAINS, DOES_NOT_CONTAIN, IS, IS_NOT, STARTS_WITH, ENDS_WITH, EQUAL_TO, NOT_EQUAL_TO, GREATER_THAN, LESS_THAN, GREATER_THAN_OR_EQUAL_TO, LESS_THAN_OR_EQUAL_TO, BETWEEN, NOT_IN, ANY_IN_SELECTION, NOT_IN_SELECTION, IN, NOT_BETWEEN]",
                                                            },
                                                            "rulename": {
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "[APP_TYPE, CLIENT_GROUP, CLIENT_GROUPLIST, USES_LIBRARY, STORAGE_POLICY, STORAGE_POLICYLIST, BACKUP_ACTIVITY_ENABLED_FOR_CLIENT, CASE_MANAGER_CLIENTS, CLIENT_PROP, CLIENT_ACTS_AS_PROXY, CLIENT_ASSOCIATED_WITH_PLAN, CLIENT_BACKUP_ACTIVITY_ENABLED, CLIENT_BY_SCHEDULE_INTERVAL, CLIENT_CIDR_RANGE, CLIENTS_RELEASE16PLUS_SP_LEVEL_CONSTRAINT, CLIENT_DESCRIPTION, CLIENT_DISPLAY_NAME, CLIENT_EXCLUDED_FROM_SLA, CLIENTGROUP_NAME, CLIENT_CIDR_RANGE_IPV6, CLIENT_ISDELETEDVM, CLIENT_HAS_CONNECTIVITY_ISSUES, CLIENT_NAME, CLIENT_NEEDS_UPDATES, CLIENT_OFFLINE, CLIENT_ONLINE, CLIENT_ONLINE_IN_LAST_30_DAYS, CLIENT_ONLINE_IN_LAST_N_DAYS, CLIENT_USES_STORAGEPOLICY, CLIENT_VERSION, COMPARE_CLIENT_CS_VERSION, CLIENT_ASSOCIATED_ADUSERGROUP, IS_COMPANY_CLIENT, CLIENT_ASSOC_SCHDPOLICY, CLIENT_ASSOCIATED_TO_TOPOLOGY_WITH_REGION, VSA_BACKEDUP_CLIENT, CLIENT_INDEX_MAINFOLIST, CLIENT_OS_INFOLIST, CLIENT_BY_PERMISSION, CLIENT_BY_ROLE, VSA_DISCOVER_CLIENT, CLIENTGROUP_TAGLIST, CLIENT_MEETS_SLA, CLIENTS_OF_RESELLER, USERGROUPLIST, CLIENT_IDA_ASSOCIATED_ADUSERGROUP, CLIENT_WITH_ASSOCIATED_REGION, CLIENTS_WITH_ATTACHED_STORAGE, CLIENT_WITH_BACKUP_SCHEDULE, CLIENT_WITH_ENABLED_BACKUP_SCHEDULE, CLIENT_ENCRYPTION_STATE, CLIENT_FETSIZE_LE_10GB, CLIENT_IMPROPER_DECONF_SUBCLIENT, CLIENT_INDEXINGV1, CLIENT_INDEXINGV2, CLIENTINFRASTRUCTUREROLE_LIST, CLIENT_LICENSELIST, CLIENT_LONG_RUNNING_JOBS_N_DAYS, CLIENT_WITH_NO_ARCHIVE_DATA, CLIENTS_WITH_ONEPASS_ENABLED, CLIENT_WITH_ASSOCIATED_SP, CLIENT_WITH_SYNTHETICFULL_BACKUP, CLIENT_TAGLIST, IS_COMMCELL_CLIENT, COMPANY_CLIENT_ASSOCIATIONS, COMPANY_INSTALLED_CLIENT_ASSOCIATIONS, ANALYZER_SERVERS_CLIENT, DAYS_SINCE_CLIENT_CREATED, DAYS_SINCE_LAST_CLIENT_BKP, EXTERNAL_GROUP_CLIENT_OWNER, HAC_CLUSTER_CLIENT, HOST_NAME, MACLIENTS_WITH_STORAGEPOOLS, CLIENT_OWNER_INACTIVE_ADUSER, SOLR_SERVERS_CLIENT, LOCAL_GROUP_CLIENT_OWNER, MEDIAAGENT_HAS_LUCENE_INDEX_ROLELIST, MEDIAAGENT_HAS_LUCENE_INDEX_ROLES, MAS_FOR_CLIENTS_IN_GROUP, MAS_FOR_CLIENTS_IN_GROUPLIST, CLIENT_GATEWAY_FOR_INSTALLATION, OS_TYPE, OS_VERSION, PACKAGE_INSTALLED, PRODUCT_VERSION, PSEUDO_CLIENTS, RESTORE_ACTIVITY_ENABLED_FOR_CLIENT, SNAP_BACKUP_CLIENTS, CLIENT_CONSECUTIVE_BACKUP_FAILURES, SUBCLIENT_NAME, TIMEZONE, TIMEZONELIST, TIMEZONE_REGIONLIST, USER_CLIENT_OWNER, USER_CLIENT_ASSOCIATIONS, USER_DESCRIPTION_CONTAINS, USERGROUP_CLIENT_ASSOCIATIONS, USERGROUP_DESCRIPTION_CONTAINS, VM_NO_CONTENT, VMHYPER_IN_CLIENTGROUP]",
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
                        "confirmrulechange": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Confirm that server group rule impacts server of a specific company smart client group only",
                        },
                        "servergroupid": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "Id of Smart Client Group",
                        },
                        "clientscope": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "entityinfo": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "guid": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "",
                                                },
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
                                    "clientscopetype": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "[COMMCELL, COMPANY, USER, USERGROUP]",
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "name": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "",
            },
            "globalconfiginfo": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Only applicable to Global CommCells. Not applicable for SaaS.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "scopefilterquery": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "CommCellEntityCache filter query string using for filtering the scope",
                        },
                        "companies": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "List of companies where the global configuration should be applied",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "guid": {
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
                        "applyonallcommcells": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Decides whether the global configuration should be applied to all the Service commcells, including the newly created ones",
                        },
                        "commcells": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "List of Service CommCells where the global configuration should be applied",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "displayname": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "",
                                    },
                                    "guid": {
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
                        "scope": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "The entity level at which the config has to be applied.",
                        },
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                        },
                        "applyonallcompanies": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Decides whether the global configuration should be applied to all the companies, including the newly created ones",
                        },
                        "actiononlocalentity": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Action that will be taken on the local entity that has the same name as the global entity that needs to be created [CREATE_NEW, TAKE_OVER, FAIL_IF_EXIST]",
                        },
                    },
                },
            },
            "description": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "",
            },
            "manualassociation": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "associatedservers": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "List of associated servers",
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
                        "operationtype": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "List operation type [NONE, OVERWRITE, ADD, DELETE, CLEAR]",
                        },
                    },
                },
            },
            "servergrouptype": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "[MANUAL, AUTOMATIC, VIRTUAL_MACHINE]",
            },
            "servergroup": {
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
    }
}

func resourceCreateServerGroup(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/ServerGroup
    var response_id = strconv.Itoa(0)
    var t_virtualassociation *handler.MsgVirtualServerGroupAssociationDetails
    if val, ok := d.GetOk("virtualassociation"); ok {
        t_virtualassociation = build_servergroup_msgvirtualservergroupassociationdetails(d, val.([]interface{}))
    }
    var t_automaticassociation *handler.MsgAutomaticServerGroupAssociationDetails
    if val, ok := d.GetOk("automaticassociation"); ok {
        t_automaticassociation = build_servergroup_msgautomaticservergroupassociationdetails(d, val.([]interface{}))
    }
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_globalconfiginfo *handler.MsgCreateGlobalConfigInfo
    if val, ok := d.GetOk("globalconfiginfo"); ok {
        t_globalconfiginfo = build_servergroup_msgcreateglobalconfiginfo(d, val.([]interface{}))
    }
    var t_description *string
    if val, ok := d.GetOk("description"); ok {
        t_description = handler.ToStringValue(val, false)
    }
    var t_manualassociation *handler.MsgManualServerGroupAssociationDetails
    if val, ok := d.GetOk("manualassociation"); ok {
        t_manualassociation = build_servergroup_msgmanualservergroupassociationdetails(d, val.([]interface{}))
    }
    var t_servergrouptype *string
    if val, ok := d.GetOk("servergrouptype"); ok {
        t_servergrouptype = handler.ToStringValue(val, false)
    }
    var req = handler.MsgCreateServerGroupsRequest{VirtualAssociation:t_virtualassociation, AutomaticAssociation:t_automaticassociation, Name:t_name, GlobalConfigInfo:t_globalconfiginfo, Description:t_description, ManualAssociation:t_manualassociation, ServerGroupType:t_servergrouptype}
    resp, err := handler.CvCreateServerGroups(req)
    if err != nil {
        return fmt.Errorf("operation [CreateServerGroups] failed, Error %s", err)
    }
    if resp.ServerGroupInfo != nil {
        if resp.ServerGroupInfo.Id != nil {
            response_id = strconv.Itoa(*resp.ServerGroupInfo.Id)
        }
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateServerGroups] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateServerGroup(d, m)
    }
}

func resourceReadServerGroup(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/ServerGroup/{serverGroupId}
    resp, err := handler.CvGetServerGroupIdDetails(d.Id())
    if err != nil {
        if strings.Contains(err.Error(), "status: 404") {
            handler.LogEntry("debug", "entity not present, removing from state")
            d.SetId("")
            return nil
        }
        return fmt.Errorf("operation [GetServerGroupIdDetails] failed, Error %s", err)
    }
    if rtn, ok := serialize_servergroup_msgidname(d, resp.ServerGroup); ok {
        d.Set("servergroup", rtn)
    } else {
        d.Set("servergroup", make([]map[string]interface{}, 0))
    }
    if resp.Description != nil {
        d.Set("description", resp.Description)
    }
    return nil
}

func resourceUpdateServerGroup(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/ServerGroup/{serverGroupId}
    var t_virtualassociation *handler.MsgVirtualServerGroupAssociationDetails
    if d.HasChange("virtualassociation") {
        val := d.Get("virtualassociation")
        t_virtualassociation = build_servergroup_msgvirtualservergroupassociationdetails(d, val.([]interface{}))
    }
    var t_automaticassociation *handler.MsgAutomaticServerGroupAssociationDetails
    if d.HasChange("automaticassociation") {
        val := d.Get("automaticassociation")
        t_automaticassociation = build_servergroup_msgautomaticservergroupassociationdetails(d, val.([]interface{}))
    }
    var t_servergroup *handler.MsgIdName
    if d.HasChange("servergroup") {
        val := d.Get("servergroup")
        t_servergroup = build_servergroup_msgidname(d, val.([]interface{}))
    }
    var t_description *string
    if d.HasChange("description") {
        val := d.Get("description")
        t_description = handler.ToStringValue(val, false)
    }
    var t_manualassociation *handler.MsgManualServerGroupAssociationDetails
    if d.HasChange("manualassociation") {
        val := d.Get("manualassociation")
        t_manualassociation = build_servergroup_msgmanualservergroupassociationdetails(d, val.([]interface{}))
    }
    var t_servergrouptype *string
    if d.HasChange("servergrouptype") {
        val := d.Get("servergrouptype")
        t_servergrouptype = handler.ToStringValue(val, false)
    }
    var req = handler.MsgUpdateServerGroupAssociationRequest{VirtualAssociation:t_virtualassociation, AutomaticAssociation:t_automaticassociation, ServerGroup:t_servergroup, Description:t_description, ManualAssociation:t_manualassociation, ServerGroupType:t_servergrouptype}
    _, err := handler.CvUpdateServerGroupAssociation(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [UpdateServerGroupAssociation] failed, Error %s", err)
    }
    return resourceReadServerGroup(d, m)
}

func resourceCreateUpdateServerGroup(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/ServerGroup/{serverGroupId}
    var execUpdate bool = false
    var t_servergroup *handler.MsgIdName
    if val, ok := d.GetOk("servergroup"); ok {
        t_servergroup = build_servergroup_msgidname(d, val.([]interface{}))
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgUpdateServerGroupAssociationRequest{ServerGroup:t_servergroup}
        _, err := handler.CvUpdateServerGroupAssociation(req, d.Id())
        if err != nil {
            return fmt.Errorf("operation [UpdateServerGroupAssociation] failed, Error %s", err)
        }
    }
    return resourceReadServerGroup(d, m)
}

func resourceDeleteServerGroup(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/ServerGroup/{serverGroupId}
    _, err := handler.CvDeleteServerGroup(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteServerGroup] failed, Error %s", err)
    }
    return nil
}

func build_servergroup_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_servergroup_msgmanualservergroupassociationdetails(d *schema.ResourceData, r []interface{}) *handler.MsgManualServerGroupAssociationDetails {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_associatedservers []handler.MsgIdNameSet
        if val, ok := tmp["associatedservers"]; ok {
            t_associatedservers = build_servergroup_msgidnameset_array(d, val.(*schema.Set).List())
        }
        var t_operationtype *string
        if val, ok := tmp["operationtype"]; ok {
            t_operationtype = handler.ToStringValue(val, true)
        }
        return &handler.MsgManualServerGroupAssociationDetails{Associatedservers:t_associatedservers, OperationType:t_operationtype}
    } else {
        return nil
    }
}

func build_servergroup_msgidnameset_array(d *schema.ResourceData, r []interface{}) []handler.MsgIdNameSet {
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

func build_servergroup_msgautomaticservergroupassociationdetails(d *schema.ResourceData, r []interface{}) *handler.MsgAutomaticServerGroupAssociationDetails {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_servergrouprule *handler.MsgServerRuleGroups
        if val, ok := tmp["servergrouprule"]; ok {
            t_servergrouprule = build_servergroup_msgserverrulegroups(d, val.([]interface{}))
        }
        var t_confirmrulechange *bool
        if val, ok := tmp["confirmrulechange"]; ok {
            t_confirmrulechange = handler.ToBooleanValue(val, true)
        }
        var t_servergroupid *int
        if val, ok := tmp["servergroupid"]; ok {
            t_servergroupid = handler.ToIntValue(val, true)
        }
        var t_clientscope *handler.MsgClientScopeDetails
        if val, ok := tmp["clientscope"]; ok {
            t_clientscope = build_servergroup_msgclientscopedetails(d, val.([]interface{}))
        }
        return &handler.MsgAutomaticServerGroupAssociationDetails{ServerGroupRule:t_servergrouprule, ConfirmRuleChange:t_confirmrulechange, ServerGroupId:t_servergroupid, ClientScope:t_clientscope}
    } else {
        return nil
    }
}

func build_servergroup_msgclientscopedetails(d *schema.ResourceData, r []interface{}) *handler.MsgClientScopeDetails {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_entityinfo *handler.MsgIdNameGUID
        if val, ok := tmp["entityinfo"]; ok {
            t_entityinfo = build_servergroup_msgidnameguid(d, val.([]interface{}))
        }
        var t_clientscopetype *string
        if val, ok := tmp["clientscopetype"]; ok {
            t_clientscopetype = handler.ToStringValue(val, true)
        }
        return &handler.MsgClientScopeDetails{EntityInfo:t_entityinfo, ClientScopeType:t_clientscopetype}
    } else {
        return nil
    }
}

func build_servergroup_msgidnameguid(d *schema.ResourceData, r []interface{}) *handler.MsgIdNameGUID {
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
        var t_id *int
        if val, ok := tmp["id"]; ok {
            t_id = handler.ToIntValue(val, true)
        }
        return &handler.MsgIdNameGUID{GUID:t_guid, Name:t_name, Id:t_id}
    } else {
        return nil
    }
}

func build_servergroup_msgserverrulegroups(d *schema.ResourceData, r []interface{}) *handler.MsgServerRuleGroups {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_match *string
        if val, ok := tmp["match"]; ok {
            t_match = handler.ToStringValue(val, true)
        }
        var t_rulegroup []handler.MsgServerRuleGroupSet
        if val, ok := tmp["rulegroup"]; ok {
            t_rulegroup = build_servergroup_msgserverrulegroupset_array(d, val.(*schema.Set).List())
        }
        return &handler.MsgServerRuleGroups{Match:t_match, RuleGroup:t_rulegroup}
    } else {
        return nil
    }
}

func build_servergroup_msgserverrulegroupset_array(d *schema.ResourceData, r []interface{}) []handler.MsgServerRuleGroupSet {
    if r != nil {
        tmp := make([]handler.MsgServerRuleGroupSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_match *string
            if val, ok := raw_a["match"]; ok {
                t_match = handler.ToStringValue(val, true)
            }
            var t_rules []handler.MsgServerRuleSet
            if val, ok := raw_a["rules"]; ok {
                t_rules = build_servergroup_msgserverruleset_array(d, val.(*schema.Set).List())
            }
            tmp[a] = handler.MsgServerRuleGroupSet{Match:t_match, Rules:t_rules}
        }
        return tmp
    } else {
        return nil
    }
}

func build_servergroup_msgserverruleset_array(d *schema.ResourceData, r []interface{}) []handler.MsgServerRuleSet {
    if r != nil {
        tmp := make([]handler.MsgServerRuleSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_rulesecvalue *string
            if val, ok := raw_a["rulesecvalue"]; ok {
                t_rulesecvalue = handler.ToStringValue(val, true)
            }
            var t_rulevalue *string
            if val, ok := raw_a["rulevalue"]; ok {
                t_rulevalue = handler.ToStringValue(val, true)
            }
            var t_matchcondition *string
            if val, ok := raw_a["matchcondition"]; ok {
                t_matchcondition = handler.ToStringValue(val, true)
            }
            var t_rulename *string
            if val, ok := raw_a["rulename"]; ok {
                t_rulename = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgServerRuleSet{RuleSecValue:t_rulesecvalue, RuleValue:t_rulevalue, MatchCondition:t_matchcondition, RuleName:t_rulename}
        }
        return tmp
    } else {
        return nil
    }
}

func build_servergroup_msgvirtualservergroupassociationdetails(d *schema.ResourceData, r []interface{}) *handler.MsgVirtualServerGroupAssociationDetails {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_clientassociation []handler.MsgclientAssociationServerGroupContentSet
        if val, ok := tmp["clientassociation"]; ok {
            t_clientassociation = build_servergroup_msgclientassociationservergroupcontentset_array(d, val.(*schema.Set).List())
        }
        var t_virtualservers []handler.MsgIdNameSet
        if val, ok := tmp["virtualservers"]; ok {
            t_virtualservers = build_servergroup_msgidnameset_array(d, val.(*schema.Set).List())
        }
        var t_rules *handler.MsgRuleGroupContent
        if val, ok := tmp["rules"]; ok {
            t_rules = build_servergroup_msgrulegroupcontent(d, val.([]interface{}))
        }
        var t_virtualinstance *string
        if val, ok := tmp["virtualinstance"]; ok {
            t_virtualinstance = handler.ToStringValue(val, true)
        }
        return &handler.MsgVirtualServerGroupAssociationDetails{ClientAssociation:t_clientassociation, VirtualServers:t_virtualservers, Rules:t_rules, VirtualInstance:t_virtualinstance}
    } else {
        return nil
    }
}

func build_servergroup_msgrulegroupcontent(d *schema.ResourceData, r []interface{}) *handler.MsgRuleGroupContent {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_guestcredentialassocid *int
        if val, ok := tmp["guestcredentialassocid"]; ok {
            t_guestcredentialassocid = handler.ToIntValue(val, true)
        }
        var t_guestcredentials *handler.MsgVMGuestCredentials
        if val, ok := tmp["guestcredentials"]; ok {
            t_guestcredentials = build_servergroup_msgvmguestcredentials(d, val.([]interface{}))
        }
        var t_matchrule *string
        if val, ok := tmp["matchrule"]; ok {
            t_matchrule = handler.ToStringValue(val, true)
        }
        var t_description *string
        if val, ok := tmp["description"]; ok {
            t_description = handler.ToStringValue(val, true)
        }
        var t_rules []handler.MsgRuleContentSet
        if val, ok := tmp["rules"]; ok {
            t_rules = build_servergroup_msgrulecontentset_array(d, val.(*schema.Set).List())
        }
        var t_existingcredential *handler.MsgVMExistingCredential
        if val, ok := tmp["existingcredential"]; ok {
            t_existingcredential = build_servergroup_msgvmexistingcredential(d, val.([]interface{}))
        }
        return &handler.MsgRuleGroupContent{GuestCredentialAssocId:t_guestcredentialassocid, GuestCredentials:t_guestcredentials, MatchRule:t_matchrule, Description:t_description, Rules:t_rules, ExistingCredential:t_existingcredential}
    } else {
        return nil
    }
}

func build_servergroup_msgvmexistingcredential(d *schema.ResourceData, r []interface{}) *handler.MsgVMExistingCredential {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_credentialid *int
        if val, ok := tmp["credentialid"]; ok {
            t_credentialid = handler.ToIntValue(val, true)
        }
        var t_credentialname *string
        if val, ok := tmp["credentialname"]; ok {
            t_credentialname = handler.ToStringValue(val, true)
        }
        return &handler.MsgVMExistingCredential{CredentialId:t_credentialid, CredentialName:t_credentialname}
    } else {
        return nil
    }
}

func build_servergroup_msgrulecontentset_array(d *schema.ResourceData, r []interface{}) []handler.MsgRuleContentSet {
    if r != nil {
        tmp := make([]handler.MsgRuleContentSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_condition *string
            if val, ok := raw_a["condition"]; ok {
                t_condition = handler.ToStringValue(val, true)
            }
            var t_name *string
            if val, ok := raw_a["name"]; ok {
                t_name = handler.ToStringValue(val, true)
            }
            var t_type *string
            if val, ok := raw_a["type"]; ok {
                t_type = handler.ToStringValue(val, true)
            }
            var t_value *string
            if val, ok := raw_a["value"]; ok {
                t_value = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgRuleContentSet{Condition:t_condition, Name:t_name, Type:t_type, Value:t_value}
        }
        return tmp
    } else {
        return nil
    }
}

func build_servergroup_msgvmguestcredentials(d *schema.ResourceData, r []interface{}) *handler.MsgVMGuestCredentials {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_password *string
        if val, ok := tmp["password"]; ok {
            t_password = handler.ToStringValue(val, true)
        }
        var t_username *string
        if val, ok := tmp["username"]; ok {
            t_username = handler.ToStringValue(val, true)
        }
        return &handler.MsgVMGuestCredentials{Password:t_password, UserName:t_username}
    } else {
        return nil
    }
}

func build_servergroup_msgclientassociationservergroupcontentset_array(d *schema.ResourceData, r []interface{}) []handler.MsgclientAssociationServerGroupContentSet {
    if r != nil {
        tmp := make([]handler.MsgclientAssociationServerGroupContentSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_associationrule []handler.MsgVirtualMachinecontentSet
            if val, ok := raw_a["associationrule"]; ok {
                t_associationrule = build_servergroup_msgvirtualmachinecontentset_array(d, val.(*schema.Set).List())
            }
            var t_associatedclient *handler.MsgassociatedClientId
            if val, ok := raw_a["associatedclient"]; ok {
                t_associatedclient = build_servergroup_msgassociatedclientid(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgclientAssociationServerGroupContentSet{AssociationRule:t_associationrule, AssociatedClient:t_associatedclient}
        }
        return tmp
    } else {
        return nil
    }
}

func build_servergroup_msgassociatedclientid(d *schema.ResourceData, r []interface{}) *handler.MsgassociatedClientId {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_clientid *int
        if val, ok := tmp["clientid"]; ok {
            t_clientid = handler.ToIntValue(val, true)
        }
        return &handler.MsgassociatedClientId{ClientID:t_clientid}
    } else {
        return nil
    }
}

func build_servergroup_msgvirtualmachinecontentset_array(d *schema.ResourceData, r []interface{}) []handler.MsgVirtualMachinecontentSet {
    if r != nil {
        tmp := make([]handler.MsgVirtualMachinecontentSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_name *string
            if val, ok := raw_a["name"]; ok {
                t_name = handler.ToStringValue(val, true)
            }
            var t_guid *string
            if val, ok := raw_a["guid"]; ok {
                t_guid = handler.ToStringValue(val, true)
            }
            var t_type *string
            if val, ok := raw_a["type"]; ok {
                t_type = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgVirtualMachinecontentSet{Name:t_name, GUID:t_guid, Type:t_type}
        }
        return tmp
    } else {
        return nil
    }
}

func build_servergroup_msgcreateglobalconfiginfo(d *schema.ResourceData, r []interface{}) *handler.MsgCreateGlobalConfigInfo {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_scopefilterquery *string
        if val, ok := tmp["scopefilterquery"]; ok {
            t_scopefilterquery = handler.ToStringValue(val, true)
        }
        var t_companies []handler.MsgGlobalConfigCompanyInfoSet
        if val, ok := tmp["companies"]; ok {
            t_companies = build_servergroup_msgglobalconfigcompanyinfoset_array(d, val.(*schema.Set).List())
        }
        var t_applyonallcommcells *bool
        if val, ok := tmp["applyonallcommcells"]; ok {
            t_applyonallcommcells = handler.ToBooleanValue(val, true)
        }
        var t_commcells []handler.MsgGlobalConfigCommcellInfoSet
        if val, ok := tmp["commcells"]; ok {
            t_commcells = build_servergroup_msgglobalconfigcommcellinfoset_array(d, val.(*schema.Set).List())
        }
        var t_scope *string
        if val, ok := tmp["scope"]; ok {
            t_scope = handler.ToStringValue(val, true)
        }
        var t_name *string
        if val, ok := tmp["name"]; ok {
            t_name = handler.ToStringValue(val, true)
        }
        var t_applyonallcompanies *bool
        if val, ok := tmp["applyonallcompanies"]; ok {
            t_applyonallcompanies = handler.ToBooleanValue(val, true)
        }
        var t_actiononlocalentity *string
        if val, ok := tmp["actiononlocalentity"]; ok {
            t_actiononlocalentity = handler.ToStringValue(val, true)
        }
        return &handler.MsgCreateGlobalConfigInfo{ScopeFilterQuery:t_scopefilterquery, Companies:t_companies, ApplyOnAllCommCells:t_applyonallcommcells, Commcells:t_commcells, Scope:t_scope, Name:t_name, ApplyOnAllCompanies:t_applyonallcompanies, ActionOnLocalEntity:t_actiononlocalentity}
    } else {
        return nil
    }
}

func build_servergroup_msgglobalconfigcommcellinfoset_array(d *schema.ResourceData, r []interface{}) []handler.MsgGlobalConfigCommcellInfoSet {
    if r != nil {
        tmp := make([]handler.MsgGlobalConfigCommcellInfoSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_displayname *string
            if val, ok := raw_a["displayname"]; ok {
                t_displayname = handler.ToStringValue(val, true)
            }
            var t_name *string
            if val, ok := raw_a["name"]; ok {
                t_name = handler.ToStringValue(val, true)
            }
            var t_guid *string
            if val, ok := raw_a["guid"]; ok {
                t_guid = handler.ToStringValue(val, true)
            }
            var t_id *int
            if val, ok := raw_a["id"]; ok {
                t_id = handler.ToIntValue(val, true)
            }
            tmp[a] = handler.MsgGlobalConfigCommcellInfoSet{DisplayName:t_displayname, Name:t_name, Guid:t_guid, Id:t_id}
        }
        return tmp
    } else {
        return nil
    }
}

func build_servergroup_msgglobalconfigcompanyinfoset_array(d *schema.ResourceData, r []interface{}) []handler.MsgGlobalConfigCompanyInfoSet {
    if r != nil {
        tmp := make([]handler.MsgGlobalConfigCompanyInfoSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_name *string
            if val, ok := raw_a["name"]; ok {
                t_name = handler.ToStringValue(val, true)
            }
            var t_guid *string
            if val, ok := raw_a["guid"]; ok {
                t_guid = handler.ToStringValue(val, true)
            }
            var t_id *int
            if val, ok := raw_a["id"]; ok {
                t_id = handler.ToIntValue(val, true)
            }
            tmp[a] = handler.MsgGlobalConfigCompanyInfoSet{Name:t_name, Guid:t_guid, Id:t_id}
        }
        return tmp
    } else {
        return nil
    }
}

func serialize_servergroup_msgidname(d *schema.ResourceData, data *handler.MsgIdName) ([]map[string]interface{}, bool) {
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
