package commvault

import (
	"fmt"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"user_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The user name for the account.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The password for the account.",
			},
			"full_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies The first, middle, and last names of the user.",
			},
			"email": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies The email address of the user.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies The description of the user account.",
			},
			"new_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Spcifies the new name for the user",
			},
			"company_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Specifies the company id for which the created user will be associated with.",
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	var createUserRequest handler.AppCreateUserRequest
	createUserRequest.Users.UserEntity.UserName = d.Get("user_name").(string)
	createUserRequest.Users.FullName = d.Get("full_name").(string)
	createUserRequest.Users.Password = d.Get("password").(string)
	createUserRequest.Users.Email = d.Get("email").(string)
	createUserRequest.Users.Description = d.Get("description").(string)
	userResp := handler.UserCreate(createUserRequest, d.Get("company_id").(int))
	errorString := userResp.Response.ErrorString
	if errorString != "Successful" {
		return fmt.Errorf("User creation failed with Error: " + userResp.ErrorMessage)
	}
	userID := userResp.Response.Entity.UserId
	d.SetId(string(userID))
	return resourceUserRead(d, m)
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	userID := d.Id()
	userProperties := handler.ReadUser(userID)
	email := userProperties.Users.Email
	if email != "" {
		return fmt.Errorf("Cannot find the user")
	}
	return nil
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	var updateUserRequest handler.AppUpdateUserPropertiesRequest
	newName := d.Get("new_name").(string)
	if newName == "" {
		updateUserRequest.Users.UserEntity.NewName = d.Get("full_name").(string)
	} else {
		updateUserRequest.Users.UserEntity.NewName = newName
	}
	updateUserRequest.Users.FullName = d.Get("full_name").(string)
	updateUserRequest.Users.Email = d.Get("email").(string)
	updateUserRequest.Users.UserEntity.UserName = d.Get("user_name").(string)
	updateUserRequest.Users.UserEntity.UserId = d.Id()
	userResp := handler.UpdateUser(updateUserRequest, d.Id())
	errorCode := userResp.Response.ErrorCode
	if errorCode != "0" {
		return fmt.Errorf("unable to update the user")
	}
	d.SetId(d.Id())
	return resourceUserRead(d, m)
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	deleteUserResp := handler.UserDelete(id)
	errorString := deleteUserResp.Response.ErrorString
	if errorString != "Successful" {
		return nil
	}
	d.SetId("")
	return nil
}
