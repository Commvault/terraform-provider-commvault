package commvault

import (
    "fmt"
    "strconv"
    "strings"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCloudConnection() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateCloudConnection,
        Read:   resourceReadCloudConnection,
        Update: resourceUpdateCloudConnection,
        Delete: resourceDeleteCloudConnection,

        Schema: map[string]*schema.Schema{
            "credentials": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Credentials for cloud connections",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "credentialtype": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: "Type of credential to be used to access cloud connection. [AWS_STS_ASSUME_ROLE, AWS_IAM_ROLE]",
                        },
                        "credentialid": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Computed:    true,
                            Description: "ID of the saved credentials.",
                        },
                    },
                },
            },
            "cloudtype": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Cloud type for the connection. [AMAZON_WEB_SERVICES, MICROSOFT_AZURE, none, aws, azure, googleCloud]",
            },
            "cloudspecificconfiguration": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "Configuration settings specific to the connections cloud type.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "googlecloud": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "Google Cloud specific cloud connection configuration.",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "projects": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "List of Google Cloud projects.",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "cloud account name",
                                                },
                                                "id": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "cloud account id",
                                                },
                                                "uuid": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "The globally unique identifier for the account",
                                                },
                                                "email": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "email for cloud account",
                                                },
                                            },
                                        },
                                    },
                                    "discoverallprojects": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Automatically discover future projects associated with the service account.",
                                    },
                                    "serviceaccount": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Service Account email address",
                                    },
                                },
                            },
                        },
                        "aws": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "AWS specific cloud connection configuration.",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "regions": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "List of AWS regions to connect to.",
                                    },
                                    "organizationconfiguration": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Configuration for AWS Organization cloud connection.",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "ownerdetectionconfiguration": {
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Owner discovery configuration.",
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "identitycenterregion": {
                                                                Type:        schema.TypeString,
                                                                Required:    true,
                                                                Description: "Region code of the region in which IAM Identity Center service is enabled in the AWS organization.",
                                                            },
                                                            "ownerpermissionsets": {
                                                                Type:        schema.TypeSet,
                                                                Required:    true,
                                                                Description: "List of IAM Identity center permission sets that identify account owners.",
                                                                Elem: &schema.Schema{
                                                                    Type:    schema.TypeString,
                                                                },
                                                            },
                                                        },
                                                    },
                                                },
                                                "enableownerdetection": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Enable owner detection for AWS accounts in the organization.",
                                                },
                                                "content": {
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "Content for AWS Organization cloud connection.",
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "discoverallaccounts": {
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "Discover all accounts in the AWS organization",
                                                            },
                                                            "accounts": {
                                                                Type:        schema.TypeSet,
                                                                Optional:    true,
                                                                Computed:    true,
                                                                Description: "List of AWS accounts",
                                                                Elem: &schema.Resource{
                                                                    Schema: map[string]*schema.Schema{
                                                                        "name": {
                                                                            Type:        schema.TypeString,
                                                                            Optional:    true,
                                                                            Computed:    true,
                                                                            Description: "cloud account name",
                                                                        },
                                                                        "id": {
                                                                            Type:        schema.TypeString,
                                                                            Optional:    true,
                                                                            Computed:    true,
                                                                            Description: "cloud account id",
                                                                        },
                                                                        "uuid": {
                                                                            Type:        schema.TypeString,
                                                                            Optional:    true,
                                                                            Computed:    true,
                                                                            Description: "The globally unique identifier for the account",
                                                                        },
                                                                        "email": {
                                                                            Type:        schema.TypeString,
                                                                            Optional:    true,
                                                                            Computed:    true,
                                                                            Description: "email for cloud account",
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
                                    "iamroleaccountid": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "ID of the AWS account which contains the IAM role to assume for authentication.",
                                    },
                                },
                            },
                        },
                        "azure": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Computed:    true,
                            Description: "Azure specific cloud connection configuration.",
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "environment": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Azure environment",
                                    },
                                    "subscriptions": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "List of Azure Subscriptions",
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "name": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "cloud account name",
                                                },
                                                "id": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "cloud account id",
                                                },
                                                "uuid": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "The globally unique identifier for the account",
                                                },
                                                "email": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Computed:    true,
                                                    Description: "email for cloud account",
                                                },
                                            },
                                        },
                                    },
                                    "tenantname": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Azure Tenant Name",
                                    },
                                    "tenantid": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Azure Tenant Id",
                                    },
                                    "discoverallsubscription": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Automatically discover future subscription in Azure tenant.",
                                    },
                                    "iscustomconfig": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Computed:    true,
                                        Description: "Flag to indicate if custom configuration is used.",
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "startdiscoveryjob": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Indicates whether the user wants to start the discovery job with the modification",
            },
            "name": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Name of the cloud connection.",
            },
            "accessnodes": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "List of servers and server groups to use to access the cloud connection.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "displayname": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "",
                        },
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
                        "accessnodetype": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: "Type of access node item [SERVER, SERVER_GROUP]",
                        },
                    },
                },
            },
            "connectiontype": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Type of the cloud connection. [CloudAccountLevel, OrganizationLevel]",
            },
            "enablediscovery": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Indicates whether the user wants to discover resources with the new cloud connection",
            },
            "displayname": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Display name of the cloud connection.",
            },
        },
    }
}

