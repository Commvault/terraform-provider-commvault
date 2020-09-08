package handler

import (
    "encoding/xml"
    "fmt"
    "net/http"
    "os"
)

func VMGroupCreate(vmgroupName string,planId string,clientid string,vms []string) *AppCreateSubClientResponse{
    file := readFile("SubClientReq.xml")
    var subclientReq AppCreateSubClientRequest
    err := xml.NewDecoder(file).Decode(&subclientReq)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
    subclientReq.SubClientProperties.SubClientEntity.SubclientName = vmgroupName
    subclientReq.SubClientProperties.SubClientEntity.ClientId = clientid
    subclientReq.SubClientProperties.PlanEntity.PlanId = planId
    subClientReqXML, _ := xml.Marshal(subclientReq)
    url := os.Getenv("CV_CSIP") + "/subclient"
    token := os.Getenv("AuthToken")
    respBody := makeHttpRequest(url,http.MethodPost,XML,subClientReqXML,XML,token)
    var subClientResp AppCreateSubClientResponse
    xml.Unmarshal(respBody,&subClientResp)
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
        updateSubClient.SubClientProperties.VmContent.Children = append(updateSubClient.SubClientProperties.VmContent.Children,child)
    }
    if subClientResp.Response.ErrorCode != "0" {
        return &subClientResp
    }
    id := subClientResp.Response.Entity.SubclientId
    url = os.Getenv("CV_CSIP") + "/subclient/"+id
    token = os.Getenv("AuthToken")
    updateSubClientXML,_ := xml.Marshal(updateSubClient)
    makeHttpRequest(url,http.MethodPost,XML,updateSubClientXML,XML,token)
    return &subClientResp
}


func VMGroupDelete(vmgroupId string) *AppDeleteSubClientResponse{
    url := os.Getenv("CV_CSIP") + "/subclient/"+vmgroupId
    token := os.Getenv("AuthToken")
    respBody := makeHttpRequest(url,http.MethodDelete,XML,nil,XML,token)
    var subClientResp AppDeleteSubClientResponse
    xml.Unmarshal(respBody,&subClientResp)
    return &subClientResp
}

func TestVMS(vms[] string) {
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
        updateSubClient.SubClientProperties.VmContent.Children = append(updateSubClient.SubClientProperties.VmContent.Children,child)
    }
    updateSubClientXML,_ := xml.Marshal(updateSubClient)
    updateSubClient.SubClientProperties.VmContent.Children = make([]Children,len(vms))
    fmt.Println(updateSubClientXML)
}