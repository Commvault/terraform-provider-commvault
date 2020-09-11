package handler

import "encoding/xml"

type AppCreatePseudoClientRequest struct {
	XMLName    xml.Name `xml:"App_CreatePseudoClientRequest"`
	Text       string   `xml:",chardata"`
	ClientInfo struct {
		Text                          string `xml:",chardata"`
		ClientType                    string `xml:"clientType,attr"`
		VirtualServerClientProperties struct {
			Text                      string `xml:",chardata"`
			VirtualServerInstanceInfo struct {
				Text           string `xml:",chardata"`
				VsInstanceType string `xml:"vsInstanceType,attr"`
				VmwareVendor   struct {
					Text            string `xml:",chardata"`
					VcenterHostName string `xml:"vcenterHostName,attr"`
					VirtualCenter   struct {
						Text     string `xml:",chardata"`
						UserName string `xml:"userName,attr"`
						Password string `xml:"password,attr"`
					} `xml:"virtualCenter"`
				} `xml:"vmwareVendor"`
				AssociatedClients struct {
					Text          string `xml:",chardata"`
					MemberServers struct {
						Text   string `xml:",chardata"`
						Client struct {
							Text       string `xml:",chardata"`
							ClientName string `xml:"clientName,attr"`
							ClientId   string `xml:"clientId,attr"`
							Type       string `xml:"_type_,attr"`
						} `xml:"client"`
					} `xml:"memberServers"`
				} `xml:"associatedClients"`
			} `xml:"virtualServerInstanceInfo"`
		} `xml:"virtualServerClientProperties"`
	} `xml:"clientInfo"`
	Entity struct {
		Text       string `xml:",chardata"`
		ClientName string `xml:"clientName,attr"`
	} `xml:"entity"`
}

type AppCreatePseudoClientResponse struct {
	XMLName                   xml.Name `xml:"App_CreatePseudoClientResponse"`
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
			Text       string `xml:",chardata"`
			ClientName string `xml:"clientName,attr"`
			ClientId   string `xml:"clientId,attr"`
			Type       string `xml:"_type_,attr"`
		} `xml:"entity"`
	} `xml:"response"`
}

type AppRetireClientResponse struct {
	XMLName  xml.Name `xml:"App_RetireClientResponse"`
	Text     string   `xml:",chardata"`
	Response struct {
		Text        string `xml:",chardata"`
		ErrorString string `xml:"errorString,attr"`
		ErrorCode   string `xml:"errorCode,attr"`
	} `xml:"response"`
}
