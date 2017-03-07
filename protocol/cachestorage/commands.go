package cachestorage


import (
    "../shared"
)

type RequestCacheNamesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    SecurityOrigin string `json:"securityOrigin"`// Security origin.
}
type RequestCacheNamesReturn struct {
    Caches []Cache `json:"caches"`// Caches for the security origin.
}

func (c *RequestCacheNamesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RequestCacheNamesCommand) Respond(r *RequestCacheNamesReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type RequestEntriesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    CacheId CacheId `json:"cacheId"`// ID of cache to get entries from.
    SkipCount int64 `json:"skipCount"`// Number of records to skip.
    PageSize int64 `json:"pageSize"`// Number of records to fetch.
}
type RequestEntriesReturn struct {
    CacheDataEntries []DataEntry `json:"cacheDataEntries"`// Array of object store data entries.
    HasMore bool `json:"hasMore"`// If true, there are more entries to fetch in the given range.
}

func (c *RequestEntriesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RequestEntriesCommand) Respond(r *RequestEntriesReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type DeleteCacheCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    CacheId CacheId `json:"cacheId"`// Id of cache for deletion.
}
type DeleteCacheReturn struct {
}

func (c *DeleteCacheCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DeleteCacheCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type DeleteEntryCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    CacheId CacheId `json:"cacheId"`// Id of cache where the entry will be deleted.
    Request string `json:"request"`// URL spec of the request.
}
type DeleteEntryReturn struct {
}

func (c *DeleteEntryCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DeleteEntryCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}
