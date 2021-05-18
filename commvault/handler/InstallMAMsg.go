package handler

type InstallMARequest struct {
	RebootClient              bool       `json:"rebootClient"`
	Entities                  []Entities `json:"entities"`
	CreatePseudoClientRequest struct {
		RegisterClient bool `json:"registerClient"`
		ClientInfo     struct {
			ClientType int `json:"clientType"`
		} `json:"clientInfo"`
	} `json:"createPseudoClientRequest"`
	Packages         []Packages `json:"packages"`
	ClientAuthForJob struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	} `json:"clientAuthForJob"`
}

type Entities struct {
	HostName   string `json:"hostName"`
	ClientName string `json:"clientName"`
	ClientID   int    `json:"clientId"`
}

type Packages struct {
	PackageID   int    `json:"packageId"`
	PackageName string `json:"packageName"`
}

type InstallMAResp struct {
	JobID    int `json:"jobId"`
	TaskID   int `json:"taskId"`
	Response struct {
		ErrorCode int `json:"errorCode"`
	} `json:"response"`
}

type JobSummaryResponse struct {
	TotalRecordsWithoutPaging int `json:"totalRecordsWithoutPaging"`
	Jobs                      []struct {
		JobSummary struct {
			SizeOfApplication      int    `json:"sizeOfApplication"`
			BackupSetName          string `json:"backupSetName"`
			TotalFailedFolders     int    `json:"totalFailedFolders"`
			TotalFailedFiles       int    `json:"totalFailedFiles"`
			IsVisible              bool   `json:"isVisible"`
			LocalizedStatus        string `json:"localizedStatus"`
			IsAged                 bool   `json:"isAged"`
			TotalNumOfFiles        int    `json:"totalNumOfFiles"`
			JobID                  int    `json:"jobId"`
			SizeOfMediaOnDisk      int    `json:"sizeOfMediaOnDisk"`
			Status                 string `json:"status"`
			LastUpdateTime         int    `json:"lastUpdateTime"`
			PercentSavings         int    `json:"percentSavings"`
			LocalizedOperationName string `json:"localizedOperationName"`
			StatusColor            string `json:"statusColor"`
			BackupLevel            int    `json:"backupLevel"`
			JobElapsedTime         int    `json:"jobElapsedTime"`
			JobStartTime           int    `json:"jobStartTime"`
			JobType                string `json:"jobType"`
			IsPreemptable          int    `json:"isPreemptable"`
			AppTypeName            string `json:"appTypeName"`
			PercentComplete        int    `json:"percentComplete"`
			DestClientName         string `json:"destClientName"`
			Subclient              struct {
				ClientName    string `json:"clientName"`
				InstanceName  string `json:"instanceName"`
				BackupsetID   int    `json:"backupsetId"`
				InstanceID    int    `json:"instanceId"`
				SubclientID   int    `json:"subclientId"`
				ClientID      int    `json:"clientId"`
				AppName       string `json:"appName"`
				BackupsetName string `json:"backupsetName"`
				ApplicationID int    `json:"applicationId"`
			} `json:"subclient"`
			ClientGroups []struct {
				Type            int    `json:"_type_"`
				ClientGroupID   int    `json:"clientGroupId"`
				ClientGroupName string `json:"clientGroupName"`
			} `json:"clientGroups"`
		} `json:"jobSummary"`
	} `json:"jobs"`
}

type ClientDetails struct {
	ClientName string `json:"clientName"`
	ClientID   int    `json:"clientId"`
	Type       int    `json:"_type_"`
}

type UninstalResp struct {
	JobID    int `json:"jobId"`
	Response struct {
		ErrorString string `json:"errorString"`
		ErrorCode   int    `json:"errorCode"`
	} `json:"response"`
}
