package commvault

import (
    "strconv"
    "fmt"

    "terraform-provider-commvault/commvault/handler"

    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceUser_V2() *schema.Resource {
    return &schema.Resource{
        Create: resourceCreateUser_V2,
        Read:   resourceReadUser_V2,
        Update: resourceUpdateUser_V2,
        Delete: resourceDeleteUser_V2,

        Schema: map[string]*schema.Schema{
            "password": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Sensitive:    true,
                Description: "Used to provide a password to the user being created. This will be accepted when the useSystemGeneratePassword tag is false. The password has to be provided in Base64 format.",
            },
            "name": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Used to provide the new user with a username. This username can be used for logging in the user instead of email-id when duplicate email-ids are present. For external user, it is necessary to provide the domain name along with the username (domainName\\username). To create a company user, the company id or name needs to be provided in the company entity.",
            },
            "fullname": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Used to provide a name to the new user.",
            },
            "company": {
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
            "usesystemgeneratepassword": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Choose to provide a system generated password to the user instead of providing your own password. An email will be sent to the user to reset the password. If it is set to true, password tag need not be provided. If it is set to false, password needs to be provided in the password tag in Base64 format.",
            },
            "inviteuser": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "User will receive an email to install backup software package on their device if this is set to true.",
            },
            "plan": {
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
            "email": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Used to provide an email-id to the new user. This email-id is used for logging in the user. Please note that email ids are compulsory for company and local users and optional for external users.",
            },
            "authenticationmethod": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Change the current authentication method of user. SAML user association can be removed using this.",
            },
            "enabled": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "enable or disable the user.",
            },
            "userprincipalname": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Change User Principal Name(UPN) for existing user. This User Principal Name can be used for logging-in.",
            },
        },
    }
}

