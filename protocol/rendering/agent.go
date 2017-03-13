package rendering


import (
    "../shared"
    "sync"
    "encoding/json"
    "fmt"
)
type SetShowPaintRectsCommandFn struct {
    mux sync.RWMutex
    fn func(SetShowPaintRectsCommand)
}

func (a *SetShowPaintRectsCommandFn) Load() func(SetShowPaintRectsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetShowPaintRectsCommandFn) Store(fn func(SetShowPaintRectsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetShowDebugBordersCommandFn struct {
    mux sync.RWMutex
    fn func(SetShowDebugBordersCommand)
}

func (a *SetShowDebugBordersCommandFn) Load() func(SetShowDebugBordersCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetShowDebugBordersCommandFn) Store(fn func(SetShowDebugBordersCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetShowFPSCounterCommandFn struct {
    mux sync.RWMutex
    fn func(SetShowFPSCounterCommand)
}

func (a *SetShowFPSCounterCommandFn) Load() func(SetShowFPSCounterCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetShowFPSCounterCommandFn) Store(fn func(SetShowFPSCounterCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetShowScrollBottleneckRectsCommandFn struct {
    mux sync.RWMutex
    fn func(SetShowScrollBottleneckRectsCommand)
}

func (a *SetShowScrollBottleneckRectsCommandFn) Load() func(SetShowScrollBottleneckRectsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetShowScrollBottleneckRectsCommandFn) Store(fn func(SetShowScrollBottleneckRectsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetShowViewportSizeOnResizeCommandFn struct {
    mux sync.RWMutex
    fn func(SetShowViewportSizeOnResizeCommand)
}

func (a *SetShowViewportSizeOnResizeCommandFn) Load() func(SetShowViewportSizeOnResizeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetShowViewportSizeOnResizeCommandFn) Store(fn func(SetShowViewportSizeOnResizeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RenderingAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        SetShowPaintRects SetShowPaintRectsCommandFn
        SetShowDebugBorders SetShowDebugBordersCommandFn
        SetShowFPSCounter SetShowFPSCounterCommandFn
        SetShowScrollBottleneckRects SetShowScrollBottleneckRectsCommandFn
        SetShowViewportSizeOnResize SetShowViewportSizeOnResizeCommandFn
    }
}
func NewAgent(conn *shared.Connection) *RenderingAgent {
    agent := &RenderingAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *RenderingAgent) Name() string {
    return "Rendering"
}

func (agent *RenderingAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    defer shared.TryRecoverFromPanic(agent.conn)
    switch (funcName) {
        case "setShowPaintRects":
            var out SetShowPaintRectsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetShowPaintRects.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setShowDebugBorders":
            var out SetShowDebugBordersCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetShowDebugBorders.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setShowFPSCounter":
            var out SetShowFPSCounterCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetShowFPSCounter.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setShowScrollBottleneckRects":
            var out SetShowScrollBottleneckRectsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetShowScrollBottleneckRects.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setShowViewportSizeOnResize":
            var out SetShowViewportSizeOnResizeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetShowViewportSizeOnResize.Load()
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
func (agent *RenderingAgent) SetSetShowPaintRectsHandler(handler func(SetShowPaintRectsCommand)) {
    agent.commandHandlers.SetShowPaintRects.Store(handler)
}
func (agent *RenderingAgent) SetSetShowDebugBordersHandler(handler func(SetShowDebugBordersCommand)) {
    agent.commandHandlers.SetShowDebugBorders.Store(handler)
}
func (agent *RenderingAgent) SetSetShowFPSCounterHandler(handler func(SetShowFPSCounterCommand)) {
    agent.commandHandlers.SetShowFPSCounter.Store(handler)
}
func (agent *RenderingAgent) SetSetShowScrollBottleneckRectsHandler(handler func(SetShowScrollBottleneckRectsCommand)) {
    agent.commandHandlers.SetShowScrollBottleneckRects.Store(handler)
}
func (agent *RenderingAgent) SetSetShowViewportSizeOnResizeHandler(handler func(SetShowViewportSizeOnResizeCommand)) {
    agent.commandHandlers.SetShowViewportSizeOnResize.Store(handler)
}
func init() {

}
