package handler

import (
	"encoding/xml"
	"net/http"
	"os"
)

func VMWareHypCreateHandler(displayName string, hostName string, userName string, password string, accessNodes string, companyID int) *AppCreatePseudoClientResponse {
	var appCreatePseudoClientReq AppCreatePseudoClientRequest
	appCreatePseudoClientReq.Entity.ClientName = displayName
	appCreatePseudoClientReq.ClientInfo.ClientType = "12"
	appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.VirtualServerInstanceInfo.VsInstanceType = "1"
	appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.VirtualServerInstanceInfo.VmwareVendor.VcenterHostName = hostName
	appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.VirtualServerInstanceInfo.VmwareVendor.VirtualCenter.UserName = userName
	appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.VirtualServerInstanceInfo.VmwareVendor.VirtualCenter.Password = password
	appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.VirtualServerInstanceInfo.AssociatedClients.MemberServers.Client.ClientName = accessNodes
	appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.VirtualServerInstanceInfo.AssociatedClients.MemberServers.Client.Type = "3"
	url := os.Getenv("CV_CSIP") + "/Client"
	token := os.Getenv("AuthToken")
	psuedoClientXML, _ := xml.Marshal(appCreatePseudoClientReq)
	respBody := makeHttpRequest(url, http.MethodPost, XML, psuedoClientXML, XML, token, companyID)
	var psuedoClientResp AppCreatePseudoClientResponse
	xml.Unmarshal(respBody, &psuedoClientResp)
	return &psuedoClientResp
}

func VMWareHypDeleteHandler(clientId string) *AppRetireClientResponse {
	url := os.Getenv("CV_CSIP") + "/Client/" + clientId + "/Retire"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodDelete, XML, nil, XML, token, 0)
	var reriteClientResp AppRetireClientResponse
	xml.Unmarshal(respBody, &reriteClientResp)
	return &reriteClientResp
}
