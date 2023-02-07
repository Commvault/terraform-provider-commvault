package handler

import (
    "encoding/json"
    "net/http"
    "os"
)

func CvCreateBackupLocation(createBackupLocationRequest MsgCreateBackupLocationRequest, storagePoolId string) (*MsgCreateBackupLocationResponse, error) {
    //API: (POST) /V4/Storage/Disk/{storagePoolId}/BackupLocation
    reqBody, _ := json.Marshal(createBackupLocationRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Disk/" + storagePoolId + "/BackupLocation"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateBackupLocationResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetBackupLocationDetails(storagePoolId string, backupLocationId string) (*MsgGetBackupLocationDetailsResponse, error) {
    //API: (GET) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Disk/" + storagePoolId + "/BackupLocation/" + backupLocationId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetBackupLocationDetailsResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteBackupLocation(storagePoolId string, backupLocationId string) (*MsgDeleteBackupLocationResponse, error) {
    //API: (DELETE) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Disk/" + storagePoolId + "/BackupLocation/" + backupLocationId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteBackupLocationResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvModifyBackupLocation(modifyBackupLocationRequest MsgModifyBackupLocationRequest, storagePoolId string, backupLocationId string) (*MsgModifyBackupLocationResponse, error) {
    //API: (PUT) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}
    reqBody, _ := json.Marshal(modifyBackupLocationRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Disk/" + storagePoolId + "/BackupLocation/" + backupLocationId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgModifyBackupLocationResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateCloudStorageS3(createCloudStorageS3Request MsgCreateCloudStorageS3Request) (*MsgCreateCloudStorageS3Response, error) {
    //API: (POST) /V4/Storage/Cloud
    reqBody, _ := json.Marshal(createCloudStorageS3Request)
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateCloudStorageS3Response
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetCloudStorageById(cloudStorageId string) (*MsgGetCloudStorageByIdResponse, error) {
    //API: (GET) /V4/Storage/Cloud/{cloudStorageId}
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud/" + cloudStorageId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetCloudStorageByIdResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteCloudStorageById(cloudStorageId string) (*MsgDeleteCloudStorageByIdResponse, error) {
    //API: (DELETE) /V4/Storage/Cloud/{cloudStorageId}
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud/" + cloudStorageId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteCloudStorageByIdResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvModifyCloudStorageById(modifyCloudStorageByIdRequest MsgModifyCloudStorageByIdRequest, cloudStorageId string) (*MsgModifyCloudStorageByIdResponse, error) {
    //API: (PUT) /V4/Storage/Cloud/{cloudStorageId}
    reqBody, _ := json.Marshal(modifyCloudStorageByIdRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud/" + cloudStorageId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgModifyCloudStorageByIdResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetDiskStorageDetails(storagePoolId string) (*MsgGetDiskStorageDetailsResponse, error) {
    //API: (GET) /V4/Storage/Disk/{storagePoolId}
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Disk/" + storagePoolId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetDiskStorageDetailsResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteDiskStorage(storagePoolId string) (*MsgDeleteDiskStorageResponse, error) {
    //API: (DELETE) /V4/Storage/Disk/{storagePoolId}
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Disk/" + storagePoolId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteDiskStorageResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvModifyDiskStorage(modifyDiskStorageRequest MsgModifyDiskStorageRequest, storagePoolId string) (*MsgModifyDiskStorageResponse, error) {
    //API: (PUT) /V4/Storage/Disk/{storagePoolId}
    reqBody, _ := json.Marshal(modifyDiskStorageRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Disk/" + storagePoolId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgModifyDiskStorageResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateDiskStorage(createDiskStorageRequest MsgCreateDiskStorageRequest) (*MsgCreateDiskStorageResponse, error) {
    //API: (POST) /V4/Storage/Disk
    reqBody, _ := json.Marshal(createDiskStorageRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Disk"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateDiskStorageResponse
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
    url := os.Getenv("CV_CSIP") + "/Client/" + hypervisorId + "/Retire"
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

func CvCreateBucketforCloudStorageS3(createBucketforCloudStorageS3Request MsgCreateBucketforCloudStorageS3Request, cloudStorageId string) (*MsgCreateBucketforCloudStorageS3Response, error) {
    //API: (POST) /V4/Storage/Cloud/{cloudStorageId}/Bucket
    reqBody, _ := json.Marshal(createBucketforCloudStorageS3Request)
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud/" + cloudStorageId + "/Bucket"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateBucketforCloudStorageS3Response
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetBucketDetailsOfCloudStorage(cloudStorageId string, bucketId string) (*MsgGetBucketDetailsOfCloudStorageResponse, error) {
    //API: (GET) /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud/" + cloudStorageId + "/Bucket/" + bucketId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetBucketDetailsOfCloudStorageResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteBucketOfCloudStorage(cloudStorageId string, bucketId string) (*MsgDeleteBucketOfCloudStorageResponse, error) {
    //API: (DELETE) /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud/" + cloudStorageId + "/Bucket/" + bucketId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteBucketOfCloudStorageResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvModifyBucketOfCloudStorage(modifyBucketOfCloudStorageRequest MsgModifyBucketOfCloudStorageRequest, cloudStorageId string, bucketId string) (*MsgModifyBucketOfCloudStorageResponse, error) {
    //API: (PUT) /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}
    reqBody, _ := json.Marshal(modifyBucketOfCloudStorageRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud/" + cloudStorageId + "/Bucket/" + bucketId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgModifyBucketOfCloudStorageResponse
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
