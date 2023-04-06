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
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "To create an active directory usergroup, the domain name should be mentioned along with the usergroup name (domainName\\usergroupName) and localUserGroup value must be given.",
            },
            "description": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "",
            },
            "enforcefsquota": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Used to determine if a backup data limit will be set for the user group being created",
            },
            "quotalimitingb": {
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: "if enforceFSQuota is set to true, the quota limit can be set in GBs",
            },
            "enablelocalauthentication": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Allows two-factor authentication to be enabled for the specific types of usergroups. it can be turned on or off based on user preferences. There will be usergroups that will not have this option. [ON, OFF, DISABLED_AT_COMPANY]",
            },
            "enabletwofactorauthentication": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Allows two-factor authentication to be enabled for the specific types of usergroups. it can be turned on or off based on user preferences. There will be usergroups that will not have this option. [ON, OFF, DISABLED_AT_COMPANY]",
            },
            "laptopadmins": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "When set to true, users in this group cannot activate or be set as server owner",
            },
            "allowmultiplecompanymembers": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "This property can be used to allow addition of users/groups from child companies. Only applicable for commcell and reseller company group.",
            },
            "enabled": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "allows the enabling/disabling of the user group.",
            },
            "users": {
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
            "restrictconsoletypes": {
                Type:        schema.TypeList,
                Optional:    true,
                Computed:    true,
                Description: "",
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "consoletype": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: "",
                            Elem: &schema.Schema{
                                Type:    schema.TypeString,
                            },
                        },
                    },
                },
            },
            "azureguid": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Azure Object ID used to link this user group to Azure AD group and manage group membership of the user during SAML login",
            },
            "donotinheritrestrictconsoletypes": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Option to not inherit the RestrictConsoleTypes from the parent. By default the value is false, parent RestrictConsoleTypes will be inherited.",
            },
            "planoperationtype": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "determines if an existing user has to be added to the user group or removed from the user group [ADD, DELETE]",
            },
            "associatedexternalgroups": {
                Type:        schema.TypeSet,
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
    }
}

func resourceCreateUserGroup(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/UserGroup
    var response_id = strconv.Itoa(0)
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_description *string
    if val, ok := d.GetOk("description"); ok {
        t_description = handler.ToStringValue(val, false)
    }
    var t_enforcefsquota *bool
    if val, ok := d.GetOk("enforcefsquota"); ok {
        t_enforcefsquota = handler.ToBooleanValue(val, false)
    }
    var t_quotalimitingb *int
    if val, ok := d.GetOk("quotalimitingb"); ok {
        t_quotalimitingb = handler.ToIntValue(val, false)
    }
    var req = handler.MsgCreateUserGroupRequest{Name:t_name, Description:t_description, EnforceFSQuota:t_enforcefsquota, QuotaLimitInGB:t_quotalimitingb}
    resp, err := handler.CvCreateUserGroup(req)
    if err != nil {
        return fmt.Errorf("operation [CreateUserGroup] failed, Error %s", err)
    }
    if resp.Id != nil {
        response_id = strconv.Itoa(*resp.Id)
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateUserGroup] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateUserGroup(d, m)
    }
}

