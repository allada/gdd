package serviceworker


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

type UnregisterCommandFn struct {
    mux sync.RWMutex
    fn func(UnregisterCommand)
}

func (a *UnregisterCommandFn) Load() func(UnregisterCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *UnregisterCommandFn) Store(fn func(UnregisterCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type UpdateRegistrationCommandFn struct {
    mux sync.RWMutex
    fn func(UpdateRegistrationCommand)
}

func (a *UpdateRegistrationCommandFn) Load() func(UpdateRegistrationCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *UpdateRegistrationCommandFn) Store(fn func(UpdateRegistrationCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StartWorkerCommandFn struct {
    mux sync.RWMutex
    fn func(StartWorkerCommand)
}

func (a *StartWorkerCommandFn) Load() func(StartWorkerCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StartWorkerCommandFn) Store(fn func(StartWorkerCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SkipWaitingCommandFn struct {
    mux sync.RWMutex
    fn func(SkipWaitingCommand)
}

func (a *SkipWaitingCommandFn) Load() func(SkipWaitingCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SkipWaitingCommandFn) Store(fn func(SkipWaitingCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StopWorkerCommandFn struct {
    mux sync.RWMutex
    fn func(StopWorkerCommand)
}

func (a *StopWorkerCommandFn) Load() func(StopWorkerCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StopWorkerCommandFn) Store(fn func(StopWorkerCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type InspectWorkerCommandFn struct {
    mux sync.RWMutex
    fn func(InspectWorkerCommand)
}

func (a *InspectWorkerCommandFn) Load() func(InspectWorkerCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *InspectWorkerCommandFn) Store(fn func(InspectWorkerCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetForceUpdateOnPageLoadCommandFn struct {
    mux sync.RWMutex
    fn func(SetForceUpdateOnPageLoadCommand)
}

func (a *SetForceUpdateOnPageLoadCommandFn) Load() func(SetForceUpdateOnPageLoadCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetForceUpdateOnPageLoadCommandFn) Store(fn func(SetForceUpdateOnPageLoadCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DeliverPushMessageCommandFn struct {
    mux sync.RWMutex
    fn func(DeliverPushMessageCommand)
}

func (a *DeliverPushMessageCommandFn) Load() func(DeliverPushMessageCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *DeliverPushMessageCommandFn) Store(fn func(DeliverPushMessageCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DispatchSyncEventCommandFn struct {
    mux sync.RWMutex
    fn func(DispatchSyncEventCommand)
}

func (a *DispatchSyncEventCommandFn) Load() func(DispatchSyncEventCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *DispatchSyncEventCommandFn) Store(fn func(DispatchSyncEventCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ServiceWorkerAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
        Unregister UnregisterCommandFn
        UpdateRegistration UpdateRegistrationCommandFn
        StartWorker StartWorkerCommandFn
        SkipWaiting SkipWaitingCommandFn
        StopWorker StopWorkerCommandFn
        InspectWorker InspectWorkerCommandFn
        SetForceUpdateOnPageLoad SetForceUpdateOnPageLoadCommandFn
        DeliverPushMessage DeliverPushMessageCommandFn
        DispatchSyncEvent DispatchSyncEventCommandFn
    }
}
func NewAgent(conn *shared.Connection) *ServiceWorkerAgent {
    agent := &ServiceWorkerAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *ServiceWorkerAgent) Name() string {
    return "ServiceWorker"
}

func (agent *ServiceWorkerAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "unregister":
            var out UnregisterCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Unregister.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "updateRegistration":
            var out UpdateRegistrationCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.UpdateRegistration.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "startWorker":
            var out StartWorkerCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StartWorker.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "skipWaiting":
            var out SkipWaitingCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SkipWaiting.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "stopWorker":
            var out StopWorkerCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StopWorker.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "inspectWorker":
            var out InspectWorkerCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.InspectWorker.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setForceUpdateOnPageLoad":
            var out SetForceUpdateOnPageLoadCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetForceUpdateOnPageLoad.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "deliverPushMessage":
            var out DeliverPushMessageCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.DeliverPushMessage.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "dispatchSyncEvent":
            var out DispatchSyncEventCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.DispatchSyncEvent.Load()
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
func (agent *ServiceWorkerAgent) FireWorkerRegistrationUpdated(event WorkerRegistrationUpdatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "ServiceWorker.workerRegistrationUpdated",
        Params: event,
    })
}
func (agent *ServiceWorkerAgent) FireWorkerRegistrationUpdatedOnTarget(targetId string,event WorkerRegistrationUpdatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "ServiceWorker.workerRegistrationUpdated",
        Params: event,
    })
}
func (agent *ServiceWorkerAgent) FireWorkerVersionUpdated(event WorkerVersionUpdatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "ServiceWorker.workerVersionUpdated",
        Params: event,
    })
}
func (agent *ServiceWorkerAgent) FireWorkerVersionUpdatedOnTarget(targetId string,event WorkerVersionUpdatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "ServiceWorker.workerVersionUpdated",
        Params: event,
    })
}
func (agent *ServiceWorkerAgent) FireWorkerErrorReported(event WorkerErrorReportedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "ServiceWorker.workerErrorReported",
        Params: event,
    })
}
func (agent *ServiceWorkerAgent) FireWorkerErrorReportedOnTarget(targetId string,event WorkerErrorReportedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "ServiceWorker.workerErrorReported",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *ServiceWorkerAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *ServiceWorkerAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *ServiceWorkerAgent) SetUnregisterHandler(handler func(UnregisterCommand)) {
    agent.commandHandlers.Unregister.Store(handler)
}
func (agent *ServiceWorkerAgent) SetUpdateRegistrationHandler(handler func(UpdateRegistrationCommand)) {
    agent.commandHandlers.UpdateRegistration.Store(handler)
}
func (agent *ServiceWorkerAgent) SetStartWorkerHandler(handler func(StartWorkerCommand)) {
    agent.commandHandlers.StartWorker.Store(handler)
}
func (agent *ServiceWorkerAgent) SetSkipWaitingHandler(handler func(SkipWaitingCommand)) {
    agent.commandHandlers.SkipWaiting.Store(handler)
}
func (agent *ServiceWorkerAgent) SetStopWorkerHandler(handler func(StopWorkerCommand)) {
    agent.commandHandlers.StopWorker.Store(handler)
}
func (agent *ServiceWorkerAgent) SetInspectWorkerHandler(handler func(InspectWorkerCommand)) {
    agent.commandHandlers.InspectWorker.Store(handler)
}
func (agent *ServiceWorkerAgent) SetSetForceUpdateOnPageLoadHandler(handler func(SetForceUpdateOnPageLoadCommand)) {
    agent.commandHandlers.SetForceUpdateOnPageLoad.Store(handler)
}
func (agent *ServiceWorkerAgent) SetDeliverPushMessageHandler(handler func(DeliverPushMessageCommand)) {
    agent.commandHandlers.DeliverPushMessage.Store(handler)
}
func (agent *ServiceWorkerAgent) SetDispatchSyncEventHandler(handler func(DispatchSyncEventCommand)) {
    agent.commandHandlers.DispatchSyncEvent.Store(handler)
}
func init() {

}
