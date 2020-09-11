package handler

import "encoding/xml"

type ApiCreatePlanReq struct {
	PlanName           string               `json:"planName"`
	BackupDestinations []BackupDestinations `json:"backupDestinations"`
}

type BackupDestinations struct {
	BackupDestinationName string `json:"backupDestinationName"`
	RetentionPeriodDays   int64  `json:"retentionPeriodDays"`
	StoragePool           struct {
		Name string `json:"name"`
	} `json:"storagePool"`
}

type ApiCreatePlanResp struct {
	Plan struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"paln"`
}

type ApiUpdatePlanReq struct {
	XMLName  xml.Name `xml:"Api_UpdatePlanReq"`
	Database struct {
		SlaInMinutes string `xml:"slaInMinutes,attr"`
		RpoInMinutes string `xml:"rpoInMinutes,attr"`
	} `xml:"database"`
	Summary struct {
		Addons struct {
			Database string `xml:"database,attr"`
		} `xml:"addons"`
		Plan struct {
			PlanId string `xml:"planId,attr"`
		} `xml:"plan"`
	} `xml:"summary"`
}

type GenericResp struct {
	ErrorString string `json:"errorString"`
	ErrorCode   int    `json:"errorCode"`
}

type AppAssociateVMToPlanRequest struct {
	XMLName xml.Name `xml:"App_AssociateVMToPlanRequest"`
	VmInfo  struct {
		Plan struct {
			PlanSubtype string `xml:"planSubtype,attr"`
			PlanType    string `xml:"planType,attr"`
			PlanSummary string `xml:"planSummary,attr"`
			PlanName    string `xml:"planName,attr"`
			PlanId      string `xml:"planId,attr"`
		} `xml:"plan"`
		VmClients struct {
			ClientId   string `xml:"clientId,attr"`
			ClientName string `xml:"clientName,attr"`
			ClientGUID string `xml:"clientGUID,attr"`
		} `xml:"vmClients"`
	} `xml:"vmInfo"`
}

type AppAssociateVMToPlanResponse struct {
	XMLName  xml.Name `xml:"App_AssociateVMToPlanResponse"`
	Response struct {
		ErrorString string `xml:"errorString,attr"`
		ErrorCode   string `xml:"errorCode,attr"`
	} `xml:"response"`
}