func resourceCreateCloudConnection(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/Cloud/CloudConnection
    var response_id = strconv.Itoa(0)
    var t_credentials *handler.MsgCloudConnectionCredentialsInfo
    if val, ok := d.GetOk("credentials"); ok {
        t_credentials = build_cloudconnection_msgcloudconnectioncredentialsinfo(d, val.([]interface{}))
    }
    var t_cloudtype *string
    if val, ok := d.GetOk("cloudtype"); ok {
        t_cloudtype = handler.ToStringValue(val, false)
    }
    var t_cloudspecificconfiguration *handler.MsgCloudSpecificCloudConnectionConfiguration
    if val, ok := d.GetOk("cloudspecificconfiguration"); ok {
        t_cloudspecificconfiguration = build_cloudconnection_msgcloudspecificcloudconnectionconfiguration(d, val.([]interface{}))
    }
    var t_startdiscoveryjob *bool
    if val, ok := d.GetOk("startdiscoveryjob"); ok {
        t_startdiscoveryjob = handler.ToBooleanValue(val, false)
    }
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_accessnodes []handler.MsgAccessNodeItemSet
    if val, ok := d.GetOk("accessnodes"); ok {
        t_accessnodes = build_cloudconnection_msgaccessnodeitemset_array(d, val.(*schema.Set).List())
    }
    var t_connectiontype *string
    if val, ok := d.GetOk("connectiontype"); ok {
        t_connectiontype = handler.ToStringValue(val, false)
    }
    var t_enablediscovery *bool
    if val, ok := d.GetOk("enablediscovery"); ok {
        t_enablediscovery = handler.ToBooleanValue(val, false)
    }
    var req = handler.MsgCreateCloudConnectionRequest{Credentials:t_credentials, CloudType:t_cloudtype, CloudSpecificConfiguration:t_cloudspecificconfiguration, StartDiscoveryJob:t_startdiscoveryjob, Name:t_name, AccessNodes:t_accessnodes, ConnectionType:t_connectiontype, EnableDiscovery:t_enablediscovery}
    resp, err := handler.CvCreateCloudConnection(req)
    if err != nil {
        return fmt.Errorf("operation [CreateCloudConnection] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateCloudConnection] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateCloudConnection(d, m)
    }
}

func resourceReadCloudConnection(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/Cloud/CloudConnection/{cloudConnectionId}
    _, err := handler.CvGetCloudConnectionDetails(d.Id())
    if err != nil {
        if strings.Contains(err.Error(), "status: 404") {
            handler.LogEntry("debug", "entity not present, removing from state")
            d.SetId("")
            return nil
        }
        return fmt.Errorf("operation [GetCloudConnectionDetails] failed, Error %s", err)
    }
    return nil
}

func resourceUpdateCloudConnection(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Cloud/CloudConnection/{cloudConnectionId}
    var t_credentials *handler.MsgCloudConnectionCredentialsInfo
    if d.HasChange("credentials") {
        val := d.Get("credentials")
        t_credentials = build_cloudconnection_msgcloudconnectioncredentialsinfo(d, val.([]interface{}))
    }
    var t_displayname *string
    if d.HasChange("displayname") {
        val := d.Get("displayname")
        t_displayname = handler.ToStringValue(val, false)
    }
    var t_cloudspecificconfiguration *handler.MsgCloudSpecificCloudConnectionConfiguration
    if d.HasChange("cloudspecificconfiguration") {
        val := d.Get("cloudspecificconfiguration")
        t_cloudspecificconfiguration = build_cloudconnection_msgcloudspecificcloudconnectionconfiguration(d, val.([]interface{}))
    }
    var t_accessnodes []handler.MsgAccessNodeItemSet
    if d.HasChange("accessnodes") {
        val := d.Get("accessnodes")
        t_accessnodes = build_cloudconnection_msgaccessnodeitemset_array(d, val.(*schema.Set).List())
    }
    var req = handler.MsgUpdateCloudConnectionRequest{Credentials:t_credentials, DisplayName:t_displayname, CloudSpecificConfiguration:t_cloudspecificconfiguration, AccessNodes:t_accessnodes}
    _, err := handler.CvUpdateCloudConnection(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [UpdateCloudConnection] failed, Error %s", err)
    }
    return resourceReadCloudConnection(d, m)
}

func resourceCreateUpdateCloudConnection(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Cloud/CloudConnection/{cloudConnectionId}
    var execUpdate bool = false
    var t_displayname *string
    if val, ok := d.GetOk("displayname"); ok {
        t_displayname = handler.ToStringValue(val, false)
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgUpdateCloudConnectionRequest{DisplayName:t_displayname}
        _, err := handler.CvUpdateCloudConnection(req, d.Id())
        if err != nil {
            return fmt.Errorf("operation [UpdateCloudConnection] failed, Error %s", err)
        }
    }
    return resourceReadCloudConnection(d, m)
}

func resourceDeleteCloudConnection(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/Cloud/CloudConnection/{cloudConnectionId}
    _, err := handler.CvDeleteCloudConnection(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteCloudConnection] failed, Error %s", err)
    }
    return nil
}

