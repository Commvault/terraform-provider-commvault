package handler

type CreateCompanyReq struct {
	OrganizationInfo struct {
		Organization struct {
			ShortName struct {
				DomainName string `json:"domainName"`
			} `json:"shortName"`
			ConnectName      string   `json:"connectName"`
			EmailDomainNames []string `json:"emailDomainNames"`
		} `json:"organization"`
		PlanDetails            []PlanDetails `json:"planDetails"`
		OrganizationProperties struct {
			PrimaryDomain       string            `json:"primaryDomain"`
			PrimaryContacts     []PrimaryContacts `json:"primaryContacts"`
			EnableAutoDiscovery bool              `json:"enableAutoDiscovery"`
		} `json:"organizationProperties"`
	} `json:"organizationInfo"`
	SendEmail bool `json:"sendEmail"`
}

type PlanDetails struct {
	Plan struct {
		PlanName string `json:"planName"`
	} `json:"plan"`
}

type PrimaryContacts struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}

type ApiCreateCompanyResp struct {
	Processinginstructioninfo struct {
		Attributes []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"attributes"`
	} `json:"processinginstructioninfo"`
	Response struct {
		ErrorString string `json:"errorString"`
		ErrorCode   int    `json:"errorCode"`
		Entity      struct {
			GUID               string `json:"GUID"`
			Type               int    `json:"_type_"`
			ProviderID         int    `json:"providerId"`
			ProviderDomainName string `json:"providerDomainName"`
		} `json:"entity"`
	} `json:"response"`
}

type DeactivateCompany struct {
	DeactivateOptions struct {
		DisableBackup  bool `json:"disableBackup"`
		DisableRestore bool `json:"disableRestore"`
		DisableLogin   bool `json:"disableLogin"`
	} `json:"deactivateOptions"`
}
