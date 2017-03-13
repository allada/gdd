package domdebugger


import (
    "../shared"
    "sync"
    "encoding/json"
    "fmt"
)
type SetDOMBreakpointCommandFn struct {
    mux sync.RWMutex
    fn func(SetDOMBreakpointCommand)
}

func (a *SetDOMBreakpointCommandFn) Load() func(SetDOMBreakpointCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetDOMBreakpointCommandFn) Store(fn func(SetDOMBreakpointCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RemoveDOMBreakpointCommandFn struct {
    mux sync.RWMutex
    fn func(RemoveDOMBreakpointCommand)
}

func (a *RemoveDOMBreakpointCommandFn) Load() func(RemoveDOMBreakpointCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RemoveDOMBreakpointCommandFn) Store(fn func(RemoveDOMBreakpointCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetEventListenerBreakpointCommandFn struct {
    mux sync.RWMutex
    fn func(SetEventListenerBreakpointCommand)
}

func (a *SetEventListenerBreakpointCommandFn) Load() func(SetEventListenerBreakpointCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetEventListenerBreakpointCommandFn) Store(fn func(SetEventListenerBreakpointCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RemoveEventListenerBreakpointCommandFn struct {
    mux sync.RWMutex
    fn func(RemoveEventListenerBreakpointCommand)
}

func (a *RemoveEventListenerBreakpointCommandFn) Load() func(RemoveEventListenerBreakpointCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RemoveEventListenerBreakpointCommandFn) Store(fn func(RemoveEventListenerBreakpointCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetInstrumentationBreakpointCommandFn struct {
    mux sync.RWMutex
    fn func(SetInstrumentationBreakpointCommand)
}

func (a *SetInstrumentationBreakpointCommandFn) Load() func(SetInstrumentationBreakpointCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetInstrumentationBreakpointCommandFn) Store(fn func(SetInstrumentationBreakpointCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RemoveInstrumentationBreakpointCommandFn struct {
    mux sync.RWMutex
    fn func(RemoveInstrumentationBreakpointCommand)
}

func (a *RemoveInstrumentationBreakpointCommandFn) Load() func(RemoveInstrumentationBreakpointCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RemoveInstrumentationBreakpointCommandFn) Store(fn func(RemoveInstrumentationBreakpointCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetXHRBreakpointCommandFn struct {
    mux sync.RWMutex
    fn func(SetXHRBreakpointCommand)
}

func (a *SetXHRBreakpointCommandFn) Load() func(SetXHRBreakpointCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetXHRBreakpointCommandFn) Store(fn func(SetXHRBreakpointCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RemoveXHRBreakpointCommandFn struct {
    mux sync.RWMutex
    fn func(RemoveXHRBreakpointCommand)
}

func (a *RemoveXHRBreakpointCommandFn) Load() func(RemoveXHRBreakpointCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RemoveXHRBreakpointCommandFn) Store(fn func(RemoveXHRBreakpointCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetEventListenersCommandFn struct {
    mux sync.RWMutex
    fn func(GetEventListenersCommand)
}

func (a *GetEventListenersCommandFn) Load() func(GetEventListenersCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetEventListenersCommandFn) Store(fn func(GetEventListenersCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DOMDebuggerAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        SetDOMBreakpoint SetDOMBreakpointCommandFn
        RemoveDOMBreakpoint RemoveDOMBreakpointCommandFn
        SetEventListenerBreakpoint SetEventListenerBreakpointCommandFn
        RemoveEventListenerBreakpoint RemoveEventListenerBreakpointCommandFn
        SetInstrumentationBreakpoint SetInstrumentationBreakpointCommandFn
        RemoveInstrumentationBreakpoint RemoveInstrumentationBreakpointCommandFn
        SetXHRBreakpoint SetXHRBreakpointCommandFn
        RemoveXHRBreakpoint RemoveXHRBreakpointCommandFn
        GetEventListeners GetEventListenersCommandFn
    }
}
func NewAgent(conn *shared.Connection) *DOMDebuggerAgent {
    agent := &DOMDebuggerAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *DOMDebuggerAgent) Name() string {
    return "DOMDebugger"
}

func (agent *DOMDebuggerAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "setDOMBreakpoint":
            var out SetDOMBreakpointCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetDOMBreakpoint.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "removeDOMBreakpoint":
            var out RemoveDOMBreakpointCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RemoveDOMBreakpoint.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setEventListenerBreakpoint":
            var out SetEventListenerBreakpointCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetEventListenerBreakpoint.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "removeEventListenerBreakpoint":
            var out RemoveEventListenerBreakpointCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RemoveEventListenerBreakpoint.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setInstrumentationBreakpoint":
            var out SetInstrumentationBreakpointCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetInstrumentationBreakpoint.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "removeInstrumentationBreakpoint":
            var out RemoveInstrumentationBreakpointCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RemoveInstrumentationBreakpoint.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setXHRBreakpoint":
            var out SetXHRBreakpointCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetXHRBreakpoint.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "removeXHRBreakpoint":
            var out RemoveXHRBreakpointCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RemoveXHRBreakpoint.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getEventListeners":
            var out GetEventListenersCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetEventListeners.Load()
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
func (agent *DOMDebuggerAgent) SetSetDOMBreakpointHandler(handler func(SetDOMBreakpointCommand)) {
    agent.commandHandlers.SetDOMBreakpoint.Store(handler)
}
func (agent *DOMDebuggerAgent) SetRemoveDOMBreakpointHandler(handler func(RemoveDOMBreakpointCommand)) {
    agent.commandHandlers.RemoveDOMBreakpoint.Store(handler)
}
func (agent *DOMDebuggerAgent) SetSetEventListenerBreakpointHandler(handler func(SetEventListenerBreakpointCommand)) {
    agent.commandHandlers.SetEventListenerBreakpoint.Store(handler)
}
func (agent *DOMDebuggerAgent) SetRemoveEventListenerBreakpointHandler(handler func(RemoveEventListenerBreakpointCommand)) {
    agent.commandHandlers.RemoveEventListenerBreakpoint.Store(handler)
}
func (agent *DOMDebuggerAgent) SetSetInstrumentationBreakpointHandler(handler func(SetInstrumentationBreakpointCommand)) {
    agent.commandHandlers.SetInstrumentationBreakpoint.Store(handler)
}
func (agent *DOMDebuggerAgent) SetRemoveInstrumentationBreakpointHandler(handler func(RemoveInstrumentationBreakpointCommand)) {
    agent.commandHandlers.RemoveInstrumentationBreakpoint.Store(handler)
}
func (agent *DOMDebuggerAgent) SetSetXHRBreakpointHandler(handler func(SetXHRBreakpointCommand)) {
    agent.commandHandlers.SetXHRBreakpoint.Store(handler)
}
func (agent *DOMDebuggerAgent) SetRemoveXHRBreakpointHandler(handler func(RemoveXHRBreakpointCommand)) {
    agent.commandHandlers.RemoveXHRBreakpoint.Store(handler)
}
func (agent *DOMDebuggerAgent) SetGetEventListenersHandler(handler func(GetEventListenersCommand)) {
    agent.commandHandlers.GetEventListeners.Store(handler)
}
func init() {

}
