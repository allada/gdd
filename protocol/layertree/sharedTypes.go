package layertree


import (
    "../dom"
)
type ScrollRectTypeEnum string
const (
    ScrollRectTypeRepaintsOnScroll ScrollRectTypeEnum = "RepaintsOnScroll"
    ScrollRectTypeTouchEventHandler ScrollRectTypeEnum = "TouchEventHandler"
    ScrollRectTypeWheelEventHandler ScrollRectTypeEnum = "WheelEventHandler"
)
type LayerId string

type SnapshotId string

type ScrollRect struct {
    Rect dom.Rect `json:"rect"`// Rectangle itself.
    Type ScrollRectTypeEnum `json:"type"`// Reason for rectangle to force scrolling on the main thread
}

type PictureTile struct {
    X float64 `json:"x"`// Offset from owning layer left boundary
    Y float64 `json:"y"`// Offset from owning layer top boundary
    Picture string `json:"picture"`// Base64-encoded snapshot data.
}

type Layer struct {
    LayerId LayerId `json:"layerId"`// The unique id for this layer.
    ParentLayerId *LayerId `json:"parentLayerId,omitempty"`// The id of parent (not present for root).
    BackendNodeId *dom.BackendNodeId `json:"backendNodeId,omitempty"`// The backend id for the node associated with this layer.
    OffsetX float64 `json:"offsetX"`// Offset from parent layer, X coordinate.
    OffsetY float64 `json:"offsetY"`// Offset from parent layer, Y coordinate.
    Width float64 `json:"width"`// Layer width.
    Height float64 `json:"height"`// Layer height.
    Transform *[]float64 `json:"transform,omitempty"`// Transformation matrix for layer, default is identity matrix
    AnchorX *float64 `json:"anchorX,omitempty"`// Transform anchor point X, absent if no transform specified
    AnchorY *float64 `json:"anchorY,omitempty"`// Transform anchor point Y, absent if no transform specified
    AnchorZ *float64 `json:"anchorZ,omitempty"`// Transform anchor point Z, absent if no transform specified
    PaintCount int64 `json:"paintCount"`// Indicates how many time this layer has painted.
    DrawsContent bool `json:"drawsContent"`// Indicates whether this layer hosts any content, rather than being used for transform/scrolling purposes only.
    Invisible *bool `json:"invisible,omitempty"`// Set if layer is not visible.
    ScrollRects *[]ScrollRect `json:"scrollRects,omitempty"`// Rectangles scrolling on main thread only.
}

type PaintProfile []float64
