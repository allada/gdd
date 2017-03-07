package memory


import (
    "../shared"
)

type GetDOMCountersCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type GetDOMCountersReturn struct {
    Documents int64 `json:"documents"`
    Nodes int64 `json:"nodes"`
    JsEventListeners int64 `json:"jsEventListeners"`
}

func (c *GetDOMCountersCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetDOMCountersCommand) Respond(r *GetDOMCountersReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetPressureNotificationsSuppressedCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Suppressed bool `json:"suppressed"`// If true, memory pressure notifications will be suppressed.
}
type SetPressureNotificationsSuppressedReturn struct {
}

func (c *SetPressureNotificationsSuppressedCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetPressureNotificationsSuppressedCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SimulatePressureNotificationCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Level PressureLevel `json:"level"`// Memory pressure level of the notification.
}
type SimulatePressureNotificationReturn struct {
}

func (c *SimulatePressureNotificationCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SimulatePressureNotificationCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}
