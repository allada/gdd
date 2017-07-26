package tracing


import (
    "github.com/allada/gdd/protocol/io"
)

type DataCollectedEvent struct {
    Value []map[string]string `json:"value"`
}
type TracingCompleteEvent struct {
    Stream *io.StreamHandle `json:"stream,omitempty"`// A handle of the stream that holds resulting trace data.
}
type BufferUsageEvent struct {
    PercentFull *float64 `json:"percentFull,omitempty"`// A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
    EventCount *float64 `json:"eventCount,omitempty"`// An approximate number of events in the trace log.
    Value *float64 `json:"value,omitempty"`// A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
}