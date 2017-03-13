package domstorage


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

type ClearCommandFn struct {
    mux sync.RWMutex
    fn func(ClearCommand)
}

func (a *ClearCommandFn) Load() func(ClearCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ClearCommandFn) Store(fn func(ClearCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetDOMStorageItemsCommandFn struct {
    mux sync.RWMutex
    fn func(GetDOMStorageItemsCommand)
}

func (a *GetDOMStorageItemsCommandFn) Load() func(GetDOMStorageItemsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetDOMStorageItemsCommandFn) Store(fn func(GetDOMStorageItemsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetDOMStorageItemCommandFn struct {
    mux sync.RWMutex
    fn func(SetDOMStorageItemCommand)
}

func (a *SetDOMStorageItemCommandFn) Load() func(SetDOMStorageItemCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetDOMStorageItemCommandFn) Store(fn func(SetDOMStorageItemCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RemoveDOMStorageItemCommandFn struct {
    mux sync.RWMutex
    fn func(RemoveDOMStorageItemCommand)
}

func (a *RemoveDOMStorageItemCommandFn) Load() func(RemoveDOMStorageItemCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RemoveDOMStorageItemCommandFn) Store(fn func(RemoveDOMStorageItemCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DOMStorageAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
        Clear ClearCommandFn
        GetDOMStorageItems GetDOMStorageItemsCommandFn
        SetDOMStorageItem SetDOMStorageItemCommandFn
        RemoveDOMStorageItem RemoveDOMStorageItemCommandFn
    }
}
func NewAgent(conn *shared.Connection) *DOMStorageAgent {
    agent := &DOMStorageAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *DOMStorageAgent) Name() string {
    return "DOMStorage"
}

func (agent *DOMStorageAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "clear":
            var out ClearCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Clear.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getDOMStorageItems":
            var out GetDOMStorageItemsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetDOMStorageItems.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setDOMStorageItem":
            var out SetDOMStorageItemCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetDOMStorageItem.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "removeDOMStorageItem":
            var out RemoveDOMStorageItemCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RemoveDOMStorageItem.Load()
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
func (agent *DOMStorageAgent) FireDomStorageItemsCleared(event DomStorageItemsClearedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOMStorage.domStorageItemsCleared",
        Params: event,
    })
}
func (agent *DOMStorageAgent) FireDomStorageItemsClearedOnTarget(targetId string,event DomStorageItemsClearedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOMStorage.domStorageItemsCleared",
        Params: event,
    })
}
func (agent *DOMStorageAgent) FireDomStorageItemRemoved(event DomStorageItemRemovedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOMStorage.domStorageItemRemoved",
        Params: event,
    })
}
func (agent *DOMStorageAgent) FireDomStorageItemRemovedOnTarget(targetId string,event DomStorageItemRemovedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOMStorage.domStorageItemRemoved",
        Params: event,
    })
}
func (agent *DOMStorageAgent) FireDomStorageItemAdded(event DomStorageItemAddedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOMStorage.domStorageItemAdded",
        Params: event,
    })
}
func (agent *DOMStorageAgent) FireDomStorageItemAddedOnTarget(targetId string,event DomStorageItemAddedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOMStorage.domStorageItemAdded",
        Params: event,
    })
}
func (agent *DOMStorageAgent) FireDomStorageItemUpdated(event DomStorageItemUpdatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOMStorage.domStorageItemUpdated",
        Params: event,
    })
}
func (agent *DOMStorageAgent) FireDomStorageItemUpdatedOnTarget(targetId string,event DomStorageItemUpdatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOMStorage.domStorageItemUpdated",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *DOMStorageAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *DOMStorageAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *DOMStorageAgent) SetClearHandler(handler func(ClearCommand)) {
    agent.commandHandlers.Clear.Store(handler)
}
func (agent *DOMStorageAgent) SetGetDOMStorageItemsHandler(handler func(GetDOMStorageItemsCommand)) {
    agent.commandHandlers.GetDOMStorageItems.Store(handler)
}
func (agent *DOMStorageAgent) SetSetDOMStorageItemHandler(handler func(SetDOMStorageItemCommand)) {
    agent.commandHandlers.SetDOMStorageItem.Store(handler)
}
func (agent *DOMStorageAgent) SetRemoveDOMStorageItemHandler(handler func(RemoveDOMStorageItemCommand)) {
    agent.commandHandlers.RemoveDOMStorageItem.Store(handler)
}
func init() {

}
