package types


import (
    "../network/types"
)
const (
    ResourceTypeDocument ResourceType = "Document"
    ResourceTypeStylesheet ResourceType = "Stylesheet"
    ResourceTypeImage ResourceType = "Image"
    ResourceTypeMedia ResourceType = "Media"
    ResourceTypeFont ResourceType = "Font"
    ResourceTypeScript ResourceType = "Script"
    ResourceTypeTextTrack ResourceType = "TextTrack"
    ResourceTypeXHR ResourceType = "XHR"
    ResourceTypeFetch ResourceType = "Fetch"
    ResourceTypeEventSource ResourceType = "EventSource"
    ResourceTypeWebSocket ResourceType = "WebSocket"
    ResourceTypeManifest ResourceType = "Manifest"
    ResourceTypeOther ResourceType = "Other"
)

const (
    DialogTypeAlert DialogType = "alert"
    DialogTypeConfirm DialogType = "confirm"
    DialogTypePrompt DialogType = "prompt"
    DialogTypeBeforeunload DialogType = "beforeunload"
)

const (
    NavigationResponseProceed NavigationResponse = "Proceed"
    NavigationResponseCancel NavigationResponse = "Cancel"
    NavigationResponseCancelAndIgnore NavigationResponse = "CancelAndIgnore"
)

type SetTouchEmulationEnabledConfigurationEnum string
const (
    SetTouchEmulationEnabledConfigurationMobile SetTouchEmulationEnabledConfigurationEnum = "mobile"
    SetTouchEmulationEnabledConfigurationDesktop SetTouchEmulationEnabledConfigurationEnum = "desktop"
)

type CaptureScreenshotFormatEnum string
const (
    CaptureScreenshotFormatJpeg CaptureScreenshotFormatEnum = "jpeg"
    CaptureScreenshotFormatPng CaptureScreenshotFormatEnum = "png"
)

type StartScreencastFormatEnum string
const (
    StartScreencastFormatJpeg StartScreencastFormatEnum = "jpeg"
    StartScreencastFormatPng StartScreencastFormatEnum = "png"
)
type ResourceType string

type FrameId string

type Frame struct {
    Id string `json:"id"`// Frame unique identifier.
    ParentId *string `json:"parentId,omitempty"`// Parent frame identifier.
    LoaderId network.LoaderId `json:"loaderId"`// Identifier of the loader associated with this frame.
    Name *string `json:"name,omitempty"`// Frame's name as specified in the tag.
    Url string `json:"url"`// Frame document's URL.
    SecurityOrigin string `json:"securityOrigin"`// Frame document's security origin.
    MimeType string `json:"mimeType"`// Frame document's mimeType as determined by the browser.
}

type FrameResource struct {
    Url string `json:"url"`// Resource URL.
    Type ResourceType `json:"type"`// Type of this resource.
    MimeType string `json:"mimeType"`// Resource mimeType as determined by the browser.
    LastModified *network.Timestamp `json:"lastModified,omitempty"`// last-modified timestamp as reported by server.
    ContentSize *float64 `json:"contentSize,omitempty"`// Resource content size.
    Failed *bool `json:"failed,omitempty"`// True if the resource failed to load.
    Canceled *bool `json:"canceled,omitempty"`// True if the resource was canceled during loading.
}

type FrameResourceTree struct {
    Frame Frame `json:"frame"`// Frame information for this tree item.
    ChildFrames *[]FrameResourceTree `json:"childFrames,omitempty"`// Child frames.
    Resources []FrameResource `json:"resources"`// Information about frame resources.
}

type ScriptIdentifier string

type NavigationEntry struct {
    Id int64 `json:"id"`// Unique id of the navigation history entry.
    Url string `json:"url"`// URL of the navigation history entry.
    Title string `json:"title"`// Title of the navigation history entry.
}

type ScreencastFrameMetadata struct {
    OffsetTop float64 `json:"offsetTop"`// [Experimental] Top offset in DIP.
    PageScaleFactor float64 `json:"pageScaleFactor"`// [Experimental] Page scale factor.
    DeviceWidth float64 `json:"deviceWidth"`// [Experimental] Device screen width in DIP.
    DeviceHeight float64 `json:"deviceHeight"`// [Experimental] Device screen height in DIP.
    ScrollOffsetX float64 `json:"scrollOffsetX"`// [Experimental] Position of horizontal scroll in CSS pixels.
    ScrollOffsetY float64 `json:"scrollOffsetY"`// [Experimental] Position of vertical scroll in CSS pixels.
    Timestamp *float64 `json:"timestamp,omitempty"`// [Experimental] Frame swap timestamp.
}

type DialogType string

type AppManifestError struct {
    Message string `json:"message"`// Error message.
    Critical int64 `json:"critical"`// If criticial, this is a non-recoverable parse error.
    Line int64 `json:"line"`// Error line.
    Column int64 `json:"column"`// Error column.
}

type NavigationResponse string

type LayoutViewport struct {
    PageX int64 `json:"pageX"`// Horizontal offset relative to the document (CSS pixels).
    PageY int64 `json:"pageY"`// Vertical offset relative to the document (CSS pixels).
    ClientWidth int64 `json:"clientWidth"`// Width (CSS pixels), excludes scrollbar if present.
    ClientHeight int64 `json:"clientHeight"`// Height (CSS pixels), excludes scrollbar if present.
}

type VisualViewport struct {
    OffsetX float64 `json:"offsetX"`// Horizontal offset relative to the layout viewport (CSS pixels).
    OffsetY float64 `json:"offsetY"`// Vertical offset relative to the layout viewport (CSS pixels).
    PageX float64 `json:"pageX"`// Horizontal offset relative to the document (CSS pixels).
    PageY float64 `json:"pageY"`// Vertical offset relative to the document (CSS pixels).
    ClientWidth float64 `json:"clientWidth"`// Width (CSS pixels), excludes scrollbar if present.
    ClientHeight float64 `json:"clientHeight"`// Height (CSS pixels), excludes scrollbar if present.
    Scale float64 `json:"scale"`// Scale relative to the ideal viewport (size at width=device-width).
}
