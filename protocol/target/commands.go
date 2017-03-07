package target


import (
    "../shared"
)

type SetDiscoverTargetsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Discover bool `json:"discover"`// Whether to discover available targets.
}
type SetDiscoverTargetsReturn struct {
}

func (c *SetDiscoverTargetsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetDiscoverTargetsCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetAutoAttachCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    AutoAttach bool `json:"autoAttach"`// Whether to auto-attach to related targets.
    WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"`// Whether to pause new targets when attaching to them. Use <code>Runtime.runIfWaitingForDebugger</code> to run paused targets.
}
type SetAutoAttachReturn struct {
}

func (c *SetAutoAttachCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetAutoAttachCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetAttachToFramesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Value bool `json:"value"`// Whether to attach to frames.
}
type SetAttachToFramesReturn struct {
}

func (c *SetAttachToFramesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetAttachToFramesCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetRemoteLocationsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Locations []RemoteLocation `json:"locations"`// List of remote locations.
}
type SetRemoteLocationsReturn struct {
}

func (c *SetRemoteLocationsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetRemoteLocationsCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SendMessageToTargetCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    TargetId string `json:"targetId"`
    Message string `json:"message"`
}
type SendMessageToTargetReturn struct {
}

func (c *SendMessageToTargetCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SendMessageToTargetCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetTargetInfoCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    TargetId TargetID `json:"targetId"`
}
type GetTargetInfoReturn struct {
    TargetInfo TargetInfo `json:"targetInfo"`
}

func (c *GetTargetInfoCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetTargetInfoCommand) Respond(r *GetTargetInfoReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type ActivateTargetCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    TargetId TargetID `json:"targetId"`
}
type ActivateTargetReturn struct {
}

func (c *ActivateTargetCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ActivateTargetCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type CloseTargetCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    TargetId TargetID `json:"targetId"`
}
type CloseTargetReturn struct {
    Success bool `json:"success"`
}

func (c *CloseTargetCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CloseTargetCommand) Respond(r *CloseTargetReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type AttachToTargetCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    TargetId TargetID `json:"targetId"`
}
type AttachToTargetReturn struct {
    Success bool `json:"success"`// Whether attach succeeded.
}

func (c *AttachToTargetCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *AttachToTargetCommand) Respond(r *AttachToTargetReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type DetachFromTargetCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    TargetId TargetID `json:"targetId"`
}
type DetachFromTargetReturn struct {
}

func (c *DetachFromTargetCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DetachFromTargetCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type CreateBrowserContextCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type CreateBrowserContextReturn struct {
    BrowserContextId BrowserContextID `json:"browserContextId"`// The id of the context created.
}

func (c *CreateBrowserContextCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CreateBrowserContextCommand) Respond(r *CreateBrowserContextReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type DisposeBrowserContextCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    BrowserContextId BrowserContextID `json:"browserContextId"`
}
type DisposeBrowserContextReturn struct {
    Success bool `json:"success"`
}

func (c *DisposeBrowserContextCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DisposeBrowserContextCommand) Respond(r *DisposeBrowserContextReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type CreateTargetCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Url string `json:"url"`// The initial URL the page will be navigated to.
    Width *int64 `json:"width,omitempty"`// Frame width in DIP (headless chrome only).
    Height *int64 `json:"height,omitempty"`// Frame height in DIP (headless chrome only).
    BrowserContextId *BrowserContextID `json:"browserContextId,omitempty"`// The browser context to create the page in (headless chrome only).
}
type CreateTargetReturn struct {
    TargetId TargetID `json:"targetId"`// The id of the page opened.
}

func (c *CreateTargetCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CreateTargetCommand) Respond(r *CreateTargetReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetTargetsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type GetTargetsReturn struct {
    TargetInfos []TargetInfo `json:"targetInfos"`// The list of targets.
}

func (c *GetTargetsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetTargetsCommand) Respond(r *GetTargetsReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
