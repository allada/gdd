package memory


import (
    "../shared"
    "sync"
    "encoding/json"
    "fmt"
)
type GetDOMCountersCommandFn struct {
    mux sync.RWMutex
    fn func(GetDOMCountersCommand)
}

func (a *GetDOMCountersCommandFn) Load() func(GetDOMCountersCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetDOMCountersCommandFn) Store(fn func(GetDOMCountersCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetPressureNotificationsSuppressedCommandFn struct {
    mux sync.RWMutex
    fn func(SetPressureNotificationsSuppressedCommand)
}

func (a *SetPressureNotificationsSuppressedCommandFn) Load() func(SetPressureNotificationsSuppressedCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetPressureNotificationsSuppressedCommandFn) Store(fn func(SetPressureNotificationsSuppressedCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SimulatePressureNotificationCommandFn struct {
    mux sync.RWMutex
    fn func(SimulatePressureNotificationCommand)
}

func (a *SimulatePressureNotificationCommandFn) Load() func(SimulatePressureNotificationCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SimulatePressureNotificationCommandFn) Store(fn func(SimulatePressureNotificationCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type MemoryAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        GetDOMCounters GetDOMCountersCommandFn
        SetPressureNotificationsSuppressed SetPressureNotificationsSuppressedCommandFn
        SimulatePressureNotification SimulatePressureNotificationCommandFn
    }
}
func NewAgent(conn *shared.Connection) *MemoryAgent {
    agent := &MemoryAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *MemoryAgent) Name() string {
    return "Memory"
}

func (agent *MemoryAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "getDOMCounters":
            var out GetDOMCountersCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetDOMCounters.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setPressureNotificationsSuppressed":
            var out SetPressureNotificationsSuppressedCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetPressureNotificationsSuppressed.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "simulatePressureNotification":
            var out SimulatePressureNotificationCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SimulatePressureNotification.Load()
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

// Commands Sent From Frontend
func (agent *MemoryAgent) SetGetDOMCountersHandler(handler func(GetDOMCountersCommand)) {
    agent.commandHandlers.GetDOMCounters.Store(handler)
}
func (agent *MemoryAgent) SetSetPressureNotificationsSuppressedHandler(handler func(SetPressureNotificationsSuppressedCommand)) {
    agent.commandHandlers.SetPressureNotificationsSuppressed.Store(handler)
}
func (agent *MemoryAgent) SetSimulatePressureNotificationHandler(handler func(SimulatePressureNotificationCommand)) {
    agent.commandHandlers.SimulatePressureNotification.Store(handler)
}
func init() {

}
