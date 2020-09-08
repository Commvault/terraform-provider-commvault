package handler

import "encoding/xml"

type AppCreateSubClientRequest struct {
    XMLName             xml.Name `xml:"App_CreateSubClientRequest"`
    Text                string   `xml:",chardata"`
    SubClientProperties struct {
        Text                   string `xml:",chardata"`
        VmContentOperationType string `xml:"vmContentOperationType,attr"`
        SubClientEntity        struct {
            Text          string `xml:",chardata"`
            ClientId      string `xml:"clientId,attr"`
            ApplicationId string `xml:"applicationId,attr"`
            AppName       string `xml:"appName,attr"`
            SubclientName string `xml:"subclientName,attr"`
        } `xml:"subClientEntity"`
        CommonProperties struct {
            Text                  string `xml:",chardata"`
            EnableBackup          string `xml:"enableBackup,attr"`
            IsSnapbackupEnabled   string `xml:"isSnapbackupEnabled,attr"`
            NumberOfBackupStreams string `xml:"numberOfBackupStreams,attr"`
            SnapCopyInfo          struct {
                Text                   string `xml:",chardata"`
                IsSnapBackupEnabled    string `xml:"isSnapBackupEnabled,attr"`
                TransportModeForVMWare string `xml:"transportModeForVMWare,attr"`
            } `xml:"snapCopyInfo"`
        } `xml:"commonProperties"`
        VmContent struct {
            Text     string `xml:",chardata"`
            children []Children `xml:"children"`
        }
        VsaSubclientProp struct {
            Text                                  string `xml:",chardata"`
            QuiesceGuestFileSystemAndApplications string `xml:"quiesceGuestFileSystemAndApplications,attr"`
            AutoDetectVMOwner                     string `xml:"autoDetectVMOwner,attr"`
        } `xml:"vsaSubclientProp"`
        PlanEntity struct {
            Text   string `xml:",chardata"`
            PlanId string `xml:"planId,attr"`
        } `xml:"planEntity"`
    } `xml:"subClientProperties"`
}

type Children struct {
    AllOrAnyChildren       string `xml:"allOrAnyChildren,attr"`
    Name                   string `xml:"name,attr"`
    Type                   string `xml:"type,attr"`
    Path                   string `xml:"path,attr"`
    DisplayName            string `xml:"displayName,attr"`
    EqualsOrNotEquals      string `xml:"equalsOrNotEquals,attr"`
    GuestCredentialAssocId string `xml:"guestCredentialAssocId,attr"`
}

type AppCreateSubClientResponse struct {
    XMLName                   xml.Name `xml:"App_CreateSubClientResponse"`
    Text                      string   `xml:",chardata"`
    Processinginstructioninfo struct {
        Text       string `xml:",chardata"`
        Attributes struct {
            Text  string `xml:",chardata"`
            Name  string `xml:"name,attr"`
            Value string `xml:"value,attr"`
        } `xml:"attributes"`
    } `xml:"processinginstructioninfo"`
    Response struct {
        Text           string `xml:",chardata"`
        WarningCode    string `xml:"warningCode,attr"`
        ErrorCode      string `xml:"errorCode,attr"`
        WarningMessage string `xml:"warningMessage,attr"`
        Entity         struct {
            Text        string `xml:",chardata"`
            SubclientId string `xml:"subclientId,attr"`
            Type        string `xml:"_type_,attr"`
        } `xml:"entity"`
    } `xml:"response"`
}


type AppDeleteSubClientResponse struct {
    XMLName                   xml.Name `xml:"App_DeleteSubClientResponse"`
    Text                      string   `xml:",chardata"`
    Processinginstructioninfo struct {
        Text       string `xml:",chardata"`
        Attributes struct {
            Text  string `xml:",chardata"`
            Name  string `xml:"name,attr"`
            Value string `xml:"value,attr"`
        } `xml:"attributes"`
    } `xml:"processinginstructioninfo"`
    Response struct {
        Text      string `xml:",chardata"`
        ErrorCode string `xml:"errorCode,attr"`
        Entity    struct {
            Text          string `xml:",chardata"`
            SubclientId   string `xml:"subclientId,attr"`
            ClientName    string `xml:"clientName,attr"`
            ClientId      string `xml:"clientId,attr"`
            SubclientName string `xml:"subclientName,attr"`
            BackupsetName string `xml:"backupsetName,attr"`
            InstanceName  string `xml:"instanceName,attr"`
            Type          string `xml:"_type_,attr"`
            AppName       string `xml:"appName,attr"`
        } `xml:"entity"`
    } `xml:"response"`
}


type AppUpdateSubClientPropertiesRequest struct {
    XMLName             xml.Name `xml:"App_UpdateSubClientPropertiesRequest"`
    Text                string   `xml:",chardata"`
    SubClientProperties struct {
        Text                      string `xml:",chardata"`
        VmContentOperationType    string `xml:"vmContentOperationType,attr"`
        VmFilterOperationType     string `xml:"vmFilterOperationType,attr"`
        VmDiskFilterOperationType string `xml:"vmDiskFilterOperationType,attr"`
        VmContent                 struct {
            Children []Children `xml:"children"`
        } `xml:"vmContent"`
        VmFilter     string `xml:"vmFilter"`
        VmDiskFilter string `xml:"vmDiskFilter"`
    } `xml:"subClientProperties"`
}
