package layertree


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

type CompositingReasonsCommandFn struct {
    mux sync.RWMutex
    fn func(CompositingReasonsCommand)
}

func (a *CompositingReasonsCommandFn) Load() func(CompositingReasonsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CompositingReasonsCommandFn) Store(fn func(CompositingReasonsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type MakeSnapshotCommandFn struct {
    mux sync.RWMutex
    fn func(MakeSnapshotCommand)
}

func (a *MakeSnapshotCommandFn) Load() func(MakeSnapshotCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *MakeSnapshotCommandFn) Store(fn func(MakeSnapshotCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type LoadSnapshotCommandFn struct {
    mux sync.RWMutex
    fn func(LoadSnapshotCommand)
}

func (a *LoadSnapshotCommandFn) Load() func(LoadSnapshotCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *LoadSnapshotCommandFn) Store(fn func(LoadSnapshotCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ReleaseSnapshotCommandFn struct {
    mux sync.RWMutex
    fn func(ReleaseSnapshotCommand)
}

func (a *ReleaseSnapshotCommandFn) Load() func(ReleaseSnapshotCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ReleaseSnapshotCommandFn) Store(fn func(ReleaseSnapshotCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ProfileSnapshotCommandFn struct {
    mux sync.RWMutex
    fn func(ProfileSnapshotCommand)
}

func (a *ProfileSnapshotCommandFn) Load() func(ProfileSnapshotCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ProfileSnapshotCommandFn) Store(fn func(ProfileSnapshotCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ReplaySnapshotCommandFn struct {
    mux sync.RWMutex
    fn func(ReplaySnapshotCommand)
}

func (a *ReplaySnapshotCommandFn) Load() func(ReplaySnapshotCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ReplaySnapshotCommandFn) Store(fn func(ReplaySnapshotCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SnapshotCommandLogCommandFn struct {
    mux sync.RWMutex
    fn func(SnapshotCommandLogCommand)
}

func (a *SnapshotCommandLogCommandFn) Load() func(SnapshotCommandLogCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SnapshotCommandLogCommandFn) Store(fn func(SnapshotCommandLogCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type LayerTreeAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
        CompositingReasons CompositingReasonsCommandFn
        MakeSnapshot MakeSnapshotCommandFn
        LoadSnapshot LoadSnapshotCommandFn
        ReleaseSnapshot ReleaseSnapshotCommandFn
        ProfileSnapshot ProfileSnapshotCommandFn
        ReplaySnapshot ReplaySnapshotCommandFn
        SnapshotCommandLog SnapshotCommandLogCommandFn
    }
}
func NewAgent(conn *shared.Connection) *LayerTreeAgent {
    agent := &LayerTreeAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *LayerTreeAgent) Name() string {
    return "LayerTree"
}

func (agent *LayerTreeAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "compositingReasons":
            var out CompositingReasonsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CompositingReasons.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "makeSnapshot":
            var out MakeSnapshotCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.MakeSnapshot.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "loadSnapshot":
            var out LoadSnapshotCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.LoadSnapshot.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "releaseSnapshot":
            var out ReleaseSnapshotCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ReleaseSnapshot.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "profileSnapshot":
            var out ProfileSnapshotCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ProfileSnapshot.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "replaySnapshot":
            var out ReplaySnapshotCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ReplaySnapshot.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "snapshotCommandLog":
            var out SnapshotCommandLogCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SnapshotCommandLog.Load()
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
func (agent *LayerTreeAgent) FireLayerTreeDidChange(event LayerTreeDidChangeEvent) {
    agent.conn.Send(shared.Notification{
        Method: "LayerTree.layerTreeDidChange",
        Params: event,
    })
}
func (agent *LayerTreeAgent) FireLayerTreeDidChangeOnTarget(targetId string,event LayerTreeDidChangeEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "LayerTree.layerTreeDidChange",
        Params: event,
    })
}
func (agent *LayerTreeAgent) FireLayerPainted(event LayerPaintedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "LayerTree.layerPainted",
        Params: event,
    })
}
func (agent *LayerTreeAgent) FireLayerPaintedOnTarget(targetId string,event LayerPaintedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "LayerTree.layerPainted",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *LayerTreeAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *LayerTreeAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *LayerTreeAgent) SetCompositingReasonsHandler(handler func(CompositingReasonsCommand)) {
    agent.commandHandlers.CompositingReasons.Store(handler)
}
func (agent *LayerTreeAgent) SetMakeSnapshotHandler(handler func(MakeSnapshotCommand)) {
    agent.commandHandlers.MakeSnapshot.Store(handler)
}
func (agent *LayerTreeAgent) SetLoadSnapshotHandler(handler func(LoadSnapshotCommand)) {
    agent.commandHandlers.LoadSnapshot.Store(handler)
}
func (agent *LayerTreeAgent) SetReleaseSnapshotHandler(handler func(ReleaseSnapshotCommand)) {
    agent.commandHandlers.ReleaseSnapshot.Store(handler)
}
func (agent *LayerTreeAgent) SetProfileSnapshotHandler(handler func(ProfileSnapshotCommand)) {
    agent.commandHandlers.ProfileSnapshot.Store(handler)
}
func (agent *LayerTreeAgent) SetReplaySnapshotHandler(handler func(ReplaySnapshotCommand)) {
    agent.commandHandlers.ReplaySnapshot.Store(handler)
}
func (agent *LayerTreeAgent) SetSnapshotCommandLogHandler(handler func(SnapshotCommandLogCommand)) {
    agent.commandHandlers.SnapshotCommandLog.Store(handler)
}
func init() {

}
