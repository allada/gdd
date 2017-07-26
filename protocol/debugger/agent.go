package debugger


import (
    "github.com/allada/gdd/protocol/shared"
    "sync"
    "encoding/json"
    "fmt"
)
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

type SetBreakpointsActiveCommandFn struct {
    mux sync.RWMutex
    fn func(SetBreakpointsActiveCommand)
}

func (a *SetBreakpointsActiveCommandFn) Load() func(SetBreakpointsActiveCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetBreakpointsActiveCommandFn) Store(fn func(SetBreakpointsActiveCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetSkipAllPausesCommandFn struct {
    mux sync.RWMutex
    fn func(SetSkipAllPausesCommand)
}

func (a *SetSkipAllPausesCommandFn) Load() func(SetSkipAllPausesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetSkipAllPausesCommandFn) Store(fn func(SetSkipAllPausesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetBreakpointByUrlCommandFn struct {
    mux sync.RWMutex
    fn func(SetBreakpointByUrlCommand)
}

func (a *SetBreakpointByUrlCommandFn) Load() func(SetBreakpointByUrlCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetBreakpointByUrlCommandFn) Store(fn func(SetBreakpointByUrlCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetBreakpointCommandFn struct {
    mux sync.RWMutex
    fn func(SetBreakpointCommand)
}

func (a *SetBreakpointCommandFn) Load() func(SetBreakpointCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetBreakpointCommandFn) Store(fn func(SetBreakpointCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RemoveBreakpointCommandFn struct {
    mux sync.RWMutex
    fn func(RemoveBreakpointCommand)
}

func (a *RemoveBreakpointCommandFn) Load() func(RemoveBreakpointCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RemoveBreakpointCommandFn) Store(fn func(RemoveBreakpointCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetPossibleBreakpointsCommandFn struct {
    mux sync.RWMutex
    fn func(GetPossibleBreakpointsCommand)
}

func (a *GetPossibleBreakpointsCommandFn) Load() func(GetPossibleBreakpointsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetPossibleBreakpointsCommandFn) Store(fn func(GetPossibleBreakpointsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ContinueToLocationCommandFn struct {
    mux sync.RWMutex
    fn func(ContinueToLocationCommand)
}

func (a *ContinueToLocationCommandFn) Load() func(ContinueToLocationCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ContinueToLocationCommandFn) Store(fn func(ContinueToLocationCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StepOverCommandFn struct {
    mux sync.RWMutex
    fn func(StepOverCommand)
}

func (a *StepOverCommandFn) Load() func(StepOverCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StepOverCommandFn) Store(fn func(StepOverCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StepIntoCommandFn struct {
    mux sync.RWMutex
    fn func(StepIntoCommand)
}

func (a *StepIntoCommandFn) Load() func(StepIntoCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StepIntoCommandFn) Store(fn func(StepIntoCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StepOutCommandFn struct {
    mux sync.RWMutex
    fn func(StepOutCommand)
}

func (a *StepOutCommandFn) Load() func(StepOutCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StepOutCommandFn) Store(fn func(StepOutCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type PauseCommandFn struct {
    mux sync.RWMutex
    fn func(PauseCommand)
}

func (a *PauseCommandFn) Load() func(PauseCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *PauseCommandFn) Store(fn func(PauseCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ResumeCommandFn struct {
    mux sync.RWMutex
    fn func(ResumeCommand)
}

func (a *ResumeCommandFn) Load() func(ResumeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ResumeCommandFn) Store(fn func(ResumeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SearchInContentCommandFn struct {
    mux sync.RWMutex
    fn func(SearchInContentCommand)
}

func (a *SearchInContentCommandFn) Load() func(SearchInContentCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SearchInContentCommandFn) Store(fn func(SearchInContentCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetScriptSourceCommandFn struct {
    mux sync.RWMutex
    fn func(SetScriptSourceCommand)
}

func (a *SetScriptSourceCommandFn) Load() func(SetScriptSourceCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetScriptSourceCommandFn) Store(fn func(SetScriptSourceCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RestartFrameCommandFn struct {
    mux sync.RWMutex
    fn func(RestartFrameCommand)
}

func (a *RestartFrameCommandFn) Load() func(RestartFrameCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RestartFrameCommandFn) Store(fn func(RestartFrameCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetScriptSourceCommandFn struct {
    mux sync.RWMutex
    fn func(GetScriptSourceCommand)
}

func (a *GetScriptSourceCommandFn) Load() func(GetScriptSourceCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetScriptSourceCommandFn) Store(fn func(GetScriptSourceCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetPauseOnExceptionsCommandFn struct {
    mux sync.RWMutex
    fn func(SetPauseOnExceptionsCommand)
}

func (a *SetPauseOnExceptionsCommandFn) Load() func(SetPauseOnExceptionsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetPauseOnExceptionsCommandFn) Store(fn func(SetPauseOnExceptionsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type EvaluateOnCallFrameCommandFn struct {
    mux sync.RWMutex
    fn func(EvaluateOnCallFrameCommand)
}

func (a *EvaluateOnCallFrameCommandFn) Load() func(EvaluateOnCallFrameCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *EvaluateOnCallFrameCommandFn) Store(fn func(EvaluateOnCallFrameCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetVariableValueCommandFn struct {
    mux sync.RWMutex
    fn func(SetVariableValueCommand)
}

func (a *SetVariableValueCommandFn) Load() func(SetVariableValueCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetVariableValueCommandFn) Store(fn func(SetVariableValueCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetAsyncCallStackDepthCommandFn struct {
    mux sync.RWMutex
    fn func(SetAsyncCallStackDepthCommand)
}

func (a *SetAsyncCallStackDepthCommandFn) Load() func(SetAsyncCallStackDepthCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetAsyncCallStackDepthCommandFn) Store(fn func(SetAsyncCallStackDepthCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetBlackboxPatternsCommandFn struct {
    mux sync.RWMutex
    fn func(SetBlackboxPatternsCommand)
}

func (a *SetBlackboxPatternsCommandFn) Load() func(SetBlackboxPatternsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetBlackboxPatternsCommandFn) Store(fn func(SetBlackboxPatternsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetBlackboxedRangesCommandFn struct {
    mux sync.RWMutex
    fn func(SetBlackboxedRangesCommand)
}

func (a *SetBlackboxedRangesCommandFn) Load() func(SetBlackboxedRangesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetBlackboxedRangesCommandFn) Store(fn func(SetBlackboxedRangesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DebuggerAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
        SetBreakpointsActive SetBreakpointsActiveCommandFn
        SetSkipAllPauses SetSkipAllPausesCommandFn
        SetBreakpointByUrl SetBreakpointByUrlCommandFn
        SetBreakpoint SetBreakpointCommandFn
        RemoveBreakpoint RemoveBreakpointCommandFn
        GetPossibleBreakpoints GetPossibleBreakpointsCommandFn
        ContinueToLocation ContinueToLocationCommandFn
        StepOver StepOverCommandFn
        StepInto StepIntoCommandFn
        StepOut StepOutCommandFn
        Pause PauseCommandFn
        Resume ResumeCommandFn
        SearchInContent SearchInContentCommandFn
        SetScriptSource SetScriptSourceCommandFn
        RestartFrame RestartFrameCommandFn
        GetScriptSource GetScriptSourceCommandFn
        SetPauseOnExceptions SetPauseOnExceptionsCommandFn
        EvaluateOnCallFrame EvaluateOnCallFrameCommandFn
        SetVariableValue SetVariableValueCommandFn
        SetAsyncCallStackDepth SetAsyncCallStackDepthCommandFn
        SetBlackboxPatterns SetBlackboxPatternsCommandFn
        SetBlackboxedRanges SetBlackboxedRangesCommandFn
    }
}
func NewAgent(conn *shared.Connection) *DebuggerAgent {
    agent := &DebuggerAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *DebuggerAgent) Name() string {
    return "Debugger"
}

func (agent *DebuggerAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    defer shared.TryRecoverFromPanic(agent.conn)
    switch (funcName) {
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
        case "setBreakpointsActive":
            var out SetBreakpointsActiveCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetBreakpointsActive.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setSkipAllPauses":
            var out SetSkipAllPausesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetSkipAllPauses.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setBreakpointByUrl":
            var out SetBreakpointByUrlCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetBreakpointByUrl.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setBreakpoint":
            var out SetBreakpointCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetBreakpoint.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "removeBreakpoint":
            var out RemoveBreakpointCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RemoveBreakpoint.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getPossibleBreakpoints":
            var out GetPossibleBreakpointsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetPossibleBreakpoints.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "continueToLocation":
            var out ContinueToLocationCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ContinueToLocation.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "stepOver":
            var out StepOverCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StepOver.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "stepInto":
            var out StepIntoCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StepInto.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "stepOut":
            var out StepOutCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StepOut.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "pause":
            var out PauseCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Pause.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "resume":
            var out ResumeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Resume.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "searchInContent":
            var out SearchInContentCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SearchInContent.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setScriptSource":
            var out SetScriptSourceCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetScriptSource.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "restartFrame":
            var out RestartFrameCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RestartFrame.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getScriptSource":
            var out GetScriptSourceCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetScriptSource.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setPauseOnExceptions":
            var out SetPauseOnExceptionsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetPauseOnExceptions.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "evaluateOnCallFrame":
            var out EvaluateOnCallFrameCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.EvaluateOnCallFrame.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setVariableValue":
            var out SetVariableValueCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetVariableValue.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setAsyncCallStackDepth":
            var out SetAsyncCallStackDepthCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetAsyncCallStackDepth.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setBlackboxPatterns":
            var out SetBlackboxPatternsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetBlackboxPatterns.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setBlackboxedRanges":
            var out SetBlackboxedRangesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetBlackboxedRanges.Load()
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
func (agent *DebuggerAgent) FireScriptParsed(event ScriptParsedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Debugger.scriptParsed",
        Params: event,
    })
}
func (agent *DebuggerAgent) FireScriptParsedOnTarget(targetId string,event ScriptParsedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Debugger.scriptParsed",
        Params: event,
    })
}
func (agent *DebuggerAgent) FireScriptFailedToParse(event ScriptFailedToParseEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Debugger.scriptFailedToParse",
        Params: event,
    })
}
func (agent *DebuggerAgent) FireScriptFailedToParseOnTarget(targetId string,event ScriptFailedToParseEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Debugger.scriptFailedToParse",
        Params: event,
    })
}
func (agent *DebuggerAgent) FireBreakpointResolved(event BreakpointResolvedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Debugger.breakpointResolved",
        Params: event,
    })
}
func (agent *DebuggerAgent) FireBreakpointResolvedOnTarget(targetId string,event BreakpointResolvedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Debugger.breakpointResolved",
        Params: event,
    })
}
func (agent *DebuggerAgent) FirePaused(event PausedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Debugger.paused",
        Params: event,
    })
}
func (agent *DebuggerAgent) FirePausedOnTarget(targetId string,event PausedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Debugger.paused",
        Params: event,
    })
}
func (agent *DebuggerAgent) FireResumed() {
    agent.conn.Send(shared.Notification{
        Method: "Debugger.resumed",
    })
}
func (agent *DebuggerAgent) FireResumedOnTarget(targetId string) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Debugger.resumed",
    })
}

// Commands Sent From Frontend
func (agent *DebuggerAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *DebuggerAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *DebuggerAgent) SetSetBreakpointsActiveHandler(handler func(SetBreakpointsActiveCommand)) {
    agent.commandHandlers.SetBreakpointsActive.Store(handler)
}
func (agent *DebuggerAgent) SetSetSkipAllPausesHandler(handler func(SetSkipAllPausesCommand)) {
    agent.commandHandlers.SetSkipAllPauses.Store(handler)
}
func (agent *DebuggerAgent) SetSetBreakpointByUrlHandler(handler func(SetBreakpointByUrlCommand)) {
    agent.commandHandlers.SetBreakpointByUrl.Store(handler)
}
func (agent *DebuggerAgent) SetSetBreakpointHandler(handler func(SetBreakpointCommand)) {
    agent.commandHandlers.SetBreakpoint.Store(handler)
}
func (agent *DebuggerAgent) SetRemoveBreakpointHandler(handler func(RemoveBreakpointCommand)) {
    agent.commandHandlers.RemoveBreakpoint.Store(handler)
}
func (agent *DebuggerAgent) SetGetPossibleBreakpointsHandler(handler func(GetPossibleBreakpointsCommand)) {
    agent.commandHandlers.GetPossibleBreakpoints.Store(handler)
}
func (agent *DebuggerAgent) SetContinueToLocationHandler(handler func(ContinueToLocationCommand)) {
    agent.commandHandlers.ContinueToLocation.Store(handler)
}
func (agent *DebuggerAgent) SetStepOverHandler(handler func(StepOverCommand)) {
    agent.commandHandlers.StepOver.Store(handler)
}
func (agent *DebuggerAgent) SetStepIntoHandler(handler func(StepIntoCommand)) {
    agent.commandHandlers.StepInto.Store(handler)
}
func (agent *DebuggerAgent) SetStepOutHandler(handler func(StepOutCommand)) {
    agent.commandHandlers.StepOut.Store(handler)
}
func (agent *DebuggerAgent) SetPauseHandler(handler func(PauseCommand)) {
    agent.commandHandlers.Pause.Store(handler)
}
func (agent *DebuggerAgent) SetResumeHandler(handler func(ResumeCommand)) {
    agent.commandHandlers.Resume.Store(handler)
}
func (agent *DebuggerAgent) SetSearchInContentHandler(handler func(SearchInContentCommand)) {
    agent.commandHandlers.SearchInContent.Store(handler)
}
func (agent *DebuggerAgent) SetSetScriptSourceHandler(handler func(SetScriptSourceCommand)) {
    agent.commandHandlers.SetScriptSource.Store(handler)
}
func (agent *DebuggerAgent) SetRestartFrameHandler(handler func(RestartFrameCommand)) {
    agent.commandHandlers.RestartFrame.Store(handler)
}
func (agent *DebuggerAgent) SetGetScriptSourceHandler(handler func(GetScriptSourceCommand)) {
    agent.commandHandlers.GetScriptSource.Store(handler)
}
func (agent *DebuggerAgent) SetSetPauseOnExceptionsHandler(handler func(SetPauseOnExceptionsCommand)) {
    agent.commandHandlers.SetPauseOnExceptions.Store(handler)
}
func (agent *DebuggerAgent) SetEvaluateOnCallFrameHandler(handler func(EvaluateOnCallFrameCommand)) {
    agent.commandHandlers.EvaluateOnCallFrame.Store(handler)
}
func (agent *DebuggerAgent) SetSetVariableValueHandler(handler func(SetVariableValueCommand)) {
    agent.commandHandlers.SetVariableValue.Store(handler)
}
func (agent *DebuggerAgent) SetSetAsyncCallStackDepthHandler(handler func(SetAsyncCallStackDepthCommand)) {
    agent.commandHandlers.SetAsyncCallStackDepth.Store(handler)
}
func (agent *DebuggerAgent) SetSetBlackboxPatternsHandler(handler func(SetBlackboxPatternsCommand)) {
    agent.commandHandlers.SetBlackboxPatterns.Store(handler)
}
func (agent *DebuggerAgent) SetSetBlackboxedRangesHandler(handler func(SetBlackboxedRangesCommand)) {
    agent.commandHandlers.SetBlackboxedRanges.Store(handler)
}
func init() {

}
