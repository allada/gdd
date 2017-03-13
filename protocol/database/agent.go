package database


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

type GetDatabaseTableNamesCommandFn struct {
    mux sync.RWMutex
    fn func(GetDatabaseTableNamesCommand)
}

func (a *GetDatabaseTableNamesCommandFn) Load() func(GetDatabaseTableNamesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetDatabaseTableNamesCommandFn) Store(fn func(GetDatabaseTableNamesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ExecuteSQLCommandFn struct {
    mux sync.RWMutex
    fn func(ExecuteSQLCommand)
}

func (a *ExecuteSQLCommandFn) Load() func(ExecuteSQLCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ExecuteSQLCommandFn) Store(fn func(ExecuteSQLCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DatabaseAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
        GetDatabaseTableNames GetDatabaseTableNamesCommandFn
        ExecuteSQL ExecuteSQLCommandFn
    }
}
func NewAgent(conn *shared.Connection) *DatabaseAgent {
    agent := &DatabaseAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *DatabaseAgent) Name() string {
    return "Database"
}

func (agent *DatabaseAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "getDatabaseTableNames":
            var out GetDatabaseTableNamesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetDatabaseTableNames.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "executeSQL":
            var out ExecuteSQLCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ExecuteSQL.Load()
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
func (agent *DatabaseAgent) FireAddDatabase(event AddDatabaseEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Database.addDatabase",
        Params: event,
    })
}
func (agent *DatabaseAgent) FireAddDatabaseOnTarget(targetId string,event AddDatabaseEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Database.addDatabase",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *DatabaseAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *DatabaseAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *DatabaseAgent) SetGetDatabaseTableNamesHandler(handler func(GetDatabaseTableNamesCommand)) {
    agent.commandHandlers.GetDatabaseTableNames.Store(handler)
}
func (agent *DatabaseAgent) SetExecuteSQLHandler(handler func(ExecuteSQLCommand)) {
    agent.commandHandlers.ExecuteSQL.Store(handler)
}
func init() {

}
