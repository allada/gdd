package tracing


import (
    "../shared"
)

type StartCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Categories *string `json:"categories,omitempty"`// Category/tag filter
    Options *string `json:"options,omitempty"`// Tracing options
    BufferUsageReportingInterval *float64 `json:"bufferUsageReportingInterval,omitempty"`// If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
    TransferMode *StartTransferModeEnum `json:"transferMode,omitempty"`// Whether to report trace events as series of dataCollected events or to save trace to a stream (defaults to <code>ReportEvents</code>).
    TraceConfig *TraceConfig `json:"traceConfig,omitempty"`
}
type StartReturn struct {
}

func (c *StartCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StartCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type EndCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type EndReturn struct {
}

func (c *EndCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *EndCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetCategoriesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type GetCategoriesReturn struct {
    Categories []string `json:"categories"`// A list of supported tracing categories.
}

func (c *GetCategoriesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetCategoriesCommand) Respond(r *GetCategoriesReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type RequestMemoryDumpCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type RequestMemoryDumpReturn struct {
    DumpGuid string `json:"dumpGuid"`// GUID of the resulting global memory dump.
    Success bool `json:"success"`// True iff the global memory dump succeeded.
}

func (c *RequestMemoryDumpCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RequestMemoryDumpCommand) Respond(r *RequestMemoryDumpReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type RecordClockSyncMarkerCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    SyncId string `json:"syncId"`// The ID of this clock sync marker
}
type RecordClockSyncMarkerReturn struct {
}

func (c *RecordClockSyncMarkerCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RecordClockSyncMarkerCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}
