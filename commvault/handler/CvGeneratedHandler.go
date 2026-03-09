package handler

import (
    "encoding/json"
    "net/http"
    "os"
)

func CvCreateAccessPathForBucketOfCloudStorage(createAccessPathForBucketOfCloudStorageRequest MsgCreateAccessPathForBucketOfCloudStorageRequest, cloudStorageId string, bucketId string) (*MsgCreateAccessPathForBucketOfCloudStorageResponse, error) {
    //API: (POST) /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}/AccessPath
    reqBody, _ := json.Marshal(createAccessPathForBucketOfCloudStorageRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud/" + cloudStorageId + "/Bucket/" + bucketId + "/AccessPath"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateAccessPathForBucketOfCloudStorageResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteAccessPathForBucketOfCloudStorage(cloudStorageId string, bucketId string, accessPathId string) (*MsgDeleteAccessPathForBucketOfCloudStorageResponse, error) {
    //API: (DELETE) /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}/AccessPath/{accessPathId}
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud/" + cloudStorageId + "/Bucket/" + bucketId + "/AccessPath/" + accessPathId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteAccessPathForBucketOfCloudStorageResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateCloudStorageAzure(createCloudStorageAzureRequest MsgCreateCloudStorageAzureRequest) (*MsgCreateCloudStorageAzureResponse, error) {
    //API: (POST) /V4/Storage/Cloud
    reqBody, _ := json.Marshal(createCloudStorageAzureRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateCloudStorageAzureResponse
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

func CvCreateServerGroups(createServerGroupsRequest MsgCreateServerGroupsRequest) (*MsgCreateServerGroupsResponse, error) {
    //API: (POST) /V4/ServerGroup
    reqBody, _ := json.Marshal(createServerGroupsRequest)
    url := os.Getenv("CV_CSIP") + "/V4/ServerGroup"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateServerGroupsResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetServerGroupIdDetails(serverGroupId string) (*MsgGetServerGroupIdDetailsResponse, error) {
    //API: (GET) /V4/ServerGroup/{serverGroupId}
    url := os.Getenv("CV_CSIP") + "/V4/ServerGroup/" + serverGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetServerGroupIdDetailsResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteServerGroup(serverGroupId string) (*MsgDeleteServerGroupResponse, error) {
    //API: (DELETE) /V4/ServerGroup/{serverGroupId}
    url := os.Getenv("CV_CSIP") + "/V4/ServerGroup/" + serverGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteServerGroupResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvUpdateServerGroupAssociation(updateServerGroupAssociationRequest MsgUpdateServerGroupAssociationRequest, serverGroupId string) (*MsgUpdateServerGroupAssociationResponse, error) {
    //API: (PUT) /V4/ServerGroup/{serverGroupId}
    reqBody, _ := json.Marshal(updateServerGroupAssociationRequest)
    url := os.Getenv("CV_CSIP") + "/V4/ServerGroup/" + serverGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgUpdateServerGroupAssociationResponse
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

func CvgetProtectionGroupProperties(protectionGroupId string) (*MsggetProtectionGroupPropertiesResponse, error) {
    //API: (GET) /V4/ProtectionGroup/{protectionGroupId}
    url := os.Getenv("CV_CSIP") + "/V4/ProtectionGroup/" + protectionGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsggetProtectionGroupPropertiesResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteWorkloadProtectionGroup(protectionGroupId string) (*MsgDeleteWorkloadProtectionGroupResponse, error) {
    //API: (DELETE) /V4/ProtectionGroup/{protectionGroupId}
    url := os.Getenv("CV_CSIP") + "/V4/ProtectionGroup/" + protectionGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteWorkloadProtectionGroupResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvUpdateWorkloadProtectionGroup(updateWorkloadProtectionGroupRequest MsgUpdateWorkloadProtectionGroupRequest, protectionGroupId string) (*MsgUpdateWorkloadProtectionGroupResponse, error) {
    //API: (PUT) /V4/ProtectionGroup/{protectionGroupId}
    reqBody, _ := json.Marshal(updateWorkloadProtectionGroupRequest)
    url := os.Getenv("CV_CSIP") + "/V4/ProtectionGroup/" + protectionGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgUpdateWorkloadProtectionGroupResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateProtectionGroup(createProtectionGroupRequest MsgCreateProtectionGroupRequest) (*MsgCreateProtectionGroupResponse, error) {
    //API: (POST) /V4/ProtectionGroup/Azure
    reqBody, _ := json.Marshal(createProtectionGroupRequest)
    url := os.Getenv("CV_CSIP") + "/V4/ProtectionGroup/Azure"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateProtectionGroupResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateKubernetesClusterOp(createKubernetesClusterOpRequest MsgCreateKubernetesClusterOpRequest) (*MsgCreateKubernetesClusterOpResponse, error) {
    //API: (POST) /V5/Kubernetes/Cluster
    reqBody, _ := json.Marshal(createKubernetesClusterOpRequest)
    url := os.Getenv("CV_CSIP") + "/V5/Kubernetes/Cluster"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateKubernetesClusterOpResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetKubernetesClusterDetails(clusterId string) (*MsgGetKubernetesClusterDetailsResponse, error) {
    //API: (GET) /V5/Kubernetes/Cluster/{clusterId}
    url := os.Getenv("CV_CSIP") + "/V5/Kubernetes/Cluster/" + clusterId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetKubernetesClusterDetailsResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvUpdateKubernetesProperties(updateKubernetesPropertiesRequest MsgUpdateKubernetesPropertiesRequest, clusterId string) (*MsgUpdateKubernetesPropertiesResponse, error) {
    //API: (PUT) /V5/Kubernetes/Cluster/{clusterId}
    reqBody, _ := json.Marshal(updateKubernetesPropertiesRequest)
    url := os.Getenv("CV_CSIP") + "/V5/Kubernetes/Cluster/" + clusterId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgUpdateKubernetesPropertiesResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvRetireKubernetesCluster(clusterId string) (*MsgRetireKubernetesClusterResponse, error) {
    //API: (DELETE) /V5/Kubernetes/Cluster/{clusterId}/Retire
    url := os.Getenv("CV_CSIP") + "/V5/Kubernetes/Cluster/" + clusterId + "/Retire"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgRetireKubernetesClusterResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvgetReplicationGroupDetailsAmazon(replicationGroupId string) (*MsggetReplicationGroupDetailsAmazonRequest, error) {
    //API: (GET) /V4/ReplicationGroup/{replicationGroupId}
    url := os.Getenv("CV_CSIP") + "/V4/ReplicationGroup/" + replicationGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsggetReplicationGroupDetailsAmazonRequest
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateReplicationGroupAmazon(createReplicationGroupAmazonRequest MsgCreateReplicationGroupAmazonRequest) (*MsgCreateReplicationGroupAmazonResponse, error) {
    //API: (POST) /V4/ReplicationGroup
    reqBody, _ := json.Marshal(createReplicationGroupAmazonRequest)
    url := os.Getenv("CV_CSIP") + "/V4/ReplicationGroup"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateReplicationGroupAmazonResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvModifyReplicationGroup(modifyReplicationGroupRequest MsgModifyReplicationGroupRequest, replicationGroupId string) (*MsgModifyReplicationGroupResponse, error) {
    //API: (PUT) /V4/ReplicationGroup/{replicationGroupId}
    reqBody, _ := json.Marshal(modifyReplicationGroupRequest)
    url := os.Getenv("CV_CSIP") + "/V4/ReplicationGroup/" + replicationGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgModifyReplicationGroupResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvgetReplicationGroupDetailsAzure(replicationGroupId string) (*MsggetReplicationGroupDetailsAzureRequest, error) {
    //API: (GET) /V4/ReplicationGroup/{replicationGroupId}
    url := os.Getenv("CV_CSIP") + "/V4/ReplicationGroup/" + replicationGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsggetReplicationGroupDetailsAzureRequest
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateReplicationGroupAzure(createReplicationGroupAzureRequest MsgCreateReplicationGroupAzureRequest) (*MsgCreateReplicationGroupAzureResponse, error) {
    //API: (POST) /V4/ReplicationGroup
    reqBody, _ := json.Marshal(createReplicationGroupAzureRequest)
    url := os.Getenv("CV_CSIP") + "/V4/ReplicationGroup"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateReplicationGroupAzureResponse
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

func CvCreateWorkloadProtectionGroupAWS(createWorkloadProtectionGroupAWSRequest MsgCreateWorkloadProtectionGroupAWSRequest) (*MsgCreateWorkloadProtectionGroupAWSResponse, error) {
    //API: (POST) /V4/ProtectionGroup/AWS
    reqBody, _ := json.Marshal(createWorkloadProtectionGroupAWSRequest)
    url := os.Getenv("CV_CSIP") + "/V4/ProtectionGroup/AWS"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateWorkloadProtectionGroupAWSResponse
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

func CvCreateDCPlan(createDCPlanRequest MsgCreateDCPlanRequest) ([]byte, error) {
    //API: (POST) /V4/DCPlan
    reqBody, _ := json.Marshal(createDCPlanRequest)
    url := os.Getenv("CV_CSIP") + "/V4/DCPlan"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    return respBody, err
}

func CvGetBlackoutWindowDetails(blackoutWindowId string) (*MsgGetBlackoutWindowDetailsResponse, error) {
    //API: (GET) /V5/BlackoutWindow/{blackoutWindowId}
    url := os.Getenv("CV_CSIP") + "/V5/BlackoutWindow/" + blackoutWindowId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetBlackoutWindowDetailsResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvModifyBlackoutWindow(modifyBlackoutWindowRequest MsgModifyBlackoutWindowRequest, blackoutWindowId string) (*MsgModifyBlackoutWindowResponse, error) {
    //API: (PUT) /V5/BlackoutWindow/{blackoutWindowId}
    reqBody, _ := json.Marshal(modifyBlackoutWindowRequest)
    url := os.Getenv("CV_CSIP") + "/V5/BlackoutWindow/" + blackoutWindowId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgModifyBlackoutWindowResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateBlackoutWindow(createBlackoutWindowRequest MsgCreateBlackoutWindowRequest) (*MsgCreateBlackoutWindowResponse, error) {
    //API: (POST) /V5/BlackoutWindow
    reqBody, _ := json.Marshal(createBlackoutWindowRequest)
    url := os.Getenv("CV_CSIP") + "/V5/BlackoutWindow"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateBlackoutWindowResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteBlackoutWindow(blackoutWindowId string) (*MsgDeleteBlackoutWindowResponse, error) {
    //API: (DELETE) /V4/BlackoutWindow/{blackoutWindowId}
    url := os.Getenv("CV_CSIP") + "/V4/BlackoutWindow/" + blackoutWindowId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteBlackoutWindowResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

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

func CvGetCredentialDetailsAzure(credentialId string) (*MsgGetCredentialDetailsAzureRequest, error) {
    //API: (GET) /V5/Credential/{credentialId}
    url := os.Getenv("CV_CSIP") + "/V5/Credential/" + credentialId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetCredentialDetailsAzureRequest
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvUpdateCredentialAzure(updateCredentialAzureRequest MsgUpdateCredentialAzureRequest, credentialId string) (*MsgUpdateCredentialAzureResponse, error) {
    //API: (PUT) /V5/Credential/{credentialId}
    reqBody, _ := json.Marshal(updateCredentialAzureRequest)
    url := os.Getenv("CV_CSIP") + "/V5/Credential/" + credentialId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgUpdateCredentialAzureResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateCredentialAzure(createCredentialAzureRequest MsgCreateCredentialAzureRequest) (*MsgCreateCredentialAzureResponse, error) {
    //API: (POST) /V4/Credential
    reqBody, _ := json.Marshal(createCredentialAzureRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Credential"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateCredentialAzureResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteCredential(credentialId string) (*MsgDeleteCredentialResponse, error) {
    //API: (DELETE) /V5/Credential/{credentialId}
    url := os.Getenv("CV_CSIP") + "/V5/Credential/" + credentialId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteCredentialResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetCredentialDetailsAWS(credentialId string) (*MsgGetCredentialDetailsAWSRequest, error) {
    //API: (GET) /V5/Credential/{credentialId}
    url := os.Getenv("CV_CSIP") + "/V5/Credential/" + credentialId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetCredentialDetailsAWSRequest
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvUpdateCredentialAWS(updateCredentialAWSRequest MsgUpdateCredentialAWSRequest, credentialId string) (*MsgUpdateCredentialAWSResponse, error) {
    //API: (PUT) /V5/Credential/{credentialId}
    reqBody, _ := json.Marshal(updateCredentialAWSRequest)
    url := os.Getenv("CV_CSIP") + "/V5/Credential/" + credentialId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgUpdateCredentialAWSResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateCredentialAWS(createCredentialAWSRequest MsgCreateCredentialAWSRequest) (*MsgCreateCredentialAWSResponse, error) {
    //API: (POST) /V4/Credential
    reqBody, _ := json.Marshal(createCredentialAWSRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Credential"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateCredentialAWSResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetCredentialDetailsAWSWithRoleArn(credentialId string) (*MsgGetCredentialDetailsAWSWithRoleArnRequest, error) {
    //API: (GET) /V5/Credential/{credentialId}
    url := os.Getenv("CV_CSIP") + "/V5/Credential/" + credentialId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetCredentialDetailsAWSWithRoleArnRequest
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvUpdateCredentialAWSWithRoleArn(updateCredentialAWSWithRoleArnRequest MsgUpdateCredentialAWSWithRoleArnRequest, credentialId string) (*MsgUpdateCredentialAWSWithRoleArnResponse, error) {
    //API: (PUT) /V5/Credential/{credentialId}
    reqBody, _ := json.Marshal(updateCredentialAWSWithRoleArnRequest)
    url := os.Getenv("CV_CSIP") + "/V5/Credential/" + credentialId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgUpdateCredentialAWSWithRoleArnResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateCredentialAWSWithRoleArn(createCredentialAWSWithRoleArnRequest MsgCreateCredentialAWSWithRoleArnRequest) (*MsgCreateCredentialAWSWithRoleArnResponse, error) {
    //API: (POST) /V4/Credential
    reqBody, _ := json.Marshal(createCredentialAWSWithRoleArnRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Credential"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateCredentialAWSWithRoleArnResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetCredentialDetailsAzureWithTenantId(credentialId string) (*MsgGetCredentialDetailsAzureWithTenantIdRequest, error) {
    //API: (GET) /V5/Credential/{credentialId}
    url := os.Getenv("CV_CSIP") + "/V5/Credential/" + credentialId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetCredentialDetailsAzureWithTenantIdRequest
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvUpdateCredentialAzureWithTenantId(updateCredentialAzureWithTenantIdRequest MsgUpdateCredentialAzureWithTenantIdRequest, credentialId string) (*MsgUpdateCredentialAzureWithTenantIdResponse, error) {
    //API: (PUT) /V5/Credential/{credentialId}
    reqBody, _ := json.Marshal(updateCredentialAzureWithTenantIdRequest)
    url := os.Getenv("CV_CSIP") + "/V5/Credential/" + credentialId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgUpdateCredentialAzureWithTenantIdResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateCredentialAzureWithTenantId(createCredentialAzureWithTenantIdRequest MsgCreateCredentialAzureWithTenantIdRequest) (*MsgCreateCredentialAzureWithTenantIdResponse, error) {
    //API: (POST) /V4/Credential
    reqBody, _ := json.Marshal(createCredentialAzureWithTenantIdRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Credential"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateCredentialAzureWithTenantIdResponse
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

func CvCreateBucketforCloudStorageAzure(createBucketforCloudStorageAzureRequest MsgCreateBucketforCloudStorageAzureRequest, cloudStorageId string) (*MsgCreateBucketforCloudStorageAzureResponse, error) {
    //API: (POST) /V4/Storage/Cloud/{cloudStorageId}/Bucket
    reqBody, _ := json.Marshal(createBucketforCloudStorageAzureRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud/" + cloudStorageId + "/Bucket"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateBucketforCloudStorageAzureResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetBucketDetailsOfCloudStorageAzure(cloudStorageId string, bucketId string) (*MsgGetBucketDetailsOfCloudStorageAzureRequest, error) {
    //API: (GET) /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud/" + cloudStorageId + "/Bucket/" + bucketId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetBucketDetailsOfCloudStorageAzureRequest
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvModifyBucketOfCloudStorageAzure(modifyBucketOfCloudStorageAzureRequest MsgModifyBucketOfCloudStorageAzureRequest, cloudStorageId string, bucketId string) (*MsgModifyBucketOfCloudStorageAzureResponse, error) {
    //API: (PUT) /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}
    reqBody, _ := json.Marshal(modifyBucketOfCloudStorageAzureRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud/" + cloudStorageId + "/Bucket/" + bucketId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgModifyBucketOfCloudStorageAzureResponse
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

func CvGetBucketDetailsOfCloudStorageS3(cloudStorageId string, bucketId string) (*MsgGetBucketDetailsOfCloudStorageS3Request, error) {
    //API: (GET) /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud/" + cloudStorageId + "/Bucket/" + bucketId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetBucketDetailsOfCloudStorageS3Request
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvModifyBucketOfCloudStorageS3(modifyBucketOfCloudStorageS3Request MsgModifyBucketOfCloudStorageS3Request, cloudStorageId string, bucketId string) (*MsgModifyBucketOfCloudStorageS3Response, error) {
    //API: (PUT) /V4/Storage/Cloud/{cloudStorageId}/Bucket/{bucketId}
    reqBody, _ := json.Marshal(modifyBucketOfCloudStorageS3Request)
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Cloud/" + cloudStorageId + "/Bucket/" + bucketId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgModifyBucketOfCloudStorageS3Response
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

func CvGetRoleDetails(roleId string) (*MsgGetRoleDetailsResponse, error) {
    //API: (GET) /V4/Role/{roleId}
    url := os.Getenv("CV_CSIP") + "/V4/Role/" + roleId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetRoleDetailsResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteRoles(roleId string) (*MsgDeleteRolesResponse, error) {
    //API: (DELETE) /V4/Role/{roleId}
    url := os.Getenv("CV_CSIP") + "/V4/Role/" + roleId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteRolesResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvModifyRole(modifyRoleRequest MsgModifyRoleRequest, roleId string) (*MsgModifyRoleResponse, error) {
    //API: (PUT) /V4/Role/{roleId}
    reqBody, _ := json.Marshal(modifyRoleRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Role/" + roleId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgModifyRoleResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateNewRole(createNewRoleRequest MsgCreateNewRoleRequest) (*MsgCreateNewRoleResponse, error) {
    //API: (POST) /V4/Role
    reqBody, _ := json.Marshal(createNewRoleRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Role"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateNewRoleResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetBackupDestinationDetailsWithoutPlanInfo(BackupDestinationId string) (*MsgGetBackupDestinationDetailsWithoutPlanInfoResponse, error) {
    //API: (GET) /V4/Plan/BackupDestination/{BackupDestinationId}
    url := os.Getenv("CV_CSIP") + "/V4/Plan/BackupDestination/" + BackupDestinationId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetBackupDestinationDetailsWithoutPlanInfoResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteBackupDestinationWithoutPlanInfo(BackupDestinationId string) (*MsgDeleteBackupDestinationWithoutPlanInfoResponse, error) {
    //API: (DELETE) /V4/Plan/BackupDestination/{BackupDestinationId}
    url := os.Getenv("CV_CSIP") + "/V4/Plan/BackupDestination/" + BackupDestinationId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteBackupDestinationWithoutPlanInfoResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvModifyBackupDestinationWithoutPlanInfo(modifyBackupDestinationWithoutPlanInfoRequest MsgModifyBackupDestinationWithoutPlanInfoRequest, BackupDestinationId string) (*MsgModifyBackupDestinationWithoutPlanInfoResponse, error) {
    //API: (PUT) /V4/Plan/BackupDestination/{BackupDestinationId}
    reqBody, _ := json.Marshal(modifyBackupDestinationWithoutPlanInfoRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Plan/BackupDestination/" + BackupDestinationId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgModifyBackupDestinationWithoutPlanInfoResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateBackupDestinationWithoutPlanInfo(createBackupDestinationWithoutPlanInfoRequest MsgCreateBackupDestinationWithoutPlanInfoRequest) (*MsgCreateBackupDestinationWithoutPlanInfoResponse, error) {
    //API: (POST) /V4/Plan/BackupDestinations
    reqBody, _ := json.Marshal(createBackupDestinationWithoutPlanInfoRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Plan/BackupDestinations"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateBackupDestinationWithoutPlanInfoResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvAddMediaAgent(addMediaAgentRequest MsgAddMediaAgentRequest, storagePoolId string, backupLocationId string) (*MsgAddMediaAgentResponse, error) {
    //API: (POST) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}/AccessPath
    reqBody, _ := json.Marshal(addMediaAgentRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Disk/" + storagePoolId + "/BackupLocation/" + backupLocationId + "/AccessPath"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgAddMediaAgentResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteDiskAccessPath(storagePoolId string, backupLocationId string, accessPathId string) (*MsgDeleteDiskAccessPathResponse, error) {
    //API: (DELETE) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}/AccessPath/{accessPathId}
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Disk/" + storagePoolId + "/BackupLocation/" + backupLocationId + "/AccessPath/" + accessPathId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteDiskAccessPathResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvModifyDiskAccessPath(modifyDiskAccessPathRequest MsgModifyDiskAccessPathRequest, storagePoolId string, backupLocationId string, accessPathId string) (*MsgModifyDiskAccessPathResponse, error) {
    //API: (PUT) /V4/Storage/Disk/{storagePoolId}/BackupLocation/{backupLocationId}/AccessPath/{accessPathId}
    reqBody, _ := json.Marshal(modifyDiskAccessPathRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Storage/Disk/" + storagePoolId + "/BackupLocation/" + backupLocationId + "/AccessPath/" + accessPathId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgModifyDiskAccessPathResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetApplicationGroupDetails(applicationGroupId string) (*MsgGetApplicationGroupDetailsResponse, error) {
    //API: (GET) /V5/Kubernetes/ApplicationGroup/{applicationGroupId}
    url := os.Getenv("CV_CSIP") + "/V5/Kubernetes/ApplicationGroup/" + applicationGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetApplicationGroupDetailsResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteKubernetesAppGroup(applicationGroupId string) (*MsgDeleteKubernetesAppGroupResponse, error) {
    //API: (DELETE) /V5/Kubernetes/ApplicationGroup/{applicationGroupId}
    url := os.Getenv("CV_CSIP") + "/V5/Kubernetes/ApplicationGroup/" + applicationGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteKubernetesAppGroupResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvUpdateKubernetesAppGroupOp(updateKubernetesAppGroupOpRequest MsgUpdateKubernetesAppGroupOpRequest, applicationGroupId string) (*MsgUpdateKubernetesAppGroupOpResponse, error) {
    //API: (PUT) /V5/Kubernetes/ApplicationGroup/{applicationGroupId}
    reqBody, _ := json.Marshal(updateKubernetesAppGroupOpRequest)
    url := os.Getenv("CV_CSIP") + "/V5/Kubernetes/ApplicationGroup/" + applicationGroupId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgUpdateKubernetesAppGroupOpResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateKubernetesApplicationGroup(createKubernetesApplicationGroupRequest MsgCreateKubernetesApplicationGroupRequest) (*MsgCreateKubernetesApplicationGroupResponse, error) {
    //API: (POST) /V5/Kubernetes/ApplicationGroup
    reqBody, _ := json.Marshal(createKubernetesApplicationGroupRequest)
    url := os.Getenv("CV_CSIP") + "/V5/Kubernetes/ApplicationGroup"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateKubernetesApplicationGroupResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvPOSTFirewallTopology(pOSTFirewallTopologyRequest MsgPOSTFirewallTopologyRequest) (*MsgPOSTFirewallTopologyResponse, error) {
    //API: (POST) /V4/NetworkTopology
    reqBody, _ := json.Marshal(pOSTFirewallTopologyRequest)
    url := os.Getenv("CV_CSIP") + "/V4/NetworkTopology"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgPOSTFirewallTopologyResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGETFirewallTopologyDetails(topologyId string) (*MsgGETFirewallTopologyDetailsResponse, error) {
    //API: (GET) /V4/NetworkTopology/{topologyId}
    url := os.Getenv("CV_CSIP") + "/V4/NetworkTopology/" + topologyId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGETFirewallTopologyDetailsResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDELETEFirewallTopology(topologyId string) (*MsgDELETEFirewallTopologyResponse, error) {
    //API: (DELETE) /V4/NetworkTopology/{topologyId}
    url := os.Getenv("CV_CSIP") + "/V4/NetworkTopology/" + topologyId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDELETEFirewallTopologyResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvPUTFirewallTopology(pUTFirewallTopologyRequest MsgPUTFirewallTopologyRequest, topologyId string) (*MsgPUTFirewallTopologyResponse, error) {
    //API: (PUT) /V4/NetworkTopology/{topologyId}
    reqBody, _ := json.Marshal(pUTFirewallTopologyRequest)
    url := os.Getenv("CV_CSIP") + "/V4/NetworkTopology/" + topologyId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgPUTFirewallTopologyResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvCreateRegion(createRegionRequest MsgCreateRegionRequest) (*MsgCreateRegionResponse, error) {
    //API: (POST) /V4/Regions
    reqBody, _ := json.Marshal(createRegionRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Regions"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateRegionResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetRegionDetails(regionId string) (*MsgGetRegionDetailsResponse, error) {
    //API: (GET) /V4/Regions/{regionId}
    url := os.Getenv("CV_CSIP") + "/V4/Regions/" + regionId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetRegionDetailsResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteRegion(regionId string) (*MsgDeleteRegionResponse, error) {
    //API: (DELETE) /V4/Regions/{regionId}
    url := os.Getenv("CV_CSIP") + "/V4/Regions/" + regionId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    var respObj MsgDeleteRegionResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvUpdateRegion(updateRegionRequest MsgUpdateRegionRequest, regionId string) (*MsgUpdateRegionResponse, error) {
    //API: (PUT) /V4/Regions/{regionId}
    reqBody, _ := json.Marshal(updateRegionRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Regions/" + regionId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    var respObj MsgUpdateRegionResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvGetCloudConnectionDetails(cloudConnectionId string) (*MsgGetCloudConnectionDetailsResponse, error) {
    //API: (GET) /V4/Cloud/CloudConnection/{cloudConnectionId}
    url := os.Getenv("CV_CSIP") + "/V4/Cloud/CloudConnection/" + cloudConnectionId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, "", token, 0)
    var respObj MsgGetCloudConnectionDetailsResponse
    json.Unmarshal(respBody, &respObj)
    return &respObj, err
}

func CvDeleteCloudConnection(cloudConnectionId string) ([]byte, error) {
    //API: (DELETE) /V4/Cloud/CloudConnection/{cloudConnectionId}
    url := os.Getenv("CV_CSIP") + "/V4/Cloud/CloudConnection/" + cloudConnectionId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, "", token, 0)
    return respBody, err
}

func CvUpdateCloudConnection(updateCloudConnectionRequest MsgUpdateCloudConnectionRequest, cloudConnectionId string) ([]byte, error) {
    //API: (PUT) /V4/Cloud/CloudConnection/{cloudConnectionId}
    reqBody, _ := json.Marshal(updateCloudConnectionRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Cloud/CloudConnection/" + cloudConnectionId
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPut, JSON, reqBody, JSON, token, 0)
    return respBody, err
}

func CvCreateCloudConnection(createCloudConnectionRequest MsgCreateCloudConnectionRequest) (*MsgCreateCloudConnectionResponse, error) {
    //API: (POST) /V4/Cloud/CloudConnection
    reqBody, _ := json.Marshal(createCloudConnectionRequest)
    url := os.Getenv("CV_CSIP") + "/V4/Cloud/CloudConnection"
    token := os.Getenv("AuthToken")
    respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
    var respObj MsgCreateCloudConnectionResponse
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
