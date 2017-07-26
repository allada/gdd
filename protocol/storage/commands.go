package storage


import (
    "github.com/allada/gdd/protocol/shared"
)

type ClearDataForOriginCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Origin string `json:"origin"`// Security origin.
    StorageTypes string `json:"storageTypes"`// Comma separated origin names.
}
type ClearDataForOriginReturn struct {
}

func (c *ClearDataForOriginCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ClearDataForOriginCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ClearDataForOriginCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}
