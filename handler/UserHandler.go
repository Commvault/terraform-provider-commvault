package handler

import (
    "encoding/xml"
    "net/http"
    "os"
)

func UserCreate(userCreateRequest AppCreateUserRequest) *AppCreateUserResponse{
    url := os.Getenv("CV_CSIP") + "/user"
    token := os.Getenv("AuthToken")
    userCreateXml,_ := xml.Marshal(userCreateRequest)
    respBody := makeHttpRequest(url,http.MethodPost,XML,userCreateXml,XML,token)
    var userResp AppCreateUserResponse
    xml.Unmarshal(respBody,&userResp)
    return &userResp
}


func UserDelete(id string) *AppDeleteUserResponse {
    url := os.Getenv("CV_CSIP") + "/user/" +id
    token := os.Getenv("AuthToken")
    if token == "" {
        token = GenerateAuthToken(os.Getenv("CV_USERNAME"),os.Getenv("CV_PASSWORD"))
    }
    respBody := makeHttpRequest(url,http.MethodDelete,"",nil,"",token)
    var deleteUserResp AppDeleteUserResponse
    xml.Unmarshal(respBody,&deleteUserResp)
    return &deleteUserResp
}


func UpdateUser(updateUserRequest AppUpdateUserPropertiesRequest,userId string) *AppUpdateUserPropertiesResponse{
    url := os.Getenv("CV_CSIP") + "/user/"+userId
    token := os.Getenv("AuthToken")
    if token == "" {
        token = GenerateAuthToken(os.Getenv("CV_USERNAME"),os.Getenv("CV_PASSWORD"))
    }
    updateUserXMLRequest,_ := xml.Marshal(updateUserRequest)
    respBody := makeHttpRequest(url,http.MethodPost,XML,updateUserXMLRequest,XML,token)
    var userResp AppUpdateUserPropertiesResponse
    xml.Unmarshal(respBody,&userResp)
    return &userResp
}

func ReadUser(userId string) *UserPropertiesResp{
    url := os.Getenv("CV_CSIP") + "/user/"+userId
    token := os.Getenv("AuthToken")
    respBody := makeHttpRequest(url,http.MethodGet,XML,nil,XML,token)
    var userResp UserPropertiesResp
    xml.Unmarshal(respBody,&userResp)
    return &userResp
}