package handler

import "encoding/xml"

type AppCreateSubClientRequest struct {
	SubClientProperties struct {
		VMContentOperationType int `json:"vmContentOperationType"`
		VMContent              struct {
		} `json:"vmContent"`
		SubClientEntity struct {
			ClientID      int    `json:"clientId"`
			ApplicationID int    `json:"applicationId"`
			SubclientName string `json:"subclientName"`
		} `json:"subClientEntity"`
		PlanEntity struct {
			PlanID int `json:"planId"`
		} `json:"planEntity"`
		CommonProperties struct {
		} `json:"commonProperties"`
		VsaSubclientProp struct {
			AutoDetectVMOwner                     bool `json:"autoDetectVMOwner"`
			QuiesceGuestFileSystemAndApplications bool `json:"quiesceGuestFileSystemAndApplications"`
		} `json:"vsaSubclientProp"`
	} `json:"subClientProperties"`
}

type Children struct {
	AllOrAnyChildren       string `xml:"allOrAnyChildren,attr"`
	Name                   string `xml:"name,attr"`
	Type                   string `xml:"type,attr"`
	Path                   string `xml:"path,attr"`
	DisplayName            string `xml:"displayName,attr"`
	EqualsOrNotEquals      string `xml:"equalsOrNotEquals,attr"`
	GuestCredentialAssocId string `xml:"guestCredentialAssocId,attr"`
}

type AppCreateSubClientResponse struct {
	XMLName                   xml.Name `xml:"App_CreateSubClientResponse"`
	Text                      string   `xml:",chardata"`
	Processinginstructioninfo struct {
		Text       string `xml:",chardata"`
		Attributes struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Value string `xml:"value,attr"`
		} `xml:"attributes"`
	} `xml:"processinginstructioninfo"`
	Response struct {
		Text           string `xml:",chardata"`
		WarningCode    string `xml:"warningCode,attr"`
		ErrorCode      string `xml:"errorCode,attr"`
		WarningMessage string `xml:"warningMessage,attr"`
		Entity         struct {
			Text        string `xml:",chardata"`
			SubclientId string `xml:"subclientId,attr"`
			Type        string `xml:"_type_,attr"`
		} `xml:"entity"`
	} `xml:"response"`
}

type AppDeleteSubClientResponse struct {
	XMLName                   xml.Name `xml:"App_DeleteSubClientResponse"`
	Text                      string   `xml:",chardata"`
	Processinginstructioninfo struct {
		Text       string `xml:",chardata"`
		Attributes struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Value string `xml:"value,attr"`
		} `xml:"attributes"`
	} `xml:"processinginstructioninfo"`
	Response struct {
		Text      string `xml:",chardata"`
		ErrorCode string `xml:"errorCode,attr"`
		Entity    struct {
			Text          string `xml:",chardata"`
			SubclientId   string `xml:"subclientId,attr"`
			ClientName    string `xml:"clientName,attr"`
			ClientId      string `xml:"clientId,attr"`
			SubclientName string `xml:"subclientName,attr"`
			BackupsetName string `xml:"backupsetName,attr"`
			InstanceName  string `xml:"instanceName,attr"`
			Type          string `xml:"_type_,attr"`
			AppName       string `xml:"appName,attr"`
		} `xml:"entity"`
	} `xml:"response"`
}

type AppUpdateSubClientPropertiesRequest struct {
	XMLName             xml.Name `xml:"App_UpdateSubClientPropertiesRequest"`
	Text                string   `xml:",chardata"`
	SubClientProperties struct {
		Text                      string `xml:",chardata"`
		VmContentOperationType    string `xml:"vmContentOperationType,attr"`
		VmFilterOperationType     string `xml:"vmFilterOperationType,attr"`
		VmDiskFilterOperationType string `xml:"vmDiskFilterOperationType,attr"`
		VmContent                 struct {
			Children []Children `xml:"children"`
		} `xml:"vmContent"`
		VmFilter     string `xml:"vmFilter"`
		VmDiskFilter string `xml:"vmDiskFilter"`
	} `xml:"subClientProperties"`
}
