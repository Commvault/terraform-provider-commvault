package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceUserGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateUserGroup,
        Read:   resourceReadUserGroup,
        Update: resourceUpdateUserGroup,
        Delete: resourceDeleteUserGroup,

        Schema: map[string]*schema.Schema{
            "name": &schema.Schema{
                Type:        schema.TypeString,
                Required:    true,
                Description: "To create an active directory usergroup, the domain name should be mentioned along with the usergroup name (domainName\\usergroupName) and localUserGroup value must be given.",
            },
            "description": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "",
            },
            "enforcefsquota": &schema.Schema{
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: "Used to determine if a backup data limit will be set for the user group being created",
            },
            "quotalimitingb": &schema.Schema{
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: "if enforceFSQuota is set to true, the quota limit can be set in GBs",
            },
            "enabletwofactorauthentication": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Allows two-factor authentication to be enabled for the specific types of usergroups. it can be turned on or off based on user preferences. There will be usergroups that will not have this option.",
            },
            "laptopadmins": &schema.Schema{
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: "When set to true, users in this group cannot activate or be set as server owner",
            },
            "allowmultiplecompanymembers": &schema.Schema{
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: "This property can be used to allow addition of users/groups from child companies. Only applicable for commcell and reseller company group.",
            },
            "enabled": &schema.Schema{
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: "allows the enabling/disabling of the user group.",
            },
            "users": &schema.Schema{
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
            "restrictconsoletypes": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "consoletype": &schema.Schema{
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "",
                            Elem: &schema.Schema{
                                Type:    schema.TypeString,
                            },
                        },
                    },
                },
            },
            "azureguid": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Azure Object ID used to link this user group to Azure AD group and manage group membership of the user during SAML login",
            },
            "planoperationtype": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "determines if an existing user has to be added to the user group or removed from the user group",
            },
            "associatedexternalgroups": &schema.Schema{
                Type:        schema.TypeSet,
                Optional:    true,
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
        },
    }
}

