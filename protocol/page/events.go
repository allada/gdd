package page


import (
    "../dom"
)

type DomContentEventFiredEvent struct {
    Timestamp float64 `json:"timestamp"`
}
type LoadEventFiredEvent struct {
    Timestamp float64 `json:"timestamp"`
}
type FrameAttachedEvent struct {
    FrameId FrameId `json:"frameId"`// Id of the frame that has been attached.
    ParentFrameId FrameId `json:"parentFrameId"`// Parent frame identifier.
}
type FrameNavigatedEvent struct {
    Frame Frame `json:"frame"`// Frame object.
}
type FrameDetachedEvent struct {
    FrameId FrameId `json:"frameId"`// Id of the frame that has been detached.
}
type FrameStartedLoadingEvent struct {
    FrameId FrameId `json:"frameId"`// Id of the frame that has started loading.
}
type FrameStoppedLoadingEvent struct {
    FrameId FrameId `json:"frameId"`// Id of the frame that has stopped loading.
}
type FrameScheduledNavigationEvent struct {
    FrameId FrameId `json:"frameId"`// Id of the frame that has scheduled a navigation.
    Delay float64 `json:"delay"`// Delay (in seconds) until the navigation is scheduled to begin. The navigation is not guaranteed to start.
}
type FrameClearedScheduledNavigationEvent struct {
    FrameId FrameId `json:"frameId"`// Id of the frame that has cleared its scheduled navigation.
}
type FrameResizedEvent struct {
}
type JavascriptDialogOpeningEvent struct {
    Message string `json:"message"`// Message that will be displayed by the dialog.
    Type DialogType `json:"type"`// Dialog type.
}
type JavascriptDialogClosedEvent struct {
    Result bool `json:"result"`// Whether dialog was confirmed.
}
type ScreencastFrameEvent struct {
    Data string `json:"data"`// Base64-encoded compressed image.
    Metadata ScreencastFrameMetadata `json:"metadata"`// Screencast frame metadata.
    SessionId int64 `json:"sessionId"`// Frame number.
}
type ScreencastVisibilityChangedEvent struct {
    Visible bool `json:"visible"`// True if the page is visible.
}
type ColorPickedEvent struct {
    Color dom.RGBA `json:"color"`// RGBA of the picked color.
}
type InterstitialShownEvent struct {
}
type InterstitialHiddenEvent struct {
}
type NavigationRequestedEvent struct {
    IsInMainFrame bool `json:"isInMainFrame"`// Whether the navigation is taking place in the main frame or in a subframe.
    IsRedirect bool `json:"isRedirect"`// Whether the navigation has encountered a server redirect or not.
    NavigationId int64 `json:"navigationId"`
    Url string `json:"url"`// URL of requested navigation.
}