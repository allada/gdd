package accessibility


import (
    "github.com/allada/gdd/protocol/shared"
    "sync"
    "encoding/json"
    "fmt"
)
type GetPartialAXTreeCommandFn struct {
    mux sync.RWMutex
    fn func(GetPartialAXTreeCommand)
}

func (a *GetPartialAXTreeCommandFn) Load() func(GetPartialAXTreeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetPartialAXTreeCommandFn) Store(fn func(GetPartialAXTreeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type AccessibilityAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        GetPartialAXTree GetPartialAXTreeCommandFn
    }
}
func NewAgent(conn *shared.Connection) *AccessibilityAgent {
    agent := &AccessibilityAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *AccessibilityAgent) Name() string {
    return "Accessibility"
}

func (agent *AccessibilityAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    defer shared.TryRecoverFromPanic(agent.conn)
    switch (funcName) {
        case "getPartialAXTree":
            var out GetPartialAXTreeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetPartialAXTree.Load()
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
func (agent *AccessibilityAgent) SetGetPartialAXTreeHandler(handler func(GetPartialAXTreeCommand)) {
    agent.commandHandlers.GetPartialAXTree.Store(handler)
}
func init() {

}
