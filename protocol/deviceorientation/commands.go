package deviceorientation


import (
    "github.com/allada/gdd/protocol/shared"
)

type SetDeviceOrientationOverrideCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Alpha float64 `json:"alpha"`// Mock alpha
    Beta float64 `json:"beta"`// Mock beta
    Gamma float64 `json:"gamma"`// Mock gamma
}
type SetDeviceOrientationOverrideReturn struct {
}

func (c *SetDeviceOrientationOverrideCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetDeviceOrientationOverrideCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetDeviceOrientationOverrideCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ClearDeviceOrientationOverrideCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type ClearDeviceOrientationOverrideReturn struct {
}

func (c *ClearDeviceOrientationOverrideCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ClearDeviceOrientationOverrideCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ClearDeviceOrientationOverrideCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}