func resourceCreateUserGroup(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/UserGroup
    var response_id = strconv.Itoa(0)
    var t_name *string
    if v, ok := d.GetOk("name"); ok {
        val := v.(string)
        t_name = new(string)
        t_name = &val
    }
    var t_description *string
    if v, ok := d.GetOk("description"); ok {
        val := v.(string)
        t_description = new(string)
        t_description = &val
    }
    var t_enforcefsquota *bool
    if v, ok := d.GetOkExists("enforcefsquota"); ok {
        val := v.(bool)
        t_enforcefsquota = new(bool)
        t_enforcefsquota = &val
    }
    var t_quotalimitingb *int
    if v, ok := d.GetOk("quotalimitingb"); ok {
        val := v.(int)
        t_quotalimitingb = new(int)
        t_quotalimitingb = &val
    }
    var req = handler.MsgCreateUserGroupRequest{Name:t_name, Description:t_description, EnforceFSQuota:t_enforcefsquota, QuotaLimitInGB:t_quotalimitingb}
    resp, err := handler.CvCreateUserGroup(req)
    if err != nil {
        return fmt.Errorf("Operation [CreateUserGroup] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("Operation [CreateUserGroup] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateUserGroup(d, m)
    }
}

func resourceReadUserGroup(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/UserGroup/{userGroupId}
    resp, err := handler.CvGetUserGroupDetails(d.Id())
    if err != nil {
        return fmt.Errorf("Operation [GetUserGroupDetails] failed, Error %s", err)
    }
    if resp.RestrictedConsoleTypes != nil {
        d.Set("restrictconsoletypes", serialize_usergroup_msgrestrictedconsoletypesset_array(resp.RestrictedConsoleTypes))
    } else {
        d.Set("restrictconsoletypes", make([]map[string]interface{}, 0))
    }
    if resp.EnableTwoFactorAuthentication != nil {
        d.Set("enabletwofactorauthentication", resp.EnableTwoFactorAuthentication)
    }
    if resp.LaptopAdmins != nil {
        d.Set("laptopadmins", resp.LaptopAdmins)
    }
    if resp.AllowMultipleCompanyMembers != nil {
        d.Set("allowmultiplecompanymembers", resp.AllowMultipleCompanyMembers)
    }
    if resp.Description != nil {
        d.Set("description", resp.Description)
    }
    if resp.EnforceFSQuota != nil {
        d.Set("enforcefsquota", resp.EnforceFSQuota)
    }
    if resp.QuotaLimitInGB != nil {
        d.Set("quotalimitingb", resp.QuotaLimitInGB)
    }
    if resp.Enabled != nil {
        d.Set("enabled", resp.Enabled)
    }
    if resp.Users != nil {
        d.Set("users", serialize_usergroup_msgidnameset_array(resp.Users))
    } else {
        d.Set("users", make([]map[string]interface{}, 0))
    }
    if resp.AzureGUID != nil {
        d.Set("azureguid", resp.AzureGUID)
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    if resp.AssociatedExternalGroups != nil {
        d.Set("associatedexternalgroups", serialize_usergroup_msgidnameset_array(resp.AssociatedExternalGroups))
    } else {
        d.Set("associatedexternalgroups", make([]map[string]interface{}, 0))
    }
    return nil
}

func resourceUpdateUserGroup(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/UserGroup/{userGroupId}
    var t_enabletwofactorauthentication *string
    if d.HasChange("enabletwofactorauthentication") {
        val := d.Get("enabletwofactorauthentication").(string)
        t_enabletwofactorauthentication = new(string)
        t_enabletwofactorauthentication = &val
    }
    var t_laptopadmins *bool
    if d.HasChange("laptopadmins") {
        val := d.Get("laptopadmins").(bool)
        t_laptopadmins = new(bool)
        t_laptopadmins = &val
    }
    var t_allowmultiplecompanymembers *bool
    if d.HasChange("allowmultiplecompanymembers") {
        val := d.Get("allowmultiplecompanymembers").(bool)
        t_allowmultiplecompanymembers = new(bool)
        t_allowmultiplecompanymembers = &val
    }
    var t_enforcefsquota *bool
    if d.HasChange("enforcefsquota") {
        val := d.Get("enforcefsquota").(bool)
        t_enforcefsquota = new(bool)
        t_enforcefsquota = &val
    }
    var t_quotalimitingb *int
    if d.HasChange("quotalimitingb") {
        val := d.Get("quotalimitingb").(int)
        t_quotalimitingb = new(int)
        t_quotalimitingb = &val
    }
    var t_externalusergroupsoperationtype *string
    if d.HasChange("associatedexternalgroups") {
        var c_externalusergroupsoperationtype string = "OVERWRITE"
        t_externalusergroupsoperationtype = new(string)
        t_externalusergroupsoperationtype = &c_externalusergroupsoperationtype
    }
    var t_newdescription *string
    if d.HasChange("description") {
        val := d.Get("description").(string)
        t_newdescription = new(string)
        t_newdescription = &val
    }
    var t_enabled *bool
    if d.HasChange("enabled") {
        val := d.Get("enabled").(bool)
        t_enabled = new(bool)
        t_enabled = &val
    }
    var t_users []handler.MsgIdNameSet
    if d.HasChange("users") {
        val := d.Get("users").(*schema.Set)
        t_users = build_usergroup_msgidnameset_array(d, val.List())
    }
    var t_useroperationtype *string
    if d.HasChange("users") {
        var c_useroperationtype string = "OVERWRITE"
        t_useroperationtype = new(string)
        t_useroperationtype = &c_useroperationtype
    }
    var t_restrictconsoletypes *handler.MsgRestrictConsoleTypes
    if d.HasChange("restrictconsoletypes") {
        val := d.Get("restrictconsoletypes").([]interface{})
        t_restrictconsoletypes = build_usergroup_msgrestrictconsoletypes(d, val)
    }
    var t_newname *string
    if d.HasChange("name") {
        val := d.Get("name").(string)
        t_newname = new(string)
        t_newname = &val
    }
    var t_azureguid *string
    if d.HasChange("azureguid") {
        val := d.Get("azureguid").(string)
        t_azureguid = new(string)
        t_azureguid = &val
    }
    var t_consoletypeoperationtype *string
    if d.HasChange("restrictconsoletypes") {
        var c_consoletypeoperationtype string = "OVERWRITE"
        t_consoletypeoperationtype = new(string)
        t_consoletypeoperationtype = &c_consoletypeoperationtype
    }
    var t_planoperationtype *string
    if d.HasChange("planoperationtype") {
        val := d.Get("planoperationtype").(string)
        t_planoperationtype = new(string)
        t_planoperationtype = &val
    }
    var t_associatedexternalgroups []handler.MsgIdNameSet
    if d.HasChange("associatedexternalgroups") {
        val := d.Get("associatedexternalgroups").(*schema.Set)
        t_associatedexternalgroups = build_usergroup_msgidnameset_array(d, val.List())
    }
    var req = handler.MsgModifyUserGroupRequest{EnableTwoFactorAuthentication:t_enabletwofactorauthentication, LaptopAdmins:t_laptopadmins, AllowMultipleCompanyMembers:t_allowmultiplecompanymembers, EnforceFSQuota:t_enforcefsquota, QuotaLimitInGB:t_quotalimitingb, ExternalUserGroupsOperationType:t_externalusergroupsoperationtype, NewDescription:t_newdescription, Enabled:t_enabled, Users:t_users, UserOperationType:t_useroperationtype, RestrictConsoleTypes:t_restrictconsoletypes, NewName:t_newname, AzureGUID:t_azureguid, ConsoleTypeOperationType:t_consoletypeoperationtype, PlanOperationType:t_planoperationtype, AssociatedExternalGroups:t_associatedexternalgroups}
    _, err := handler.CvModifyUserGroup(req, d.Id())
    if err != nil {
        return fmt.Errorf("Operation [ModifyUserGroup] failed, Error %s", err)
    }
    return resourceReadUserGroup(d, m)
}

func resourceCreateUpdateUserGroup(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/UserGroup/{userGroupId}
    var execUpdate bool = false
    var t_enabletwofactorauthentication *string
    if v, ok := d.GetOk("enabletwofactorauthentication"); ok {
        val := v.(string)
        t_enabletwofactorauthentication = new(string)
        t_enabletwofactorauthentication = &val
        execUpdate = true
    }
    var t_laptopadmins *bool
    if v, ok := d.GetOkExists("laptopadmins"); ok {
        val := v.(bool)
        t_laptopadmins = new(bool)
        t_laptopadmins = &val
        execUpdate = true
    }
    var t_allowmultiplecompanymembers *bool
    if v, ok := d.GetOkExists("allowmultiplecompanymembers"); ok {
        val := v.(bool)
        t_allowmultiplecompanymembers = new(bool)
        t_allowmultiplecompanymembers = &val
        execUpdate = true
    }
    var t_externalusergroupsoperationtype *string
    if d.HasChange("associatedexternalgroups") {
        var c_externalusergroupsoperationtype string = "OVERWRITE"
        t_externalusergroupsoperationtype = new(string)
        t_externalusergroupsoperationtype = &c_externalusergroupsoperationtype
    }
    var t_enabled *bool
    if v, ok := d.GetOkExists("enabled"); ok {
        val := v.(bool)
        t_enabled = new(bool)
        t_enabled = &val
        execUpdate = true
    }
    var t_users []handler.MsgIdNameSet
    if v, ok := d.GetOk("users"); ok {
        val := v.(*schema.Set)
        t_users = build_usergroup_msgidnameset_array(d, val.List())
        execUpdate = true
    }
    var t_useroperationtype *string
    if d.HasChange("users") {
        var c_useroperationtype string = "OVERWRITE"
        t_useroperationtype = new(string)
        t_useroperationtype = &c_useroperationtype
    }
    var t_restrictconsoletypes *handler.MsgRestrictConsoleTypes
    if v, ok := d.GetOk("restrictconsoletypes"); ok {
        val := v.([]interface{})
        t_restrictconsoletypes = build_usergroup_msgrestrictconsoletypes(d, val)
        execUpdate = true
    }
    var t_azureguid *string
    if v, ok := d.GetOk("azureguid"); ok {
        val := v.(string)
        t_azureguid = new(string)
        t_azureguid = &val
        execUpdate = true
    }
    var t_consoletypeoperationtype *string
    if d.HasChange("restrictconsoletypes") {
        var c_consoletypeoperationtype string = "OVERWRITE"
        t_consoletypeoperationtype = new(string)
        t_consoletypeoperationtype = &c_consoletypeoperationtype
    }
    var t_planoperationtype *string
    if v, ok := d.GetOk("planoperationtype"); ok {
        val := v.(string)
        t_planoperationtype = new(string)
        t_planoperationtype = &val
        execUpdate = true
    }
    var t_associatedexternalgroups []handler.MsgIdNameSet
    if v, ok := d.GetOk("associatedexternalgroups"); ok {
        val := v.(*schema.Set)
        t_associatedexternalgroups = build_usergroup_msgidnameset_array(d, val.List())
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgModifyUserGroupRequest{EnableTwoFactorAuthentication:t_enabletwofactorauthentication, LaptopAdmins:t_laptopadmins, AllowMultipleCompanyMembers:t_allowmultiplecompanymembers, ExternalUserGroupsOperationType:t_externalusergroupsoperationtype, Enabled:t_enabled, Users:t_users, UserOperationType:t_useroperationtype, RestrictConsoleTypes:t_restrictconsoletypes, AzureGUID:t_azureguid, ConsoleTypeOperationType:t_consoletypeoperationtype, PlanOperationType:t_planoperationtype, AssociatedExternalGroups:t_associatedexternalgroups}
        _, err := handler.CvModifyUserGroup(req, d.Id())
        if err != nil {
            return fmt.Errorf("Operation [ModifyUserGroup] failed, Error %s", err)
        }
    }
    return resourceReadUserGroup(d, m)
}

func resourceDeleteUserGroup(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/UserGroup/{userGroupId}
    _, err := handler.CvDeleteUserGroup(d.Id())
    if err != nil {
        return fmt.Errorf("Operation [DeleteUserGroup] failed, Error %s", err)
    }
    return nil
}

func build_usergroup_msgidnameset_array(d *schema.ResourceData, r []interface{}) []handler.MsgIdNameSet {
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

func build_usergroup_msgrestrictconsoletypes(d *schema.ResourceData, r []interface{}) *handler.MsgRestrictConsoleTypes {
    if r != nil && len(r) > 0 {
        tmp := r[0].(map[string]interface{})
        var t_consoletype []string
        if val, ok := tmp["consoletype"].(*schema.Set); ok {
            t_consoletype = handler.ToStringArray(val.List())
        }
        return &handler.MsgRestrictConsoleTypes{ConsoleType:t_consoletype}
    } else {
        return nil
    }
}

func serialize_usergroup_msgidnameset_array(data []handler.MsgIdNameSet) []map[string]interface{} {
    val := make([]map[string]interface{}, len(data))
    for i := range data {
        val[i] = make(map[string]interface{})
        if data[i].Id != nil {
            val[i]["id"] = data[i].Id
        }
    }
    return val
}

func serialize_usergroup_msgrestrictedconsoletypesset_array(data []handler.MsgRestrictedConsoleTypesSet) []map[string]interface{} {
    return handler.GetConsoleTypes(data)
}
