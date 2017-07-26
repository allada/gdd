package serviceworker


import (
    "github.com/allada/gdd/protocol/target"
)
const (
    ServiceWorkerVersionRunningStatusStopped ServiceWorkerVersionRunningStatus = "stopped"
    ServiceWorkerVersionRunningStatusStarting ServiceWorkerVersionRunningStatus = "starting"
    ServiceWorkerVersionRunningStatusRunning ServiceWorkerVersionRunningStatus = "running"
    ServiceWorkerVersionRunningStatusStopping ServiceWorkerVersionRunningStatus = "stopping"
)

const (
    ServiceWorkerVersionStatusNew ServiceWorkerVersionStatus = "new"
    ServiceWorkerVersionStatusInstalling ServiceWorkerVersionStatus = "installing"
    ServiceWorkerVersionStatusInstalled ServiceWorkerVersionStatus = "installed"
    ServiceWorkerVersionStatusActivating ServiceWorkerVersionStatus = "activating"
    ServiceWorkerVersionStatusActivated ServiceWorkerVersionStatus = "activated"
    ServiceWorkerVersionStatusRedundant ServiceWorkerVersionStatus = "redundant"
)
type ServiceWorkerRegistration struct {
    RegistrationId string `json:"registrationId"`
    ScopeURL string `json:"scopeURL"`
    IsDeleted bool `json:"isDeleted"`
}

type ServiceWorkerVersionRunningStatus string

type ServiceWorkerVersionStatus string

type ServiceWorkerVersion struct {
    VersionId string `json:"versionId"`
    RegistrationId string `json:"registrationId"`
    ScriptURL string `json:"scriptURL"`
    RunningStatus ServiceWorkerVersionRunningStatus `json:"runningStatus"`
    Status ServiceWorkerVersionStatus `json:"status"`
    ScriptLastModified *float64 `json:"scriptLastModified,omitempty"`// The Last-Modified header value of the main script.
    ScriptResponseTime *float64 `json:"scriptResponseTime,omitempty"`// The time at which the response headers of the main script were received from the server.  For cached script it is the last time the cache entry was validated.
    ControlledClients *[]target.TargetID `json:"controlledClients,omitempty"`
    TargetId *target.TargetID `json:"targetId,omitempty"`
}

type ServiceWorkerErrorMessage struct {
    ErrorMessage string `json:"errorMessage"`
    RegistrationId string `json:"registrationId"`
    VersionId string `json:"versionId"`
    SourceURL string `json:"sourceURL"`
    LineNumber int64 `json:"lineNumber"`
    ColumnNumber int64 `json:"columnNumber"`
}
