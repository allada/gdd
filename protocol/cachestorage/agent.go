package cachestorage


import (
    "github.com/allada/gdd/protocol/shared"
    "sync"
    "encoding/json"
    "fmt"
)
type RequestCacheNamesCommandFn struct {
    mux sync.RWMutex
    fn func(RequestCacheNamesCommand)
}

func (a *RequestCacheNamesCommandFn) Load() func(RequestCacheNamesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RequestCacheNamesCommandFn) Store(fn func(RequestCacheNamesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RequestEntriesCommandFn struct {
    mux sync.RWMutex
    fn func(RequestEntriesCommand)
}

func (a *RequestEntriesCommandFn) Load() func(RequestEntriesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RequestEntriesCommandFn) Store(fn func(RequestEntriesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DeleteCacheCommandFn struct {
    mux sync.RWMutex
    fn func(DeleteCacheCommand)
}

func (a *DeleteCacheCommandFn) Load() func(DeleteCacheCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *DeleteCacheCommandFn) Store(fn func(DeleteCacheCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DeleteEntryCommandFn struct {
    mux sync.RWMutex
    fn func(DeleteEntryCommand)
}

func (a *DeleteEntryCommandFn) Load() func(DeleteEntryCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *DeleteEntryCommandFn) Store(fn func(DeleteEntryCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CacheStorageAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        RequestCacheNames RequestCacheNamesCommandFn
        RequestEntries RequestEntriesCommandFn
        DeleteCache DeleteCacheCommandFn
        DeleteEntry DeleteEntryCommandFn
    }
}
func NewAgent(conn *shared.Connection) *CacheStorageAgent {
    agent := &CacheStorageAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *CacheStorageAgent) Name() string {
    return "CacheStorage"
}

func (agent *CacheStorageAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    defer shared.TryRecoverFromPanic(agent.conn)
    switch (funcName) {
        case "requestCacheNames":
            var out RequestCacheNamesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RequestCacheNames.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "requestEntries":
            var out RequestEntriesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RequestEntries.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "deleteCache":
            var out DeleteCacheCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.DeleteCache.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "deleteEntry":
            var out DeleteEntryCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.DeleteEntry.Load()
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
func (agent *CacheStorageAgent) SetRequestCacheNamesHandler(handler func(RequestCacheNamesCommand)) {
    agent.commandHandlers.RequestCacheNames.Store(handler)
}
func (agent *CacheStorageAgent) SetRequestEntriesHandler(handler func(RequestEntriesCommand)) {
    agent.commandHandlers.RequestEntries.Store(handler)
}
func (agent *CacheStorageAgent) SetDeleteCacheHandler(handler func(DeleteCacheCommand)) {
    agent.commandHandlers.DeleteCache.Store(handler)
}
func (agent *CacheStorageAgent) SetDeleteEntryHandler(handler func(DeleteEntryCommand)) {
    agent.commandHandlers.DeleteEntry.Store(handler)
}
func init() {

}
