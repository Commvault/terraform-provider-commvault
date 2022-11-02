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
            "password": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Sensitive:    true,
                Description: "Used to provide a password to the user being created. This will be accepted when the useSystemGeneratePassword tag is false. The password has to be provided in Base64 format.",
            },
            "name": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Used to provide the new user with a username. This username can be used for logging in the user instead of email-id when duplicate email-ids are present. For external user, it is necessary to provide the domain name along with the username (domainName\\username). To create a company user, the company id or name needs to be provided in the company entity.",
            },
            "fullname": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: "Used to provide a name to the new user.",
            },
            "company": &schema.Schema{
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
            "usesystemgeneratepassword": &schema.Schema{
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: "Choose to provide a system generated password to the user instead of providing your own password. An email will be sent to the user to reset the password. If it is set to true, password tag need not be provided. If it is set to false, password needs to be provided in the password tag in Base64 format.",
            },
            "inviteuser": &schema.Schema{
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: "User will receive an email to install backup software package on their device if this is set to true.",
            },
            "plan": &schema.Schema{
                Type:        schema.TypeList,
                Optional:    true,
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
            "email": &schema.Schema{
                Type:        schema.TypeString,
                Required:    true,
                Description: "Used to provide an email-id to the new user. This email-id is used for logging in the user. Please note that email ids are compulsory for company and local users and optional for external users.",
            },
            "enabled": &schema.Schema{
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: "enable or disable the user.",
            },
            "userprincipalname": &schema.Schema{
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
        return fmt.Errorf("Operation [CreateUser] failed, Error %s", err)
    }
    if resp.Users != nil && len(resp.Users) > 0 {
        if resp.Users[0].Id != nil {
            response_id = strconv.Itoa(*resp.Users[0].Id)
        }
    }
    if response_id == "0" {
        return fmt.Errorf("Operation [CreateUser] failed")
    } else {
        d.SetId(response_id)
        return resourceCreateUpdateUser_V2(d, m)
    }
}

func resourceReadUser_V2(d *schema.ResourceData, m interface{}) error {
    //API: (GET) /V4/user/{userId}
    resp, err := handler.CvGetUserDetails(d.Id())
    if err != nil {
        return fmt.Errorf("Operation [GetUserDetails] failed, Error %s", err)
    }
    if resp.FullName != nil {
        d.Set("fullname", resp.FullName)
    }
    if resp.Enabled != nil {
        d.Set("enabled", resp.Enabled)
    }
    if resp.Name != nil {
        d.Set("name", resp.Name)
    }
    if resp.Company != nil {
        d.Set("company", serialize_user_v2_msgidname(resp.Company))
    } else {
        d.Set("company", make([]map[string]interface{}, 0))
    }
    if resp.Plan != nil {
        d.Set("plan", serialize_user_v2_msgidname(resp.Plan))
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
        val := d.Get("name").(string)
        t_newname = new(string)
        t_newname = &val
    }
    var t_fullname *string
    if d.HasChange("fullname") {
        val := d.Get("fullname").(string)
        t_fullname = new(string)
        t_fullname = &val
    }
    var t_newpassword *string
    if d.HasChange("password") {
        val := d.Get("password").(string)
        t_newpassword = new(string)
        t_newpassword = &val
    }
    var t_plan *handler.MsgIdName
    if d.HasChange("plan") {
        val := d.Get("plan").([]interface{})
        t_plan = build_user_v2_msgidname(d, val)
    }
    var t_email *string
    if d.HasChange("email") {
        val := d.Get("email").(string)
        t_email = new(string)
        t_email = &val
    }
    var t_enabled *bool
    if d.HasChange("enabled") {
        val := d.Get("enabled").(bool)
        t_enabled = new(bool)
        t_enabled = &val
    }
    var t_userprincipalname *string
    if d.HasChange("userprincipalname") {
        val := d.Get("userprincipalname").(string)
        t_userprincipalname = new(string)
        t_userprincipalname = &val
    }
    var req = handler.MsgModifyUserRequest{NewName:t_newname, FullName:t_fullname, NewPassword:t_newpassword, Plan:t_plan, Email:t_email, Enabled:t_enabled, UserPrincipalName:t_userprincipalname}
    h_err := handler.UpdateUserRequest(&req, d, m)
    if h_err != nil {
        return fmt.Errorf("Operation [ModifyUser] failed, Error %s", h_err)
    }
    _, err := handler.CvModifyUser(req, d.Id())
    if err != nil {
        return fmt.Errorf("Operation [ModifyUser] failed, Error %s", err)
    }
    return resourceReadUser_V2(d, m)
}

func resourceCreateUpdateUser_V2(d *schema.ResourceData, m interface{}) error {
    //API: (PUT) /V4/user/{userId}
    var execUpdate bool = false
    var t_enabled *bool
    if v, ok := d.GetOkExists("enabled"); ok {
        val := v.(bool)
        t_enabled = new(bool)
        t_enabled = &val
        execUpdate = true
    }
    var t_userprincipalname *string
    if v, ok := d.GetOk("userprincipalname"); ok {
        val := v.(string)
        t_userprincipalname = new(string)
        t_userprincipalname = &val
        execUpdate = true
    }
    if execUpdate {
        var req = handler.MsgModifyUserRequest{Enabled:t_enabled, UserPrincipalName:t_userprincipalname}
        h_err := handler.UpdateUserRequest(&req, d, m)
        if h_err != nil {
            return fmt.Errorf("Operation [ModifyUser] failed, Error %s", h_err)
        }
        _, err := handler.CvModifyUser(req, d.Id())
        if err != nil {
            return fmt.Errorf("Operation [ModifyUser] failed, Error %s", err)
        }
    }
    return resourceReadUser_V2(d, m)
}

func resourceDeleteUser_V2(d *schema.ResourceData, m interface{}) error {
    //API: (DELETE) /V4/user/{userId}
    _, err := handler.CvDeleteUser(d.Id())
    if err != nil {
        return fmt.Errorf("Operation [DeleteUser] failed, Error %s", err)
    }
    return nil
}

func build_user_v2_msgidname(d *schema.ResourceData, r []interface{}) *handler.MsgIdName {
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

func build_user_v2_msgcreateuserset_array(d *schema.ResourceData) []handler.MsgCreateUserSet {
    tmp := make([]handler.MsgCreateUserSet, 1)
    var t_password *string
    if v, ok := d.GetOk("password"); ok {
        val := v.(string)
        t_password = new(string)
        t_password = &val
    }
    var t_name *string
    if v, ok := d.GetOk("name"); ok {
        val := v.(string)
        t_name = new(string)
        t_name = &val
    }
    var t_fullname *string
    if v, ok := d.GetOk("fullname"); ok {
        val := v.(string)
        t_fullname = new(string)
        t_fullname = &val
    }
    var t_company *handler.MsgIdName
    if v, ok := d.GetOk("company"); ok {
        val := v.([]interface{})
        t_company = build_user_v2_msgidname(d, val)
    }
    var t_usesystemgeneratepassword *bool
    if v, ok := d.GetOkExists("usesystemgeneratepassword"); ok {
        val := v.(bool)
        t_usesystemgeneratepassword = new(bool)
        t_usesystemgeneratepassword = &val
    }
    var t_inviteuser *bool
    if v, ok := d.GetOkExists("inviteuser"); ok {
        val := v.(bool)
        t_inviteuser = new(bool)
        t_inviteuser = &val
    }
    var t_plan *handler.MsgIdName
    if v, ok := d.GetOk("plan"); ok {
        val := v.([]interface{})
        t_plan = build_user_v2_msgidname(d, val)
    }
    var t_email *string
    if v, ok := d.GetOk("email"); ok {
        val := v.(string)
        t_email = new(string)
        t_email = &val
    }
    tmp[0] = handler.MsgCreateUserSet{Password:t_password, Name:t_name, FullName:t_fullname, Company:t_company, UseSystemGeneratePassword:t_usesystemgeneratepassword, InviteUser:t_inviteuser, Plan:t_plan, Email:t_email}
    return tmp
}

func serialize_user_v2_msgidname(data *handler.MsgIdName) map[string]interface{} {
    val := make(map[string]interface{})
    if data.Name != nil {
        val["name"] = data.Name
    }
    if data.Id != nil {
        val["id"] = data.Id
    }
    return val
}
