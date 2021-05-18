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
