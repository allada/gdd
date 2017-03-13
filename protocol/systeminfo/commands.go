package systeminfo


import (
    "../shared"
)

type GetInfoCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type GetInfoReturn struct {
    Gpu GPUInfo `json:"gpu"`// Information about the GPUs on the system.
    ModelName string `json:"modelName"`// A platform-dependent description of the model of the machine. On Mac OS, this is, for example, 'MacBookPro'. Will be the empty string if not supported.
    ModelVersion string `json:"modelVersion"`// A platform-dependent description of the version of the machine. On Mac OS, this is, for example, '10.1'. Will be the empty string if not supported.
}

func (c *GetInfoCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetInfoCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetInfoCommand) Respond(r *GetInfoReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
