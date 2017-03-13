package runtime

type ConsoleAPICalledTypeEnum string
const (
    ConsoleAPICalledTypeLog ConsoleAPICalledTypeEnum = "log"
    ConsoleAPICalledTypeDebug ConsoleAPICalledTypeEnum = "debug"
    ConsoleAPICalledTypeInfo ConsoleAPICalledTypeEnum = "info"
    ConsoleAPICalledTypeError ConsoleAPICalledTypeEnum = "error"
    ConsoleAPICalledTypeWarning ConsoleAPICalledTypeEnum = "warning"
    ConsoleAPICalledTypeDir ConsoleAPICalledTypeEnum = "dir"
    ConsoleAPICalledTypeDirxml ConsoleAPICalledTypeEnum = "dirxml"
    ConsoleAPICalledTypeTable ConsoleAPICalledTypeEnum = "table"
    ConsoleAPICalledTypeTrace ConsoleAPICalledTypeEnum = "trace"
    ConsoleAPICalledTypeClear ConsoleAPICalledTypeEnum = "clear"
    ConsoleAPICalledTypeStartGroup ConsoleAPICalledTypeEnum = "startGroup"
    ConsoleAPICalledTypeStartGroupCollapsed ConsoleAPICalledTypeEnum = "startGroupCollapsed"
    ConsoleAPICalledTypeEndGroup ConsoleAPICalledTypeEnum = "endGroup"
    ConsoleAPICalledTypeAssert ConsoleAPICalledTypeEnum = "assert"
    ConsoleAPICalledTypeProfile ConsoleAPICalledTypeEnum = "profile"
    ConsoleAPICalledTypeProfileEnd ConsoleAPICalledTypeEnum = "profileEnd"
    ConsoleAPICalledTypeCount ConsoleAPICalledTypeEnum = "count"
    ConsoleAPICalledTypeTimeEnd ConsoleAPICalledTypeEnum = "timeEnd"
)

const (
    ResourceTypeDocument ResourceType = "Document"
    ResourceTypeStylesheet ResourceType = "Stylesheet"
    ResourceTypeImage ResourceType = "Image"
    ResourceTypeMedia ResourceType = "Media"
    ResourceTypeFont ResourceType = "Font"
    ResourceTypeScript ResourceType = "Script"
    ResourceTypeTextTrack ResourceType = "TextTrack"
    ResourceTypeXHR ResourceType = "XHR"
    ResourceTypeFetch ResourceType = "Fetch"
    ResourceTypeEventSource ResourceType = "EventSource"
    ResourceTypeWebSocket ResourceType = "WebSocket"
    ResourceTypeManifest ResourceType = "Manifest"
    ResourceTypeOther ResourceType = "Other"
)

const (
    UnserializableValueInfinity UnserializableValue = "Infinity"
    UnserializableValueNaN UnserializableValue = "NaN"
    UnserializableValueDashInfinity UnserializableValue = "-Infinity"
    UnserializableValueDash0 UnserializableValue = "-0"
)

type RemoteObjectTypeEnum string
const (
    RemoteObjectTypeObject RemoteObjectTypeEnum = "object"
    RemoteObjectTypeFunction RemoteObjectTypeEnum = "function"
    RemoteObjectTypeUndefined RemoteObjectTypeEnum = "undefined"
    RemoteObjectTypeString RemoteObjectTypeEnum = "string"
    RemoteObjectTypeNumber RemoteObjectTypeEnum = "number"
    RemoteObjectTypeBoolean RemoteObjectTypeEnum = "boolean"
    RemoteObjectTypeSymbol RemoteObjectTypeEnum = "symbol"
)

type RemoteObjectSubtypeEnum string
const (
    RemoteObjectSubtypeArray RemoteObjectSubtypeEnum = "array"
    RemoteObjectSubtypeNull RemoteObjectSubtypeEnum = "null"
    RemoteObjectSubtypeNode RemoteObjectSubtypeEnum = "node"
    RemoteObjectSubtypeRegexp RemoteObjectSubtypeEnum = "regexp"
    RemoteObjectSubtypeDate RemoteObjectSubtypeEnum = "date"
    RemoteObjectSubtypeMap RemoteObjectSubtypeEnum = "map"
    RemoteObjectSubtypeSet RemoteObjectSubtypeEnum = "set"
    RemoteObjectSubtypeIterator RemoteObjectSubtypeEnum = "iterator"
    RemoteObjectSubtypeGenerator RemoteObjectSubtypeEnum = "generator"
    RemoteObjectSubtypeError RemoteObjectSubtypeEnum = "error"
    RemoteObjectSubtypeProxy RemoteObjectSubtypeEnum = "proxy"
    RemoteObjectSubtypePromise RemoteObjectSubtypeEnum = "promise"
    RemoteObjectSubtypeTypedarray RemoteObjectSubtypeEnum = "typedarray"
)

