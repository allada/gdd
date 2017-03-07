package types


import (
    "../runtime/types"
    "../dom/types"
)
const (
    DOMBreakpointTypeSubtreeDashmodified DOMBreakpointType = "subtree-modified"
    DOMBreakpointTypeAttributeDashmodified DOMBreakpointType = "attribute-modified"
    DOMBreakpointTypeNodeDashremoved DOMBreakpointType = "node-removed"
)
type DOMBreakpointType string

type EventListener struct {
    Type string `json:"type"`// <code>EventListener</code>'s type.
    UseCapture bool `json:"useCapture"`// <code>EventListener</code>'s useCapture.
    Passive bool `json:"passive"`// <code>EventListener</code>'s passive flag.
    Once bool `json:"once"`// <code>EventListener</code>'s once flag.
    ScriptId runtime.ScriptId `json:"scriptId"`// Script id of the handler code.
    LineNumber int64 `json:"lineNumber"`// Line number in the script (0-based).
    ColumnNumber int64 `json:"columnNumber"`// Column number in the script (0-based).
    Handler *runtime.RemoteObject `json:"handler,omitempty"`// Event handler function value.
    OriginalHandler *runtime.RemoteObject `json:"originalHandler,omitempty"`// Event original handler function value.
    BackendNodeId *dom.BackendNodeId `json:"backendNodeId,omitempty"`// Node the listener is added to (if any).
}
