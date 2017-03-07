package debugger


import (
    "../runtime"
)

type ScriptParsedEvent struct {
    ScriptId runtime.ScriptId `json:"scriptId"`// Identifier of the script parsed.
    Url string `json:"url"`// URL or name of the script parsed (if any).
    StartLine int64 `json:"startLine"`// Line offset of the script within the resource with given URL (for script tags).
    StartColumn int64 `json:"startColumn"`// Column offset of the script within the resource with given URL.
    EndLine int64 `json:"endLine"`// Last line of the script.
    EndColumn int64 `json:"endColumn"`// Length of the last line of the script.
    ExecutionContextId runtime.ExecutionContextId `json:"executionContextId"`// Specifies script creation context.
    Hash string `json:"hash"`// Content hash of the script.
    ExecutionContextAuxData *map[string]string `json:"executionContextAuxData,omitempty"`// Embedder-specific auxiliary data.
    IsLiveEdit *bool `json:"isLiveEdit,omitempty"`// [Experimental] True, if this script is generated as a result of the live edit operation.
    SourceMapURL *string `json:"sourceMapURL,omitempty"`// URL of source map associated with script (if any).
    HasSourceURL *bool `json:"hasSourceURL,omitempty"`// [Experimental] True, if this script has sourceURL.
    IsModule *bool `json:"isModule,omitempty"`// [Experimental] True, if this script is ES6 module.
}
type ScriptFailedToParseEvent struct {
    ScriptId runtime.ScriptId `json:"scriptId"`// Identifier of the script parsed.
    Url string `json:"url"`// URL or name of the script parsed (if any).
    StartLine int64 `json:"startLine"`// Line offset of the script within the resource with given URL (for script tags).
    StartColumn int64 `json:"startColumn"`// Column offset of the script within the resource with given URL.
    EndLine int64 `json:"endLine"`// Last line of the script.
    EndColumn int64 `json:"endColumn"`// Length of the last line of the script.
    ExecutionContextId runtime.ExecutionContextId `json:"executionContextId"`// Specifies script creation context.
    Hash string `json:"hash"`// Content hash of the script.
    ExecutionContextAuxData *map[string]string `json:"executionContextAuxData,omitempty"`// Embedder-specific auxiliary data.
    SourceMapURL *string `json:"sourceMapURL,omitempty"`// URL of source map associated with script (if any).
    HasSourceURL *bool `json:"hasSourceURL,omitempty"`// [Experimental] True, if this script has sourceURL.
    IsModule *bool `json:"isModule,omitempty"`// [Experimental] True, if this script is ES6 module.
}
type BreakpointResolvedEvent struct {
    BreakpointId BreakpointId `json:"breakpointId"`// Breakpoint unique identifier.
    Location Location `json:"location"`// Actual breakpoint location.
}
type PausedEvent struct {
    CallFrames []CallFrame `json:"callFrames"`// Call stack the virtual machine stopped on.
    Reason PausedReasonEnum `json:"reason"`// Pause reason.
    Data *map[string]string `json:"data,omitempty"`// Object containing break-specific auxiliary properties.
    HitBreakpoints *[]string `json:"hitBreakpoints,omitempty"`// Hit breakpoints IDs
    AsyncStackTrace *runtime.StackTrace `json:"asyncStackTrace,omitempty"`// Async stack trace, if any.
}
type ResumedEvent struct {
}