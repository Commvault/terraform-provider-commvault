package handler

type SecurityAssociations struct {
	EntityAssociated struct {
		Entity []Entity `json:"entity"`
	} `json:"entityAssociated"`
	SecurityAssociations struct {
		AssociationsOperationType string         `json:"associationsOperationType"`
		Associations              []Associations `json:"associations"`
	} `json:"securityAssociations"`
}

type Entity struct {
	ClientID int `json:"clientId"`
}

type Associations struct {
	UserOrGroup []UserOrGroup `json:"userOrGroup"`
	Properties  struct {
		IsCreatorAssociation bool `json:"isCreatorAssociation"`
		CategoryPermission   struct {
			CategoriesPermissionList []CategoriesPermissionList `json:"categoriesPermissionList"`
		} `json:"categoryPermission"`
	} `json:"properties"`
}

type UserOrGroup struct {
	UshhherGroupName string `json:"userGroupName"`
}

type CategoriesPermissionList struct {
	PermissionName string `json:"permissionName"`
}

type SecurityAssociationsV2 struct {
	EntityAssociated struct {
		Entity [1]EntityV2 `json:"entity"`
	} `json:"entityAssociated"`
	SecurityAssociations struct {
		AssociationsOperationType string            `json:"associationsOperationType"`
		Associations              [1]AssociationsV2 `json:"associations"`
	} `json:"securityAssociations"`
}

type EntityV2 struct {
	EntityID       int    `json:"entityId,omitempty"`
	EntityName     string `json:"entityName,omitempty"`
	EntityType     int    `json:"entityType,omitempty"`
	EntityTypeName string `json:"entityTypeName,omitempty"`
	Type           int    `json:"_type_,omitempty"`
}

type AssociationsV2 struct {
	UserOrGroup [1]UserOrGroupV2 `json:"userOrGroup"`
	Properties  struct {
		Role struct {
			RoleId   int    `json:"roleId,omitempty"`
			RoleName string `json:"roleName,omitempty"`
		} `json:"role"`
	} `json:"properties"`
}

type UserOrGroupV2 struct {
	UserGroupId   int    `json:"userGroupId,omitempty"`
	UserGroupName string `json:"userGroupName,omitempty"`
}

type ApiAssociateResp struct {
	Processinginstructioninfo struct {
		Attributes []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"attributes"`
	} `json:"processinginstructioninfo"`
	Response []struct {
		WarningCode    int    `json:"warningCode"`
		ErrorCode      int    `json:"errorCode"`
		WarningMessage string `json:"warningMessage"`
	} `json:"response"`
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    int    `json:"errorCode"`
}

type SecurityAssociationsResp struct {
	SecurityAssociations []struct {
		EntityAssociated struct {
			ClientGroupId int `json:"clientGroupId"`
		} `json:"entityAssociated"`
		SecurityAssociations struct {
			Associations []SecurityAssociationAssociations `json:"associations"`
		} `json:"securityAssociations"`
	} `json:"securityAssociations"`
}

type SecurityAssociationAssociations struct {
	UserOrGroup []SecurityUserOrGroup `json:"userOrGroup"`
	Properties  struct {
		Role struct {
			RoleId   int    `json:"roleId"`
			RoleName string `json:"roleName"`
		} `json:"role"`
	} `json:"properties"`
}

type SecurityUserOrGroup struct {
	UserGroupId        int    `json:"userGroupId"`
	UserGroupName      string `json:"userGroupName"`
	GroupId            int    `json:"groupId"`
	ExternalGroupName  string `json:"externalGroupName"`
	Type               int    `json:"_type_"`
	ProviderDomainName string `json:"providerDomainName"`
}
