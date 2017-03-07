package accessibility


import (
    "../shared"
    "../dom"
)

type GetPartialAXTreeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId dom.NodeId `json:"nodeId"`// ID of node to get the partial accessibility tree for.
    FetchRelatives *bool `json:"fetchRelatives,omitempty"`// Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
}
type GetPartialAXTreeReturn struct {
    Nodes []AXNode `json:"nodes"`// The <code>Accessibility.AXNode</code> for this DOM node, if it exists, plus its ancestors, siblings and children, if requested.
}

func (c *GetPartialAXTreeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetPartialAXTreeCommand) Respond(r *GetPartialAXTreeReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
