package database


type AddDatabaseEvent struct {
    Database Database `json:"database"`
}