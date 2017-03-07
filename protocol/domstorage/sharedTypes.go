package domstorage

type StorageId struct {
    SecurityOrigin string `json:"securityOrigin"`// Security origin for the storage.
    IsLocalStorage bool `json:"isLocalStorage"`// Whether the storage is local storage (not session storage).
}

type Item []string
