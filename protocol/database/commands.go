package database


import (
    "github.com/allada/gdd/protocol/shared"
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

type GetDatabaseTableNamesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    DatabaseId DatabaseId `json:"databaseId"`
}
type GetDatabaseTableNamesReturn struct {
    TableNames []string `json:"tableNames"`
}

func (c *GetDatabaseTableNamesCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetDatabaseTableNamesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetDatabaseTableNamesCommand) Respond(r *GetDatabaseTableNamesReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type ExecuteSQLCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    DatabaseId DatabaseId `json:"databaseId"`
    Query string `json:"query"`
}
type ExecuteSQLReturn struct {
    ColumnNames *[]string `json:"columnNames,omitempty"`
    Values *[]interface{} `json:"values,omitempty"`
    SqlError *Error `json:"sqlError,omitempty"`
}

func (c *ExecuteSQLCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ExecuteSQLCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ExecuteSQLCommand) Respond(r *ExecuteSQLReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
