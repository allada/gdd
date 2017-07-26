package rendering


import (
    "github.com/allada/gdd/protocol/shared"
)

type SetShowPaintRectsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Result bool `json:"result"`// True for showing paint rectangles
}
type SetShowPaintRectsReturn struct {
}

func (c *SetShowPaintRectsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetShowPaintRectsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetShowPaintRectsCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetShowDebugBordersCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Show bool `json:"show"`// True for showing debug borders
}
type SetShowDebugBordersReturn struct {
}

func (c *SetShowDebugBordersCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetShowDebugBordersCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetShowDebugBordersCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetShowFPSCounterCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Show bool `json:"show"`// True for showing the FPS counter
}
type SetShowFPSCounterReturn struct {
}

func (c *SetShowFPSCounterCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetShowFPSCounterCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetShowFPSCounterCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetShowScrollBottleneckRectsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Show bool `json:"show"`// True for showing scroll bottleneck rects
}
type SetShowScrollBottleneckRectsReturn struct {
}

func (c *SetShowScrollBottleneckRectsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetShowScrollBottleneckRectsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetShowScrollBottleneckRectsCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetShowViewportSizeOnResizeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Show bool `json:"show"`// Whether to paint size or not.
}
type SetShowViewportSizeOnResizeReturn struct {
}

func (c *SetShowViewportSizeOnResizeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetShowViewportSizeOnResizeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetShowViewportSizeOnResizeCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}
