package schema


import (
    "../shared"
)

type GetDomainsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type GetDomainsReturn struct {
    Domains []Domain `json:"domains"`// List of supported domains.
}

func (c *GetDomainsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetDomainsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetDomainsCommand) Respond(r *GetDomainsReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
