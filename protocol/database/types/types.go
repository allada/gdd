package types

type DatabaseId string

type Database struct {
    Id DatabaseId `json:"id"`// Database ID.
    Domain string `json:"domain"`// Database domain.
    Name string `json:"name"`// Database name.
    Version string `json:"version"`// Database version.
}

type Error struct {
    Message string `json:"message"`// Error message.
    Code int64 `json:"code"`// Error code.
}
