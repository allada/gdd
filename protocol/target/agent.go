package target


import (
    "../shared"
    "sync"
    "encoding/json"
    "fmt"
)
type SetDiscoverTargetsCommandFn struct {
    mux sync.RWMutex
    fn func(SetDiscoverTargetsCommand)
}

func (a *SetDiscoverTargetsCommandFn) Load() func(SetDiscoverTargetsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetDiscoverTargetsCommandFn) Store(fn func(SetDiscoverTargetsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetAutoAttachCommandFn struct {
    mux sync.RWMutex
    fn func(SetAutoAttachCommand)
}

func (a *SetAutoAttachCommandFn) Load() func(SetAutoAttachCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetAutoAttachCommandFn) Store(fn func(SetAutoAttachCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetAttachToFramesCommandFn struct {
    mux sync.RWMutex
    fn func(SetAttachToFramesCommand)
}

func (a *SetAttachToFramesCommandFn) Load() func(SetAttachToFramesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetAttachToFramesCommandFn) Store(fn func(SetAttachToFramesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetRemoteLocationsCommandFn struct {
    mux sync.RWMutex
    fn func(SetRemoteLocationsCommand)
}

func (a *SetRemoteLocationsCommandFn) Load() func(SetRemoteLocationsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetRemoteLocationsCommandFn) Store(fn func(SetRemoteLocationsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SendMessageToTargetCommandFn struct {
    mux sync.RWMutex
    fn func(SendMessageToTargetCommand)
}

func (a *SendMessageToTargetCommandFn) Load() func(SendMessageToTargetCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SendMessageToTargetCommandFn) Store(fn func(SendMessageToTargetCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetTargetInfoCommandFn struct {
    mux sync.RWMutex
    fn func(GetTargetInfoCommand)
}

func (a *GetTargetInfoCommandFn) Load() func(GetTargetInfoCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetTargetInfoCommandFn) Store(fn func(GetTargetInfoCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ActivateTargetCommandFn struct {
    mux sync.RWMutex
    fn func(ActivateTargetCommand)
}

func (a *ActivateTargetCommandFn) Load() func(ActivateTargetCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ActivateTargetCommandFn) Store(fn func(ActivateTargetCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CloseTargetCommandFn struct {
    mux sync.RWMutex
    fn func(CloseTargetCommand)
}

func (a *CloseTargetCommandFn) Load() func(CloseTargetCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CloseTargetCommandFn) Store(fn func(CloseTargetCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type AttachToTargetCommandFn struct {
    mux sync.RWMutex
    fn func(AttachToTargetCommand)
}

func (a *AttachToTargetCommandFn) Load() func(AttachToTargetCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *AttachToTargetCommandFn) Store(fn func(AttachToTargetCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DetachFromTargetCommandFn struct {
    mux sync.RWMutex
    fn func(DetachFromTargetCommand)
}

func (a *DetachFromTargetCommandFn) Load() func(DetachFromTargetCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *DetachFromTargetCommandFn) Store(fn func(DetachFromTargetCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CreateBrowserContextCommandFn struct {
    mux sync.RWMutex
    fn func(CreateBrowserContextCommand)
}

func (a *CreateBrowserContextCommandFn) Load() func(CreateBrowserContextCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CreateBrowserContextCommandFn) Store(fn func(CreateBrowserContextCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DisposeBrowserContextCommandFn struct {
    mux sync.RWMutex
    fn func(DisposeBrowserContextCommand)
}

func (a *DisposeBrowserContextCommandFn) Load() func(DisposeBrowserContextCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *DisposeBrowserContextCommandFn) Store(fn func(DisposeBrowserContextCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CreateTargetCommandFn struct {
    mux sync.RWMutex
    fn func(CreateTargetCommand)
}

func (a *CreateTargetCommandFn) Load() func(CreateTargetCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CreateTargetCommandFn) Store(fn func(CreateTargetCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetTargetsCommandFn struct {
    mux sync.RWMutex
    fn func(GetTargetsCommand)
}

func (a *GetTargetsCommandFn) Load() func(GetTargetsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetTargetsCommandFn) Store(fn func(GetTargetsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type TargetAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        SetDiscoverTargets SetDiscoverTargetsCommandFn
        SetAutoAttach SetAutoAttachCommandFn
        SetAttachToFrames SetAttachToFramesCommandFn
        SetRemoteLocations SetRemoteLocationsCommandFn
        SendMessageToTarget SendMessageToTargetCommandFn
        GetTargetInfo GetTargetInfoCommandFn
        ActivateTarget ActivateTargetCommandFn
        CloseTarget CloseTargetCommandFn
        AttachToTarget AttachToTargetCommandFn
        DetachFromTarget DetachFromTargetCommandFn
        CreateBrowserContext CreateBrowserContextCommandFn
        DisposeBrowserContext DisposeBrowserContextCommandFn
        CreateTarget CreateTargetCommandFn
        GetTargets GetTargetsCommandFn
    }
}
func NewAgent(conn *shared.Connection) *TargetAgent {
    agent := &TargetAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *TargetAgent) Name() string {
    return "Target"
}

func (agent *TargetAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    defer shared.TryRecoverFromPanic(agent.conn)
    switch (funcName) {
        case "setDiscoverTargets":
            var out SetDiscoverTargetsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetDiscoverTargets.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setAutoAttach":
            var out SetAutoAttachCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetAutoAttach.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setAttachToFrames":
            var out SetAttachToFramesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetAttachToFrames.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setRemoteLocations":
            var out SetRemoteLocationsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetRemoteLocations.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "sendMessageToTarget":
            var out SendMessageToTargetCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SendMessageToTarget.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getTargetInfo":
            var out GetTargetInfoCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetTargetInfo.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "activateTarget":
            var out ActivateTargetCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ActivateTarget.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "closeTarget":
            var out CloseTargetCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CloseTarget.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "attachToTarget":
            var out AttachToTargetCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.AttachToTarget.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "detachFromTarget":
            var out DetachFromTargetCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.DetachFromTarget.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "createBrowserContext":
            var out CreateBrowserContextCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CreateBrowserContext.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "disposeBrowserContext":
            var out DisposeBrowserContextCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.DisposeBrowserContext.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "createTarget":
            var out CreateTargetCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CreateTarget.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getTargets":
            var out GetTargetsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetTargets.Load()
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
func (agent *TargetAgent) FireTargetCreated(event TargetCreatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Target.targetCreated",
        Params: event,
    })
}
func (agent *TargetAgent) FireTargetCreatedOnTarget(targetId string,event TargetCreatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Target.targetCreated",
        Params: event,
    })
}
func (agent *TargetAgent) FireTargetDestroyed(event TargetDestroyedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Target.targetDestroyed",
        Params: event,
    })
}
func (agent *TargetAgent) FireTargetDestroyedOnTarget(targetId string,event TargetDestroyedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Target.targetDestroyed",
        Params: event,
    })
}
func (agent *TargetAgent) FireAttachedToTarget(event AttachedToTargetEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Target.attachedToTarget",
        Params: event,
    })
}
func (agent *TargetAgent) FireAttachedToTargetOnTarget(targetId string,event AttachedToTargetEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Target.attachedToTarget",
        Params: event,
    })
}
func (agent *TargetAgent) FireDetachedFromTarget(event DetachedFromTargetEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Target.detachedFromTarget",
        Params: event,
    })
}
func (agent *TargetAgent) FireDetachedFromTargetOnTarget(targetId string,event DetachedFromTargetEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Target.detachedFromTarget",
        Params: event,
    })
}
func (agent *TargetAgent) FireReceivedMessageFromTarget(event ReceivedMessageFromTargetEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Target.receivedMessageFromTarget",
        Params: event,
    })
}
func (agent *TargetAgent) FireReceivedMessageFromTargetOnTarget(targetId string,event ReceivedMessageFromTargetEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Target.receivedMessageFromTarget",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *TargetAgent) SetSetDiscoverTargetsHandler(handler func(SetDiscoverTargetsCommand)) {
    agent.commandHandlers.SetDiscoverTargets.Store(handler)
}
func (agent *TargetAgent) SetSetAutoAttachHandler(handler func(SetAutoAttachCommand)) {
    agent.commandHandlers.SetAutoAttach.Store(handler)
}
func (agent *TargetAgent) SetSetAttachToFramesHandler(handler func(SetAttachToFramesCommand)) {
    agent.commandHandlers.SetAttachToFrames.Store(handler)
}
func (agent *TargetAgent) SetSetRemoteLocationsHandler(handler func(SetRemoteLocationsCommand)) {
    agent.commandHandlers.SetRemoteLocations.Store(handler)
}
func (agent *TargetAgent) SetSendMessageToTargetHandler(handler func(SendMessageToTargetCommand)) {
    agent.commandHandlers.SendMessageToTarget.Store(handler)
}
func (agent *TargetAgent) SetGetTargetInfoHandler(handler func(GetTargetInfoCommand)) {
    agent.commandHandlers.GetTargetInfo.Store(handler)
}
func (agent *TargetAgent) SetActivateTargetHandler(handler func(ActivateTargetCommand)) {
    agent.commandHandlers.ActivateTarget.Store(handler)
}
func (agent *TargetAgent) SetCloseTargetHandler(handler func(CloseTargetCommand)) {
    agent.commandHandlers.CloseTarget.Store(handler)
}
func (agent *TargetAgent) SetAttachToTargetHandler(handler func(AttachToTargetCommand)) {
    agent.commandHandlers.AttachToTarget.Store(handler)
}
func (agent *TargetAgent) SetDetachFromTargetHandler(handler func(DetachFromTargetCommand)) {
    agent.commandHandlers.DetachFromTarget.Store(handler)
}
func (agent *TargetAgent) SetCreateBrowserContextHandler(handler func(CreateBrowserContextCommand)) {
    agent.commandHandlers.CreateBrowserContext.Store(handler)
}
func (agent *TargetAgent) SetDisposeBrowserContextHandler(handler func(DisposeBrowserContextCommand)) {
    agent.commandHandlers.DisposeBrowserContext.Store(handler)
}
func (agent *TargetAgent) SetCreateTargetHandler(handler func(CreateTargetCommand)) {
    agent.commandHandlers.CreateTarget.Store(handler)
}
func (agent *TargetAgent) SetGetTargetsHandler(handler func(GetTargetsCommand)) {
    agent.commandHandlers.GetTargets.Store(handler)
}
func init() {

}
