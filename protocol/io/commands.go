package io


import (
    "../shared"
)

type ReadCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Handle StreamHandle `json:"handle"`// Handle of the stream to read.
    Offset *int64 `json:"offset,omitempty"`// Seek to the specified offset before reading (if not specificed, proceed with offset following the last read).
    Size *int64 `json:"size,omitempty"`// Maximum number of bytes to read (left upon the agent discretion if not specified).
}
type ReadReturn struct {
    Data string `json:"data"`// Data that were read.
    Eof bool `json:"eof"`// Set if the end-of-file condition occured while reading.
}

func (c *ReadCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ReadCommand) Respond(r *ReadReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type CloseCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Handle StreamHandle `json:"handle"`// Handle of the stream to close.
}
type CloseReturn struct {
}

func (c *CloseCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CloseCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}
