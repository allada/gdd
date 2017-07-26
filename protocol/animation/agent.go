package animation


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

type GetPlaybackRateCommandFn struct {
    mux sync.RWMutex
    fn func(GetPlaybackRateCommand)
}

func (a *GetPlaybackRateCommandFn) Load() func(GetPlaybackRateCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetPlaybackRateCommandFn) Store(fn func(GetPlaybackRateCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetPlaybackRateCommandFn struct {
    mux sync.RWMutex
    fn func(SetPlaybackRateCommand)
}

func (a *SetPlaybackRateCommandFn) Load() func(SetPlaybackRateCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetPlaybackRateCommandFn) Store(fn func(SetPlaybackRateCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetCurrentTimeCommandFn struct {
    mux sync.RWMutex
    fn func(GetCurrentTimeCommand)
}

func (a *GetCurrentTimeCommandFn) Load() func(GetCurrentTimeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetCurrentTimeCommandFn) Store(fn func(GetCurrentTimeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetPausedCommandFn struct {
    mux sync.RWMutex
    fn func(SetPausedCommand)
}

func (a *SetPausedCommandFn) Load() func(SetPausedCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetPausedCommandFn) Store(fn func(SetPausedCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetTimingCommandFn struct {
    mux sync.RWMutex
    fn func(SetTimingCommand)
}

func (a *SetTimingCommandFn) Load() func(SetTimingCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetTimingCommandFn) Store(fn func(SetTimingCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SeekAnimationsCommandFn struct {
    mux sync.RWMutex
    fn func(SeekAnimationsCommand)
}

func (a *SeekAnimationsCommandFn) Load() func(SeekAnimationsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SeekAnimationsCommandFn) Store(fn func(SeekAnimationsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ReleaseAnimationsCommandFn struct {
    mux sync.RWMutex
    fn func(ReleaseAnimationsCommand)
}

func (a *ReleaseAnimationsCommandFn) Load() func(ReleaseAnimationsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ReleaseAnimationsCommandFn) Store(fn func(ReleaseAnimationsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ResolveAnimationCommandFn struct {
    mux sync.RWMutex
    fn func(ResolveAnimationCommand)
}

func (a *ResolveAnimationCommandFn) Load() func(ResolveAnimationCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ResolveAnimationCommandFn) Store(fn func(ResolveAnimationCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type AnimationAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
        GetPlaybackRate GetPlaybackRateCommandFn
        SetPlaybackRate SetPlaybackRateCommandFn
        GetCurrentTime GetCurrentTimeCommandFn
        SetPaused SetPausedCommandFn
        SetTiming SetTimingCommandFn
        SeekAnimations SeekAnimationsCommandFn
        ReleaseAnimations ReleaseAnimationsCommandFn
        ResolveAnimation ResolveAnimationCommandFn
    }
}
func NewAgent(conn *shared.Connection) *AnimationAgent {
    agent := &AnimationAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *AnimationAgent) Name() string {
    return "Animation"
}

func (agent *AnimationAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "getPlaybackRate":
            var out GetPlaybackRateCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetPlaybackRate.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setPlaybackRate":
            var out SetPlaybackRateCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetPlaybackRate.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getCurrentTime":
            var out GetCurrentTimeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetCurrentTime.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setPaused":
            var out SetPausedCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetPaused.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setTiming":
            var out SetTimingCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetTiming.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "seekAnimations":
            var out SeekAnimationsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SeekAnimations.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "releaseAnimations":
            var out ReleaseAnimationsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ReleaseAnimations.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "resolveAnimation":
            var out ResolveAnimationCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ResolveAnimation.Load()
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
func (agent *AnimationAgent) FireAnimationCreated(event AnimationCreatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Animation.animationCreated",
        Params: event,
    })
}
func (agent *AnimationAgent) FireAnimationCreatedOnTarget(targetId string,event AnimationCreatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Animation.animationCreated",
        Params: event,
    })
}
func (agent *AnimationAgent) FireAnimationStarted(event AnimationStartedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Animation.animationStarted",
        Params: event,
    })
}
func (agent *AnimationAgent) FireAnimationStartedOnTarget(targetId string,event AnimationStartedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Animation.animationStarted",
        Params: event,
    })
}
func (agent *AnimationAgent) FireAnimationCanceled(event AnimationCanceledEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Animation.animationCanceled",
        Params: event,
    })
}
func (agent *AnimationAgent) FireAnimationCanceledOnTarget(targetId string,event AnimationCanceledEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Animation.animationCanceled",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *AnimationAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *AnimationAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *AnimationAgent) SetGetPlaybackRateHandler(handler func(GetPlaybackRateCommand)) {
    agent.commandHandlers.GetPlaybackRate.Store(handler)
}
func (agent *AnimationAgent) SetSetPlaybackRateHandler(handler func(SetPlaybackRateCommand)) {
    agent.commandHandlers.SetPlaybackRate.Store(handler)
}
func (agent *AnimationAgent) SetGetCurrentTimeHandler(handler func(GetCurrentTimeCommand)) {
    agent.commandHandlers.GetCurrentTime.Store(handler)
}
func (agent *AnimationAgent) SetSetPausedHandler(handler func(SetPausedCommand)) {
    agent.commandHandlers.SetPaused.Store(handler)
}
func (agent *AnimationAgent) SetSetTimingHandler(handler func(SetTimingCommand)) {
    agent.commandHandlers.SetTiming.Store(handler)
}
func (agent *AnimationAgent) SetSeekAnimationsHandler(handler func(SeekAnimationsCommand)) {
    agent.commandHandlers.SeekAnimations.Store(handler)
}
func (agent *AnimationAgent) SetReleaseAnimationsHandler(handler func(ReleaseAnimationsCommand)) {
    agent.commandHandlers.ReleaseAnimations.Store(handler)
}
func (agent *AnimationAgent) SetResolveAnimationHandler(handler func(ResolveAnimationCommand)) {
    agent.commandHandlers.ResolveAnimation.Store(handler)
}
func init() {

}
