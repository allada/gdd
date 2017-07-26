package debugger


import (
    "github.com/allada/gdd/protocol/shared"
    "github.com/allada/gdd/protocol/runtime"
)

type EnableCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type EnableReturn struct {
}

func (c *EnableCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *EnableCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *EnableCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type DisableCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type DisableReturn struct {
}

func (c *DisableCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *DisableCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DisableCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetBreakpointsActiveCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Active bool `json:"active"`// New value for breakpoints active state.
}
type SetBreakpointsActiveReturn struct {
}

func (c *SetBreakpointsActiveCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetBreakpointsActiveCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetBreakpointsActiveCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetSkipAllPausesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Skip bool `json:"skip"`// New value for skip pauses state.
}
type SetSkipAllPausesReturn struct {
}

func (c *SetSkipAllPausesCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetSkipAllPausesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetSkipAllPausesCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetBreakpointByUrlCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    LineNumber int64 `json:"lineNumber"`// Line number to set breakpoint at.
    Url *string `json:"url,omitempty"`// URL of the resources to set breakpoint on.
    UrlRegex *string `json:"urlRegex,omitempty"`// Regex pattern for the URLs of the resources to set breakpoints on. Either <code>url</code> or <code>urlRegex</code> must be specified.
    ColumnNumber *int64 `json:"columnNumber,omitempty"`// Offset in the line to set breakpoint at.
    Condition *string `json:"condition,omitempty"`// Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
}
type SetBreakpointByUrlReturn struct {
    BreakpointId BreakpointId `json:"breakpointId"`// Id of the created breakpoint for further reference.
    Locations []Location `json:"locations"`// List of the locations this breakpoint resolved into upon addition.
}

func (c *SetBreakpointByUrlCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetBreakpointByUrlCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetBreakpointByUrlCommand) Respond(r *SetBreakpointByUrlReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetBreakpointCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Location Location `json:"location"`// Location to set breakpoint in.
    Condition *string `json:"condition,omitempty"`// Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
}
type SetBreakpointReturn struct {
    BreakpointId BreakpointId `json:"breakpointId"`// Id of the created breakpoint for further reference.
    ActualLocation Location `json:"actualLocation"`// Location this breakpoint resolved into.
}

func (c *SetBreakpointCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetBreakpointCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetBreakpointCommand) Respond(r *SetBreakpointReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type RemoveBreakpointCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    BreakpointId BreakpointId `json:"breakpointId"`
}
type RemoveBreakpointReturn struct {
}

func (c *RemoveBreakpointCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RemoveBreakpointCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RemoveBreakpointCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetPossibleBreakpointsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Start Location `json:"start"`// Start of range to search possible breakpoint locations in.
    End *Location `json:"end,omitempty"`// End of range to search possible breakpoint locations in (excluding). When not specifed, end of scripts is used as end of range.
}
type GetPossibleBreakpointsReturn struct {
    Locations []Location `json:"locations"`// List of the possible breakpoint locations.
}

func (c *GetPossibleBreakpointsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetPossibleBreakpointsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetPossibleBreakpointsCommand) Respond(r *GetPossibleBreakpointsReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type ContinueToLocationCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Location Location `json:"location"`// Location to continue to.
}
type ContinueToLocationReturn struct {
}

func (c *ContinueToLocationCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ContinueToLocationCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ContinueToLocationCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type StepOverCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type StepOverReturn struct {
}

func (c *StepOverCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *StepOverCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StepOverCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type StepIntoCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type StepIntoReturn struct {
}

func (c *StepIntoCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *StepIntoCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StepIntoCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type StepOutCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type StepOutReturn struct {
}

func (c *StepOutCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *StepOutCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StepOutCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type PauseCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type PauseReturn struct {
}

func (c *PauseCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *PauseCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *PauseCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ResumeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type ResumeReturn struct {
}

func (c *ResumeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ResumeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ResumeCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SearchInContentCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ScriptId runtime.ScriptId `json:"scriptId"`// Id of the script to search in.
    Query string `json:"query"`// String to search for.
    CaseSensitive *bool `json:"caseSensitive,omitempty"`// If true, search is case sensitive.
    IsRegex *bool `json:"isRegex,omitempty"`// If true, treats string parameter as regex.
}
type SearchInContentReturn struct {
    Result []SearchMatch `json:"result"`// List of search matches.
}

func (c *SearchInContentCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SearchInContentCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SearchInContentCommand) Respond(r *SearchInContentReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetScriptSourceCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ScriptId runtime.ScriptId `json:"scriptId"`// Id of the script to edit.
    ScriptSource string `json:"scriptSource"`// New content of the script.
    DryRun *bool `json:"dryRun,omitempty"`//  If true the change will not actually be applied. Dry run may be used to get result description without actually modifying the code.
}
type SetScriptSourceReturn struct {
    CallFrames *[]CallFrame `json:"callFrames,omitempty"`// New stack trace in case editing has happened while VM was stopped.
    StackChanged *bool `json:"stackChanged,omitempty"`// Whether current call stack  was modified after applying the changes.
    AsyncStackTrace *runtime.StackTrace `json:"asyncStackTrace,omitempty"`// Async stack trace, if any.
    ExceptionDetails *runtime.ExceptionDetails `json:"exceptionDetails,omitempty"`// Exception details if any.
}

func (c *SetScriptSourceCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetScriptSourceCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetScriptSourceCommand) Respond(r *SetScriptSourceReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type RestartFrameCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    CallFrameId CallFrameId `json:"callFrameId"`// Call frame identifier to evaluate on.
}
type RestartFrameReturn struct {
    CallFrames []CallFrame `json:"callFrames"`// New stack trace.
    AsyncStackTrace *runtime.StackTrace `json:"asyncStackTrace,omitempty"`// Async stack trace, if any.
}

func (c *RestartFrameCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RestartFrameCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RestartFrameCommand) Respond(r *RestartFrameReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetScriptSourceCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ScriptId runtime.ScriptId `json:"scriptId"`// Id of the script to get source for.
}
type GetScriptSourceReturn struct {
    ScriptSource string `json:"scriptSource"`// Script source.
}

func (c *GetScriptSourceCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetScriptSourceCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetScriptSourceCommand) Respond(r *GetScriptSourceReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetPauseOnExceptionsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    State SetPauseOnExceptionsStateEnum `json:"state"`// Pause on exceptions mode.
}
type SetPauseOnExceptionsReturn struct {
}

func (c *SetPauseOnExceptionsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetPauseOnExceptionsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetPauseOnExceptionsCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type EvaluateOnCallFrameCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    CallFrameId CallFrameId `json:"callFrameId"`// Call frame identifier to evaluate on.
    Expression string `json:"expression"`// Expression to evaluate.
    ObjectGroup *string `json:"objectGroup,omitempty"`// String object group name to put result into (allows rapid releasing resulting object handles using <code>releaseObjectGroup</code>).
    IncludeCommandLineAPI *bool `json:"includeCommandLineAPI,omitempty"`// Specifies whether command line API should be available to the evaluated expression, defaults to false.
    Silent *bool `json:"silent,omitempty"`// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
    ReturnByValue *bool `json:"returnByValue,omitempty"`// Whether the result is expected to be a JSON object that should be sent by value.
    GeneratePreview *bool `json:"generatePreview,omitempty"`// [Experimental] Whether preview should be generated for the result.
    ThrowOnSideEffect *bool `json:"throwOnSideEffect,omitempty"`// [Experimental] Whether to throw an exception if side effect cannot be ruled out during evaluation.
}
type EvaluateOnCallFrameReturn struct {
    Result runtime.RemoteObject `json:"result"`// Object wrapper for the evaluation result.
    ExceptionDetails *runtime.ExceptionDetails `json:"exceptionDetails,omitempty"`// Exception details.
}

func (c *EvaluateOnCallFrameCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *EvaluateOnCallFrameCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *EvaluateOnCallFrameCommand) Respond(r *EvaluateOnCallFrameReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetVariableValueCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ScopeNumber int64 `json:"scopeNumber"`// 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch' scope types are allowed. Other scopes could be manipulated manually.
    VariableName string `json:"variableName"`// Variable name.
    NewValue runtime.CallArgument `json:"newValue"`// New variable value.
    CallFrameId CallFrameId `json:"callFrameId"`// Id of callframe that holds variable.
}
type SetVariableValueReturn struct {
}

func (c *SetVariableValueCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetVariableValueCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetVariableValueCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetAsyncCallStackDepthCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    MaxDepth int64 `json:"maxDepth"`// Maximum depth of async call stacks. Setting to <code>0</code> will effectively disable collecting async call stacks (default).
}
type SetAsyncCallStackDepthReturn struct {
}

func (c *SetAsyncCallStackDepthCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetAsyncCallStackDepthCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetAsyncCallStackDepthCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetBlackboxPatternsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Patterns []string `json:"patterns"`// Array of regexps that will be used to check script url for blackbox state.
}
type SetBlackboxPatternsReturn struct {
}

func (c *SetBlackboxPatternsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetBlackboxPatternsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetBlackboxPatternsCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetBlackboxedRangesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ScriptId runtime.ScriptId `json:"scriptId"`// Id of the script.
    Positions []ScriptPosition `json:"positions"`
}
type SetBlackboxedRangesReturn struct {
}

func (c *SetBlackboxedRangesCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetBlackboxedRangesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetBlackboxedRangesCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}
