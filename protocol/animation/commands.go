package animation


import (
    "github.com/allada/gdd/protocol/shared"
    "github.com/allada/gdd/protocol/runtime"
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

type GetPlaybackRateCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type GetPlaybackRateReturn struct {
    PlaybackRate float64 `json:"playbackRate"`// Playback rate for animations on page.
}

func (c *GetPlaybackRateCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetPlaybackRateCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetPlaybackRateCommand) Respond(r *GetPlaybackRateReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetPlaybackRateCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    PlaybackRate float64 `json:"playbackRate"`// Playback rate for animations on page
}
type SetPlaybackRateReturn struct {
}

func (c *SetPlaybackRateCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetPlaybackRateCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetPlaybackRateCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetCurrentTimeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Id string `json:"id"`// Id of animation.
}
type GetCurrentTimeReturn struct {
    CurrentTime float64 `json:"currentTime"`// Current time of the page.
}

func (c *GetCurrentTimeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetCurrentTimeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetCurrentTimeCommand) Respond(r *GetCurrentTimeReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetPausedCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Animations []string `json:"animations"`// Animations to set the pause state of.
    Paused bool `json:"paused"`// Paused state to set to.
}
type SetPausedReturn struct {
}

func (c *SetPausedCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetPausedCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetPausedCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetTimingCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    AnimationId string `json:"animationId"`// Animation id.
    Duration float64 `json:"duration"`// Duration of the animation.
    Delay float64 `json:"delay"`// Delay of the animation.
}
type SetTimingReturn struct {
}

func (c *SetTimingCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetTimingCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetTimingCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SeekAnimationsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Animations []string `json:"animations"`// List of animation ids to seek.
    CurrentTime float64 `json:"currentTime"`// Set the current time of each animation.
}
type SeekAnimationsReturn struct {
}

func (c *SeekAnimationsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SeekAnimationsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SeekAnimationsCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ReleaseAnimationsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Animations []string `json:"animations"`// List of animation ids to seek.
}
type ReleaseAnimationsReturn struct {
}

func (c *ReleaseAnimationsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ReleaseAnimationsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ReleaseAnimationsCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ResolveAnimationCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    AnimationId string `json:"animationId"`// Animation id.
}
type ResolveAnimationReturn struct {
    RemoteObject runtime.RemoteObject `json:"remoteObject"`// Corresponding remote object.
}

func (c *ResolveAnimationCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ResolveAnimationCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ResolveAnimationCommand) Respond(r *ResolveAnimationReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
