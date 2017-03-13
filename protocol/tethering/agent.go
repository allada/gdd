package tethering


import (
    "../shared"
    "sync"
    "encoding/json"
    "fmt"
)
type BindCommandFn struct {
    mux sync.RWMutex
    fn func(BindCommand)
}

func (a *BindCommandFn) Load() func(BindCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *BindCommandFn) Store(fn func(BindCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type UnbindCommandFn struct {
    mux sync.RWMutex
    fn func(UnbindCommand)
}

func (a *UnbindCommandFn) Load() func(UnbindCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *UnbindCommandFn) Store(fn func(UnbindCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type TetheringAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Bind BindCommandFn
        Unbind UnbindCommandFn
    }
}
func NewAgent(conn *shared.Connection) *TetheringAgent {
    agent := &TetheringAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *TetheringAgent) Name() string {
    return "Tethering"
}

func (agent *TetheringAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "bind":
            var out BindCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Bind.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "unbind":
            var out UnbindCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Unbind.Load()
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
func (agent *TetheringAgent) FireAccepted(event AcceptedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Tethering.accepted",
        Params: event,
    })
}
func (agent *TetheringAgent) FireAcceptedOnTarget(targetId string,event AcceptedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Tethering.accepted",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *TetheringAgent) SetBindHandler(handler func(BindCommand)) {
    agent.commandHandlers.Bind.Store(handler)
}
func (agent *TetheringAgent) SetUnbindHandler(handler func(UnbindCommand)) {
    agent.commandHandlers.Unbind.Store(handler)
}
func init() {

}