type ObjectPreviewTypeEnum string
const (
    ObjectPreviewTypeObject ObjectPreviewTypeEnum = "object"
    ObjectPreviewTypeFunction ObjectPreviewTypeEnum = "function"
    ObjectPreviewTypeUndefined ObjectPreviewTypeEnum = "undefined"
    ObjectPreviewTypeString ObjectPreviewTypeEnum = "string"
    ObjectPreviewTypeNumber ObjectPreviewTypeEnum = "number"
    ObjectPreviewTypeBoolean ObjectPreviewTypeEnum = "boolean"
    ObjectPreviewTypeSymbol ObjectPreviewTypeEnum = "symbol"
)

type ObjectPreviewSubtypeEnum string
const (
    ObjectPreviewSubtypeArray ObjectPreviewSubtypeEnum = "array"
    ObjectPreviewSubtypeNull ObjectPreviewSubtypeEnum = "null"
    ObjectPreviewSubtypeNode ObjectPreviewSubtypeEnum = "node"
    ObjectPreviewSubtypeRegexp ObjectPreviewSubtypeEnum = "regexp"
    ObjectPreviewSubtypeDate ObjectPreviewSubtypeEnum = "date"
    ObjectPreviewSubtypeMap ObjectPreviewSubtypeEnum = "map"
    ObjectPreviewSubtypeSet ObjectPreviewSubtypeEnum = "set"
    ObjectPreviewSubtypeIterator ObjectPreviewSubtypeEnum = "iterator"
    ObjectPreviewSubtypeGenerator ObjectPreviewSubtypeEnum = "generator"
    ObjectPreviewSubtypeError ObjectPreviewSubtypeEnum = "error"
)

type PropertyPreviewTypeEnum string
const (
    PropertyPreviewTypeObject PropertyPreviewTypeEnum = "object"
    PropertyPreviewTypeFunction PropertyPreviewTypeEnum = "function"
    PropertyPreviewTypeUndefined PropertyPreviewTypeEnum = "undefined"
    PropertyPreviewTypeString PropertyPreviewTypeEnum = "string"
    PropertyPreviewTypeNumber PropertyPreviewTypeEnum = "number"
    PropertyPreviewTypeBoolean PropertyPreviewTypeEnum = "boolean"
    PropertyPreviewTypeSymbol PropertyPreviewTypeEnum = "symbol"
    PropertyPreviewTypeAccessor PropertyPreviewTypeEnum = "accessor"
)

type PropertyPreviewSubtypeEnum string
const (
    PropertyPreviewSubtypeArray PropertyPreviewSubtypeEnum = "array"
    PropertyPreviewSubtypeNull PropertyPreviewSubtypeEnum = "null"
    PropertyPreviewSubtypeNode PropertyPreviewSubtypeEnum = "node"
    PropertyPreviewSubtypeRegexp PropertyPreviewSubtypeEnum = "regexp"
    PropertyPreviewSubtypeDate PropertyPreviewSubtypeEnum = "date"
    PropertyPreviewSubtypeMap PropertyPreviewSubtypeEnum = "map"
    PropertyPreviewSubtypeSet PropertyPreviewSubtypeEnum = "set"
    PropertyPreviewSubtypeIterator PropertyPreviewSubtypeEnum = "iterator"
    PropertyPreviewSubtypeGenerator PropertyPreviewSubtypeEnum = "generator"
    PropertyPreviewSubtypeError PropertyPreviewSubtypeEnum = "error"
)
type ResourceType string

type ScriptId string

type RemoteObjectId string

type UnserializableValue string

type RemoteObject struct {
    Type RemoteObjectTypeEnum `json:"type"`// Object type.
    Subtype *RemoteObjectSubtypeEnum `json:"subtype,omitempty"`// Object subtype hint. Specified for <code>object</code> type values only.
    ClassName *string `json:"className,omitempty"`// Object class (constructor) name. Specified for <code>object</code> type values only.
    Value interface{} `json:"value"`// Remote object value in case of primitive values or JSON values (if it was requested).
    UnserializableValue *UnserializableValue `json:"unserializableValue,omitempty"`// Primitive value which can not be JSON-stringified does not have <code>value</code>, but gets this property.
    Description *string `json:"description,omitempty"`// String representation of the object.
    ObjectId *RemoteObjectId `json:"objectId,omitempty"`// Unique object identifier (for non-primitive values).
    Preview *ObjectPreview `json:"preview,omitempty"`// [Experimental] Preview containing abbreviated property values. Specified for <code>object</code> type values only.
    CustomPreview *CustomPreview `json:"customPreview,omitempty"`// [Experimental]
}

type CustomPreview struct {
    Header string `json:"header"`
    HasBody bool `json:"hasBody"`
    FormatterObjectId RemoteObjectId `json:"formatterObjectId"`
    BindRemoteObjectFunctionId RemoteObjectId `json:"bindRemoteObjectFunctionId"`
    ConfigObjectId *RemoteObjectId `json:"configObjectId,omitempty"`
}

