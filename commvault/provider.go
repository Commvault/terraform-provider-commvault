package commvault

import (
	"os"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"net/url"

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
			"secured": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Specifies if the connection should be secured https or non secured http",
			},
			"api_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the encrypted token for the user to authentication to Web Server.",
			},
			"logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "",
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
			"commvault_user_v2":              resourceUser_V2(),
			"commvault_usergroup":            resourceUserGroup(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"commvault_user":        datasourceUser(),
			"commvault_credential":  datasourceCredential(),
			"commvault_client":      datasourceClient(),
			"commvault_clientgroup": datasourceClientGroup(),
			"commvault_company":     datasourceCompany(),
			"commvault_plan":        datasourcePlan(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (i interface{}, err error) {
	cvCsip := data.Get("web_service_url").(string)
	username := data.Get("user_name").(string)
	password := data.Get("password").(string)
	api_token := data.Get("api_token").(string)
	logging := data.Get("logging").(bool)

	os.Setenv("CV_CSIP", cvCsip)
	os.Setenv("CV_USERNAME", username)
	os.Setenv("CV_PASSWORD", password)
	os.Setenv("CV_LOGGING", strconv.FormatBool(logging))

	if api_token != "" {
		os.Setenv("AuthToken", api_token)
	} else if os.Getenv("CV_TER_TOKEN") != "" {
		os.Setenv("AuthToken", os.Getenv("CV_TER_TOKEN"))
	} else if os.Getenv("CV_TER_PASSWORD") != "" {
		handler.LoginWithProviderCredentials(username, os.Getenv("CV_TER_PASSWORD"))
	} else {
		handler.LoginWithProviderCredentials(username, password)
	}

	return i, nil
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
