package commvault

import (
	"fmt"
	"strconv"
	"time"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceInstallMA() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstallMACreate,
		Read:   resourceInstallMARead,
		Update: resourceInstallMAUpdate,
		Delete: resourceInstallMADelete,

		Schema: map[string]*schema.Schema{
			"mediaagent_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the Media Agent name used for installation.",
			},
			"hostname": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the Media Agent Hostname user for the installation",
			},
			"user_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the User name of the host computer for the installation.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the password for the host computer for the installation.",
			},
			"is_unix": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Specifies whether OS is Unix or not",
			},
			"company_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Specifies the company id to which the installed MA should be associated with.",
			},
		},
	}
}

func resourceInstallMACreate(d *schema.ResourceData, m interface{}) error {
	var installMAReq handler.InstallMARequest
	var entities handler.Entities
	entities.ClientName = d.Get("mediaagent_name").(string)
	entities.HostName = d.Get("hostname").(string)
	installMAReq.Entities = append(installMAReq.Entities, entities)
	installMAReq.ClientAuthForJob.UserName = d.Get("user_name").(string)
	installMAReq.ClientAuthForJob.Password = d.Get("password").(string)
	installMAReq.RebootClient = true
	installMAReq.CreatePseudoClientRequest.RegisterClient = true
	isunix := d.Get("is_unix").(bool)
	var packages handler.Packages
	if isunix {
		installMAReq.CreatePseudoClientRequest.ClientInfo.ClientType = 1
		packages.PackageID = 1301
	} else {
		installMAReq.CreatePseudoClientRequest.ClientInfo.ClientType = 0
		packages.PackageID = 51
	}
	installMAReq.Packages = append(installMAReq.Packages, packages)
	InstallMAResp := handler.InstallMA(installMAReq, d.Get("company_id").(int))
	errorCode := InstallMAResp.Response.ErrorCode
	if errorCode != 0 {
		return fmt.Errorf("Installing MA failed")
	}
	var Waittime = 0
	jobId := InstallMAResp.JobID
	for Waittime < 30 {
		jobResponse := handler.JobStatus(strconv.Itoa(jobId))
		if jobResponse.Jobs[0].JobSummary.Status == "Completed" {
			break
		} else if jobResponse.Jobs[0].JobSummary.Status == "Peding" {
			Waittime += 1
		} else if jobResponse.Jobs[0].JobSummary.Status == "Running" {
			Waittime = 0
		} else if jobResponse.Jobs[0].JobSummary.Status == "Waiting" {
			Waittime = 0
		} else {
			return fmt.Errorf("Install MA Job Failed / Killed")
		}
		time.Sleep(1 * time.Minute)
	}
	if Waittime > 30 {
		return fmt.Errorf("Install MA Job is in Pending State since more than 30 minutes")
	}
	clientDetailsResp := handler.GetClientID(d.Get("mediaagent_name").(string))
	if clientDetailsResp.ClientID > 0 {
		clientID := strconv.Itoa(clientDetailsResp.ClientID)
		d.SetId(clientID)
		return resourceInstallMARead(d, m)
	}
	return fmt.Errorf("Unable to get client ID")
}

func resourceInstallMARead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInstallMAUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInstallMADelete(d *schema.ResourceData, m interface{}) error {
	clientID := d.Id()
	uninstalResp := handler.UninstallMA(clientID)
	if uninstalResp.Response.ErrorCode != 0 {
		return fmt.Errorf("Error in uninstallation of Client")
	}
	var Waittime = 0
	jobId := uninstalResp.JobID
	for Waittime < 30 {
		jobResponse := handler.JobStatus(strconv.Itoa(jobId))
		if jobResponse.Jobs[0].JobSummary.Status == "Completed" {
			break
		} else if jobResponse.Jobs[0].JobSummary.Status == "Peding" {
			Waittime += 1
		} else if jobResponse.Jobs[0].JobSummary.Status == "Running" {
			Waittime = 0
		} else if jobResponse.Jobs[0].JobSummary.Status == "Waiting" {
			Waittime = 0
		} else {
			return fmt.Errorf("Un-Install MA Job Failed / Killed")
		}
		time.Sleep(1 * time.Minute)
	}
	if Waittime > 30 {
		return fmt.Errorf("Un-Install MA Job is in Pending State since more than 30 minutes")
	}
	d.SetId("")
	return nil
}