type ObjectPreview struct {
    Type ObjectPreviewTypeEnum `json:"type"`// Object type.
    Subtype *ObjectPreviewSubtypeEnum `json:"subtype,omitempty"`// Object subtype hint. Specified for <code>object</code> type values only.
    Description *string `json:"description,omitempty"`// String representation of the object.
    Overflow bool `json:"overflow"`// True iff some of the properties or entries of the original object did not fit.
    Properties []PropertyPreview `json:"properties"`// List of the properties.
    Entries *[]EntryPreview `json:"entries,omitempty"`// List of the entries. Specified for <code>map</code> and <code>set</code> subtype values only.
}

type PropertyPreview struct {
    Name string `json:"name"`// Property name.
    Type PropertyPreviewTypeEnum `json:"type"`// Object type. Accessor means that the property itself is an accessor property.
    Value *string `json:"value,omitempty"`// User-friendly property value string.
    ValuePreview *ObjectPreview `json:"valuePreview,omitempty"`// Nested value preview.
    Subtype *PropertyPreviewSubtypeEnum `json:"subtype,omitempty"`// Object subtype hint. Specified for <code>object</code> type values only.
}

type EntryPreview struct {
    Key *ObjectPreview `json:"key,omitempty"`// Preview of the key. Specified for map-like collection entries.
    Value ObjectPreview `json:"value"`// Preview of the value.
}

type PropertyDescriptor struct {
    Name string `json:"name"`// Property name or symbol description.
    Value *RemoteObject `json:"value,omitempty"`// The value associated with the property.
    Writable *bool `json:"writable,omitempty"`// True if the value associated with the property may be changed (data descriptors only).
    Get *RemoteObject `json:"get,omitempty"`// A function which serves as a getter for the property, or <code>undefined</code> if there is no getter (accessor descriptors only).
    Set *RemoteObject `json:"set,omitempty"`// A function which serves as a setter for the property, or <code>undefined</code> if there is no setter (accessor descriptors only).
    Configurable bool `json:"configurable"`// True if the type of this property descriptor may be changed and if the property may be deleted from the corresponding object.
    Enumerable bool `json:"enumerable"`// True if this property shows up during enumeration of the properties on the corresponding object.
    WasThrown *bool `json:"wasThrown,omitempty"`// True if the result was thrown during the evaluation.
    IsOwn *bool `json:"isOwn,omitempty"`// True if the property is owned for the object.
    Symbol *RemoteObject `json:"symbol,omitempty"`// Property symbol object, if the property is of the <code>symbol</code> type.
}

type InternalPropertyDescriptor struct {
    Name string `json:"name"`// Conventional property name.
    Value *RemoteObject `json:"value,omitempty"`// The value associated with the property.
}

type CallArgument struct {
    Value interface{} `json:"value"`// Primitive value.
    UnserializableValue *UnserializableValue `json:"unserializableValue,omitempty"`// Primitive value which can not be JSON-stringified.
    ObjectId *RemoteObjectId `json:"objectId,omitempty"`// Remote object handle.
}

type ExecutionContextId int64

type ExecutionContextDescription struct {
    Id ExecutionContextId `json:"id"`// Unique id of the execution context. It can be used to specify in which execution context script evaluation should be performed.
    Origin string `json:"origin"`// Execution context origin.
    Name string `json:"name"`// Human readable name describing given context.
    AuxData *map[string]string `json:"auxData,omitempty"`// Embedder-specific auxiliary data.
}

type ExceptionDetails struct {
    ExceptionId int64 `json:"exceptionId"`// Exception id.
    Text string `json:"text"`// Exception text, which should be used together with exception object when available.
    LineNumber int64 `json:"lineNumber"`// Line number of the exception location (0-based).
    ColumnNumber int64 `json:"columnNumber"`// Column number of the exception location (0-based).
    ScriptId *ScriptId `json:"scriptId,omitempty"`// Script ID of the exception location.
    Url *string `json:"url,omitempty"`// URL of the exception location, to be used when the script was not reported.
    StackTrace *StackTrace `json:"stackTrace,omitempty"`// JavaScript stack trace if available.
    Exception *RemoteObject `json:"exception,omitempty"`// Exception object if available.
    ExecutionContextId *ExecutionContextId `json:"executionContextId,omitempty"`// Identifier of the context where exception happened.
}

type Timestamp float64

type CallFrame struct {
    FunctionName string `json:"functionName"`// JavaScript function name.
    ScriptId ScriptId `json:"scriptId"`// JavaScript script id.
    Url string `json:"url"`// JavaScript script name or url.
    LineNumber int64 `json:"lineNumber"`// JavaScript script line number (0-based).
    ColumnNumber int64 `json:"columnNumber"`// JavaScript script column number (0-based).
}

type StackTrace struct {
    Description *string `json:"description,omitempty"`// String label of this stack trace. For async traces this may be a name of the function that initiated the async call.
    CallFrames []CallFrame `json:"callFrames"`// JavaScript function name.
    Parent *StackTrace `json:"parent,omitempty"`// Asynchronous JavaScript stack trace that preceded this stack, if available.
    PromiseCreationFrame *CallFrame `json:"promiseCreationFrame,omitempty"`// [Experimental] Creation frame of the Promise which produced the next synchronous trace when resolved, if available.
}
