package main

import (
	"os"

	"github.com/TerraformProvider/handler"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"web_service_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CV_CSIP", os.Getenv("CV_CSIP")),
			},
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CV_USERNAME", os.Getenv("CV_USERNAME")),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CV_PASSWORD", os.Getenv("CV_PASSWORD")),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"commvault_plan":          resourcePlan(),
			"commvault_user":          resourceUser(),
			"commvault_login":         resourceLogin(),
			"commvault_vm_group":      resourceVMGroup(),
			"commvault_vmware_hyperv": resourceVMWareHypervisor(),
			"commvault_plan_to_vm":    resourceAssociateVMToPlan(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (i interface{}, err error) {
	cvCsip := data.Get("web_service_url").(string)
	username := data.Get("user_name").(string)
	password := data.Get("password").(string)
	os.Setenv("CV_CSIP", cvCsip)
	os.Setenv("CV_USERNAME", username)
	os.Setenv("CV_PASSWORD", password)
	handler.LoginWithProviderCredentials(username, password)
	return i, nil
}
