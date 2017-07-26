package page


import (
    "github.com/allada/gdd/protocol/shared"
    "github.com/allada/gdd/protocol/debugger"
    "github.com/allada/gdd/protocol/emulation"
    "github.com/allada/gdd/protocol/dom"
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

type AddScriptToEvaluateOnLoadCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ScriptSource string `json:"scriptSource"`
}
type AddScriptToEvaluateOnLoadReturn struct {
    Identifier ScriptIdentifier `json:"identifier"`// Identifier of the added script.
}

func (c *AddScriptToEvaluateOnLoadCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *AddScriptToEvaluateOnLoadCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *AddScriptToEvaluateOnLoadCommand) Respond(r *AddScriptToEvaluateOnLoadReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type RemoveScriptToEvaluateOnLoadCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Identifier ScriptIdentifier `json:"identifier"`
}
type RemoveScriptToEvaluateOnLoadReturn struct {
}

func (c *RemoveScriptToEvaluateOnLoadCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RemoveScriptToEvaluateOnLoadCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RemoveScriptToEvaluateOnLoadCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetAutoAttachToCreatedPagesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    AutoAttach bool `json:"autoAttach"`// If true, browser will open a new inspector window for every page created from this one.
}
type SetAutoAttachToCreatedPagesReturn struct {
}

func (c *SetAutoAttachToCreatedPagesCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetAutoAttachToCreatedPagesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetAutoAttachToCreatedPagesCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ReloadCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    IgnoreCache *bool `json:"ignoreCache,omitempty"`// If true, browser cache is ignored (as if the user pressed Shift+refresh).
    ScriptToEvaluateOnLoad *string `json:"scriptToEvaluateOnLoad,omitempty"`// If set, the script will be injected into all frames of the inspected page after reload.
}
type ReloadReturn struct {
}

func (c *ReloadCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ReloadCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ReloadCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type NavigateCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Url string `json:"url"`// URL to navigate the page to.
    Referrer *string `json:"referrer,omitempty"`// [Experimental] Referrer URL.
}
type NavigateReturn struct {
    FrameId FrameId `json:"frameId"`// [Experimental] Frame id that will be navigated.
}

