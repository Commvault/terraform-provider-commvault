package main

import (
	"fmt"

	"github.com/TerraformProvider/handler"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"fullname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	var createUserRequest handler.AppCreateUserRequest
	createUserRequest.Users.UserEntity.UserName = d.Get("username").(string)
	createUserRequest.Users.FullName = d.Get("fullname").(string)
	createUserRequest.Users.Password = d.Get("password").(string)
	createUserRequest.Users.Email = d.Get("email").(string)
	createUserRequest.Users.Description = d.Get("description").(string)
	userResp := handler.UserCreate(createUserRequest)
	errorString := userResp.Response.ErrorString
	if errorString != "Successful" {
		return fmt.Errorf("User creation failed")
	}
	userId := userResp.Response.Entity.UserId
	d.SetId(string(userId))
	return resourceUserRead(d, m)
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	userId := d.Id()
	userProperties := handler.ReadUser(userId)
	email := userProperties.Users.Email
	if email != "" {
		return fmt.Errorf("Cannot find the user")
		d.SetId("")
	}
	return nil
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	var updateUserRequest handler.AppUpdateUserPropertiesRequest
	newName := d.Get("newname").(string)
	if newName == "" {
		updateUserRequest.Users.UserEntity.NewName = d.Get("fullname").(string)
	} else {
		updateUserRequest.Users.UserEntity.NewName = newName
	}
	updateUserRequest.Users.FullName = d.Get("fullname").(string)
	updateUserRequest.Users.Email = d.Get("email").(string)
	updateUserRequest.Users.UserEntity.UserName = d.Get("username").(string)
	updateUserRequest.Users.UserEntity.UserId = d.Id()
	userResp := handler.UpdateUser(updateUserRequest, d.Id())
	errorCode := userResp.Response.ErrorCode
	if errorCode != "0" {
		fmt.Errorf("unable to update the user")
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
