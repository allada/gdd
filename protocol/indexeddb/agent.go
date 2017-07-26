package indexeddb


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

type RequestDatabaseNamesCommandFn struct {
    mux sync.RWMutex
    fn func(RequestDatabaseNamesCommand)
}

func (a *RequestDatabaseNamesCommandFn) Load() func(RequestDatabaseNamesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RequestDatabaseNamesCommandFn) Store(fn func(RequestDatabaseNamesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RequestDatabaseCommandFn struct {
    mux sync.RWMutex
    fn func(RequestDatabaseCommand)
}

func (a *RequestDatabaseCommandFn) Load() func(RequestDatabaseCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RequestDatabaseCommandFn) Store(fn func(RequestDatabaseCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RequestDataCommandFn struct {
    mux sync.RWMutex
    fn func(RequestDataCommand)
}

func (a *RequestDataCommandFn) Load() func(RequestDataCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RequestDataCommandFn) Store(fn func(RequestDataCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ClearObjectStoreCommandFn struct {
    mux sync.RWMutex
    fn func(ClearObjectStoreCommand)
}

func (a *ClearObjectStoreCommandFn) Load() func(ClearObjectStoreCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ClearObjectStoreCommandFn) Store(fn func(ClearObjectStoreCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DeleteDatabaseCommandFn struct {
    mux sync.RWMutex
    fn func(DeleteDatabaseCommand)
}

func (a *DeleteDatabaseCommandFn) Load() func(DeleteDatabaseCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *DeleteDatabaseCommandFn) Store(fn func(DeleteDatabaseCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type IndexedDBAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
        RequestDatabaseNames RequestDatabaseNamesCommandFn
        RequestDatabase RequestDatabaseCommandFn
        RequestData RequestDataCommandFn
        ClearObjectStore ClearObjectStoreCommandFn
        DeleteDatabase DeleteDatabaseCommandFn
    }
}
func NewAgent(conn *shared.Connection) *IndexedDBAgent {
    agent := &IndexedDBAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *IndexedDBAgent) Name() string {
    return "IndexedDB"
}

func (agent *IndexedDBAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "requestDatabaseNames":
            var out RequestDatabaseNamesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RequestDatabaseNames.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "requestDatabase":
            var out RequestDatabaseCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RequestDatabase.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "requestData":
            var out RequestDataCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RequestData.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "clearObjectStore":
            var out ClearObjectStoreCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ClearObjectStore.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "deleteDatabase":
            var out DeleteDatabaseCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.DeleteDatabase.Load()
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
func (agent *IndexedDBAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *IndexedDBAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *IndexedDBAgent) SetRequestDatabaseNamesHandler(handler func(RequestDatabaseNamesCommand)) {
    agent.commandHandlers.RequestDatabaseNames.Store(handler)
}
func (agent *IndexedDBAgent) SetRequestDatabaseHandler(handler func(RequestDatabaseCommand)) {
    agent.commandHandlers.RequestDatabase.Store(handler)
}
func (agent *IndexedDBAgent) SetRequestDataHandler(handler func(RequestDataCommand)) {
    agent.commandHandlers.RequestData.Store(handler)
}
func (agent *IndexedDBAgent) SetClearObjectStoreHandler(handler func(ClearObjectStoreCommand)) {
    agent.commandHandlers.ClearObjectStore.Store(handler)
}
func (agent *IndexedDBAgent) SetDeleteDatabaseHandler(handler func(DeleteDatabaseCommand)) {
    agent.commandHandlers.DeleteDatabase.Store(handler)
}
func init() {

}
