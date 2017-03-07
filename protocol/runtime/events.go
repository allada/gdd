package runtime


type ExecutionContextCreatedEvent struct {
    Context ExecutionContextDescription `json:"context"`// A newly created execution contex.
}
type ExecutionContextDestroyedEvent struct {
    ExecutionContextId ExecutionContextId `json:"executionContextId"`// Id of the destroyed context
}
type ExecutionContextsClearedEvent struct {
}
type ExceptionThrownEvent struct {
    Timestamp Timestamp `json:"timestamp"`// Timestamp of the exception.
    ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}
type ExceptionRevokedEvent struct {
    Reason string `json:"reason"`// Reason describing why exception was revoked.
    ExceptionId int64 `json:"exceptionId"`// The id of revoked exception, as reported in <code>exceptionUnhandled</code>.
}
type ConsoleAPICalledEvent struct {
    Type ConsoleAPICalledTypeEnum `json:"type"`// Type of the call.
    Args []RemoteObject `json:"args"`// Call arguments.
    ExecutionContextId ExecutionContextId `json:"executionContextId"`// Identifier of the context where the call was made.
    Timestamp Timestamp `json:"timestamp"`// Call timestamp.
    StackTrace *StackTrace `json:"stackTrace,omitempty"`// Stack trace captured when the call was made.
}
type InspectRequestedEvent struct {
    Object RemoteObject `json:"object"`
    Hints map[string]string `json:"hints"`
}