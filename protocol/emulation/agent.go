package emulation


import (
    "github.com/allada/gdd/protocol/shared"
    "sync"
    "encoding/json"
    "fmt"
)
type SetDeviceMetricsOverrideCommandFn struct {
    mux sync.RWMutex
    fn func(SetDeviceMetricsOverrideCommand)
}

func (a *SetDeviceMetricsOverrideCommandFn) Load() func(SetDeviceMetricsOverrideCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetDeviceMetricsOverrideCommandFn) Store(fn func(SetDeviceMetricsOverrideCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ClearDeviceMetricsOverrideCommandFn struct {
    mux sync.RWMutex
    fn func(ClearDeviceMetricsOverrideCommand)
}

func (a *ClearDeviceMetricsOverrideCommandFn) Load() func(ClearDeviceMetricsOverrideCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ClearDeviceMetricsOverrideCommandFn) Store(fn func(ClearDeviceMetricsOverrideCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ForceViewportCommandFn struct {
    mux sync.RWMutex
    fn func(ForceViewportCommand)
}

func (a *ForceViewportCommandFn) Load() func(ForceViewportCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ForceViewportCommandFn) Store(fn func(ForceViewportCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ResetViewportCommandFn struct {
    mux sync.RWMutex
    fn func(ResetViewportCommand)
}

func (a *ResetViewportCommandFn) Load() func(ResetViewportCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ResetViewportCommandFn) Store(fn func(ResetViewportCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ResetPageScaleFactorCommandFn struct {
    mux sync.RWMutex
    fn func(ResetPageScaleFactorCommand)
}

func (a *ResetPageScaleFactorCommandFn) Load() func(ResetPageScaleFactorCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ResetPageScaleFactorCommandFn) Store(fn func(ResetPageScaleFactorCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetPageScaleFactorCommandFn struct {
    mux sync.RWMutex
    fn func(SetPageScaleFactorCommand)
}

func (a *SetPageScaleFactorCommandFn) Load() func(SetPageScaleFactorCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetPageScaleFactorCommandFn) Store(fn func(SetPageScaleFactorCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetVisibleSizeCommandFn struct {
    mux sync.RWMutex
    fn func(SetVisibleSizeCommand)
}

func (a *SetVisibleSizeCommandFn) Load() func(SetVisibleSizeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetVisibleSizeCommandFn) Store(fn func(SetVisibleSizeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetScriptExecutionDisabledCommandFn struct {
    mux sync.RWMutex
    fn func(SetScriptExecutionDisabledCommand)
}

func (a *SetScriptExecutionDisabledCommandFn) Load() func(SetScriptExecutionDisabledCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetScriptExecutionDisabledCommandFn) Store(fn func(SetScriptExecutionDisabledCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetGeolocationOverrideCommandFn struct {
    mux sync.RWMutex
    fn func(SetGeolocationOverrideCommand)
}

func (a *SetGeolocationOverrideCommandFn) Load() func(SetGeolocationOverrideCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetGeolocationOverrideCommandFn) Store(fn func(SetGeolocationOverrideCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ClearGeolocationOverrideCommandFn struct {
    mux sync.RWMutex
    fn func(ClearGeolocationOverrideCommand)
}

func (a *ClearGeolocationOverrideCommandFn) Load() func(ClearGeolocationOverrideCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ClearGeolocationOverrideCommandFn) Store(fn func(ClearGeolocationOverrideCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetTouchEmulationEnabledCommandFn struct {
    mux sync.RWMutex
    fn func(SetTouchEmulationEnabledCommand)
}

func (a *SetTouchEmulationEnabledCommandFn) Load() func(SetTouchEmulationEnabledCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetTouchEmulationEnabledCommandFn) Store(fn func(SetTouchEmulationEnabledCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetEmulatedMediaCommandFn struct {
    mux sync.RWMutex
    fn func(SetEmulatedMediaCommand)
}

func (a *SetEmulatedMediaCommandFn) Load() func(SetEmulatedMediaCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetEmulatedMediaCommandFn) Store(fn func(SetEmulatedMediaCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetCPUThrottlingRateCommandFn struct {
    mux sync.RWMutex
    fn func(SetCPUThrottlingRateCommand)
}

func (a *SetCPUThrottlingRateCommandFn) Load() func(SetCPUThrottlingRateCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetCPUThrottlingRateCommandFn) Store(fn func(SetCPUThrottlingRateCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CanEmulateCommandFn struct {
    mux sync.RWMutex
    fn func(CanEmulateCommand)
}

func (a *CanEmulateCommandFn) Load() func(CanEmulateCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CanEmulateCommandFn) Store(fn func(CanEmulateCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetVirtualTimePolicyCommandFn struct {
    mux sync.RWMutex
    fn func(SetVirtualTimePolicyCommand)
}

func (a *SetVirtualTimePolicyCommandFn) Load() func(SetVirtualTimePolicyCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetVirtualTimePolicyCommandFn) Store(fn func(SetVirtualTimePolicyCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetDefaultBackgroundColorOverrideCommandFn struct {
    mux sync.RWMutex
    fn func(SetDefaultBackgroundColorOverrideCommand)
}

func (a *SetDefaultBackgroundColorOverrideCommandFn) Load() func(SetDefaultBackgroundColorOverrideCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetDefaultBackgroundColorOverrideCommandFn) Store(fn func(SetDefaultBackgroundColorOverrideCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type EmulationAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        SetDeviceMetricsOverride SetDeviceMetricsOverrideCommandFn
        ClearDeviceMetricsOverride ClearDeviceMetricsOverrideCommandFn
        ForceViewport ForceViewportCommandFn
        ResetViewport ResetViewportCommandFn
        ResetPageScaleFactor ResetPageScaleFactorCommandFn
        SetPageScaleFactor SetPageScaleFactorCommandFn
        SetVisibleSize SetVisibleSizeCommandFn
        SetScriptExecutionDisabled SetScriptExecutionDisabledCommandFn
        SetGeolocationOverride SetGeolocationOverrideCommandFn
        ClearGeolocationOverride ClearGeolocationOverrideCommandFn
        SetTouchEmulationEnabled SetTouchEmulationEnabledCommandFn
        SetEmulatedMedia SetEmulatedMediaCommandFn
        SetCPUThrottlingRate SetCPUThrottlingRateCommandFn
        CanEmulate CanEmulateCommandFn
        SetVirtualTimePolicy SetVirtualTimePolicyCommandFn
        SetDefaultBackgroundColorOverride SetDefaultBackgroundColorOverrideCommandFn
    }
}
func NewAgent(conn *shared.Connection) *EmulationAgent {
    agent := &EmulationAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *EmulationAgent) Name() string {
    return "Emulation"
}

func (agent *EmulationAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    defer shared.TryRecoverFromPanic(agent.conn)
    switch (funcName) {
        case "setDeviceMetricsOverride":
            var out SetDeviceMetricsOverrideCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetDeviceMetricsOverride.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "clearDeviceMetricsOverride":
            var out ClearDeviceMetricsOverrideCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ClearDeviceMetricsOverride.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "forceViewport":
            var out ForceViewportCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ForceViewport.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "resetViewport":
            var out ResetViewportCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ResetViewport.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "resetPageScaleFactor":
            var out ResetPageScaleFactorCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ResetPageScaleFactor.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setPageScaleFactor":
            var out SetPageScaleFactorCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetPageScaleFactor.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setVisibleSize":
            var out SetVisibleSizeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetVisibleSize.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setScriptExecutionDisabled":
            var out SetScriptExecutionDisabledCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetScriptExecutionDisabled.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setGeolocationOverride":
            var out SetGeolocationOverrideCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetGeolocationOverride.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "clearGeolocationOverride":
            var out ClearGeolocationOverrideCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ClearGeolocationOverride.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setTouchEmulationEnabled":
            var out SetTouchEmulationEnabledCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetTouchEmulationEnabled.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setEmulatedMedia":
            var out SetEmulatedMediaCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetEmulatedMedia.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setCPUThrottlingRate":
            var out SetCPUThrottlingRateCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetCPUThrottlingRate.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "canEmulate":
            var out CanEmulateCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CanEmulate.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setVirtualTimePolicy":
            var out SetVirtualTimePolicyCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetVirtualTimePolicy.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setDefaultBackgroundColorOverride":
            var out SetDefaultBackgroundColorOverrideCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetDefaultBackgroundColorOverride.Load()
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
func (agent *EmulationAgent) FireVirtualTimeBudgetExpired() {
    agent.conn.Send(shared.Notification{
        Method: "Emulation.virtualTimeBudgetExpired",
    })
}
func (agent *EmulationAgent) FireVirtualTimeBudgetExpiredOnTarget(targetId string) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Emulation.virtualTimeBudgetExpired",
    })
}

// Commands Sent From Frontend
func (agent *EmulationAgent) SetSetDeviceMetricsOverrideHandler(handler func(SetDeviceMetricsOverrideCommand)) {
    agent.commandHandlers.SetDeviceMetricsOverride.Store(handler)
}
func (agent *EmulationAgent) SetClearDeviceMetricsOverrideHandler(handler func(ClearDeviceMetricsOverrideCommand)) {
    agent.commandHandlers.ClearDeviceMetricsOverride.Store(handler)
}
func (agent *EmulationAgent) SetForceViewportHandler(handler func(ForceViewportCommand)) {
    agent.commandHandlers.ForceViewport.Store(handler)
}
func (agent *EmulationAgent) SetResetViewportHandler(handler func(ResetViewportCommand)) {
    agent.commandHandlers.ResetViewport.Store(handler)
}
func (agent *EmulationAgent) SetResetPageScaleFactorHandler(handler func(ResetPageScaleFactorCommand)) {
    agent.commandHandlers.ResetPageScaleFactor.Store(handler)
}
func (agent *EmulationAgent) SetSetPageScaleFactorHandler(handler func(SetPageScaleFactorCommand)) {
    agent.commandHandlers.SetPageScaleFactor.Store(handler)
}
func (agent *EmulationAgent) SetSetVisibleSizeHandler(handler func(SetVisibleSizeCommand)) {
    agent.commandHandlers.SetVisibleSize.Store(handler)
}
func (agent *EmulationAgent) SetSetScriptExecutionDisabledHandler(handler func(SetScriptExecutionDisabledCommand)) {
    agent.commandHandlers.SetScriptExecutionDisabled.Store(handler)
}
func (agent *EmulationAgent) SetSetGeolocationOverrideHandler(handler func(SetGeolocationOverrideCommand)) {
    agent.commandHandlers.SetGeolocationOverride.Store(handler)
}
func (agent *EmulationAgent) SetClearGeolocationOverrideHandler(handler func(ClearGeolocationOverrideCommand)) {
    agent.commandHandlers.ClearGeolocationOverride.Store(handler)
}
func (agent *EmulationAgent) SetSetTouchEmulationEnabledHandler(handler func(SetTouchEmulationEnabledCommand)) {
    agent.commandHandlers.SetTouchEmulationEnabled.Store(handler)
}
func (agent *EmulationAgent) SetSetEmulatedMediaHandler(handler func(SetEmulatedMediaCommand)) {
    agent.commandHandlers.SetEmulatedMedia.Store(handler)
}
func (agent *EmulationAgent) SetSetCPUThrottlingRateHandler(handler func(SetCPUThrottlingRateCommand)) {
    agent.commandHandlers.SetCPUThrottlingRate.Store(handler)
}
func (agent *EmulationAgent) SetCanEmulateHandler(handler func(CanEmulateCommand)) {
    agent.commandHandlers.CanEmulate.Store(handler)
}
func (agent *EmulationAgent) SetSetVirtualTimePolicyHandler(handler func(SetVirtualTimePolicyCommand)) {
    agent.commandHandlers.SetVirtualTimePolicy.Store(handler)
}
func (agent *EmulationAgent) SetSetDefaultBackgroundColorOverrideHandler(handler func(SetDefaultBackgroundColorOverrideCommand)) {
    agent.commandHandlers.SetDefaultBackgroundColorOverride.Store(handler)
}
func init() {

}
