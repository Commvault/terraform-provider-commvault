package commvault

import (
	"encoding/json"
	"fmt"
	"strconv"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func securityAssociation_v2() *schema.Resource {
	return &schema.Resource{
		Create: securityAssociationCreate_v2,
		Read:   securityAssociationRead_v2,
		Update: securityAssociationUpdate_v2,
		Delete: securityAssociationDelete_v2,

		Schema: map[string]*schema.Schema{
			"entity": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "",
						},
						"typeid": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "",
						},
						"typename": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "",
						},
					},
				},
			},
			"usergroup": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "",
						},
					},
				},
			},
			"role": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func securityAssociationCreate_v2(d *schema.ResourceData, m interface{}) error {
	securityAssociations := buildSecurityPayload(d, m)
	securityAssociations.SecurityAssociations.AssociationsOperationType = "ADD"

	resp, err1 := handler.CreateSecurityAssociationV2(securityAssociations)

	if err1 != nil {
		return fmt.Errorf("operation [Security] failed, Error %s", err1)
	}

	var apiAssociateResp handler.ApiAssociateResp
	err2 := json.Unmarshal(resp, &apiAssociateResp)

	if err2 != nil {
		return fmt.Errorf("security association failed")
	} else {
		if apiAssociateResp.ErrorCode > 0 {
			return fmt.Errorf("security association failed with Error:" + apiAssociateResp.ErrorMessage)
		}
		if apiAssociateResp.Response[0].ErrorCode == 0 {
			//associationID := "v2_" + strconv.Itoa(securityAssociations.EntityAssociated.Entity[0].EntityID) + "_" + strconv.Itoa(securityAssociations.EntityAssociated.Entity[0].EntityType) + "_" + strconv.Itoa(userGroupId) + "_" + strconv.Itoa(roleId)
			d.SetId("V2")
			return securityAssociationRead_v2(d, m)
		} else {
			return fmt.Errorf("security association failed, error code: " + strconv.Itoa(apiAssociateResp.Response[0].ErrorCode))
		}
	}
}

func buildSecurityPayload(d *schema.ResourceData, m interface{}) handler.SecurityAssociationsV2 {
	var securityAssociations handler.SecurityAssociationsV2

	var entityId int
	var entityType int
	var entityName string
	var entityTypeName string
	var userGroupId int
	var userGroupName string
	var roleId int
	var roleName string

	entity := d.Get("entity").([]interface{})
	if len(entity) > 0 && entity[0] != nil {
		tmp := entity[0].(map[string]interface{})
		entityId = tmp["id"].(int)
		entityType = tmp["typeid"].(int)
		entityName = tmp["name"].(string)
		entityTypeName = tmp["typename"].(string)
	}

	usergroup := d.Get("usergroup").([]interface{})
	if len(usergroup) > 0 && usergroup[0] != nil {
		tmp := usergroup[0].(map[string]interface{})
		userGroupId = tmp["id"].(int)
		userGroupName = tmp["name"].(string)
	}

	role := d.Get("role").([]interface{})
	if len(role) > 0 && role[0] != nil {
		tmp := role[0].(map[string]interface{})
		roleId = tmp["id"].(int)
		roleName = tmp["name"].(string)
	}

	//entityId := d.Get("entityid").(int)
	//entityType := d.Get("entitytype").(int)
	//userGroupId := d.Get("usergroupid").(int)
	//roleId := d.Get("roleid").(int)

	securityAssociations.EntityAssociated.Entity[0].EntityID = entityId
	securityAssociations.EntityAssociated.Entity[0].EntityName = entityName
	securityAssociations.EntityAssociated.Entity[0].EntityType = entityType
	securityAssociations.EntityAssociated.Entity[0].EntityTypeName = entityTypeName
	securityAssociations.EntityAssociated.Entity[0].Type = 150
	//securityAssociations.SecurityAssociations.AssociationsOperationType = "ADD"
	securityAssociations.SecurityAssociations.Associations[0].UserOrGroup[0].UserGroupId = userGroupId
	securityAssociations.SecurityAssociations.Associations[0].UserOrGroup[0].UserGroupName = userGroupName
	securityAssociations.SecurityAssociations.Associations[0].Properties.Role.RoleId = roleId
	securityAssociations.SecurityAssociations.Associations[0].Properties.Role.RoleName = roleName

	return securityAssociations
}

func securityAssociationRead_v2(d *schema.ResourceData, m interface{}) error {
	var entityId int
	var entityType int

	entity := d.Get("entity").([]interface{})
	if len(entity) > 0 && entity[0] != nil {
		tmp := entity[0].(map[string]interface{})
		entityId = tmp["id"].(int)
		entityType = tmp["typeid"].(int)
	}

	resp, _ := handler.GetSecurityAssociationV2(entityType, entityId)

	var sec handler.SecurityAssociationsResp
	json.Unmarshal(resp, &sec)

	found := false

	if len(sec.SecurityAssociations) > 0 {
		for _, iter_a := range sec.SecurityAssociations {
			if checkAssociations(iter_a.SecurityAssociations.Associations, d, m) {
				found = true
			}
		}
	}

	if !found {
		handler.LogEntry(">>>>>>>>>>>>>>>>", "MISSING ASSOCIATION, CLEARING")
		clear(d, m)
	}

	return nil
}

func checkAssociations(r []handler.SecurityAssociationAssociations, d *schema.ResourceData, m interface{}) bool {
	var roleId int
	var roleName string

	role := d.Get("role").([]interface{})
	if len(role) > 0 && role[0] != nil {
		tmp := role[0].(map[string]interface{})
		roleId = tmp["id"].(int)
		roleName = tmp["name"].(string)
	}

	found := false

	for _, iter_a := range r {
		if iter_a.Properties.Role.RoleId == roleId || iter_a.Properties.Role.RoleName == roleName {
			if checkAssociatedGroup(iter_a.UserOrGroup, d, m) {
				found = true
			}
		}
	}

	return found
}

func checkAssociatedGroup(r []handler.SecurityUserOrGroup, d *schema.ResourceData, m interface{}) bool {
	var userGroupId int
	var userGroupName string

	usergroup := d.Get("usergroup").([]interface{})
	if len(usergroup) > 0 && usergroup[0] != nil {
		tmp := usergroup[0].(map[string]interface{})
		userGroupId = tmp["id"].(int)
		userGroupName = tmp["name"].(string)
	}

	found := false

	for _, iter_a := range r {
		if iter_a.Type == 62 {
			externalGroup := iter_a.ProviderDomainName + "\\" + iter_a.ExternalGroupName
			if iter_a.GroupId == userGroupId || externalGroup == userGroupName {
				found = true
			}
		} else {
			if iter_a.UserGroupId == userGroupId || iter_a.UserGroupName == userGroupName {
				found = true
			}
		}
	}

	return found
}

func clear(d *schema.ResourceData, m interface{}) {
	d.Set("entity", make([]map[string]interface{}, 0))
	d.Set("entity", make([]map[string]interface{}, 0))
	d.Set("entity", make([]map[string]interface{}, 0))
}

func securityAssociationUpdate_v2(d *schema.ResourceData, m interface{}) error {
	return securityAssociationCreate_v2(d, m)
}

func securityAssociationDelete_v2(d *schema.ResourceData, m interface{}) error {
	securityAssociations := buildSecurityPayload(d, m)
	securityAssociations.SecurityAssociations.AssociationsOperationType = "DELETE"
	handler.DeleteSecurityAssociationV2(securityAssociations)
	return nil
}
