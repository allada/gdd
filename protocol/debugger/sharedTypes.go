package debugger


import (
    "github.com/allada/gdd/protocol/runtime"
)
type PausedReasonEnum string
const (
    PausedReasonXHR PausedReasonEnum = "XHR"
    PausedReasonDOM PausedReasonEnum = "DOM"
    PausedReasonEventListener PausedReasonEnum = "EventListener"
    PausedReasonException PausedReasonEnum = "exception"
    PausedReasonAssert PausedReasonEnum = "assert"
    PausedReasonDebugCommand PausedReasonEnum = "debugCommand"
    PausedReasonPromiseRejection PausedReasonEnum = "promiseRejection"
    PausedReasonOOM PausedReasonEnum = "OOM"
    PausedReasonOther PausedReasonEnum = "other"
    PausedReasonAmbiguous PausedReasonEnum = "ambiguous"
)

type ScopeTypeEnum string
const (
    ScopeTypeGlobal ScopeTypeEnum = "global"
    ScopeTypeLocal ScopeTypeEnum = "local"
    ScopeTypeWith ScopeTypeEnum = "with"
    ScopeTypeClosure ScopeTypeEnum = "closure"
    ScopeTypeCatch ScopeTypeEnum = "catch"
    ScopeTypeBlock ScopeTypeEnum = "block"
    ScopeTypeScript ScopeTypeEnum = "script"
    ScopeTypeEval ScopeTypeEnum = "eval"
    ScopeTypeModule ScopeTypeEnum = "module"
)

type SetPauseOnExceptionsStateEnum string
const (
    SetPauseOnExceptionsStateNone SetPauseOnExceptionsStateEnum = "none"
    SetPauseOnExceptionsStateUncaught SetPauseOnExceptionsStateEnum = "uncaught"
    SetPauseOnExceptionsStateAll SetPauseOnExceptionsStateEnum = "all"
)
type BreakpointId string

type CallFrameId string

type Location struct {
    ScriptId runtime.ScriptId `json:"scriptId"`// Script identifier as reported in the <code>Debugger.scriptParsed</code>.
    LineNumber int64 `json:"lineNumber"`// Line number in the script (0-based).
    ColumnNumber *int64 `json:"columnNumber,omitempty"`// Column number in the script (0-based).
}

type ScriptPosition struct {
    LineNumber int64 `json:"lineNumber"`
    ColumnNumber int64 `json:"columnNumber"`
}

type CallFrame struct {
    CallFrameId CallFrameId `json:"callFrameId"`// Call frame identifier. This identifier is only valid while the virtual machine is paused.
    FunctionName string `json:"functionName"`// Name of the JavaScript function called on this call frame.
    FunctionLocation *Location `json:"functionLocation,omitempty"`// [Experimental] Location in the source code.
    Location Location `json:"location"`// Location in the source code.
    ScopeChain []Scope `json:"scopeChain"`// Scope chain for this call frame.
    This runtime.RemoteObject `json:"this"`// <code>this</code> object for this call frame.
    ReturnValue *runtime.RemoteObject `json:"returnValue,omitempty"`// The value being returned, if the function is at return point.
}

type Scope struct {
    Type ScopeTypeEnum `json:"type"`// Scope type.
    Object runtime.RemoteObject `json:"object"`// Object representing the scope. For <code>global</code> and <code>with</code> scopes it represents the actual object; for the rest of the scopes, it is artificial transient object enumerating scope variables as its properties.
    Name *string `json:"name,omitempty"`
    StartLocation *Location `json:"startLocation,omitempty"`// Location in the source code where scope starts
    EndLocation *Location `json:"endLocation,omitempty"`// Location in the source code where scope ends
}

type SearchMatch struct {
    LineNumber float64 `json:"lineNumber"`// Line number in resource content.
    LineContent string `json:"lineContent"`// Line with match content.
}
