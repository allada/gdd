package io


import (
    "github.com/allada/gdd/protocol/shared"
    "sync"
    "encoding/json"
    "fmt"
)
type ReadCommandFn struct {
    mux sync.RWMutex
    fn func(ReadCommand)
}

func (a *ReadCommandFn) Load() func(ReadCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ReadCommandFn) Store(fn func(ReadCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CloseCommandFn struct {
    mux sync.RWMutex
    fn func(CloseCommand)
}

func (a *CloseCommandFn) Load() func(CloseCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CloseCommandFn) Store(fn func(CloseCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type IOAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Read ReadCommandFn
        Close CloseCommandFn
    }
}
func NewAgent(conn *shared.Connection) *IOAgent {
    agent := &IOAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *IOAgent) Name() string {
    return "IO"
}

func (agent *IOAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    defer shared.TryRecoverFromPanic(agent.conn)
    switch (funcName) {
        case "read":
            var out ReadCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Read.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "close":
            var out CloseCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Close.Load()
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
func (agent *IOAgent) SetReadHandler(handler func(ReadCommand)) {
    agent.commandHandlers.Read.Store(handler)
}
func (agent *IOAgent) SetCloseHandler(handler func(CloseCommand)) {
    agent.commandHandlers.Close.Store(handler)
}
func init() {

}
