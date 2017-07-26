package input


import (
    "github.com/allada/gdd/protocol/shared"
    "sync"
    "encoding/json"
    "fmt"
)
type DispatchKeyEventCommandFn struct {
    mux sync.RWMutex
    fn func(DispatchKeyEventCommand)
}

func (a *DispatchKeyEventCommandFn) Load() func(DispatchKeyEventCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *DispatchKeyEventCommandFn) Store(fn func(DispatchKeyEventCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DispatchMouseEventCommandFn struct {
    mux sync.RWMutex
    fn func(DispatchMouseEventCommand)
}

func (a *DispatchMouseEventCommandFn) Load() func(DispatchMouseEventCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *DispatchMouseEventCommandFn) Store(fn func(DispatchMouseEventCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DispatchTouchEventCommandFn struct {
    mux sync.RWMutex
    fn func(DispatchTouchEventCommand)
}

func (a *DispatchTouchEventCommandFn) Load() func(DispatchTouchEventCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *DispatchTouchEventCommandFn) Store(fn func(DispatchTouchEventCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type EmulateTouchFromMouseEventCommandFn struct {
    mux sync.RWMutex
    fn func(EmulateTouchFromMouseEventCommand)
}

func (a *EmulateTouchFromMouseEventCommandFn) Load() func(EmulateTouchFromMouseEventCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *EmulateTouchFromMouseEventCommandFn) Store(fn func(EmulateTouchFromMouseEventCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SynthesizePinchGestureCommandFn struct {
    mux sync.RWMutex
    fn func(SynthesizePinchGestureCommand)
}

func (a *SynthesizePinchGestureCommandFn) Load() func(SynthesizePinchGestureCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SynthesizePinchGestureCommandFn) Store(fn func(SynthesizePinchGestureCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SynthesizeScrollGestureCommandFn struct {
    mux sync.RWMutex
    fn func(SynthesizeScrollGestureCommand)
}

func (a *SynthesizeScrollGestureCommandFn) Load() func(SynthesizeScrollGestureCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SynthesizeScrollGestureCommandFn) Store(fn func(SynthesizeScrollGestureCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SynthesizeTapGestureCommandFn struct {
    mux sync.RWMutex
    fn func(SynthesizeTapGestureCommand)
}

func (a *SynthesizeTapGestureCommandFn) Load() func(SynthesizeTapGestureCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SynthesizeTapGestureCommandFn) Store(fn func(SynthesizeTapGestureCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type InputAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        DispatchKeyEvent DispatchKeyEventCommandFn
        DispatchMouseEvent DispatchMouseEventCommandFn
        DispatchTouchEvent DispatchTouchEventCommandFn
        EmulateTouchFromMouseEvent EmulateTouchFromMouseEventCommandFn
        SynthesizePinchGesture SynthesizePinchGestureCommandFn
        SynthesizeScrollGesture SynthesizeScrollGestureCommandFn
        SynthesizeTapGesture SynthesizeTapGestureCommandFn
    }
}
func NewAgent(conn *shared.Connection) *InputAgent {
    agent := &InputAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *InputAgent) Name() string {
    return "Input"
}

func (agent *InputAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    defer shared.TryRecoverFromPanic(agent.conn)
    switch (funcName) {
        case "dispatchKeyEvent":
            var out DispatchKeyEventCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.DispatchKeyEvent.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "dispatchMouseEvent":
            var out DispatchMouseEventCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.DispatchMouseEvent.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "dispatchTouchEvent":
            var out DispatchTouchEventCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.DispatchTouchEvent.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "emulateTouchFromMouseEvent":
            var out EmulateTouchFromMouseEventCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.EmulateTouchFromMouseEvent.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "synthesizePinchGesture":
            var out SynthesizePinchGestureCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SynthesizePinchGesture.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "synthesizeScrollGesture":
            var out SynthesizeScrollGestureCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SynthesizeScrollGesture.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "synthesizeTapGesture":
            var out SynthesizeTapGestureCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SynthesizeTapGesture.Load()
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
func (agent *InputAgent) SetDispatchKeyEventHandler(handler func(DispatchKeyEventCommand)) {
    agent.commandHandlers.DispatchKeyEvent.Store(handler)
}
func (agent *InputAgent) SetDispatchMouseEventHandler(handler func(DispatchMouseEventCommand)) {
    agent.commandHandlers.DispatchMouseEvent.Store(handler)
}
func (agent *InputAgent) SetDispatchTouchEventHandler(handler func(DispatchTouchEventCommand)) {
    agent.commandHandlers.DispatchTouchEvent.Store(handler)
}
func (agent *InputAgent) SetEmulateTouchFromMouseEventHandler(handler func(EmulateTouchFromMouseEventCommand)) {
    agent.commandHandlers.EmulateTouchFromMouseEvent.Store(handler)
}
func (agent *InputAgent) SetSynthesizePinchGestureHandler(handler func(SynthesizePinchGestureCommand)) {
    agent.commandHandlers.SynthesizePinchGesture.Store(handler)
}
func (agent *InputAgent) SetSynthesizeScrollGestureHandler(handler func(SynthesizeScrollGestureCommand)) {
    agent.commandHandlers.SynthesizeScrollGesture.Store(handler)
}
func (agent *InputAgent) SetSynthesizeTapGestureHandler(handler func(SynthesizeTapGestureCommand)) {
    agent.commandHandlers.SynthesizeTapGesture.Store(handler)
}
func init() {

}
