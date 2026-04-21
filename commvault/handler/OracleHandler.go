package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

// CvCreateOracleInstance creates an Oracle instance
func CvCreateOracleInstance(req MsgCreateOracleInstanceRequest) (*MsgCreateOracleInstanceResponse, error) {
	// API: (POST) /instance
	reqBody, _ := json.Marshal(req)
	url := os.Getenv("CV_CSIP") + "/instance"
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
	var respObj MsgCreateOracleInstanceResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvFetchOracleInstances fetches Oracle instances associated to a client
func CvFetchOracleInstances(clientName string) (*MsgFetchOracleInstancesResponse, error) {
	// API: (GET) /instance?clientName={clientName}&appName=Oracle
	url := os.Getenv("CV_CSIP") + "/instance?clientName=" + urlEscape(clientName) + "&appName=Oracle"
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgFetchOracleInstancesResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvGetOracleInstanceProperties fetches Oracle instance properties
func CvGetOracleInstanceProperties(instanceId string) (*MsgGetOracleInstancePropertiesResponse, error) {
	// API: (GET) /instance/{instanceId}
	url := os.Getenv("CV_CSIP") + "/instance/" + instanceId
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgGetOracleInstancePropertiesResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvModifyOracleInstance modifies an Oracle instance
func CvModifyOracleInstance(req MsgModifyOracleInstanceRequest, instanceId string) (*MsgModifyOracleInstanceResponse, error) {
	// API: (POST) /instance/{instanceId}
	reqBody, _ := json.Marshal(req)
	url := os.Getenv("CV_CSIP") + "/instance/" + instanceId
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
	var respObj MsgModifyOracleInstanceResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvDeleteOracleInstance deletes an Oracle instance
func CvDeleteOracleInstance(instanceId string) (*MsgDeleteOracleInstanceResponse, error) {
	// API: (DELETE) /instance/{instanceId}
	url := os.Getenv("CV_CSIP") + "/instance/" + instanceId
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, JSON, token, 0)
	var respObj MsgDeleteOracleInstanceResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvOracleInstanceDiscovery discovers Oracle instances on a client
func CvOracleInstanceDiscovery(clientId string) (*MsgOracleInstanceDiscoveryResponse, error) {
	// API: (GET) /client/{clientId}/instance/oracle/discover
	url := os.Getenv("CV_CSIP") + "/client/" + clientId + "/instance/oracle/discover"
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgOracleInstanceDiscoveryResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvGetOracleBackupPieces fetches Oracle backup pieces for a given time duration
func CvGetOracleBackupPieces(instanceId string, fromTime int, toTime int) (*MsgGetOracleBackupPiecesResponse, error) {
	// API: (GET) /oracle/instance/{instanceId}/backupPieces
	url := os.Getenv("CV_CSIP") + "/oracle/instance/" + instanceId + "/backupPieces"
	if fromTime > 0 {
		url += "?fromTime=" + strconv.Itoa(fromTime)
		if toTime > 0 {
			url += "&toTime=" + strconv.Itoa(toTime)
		}
	} else if toTime > 0 {
		url += "?toTime=" + strconv.Itoa(toTime)
	}
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgGetOracleBackupPiecesResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvBrowseOracleDB browses Oracle database instance
func CvBrowseOracleDB(req MsgBrowseOracleDBRequest, instanceId string) (*MsgBrowseOracleDBResponse, error) {
	// API: (POST) /instance/DBBrowse/{instanceId}
	reqBody, _ := json.Marshal(req)
	url := os.Getenv("CV_CSIP") + "/instance/DBBrowse/" + instanceId
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
	var respObj MsgBrowseOracleDBResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvCreateOracleSubclient creates an Oracle subclient
func CvCreateOracleSubclient(req MsgCreateOracleSubclientRequest) (*MsgCreateOracleSubclientResponse, error) {
	// API: (POST) /subclient
	reqBody, _ := json.Marshal(req)
	url := os.Getenv("CV_CSIP") + "/subclient"
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
	var respObj MsgCreateOracleSubclientResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvFetchOracleSubclients fetches subclients associated to an Oracle instance
func CvFetchOracleSubclients(clientId int, instanceId string) (*MsgFetchOracleSubclientsResponse, error) {
	// API: (GET) /subclient?clientId={clientId}&applicationId=22&instanceId={instanceId}
	url := os.Getenv("CV_CSIP") + "/subclient?clientId=" + strconv.Itoa(clientId) + "&applicationId=22&instanceId=" + instanceId
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgFetchOracleSubclientsResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvGetOracleSubclientProperties fetches Oracle subclient properties
func CvGetOracleSubclientProperties(subclientId string) (*MsgGetOracleSubclientPropertiesResponse, error) {
	// API: (GET) /subclient/{subclientId}
	url := os.Getenv("CV_CSIP") + "/subclient/" + subclientId
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgGetOracleSubclientPropertiesResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvModifyOracleSubclient modifies an Oracle subclient
func CvModifyOracleSubclient(req MsgModifyOracleSubclientRequest, subclientId string) (*MsgModifyOracleSubclientResponse, error) {
	// API: (POST) /subclient/{subclientId}
	reqBody, _ := json.Marshal(req)
	url := os.Getenv("CV_CSIP") + "/subclient/" + subclientId
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
	var respObj MsgModifyOracleSubclientResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvDeleteOracleSubclient deletes an Oracle subclient
func CvDeleteOracleSubclient(subclientId string) (*MsgDeleteOracleSubclientResponse, error) {
	// API: (DELETE) /subclient/{subclientId}
	url := os.Getenv("CV_CSIP") + "/subclient/" + subclientId
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodDelete, JSON, nil, JSON, token, 0)
	var respObj MsgDeleteOracleSubclientResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvOracleBackup triggers an Oracle backup
func CvOracleBackup(req MsgOracleBackupRequest) (*MsgOracleBackupResponse, error) {
	// API: (POST) /CreateTask
	reqBody, _ := json.Marshal(req)
	url := os.Getenv("CV_CSIP") + "/CreateTask"
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
	var respObj MsgOracleBackupResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvOracleRestore triggers an Oracle restore
func CvOracleRestore(req MsgOracleBackupRequest) (*MsgOracleRestoreResponse, error) {
	// API: (POST) /createTask
	reqBody, _ := json.Marshal(req)
	url := os.Getenv("CV_CSIP") + "/createTask"
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
	var respObj MsgOracleRestoreResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvFetchRMANLogs fetches RMAN logs for a job
func CvFetchRMANLogs(jobId string) (*MsgFetchRMANLogsResponse, error) {
	// API: (GET) /Job/{JobId}/RMANLogs
	url := os.Getenv("CV_CSIP") + "/Job/" + jobId + "/RMANLogs"
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodGet, JSON, nil, JSON, token, 0)
	var respObj MsgFetchRMANLogsResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}

// CvFetchOracleEntityId fetches Oracle entity IDs using the /GetId endpoint.
// Commvault returns HTTP 200 when all IDs are found, or HTTP 404 when some are
// not found — but in both cases the body contains the resolved IDs (with -32000
// as the sentinel for "not found").  We use makeHttpRequestBody so the body is
// never discarded.  Callers should check id > 0 to detect a valid ID.
func CvFetchOracleEntityId(clientName string, instanceName string, subclient string) (*MsgFetchOracleEntityIdResponse, error) {
	// API: (GET) /GetId?clientName={clientName}&agent=Oracle&instanceName={instanceName}&subclient={subclient}
	url := os.Getenv("CV_CSIP") + "/GetId?agent=Oracle"
	if clientName != "" {
		url += "&clientName=" + urlEscape(clientName)
	}
	if instanceName != "" {
		url += "&instanceName=" + urlEscape(instanceName)
	}
	if subclient != "" {
		url += "&subclient=" + urlEscape(subclient)
	}
	token := os.Getenv("AuthToken")
	respBody, statusCode, err := makeHttpRequestBody(url, http.MethodGet, JSON, nil, JSON, token, 0)
	if err != nil {
		return nil, err
	}
	var respObj MsgFetchOracleEntityIdResponse
	json.Unmarshal(respBody, &respObj)
	// HTTP 404 with -32000 means the entity was genuinely not found — return a
	// clear error rather than a confusing -32000 value.
	if statusCode == 404 && respObj.InstanceId <= 0 && respObj.ClientId <= 0 {
		return &respObj, nil // caller checks id > 0
	}
	return &respObj, nil
}

// CvInstallOracleAgent installs Oracle agent on database server
func CvInstallOracleAgent(req MsgInstallOracleAgentRequest) (*MsgInstallOracleAgentResponse, error) {
	// API: (POST) /Createtask
	reqBody, _ := json.Marshal(req)
	url := os.Getenv("CV_CSIP") + "/Createtask"
	token := os.Getenv("AuthToken")
	respBody, err := makeHttpRequestErr(url, http.MethodPost, JSON, reqBody, JSON, token, 0)
	var respObj MsgInstallOracleAgentResponse
	json.Unmarshal(respBody, &respObj)
	return &respObj, err
}
