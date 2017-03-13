package domstorage


type DomStorageItemsClearedEvent struct {
    StorageId StorageId `json:"storageId"`
}
type DomStorageItemRemovedEvent struct {
    StorageId StorageId `json:"storageId"`
    Key string `json:"key"`
}
type DomStorageItemAddedEvent struct {
    StorageId StorageId `json:"storageId"`
    Key string `json:"key"`
    NewValue string `json:"newValue"`
}
type DomStorageItemUpdatedEvent struct {
    StorageId StorageId `json:"storageId"`
    Key string `json:"key"`
    OldValue string `json:"oldValue"`
    NewValue string `json:"newValue"`
}