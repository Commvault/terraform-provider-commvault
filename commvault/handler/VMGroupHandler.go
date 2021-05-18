package handler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func VMGroupCreate(vmgroupName string, planID int, clientid int, vms []string, tags []string, companyID int) *AppCreateSubClientResponse {
	var subclientReq AppCreateSubClientRequest
	subclientReq.SubClientProperties.SubClientEntity.SubclientName = vmgroupName
	subclientReq.SubClientProperties.SubClientEntity.ClientID = clientid
	subclientReq.SubClientProperties.PlanEntity.PlanID = planID
	subclientReq.SubClientProperties.SubClientEntity.ApplicationID = 106
	subclientReq.SubClientProperties.VMContentOperationType = 2
	subclientReq.SubClientProperties.VsaSubclientProp.AutoDetectVMOwner = false
	subclientReq.SubClientProperties.VsaSubclientProp.QuiesceGuestFileSystemAndApplications = true
	subClientReqXML, _ := json.Marshal(subclientReq)
	url := os.Getenv("CV_CSIP") + "/subclient"
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodPost, XML, subClientReqXML, JSON, token, companyID)
	var subClientResp AppCreateSubClientResponse
	xml.Unmarshal(respBody, &subClientResp)

	urlc := os.Getenv("CV_CSIP") + "/Client/" + strconv.Itoa(clientid)
	tokenc := os.Getenv("AuthToken")
	respBodyC := makeHttpRequest(urlc, http.MethodGet, JSON, nil, JSON, tokenc, companyID)
	var clientDetailsResp GetClientDetailsResp
	json.Unmarshal(respBodyC, &clientDetailsResp)

	var vsaType int
	vsaType = clientDetailsResp.ClientProperties[0].PseudoClientInfo.VirtualServerClientProperties.VirtualServerInstanceInfo.VsInstanceType

	var updateSubClient AppUpdateSubClientPropertiesRequest
	updateSubClient.SubClientProperties.VmContentOperationType = "1"
	updateSubClient.SubClientProperties.VmDiskFilterOperationType = "1"
	updateSubClient.SubClientProperties.VmFilterOperationType = "1"
	if len(vms) != 0 {
		for i := range vms {
			var child Children
			child.Name = ""
			child.Type = "10"
			child.AllOrAnyChildren = "1"
			child.DisplayName = vms[i]
			child.EqualsOrNotEquals = "1"
			child.GuestCredentialAssocId = "0"
			updateSubClient.SubClientProperties.VmContent.Children = append(updateSubClient.SubClientProperties.VmContent.Children, child)
		}
	}
	if len(tags) != 0 {
		for i := range tags {
			var child Children
			child.Type = "Tag"
			child.AllOrAnyChildren = "1"
			child.EqualsOrNotEquals = "1"
			child.GuestCredentialAssocId = "0"
			if vsaType == 4 {
				child.Name = tags[i]
				index := strings.LastIndex(tags[i], "=")
				child.DisplayName = tags[i][index+1:]
			} else {
				child.DisplayName = tags[i]
			}
			updateSubClient.SubClientProperties.VmContent.Children = append(updateSubClient.SubClientProperties.VmContent.Children, child)
		}
	}
	if subClientResp.Response.ErrorCode != "0" {
		return &subClientResp
	}
	id := subClientResp.Response.Entity.SubclientId
	url = os.Getenv("CV_CSIP") + "/subclient/" + id
	token = os.Getenv("AuthToken")
	updateSubClientXML, _ := xml.Marshal(updateSubClient)
	makeHttpRequest(url, http.MethodPost, XML, updateSubClientXML, XML, token, companyID)
	return &subClientResp
}

func VMGroupDelete(vmgroupId string) *AppDeleteSubClientResponse {
	url := os.Getenv("CV_CSIP") + "/subclient/" + vmgroupId
	token := os.Getenv("AuthToken")
	respBody := makeHttpRequest(url, http.MethodDelete, XML, nil, XML, token, 0)
	var subClientResp AppDeleteSubClientResponse
	xml.Unmarshal(respBody, &subClientResp)
	return &subClientResp
}

func TestVMS(vms []string) {
	var updateSubClient AppUpdateSubClientPropertiesRequest
	updateSubClient.SubClientProperties.VmContentOperationType = "1"
	updateSubClient.SubClientProperties.VmDiskFilterOperationType = "1"
	updateSubClient.SubClientProperties.VmFilterOperationType = "1"
	for i := range vms {
		var child Children
		child.Name = ""
		child.Type = "10"
		child.AllOrAnyChildren = "1"
		child.DisplayName = vms[i]
		child.EqualsOrNotEquals = "1"
		child.GuestCredentialAssocId = "0"
		updateSubClient.SubClientProperties.VmContent.Children = append(updateSubClient.SubClientProperties.VmContent.Children, child)
	}
	updateSubClientXML, _ := xml.Marshal(updateSubClient)
	updateSubClient.SubClientProperties.VmContent.Children = make([]Children, len(vms))
	fmt.Println(updateSubClientXML)
}