func resourceReadUserGroup(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/UserGroup/{userGroupId}
    resp, err := handler.CvGetUserGroupDetails(d.Id())
    if err != nil {
        return fmt.Errorf("operation [GetUserGroupDetails] failed, Error %s", err)
    }
    if rtn, ok := handler.GetConsoleTypes(d, resp.RestrictedConsoleTypes); ok {
        d.Set("restrictconsoletypes", rtn)
    } else {
        d.Set("restrictconsoletypes", make([]map[string]interface{}, 0))
    }
    if resp.EnableLocalAuthentication != nil {
        d.Set("enablelocalauthentication", resp.EnableLocalAuthentication)
    }
    if resp.EnableTwoFactorAuthentication != nil {
        d.Set("enabletwofactorauthentication", resp.EnableTwoFactorAuthentication)
    }
    if resp.LaptopAdmins != nil {
        d.Set("laptopadmins", strconv.FormatBool(*resp.LaptopAdmins))
    }
    if resp.AllowMultipleCompanyMembers != nil {
        d.Set("allowmultiplecompanymembers", strconv.FormatBool(*resp.AllowMultipleCompanyMembers))
    }
    if resp.Description != nil {
        d.Set("description", resp.Description)
    }
    if resp.EnforceFSQuota != nil {
        d.Set("enforcefsquota", strconv.FormatBool(*resp.EnforceFSQuota))
    }
    if resp.QuotaLimitInGB != nil {
        d.Set("quotalimitingb", resp.QuotaLimitInGB)
    }
    if resp.Enabled != nil {
        d.Set("enabled", strconv.FormatBool(*resp.Enabled))
    }
    if rtn, ok := serialize_usergroup_msgidnameset_array(d, resp.Users); ok {
        d.Set("users", rtn)
    } else {
        d.Set("users", make([]map[string]interface{}, 0))
    }
    if resp.AzureGUID != nil {
        d.Set("azureguid", resp.AzureGUID)
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    if resp.DoNotInheritRestrictConsoleTypes != nil {
        d.Set("donotinheritrestrictconsoletypes", strconv.FormatBool(*resp.DoNotInheritRestrictConsoleTypes))
    }
    if rtn, ok := serialize_usergroup_msgidnameset_array(d, resp.AssociatedExternalGroups); ok {
        d.Set("associatedexternalgroups", rtn)
    } else {
        d.Set("associatedexternalgroups", make([]map[string]interface{}, 0))
    }
    return nil
}

func resourceUpdateUserGroup(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/UserGroup/{userGroupId}
    var t_enablelocalauthentication *string
    if d.HasChange("enablelocalauthentication") {
        val := d.Get("enablelocalauthentication")
        t_enablelocalauthentication = handler.ToStringValue(val, false)
    }
    var t_enabletwofactorauthentication *string
    if d.HasChange("enabletwofactorauthentication") {
        val := d.Get("enabletwofactorauthentication")
        t_enabletwofactorauthentication = handler.ToStringValue(val, false)
    }
    var t_laptopadmins *bool
    if d.HasChange("laptopadmins") {
        val := d.Get("laptopadmins")
        t_laptopadmins = handler.ToBooleanValue(val, false)
    }
    var t_allowmultiplecompanymembers *bool
    if d.HasChange("allowmultiplecompanymembers") {
        val := d.Get("allowmultiplecompanymembers")
        t_allowmultiplecompanymembers = handler.ToBooleanValue(val, false)
    }
    var t_enforcefsquota *bool
    if d.HasChange("enforcefsquota") {
        val := d.Get("enforcefsquota")
        t_enforcefsquota = handler.ToBooleanValue(val, false)
    }
    var t_quotalimitingb *int
    if d.HasChange("quotalimitingb") {
        val := d.Get("quotalimitingb")
        t_quotalimitingb = handler.ToIntValue(val, false)
    }
    var t_externalusergroupsoperationtype *string
    if d.HasChange("associatedexternalgroups") {
        var c_externalusergroupsoperationtype string = "OVERWRITE"
        t_externalusergroupsoperationtype = &c_externalusergroupsoperationtype
    }
    var t_newdescription *string
    if d.HasChange("description") {
        val := d.Get("description")
        t_newdescription = handler.ToStringValue(val, false)
    }
    var t_enabled *bool
    if d.HasChange("enabled") {
        val := d.Get("enabled")
        t_enabled = handler.ToBooleanValue(val, false)
    }
    var t_users []handler.MsgIdNameSet
    if d.HasChange("users") {
        val := d.Get("users")
        t_users = build_usergroup_msgidnameset_array(d, val.(*schema.Set).List())
    }
    var t_useroperationtype *string
    if d.HasChange("users") {
        var c_useroperationtype string = "OVERWRITE"
        t_useroperationtype = &c_useroperationtype
    }
    var t_restrictconsoletypes *handler.MsgRestrictConsoleTypes
    if d.HasChange("restrictconsoletypes") {
        val := d.Get("restrictconsoletypes")
        t_restrictconsoletypes = build_usergroup_msgrestrictconsoletypes(d, val.([]interface{}))
    }
    var t_newname *string
    if d.HasChange("name") {
        val := d.Get("name")
        t_newname = handler.ToStringValue(val, false)
    }
    var t_azureguid *string
    if d.HasChange("azureguid") {
        val := d.Get("azureguid")
        t_azureguid = handler.ToStringValue(val, false)
    }
    var t_donotinheritrestrictconsoletypes *bool
    if d.HasChange("donotinheritrestrictconsoletypes") {
        val := d.Get("donotinheritrestrictconsoletypes")
        t_donotinheritrestrictconsoletypes = handler.ToBooleanValue(val, false)
    }
    var t_consoletypeoperationtype *string
    if d.HasChange("restrictconsoletypes") {
        var c_consoletypeoperationtype string = "OVERWRITE"
        t_consoletypeoperationtype = &c_consoletypeoperationtype
    }
    var t_planoperationtype *string
    if d.HasChange("planoperationtype") {
        val := d.Get("planoperationtype")
        t_planoperationtype = handler.ToStringValue(val, false)
    }
    var t_associatedexternalgroups []handler.MsgIdNameSet
    if d.HasChange("associatedexternalgroups") {
        val := d.Get("associatedexternalgroups")
        t_associatedexternalgroups = build_usergroup_msgidnameset_array(d, val.(*schema.Set).List())
    }
    var req = handler.MsgModifyUserGroupRequest{EnableLocalAuthentication:t_enablelocalauthentication, EnableTwoFactorAuthentication:t_enabletwofactorauthentication, LaptopAdmins:t_laptopadmins, AllowMultipleCompanyMembers:t_allowmultiplecompanymembers, EnforceFSQuota:t_enforcefsquota, QuotaLimitInGB:t_quotalimitingb, ExternalUserGroupsOperationType:t_externalusergroupsoperationtype, NewDescription:t_newdescription, Enabled:t_enabled, Users:t_users, UserOperationType:t_useroperationtype, RestrictConsoleTypes:t_restrictconsoletypes, NewName:t_newname, AzureGUID:t_azureguid, DoNotInheritRestrictConsoleTypes:t_donotinheritrestrictconsoletypes, ConsoleTypeOperationType:t_consoletypeoperationtype, PlanOperationType:t_planoperationtype, AssociatedExternalGroups:t_associatedexternalgroups}
    _, err := handler.CvModifyUserGroup(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [ModifyUserGroup] failed, Error %s", err)
    }
    return resourceReadUserGroup(d, m)
}

func resourceCreateUpdateUserGroup(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/UserGroup/{userGroupId}
    var execUpdate bool = false
    var t_enablelocalauthentication *string
    if val, ok := d.GetOk("enablelocalauthentication"); ok {
        t_enablelocalauthentication = handler.ToStringValue(val, false)
        execUpdate = true
    }
    var t_enabletwofactorauthentication *string
    if val, ok := d.GetOk("enabletwofactorauthentication"); ok {
        t_enabletwofactorauthentication = handler.ToStringValue(val, false)
        execUpdate = true
    }
    var t_laptopadmins *bool
    if val, ok := d.GetOk("laptopadmins"); ok {
        t_laptopadmins = handler.ToBooleanValue(val, false)
        execUpdate = true
    }
    var t_allowmultiplecompanymembers *bool
    if val, ok := d.GetOk("allowmultiplecompanymembers"); ok {
        t_allowmultiplecompanymembers = handler.ToBooleanValue(val, false)
        execUpdate = true
    }
    var t_externalusergroupsoperationtype *string
    if d.HasChange("associatedexternalgroups") {
        var c_externalusergroupsoperationtype string = "OVERWRITE"
        t_externalusergroupsoperationtype = &c_externalusergroupsoperationtype
    }
    var t_enabled *bool
    if val, ok := d.GetOk("enabled"); ok {
        t_enabled = handler.ToBooleanValue(val, false)
        execUpdate = true
    }
    var t_users []handler.MsgIdNameSet
    if val, ok := d.GetOk("users"); ok {
        t_users = build_usergroup_msgidnameset_array(d, val.(*schema.Set).List())
        execUpdate = true
    }
    var t_useroperationtype *string
    if d.HasChange("users") {
        var c_useroperationtype string = "OVERWRITE"
        t_useroperationtype = &c_useroperationtype
    }
    var t_restrictconsoletypes *handler.MsgRestrictConsoleTypes
    if val, ok := d.GetOk("restrictconsoletypes"); ok {
        t_restrictconsoletypes = build_usergroup_msgrestrictconsoletypes(d, val.([]interface{}))
        execUpdate = true
    }
    var t_azureguid *string
    if val, ok := d.GetOk("azureguid"); ok {
        t_azureguid = handler.ToStringValue(val, false)
        execUpdate = true
    }
    var t_donotinheritrestrictconsoletypes *bool
    if val, ok := d.GetOk("donotinheritrestrictconsoletypes"); ok {
        t_donotinheritrestrictconsoletypes = handler.ToBooleanValue(val, false)
        execUpdate = true
    }
    var t_consoletypeoperationtype *string
    if d.HasChange("restrictconsoletypes") {
        var c_consoletypeoperationtype string = "OVERWRITE"
        t_consoletypeoperationtype = &c_consoletypeoperationtype
    }
    var t_planoperationtype *string
    if val, ok := d.GetOk("planoperationtype"); ok {
        t_planoperationtype = handler.ToStringValue(val, false)
        execUpdate = true
    }
    var t_associatedexternalgroups []handler.MsgIdNameSet
    if val, ok := d.GetOk("associatedexternalgroups"); ok {
        t_associatedexternalgroups = build_usergroup_msgidnameset_array(d, val.(*schema.Set).List())
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgModifyUserGroupRequest{EnableLocalAuthentication:t_enablelocalauthentication, EnableTwoFactorAuthentication:t_enabletwofactorauthentication, LaptopAdmins:t_laptopadmins, AllowMultipleCompanyMembers:t_allowmultiplecompanymembers, ExternalUserGroupsOperationType:t_externalusergroupsoperationtype, Enabled:t_enabled, Users:t_users, UserOperationType:t_useroperationtype, RestrictConsoleTypes:t_restrictconsoletypes, AzureGUID:t_azureguid, DoNotInheritRestrictConsoleTypes:t_donotinheritrestrictconsoletypes, ConsoleTypeOperationType:t_consoletypeoperationtype, PlanOperationType:t_planoperationtype, AssociatedExternalGroups:t_associatedexternalgroups}
        _, err := handler.CvModifyUserGroup(req, d.Id())
        if err != nil {
            return fmt.Errorf("operation [ModifyUserGroup] failed, Error %s", err)
        }
    }
    return resourceReadUserGroup(d, m)
}

func resourceDeleteUserGroup(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/UserGroup/{userGroupId}
    _, err := handler.CvDeleteUserGroup(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteUserGroup] failed, Error %s", err)
    }
    return nil
}

func build_usergroup_msgidnameset_array(d *schema.ResourceData, r []interface{}) []handler.MsgIdNameSet {
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

func build_usergroup_msgrestrictconsoletypes(d *schema.ResourceData, r []interface{}) *handler.MsgRestrictConsoleTypes {
    if len(r) > 0 && r[0] != nil {
        tmp := r[0].(map[string]interface{})
        var t_consoletype []string
        if val, ok := tmp["consoletype"]; ok {
            t_consoletype = handler.ToStringArray(val.(*schema.Set).List())
        }
        return &handler.MsgRestrictConsoleTypes{ConsoleType:t_consoletype}
    } else {
        return nil
    }
}

func serialize_usergroup_msgidnameset_array(d *schema.ResourceData, data []handler.MsgIdNameSet) ([]map[string]interface{}, bool) {
    //MsgIdNameSet
    //MsgIdNameSet
    if data == nil {
        return nil, false
    }
    val := make([]map[string]interface{}, 0)
    for i := range data {
        tmp := make(map[string]interface{})
        added := false
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
