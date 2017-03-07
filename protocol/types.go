package protocol

const (
    ErrorCodeParseError = -32700
    ErrorCodeInvalidRequest = -32600
    ErrorCodeMethodNotFound = -32601
    ErrorCodeInvalidParams = -32602
    ErrorCodeInternalError = -32603
    ErrorCodeServerError = -32000
)

const (
    StatusCodeSuccess = 0
    StatusCodeError = 1
    StatusCodeFallThrough = 2
    StatusCodeAsync = 3
)

type ErrorResponse struct {
    Code int `json:"code"`
    Message string `json:"message"`
}

type Response struct {
    Method string `json:"method"`
    Params interface{} `json:"params"`
}

type ErrorResult struct {
    Id int `json:"id"`
    Status int `json:"status"`
    Error interface{} `json:"error,omitempty"`
}

type Result struct {
    Id int `json:"id,omitempty"`
    Status int `json:"status,omitempty"`
    Result interface{} `json:"result"`
}

type ScriptParsed struct {
    ScriptId string `json:"scriptId"`
    Url string `json:"url"`
    StartLine int `json:"startLine,omitempty"`
    StartColumn int `json:"startColumn,omitempty"`
    EndLine int `json:"endLine,omitempty"`
    EndColumn int `json:"endColumn,omitempty"`
    ExecutionContextId int `json:"executionContextId"`
    Hash string `json:"hash"`
    SourceMapURL string `json:"sourceMapURL,omitempty"`
}

type ExecutionContextDescription struct {
    Id int `json:"id"`
    Origin string `json:"origin"`
    Name string `json:"name"`
    AuxData ExecutionContextDescriptionAuxData `json:"auxData,omitempty"`
}

type ExecutionContextDescriptionAuxData struct {
    IsDefault bool `json:"isDefault"`
    FrameId string `json:"frameId,omitempty"`
}

type CreateExecutionContext struct {
    Context ExecutionContextDescription `json:"context"`
}

type ScriptSourceResult struct {
    ScriptSource string `json:"scriptSource"`
}

type DebuggerLocation struct {
    ScriptId string `json:"scriptId"`
    LineNumber int `json:"lineNumber"`
    ColumnNumber int `json:ColumnNumber,omitempty"`
}

type DebuggerSetBreakpointByUrlResponse struct {
    BreakpointId string `json:"breakpointId,omitempty"`
    ActualLocation []DebuggerLocation `json:"actualLocation"`
}

type RuntimeRemoteObject struct {
    Type string `json:"type"`
    Subtype string `json:"subtype,omitempty"`
    ClassName string `json:"className,omitempty"`
    Value interface{} `json:"value,omitempty"`
    UnserializableValue string `json:"unserializableValue,omitempty"`
    Description string `json:"description,omitempty"`
    ObjectId string `json:"objectId,omitempty"`
}

type DebuggerScope struct {
    Type string `json:"type"`
    Object RuntimeRemoteObject `json:"object"`
    Name string `json:"string,omitempty"`
    StartLocation DebuggerLocation `json:"startLocation"`
    EndLocation DebuggerLocation `json:"endLocation"`
}

type RuntimeCallFrame struct {
    FunctionName string `json:"functionName"`
    ScriptId string `json:"scriptId"`
    Url string `json:"url"`
    LineNumber int `json:"lineNumber"`
    ColumnNumber int `json:"columnNumber"`
}

type DebuggerCallFrame struct {
    CallFrameId string `json:"callFrameId"`
    FunctionName string `json:"functionName"`
    Location DebuggerLocation `json:"location"`
    ScopeChain []DebuggerScope `json:"scopeChain"`
    This RuntimeRemoteObject `json:"runtimeRemoteObject"`
    ReturnValue *RuntimeRemoteObject `json:"returnValue,omitempty"`
}

type RuntimeStackTrace struct {
    Description string `json:"description"`
    CallFrames []RuntimeCallFrame `json:"callFrames"`
    Parent *RuntimeStackTrace `json:parent"`
}

type DebuggerPausedEvent struct {
    CallFrames []DebuggerCallFrame `json:"callFrames"`
    Reason string `json:"reason"`
    Data interface{} `json:"data,omitempty"`
    HitBreakpoints []string `json:"hitBreakpoints,omitempty"`
    AsyncStackTrace []RuntimeStackTrace `json:"asyncStackTrace,omitempty"`
}

type TargetAttachedToTargetEvent struct {
    TargetInfo TargetTargetInfo`json:"targetInfo"`
}

type TargetTargetInfo struct {
    ID string `json:"target"`
    Type string `json:"type"`
    Title string `json:"title"`
    Url string `json:"url"`
}
