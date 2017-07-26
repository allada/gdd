package heapprofiler


import (
    "github.com/allada/gdd/protocol/runtime"
)
type HeapSnapshotObjectId string

type SamplingHeapProfileNode struct {
    CallFrame runtime.CallFrame `json:"callFrame"`// Function location.
    SelfSize float64 `json:"selfSize"`// Allocations size in bytes for the node excluding children.
    Children []SamplingHeapProfileNode `json:"children"`// Child nodes.
}

type SamplingHeapProfile struct {
    Head SamplingHeapProfileNode `json:"head"`
}
