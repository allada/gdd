package schema


import (
    "github.com/allada/gdd/protocol/shared"
    "sync"
    "encoding/json"
    "fmt"
)
type GetDomainsCommandFn struct {
    mux sync.RWMutex
    fn func(GetDomainsCommand)
}

func (a *GetDomainsCommandFn) Load() func(GetDomainsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetDomainsCommandFn) Store(fn func(GetDomainsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SchemaAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        GetDomains GetDomainsCommandFn
    }
}
func NewAgent(conn *shared.Connection) *SchemaAgent {
    agent := &SchemaAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *SchemaAgent) Name() string {
    return "Schema"
}

func (agent *SchemaAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    defer shared.TryRecoverFromPanic(agent.conn)
    switch (funcName) {
        case "getDomains":
            var out GetDomainsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetDomains.Load()
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
func (agent *SchemaAgent) SetGetDomainsHandler(handler func(GetDomainsCommand)) {
    agent.commandHandlers.GetDomains.Store(handler)
}
func init() {

}
