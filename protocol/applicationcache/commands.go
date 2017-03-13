package applicationcache


import (
    "../shared"
    "../page"
)

type GetFramesWithManifestsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type GetFramesWithManifestsReturn struct {
    FrameIds []FrameWithManifest `json:"frameIds"`// Array of frame identifiers with manifest urls for each frame containing a document associated with some application cache.
}

func (c *GetFramesWithManifestsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetFramesWithManifestsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetFramesWithManifestsCommand) Respond(r *GetFramesWithManifestsReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

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

type GetManifestForFrameCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    FrameId page.FrameId `json:"frameId"`// Identifier of the frame containing document whose manifest is retrieved.
}
type GetManifestForFrameReturn struct {
    ManifestURL string `json:"manifestURL"`// Manifest URL for document in the given frame.
}

func (c *GetManifestForFrameCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetManifestForFrameCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetManifestForFrameCommand) Respond(r *GetManifestForFrameReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetApplicationCacheForFrameCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    FrameId page.FrameId `json:"frameId"`// Identifier of the frame containing document whose application cache is retrieved.
}
type GetApplicationCacheForFrameReturn struct {
    ApplicationCache ApplicationCache `json:"applicationCache"`// Relevant application cache data for the document in given frame.
}

func (c *GetApplicationCacheForFrameCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetApplicationCacheForFrameCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetApplicationCacheForFrameCommand) Respond(r *GetApplicationCacheForFrameReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
