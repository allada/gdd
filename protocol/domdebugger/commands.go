package domdebugger


import (
    "../shared"
    "../dom"
    "../runtime"
)

type SetDOMBreakpointCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId dom.NodeId `json:"nodeId"`// Identifier of the node to set breakpoint on.
    Type DOMBreakpointType `json:"type"`// Type of the operation to stop upon.
}
type SetDOMBreakpointReturn struct {
}

func (c *SetDOMBreakpointCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetDOMBreakpointCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetDOMBreakpointCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type RemoveDOMBreakpointCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId dom.NodeId `json:"nodeId"`// Identifier of the node to remove breakpoint from.
    Type DOMBreakpointType `json:"type"`// Type of the breakpoint to remove.
}
type RemoveDOMBreakpointReturn struct {
}

func (c *RemoveDOMBreakpointCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RemoveDOMBreakpointCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RemoveDOMBreakpointCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetEventListenerBreakpointCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    EventName string `json:"eventName"`// DOM Event name to stop on (any DOM event will do).
    TargetName *string `json:"targetName,omitempty"`// [Experimental] EventTarget interface name to stop on. If equal to <code>"*"</code> or not provided, will stop on any EventTarget.
}
type SetEventListenerBreakpointReturn struct {
}

func (c *SetEventListenerBreakpointCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetEventListenerBreakpointCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetEventListenerBreakpointCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type RemoveEventListenerBreakpointCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    EventName string `json:"eventName"`// Event name.
    TargetName *string `json:"targetName,omitempty"`// [Experimental] EventTarget interface name.
}
type RemoveEventListenerBreakpointReturn struct {
}

func (c *RemoveEventListenerBreakpointCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RemoveEventListenerBreakpointCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RemoveEventListenerBreakpointCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetInstrumentationBreakpointCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    EventName string `json:"eventName"`// Instrumentation name to stop on.
}
type SetInstrumentationBreakpointReturn struct {
}

func (c *SetInstrumentationBreakpointCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetInstrumentationBreakpointCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetInstrumentationBreakpointCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type RemoveInstrumentationBreakpointCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    EventName string `json:"eventName"`// Instrumentation name to stop on.
}
type RemoveInstrumentationBreakpointReturn struct {
}

func (c *RemoveInstrumentationBreakpointCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RemoveInstrumentationBreakpointCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RemoveInstrumentationBreakpointCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetXHRBreakpointCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Url string `json:"url"`// Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
}
type SetXHRBreakpointReturn struct {
}

func (c *SetXHRBreakpointCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetXHRBreakpointCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetXHRBreakpointCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type RemoveXHRBreakpointCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Url string `json:"url"`// Resource URL substring.
}
type RemoveXHRBreakpointReturn struct {
}

func (c *RemoveXHRBreakpointCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RemoveXHRBreakpointCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RemoveXHRBreakpointCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetEventListenersCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ObjectId runtime.RemoteObjectId `json:"objectId"`// Identifier of the object to return listeners for.
    Depth *int64 `json:"depth,omitempty"`// [Experimental] The maximum depth at which Node children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
    Pierce *bool `json:"pierce,omitempty"`// [Experimental] Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false). Reports listeners for all contexts if pierce is enabled.
}
type GetEventListenersReturn struct {
    Listeners []EventListener `json:"listeners"`// Array of relevant listeners.
}

func (c *GetEventListenersCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetEventListenersCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetEventListenersCommand) Respond(r *GetEventListenersReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
