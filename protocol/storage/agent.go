package storage


import (
    "../shared"
    "sync"
    "encoding/json"
    "fmt"
)
type ClearDataForOriginCommandFn struct {
    mux sync.RWMutex
    fn func(ClearDataForOriginCommand)
}

func (a *ClearDataForOriginCommandFn) Load() func(ClearDataForOriginCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ClearDataForOriginCommandFn) Store(fn func(ClearDataForOriginCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StorageAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        ClearDataForOrigin ClearDataForOriginCommandFn
    }
}
func NewAgent(conn *shared.Connection) *StorageAgent {
    agent := &StorageAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *StorageAgent) Name() string {
    return "Storage"
}

func (agent *StorageAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    defer shared.TryRecoverFromPanic(agent.conn)
    switch (funcName) {
        case "clearDataForOrigin":
            var out ClearDataForOriginCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ClearDataForOrigin.Load()
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
func (agent *StorageAgent) SetClearDataForOriginHandler(handler func(ClearDataForOriginCommand)) {
    agent.commandHandlers.ClearDataForOrigin.Store(handler)
}
func init() {

}
