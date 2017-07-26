package inspector


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

type InspectorAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
    }
}
func NewAgent(conn *shared.Connection) *InspectorAgent {
    agent := &InspectorAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *InspectorAgent) Name() string {
    return "Inspector"
}

func (agent *InspectorAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        default:
            fmt.Printf("Command %s unknown\n", funcName)
            agent.conn.SendErrorResult(id, targetId, shared.ErrorCodeMethodNotFound, "")
    }
}

// Dispatchable Events
func (agent *InspectorAgent) FireDetached(event DetachedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Inspector.detached",
        Params: event,
    })
}
func (agent *InspectorAgent) FireDetachedOnTarget(targetId string,event DetachedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Inspector.detached",
        Params: event,
    })
}
func (agent *InspectorAgent) FireTargetCrashed() {
    agent.conn.Send(shared.Notification{
        Method: "Inspector.targetCrashed",
    })
}
func (agent *InspectorAgent) FireTargetCrashedOnTarget(targetId string) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Inspector.targetCrashed",
    })
}

// Commands Sent From Frontend
func (agent *InspectorAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *InspectorAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func init() {

}
