package types

type TouchPointStateEnum string
const (
    TouchPointStateTouchPressed TouchPointStateEnum = "touchPressed"
    TouchPointStateTouchReleased TouchPointStateEnum = "touchReleased"
    TouchPointStateTouchMoved TouchPointStateEnum = "touchMoved"
    TouchPointStateTouchStationary TouchPointStateEnum = "touchStationary"
    TouchPointStateTouchCancelled TouchPointStateEnum = "touchCancelled"
)

const (
    GestureSourceTypeDefault GestureSourceType = "default"
    GestureSourceTypeTouch GestureSourceType = "touch"
    GestureSourceTypeMouse GestureSourceType = "mouse"
)

type DispatchKeyEventTypeEnum string
const (
    DispatchKeyEventTypeKeyDown DispatchKeyEventTypeEnum = "keyDown"
    DispatchKeyEventTypeKeyUp DispatchKeyEventTypeEnum = "keyUp"
    DispatchKeyEventTypeRawKeyDown DispatchKeyEventTypeEnum = "rawKeyDown"
    DispatchKeyEventTypeChar DispatchKeyEventTypeEnum = "char"
)

type DispatchMouseEventTypeEnum string
const (
    DispatchMouseEventTypeMousePressed DispatchMouseEventTypeEnum = "mousePressed"
    DispatchMouseEventTypeMouseReleased DispatchMouseEventTypeEnum = "mouseReleased"
    DispatchMouseEventTypeMouseMoved DispatchMouseEventTypeEnum = "mouseMoved"
)

type DispatchMouseEventButtonEnum string
const (
    DispatchMouseEventButtonNone DispatchMouseEventButtonEnum = "none"
    DispatchMouseEventButtonLeft DispatchMouseEventButtonEnum = "left"
    DispatchMouseEventButtonMiddle DispatchMouseEventButtonEnum = "middle"
    DispatchMouseEventButtonRight DispatchMouseEventButtonEnum = "right"
)

type DispatchTouchEventTypeEnum string
const (
    DispatchTouchEventTypeTouchStart DispatchTouchEventTypeEnum = "touchStart"
    DispatchTouchEventTypeTouchEnd DispatchTouchEventTypeEnum = "touchEnd"
    DispatchTouchEventTypeTouchMove DispatchTouchEventTypeEnum = "touchMove"
)

type EmulateTouchFromMouseEventTypeEnum string
const (
    EmulateTouchFromMouseEventTypeMousePressed EmulateTouchFromMouseEventTypeEnum = "mousePressed"
    EmulateTouchFromMouseEventTypeMouseReleased EmulateTouchFromMouseEventTypeEnum = "mouseReleased"
    EmulateTouchFromMouseEventTypeMouseMoved EmulateTouchFromMouseEventTypeEnum = "mouseMoved"
    EmulateTouchFromMouseEventTypeMouseWheel EmulateTouchFromMouseEventTypeEnum = "mouseWheel"
)

type EmulateTouchFromMouseEventButtonEnum string
const (
    EmulateTouchFromMouseEventButtonNone EmulateTouchFromMouseEventButtonEnum = "none"
    EmulateTouchFromMouseEventButtonLeft EmulateTouchFromMouseEventButtonEnum = "left"
    EmulateTouchFromMouseEventButtonMiddle EmulateTouchFromMouseEventButtonEnum = "middle"
    EmulateTouchFromMouseEventButtonRight EmulateTouchFromMouseEventButtonEnum = "right"
)
type TouchPoint struct {
    State TouchPointStateEnum `json:"state"`// State of the touch point.
    X int64 `json:"x"`// X coordinate of the event relative to the main frame's viewport.
    Y int64 `json:"y"`// Y coordinate of the event relative to the main frame's viewport. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
    RadiusX *int64 `json:"radiusX,omitempty"`// X radius of the touch area (default: 1).
    RadiusY *int64 `json:"radiusY,omitempty"`// Y radius of the touch area (default: 1).
    RotationAngle *float64 `json:"rotationAngle,omitempty"`// Rotation angle (default: 0.0).
    Force *float64 `json:"force,omitempty"`// Force (default: 1.0).
    Id *float64 `json:"id,omitempty"`// Identifier used to track touch sources between events, must be unique within an event.
}

type GestureSourceType string
