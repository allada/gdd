package runtime


import (
    "github.com/allada/gdd/protocol/shared"
    "sync"
    "encoding/json"
    "fmt"
)
type EvaluateCommandFn struct {
    mux sync.RWMutex
    fn func(EvaluateCommand)
}

func (a *EvaluateCommandFn) Load() func(EvaluateCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *EvaluateCommandFn) Store(fn func(EvaluateCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type AwaitPromiseCommandFn struct {
    mux sync.RWMutex
    fn func(AwaitPromiseCommand)
}

func (a *AwaitPromiseCommandFn) Load() func(AwaitPromiseCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *AwaitPromiseCommandFn) Store(fn func(AwaitPromiseCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CallFunctionOnCommandFn struct {
    mux sync.RWMutex
    fn func(CallFunctionOnCommand)
}

func (a *CallFunctionOnCommandFn) Load() func(CallFunctionOnCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CallFunctionOnCommandFn) Store(fn func(CallFunctionOnCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetPropertiesCommandFn struct {
    mux sync.RWMutex
    fn func(GetPropertiesCommand)
}

func (a *GetPropertiesCommandFn) Load() func(GetPropertiesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetPropertiesCommandFn) Store(fn func(GetPropertiesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ReleaseObjectCommandFn struct {
    mux sync.RWMutex
    fn func(ReleaseObjectCommand)
}

func (a *ReleaseObjectCommandFn) Load() func(ReleaseObjectCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ReleaseObjectCommandFn) Store(fn func(ReleaseObjectCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ReleaseObjectGroupCommandFn struct {
    mux sync.RWMutex
    fn func(ReleaseObjectGroupCommand)
}

func (a *ReleaseObjectGroupCommandFn) Load() func(ReleaseObjectGroupCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ReleaseObjectGroupCommandFn) Store(fn func(ReleaseObjectGroupCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RunIfWaitingForDebuggerCommandFn struct {
    mux sync.RWMutex
    fn func(RunIfWaitingForDebuggerCommand)
}

func (a *RunIfWaitingForDebuggerCommandFn) Load() func(RunIfWaitingForDebuggerCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RunIfWaitingForDebuggerCommandFn) Store(fn func(RunIfWaitingForDebuggerCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type EnableCommandFn struct {
    mux sync.RWMutex
    fn func(EnableCommand)
}

func (a *EnableCommandFn) Load() func(EnableCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *EnableCommandFn) Store(fn func(EnableCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DisableCommandFn struct {
    mux sync.RWMutex
    fn func(DisableCommand)
}

func (a *DisableCommandFn) Load() func(DisableCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *DisableCommandFn) Store(fn func(DisableCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DiscardConsoleEntriesCommandFn struct {
    mux sync.RWMutex
    fn func(DiscardConsoleEntriesCommand)
}

func (a *DiscardConsoleEntriesCommandFn) Load() func(DiscardConsoleEntriesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *DiscardConsoleEntriesCommandFn) Store(fn func(DiscardConsoleEntriesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetCustomObjectFormatterEnabledCommandFn struct {
    mux sync.RWMutex
    fn func(SetCustomObjectFormatterEnabledCommand)
}

func (a *SetCustomObjectFormatterEnabledCommandFn) Load() func(SetCustomObjectFormatterEnabledCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetCustomObjectFormatterEnabledCommandFn) Store(fn func(SetCustomObjectFormatterEnabledCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CompileScriptCommandFn struct {
    mux sync.RWMutex
    fn func(CompileScriptCommand)
}

func (a *CompileScriptCommandFn) Load() func(CompileScriptCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CompileScriptCommandFn) Store(fn func(CompileScriptCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RunScriptCommandFn struct {
    mux sync.RWMutex
    fn func(RunScriptCommand)
}

func (a *RunScriptCommandFn) Load() func(RunScriptCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RunScriptCommandFn) Store(fn func(RunScriptCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RuntimeAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Evaluate EvaluateCommandFn
        AwaitPromise AwaitPromiseCommandFn
        CallFunctionOn CallFunctionOnCommandFn
        GetProperties GetPropertiesCommandFn
        ReleaseObject ReleaseObjectCommandFn
        ReleaseObjectGroup ReleaseObjectGroupCommandFn
        RunIfWaitingForDebugger RunIfWaitingForDebuggerCommandFn
        Enable EnableCommandFn
        Disable DisableCommandFn
        DiscardConsoleEntries DiscardConsoleEntriesCommandFn
        SetCustomObjectFormatterEnabled SetCustomObjectFormatterEnabledCommandFn
        CompileScript CompileScriptCommandFn
        RunScript RunScriptCommandFn
    }
}
func NewAgent(conn *shared.Connection) *RuntimeAgent {
    agent := &RuntimeAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *RuntimeAgent) Name() string {
    return "Runtime"
}

func (agent *RuntimeAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    defer shared.TryRecoverFromPanic(agent.conn)
    switch (funcName) {
        case "evaluate":
            var out EvaluateCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Evaluate.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "awaitPromise":
            var out AwaitPromiseCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.AwaitPromise.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "callFunctionOn":
            var out CallFunctionOnCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CallFunctionOn.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getProperties":
            var out GetPropertiesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetProperties.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "releaseObject":
            var out ReleaseObjectCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ReleaseObject.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "releaseObjectGroup":
            var out ReleaseObjectGroupCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ReleaseObjectGroup.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "runIfWaitingForDebugger":
            var out RunIfWaitingForDebuggerCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RunIfWaitingForDebugger.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "enable":
            var out EnableCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Enable.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "disable":
            var out DisableCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Disable.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "discardConsoleEntries":
            var out DiscardConsoleEntriesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.DiscardConsoleEntries.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setCustomObjectFormatterEnabled":
            var out SetCustomObjectFormatterEnabledCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetCustomObjectFormatterEnabled.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "compileScript":
            var out CompileScriptCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CompileScript.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "runScript":
            var out RunScriptCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RunScript.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        default:
            fmt.Printf("Command %s unknown\n", funcName)
            agent.conn.SendErrorResult(id, targetId, shared.ErrorCodeMethodNotFound, "")
    }
}

// Dispatchable Events
func (agent *RuntimeAgent) FireExecutionContextCreated(event ExecutionContextCreatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Runtime.executionContextCreated",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExecutionContextCreatedOnTarget(targetId string,event ExecutionContextCreatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Runtime.executionContextCreated",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExecutionContextDestroyed(event ExecutionContextDestroyedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Runtime.executionContextDestroyed",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExecutionContextDestroyedOnTarget(targetId string,event ExecutionContextDestroyedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Runtime.executionContextDestroyed",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExecutionContextsCleared() {
    agent.conn.Send(shared.Notification{
        Method: "Runtime.executionContextsCleared",
    })
}
func (agent *RuntimeAgent) FireExecutionContextsClearedOnTarget(targetId string) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Runtime.executionContextsCleared",
    })
}
func (agent *RuntimeAgent) FireExceptionThrown(event ExceptionThrownEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Runtime.exceptionThrown",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExceptionThrownOnTarget(targetId string,event ExceptionThrownEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Runtime.exceptionThrown",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExceptionRevoked(event ExceptionRevokedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Runtime.exceptionRevoked",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExceptionRevokedOnTarget(targetId string,event ExceptionRevokedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Runtime.exceptionRevoked",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireConsoleAPICalled(event ConsoleAPICalledEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Runtime.consoleAPICalled",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireConsoleAPICalledOnTarget(targetId string,event ConsoleAPICalledEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Runtime.consoleAPICalled",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireInspectRequested(event InspectRequestedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Runtime.inspectRequested",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireInspectRequestedOnTarget(targetId string,event InspectRequestedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Runtime.inspectRequested",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *RuntimeAgent) SetEvaluateHandler(handler func(EvaluateCommand)) {
    agent.commandHandlers.Evaluate.Store(handler)
}
func (agent *RuntimeAgent) SetAwaitPromiseHandler(handler func(AwaitPromiseCommand)) {
    agent.commandHandlers.AwaitPromise.Store(handler)
}
func (agent *RuntimeAgent) SetCallFunctionOnHandler(handler func(CallFunctionOnCommand)) {
    agent.commandHandlers.CallFunctionOn.Store(handler)
}
func (agent *RuntimeAgent) SetGetPropertiesHandler(handler func(GetPropertiesCommand)) {
    agent.commandHandlers.GetProperties.Store(handler)
}
func (agent *RuntimeAgent) SetReleaseObjectHandler(handler func(ReleaseObjectCommand)) {
    agent.commandHandlers.ReleaseObject.Store(handler)
}
func (agent *RuntimeAgent) SetReleaseObjectGroupHandler(handler func(ReleaseObjectGroupCommand)) {
    agent.commandHandlers.ReleaseObjectGroup.Store(handler)
}
func (agent *RuntimeAgent) SetRunIfWaitingForDebuggerHandler(handler func(RunIfWaitingForDebuggerCommand)) {
    agent.commandHandlers.RunIfWaitingForDebugger.Store(handler)
}
func (agent *RuntimeAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *RuntimeAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *RuntimeAgent) SetDiscardConsoleEntriesHandler(handler func(DiscardConsoleEntriesCommand)) {
    agent.commandHandlers.DiscardConsoleEntries.Store(handler)
}
func (agent *RuntimeAgent) SetSetCustomObjectFormatterEnabledHandler(handler func(SetCustomObjectFormatterEnabledCommand)) {
    agent.commandHandlers.SetCustomObjectFormatterEnabled.Store(handler)
}
func (agent *RuntimeAgent) SetCompileScriptHandler(handler func(CompileScriptCommand)) {
    agent.commandHandlers.CompileScript.Store(handler)
}
func (agent *RuntimeAgent) SetRunScriptHandler(handler func(RunScriptCommand)) {
    agent.commandHandlers.RunScript.Store(handler)
}
func init() {

}
