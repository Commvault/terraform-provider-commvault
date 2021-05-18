package commvault

import (
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCompany() *schema.Resource {
	return &schema.Resource{
		Create: resourceCompanyCreate,
		Read:   resourceCompanyRead,
		Update: resourceCompanyUpdate,
		Delete: resourceCompanyDelete,

		Schema: map[string]*schema.Schema{
			"company_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"contact_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"plans": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"company_alias": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"associated_smtp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"send_email": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"company_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceCompanyCreate(d *schema.ResourceData, m interface{}) error {
	var createCompanyReq handler.CreateCompanyReq
	createCompanyReq.OrganizationInfo.Organization.ConnectName = d.Get("company_name").(string)
	createCompanyReq.OrganizationInfo.Organization.EmailDomainNames = append(createCompanyReq.OrganizationInfo.Organization.EmailDomainNames, d.Get("associated_smtp").(string))
	var primaryContacts handler.PrimaryContacts
	primaryContacts.FullName = d.Get("contact_name").(string)
	primaryContacts.Email = d.Get("email").(string)
	createCompanyReq.OrganizationInfo.OrganizationProperties.PrimaryContacts = append(createCompanyReq.OrganizationInfo.OrganizationProperties.PrimaryContacts, primaryContacts)
	createCompanyReq.OrganizationInfo.Organization.ShortName.DomainName = d.Get("company_alias").(string)
	createCompanyReq.SendEmail = d.Get("send_email").(bool)
	plans := d.Get("plans").(*schema.Set).List()
	plannames := make([]string, len(plans))
	for i, n := range plans {
		plannames[i] = n.(string)
	}
	if len(plannames) != 0 {
		for i := range plannames {
			var planDetails handler.PlanDetails
			planDetails.Plan.PlanName = plannames[i]
			createCompanyReq.OrganizationInfo.PlanDetails = append(createCompanyReq.OrganizationInfo.PlanDetails, planDetails)
		}
	} else {
		createCompanyReq.OrganizationInfo.PlanDetails = make([]handler.PlanDetails, 0)
	}
	companyResp := handler.CompanyCreate(createCompanyReq, d.Get("company_id").(int))
	errorCode := companyResp.Response.ErrorCode
	if errorCode != 0 {
		return fmt.Errorf("Company creation failed")
	}
	providerID := strconv.Itoa(companyResp.Response.Entity.ProviderID)
	d.SetId(providerID)
	return resourceCompanyRead(d, m)
}

func resourceCompanyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCompanyUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCompanyDelete(d *schema.ResourceData, m interface{}) error {
	providerID := d.Id()
	deactivateResp := handler.CompanyDeactivate(providerID)
	if deactivateResp.Response.ErrorCode != 0 {
		return fmt.Errorf("Error in Deactivatoin of Company")
	}
	genericResp := handler.CompanyDelete(providerID)
	if genericResp.ErrorCode != 0 {
		return fmt.Errorf("Error in Deletion of Company")
	}
	d.SetId("")
	return nil
}
