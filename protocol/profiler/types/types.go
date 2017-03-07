package types


import (
    "../runtime/types"
)
type ProfileNode struct {
    Id int64 `json:"id"`// Unique id of the node.
    CallFrame runtime.CallFrame `json:"callFrame"`// Function location.
    HitCount *int64 `json:"hitCount,omitempty"`// [Experimental] Number of samples where this node was on top of the call stack.
    Children *[]int64 `json:"children,omitempty"`// Child node ids.
    DeoptReason *string `json:"deoptReason,omitempty"`// The reason of being not optimized. The function may be deoptimized or marked as don't optimize.
    PositionTicks *[]PositionTickInfo `json:"positionTicks,omitempty"`// [Experimental] An array of source position ticks.
}

type Profile struct {
    Nodes []ProfileNode `json:"nodes"`// The list of profile nodes. First item is the root node.
    StartTime float64 `json:"startTime"`// Profiling start timestamp in microseconds.
    EndTime float64 `json:"endTime"`// Profiling end timestamp in microseconds.
    Samples *[]int64 `json:"samples,omitempty"`// Ids of samples top nodes.
    TimeDeltas *[]int64 `json:"timeDeltas,omitempty"`// Time intervals between adjacent samples in microseconds. The first delta is relative to the profile startTime.
}

type PositionTickInfo struct {
    Line int64 `json:"line"`// Source line number (1-based).
    Ticks int64 `json:"ticks"`// Number of samples attributed to the source line.
}

type CoverageRange struct {
    StartLineNumber int64 `json:"startLineNumber"`// JavaScript script line number (0-based) for the range start.
    StartColumnNumber int64 `json:"startColumnNumber"`// JavaScript script column number (0-based) for the range start.
    EndLineNumber int64 `json:"endLineNumber"`// JavaScript script line number (0-based) for the range end.
    EndColumnNumber int64 `json:"endColumnNumber"`// JavaScript script column number (0-based) for the range end.
    Count int64 `json:"count"`// Collected execution count of the source range.
}

type FunctionCoverage struct {
    FunctionName string `json:"functionName"`// JavaScript function name.
    Ranges []CoverageRange `json:"ranges"`// Source ranges inside the function with coverage data.
}

type ScriptCoverage struct {
    ScriptId runtime.ScriptId `json:"scriptId"`// JavaScript script id.
    Url string `json:"url"`// JavaScript script name or url.
    Functions []FunctionCoverage `json:"functions"`// Functions contained in the script that has coverage data.
}
