package deviceorientation


import (
    "../shared"
    "sync"
    "encoding/json"
    "fmt"
)
type SetDeviceOrientationOverrideCommandFn struct {
    mux sync.RWMutex
    fn func(SetDeviceOrientationOverrideCommand)
}

func (a *SetDeviceOrientationOverrideCommandFn) Load() func(SetDeviceOrientationOverrideCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetDeviceOrientationOverrideCommandFn) Store(fn func(SetDeviceOrientationOverrideCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ClearDeviceOrientationOverrideCommandFn struct {
    mux sync.RWMutex
    fn func(ClearDeviceOrientationOverrideCommand)
}

func (a *ClearDeviceOrientationOverrideCommandFn) Load() func(ClearDeviceOrientationOverrideCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ClearDeviceOrientationOverrideCommandFn) Store(fn func(ClearDeviceOrientationOverrideCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DeviceOrientationAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        SetDeviceOrientationOverride SetDeviceOrientationOverrideCommandFn
        ClearDeviceOrientationOverride ClearDeviceOrientationOverrideCommandFn
    }
}
func NewAgent(conn *shared.Connection) *DeviceOrientationAgent {
    agent := &DeviceOrientationAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *DeviceOrientationAgent) Name() string {
    return "DeviceOrientation"
}

func (agent *DeviceOrientationAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    defer shared.TryRecoverFromPanic(agent.conn)
    switch (funcName) {
        case "setDeviceOrientationOverride":
            var out SetDeviceOrientationOverrideCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetDeviceOrientationOverride.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "clearDeviceOrientationOverride":
            var out ClearDeviceOrientationOverrideCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ClearDeviceOrientationOverride.Load()
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
func (agent *DeviceOrientationAgent) SetSetDeviceOrientationOverrideHandler(handler func(SetDeviceOrientationOverrideCommand)) {
    agent.commandHandlers.SetDeviceOrientationOverride.Store(handler)
}
func (agent *DeviceOrientationAgent) SetClearDeviceOrientationOverrideHandler(handler func(ClearDeviceOrientationOverrideCommand)) {
    agent.commandHandlers.ClearDeviceOrientationOverride.Store(handler)
}
func init() {

}