func build_cloudconnection_msgaccessnodeitemset_array(d *schema.ResourceData, r []interface{}) []handler.MsgAccessNodeItemSet {
    if r != nil {
        tmp := make([]handler.MsgAccessNodeItemSet, len(r))
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
            var t_id *int
            if val, ok := raw_a["id"]; ok {
                t_id = handler.ToIntValue(val, true)
            }
            var t_accessnodetype *string
            if val, ok := raw_a["accessnodetype"]; ok {
                t_accessnodetype = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgAccessNodeItemSet{DisplayName:t_displayname, Name:t_name, Id:t_id, AccessNodeType:t_accessnodetype}
        }
        return tmp
    } else {
        return nil
    }
}

func build_cloudconnection_msgcloudspecificcloudconnectionconfiguration(d *schema.ResourceData, r []interface{}) *handler.MsgCloudSpecificCloudConnectionConfiguration {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_googlecloud *handler.MsgGCPCloudConnectionConfiguration
        if val, ok := tmp["googlecloud"]; ok {
            t_googlecloud = build_cloudconnection_msggcpcloudconnectionconfiguration(d, val.([]interface{}))
        }
        var t_aws *handler.MsgAWSCloudConnectionConfiguration
        if val, ok := tmp["aws"]; ok {
            t_aws = build_cloudconnection_msgawscloudconnectionconfiguration(d, val.([]interface{}))
        }
        var t_azure *handler.MsgAzureCloudConnectionConfiguration
        if val, ok := tmp["azure"]; ok {
            t_azure = build_cloudconnection_msgazurecloudconnectionconfiguration(d, val.([]interface{}))
        }
        return &handler.MsgCloudSpecificCloudConnectionConfiguration{GoogleCloud:t_googlecloud, Aws:t_aws, Azure:t_azure}
    } else {
        return nil
    }
}

