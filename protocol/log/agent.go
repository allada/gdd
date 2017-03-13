package log


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

type ClearCommandFn struct {
    mux sync.RWMutex
    fn func(ClearCommand)
}

func (a *ClearCommandFn) Load() func(ClearCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ClearCommandFn) Store(fn func(ClearCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StartViolationsReportCommandFn struct {
    mux sync.RWMutex
    fn func(StartViolationsReportCommand)
}

func (a *StartViolationsReportCommandFn) Load() func(StartViolationsReportCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StartViolationsReportCommandFn) Store(fn func(StartViolationsReportCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StopViolationsReportCommandFn struct {
    mux sync.RWMutex
    fn func(StopViolationsReportCommand)
}

func (a *StopViolationsReportCommandFn) Load() func(StopViolationsReportCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StopViolationsReportCommandFn) Store(fn func(StopViolationsReportCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type LogAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
        Clear ClearCommandFn
        StartViolationsReport StartViolationsReportCommandFn
        StopViolationsReport StopViolationsReportCommandFn
    }
}
func NewAgent(conn *shared.Connection) *LogAgent {
    agent := &LogAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *LogAgent) Name() string {
    return "Log"
}

func (agent *LogAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    defer func() {
        data := recover()
        switch data.(type) {
        case nil:
            return
        case shared.Warning:
            fmt.Println(data)
        case shared.Error:
            fmt.Println("Closing websocket because of following Error:")
            fmt.Println(data)
            agent.conn.Close()
        case error:
            fmt.Println("Closing websocket because of following Error:")
            fmt.Println(data)
            agent.conn.Close()
        default:
            fmt.Println("Should probably use shared.Warning or shared.Error instead to panic()")
            panic(data)
        }
    }()
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
        case "clear":
            var out ClearCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Clear.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "startViolationsReport":
            var out StartViolationsReportCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StartViolationsReport.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "stopViolationsReport":
            var out StopViolationsReportCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StopViolationsReport.Load()
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
func (agent *LogAgent) FireEntryAdded(event EntryAddedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Log.entryAdded",
        Params: event,
    })
}
func (agent *LogAgent) FireEntryAddedOnTarget(targetId string,event EntryAddedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Log.entryAdded",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *LogAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *LogAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *LogAgent) SetClearHandler(handler func(ClearCommand)) {
    agent.commandHandlers.Clear.Store(handler)
}
func (agent *LogAgent) SetStartViolationsReportHandler(handler func(StartViolationsReportCommand)) {
    agent.commandHandlers.StartViolationsReport.Store(handler)
}
func (agent *LogAgent) SetStopViolationsReportHandler(handler func(StopViolationsReportCommand)) {
    agent.commandHandlers.StopViolationsReport.Store(handler)
}
func init() {

}
