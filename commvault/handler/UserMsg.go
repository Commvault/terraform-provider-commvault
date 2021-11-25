package handler

type AppCreateUserResponse struct {
	Text         string `xml:",chardata"`
	ErrorMessage string `xml:"errorMessage,attr"`
	Response     struct {
		Text           string `xml:",chardata"`
		ErrorString    string `xml:"errorString,attr"`
		ErrorCode      string `xml:"errorCode,attr"`
		WarningMessage string `xml:"warningMessage,attr"`
		Entity         struct {
			Text     string `xml:",chardata"`
			UserId   string `xml:"userId,attr"`
			UserGUID string `xml:"userGUID,attr"`
			UserName string `xml:"userName,attr"`
		} `xml:"entity"`
	} `xml:"response"`
}

type AppDeleteUserResponse struct {
	Response struct {
		Text           string `xml:",chardata"`
		ErrorString    string `xml:"errorString,attr"`
		ErrorCode      string `xml:"errorCode,attr"`
		WarningMessage string `xml:"warningMessage,attr"`
		Entity         struct {
			Text     string `xml:",chardata"`
			UserId   string `xml:"userId,attr"`
			UserName string `xml:"userName,attr"`
		} `xml:"entity"`
	} `xml:"response"`
}

type AppUpdateUserPropertiesRequest struct {
	Users struct {
		Text                                      string `xml:",chardata"`
		AssociatedUserGroupsOperationType         string `xml:"associatedUserGroupsOperationType,attr"`
		Email                                     string `xml:"email,attr"`
		SystemGeneratePassword                    string `xml:"systemGeneratePassword,attr"`
		AssociatedExternalUserGroupsOperationType string `xml:"associatedExternalUserGroupsOperationType,attr"`
		FullName                                  string `xml:"fullName,attr"`
		EnableUser                                string `xml:"enableUser,attr"`
		UserEntity                                struct {
			Text     string `xml:",chardata"`
			Type     string `xml:"_type_,attr"`
			NewName  string `xml:"newName,attr"`
			UserGUID string `xml:"userGUID,attr"`
			UserName string `xml:"userName,attr"`
			UserId   string `xml:"userId,attr"`
		} `xml:"userEntity"`
	} `xml:"users"`
}

type AppUpdateUserPropertiesResponse struct {
	Text     string `xml:",chardata"`
	Response struct {
		Text           string `xml:",chardata"`
		ErrorString    string `xml:"errorString,attr"`
		ErrorCode      string `xml:"errorCode,attr"`
		WarningMessage string `xml:"warningMessage,attr"`
		Entity         struct {
			Text     string `xml:",chardata"`
			UserId   string `xml:"userId,attr"`
			UserName string `xml:"userName,attr"`
		} `xml:"entity"`
	} `xml:"response"`
}

type UserPropertiesResp struct {
	Users struct {
		Description                        string `json:"description"`
		EnforceFSQuota                     bool   `json:"enforceFSQuota"`
		IdleTime                           int    `json:"idleTime"`
		AgePasswordDays                    int    `json:"agePasswordDays"`
		InheritGroupEdgeDriveQuotaSettings bool   `json:"inheritGroupEdgeDriveQuotaSettings"`
		Email                              string `json:"email"`
		EdgeDriveQuotaLimitInGB            int    `json:"edgeDriveQuotaLimitInGB"`
		LastLogIntime                      int    `json:"lastLogIntime"`
		EnforceEdgeDriveQuota              bool   `json:"enforceEdgeDriveQuota"`
		FullName                           string `json:"fullName"`
		QuotaLimitInGB                     int    `json:"quotaLimitInGB"`
		LoggedInMode                       int    `json:"loggedInMode"`
		InheritGroupQuotaSettings          bool   `json:"inheritGroupQuotaSettings"`
		EnableUser                         bool   `json:"enableUser"`
		APIQuota                           struct {
			APILimit     int `json:"APILimit"`
			APItimeFrame int `json:"APItimeFrame"`
		} `json:"apiQuota"`
		UserEntity struct {
			UserGUID string `json:"userGUID"`
			UserName string `json:"userName"`
			UserID   int    `json:"userId"`
		} `json:"userEntity"`
		LinkedCommvaultUser struct {
		} `json:"LinkedCommvaultUser"`
	} `json:"users"`
}

type AppCreateUserRequest struct {
	Users struct {
		UserEntity struct {
			UserName string `xml:"userName"`
		} `xml:"userEntity"`
		EnableUser      string `xml:"enableUser"`
		AgePasswordDays string `xml:"agePasswordDays"`
		Email           string `xml:"email"`
		Password        string `xml:"password"`
		FullName        string `xml:"fullName"`
		Description     string `xml:"description"`
	} `xml:"users"`
}