func build_cloudconnection_msgazurecloudconnectionconfiguration(d *schema.ResourceData, r []interface{}) *handler.MsgAzureCloudConnectionConfiguration {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_environment *string
        if val, ok := tmp["environment"]; ok {
            t_environment = handler.ToStringValue(val, true)
        }
        var t_subscriptions []handler.MsgCloudAccountSet
        if val, ok := tmp["subscriptions"]; ok {
            t_subscriptions = build_cloudconnection_msgcloudaccountset_array(d, val.(*schema.Set).List())
        }
        var t_tenantname *string
        if val, ok := tmp["tenantname"]; ok {
            t_tenantname = handler.ToStringValue(val, true)
        }
        var t_tenantid *string
        if val, ok := tmp["tenantid"]; ok {
            t_tenantid = handler.ToStringValue(val, true)
        }
        var t_discoverallsubscription *bool
        if val, ok := tmp["discoverallsubscription"]; ok {
            t_discoverallsubscription = handler.ToBooleanValue(val, true)
        }
        var t_iscustomconfig *bool
        if val, ok := tmp["iscustomconfig"]; ok {
            t_iscustomconfig = handler.ToBooleanValue(val, true)
        }
        return &handler.MsgAzureCloudConnectionConfiguration{Environment:t_environment, Subscriptions:t_subscriptions, TenantName:t_tenantname, TenantId:t_tenantid, DiscoverAllSubscription:t_discoverallsubscription, IsCustomConfig:t_iscustomconfig}
    } else {
        return nil
    }
}

func build_cloudconnection_msgcloudaccountset_array(d *schema.ResourceData, r []interface{}) []handler.MsgCloudAccountSet {
    if r != nil {
        tmp := make([]handler.MsgCloudAccountSet, len(r))
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
            var t_uuid *string
            if val, ok := raw_a["uuid"]; ok {
                t_uuid = handler.ToStringValue(val, true)
            }
            var t_email *string
            if val, ok := raw_a["email"]; ok {
                t_email = handler.ToStringValue(val, true)
            }
            tmp[a] = handler.MsgCloudAccountSet{Name:t_name, Id:t_id, Uuid:t_uuid, Email:t_email}
        }
        return tmp
    } else {
        return nil
    }
}

func build_cloudconnection_msgawscloudconnectionconfiguration(d *schema.ResourceData, r []interface{}) *handler.MsgAWSCloudConnectionConfiguration {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_regions *string
        if val, ok := tmp["regions"]; ok {
            t_regions = handler.ToStringValue(val, true)
        }
        var t_organizationconfiguration *handler.MsgAWSCloudConnectionOrganizationConfiguration
        if val, ok := tmp["organizationconfiguration"]; ok {
            t_organizationconfiguration = build_cloudconnection_msgawscloudconnectionorganizationconfiguration(d, val.([]interface{}))
        }
        var t_iamroleaccountid *string
        if val, ok := tmp["iamroleaccountid"]; ok {
            t_iamroleaccountid = handler.ToStringValue(val, true)
        }
        return &handler.MsgAWSCloudConnectionConfiguration{Regions:t_regions, OrganizationConfiguration:t_organizationconfiguration, IamRoleAccountId:t_iamroleaccountid}
    } else {
        return nil
    }
}

