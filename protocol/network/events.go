package network


import (
    "../page"
)

type ResourceChangedPriorityEvent struct {
    RequestId RequestId `json:"requestId"`// Request identifier.
    NewPriority ResourcePriority `json:"newPriority"`// New priority
    Timestamp Timestamp `json:"timestamp"`// Timestamp.
}
type RequestWillBeSentEvent struct {
    RequestId RequestId `json:"requestId"`// Request identifier.
    FrameId page.FrameId `json:"frameId"`// [Experimental] Frame identifier.
    LoaderId LoaderId `json:"loaderId"`// Loader identifier.
    DocumentURL string `json:"documentURL"`// URL of the document this request is loaded for.
    Request Request `json:"request"`// Request data.
    Timestamp Timestamp `json:"timestamp"`// Timestamp.
    WallTime Timestamp `json:"wallTime"`// [Experimental] UTC Timestamp.
    Initiator Initiator `json:"initiator"`// Request initiator.
    RedirectResponse *Response `json:"redirectResponse,omitempty"`// Redirect response data.
    Type *page.ResourceType `json:"type,omitempty"`// [Experimental] Type of this resource.
}
type RequestServedFromCacheEvent struct {
    RequestId RequestId `json:"requestId"`// Request identifier.
}
type ResponseReceivedEvent struct {
    RequestId RequestId `json:"requestId"`// Request identifier.
    FrameId page.FrameId `json:"frameId"`// [Experimental] Frame identifier.
    LoaderId LoaderId `json:"loaderId"`// Loader identifier.
    Timestamp Timestamp `json:"timestamp"`// Timestamp.
    Type page.ResourceType `json:"type"`// Resource type.
    Response Response `json:"response"`// Response data.
}
type DataReceivedEvent struct {
    RequestId RequestId `json:"requestId"`// Request identifier.
    Timestamp Timestamp `json:"timestamp"`// Timestamp.
    DataLength int64 `json:"dataLength"`// Data chunk length.
    EncodedDataLength int64 `json:"encodedDataLength"`// Actual bytes received (might be less than dataLength for compressed encodings).
}
type LoadingFinishedEvent struct {
    RequestId RequestId `json:"requestId"`// Request identifier.
    Timestamp Timestamp `json:"timestamp"`// Timestamp.
    EncodedDataLength float64 `json:"encodedDataLength"`// Total number of bytes received for this request.
}
type LoadingFailedEvent struct {
    RequestId RequestId `json:"requestId"`// Request identifier.
    Timestamp Timestamp `json:"timestamp"`// Timestamp.
    Type page.ResourceType `json:"type"`// Resource type.
    ErrorText string `json:"errorText"`// User friendly error message.
    Canceled *bool `json:"canceled,omitempty"`// True if loading was canceled.
    BlockedReason *BlockedReason `json:"blockedReason,omitempty"`// [Experimental] The reason why loading was blocked, if any.
}
type WebSocketWillSendHandshakeRequestEvent struct {
    RequestId RequestId `json:"requestId"`// Request identifier.
    Timestamp Timestamp `json:"timestamp"`// Timestamp.
    WallTime Timestamp `json:"wallTime"`// [Experimental] UTC Timestamp.
    Request WebSocketRequest `json:"request"`// WebSocket request data.
}
type WebSocketHandshakeResponseReceivedEvent struct {
    RequestId RequestId `json:"requestId"`// Request identifier.
    Timestamp Timestamp `json:"timestamp"`// Timestamp.
    Response WebSocketResponse `json:"response"`// WebSocket response data.
}
type WebSocketCreatedEvent struct {
    RequestId RequestId `json:"requestId"`// Request identifier.
    Url string `json:"url"`// WebSocket request URL.
    Initiator *Initiator `json:"initiator,omitempty"`// Request initiator.
}
type WebSocketClosedEvent struct {
    RequestId RequestId `json:"requestId"`// Request identifier.
    Timestamp Timestamp `json:"timestamp"`// Timestamp.
}
type WebSocketFrameReceivedEvent struct {
    RequestId RequestId `json:"requestId"`// Request identifier.
    Timestamp Timestamp `json:"timestamp"`// Timestamp.
    Response WebSocketFrame `json:"response"`// WebSocket response data.
}
type WebSocketFrameErrorEvent struct {
    RequestId RequestId `json:"requestId"`// Request identifier.
    Timestamp Timestamp `json:"timestamp"`// Timestamp.
    ErrorMessage string `json:"errorMessage"`// WebSocket frame error message.
}
type WebSocketFrameSentEvent struct {
    RequestId RequestId `json:"requestId"`// Request identifier.
    Timestamp Timestamp `json:"timestamp"`// Timestamp.
    Response WebSocketFrame `json:"response"`// WebSocket response data.
}
type EventSourceMessageReceivedEvent struct {
    RequestId RequestId `json:"requestId"`// Request identifier.
    Timestamp Timestamp `json:"timestamp"`// Timestamp.
    EventName string `json:"eventName"`// Message type.
    EventId string `json:"eventId"`// Message identifier.
    Data string `json:"data"`// Message content.
}