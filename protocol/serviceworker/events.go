package serviceworker


type WorkerRegistrationUpdatedEvent struct {
    Registrations []ServiceWorkerRegistration `json:"registrations"`
}
type WorkerVersionUpdatedEvent struct {
    Versions []ServiceWorkerVersion `json:"versions"`
}
type WorkerErrorReportedEvent struct {
    ErrorMessage ServiceWorkerErrorMessage `json:"errorMessage"`
}