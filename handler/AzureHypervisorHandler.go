package handler

import (
	"encoding/xml"
	"net/http"
	"os"
)

func AzureHypCreateHandler(displayName string, subscriptionID string, tenantID string, applicationID string, applicationPWD string, accessNodes string) *AppCreatePseudoClientResponse {
	var appCreatePseudoClientReq AppCreateAzurePseudoClientRequest
	var azureResourceManager AzureResourceManagerInfo
	appCreatePseudoClientReq.Entity.ClientName = displayName
	appCreatePseudoClientReq.ClientInfo.ClientType = "12"
	appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.VirtualServerInstanceInfo.VsInstanceType = "AZURE_V2"
	azureResourceManager.Credentials.UserName = applicationID
	azureResourceManager.Credentials.Password = applicationPWD
	azureResourceManager.ServerName = displayName
	azureResourceManager.SubscriptionID = subscriptionID
	azureResourceManager.TenantID = tenantID
	appCreatePseudoClientReq.ClientInfo.VirtualServerClientProperties.VirtualServerInstanceInfo.AzureResourceManager = azureResourceManager
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

func AzureHypDeleteHandler(clientId string) *AppRetireClientResponse {
	url := os.Getenv("CV_CSIP") + "/Client/" + clientId + "/Retire"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodDelete, XML, nil, XML, token)
	var reriteClientResp AppRetireClientResponse
	xml.Unmarshal(respBody, &reriteClientResp)
	return &reriteClientResp
}
