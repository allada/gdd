package profiler


import (
    "../shared"
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

type SetSamplingIntervalCommandFn struct {
    mux sync.RWMutex
    fn func(SetSamplingIntervalCommand)
}

func (a *SetSamplingIntervalCommandFn) Load() func(SetSamplingIntervalCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetSamplingIntervalCommandFn) Store(fn func(SetSamplingIntervalCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StartCommandFn struct {
    mux sync.RWMutex
    fn func(StartCommand)
}

func (a *StartCommandFn) Load() func(StartCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StartCommandFn) Store(fn func(StartCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StopCommandFn struct {
    mux sync.RWMutex
    fn func(StopCommand)
}

func (a *StopCommandFn) Load() func(StopCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StopCommandFn) Store(fn func(StopCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StartPreciseCoverageCommandFn struct {
    mux sync.RWMutex
    fn func(StartPreciseCoverageCommand)
}

func (a *StartPreciseCoverageCommandFn) Load() func(StartPreciseCoverageCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StartPreciseCoverageCommandFn) Store(fn func(StartPreciseCoverageCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StopPreciseCoverageCommandFn struct {
    mux sync.RWMutex
    fn func(StopPreciseCoverageCommand)
}

func (a *StopPreciseCoverageCommandFn) Load() func(StopPreciseCoverageCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StopPreciseCoverageCommandFn) Store(fn func(StopPreciseCoverageCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type TakePreciseCoverageCommandFn struct {
    mux sync.RWMutex
    fn func(TakePreciseCoverageCommand)
}

func (a *TakePreciseCoverageCommandFn) Load() func(TakePreciseCoverageCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *TakePreciseCoverageCommandFn) Store(fn func(TakePreciseCoverageCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetBestEffortCoverageCommandFn struct {
    mux sync.RWMutex
    fn func(GetBestEffortCoverageCommand)
}

func (a *GetBestEffortCoverageCommandFn) Load() func(GetBestEffortCoverageCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetBestEffortCoverageCommandFn) Store(fn func(GetBestEffortCoverageCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ProfilerAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
        SetSamplingInterval SetSamplingIntervalCommandFn
        Start StartCommandFn
        Stop StopCommandFn
        StartPreciseCoverage StartPreciseCoverageCommandFn
        StopPreciseCoverage StopPreciseCoverageCommandFn
        TakePreciseCoverage TakePreciseCoverageCommandFn
        GetBestEffortCoverage GetBestEffortCoverageCommandFn
    }
}
func NewAgent(conn *shared.Connection) *ProfilerAgent {
    agent := &ProfilerAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *ProfilerAgent) Name() string {
    return "Profiler"
}

func (agent *ProfilerAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "setSamplingInterval":
            var out SetSamplingIntervalCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetSamplingInterval.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "start":
            var out StartCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Start.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "stop":
            var out StopCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Stop.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "startPreciseCoverage":
            var out StartPreciseCoverageCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StartPreciseCoverage.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "stopPreciseCoverage":
            var out StopPreciseCoverageCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StopPreciseCoverage.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "takePreciseCoverage":
            var out TakePreciseCoverageCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.TakePreciseCoverage.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getBestEffortCoverage":
            var out GetBestEffortCoverageCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetBestEffortCoverage.Load()
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
func (agent *ProfilerAgent) FireConsoleProfileStarted(event ConsoleProfileStartedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Profiler.consoleProfileStarted",
        Params: event,
    })
}
func (agent *ProfilerAgent) FireConsoleProfileStartedOnTarget(targetId string,event ConsoleProfileStartedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Profiler.consoleProfileStarted",
        Params: event,
    })
}
func (agent *ProfilerAgent) FireConsoleProfileFinished(event ConsoleProfileFinishedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Profiler.consoleProfileFinished",
        Params: event,
    })
}
func (agent *ProfilerAgent) FireConsoleProfileFinishedOnTarget(targetId string,event ConsoleProfileFinishedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Profiler.consoleProfileFinished",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *ProfilerAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *ProfilerAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *ProfilerAgent) SetSetSamplingIntervalHandler(handler func(SetSamplingIntervalCommand)) {
    agent.commandHandlers.SetSamplingInterval.Store(handler)
}
func (agent *ProfilerAgent) SetStartHandler(handler func(StartCommand)) {
    agent.commandHandlers.Start.Store(handler)
}
func (agent *ProfilerAgent) SetStopHandler(handler func(StopCommand)) {
    agent.commandHandlers.Stop.Store(handler)
}
func (agent *ProfilerAgent) SetStartPreciseCoverageHandler(handler func(StartPreciseCoverageCommand)) {
    agent.commandHandlers.StartPreciseCoverage.Store(handler)
}
func (agent *ProfilerAgent) SetStopPreciseCoverageHandler(handler func(StopPreciseCoverageCommand)) {
    agent.commandHandlers.StopPreciseCoverage.Store(handler)
}
func (agent *ProfilerAgent) SetTakePreciseCoverageHandler(handler func(TakePreciseCoverageCommand)) {
    agent.commandHandlers.TakePreciseCoverage.Store(handler)
}
func (agent *ProfilerAgent) SetGetBestEffortCoverageHandler(handler func(GetBestEffortCoverageCommand)) {
    agent.commandHandlers.GetBestEffortCoverage.Store(handler)
}
func init() {

}
