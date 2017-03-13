package profiler


import (
    "../shared"
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

type SetSamplingIntervalCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Interval int64 `json:"interval"`// New sampling interval in microseconds.
}
type SetSamplingIntervalReturn struct {
}

func (c *SetSamplingIntervalCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetSamplingIntervalCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetSamplingIntervalCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type StartCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type StartReturn struct {
}

func (c *StartCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *StartCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StartCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type StopCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type StopReturn struct {
    Profile Profile `json:"profile"`// Recorded profile.
}

func (c *StopCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *StopCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StopCommand) Respond(r *StopReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type StartPreciseCoverageCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type StartPreciseCoverageReturn struct {
}

func (c *StartPreciseCoverageCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *StartPreciseCoverageCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StartPreciseCoverageCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type StopPreciseCoverageCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type StopPreciseCoverageReturn struct {
}

func (c *StopPreciseCoverageCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *StopPreciseCoverageCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StopPreciseCoverageCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type TakePreciseCoverageCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type TakePreciseCoverageReturn struct {
    Result []ScriptCoverage `json:"result"`// Coverage data for the current isolate.
}

func (c *TakePreciseCoverageCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *TakePreciseCoverageCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *TakePreciseCoverageCommand) Respond(r *TakePreciseCoverageReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetBestEffortCoverageCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type GetBestEffortCoverageReturn struct {
    Result []ScriptCoverage `json:"result"`// Coverage data for the current isolate.
}

func (c *GetBestEffortCoverageCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetBestEffortCoverageCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetBestEffortCoverageCommand) Respond(r *GetBestEffortCoverageReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
