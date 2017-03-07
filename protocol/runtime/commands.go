package runtime


import (
    "../shared"
)

type EvaluateCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Expression string `json:"expression"`// Expression to evaluate.
    ObjectGroup *string `json:"objectGroup,omitempty"`// Symbolic group name that can be used to release multiple objects.
    IncludeCommandLineAPI *bool `json:"includeCommandLineAPI,omitempty"`// Determines whether Command Line API should be available during the evaluation.
    Silent *bool `json:"silent,omitempty"`// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
    ContextId *ExecutionContextId `json:"contextId,omitempty"`// Specifies in which execution context to perform evaluation. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
    ReturnByValue *bool `json:"returnByValue,omitempty"`// Whether the result is expected to be a JSON object that should be sent by value.
    GeneratePreview *bool `json:"generatePreview,omitempty"`// [Experimental] Whether preview should be generated for the result.
    UserGesture *bool `json:"userGesture,omitempty"`// [Experimental] Whether execution should be treated as initiated by user in the UI.
    AwaitPromise *bool `json:"awaitPromise,omitempty"`// Whether execution should wait for promise to be resolved. If the result of evaluation is not a Promise, it's considered to be an error.
}
type EvaluateReturn struct {
    Result RemoteObject `json:"result"`// Evaluation result.
    ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`// Exception details.
}

func (c *EvaluateCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *EvaluateCommand) Respond(r *EvaluateReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type AwaitPromiseCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    PromiseObjectId RemoteObjectId `json:"promiseObjectId"`// Identifier of the promise.
    ReturnByValue *bool `json:"returnByValue,omitempty"`// Whether the result is expected to be a JSON object that should be sent by value.
    GeneratePreview *bool `json:"generatePreview,omitempty"`// Whether preview should be generated for the result.
}
type AwaitPromiseReturn struct {
    Result RemoteObject `json:"result"`// Promise result. Will contain rejected value if promise was rejected.
    ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`// Exception details if stack strace is available.
}

func (c *AwaitPromiseCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *AwaitPromiseCommand) Respond(r *AwaitPromiseReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type CallFunctionOnCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ObjectId RemoteObjectId `json:"objectId"`// Identifier of the object to call function on.
    FunctionDeclaration string `json:"functionDeclaration"`// Declaration of the function to call.
    Arguments *[]CallArgument `json:"arguments,omitempty"`// Call arguments. All call arguments must belong to the same JavaScript world as the target object.
    Silent *bool `json:"silent,omitempty"`// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
    ReturnByValue *bool `json:"returnByValue,omitempty"`// Whether the result is expected to be a JSON object which should be sent by value.
    GeneratePreview *bool `json:"generatePreview,omitempty"`// [Experimental] Whether preview should be generated for the result.
    UserGesture *bool `json:"userGesture,omitempty"`// [Experimental] Whether execution should be treated as initiated by user in the UI.
    AwaitPromise *bool `json:"awaitPromise,omitempty"`// Whether execution should wait for promise to be resolved. If the result of evaluation is not a Promise, it's considered to be an error.
}
type CallFunctionOnReturn struct {
    Result RemoteObject `json:"result"`// Call result.
    ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`// Exception details.
}

func (c *CallFunctionOnCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CallFunctionOnCommand) Respond(r *CallFunctionOnReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetPropertiesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ObjectId RemoteObjectId `json:"objectId"`// Identifier of the object to return properties for.
    OwnProperties *bool `json:"ownProperties,omitempty"`// If true, returns properties belonging only to the element itself, not to its prototype chain.
    AccessorPropertiesOnly *bool `json:"accessorPropertiesOnly,omitempty"`// [Experimental] If true, returns accessor properties (with getter/setter) only; internal properties are not returned either.
    GeneratePreview *bool `json:"generatePreview,omitempty"`// [Experimental] Whether preview should be generated for the results.
}
type GetPropertiesReturn struct {
    Result []PropertyDescriptor `json:"result"`// Object properties.
    InternalProperties *[]InternalPropertyDescriptor `json:"internalProperties,omitempty"`// Internal object properties (only of the element itself).
    ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`// Exception details.
}

func (c *GetPropertiesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetPropertiesCommand) Respond(r *GetPropertiesReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type ReleaseObjectCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ObjectId RemoteObjectId `json:"objectId"`// Identifier of the object to release.
}
type ReleaseObjectReturn struct {
}

func (c *ReleaseObjectCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ReleaseObjectCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ReleaseObjectGroupCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ObjectGroup string `json:"objectGroup"`// Symbolic object group name.
}
type ReleaseObjectGroupReturn struct {
}

func (c *ReleaseObjectGroupCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ReleaseObjectGroupCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type RunIfWaitingForDebuggerCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type RunIfWaitingForDebuggerReturn struct {
}

func (c *RunIfWaitingForDebuggerCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RunIfWaitingForDebuggerCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type EnableCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type EnableReturn struct {
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

func (c *DisableCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DisableCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type DiscardConsoleEntriesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type DiscardConsoleEntriesReturn struct {
}

func (c *DiscardConsoleEntriesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DiscardConsoleEntriesCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetCustomObjectFormatterEnabledCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Enabled bool `json:"enabled"`
}
type SetCustomObjectFormatterEnabledReturn struct {
}

func (c *SetCustomObjectFormatterEnabledCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetCustomObjectFormatterEnabledCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type CompileScriptCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Expression string `json:"expression"`// Expression to compile.
    SourceURL string `json:"sourceURL"`// Source url to be set for the script.
    PersistScript bool `json:"persistScript"`// Specifies whether the compiled script should be persisted.
    ExecutionContextId *ExecutionContextId `json:"executionContextId,omitempty"`// Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
}
type CompileScriptReturn struct {
    ScriptId *ScriptId `json:"scriptId,omitempty"`// Id of the script.
    ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`// Exception details.
}

func (c *CompileScriptCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CompileScriptCommand) Respond(r *CompileScriptReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type RunScriptCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ScriptId ScriptId `json:"scriptId"`// Id of the script to run.
    ExecutionContextId *ExecutionContextId `json:"executionContextId,omitempty"`// Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
    ObjectGroup *string `json:"objectGroup,omitempty"`// Symbolic group name that can be used to release multiple objects.
    Silent *bool `json:"silent,omitempty"`// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
    IncludeCommandLineAPI *bool `json:"includeCommandLineAPI,omitempty"`// Determines whether Command Line API should be available during the evaluation.
    ReturnByValue *bool `json:"returnByValue,omitempty"`// Whether the result is expected to be a JSON object which should be sent by value.
    GeneratePreview *bool `json:"generatePreview,omitempty"`// Whether preview should be generated for the result.
    AwaitPromise *bool `json:"awaitPromise,omitempty"`// Whether execution should wait for promise to be resolved. If the result of evaluation is not a Promise, it's considered to be an error.
}
type RunScriptReturn struct {
    Result RemoteObject `json:"result"`// Run result.
    ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"`// Exception details.
}

func (c *RunScriptCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RunScriptCommand) Respond(r *RunScriptReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
