package tethering


import (
    "github.com/allada/gdd/protocol/shared"
)

type BindCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Port int64 `json:"port"`// Port number to bind.
}
type BindReturn struct {
}

func (c *BindCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *BindCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *BindCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type UnbindCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Port int64 `json:"port"`// Port number to unbind.
}
type UnbindReturn struct {
}

func (c *UnbindCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *UnbindCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *UnbindCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}
