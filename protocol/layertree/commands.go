package layertree


import (
    "../shared"
    "../dom"
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

type CompositingReasonsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    LayerId LayerId `json:"layerId"`// The id of the layer for which we want to get the reasons it was composited.
}
type CompositingReasonsReturn struct {
    CompositingReasons []string `json:"compositingReasons"`// A list of strings specifying reasons for the given layer to become composited.
}

func (c *CompositingReasonsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *CompositingReasonsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CompositingReasonsCommand) Respond(r *CompositingReasonsReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type MakeSnapshotCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    LayerId LayerId `json:"layerId"`// The id of the layer.
}
type MakeSnapshotReturn struct {
    SnapshotId SnapshotId `json:"snapshotId"`// The id of the layer snapshot.
}

func (c *MakeSnapshotCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *MakeSnapshotCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *MakeSnapshotCommand) Respond(r *MakeSnapshotReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type LoadSnapshotCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Tiles []PictureTile `json:"tiles"`// An array of tiles composing the snapshot.
}
type LoadSnapshotReturn struct {
    SnapshotId SnapshotId `json:"snapshotId"`// The id of the snapshot.
}

func (c *LoadSnapshotCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *LoadSnapshotCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *LoadSnapshotCommand) Respond(r *LoadSnapshotReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type ReleaseSnapshotCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    SnapshotId SnapshotId `json:"snapshotId"`// The id of the layer snapshot.
}
type ReleaseSnapshotReturn struct {
}

func (c *ReleaseSnapshotCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ReleaseSnapshotCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ReleaseSnapshotCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ProfileSnapshotCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    SnapshotId SnapshotId `json:"snapshotId"`// The id of the layer snapshot.
    MinRepeatCount *int64 `json:"minRepeatCount,omitempty"`// The maximum number of times to replay the snapshot (1, if not specified).
    MinDuration *float64 `json:"minDuration,omitempty"`// The minimum duration (in seconds) to replay the snapshot.
    ClipRect *dom.Rect `json:"clipRect,omitempty"`// The clip rectangle to apply when replaying the snapshot.
}
type ProfileSnapshotReturn struct {
    Timings []PaintProfile `json:"timings"`// The array of paint profiles, one per run.
}

func (c *ProfileSnapshotCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ProfileSnapshotCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ProfileSnapshotCommand) Respond(r *ProfileSnapshotReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type ReplaySnapshotCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    SnapshotId SnapshotId `json:"snapshotId"`// The id of the layer snapshot.
    FromStep *int64 `json:"fromStep,omitempty"`// The first step to replay from (replay from the very start if not specified).
    ToStep *int64 `json:"toStep,omitempty"`// The last step to replay to (replay till the end if not specified).
    Scale *float64 `json:"scale,omitempty"`// The scale to apply while replaying (defaults to 1).
}
type ReplaySnapshotReturn struct {
    DataURL string `json:"dataURL"`// A data: URL for resulting image.
}

func (c *ReplaySnapshotCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ReplaySnapshotCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ReplaySnapshotCommand) Respond(r *ReplaySnapshotReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SnapshotCommandLogCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    SnapshotId SnapshotId `json:"snapshotId"`// The id of the layer snapshot.
}
type SnapshotCommandLogReturn struct {
    CommandLog []map[string]string `json:"commandLog"`// The array of canvas function calls.
}

func (c *SnapshotCommandLogCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SnapshotCommandLogCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SnapshotCommandLogCommand) Respond(r *SnapshotCommandLogReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
