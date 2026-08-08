package commvault

import (
    "fmt"
    "strconv"
    "strings"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRole() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateRole,
        Read:   resourceReadRole,
        Update: resourceUpdateRole,
        Delete: resourceDeleteRole,

        Schema: map[string]*schema.Schema{
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Name of the new role",
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
            "permissionlist": {
                Type:        schema.TypeSet,
                Required:    true,
                Description: "Used to provide the list of permissions associated with the role.",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "permission": {
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
                        "category": {
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
            "description": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Description of the new role",
            },
            "visibletoall": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Determines if the role is visible to everyone. if not provided, it will be set to false by default.",
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: "Used to determine if the role is enabled or disabled. If not provided, role will be enabled by default.",
            },
            "security": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "Used to update the security association for the role",
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
        },
    }
}

func resourceCreateRole(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/Role
    var response_id = strconv.Itoa(0)
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_globalconfiginfo *handler.MsgCreateGlobalConfigInfo
    if val, ok := d.GetOk("globalconfiginfo"); ok {
        t_globalconfiginfo = build_role_msgcreateglobalconfiginfo(d, val.([]interface{}))
    }
    var t_permissionlist []handler.MsgPermissionsSet
    if val, ok := d.GetOk("permissionlist"); ok {
        t_permissionlist = build_role_msgpermissionsset_array(d, val.(*schema.Set).List())
    }
    var t_description *string
    if val, ok := d.GetOk("description"); ok {
        t_description = handler.ToStringValue(val, false)
    }
    var t_visibletoall *bool
    if val, ok := d.GetOk("visibletoall"); ok {
        t_visibletoall = handler.ToBooleanValue(val, false)
    }
    var t_enabled *bool
    if val, ok := d.GetOk("enabled"); ok {
        t_enabled = handler.ToBooleanValue(val, false)
    }
    var req = handler.MsgCreateNewRoleRequest{Name:t_name, GlobalConfigInfo:t_globalconfiginfo, PermissionList:t_permissionlist, Description:t_description, VisibleToAll:t_visibletoall, Enabled:t_enabled}
    resp, err := handler.CvCreateNewRole(req)
    if err != nil {
        return fmt.Errorf("operation [CreateNewRole] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateNewRole] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateRole(d, m)
    }
}

func resourceReadRole(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/Role/{roleId}
    resp, err := handler.CvGetRoleDetails(d.Id())
    if err != nil {
        if strings.Contains(err.Error(), "status: 404") {
            handler.LogEntry("debug", "entity not present, removing from state")
            d.SetId("")
            return nil
        }
        return fmt.Errorf("operation [GetRoleDetails] failed, Error %s", err)
    }
    if rtn, ok := serialize_role_msgsecurityassocset_array(d, resp.Security); ok {
        d.Set("security", rtn)
    } else {
        d.Set("security", make([]map[string]interface{}, 0))
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    if rtn, ok := serialize_role_msgglobalconfiginfo(d, resp.GlobalConfigInfo); ok {
        d.Set("globalconfiginfo", rtn)
    } else {
        d.Set("globalconfiginfo", make([]map[string]interface{}, 0))
    }
    if resp.Description != nil {
        d.Set("description", resp.Description)
    }
    if resp.VisibleToAll != nil {
        d.Set("visibletoall", strconv.FormatBool(*resp.VisibleToAll))
    }
    return nil
}

func resourceUpdateRole(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Role/{roleId}
    var t_permissionoperationtype *string
    var c_permissionoperationtype string = "OVERWRITE"
    t_permissionoperationtype = &c_permissionoperationtype
    var t_security []handler.MsgUpdateSecurityAssocSet
    if d.HasChange("security") {
        val := d.Get("security")
        t_security = build_role_msgupdatesecurityassocset_array(d, val.(*schema.Set).List())
    }
    var t_newname *string
    if d.HasChange("name") {
        val := d.Get("name")
        t_newname = handler.ToStringValue(val, false)
    }
    var t_permissionlist []handler.MsgPermissionsSet
    if d.HasChange("permissionlist") {
        val := d.Get("permissionlist")
        t_permissionlist = build_role_msgpermissionsset_array(d, val.(*schema.Set).List())
    }
    var t_description *string
    if d.HasChange("description") {
        val := d.Get("description")
        t_description = handler.ToStringValue(val, false)
    }
    var t_visibletoall *bool
    if d.HasChange("visibletoall") {
        val := d.Get("visibletoall")
        t_visibletoall = handler.ToBooleanValue(val, false)
    }
    var t_enabled *bool
    if d.HasChange("enabled") {
        val := d.Get("enabled")
        t_enabled = handler.ToBooleanValue(val, false)
    }
    var req = handler.MsgModifyRoleRequest{PermissionOperationType:t_permissionoperationtype, Security:t_security, NewName:t_newname, PermissionList:t_permissionlist, Description:t_description, VisibleToAll:t_visibletoall, Enabled:t_enabled}
    _, err := handler.CvModifyRole(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [ModifyRole] failed, Error %s", err)
    }
    return resourceReadRole(d, m)
}

func resourceCreateUpdateRole(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/Role/{roleId}
    var execUpdate bool = false
    var t_permissionoperationtype *string
    var c_permissionoperationtype string = "OVERWRITE"
    t_permissionoperationtype = &c_permissionoperationtype
    var t_security []handler.MsgUpdateSecurityAssocSet
    if val, ok := d.GetOk("security"); ok {
        t_security = build_role_msgupdatesecurityassocset_array(d, val.(*schema.Set).List())
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgModifyRoleRequest{PermissionOperationType:t_permissionoperationtype, Security:t_security}
        _, err := handler.CvModifyRole(req, d.Id())
        if err != nil {
            return fmt.Errorf("operation [ModifyRole] failed, Error %s", err)
        }
    }
    return resourceReadRole(d, m)
}

func resourceDeleteRole(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/Role/{roleId}
    _, err := handler.CvDeleteRoles(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteRoles] failed, Error %s", err)
    }
    return nil
}

func build_role_msgupdatesecurityassocset_array(d *schema.ResourceData, r []interface{}) []handler.MsgUpdateSecurityAssocSet {
    if r != nil {
        tmp := make([]handler.MsgUpdateSecurityAssocSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_role *handler.MsgIdName
            if val, ok := raw_a["role"]; ok {
                t_role = build_role_msgidname(d, val.([]interface{}))
            }
            var t_user *handler.MsgIdName
            if val, ok := raw_a["user"]; ok {
                t_user = build_role_msgidname(d, val.([]interface{}))
            }
            var t_usergroup *handler.MsgIdName
            if val, ok := raw_a["usergroup"]; ok {
                t_usergroup = build_role_msgidname(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgUpdateSecurityAssocSet{Role:t_role, User:t_user, UserGroup:t_usergroup}
        }
        return tmp
    } else {
        return nil
    }
}

func build_role_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_id *int
        if val, ok := tmp["id"]; ok {
            t_id = handler.ToIntValue(val, true)
        }
        return &handler.MsgIdName{Id:t_id}
    } else {
        return nil
    }
}

func build_role_msgpermissionsset_array(d *schema.ResourceData, r []interface{}) []handler.MsgPermissionsSet {
    if r != nil {
        tmp := make([]handler.MsgPermissionsSet, len(r))
        for a, iter_a := range r {
            raw_a := iter_a.(map[string]interface{})
            var t_permission *handler.MsgIdName
            if val, ok := raw_a["permission"]; ok {
                t_permission = build_role_msgidname(d, val.([]interface{}))
            }
            var t_category *handler.MsgIdName
            if val, ok := raw_a["category"]; ok {
                t_category = build_role_msgidname(d, val.([]interface{}))
            }
            tmp[a] = handler.MsgPermissionsSet{Permission:t_permission, Category:t_category}
        }
        return tmp
    } else {
        return nil
    }
}

func build_role_msgcreateglobalconfiginfo(d *schema.ResourceData, r []interface{}) *handler.MsgCreateGlobalConfigInfo {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_scopefilterquery *string
        if val, ok := tmp["scopefilterquery"]; ok {
            t_scopefilterquery = handler.ToStringValue(val, true)
        }
        var t_companies []handler.MsgGlobalConfigCompanyInfoSet
        if val, ok := tmp["companies"]; ok {
            t_companies = build_role_msgglobalconfigcompanyinfoset_array(d, val.(*schema.Set).List())
        }
        var t_applyonallcommcells *bool
        if val, ok := tmp["applyonallcommcells"]; ok {
            t_applyonallcommcells = handler.ToBooleanValue(val, true)
        }
        var t_commcells []handler.MsgGlobalConfigCommcellInfoSet
        if val, ok := tmp["commcells"]; ok {
            t_commcells = build_role_msgglobalconfigcommcellinfoset_array(d, val.(*schema.Set).List())
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

func build_role_msgglobalconfigcommcellinfoset_array(d *schema.ResourceData, r []interface{}) []handler.MsgGlobalConfigCommcellInfoSet {
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

func build_role_msgglobalconfigcompanyinfoset_array(d *schema.ResourceData, r []interface{}) []handler.MsgGlobalConfigCompanyInfoSet {
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

func serialize_role_msgglobalconfiginfo(d *schema.ResourceData, data *handler.MsgGlobalConfigInfo) ([]map[string]interface{}, bool) {
    //MsgCreateGlobalConfigInfo
    //MsgGlobalConfigInfo
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
    if data.ScopeFilterQuery != nil {
        val[0]["scopefilterquery"] = data.ScopeFilterQuery
        added = true
    }
    if rtn, ok := serialize_role_msgglobalconfigcompanyinfoset_array(d, data.Companies); ok {
        val[0]["companies"] = rtn
        added = true
    }
    if data.ApplyOnAllCommCells != nil {
        val[0]["applyonallcommcells"] = strconv.FormatBool(*data.ApplyOnAllCommCells)
        added = true
    }
    if rtn, ok := serialize_role_msgglobalconfigcommcellinfoset_array(d, data.Commcells); ok {
        val[0]["commcells"] = rtn
        added = true
    }
    if data.Scope != nil {
        val[0]["scope"] = data.Scope
        added = true
    }
    if data.Name != nil {
        val[0]["name"] = data.Name
        added = true
    }
    if data.ApplyOnAllCompanies != nil {
        val[0]["applyonallcompanies"] = strconv.FormatBool(*data.ApplyOnAllCompanies)
        added = true
    }
    if added {
        return val, true
    } else {
        return nil, false
    }
}

func serialize_role_msgglobalconfigcommcellinfoset_array(d *schema.ResourceData, data []handler.MsgGlobalConfigCommcellInfoSet) ([]map[string]interface{}, bool) {
    //MsgCreateGlobalConfigInfo -> MsgGlobalConfigCommcellInfoSet
    //MsgGlobalConfigInfo -> MsgGlobalConfigCommcellInfoSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].DisplayName != nil {
            tmp["displayname"] = data[i].DisplayName
            added = true
        }
        if data[i].Name != nil {
            tmp["name"] = data[i].Name
            added = true
        }
        if data[i].Guid != nil {
            tmp["guid"] = data[i].Guid
            added = true
        }
        if data[i].Id != nil {
            tmp["id"] = data[i].Id
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_role_msgglobalconfigcompanyinfoset_array(d *schema.ResourceData, data []handler.MsgGlobalConfigCompanyInfoSet) ([]map[string]interface{}, bool) {
    //MsgCreateGlobalConfigInfo -> MsgGlobalConfigCompanyInfoSet
    //MsgGlobalConfigInfo -> MsgGlobalConfigCompanyInfoSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if data[i].Name != nil {
            tmp["name"] = data[i].Name
            added = true
        }
        if data[i].Guid != nil {
            tmp["guid"] = data[i].Guid
            added = true
        }
        if data[i].Id != nil {
            tmp["id"] = data[i].Id
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_role_msgsecurityassocset_array(d *schema.ResourceData, data []handler.MsgSecurityAssocSet) ([]map[string]interface{}, bool) {
    //MsgUpdateSecurityAssocSet
    //MsgSecurityAssocSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
        if rtn, ok := serialize_role_msgidname(d, data[i].Role); ok {
            tmp["role"] = rtn
            added = true
        }
        if rtn, ok := serialize_role_msgidname(d, data[i].User); ok {
            tmp["user"] = rtn
            added = true
        }
        if rtn, ok := serialize_role_msgidname(d, data[i].UserGroup); ok {
            tmp["usergroup"] = rtn
            added = true
        }
        if added {
            val = append(val, tmp)
        }
    }
    return val, true
}

func serialize_role_msgidname(d *schema.ResourceData, data *handler.MsgIdName) ([]map[string]interface{}, bool) {
    //MsgUpdateSecurityAssocSet -> MsgIdName
    //MsgSecurityAssocSet -> MsgIdName
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 1)
    val[0] = make(map[string]interface{})
    added := false
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
