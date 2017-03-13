package network


import (
    "../shared"
)

type EnableCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    MaxTotalBufferSize *int64 `json:"maxTotalBufferSize,omitempty"`// [Experimental] Buffer size in bytes to use when preserving network payloads (XHRs, etc).
    MaxResourceBufferSize *int64 `json:"maxResourceBufferSize,omitempty"`// [Experimental] Per-resource buffer size in bytes to use when preserving network payloads (XHRs, etc).
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

type SetUserAgentOverrideCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    UserAgent string `json:"userAgent"`// User agent to use.
}
type SetUserAgentOverrideReturn struct {
}

func (c *SetUserAgentOverrideCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetUserAgentOverrideCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetUserAgentOverrideCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetExtraHTTPHeadersCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Headers Headers `json:"headers"`// Map with extra HTTP headers.
}
type SetExtraHTTPHeadersReturn struct {
}

func (c *SetExtraHTTPHeadersCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetExtraHTTPHeadersCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetExtraHTTPHeadersCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetResponseBodyCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    RequestId RequestId `json:"requestId"`// Identifier of the network request to get content for.
}
type GetResponseBodyReturn struct {
    Body string `json:"body"`// Response body.
    Base64Encoded bool `json:"base64Encoded"`// True, if content was sent as base64.
}

func (c *GetResponseBodyCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetResponseBodyCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetResponseBodyCommand) Respond(r *GetResponseBodyReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type AddBlockedURLCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Url string `json:"url"`// URL to block.
}
type AddBlockedURLReturn struct {
}

func (c *AddBlockedURLCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *AddBlockedURLCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *AddBlockedURLCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type RemoveBlockedURLCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Url string `json:"url"`// URL to stop blocking.
}
type RemoveBlockedURLReturn struct {
}

func (c *RemoveBlockedURLCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RemoveBlockedURLCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RemoveBlockedURLCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ReplayXHRCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    RequestId RequestId `json:"requestId"`// Identifier of XHR to replay.
}
type ReplayXHRReturn struct {
}

func (c *ReplayXHRCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ReplayXHRCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ReplayXHRCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetMonitoringXHREnabledCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Enabled bool `json:"enabled"`// Monitoring enabled state.
}
type SetMonitoringXHREnabledReturn struct {
}

func (c *SetMonitoringXHREnabledCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetMonitoringXHREnabledCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetMonitoringXHREnabledCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type CanClearBrowserCacheCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type CanClearBrowserCacheReturn struct {
    Result bool `json:"result"`// True if browser cache can be cleared.
}

func (c *CanClearBrowserCacheCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *CanClearBrowserCacheCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CanClearBrowserCacheCommand) Respond(r *CanClearBrowserCacheReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type ClearBrowserCacheCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type ClearBrowserCacheReturn struct {
}

func (c *ClearBrowserCacheCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ClearBrowserCacheCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ClearBrowserCacheCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type CanClearBrowserCookiesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type CanClearBrowserCookiesReturn struct {
    Result bool `json:"result"`// True if browser cookies can be cleared.
}

func (c *CanClearBrowserCookiesCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *CanClearBrowserCookiesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CanClearBrowserCookiesCommand) Respond(r *CanClearBrowserCookiesReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type ClearBrowserCookiesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type ClearBrowserCookiesReturn struct {
}

func (c *ClearBrowserCookiesCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ClearBrowserCookiesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ClearBrowserCookiesCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetCookiesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Urls *[]string `json:"urls,omitempty"`// The list of URLs for which applicable cookies will be fetched
}
type GetCookiesReturn struct {
    Cookies []Cookie `json:"cookies"`// Array of cookie objects.
}

func (c *GetCookiesCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetCookiesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetCookiesCommand) Respond(r *GetCookiesReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetAllCookiesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type GetAllCookiesReturn struct {
    Cookies []Cookie `json:"cookies"`// Array of cookie objects.
}

func (c *GetAllCookiesCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetAllCookiesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetAllCookiesCommand) Respond(r *GetAllCookiesReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type DeleteCookieCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    CookieName string `json:"cookieName"`// Name of the cookie to remove.
    Url string `json:"url"`// URL to match cooke domain and path.
}
type DeleteCookieReturn struct {
}

func (c *DeleteCookieCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *DeleteCookieCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DeleteCookieCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetCookieCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Url string `json:"url"`// The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
    Name string `json:"name"`// The name of the cookie.
    Value string `json:"value"`// The value of the cookie.
    Domain *string `json:"domain,omitempty"`// If omitted, the cookie becomes a host-only cookie.
    Path *string `json:"path,omitempty"`// Defaults to the path portion of the url parameter.
    Secure *bool `json:"secure,omitempty"`// Defaults ot false.
    HttpOnly *bool `json:"httpOnly,omitempty"`// Defaults to false.
    SameSite *CookieSameSite `json:"sameSite,omitempty"`// Defaults to browser default behavior.
    ExpirationDate *Timestamp `json:"expirationDate,omitempty"`// If omitted, the cookie becomes a session cookie.
}
type SetCookieReturn struct {
    Success bool `json:"success"`// True if successfully set cookie.
}

func (c *SetCookieCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetCookieCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetCookieCommand) Respond(r *SetCookieReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type CanEmulateNetworkConditionsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type CanEmulateNetworkConditionsReturn struct {
    Result bool `json:"result"`// True if emulation of network conditions is supported.
}

func (c *CanEmulateNetworkConditionsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *CanEmulateNetworkConditionsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CanEmulateNetworkConditionsCommand) Respond(r *CanEmulateNetworkConditionsReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type EmulateNetworkConditionsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Offline bool `json:"offline"`// True to emulate internet disconnection.
    Latency float64 `json:"latency"`// Additional latency (ms).
    DownloadThroughput float64 `json:"downloadThroughput"`// Maximal aggregated download throughput.
    UploadThroughput float64 `json:"uploadThroughput"`// Maximal aggregated upload throughput.
    ConnectionType *ConnectionType `json:"connectionType,omitempty"`// Connection type if known.
}
type EmulateNetworkConditionsReturn struct {
}

func (c *EmulateNetworkConditionsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *EmulateNetworkConditionsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *EmulateNetworkConditionsCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetCacheDisabledCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    CacheDisabled bool `json:"cacheDisabled"`// Cache disabled state.
}
type SetCacheDisabledReturn struct {
}

func (c *SetCacheDisabledCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetCacheDisabledCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetCacheDisabledCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetBypassServiceWorkerCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Bypass bool `json:"bypass"`// Bypass service worker and load from network.
}
type SetBypassServiceWorkerReturn struct {
}

func (c *SetBypassServiceWorkerCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetBypassServiceWorkerCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetBypassServiceWorkerCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetDataSizeLimitsForTestCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    MaxTotalSize int64 `json:"maxTotalSize"`// Maximum total buffer size.
    MaxResourceSize int64 `json:"maxResourceSize"`// Maximum per-resource size.
}
type SetDataSizeLimitsForTestReturn struct {
}

func (c *SetDataSizeLimitsForTestCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetDataSizeLimitsForTestCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetDataSizeLimitsForTestCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetCertificateCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Origin string `json:"origin"`// Origin to get certificate for.
}
type GetCertificateReturn struct {
    TableNames []string `json:"tableNames"`
}

func (c *GetCertificateCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetCertificateCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetCertificateCommand) Respond(r *GetCertificateReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
