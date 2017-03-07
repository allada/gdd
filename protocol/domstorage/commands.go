package domstorage


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

func (c *DisableCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DisableCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ClearCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    StorageId StorageId `json:"storageId"`
}
type ClearReturn struct {
}

func (c *ClearCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ClearCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetDOMStorageItemsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    StorageId StorageId `json:"storageId"`
}
type GetDOMStorageItemsReturn struct {
    Entries []Item `json:"entries"`
}

func (c *GetDOMStorageItemsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetDOMStorageItemsCommand) Respond(r *GetDOMStorageItemsReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetDOMStorageItemCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    StorageId StorageId `json:"storageId"`
    Key string `json:"key"`
    Value string `json:"value"`
}
type SetDOMStorageItemReturn struct {
}

func (c *SetDOMStorageItemCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetDOMStorageItemCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type RemoveDOMStorageItemCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    StorageId StorageId `json:"storageId"`
    Key string `json:"key"`
}
type RemoveDOMStorageItemReturn struct {
}

func (c *RemoveDOMStorageItemCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RemoveDOMStorageItemCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}