func resourceCreateUser_V2(d *schema.ResourceData, m interface{}) error {
    //API: (POST) /V4/user
    var response_id = strconv.Itoa(0)
    t_users := build_user_v2_msgcreateuserset_array(d)
    var req = handler.MsgCreateUserRequest{Users:t_users}
    resp, err := handler.CvCreateUser(req)
    if err != nil {
        return fmt.Errorf("operation [CreateUser] failed, Error %s", err)
    }
    if resp.Users != nil && len(resp.Users) > 0 {
        if resp.Users[0].Id != nil {
            response_id = strconv.Itoa(*resp.Users[0].Id)
        }
    }
    if response_id == "0" {
        return fmt.Errorf("operation [CreateUser] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateUser_V2(d, m)
    }
}

func resourceReadUser_V2(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/user/{userId}
    resp, err := handler.CvGetUserDetails(d.Id())
    if err != nil {
        return fmt.Errorf("operation [GetUserDetails] failed, Error %s", err)
    }
    if resp.FullName != nil {
        d.Set("fullname", resp.FullName)
    }
    if resp.Enabled != nil {
        d.Set("enabled", strconv.FormatBool(*resp.Enabled))
    }
    if resp.AuthenticationMethod != nil {
        d.Set("authenticationmethod", resp.AuthenticationMethod)
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    if rtn, ok := serialize_user_v2_msgidname(d, resp.Company); ok {
        d.Set("company", rtn)
    } else {
        d.Set("company", make([]map[string]interface{}, 0))
    }
    if rtn, ok := serialize_user_v2_msgidname(d, resp.Plan); ok {
        d.Set("plan", rtn)
    } else {
        d.Set("plan", make([]map[string]interface{}, 0))
    }
    if resp.Email != nil {
        d.Set("email", resp.Email)
    }
    if resp.UserPrincipalName != nil {
        d.Set("userprincipalname", resp.UserPrincipalName)
    }
    return nil
}

func resourceUpdateUser_V2(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/user/{userId}
    var t_newname *string
    if d.HasChange("name") {
        val := d.Get("name")
        t_newname = handler.ToStringValue(val, false)
    }
    var t_authenticationmethod *string
    if d.HasChange("authenticationmethod") {
        val := d.Get("authenticationmethod")
        t_authenticationmethod = handler.ToStringValue(val, false)
    }
    var t_fullname *string
    if d.HasChange("fullname") {
        val := d.Get("fullname")
        t_fullname = handler.ToStringValue(val, false)
    }
    var t_newpassword *string
    if d.HasChange("password") {
        val := d.Get("password")
        t_newpassword = handler.ToStringValue(val, false)
    }
    var t_plan *handler.MsgIdName
    if d.HasChange("plan") {
        val := d.Get("plan")
        t_plan = build_user_v2_msgidname(d, val.([]interface{}))
    }
    var t_email *string
    if d.HasChange("email") {
        val := d.Get("email")
        t_email = handler.ToStringValue(val, false)
    }
    var t_enabled *bool
    if d.HasChange("enabled") {
        val := d.Get("enabled")
        t_enabled = handler.ToBooleanValue(val, false)
    }
    var t_userprincipalname *string
    if d.HasChange("userprincipalname") {
        val := d.Get("userprincipalname")
        t_userprincipalname = handler.ToStringValue(val, false)
    }
    var req = handler.MsgModifyUserRequest{NewName:t_newname, AuthenticationMethod:t_authenticationmethod, FullName:t_fullname, NewPassword:t_newpassword, Plan:t_plan, Email:t_email, Enabled:t_enabled, UserPrincipalName:t_userprincipalname}
    h_err := handler.UpdateUserRequest(&req, d, m)
    if h_err != nil {
        return fmt.Errorf("operation [ModifyUser] failed, Error %s", h_err)
    }
    _, err := handler.CvModifyUser(req, d.Id())
    if err != nil {
        return fmt.Errorf("operation [ModifyUser] failed, Error %s", err)
    }
    return resourceReadUser_V2(d, m)
}

func resourceCreateUpdateUser_V2(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/user/{userId}
    var execUpdate bool = false
    var t_authenticationmethod *string
    if val, ok := d.GetOk("authenticationmethod"); ok {
        t_authenticationmethod = handler.ToStringValue(val, false)
        execUpdate = true
    }
    var t_enabled *bool
    if val, ok := d.GetOk("enabled"); ok {
        t_enabled = handler.ToBooleanValue(val, false)
        execUpdate = true
    }
    var t_userprincipalname *string
    if val, ok := d.GetOk("userprincipalname"); ok {
        t_userprincipalname = handler.ToStringValue(val, false)
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgModifyUserRequest{AuthenticationMethod:t_authenticationmethod, Enabled:t_enabled, UserPrincipalName:t_userprincipalname}
        h_err := handler.UpdateUserRequest(&req, d, m)
        if h_err != nil {
            return fmt.Errorf("operation [ModifyUser] failed, Error %s", h_err)
        }
        _, err := handler.CvModifyUser(req, d.Id())
        if err != nil {
            return fmt.Errorf("operation [ModifyUser] failed, Error %s", err)
        }
    }
    return resourceReadUser_V2(d, m)
}

func resourceDeleteUser_V2(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/user/{userId}
    _, err := handler.CvDeleteUser(d.Id())
    if err != nil {
        return fmt.Errorf("operation [DeleteUser] failed, Error %s", err)
    }
    return nil
}

func build_user_v2_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_user_v2_msgcreateuserset_array(d *schema.ResourceData) []handler.MsgCreateUserSet {
    tmp := make([]handler.MsgCreateUserSet, 1)
    var t_password *string
    if val, ok := d.GetOk("password"); ok {
        t_password = handler.ToStringValue(val, false)
    }
    var t_name *string
    if val, ok := d.GetOk("name"); ok {
        t_name = handler.ToStringValue(val, false)
    }
    var t_fullname *string
    if val, ok := d.GetOk("fullname"); ok {
        t_fullname = handler.ToStringValue(val, false)
    }
    var t_company *handler.MsgIdName
    if val, ok := d.GetOk("company"); ok {
        t_company = build_user_v2_msgidname(d, val.([]interface{}))
    }
    var t_usesystemgeneratepassword *bool
    if val, ok := d.GetOk("usesystemgeneratepassword"); ok {
        t_usesystemgeneratepassword = handler.ToBooleanValue(val, false)
    }
    var t_inviteuser *bool
    if val, ok := d.GetOk("inviteuser"); ok {
        t_inviteuser = handler.ToBooleanValue(val, false)
    }
    var t_plan *handler.MsgIdName
    if val, ok := d.GetOk("plan"); ok {
        t_plan = build_user_v2_msgidname(d, val.([]interface{}))
    }
    var t_email *string
    if val, ok := d.GetOk("email"); ok {
        t_email = handler.ToStringValue(val, false)
    }
    tmp[0] = handler.MsgCreateUserSet{Password:t_password, Name:t_name, FullName:t_fullname, Company:t_company, UseSystemGeneratePassword:t_usesystemgeneratepassword, InviteUser:t_inviteuser, Plan:t_plan, Email:t_email}
    return tmp
}

func serialize_user_v2_msgidname(d *schema.ResourceData, data *handler.MsgIdName) ([]map[string]interface{}, bool) {
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
