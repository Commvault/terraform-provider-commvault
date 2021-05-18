package commvault

import (
	"os"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"web_service_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CV_CSIP", os.Getenv("CV_CSIP")),
				Description: "Specifies the Web Server URL of the commserver for performing Terraform Operations.",
			},
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CV_USERNAME", os.Getenv("CV_USERNAME")),
				Description: "Specifies the User name used for authentication to Web Server",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CV_PASSWORD", os.Getenv("CV_PASSWORD")),
				Description: "Specifies the Password for the user name to authentication to Web Server.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"commvault_plan":                 resourcePlan(),
			"commvault_user":                 resourceUser(),
			"commvault_login":                resourceLogin(),
			"commvault_vm_group":             resourceVMGroup(),
			"commvault_vmware_hypervisor":    resourceVMWareHypervisor(),
			"commvault_amazon_hypervisor":    resourceAmazonHypervisor(),
			"commvault_azure_hypervisor":     resourceAzureHypervisor(),
			"commvault_plan_to_vm":           resourceAssociateVMToPlan(),
			"commvault_company":              resourceCompany(),
			"commvault_disk_storage":         resourceDiskStorage(),
			"commvault_aws_storage":          resourceAWSStorage(),
			"commvault_azure_storage":        resourceAzureStorage(),
			"commvault_google_storage":       resourceGoogleStorage(),
			"commvault_install_ma":           resourceInstallMA(),
			"commvault_security_association": securityAssociation(),
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
