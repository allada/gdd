package applicationcache


import (
    "../shared"
    "sync"
    "encoding/json"
    "fmt"
)
type GetFramesWithManifestsCommandFn struct {
    mux sync.RWMutex
    fn func(GetFramesWithManifestsCommand)
}

func (a *GetFramesWithManifestsCommandFn) Load() func(GetFramesWithManifestsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetFramesWithManifestsCommandFn) Store(fn func(GetFramesWithManifestsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

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

type GetManifestForFrameCommandFn struct {
    mux sync.RWMutex
    fn func(GetManifestForFrameCommand)
}

func (a *GetManifestForFrameCommandFn) Load() func(GetManifestForFrameCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetManifestForFrameCommandFn) Store(fn func(GetManifestForFrameCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetApplicationCacheForFrameCommandFn struct {
    mux sync.RWMutex
    fn func(GetApplicationCacheForFrameCommand)
}

func (a *GetApplicationCacheForFrameCommandFn) Load() func(GetApplicationCacheForFrameCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetApplicationCacheForFrameCommandFn) Store(fn func(GetApplicationCacheForFrameCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ApplicationCacheAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        GetFramesWithManifests GetFramesWithManifestsCommandFn
        Enable EnableCommandFn
        GetManifestForFrame GetManifestForFrameCommandFn
        GetApplicationCacheForFrame GetApplicationCacheForFrameCommandFn
    }
}
func NewAgent(conn *shared.Connection) *ApplicationCacheAgent {
    agent := &ApplicationCacheAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *ApplicationCacheAgent) Name() string {
    return "ApplicationCache"
}

func (agent *ApplicationCacheAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "getFramesWithManifests":
            var out GetFramesWithManifestsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetFramesWithManifests.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "enable":
            var out EnableCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Enable.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getManifestForFrame":
            var out GetManifestForFrameCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetManifestForFrame.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getApplicationCacheForFrame":
            var out GetApplicationCacheForFrameCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetApplicationCacheForFrame.Load()
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
func (agent *ApplicationCacheAgent) FireApplicationCacheStatusUpdated(event ApplicationCacheStatusUpdatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "ApplicationCache.applicationCacheStatusUpdated",
        Params: event,
    })
}
func (agent *ApplicationCacheAgent) FireApplicationCacheStatusUpdatedOnTarget(targetId string,event ApplicationCacheStatusUpdatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "ApplicationCache.applicationCacheStatusUpdated",
        Params: event,
    })
}
func (agent *ApplicationCacheAgent) FireNetworkStateUpdated(event NetworkStateUpdatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "ApplicationCache.networkStateUpdated",
        Params: event,
    })
}
func (agent *ApplicationCacheAgent) FireNetworkStateUpdatedOnTarget(targetId string,event NetworkStateUpdatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "ApplicationCache.networkStateUpdated",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *ApplicationCacheAgent) SetGetFramesWithManifestsHandler(handler func(GetFramesWithManifestsCommand)) {
    agent.commandHandlers.GetFramesWithManifests.Store(handler)
}
func (agent *ApplicationCacheAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *ApplicationCacheAgent) SetGetManifestForFrameHandler(handler func(GetManifestForFrameCommand)) {
    agent.commandHandlers.GetManifestForFrame.Store(handler)
}
func (agent *ApplicationCacheAgent) SetGetApplicationCacheForFrameHandler(handler func(GetApplicationCacheForFrameCommand)) {
    agent.commandHandlers.GetApplicationCacheForFrame.Store(handler)
}
func init() {

}
