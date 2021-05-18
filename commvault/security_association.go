package commvault

import (
	"encoding/json"
	"fmt"

	"terraform-provider-commvault/commvault/handler"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func securityAssociation() *schema.Resource {
	return &schema.Resource{
		Create: securityAssociationCreate,
		Read:   securityAssociationRead,
		Update: securityAssociationUpdate,
		Delete: securityAssociationDelete,

		Schema: map[string]*schema.Schema{
			"client_list": &schema.Schema{
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Specifies the list of clients for association.",
			},
			"user_group_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the user group name used for association.",
			},
			"permissions_list": &schema.Schema{
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Specifies the permission names list used for the association.",
			},
		},
	}
}

func securityAssociationCreate(d *schema.ResourceData, m interface{}) error {
	var securityAssociations handler.SecurityAssociations

	clients := d.Get("client_list").(*schema.Set).List()
	clientList := make([]string, len(clients))
	for i, n := range clients {
		clientList[i] = n.(string)
	}

	clientIdCount := 0
	if len(clientList) != 0 {
		for i := range clientList {

			clientDetailsResp := handler.GetClientID(clientList[i])
			if clientDetailsResp.ClientID > 0 {
				var entity handler.Entity
				entity.ClientID = clientDetailsResp.ClientID
				securityAssociations.EntityAssociated.Entity = append(securityAssociations.EntityAssociated.Entity, entity)
				clientIdCount++
			} else {
				return fmt.Errorf("Invalid client id for client: " + clientList[i])
			}
		}
	} else {
		securityAssociations.EntityAssociated.Entity = make([]handler.Entity, 0)
	}
	if clientIdCount == 0 {
		securityAssociations.EntityAssociated.Entity = make([]handler.Entity, 0)
	}

	permissions := d.Get("permissions_list").(*schema.Set).List()
	permissionsList := make([]string, len(permissions))
	for i, n := range permissions {
		permissionsList[i] = n.(string)
	}

	var associations handler.Associations
	var userorgroup handler.UserOrGroup
	userorgroup.UshhherGroupName = d.Get("user_group_name").(string)
	associations.UserOrGroup = append(associations.UserOrGroup, userorgroup)
	associations.Properties.IsCreatorAssociation = false

	if len(permissionsList) != 0 {
		for i := range permissionsList {
			var categoriesPermissionList handler.CategoriesPermissionList
			categoriesPermissionList.PermissionName = permissionsList[i]
			associations.Properties.CategoryPermission.CategoriesPermissionList = append(associations.Properties.CategoryPermission.CategoriesPermissionList, categoriesPermissionList)
		}

	} else {
		associations.Properties.CategoryPermission.CategoriesPermissionList = make([]handler.CategoriesPermissionList, 0)
	}

	securityAssociations.SecurityAssociations.AssociationsOperationType = "ADD"
	securityAssociations.SecurityAssociations.Associations = append(securityAssociations.SecurityAssociations.Associations, associations)
	repoBody := handler.SecurityAssociation(securityAssociations)
	var apiAssociateResp handler.ApiAssociateResp
	err := json.Unmarshal(repoBody, &apiAssociateResp)
	if err != nil {
		return fmt.Errorf("Security association failed")
	} else {
		if apiAssociateResp.ErrorCode > 0 {
			return fmt.Errorf("Security association failed with Error:" + apiAssociateResp.ErrorMessage)
		}
		if apiAssociateResp.Response[0].ErrorCode == 0 {
			associationID := d.Get("user_group_name").(string) + clientList[0] + permissionsList[0]
			d.SetId(associationID)
			return securityAssociationRead(d, m)
		} else {
			return fmt.Errorf("Security association failed")
		}
	}
}

func securityAssociationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func securityAssociationUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func securityAssociationDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