func (c *NavigateCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *NavigateCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *NavigateCommand) Respond(r *NavigateReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type StopLoadingCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type StopLoadingReturn struct {
}

func (c *StopLoadingCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *StopLoadingCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StopLoadingCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetNavigationHistoryCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type GetNavigationHistoryReturn struct {
    CurrentIndex int64 `json:"currentIndex"`// Index of the current navigation history entry.
    Entries []NavigationEntry `json:"entries"`// Array of navigation history entries.
}

func (c *GetNavigationHistoryCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetNavigationHistoryCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetNavigationHistoryCommand) Respond(r *GetNavigationHistoryReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type NavigateToHistoryEntryCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    EntryId int64 `json:"entryId"`// Unique id of the entry to navigate to.
}
type NavigateToHistoryEntryReturn struct {
}

func (c *NavigateToHistoryEntryCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *NavigateToHistoryEntryCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *NavigateToHistoryEntryCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
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

type GetResourceTreeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type GetResourceTreeReturn struct {
    FrameTree FrameResourceTree `json:"frameTree"`// Present frame / resource tree structure.
}

func (c *GetResourceTreeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetResourceTreeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetResourceTreeCommand) Respond(r *GetResourceTreeReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetResourceContentCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    FrameId FrameId `json:"frameId"`// Frame id to get resource for.
    Url string `json:"url"`// URL of the resource to get content for.
}
type GetResourceContentReturn struct {
    Content string `json:"content"`// Resource content.
    Base64Encoded bool `json:"base64Encoded"`// True, if content was served as base64.
}

func (c *GetResourceContentCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetResourceContentCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetResourceContentCommand) Respond(r *GetResourceContentReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SearchInResourceCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    FrameId FrameId `json:"frameId"`// Frame id for resource to search in.
    Url string `json:"url"`// URL of the resource to search in.
    Query string `json:"query"`// String to search for.
    CaseSensitive *bool `json:"caseSensitive,omitempty"`// If true, search is case sensitive.
    IsRegex *bool `json:"isRegex,omitempty"`// If true, treats string parameter as regex.
}
type SearchInResourceReturn struct {
    Result []debugger.SearchMatch `json:"result"`// List of search matches.
}

func (c *SearchInResourceCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SearchInResourceCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SearchInResourceCommand) Respond(r *SearchInResourceReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetDocumentContentCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    FrameId FrameId `json:"frameId"`// Frame id to set HTML for.
    Html string `json:"html"`// HTML content to set.
}
type SetDocumentContentReturn struct {
}

func (c *SetDocumentContentCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetDocumentContentCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetDocumentContentCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetDeviceMetricsOverrideCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Width int64 `json:"width"`// Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
    Height int64 `json:"height"`// Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
    DeviceScaleFactor float64 `json:"deviceScaleFactor"`// Overriding device scale factor value. 0 disables the override.
    Mobile bool `json:"mobile"`// Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
    FitWindow bool `json:"fitWindow"`// Whether a view that exceeds the available browser window area should be scaled down to fit.
    Scale *float64 `json:"scale,omitempty"`// Scale to apply to resulting view image. Ignored in |fitWindow| mode.
    OffsetX *float64 `json:"offsetX,omitempty"`// X offset to shift resulting view image by. Ignored in |fitWindow| mode.
    OffsetY *float64 `json:"offsetY,omitempty"`// Y offset to shift resulting view image by. Ignored in |fitWindow| mode.
    ScreenWidth *int64 `json:"screenWidth,omitempty"`// Overriding screen width value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
    ScreenHeight *int64 `json:"screenHeight,omitempty"`// Overriding screen height value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
    PositionX *int64 `json:"positionX,omitempty"`// Overriding view X position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
    PositionY *int64 `json:"positionY,omitempty"`// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
    ScreenOrientation *emulation.ScreenOrientation `json:"screenOrientation,omitempty"`// Screen orientation override.
}
type SetDeviceMetricsOverrideReturn struct {
}

func (c *SetDeviceMetricsOverrideCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetDeviceMetricsOverrideCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetDeviceMetricsOverrideCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ClearDeviceMetricsOverrideCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type ClearDeviceMetricsOverrideReturn struct {
}

func (c *ClearDeviceMetricsOverrideCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ClearDeviceMetricsOverrideCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ClearDeviceMetricsOverrideCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetGeolocationOverrideCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Latitude *float64 `json:"latitude,omitempty"`// Mock latitude
    Longitude *float64 `json:"longitude,omitempty"`// Mock longitude
    Accuracy *float64 `json:"accuracy,omitempty"`// Mock accuracy
}
type SetGeolocationOverrideReturn struct {
}

func (c *SetGeolocationOverrideCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetGeolocationOverrideCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetGeolocationOverrideCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ClearGeolocationOverrideCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type ClearGeolocationOverrideReturn struct {
}

func (c *ClearGeolocationOverrideCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ClearGeolocationOverrideCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ClearGeolocationOverrideCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetDeviceOrientationOverrideCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Alpha float64 `json:"alpha"`// Mock alpha
    Beta float64 `json:"beta"`// Mock beta
    Gamma float64 `json:"gamma"`// Mock gamma
}
type SetDeviceOrientationOverrideReturn struct {
}

func (c *SetDeviceOrientationOverrideCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetDeviceOrientationOverrideCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetDeviceOrientationOverrideCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ClearDeviceOrientationOverrideCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type ClearDeviceOrientationOverrideReturn struct {
}

func (c *ClearDeviceOrientationOverrideCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ClearDeviceOrientationOverrideCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ClearDeviceOrientationOverrideCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetTouchEmulationEnabledCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Enabled bool `json:"enabled"`// Whether the touch event emulation should be enabled.
    Configuration *SetTouchEmulationEnabledConfigurationEnum `json:"configuration,omitempty"`// Touch/gesture events configuration. Default: current platform.
}
type SetTouchEmulationEnabledReturn struct {
}

func (c *SetTouchEmulationEnabledCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetTouchEmulationEnabledCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetTouchEmulationEnabledCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type CaptureScreenshotCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Format *CaptureScreenshotFormatEnum `json:"format,omitempty"`// Image compression format (defaults to png).
    Quality *int64 `json:"quality,omitempty"`// Compression quality from range [0..100] (jpeg only).
}
type CaptureScreenshotReturn struct {
    Data string `json:"data"`// Base64-encoded image data.
}

func (c *CaptureScreenshotCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *CaptureScreenshotCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CaptureScreenshotCommand) Respond(r *CaptureScreenshotReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type PrintToPDFCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type PrintToPDFReturn struct {
    Data string `json:"data"`// Base64-encoded pdf data.
}

func (c *PrintToPDFCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *PrintToPDFCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *PrintToPDFCommand) Respond(r *PrintToPDFReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type StartScreencastCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Format *StartScreencastFormatEnum `json:"format,omitempty"`// Image compression format.
    Quality *int64 `json:"quality,omitempty"`// Compression quality from range [0..100].
    MaxWidth *int64 `json:"maxWidth,omitempty"`// Maximum screenshot width.
    MaxHeight *int64 `json:"maxHeight,omitempty"`// Maximum screenshot height.
    EveryNthFrame *int64 `json:"everyNthFrame,omitempty"`// Send every n-th frame.
}
type StartScreencastReturn struct {
}

func (c *StartScreencastCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *StartScreencastCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StartScreencastCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type StopScreencastCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type StopScreencastReturn struct {
}

func (c *StopScreencastCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *StopScreencastCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StopScreencastCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ScreencastFrameAckCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    SessionId int64 `json:"sessionId"`// Frame number.
}
type ScreencastFrameAckReturn struct {
}

func (c *ScreencastFrameAckCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ScreencastFrameAckCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ScreencastFrameAckCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type HandleJavaScriptDialogCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Accept bool `json:"accept"`// Whether to accept or dismiss the dialog.
    PromptText *string `json:"promptText,omitempty"`// The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
}
type HandleJavaScriptDialogReturn struct {
}

func (c *HandleJavaScriptDialogCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *HandleJavaScriptDialogCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *HandleJavaScriptDialogCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetColorPickerEnabledCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Enabled bool `json:"enabled"`// Shows / hides color picker
}
type SetColorPickerEnabledReturn struct {
}

func (c *SetColorPickerEnabledCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetColorPickerEnabledCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetColorPickerEnabledCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ConfigureOverlayCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Suspended *bool `json:"suspended,omitempty"`// Whether overlay should be suspended and not consume any resources.
    Message *string `json:"message,omitempty"`// Overlay message to display.
}
type ConfigureOverlayReturn struct {
}

func (c *ConfigureOverlayCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ConfigureOverlayCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ConfigureOverlayCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetAppManifestCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type GetAppManifestReturn struct {
    Url string `json:"url"`// Manifest location.
    Errors []AppManifestError `json:"errors"`
    Data *string `json:"data,omitempty"`// Manifest content.
}

func (c *GetAppManifestCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetAppManifestCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetAppManifestCommand) Respond(r *GetAppManifestReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type RequestAppBannerCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type RequestAppBannerReturn struct {
}

func (c *RequestAppBannerCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RequestAppBannerCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RequestAppBannerCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetControlNavigationsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Enabled bool `json:"enabled"`
}
type SetControlNavigationsReturn struct {
}

func (c *SetControlNavigationsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetControlNavigationsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetControlNavigationsCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ProcessNavigationCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Response NavigationResponse `json:"response"`
    NavigationId int64 `json:"navigationId"`
}
type ProcessNavigationReturn struct {
}

func (c *ProcessNavigationCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ProcessNavigationCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ProcessNavigationCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetLayoutMetricsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type GetLayoutMetricsReturn struct {
    LayoutViewport LayoutViewport `json:"layoutViewport"`// Metrics relating to the layout viewport.
    VisualViewport VisualViewport `json:"visualViewport"`// Metrics relating to the visual viewport.
    ContentSize dom.Rect `json:"contentSize"`// Size of scrollable area.
}

func (c *GetLayoutMetricsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetLayoutMetricsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetLayoutMetricsCommand) Respond(r *GetLayoutMetricsReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
