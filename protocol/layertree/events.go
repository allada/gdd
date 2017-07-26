package layertree


import (
    "github.com/allada/gdd/protocol/dom"
)

type LayerTreeDidChangeEvent struct {
    Layers *[]Layer `json:"layers,omitempty"`// Layer tree, absent if not in the comspositing mode.
}
type LayerPaintedEvent struct {
    LayerId LayerId `json:"layerId"`// The id of the painted layer.
    Clip dom.Rect `json:"clip"`// Clip rectangle.
}