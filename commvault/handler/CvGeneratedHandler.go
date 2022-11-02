package handler

import (
    "encoding/json"
    "net/http"
    "os"
)

func CvCreateUser(createUserRequest MsgCreateUserRequest) (*MsgCreateUserResponse, error) {
    //API: (POST) /V4/user
    reqBody, _ := json.Marshal(createUserRequest)
    url := os.Getenv("CV_CSIP") + "/V4/user"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateUserResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetUserDetails(userId string) (*MsgGetUserDetailsResponse, error) {
    //API: (GET) /V4/user/{userId}
    url := os.Getenv("CV_CSIP") + "/V4/user/" + userId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetUserDetailsResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteUser(userId string) (*MsgDeleteUserResponse, error) {
    //API: (DELETE) /V4/user/{userId}
    url := os.Getenv("CV_CSIP") + "/V4/user/" + userId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteUserResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvModifyUser(modifyUserRequest MsgModifyUserRequest, userId string) (*MsgModifyUserResponse, error) {
    //API: (PUT) /V4/user/{userId}
    reqBody, _ := json.Marshal(modifyUserRequest)
    url := os.Getenv("CV_CSIP") + "/V4/user/" + userId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgModifyUserResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateHypervisorAWS(createHypervisorAWSRequest MsgCreateHypervisorAWSRequest) (*MsgCreateHypervisorAWSResponse, error) {
    //API: (POST) /V4/Hypervisor
    reqBody, _ := json.Marshal(createHypervisorAWSRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Hypervisor"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateHypervisorAWSResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvupdateHypervisorAWS(updateHypervisorAWSRequest MsgupdateHypervisorAWSRequest, hypervisorId string) (*MsgupdateHypervisorAWSResponse, error) {
    //API: (PUT) /V4/Hypervisor/{hypervisorId}
    reqBody, _ := json.Marshal(updateHypervisorAWSRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Hypervisor/" + hypervisorId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgupdateHypervisorAWSResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetHypervisors(hypervisorId string) (*MsgGetHypervisorsResponse, error) {
    //API: (GET) /V4/Hypervisor/{hypervisorId}
    url := os.Getenv("CV_CSIP") + "/V4/Hypervisor/" + hypervisorId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetHypervisorsResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteHypervisor(hypervisorId string) (*MsgDeleteHypervisorResponse, error) {
    //API: (DELETE) /V4/Hypervisor/{hypervisorId}
    url := os.Getenv("CV_CSIP") + "/V4/Hypervisor/" + hypervisorId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteHypervisorResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateHypervisorAzure(createHypervisorAzureRequest MsgCreateHypervisorAzureRequest) (*MsgCreateHypervisorAzureResponse, error) {
    //API: (POST) /V4/Hypervisor
    reqBody, _ := json.Marshal(createHypervisorAzureRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Hypervisor"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateHypervisorAzureResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvupdateHypervisorAzure(updateHypervisorAzureRequest MsgupdateHypervisorAzureRequest, hypervisorId string) (*MsgupdateHypervisorAzureResponse, error) {
    //API: (PUT) /V4/Hypervisor/{hypervisorId}
    reqBody, _ := json.Marshal(updateHypervisorAzureRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Hypervisor/" + hypervisorId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgupdateHypervisorAzureResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetVMGroup(VmGroupId string) (*MsgGetVMGroupResponse, error) {
    //API: (GET) /V4/VmGroup/{VmGroupId}
    url := os.Getenv("CV_CSIP") + "/V4/VmGroup/" + VmGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetVMGroupResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteVMGroup(VmGroupId string) (*MsgDeleteVMGroupResponse, error) {
    //API: (DELETE) /V4/VmGroup/{VmGroupId}
    url := os.Getenv("CV_CSIP") + "/V4/VmGroup/" + VmGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteVMGroupResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvUpdateVMGroup(updateVMGroupRequest MsgUpdateVMGroupRequest, VmGroupId string) (*MsgUpdateVMGroupResponse, error) {
    //API: (PUT) /V4/VmGroup/{VmGroupId}
    reqBody, _ := json.Marshal(updateVMGroupRequest)
    url := os.Getenv("CV_CSIP") + "/V4/VmGroup/" + VmGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgUpdateVMGroupResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateVMGroup(createVMGroupRequest MsgCreateVMGroupRequest) (*MsgCreateVMGroupResponse, error) {
    //API: (POST) /V4/VMGroup
    reqBody, _ := json.Marshal(createVMGroupRequest)
    url := os.Getenv("CV_CSIP") + "/V4/VMGroup"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateVMGroupResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetPlanById(planId string) (*MsgGetPlanByIdResponse, error) {
    //API: (GET) /V4/ServerPlan/{planId}
    url := os.Getenv("CV_CSIP") + "/V4/ServerPlan/" + planId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetPlanByIdResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeletePlan(planId string) (*MsgDeletePlanResponse, error) {
    //API: (DELETE) /V4/ServerPlan/{planId}
    url := os.Getenv("CV_CSIP") + "/V4/ServerPlan/" + planId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeletePlanResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvModifyPlan(modifyPlanRequest MsgModifyPlanRequest, planId string) (*MsgModifyPlanResponse, error) {
    //API: (PUT) /V4/ServerPlan/{planId}
    reqBody, _ := json.Marshal(modifyPlanRequest)
    url := os.Getenv("CV_CSIP") + "/V4/ServerPlan/" + planId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgModifyPlanResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateServerPlan(createServerPlanRequest MsgCreateServerPlanRequest) (*MsgCreateServerPlanResponse, error) {
    //API: (POST) /V4/ServerPlan
    reqBody, _ := json.Marshal(createServerPlanRequest)
    url := os.Getenv("CV_CSIP") + "/V4/ServerPlan"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateServerPlanResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateUserGroup(createUserGroupRequest MsgCreateUserGroupRequest) (*MsgCreateUserGroupResponse, error) {
    //API: (POST) /V4/UserGroup
    reqBody, _ := json.Marshal(createUserGroupRequest)
    url := os.Getenv("CV_CSIP") + "/V4/UserGroup"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateUserGroupResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetUserGroupDetails(userGroupId string) (*MsgGetUserGroupDetailsResponse, error) {
    //API: (GET) /V4/UserGroup/{userGroupId}
    url := os.Getenv("CV_CSIP") + "/V4/UserGroup/" + userGroupId + "?additionalproperties=true"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetUserGroupDetailsResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteUserGroup(userGroupId string) (*MsgDeleteUserGroupResponse, error) {
    //API: (DELETE) /V4/UserGroup/{userGroupId}
    url := os.Getenv("CV_CSIP") + "/V4/UserGroup/" + userGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteUserGroupResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvModifyUserGroup(modifyUserGroupRequest MsgModifyUserGroupRequest, userGroupId string) (*MsgModifyUserGroupResponse, error) {
    //API: (PUT) /V4/UserGroup/{userGroupId}
    reqBody, _ := json.Marshal(modifyUserGroupRequest)
    url := os.Getenv("CV_CSIP") + "/V4/UserGroup/" + userGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgModifyUserGroupResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}
