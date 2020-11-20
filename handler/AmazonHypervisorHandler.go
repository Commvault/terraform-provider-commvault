package handler

import (
	"encoding/xml"
	"net/http"
	"os"
)

func AmazonHypCreateHandler(clientName string, regions string, useIAMrole bool, accessKey string, secretKey string, accessNodes string) *AppCreatePseudoClientResponse {
	var appCreatePseudoClientReq AppCreateAmazonPseudoClientRequest
	var amazonInstanceInfo AmazonInstanceInfo
	var appTypes1 AppTypes
	var appTypes2 AppTypes
	appCreatePseudoClientReq.Entity.ClientName = clientName
	appCreatePseudoClientReq.ClientInfo.ClientType = "12"
	appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.VirtualServerInstanceInfo.VsInstanceType = "4"
	if useIAMrole {
		amazonInstanceInfo.AccessKey = ""
		amazonInstanceInfo.Secretkey = ""
		amazonInstanceInfo.RegionEndPoints = "default"
		amazonInstanceInfo.UseIamRole = "1"
	} else {
		amazonInstanceInfo.AccessKey = accessKey
		amazonInstanceInfo.Secretkey = secretKey
		amazonInstanceInfo.RegionEndPoints = "default"
		amazonInstanceInfo.UseIamRole = "0"
	}
	appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.VirtualServerInstanceInfo.AmazonInstanceInfo = amazonInstanceInfo
	appTypes1.ApplicationID = "134"
	appTypes2.ApplicationID = "106"
	appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.VirtualServerInstanceInfo.VmwareVendor.VcenterHostName = "default"
	appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.AppTypes = append(appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.AppTypes, appTypes1)
	appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.AppTypes = append(appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.AppTypes, appTypes2)
	appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.VirtualServerInstanceInfo.AssociatedClients.MemberServers.Client.ClientName = accessNodes
	appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.VirtualServerInstanceInfo.AssociatedClients.MemberServers.Client.Type = "3"
	url := os.Getenv("CV_CSIP") + "/Client"
	token := os.Getenv("AuthToken")
	psuedoClientXML, _ := xml.Marshal(appCreatePseudoClientReq)
	respBody := makeHttpRequest(url, http.MethodPost, XML, psuedoClientXML, XML, token)
	var psuedoClientResp AppCreatePseudoClientResponse
	xml.Unmarshal(respBody, &psuedoClientResp)
	return &psuedoClientResp
}

func AmazonHypDeleteHandler(clientId string) *AppRetireClientResponse {
	url := os.Getenv("CV_CSIP") + "/Client/" + clientId + "/Retire"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodDelete, XML, nil, XML, token)
	var reriteClientResp AppRetireClientResponse
	xml.Unmarshal(respBody, &reriteClientResp)
	return &reriteClientResp
}