func build_cloudconnection_msgawscloudconnectionorganizationconfiguration(d *schema.ResourceData, r []interface{}) *handler.MsgAWSCloudConnectionOrganizationConfiguration {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_ownerdetectionconfiguration *handler.MsgAWSCloudConnectionOwnerDiscoveryConfiguration
        if val, ok := tmp["ownerdetectionconfiguration"]; ok {
            t_ownerdetectionconfiguration = build_cloudconnection_msgawscloudconnectionownerdiscoveryconfiguration(d, val.([]interface{}))
        }
        var t_enableownerdetection *bool
        if val, ok := tmp["enableownerdetection"]; ok {
            t_enableownerdetection = handler.ToBooleanValue(val, true)
        }
        var t_content *handler.MsgAWSCloudConnectionOrganizationContent
        if val, ok := tmp["content"]; ok {
            t_content = build_cloudconnection_msgawscloudconnectionorganizationcontent(d, val.([]interface{}))
        }
        return &handler.MsgAWSCloudConnectionOrganizationConfiguration{OwnerDetectionConfiguration:t_ownerdetectionconfiguration, EnableOwnerDetection:t_enableownerdetection, Content:t_content}
    } else {
        return nil
    }
}

func build_cloudconnection_msgawscloudconnectionorganizationcontent(d *schema.ResourceData, r []interface{}) *handler.MsgAWSCloudConnectionOrganizationContent {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_discoverallaccounts *bool
        if val, ok := tmp["discoverallaccounts"]; ok {
            t_discoverallaccounts = handler.ToBooleanValue(val, true)
        }
        var t_accounts []handler.MsgCloudAccountSet
        if val, ok := tmp["accounts"]; ok {
            t_accounts = build_cloudconnection_msgcloudaccountset_array(d, val.(*schema.Set).List())
        }
        return &handler.MsgAWSCloudConnectionOrganizationContent{DiscoverAllAccounts:t_discoverallaccounts, Accounts:t_accounts}
    } else {
        return nil
    }
}

func build_cloudconnection_msgawscloudconnectionownerdiscoveryconfiguration(d *schema.ResourceData, r []interface{}) *handler.MsgAWSCloudConnectionOwnerDiscoveryConfiguration {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_identitycenterregion *string
        if val, ok := tmp["identitycenterregion"]; ok {
            t_identitycenterregion = handler.ToStringValue(val, true)
        }
        var t_ownerpermissionsets []string
        if val, ok := tmp["ownerpermissionsets"]; ok {
            t_ownerpermissionsets = handler.ToStringArray(val.(*schema.Set).List())
        }
        return &handler.MsgAWSCloudConnectionOwnerDiscoveryConfiguration{IdentityCenterRegion:t_identitycenterregion, OwnerPermissionSets:t_ownerpermissionsets}
    } else {
        return nil
    }
}

func build_cloudconnection_msggcpcloudconnectionconfiguration(d *schema.ResourceData, r []interface{}) *handler.MsgGCPCloudConnectionConfiguration {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_projects []handler.MsgCloudAccountSet
        if val, ok := tmp["projects"]; ok {
            t_projects = build_cloudconnection_msgcloudaccountset_array(d, val.(*schema.Set).List())
        }
        var t_discoverallprojects *bool
        if val, ok := tmp["discoverallprojects"]; ok {
            t_discoverallprojects = handler.ToBooleanValue(val, true)
        }
        var t_serviceaccount *string
        if val, ok := tmp["serviceaccount"]; ok {
            t_serviceaccount = handler.ToStringValue(val, true)
        }
        return &handler.MsgGCPCloudConnectionConfiguration{Projects:t_projects, DiscoverAllProjects:t_discoverallprojects, ServiceAccount:t_serviceaccount}
    } else {
        return nil
    }
}

func build_cloudconnection_msgcloudconnectioncredentialsinfo(d *schema.ResourceData, r []interface{}) *handler.MsgCloudConnectionCredentialsInfo {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_credentialtype *string
        if val, ok := tmp["credentialtype"]; ok {
            t_credentialtype = handler.ToStringValue(val, true)
        }
        var t_credentialid *int
        if val, ok := tmp["credentialid"]; ok {
            t_credentialid = handler.ToIntValue(val, true)
        }
        return &handler.MsgCloudConnectionCredentialsInfo{CredentialType:t_credentialtype, CredentialId:t_credentialid}
    } else {
        return nil
    }
}
