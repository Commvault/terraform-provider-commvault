package commvault

import (
	"os"

	"terraform-provider-commvault/commvault/handler"

	"net/url"
	"regexp"

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
	secured := data.Get("secured").(bool)

	CSUrl := ""
	if isValidUrl(cvCsip) {
		u, err := url.Parse(cvCsip)
		if err != nil {
			panic(err)
		}
		CSUrl = u.Hostname()
	} else {
		re := regexp.MustCompile(`^(?:http?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
		submatchall := re.FindAllString(cvCsip, -1)
		for _, element := range submatchall {
			CSUrl = CSUrl + element
		}
	}

	if secured {
		CSUrl = "https://" + CSUrl + ":81/SearchSvc/CVWebService.svc"
	} else {
		CSUrl = "http://" + CSUrl + ":81/SearchSvc/CVWebService.svc"
	}
	os.Setenv("CV_CSIP", CSUrl)
	os.Setenv("CV_USERNAME", username)
	os.Setenv("CV_PASSWORD", password)
	handler.LoginWithProviderCredentials(username, password)
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
