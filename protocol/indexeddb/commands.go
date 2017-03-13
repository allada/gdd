package indexeddb


import (
    "../shared"
)

type EnableCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type EnableReturn struct {
}

func (c *EnableCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *EnableCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *EnableCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type DisableCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type DisableReturn struct {
}

func (c *DisableCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *DisableCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DisableCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type RequestDatabaseNamesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    SecurityOrigin string `json:"securityOrigin"`// Security origin.
}
type RequestDatabaseNamesReturn struct {
    DatabaseNames []string `json:"databaseNames"`// Database names for origin.
}

func (c *RequestDatabaseNamesCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RequestDatabaseNamesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RequestDatabaseNamesCommand) Respond(r *RequestDatabaseNamesReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type RequestDatabaseCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    SecurityOrigin string `json:"securityOrigin"`// Security origin.
    DatabaseName string `json:"databaseName"`// Database name.
}
type RequestDatabaseReturn struct {
    DatabaseWithObjectStores DatabaseWithObjectStores `json:"databaseWithObjectStores"`// Database with an array of object stores.
}

func (c *RequestDatabaseCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RequestDatabaseCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RequestDatabaseCommand) Respond(r *RequestDatabaseReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type RequestDataCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    SecurityOrigin string `json:"securityOrigin"`// Security origin.
    DatabaseName string `json:"databaseName"`// Database name.
    ObjectStoreName string `json:"objectStoreName"`// Object store name.
    IndexName string `json:"indexName"`// Index name, empty string for object store data requests.
    SkipCount int64 `json:"skipCount"`// Number of records to skip.
    PageSize int64 `json:"pageSize"`// Number of records to fetch.
    KeyRange *KeyRange `json:"keyRange,omitempty"`// Key range.
}
type RequestDataReturn struct {
    ObjectStoreDataEntries []DataEntry `json:"objectStoreDataEntries"`// Array of object store data entries.
    HasMore bool `json:"hasMore"`// If true, there are more entries to fetch in the given range.
}

func (c *RequestDataCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RequestDataCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RequestDataCommand) Respond(r *RequestDataReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type ClearObjectStoreCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    SecurityOrigin string `json:"securityOrigin"`// Security origin.
    DatabaseName string `json:"databaseName"`// Database name.
    ObjectStoreName string `json:"objectStoreName"`// Object store name.
}
type ClearObjectStoreReturn struct {
}

func (c *ClearObjectStoreCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ClearObjectStoreCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ClearObjectStoreCommand) Respond(r *ClearObjectStoreReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type DeleteDatabaseCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    SecurityOrigin string `json:"securityOrigin"`// Security origin.
    DatabaseName string `json:"databaseName"`// Database name.
}
type DeleteDatabaseReturn struct {
}

func (c *DeleteDatabaseCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *DeleteDatabaseCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DeleteDatabaseCommand) Respond(r *DeleteDatabaseReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
