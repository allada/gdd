package emulation


import (
    "../shared"
    "../dom"
)

type SetDeviceMetricsOverrideCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Width int64 `json:"width"`// Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
    Height int64 `json:"height"`// Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
    DeviceScaleFactor float64 `json:"deviceScaleFactor"`// Overriding device scale factor value. 0 disables the override.
    Mobile bool `json:"mobile"`// Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
    FitWindow bool `json:"fitWindow"`// Whether a view that exceeds the available browser window area should be scaled down to fit.
    Scale *float64 `json:"scale,omitempty"`// [Experimental] Scale to apply to resulting view image. Ignored in |fitWindow| mode.
    OffsetX *float64 `json:"offsetX,omitempty"`// [Experimental] Not used.
    OffsetY *float64 `json:"offsetY,omitempty"`// [Experimental] Not used.
    ScreenWidth *int64 `json:"screenWidth,omitempty"`// [Experimental] Overriding screen width value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
    ScreenHeight *int64 `json:"screenHeight,omitempty"`// [Experimental] Overriding screen height value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
    PositionX *int64 `json:"positionX,omitempty"`// [Experimental] Overriding view X position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
    PositionY *int64 `json:"positionY,omitempty"`// [Experimental] Overriding view Y position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
    ScreenOrientation *ScreenOrientation `json:"screenOrientation,omitempty"`// Screen orientation override.
}
type SetDeviceMetricsOverrideReturn struct {
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

func (c *ClearDeviceMetricsOverrideCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ClearDeviceMetricsOverrideCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ForceViewportCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    X float64 `json:"x"`// X coordinate of top-left corner of the area (CSS pixels).
    Y float64 `json:"y"`// Y coordinate of top-left corner of the area (CSS pixels).
    Scale float64 `json:"scale"`// Scale to apply to the area (relative to a page scale of 1.0).
}
type ForceViewportReturn struct {
}

func (c *ForceViewportCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ForceViewportCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ResetViewportCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type ResetViewportReturn struct {
}

func (c *ResetViewportCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ResetViewportCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ResetPageScaleFactorCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type ResetPageScaleFactorReturn struct {
}

func (c *ResetPageScaleFactorCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ResetPageScaleFactorCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetPageScaleFactorCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    PageScaleFactor float64 `json:"pageScaleFactor"`// Page scale factor.
}
type SetPageScaleFactorReturn struct {
}

func (c *SetPageScaleFactorCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetPageScaleFactorCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetVisibleSizeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Width int64 `json:"width"`// Frame width (DIP).
    Height int64 `json:"height"`// Frame height (DIP).
}
type SetVisibleSizeReturn struct {
}

func (c *SetVisibleSizeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetVisibleSizeCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetScriptExecutionDisabledCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Value bool `json:"value"`// Whether script execution should be disabled in the page.
}
type SetScriptExecutionDisabledReturn struct {
}

func (c *SetScriptExecutionDisabledCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetScriptExecutionDisabledCommand) Respond() {
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

func (c *ClearGeolocationOverrideCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ClearGeolocationOverrideCommand) Respond() {
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

func (c *SetTouchEmulationEnabledCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetTouchEmulationEnabledCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetEmulatedMediaCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Media string `json:"media"`// Media type to emulate. Empty string disables the override.
}
type SetEmulatedMediaReturn struct {
}

func (c *SetEmulatedMediaCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetEmulatedMediaCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetCPUThrottlingRateCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Rate float64 `json:"rate"`// Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
}
type SetCPUThrottlingRateReturn struct {
}

func (c *SetCPUThrottlingRateCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetCPUThrottlingRateCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type CanEmulateCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type CanEmulateReturn struct {
    Result bool `json:"result"`// True if emulation is supported.
}

func (c *CanEmulateCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CanEmulateCommand) Respond(r *CanEmulateReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetVirtualTimePolicyCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Policy VirtualTimePolicy `json:"policy"`
    Budget *int64 `json:"budget,omitempty"`// If set, after this many virtual milliseconds have elapsed virtual time will be paused and a virtualTimeBudgetExpired event is sent.
}
type SetVirtualTimePolicyReturn struct {
}

func (c *SetVirtualTimePolicyCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetVirtualTimePolicyCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetDefaultBackgroundColorOverrideCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Color *dom.RGBA `json:"color,omitempty"`// RGBA of the default background color. If not specified, any existing override will be cleared.
}
type SetDefaultBackgroundColorOverrideReturn struct {
}

func (c *SetDefaultBackgroundColorOverrideCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetDefaultBackgroundColorOverrideCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}
