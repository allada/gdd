package network


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

type SetUserAgentOverrideCommandFn struct {
    mux sync.RWMutex
    fn func(SetUserAgentOverrideCommand)
}

func (a *SetUserAgentOverrideCommandFn) Load() func(SetUserAgentOverrideCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetUserAgentOverrideCommandFn) Store(fn func(SetUserAgentOverrideCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetExtraHTTPHeadersCommandFn struct {
    mux sync.RWMutex
    fn func(SetExtraHTTPHeadersCommand)
}

func (a *SetExtraHTTPHeadersCommandFn) Load() func(SetExtraHTTPHeadersCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetExtraHTTPHeadersCommandFn) Store(fn func(SetExtraHTTPHeadersCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetResponseBodyCommandFn struct {
    mux sync.RWMutex
    fn func(GetResponseBodyCommand)
}

func (a *GetResponseBodyCommandFn) Load() func(GetResponseBodyCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetResponseBodyCommandFn) Store(fn func(GetResponseBodyCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type AddBlockedURLCommandFn struct {
    mux sync.RWMutex
    fn func(AddBlockedURLCommand)
}

func (a *AddBlockedURLCommandFn) Load() func(AddBlockedURLCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *AddBlockedURLCommandFn) Store(fn func(AddBlockedURLCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RemoveBlockedURLCommandFn struct {
    mux sync.RWMutex
    fn func(RemoveBlockedURLCommand)
}

func (a *RemoveBlockedURLCommandFn) Load() func(RemoveBlockedURLCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RemoveBlockedURLCommandFn) Store(fn func(RemoveBlockedURLCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ReplayXHRCommandFn struct {
    mux sync.RWMutex
    fn func(ReplayXHRCommand)
}

func (a *ReplayXHRCommandFn) Load() func(ReplayXHRCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ReplayXHRCommandFn) Store(fn func(ReplayXHRCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetMonitoringXHREnabledCommandFn struct {
    mux sync.RWMutex
    fn func(SetMonitoringXHREnabledCommand)
}

func (a *SetMonitoringXHREnabledCommandFn) Load() func(SetMonitoringXHREnabledCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetMonitoringXHREnabledCommandFn) Store(fn func(SetMonitoringXHREnabledCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CanClearBrowserCacheCommandFn struct {
    mux sync.RWMutex
    fn func(CanClearBrowserCacheCommand)
}

func (a *CanClearBrowserCacheCommandFn) Load() func(CanClearBrowserCacheCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CanClearBrowserCacheCommandFn) Store(fn func(CanClearBrowserCacheCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ClearBrowserCacheCommandFn struct {
    mux sync.RWMutex
    fn func(ClearBrowserCacheCommand)
}

func (a *ClearBrowserCacheCommandFn) Load() func(ClearBrowserCacheCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ClearBrowserCacheCommandFn) Store(fn func(ClearBrowserCacheCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CanClearBrowserCookiesCommandFn struct {
    mux sync.RWMutex
    fn func(CanClearBrowserCookiesCommand)
}

func (a *CanClearBrowserCookiesCommandFn) Load() func(CanClearBrowserCookiesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CanClearBrowserCookiesCommandFn) Store(fn func(CanClearBrowserCookiesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ClearBrowserCookiesCommandFn struct {
    mux sync.RWMutex
    fn func(ClearBrowserCookiesCommand)
}

func (a *ClearBrowserCookiesCommandFn) Load() func(ClearBrowserCookiesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ClearBrowserCookiesCommandFn) Store(fn func(ClearBrowserCookiesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetCookiesCommandFn struct {
    mux sync.RWMutex
    fn func(GetCookiesCommand)
}

func (a *GetCookiesCommandFn) Load() func(GetCookiesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetCookiesCommandFn) Store(fn func(GetCookiesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetAllCookiesCommandFn struct {
    mux sync.RWMutex
    fn func(GetAllCookiesCommand)
}

func (a *GetAllCookiesCommandFn) Load() func(GetAllCookiesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetAllCookiesCommandFn) Store(fn func(GetAllCookiesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DeleteCookieCommandFn struct {
    mux sync.RWMutex
    fn func(DeleteCookieCommand)
}

func (a *DeleteCookieCommandFn) Load() func(DeleteCookieCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *DeleteCookieCommandFn) Store(fn func(DeleteCookieCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetCookieCommandFn struct {
    mux sync.RWMutex
    fn func(SetCookieCommand)
}

func (a *SetCookieCommandFn) Load() func(SetCookieCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetCookieCommandFn) Store(fn func(SetCookieCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CanEmulateNetworkConditionsCommandFn struct {
    mux sync.RWMutex
    fn func(CanEmulateNetworkConditionsCommand)
}

func (a *CanEmulateNetworkConditionsCommandFn) Load() func(CanEmulateNetworkConditionsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CanEmulateNetworkConditionsCommandFn) Store(fn func(CanEmulateNetworkConditionsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type EmulateNetworkConditionsCommandFn struct {
    mux sync.RWMutex
    fn func(EmulateNetworkConditionsCommand)
}

func (a *EmulateNetworkConditionsCommandFn) Load() func(EmulateNetworkConditionsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *EmulateNetworkConditionsCommandFn) Store(fn func(EmulateNetworkConditionsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetCacheDisabledCommandFn struct {
    mux sync.RWMutex
    fn func(SetCacheDisabledCommand)
}

func (a *SetCacheDisabledCommandFn) Load() func(SetCacheDisabledCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetCacheDisabledCommandFn) Store(fn func(SetCacheDisabledCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetBypassServiceWorkerCommandFn struct {
    mux sync.RWMutex
    fn func(SetBypassServiceWorkerCommand)
}

func (a *SetBypassServiceWorkerCommandFn) Load() func(SetBypassServiceWorkerCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetBypassServiceWorkerCommandFn) Store(fn func(SetBypassServiceWorkerCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetDataSizeLimitsForTestCommandFn struct {
    mux sync.RWMutex
    fn func(SetDataSizeLimitsForTestCommand)
}

func (a *SetDataSizeLimitsForTestCommandFn) Load() func(SetDataSizeLimitsForTestCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetDataSizeLimitsForTestCommandFn) Store(fn func(SetDataSizeLimitsForTestCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetCertificateCommandFn struct {
    mux sync.RWMutex
    fn func(GetCertificateCommand)
}

func (a *GetCertificateCommandFn) Load() func(GetCertificateCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetCertificateCommandFn) Store(fn func(GetCertificateCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type NetworkAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
        SetUserAgentOverride SetUserAgentOverrideCommandFn
        SetExtraHTTPHeaders SetExtraHTTPHeadersCommandFn
        GetResponseBody GetResponseBodyCommandFn
        AddBlockedURL AddBlockedURLCommandFn
        RemoveBlockedURL RemoveBlockedURLCommandFn
        ReplayXHR ReplayXHRCommandFn
        SetMonitoringXHREnabled SetMonitoringXHREnabledCommandFn
        CanClearBrowserCache CanClearBrowserCacheCommandFn
        ClearBrowserCache ClearBrowserCacheCommandFn
        CanClearBrowserCookies CanClearBrowserCookiesCommandFn
        ClearBrowserCookies ClearBrowserCookiesCommandFn
        GetCookies GetCookiesCommandFn
        GetAllCookies GetAllCookiesCommandFn
        DeleteCookie DeleteCookieCommandFn
        SetCookie SetCookieCommandFn
        CanEmulateNetworkConditions CanEmulateNetworkConditionsCommandFn
        EmulateNetworkConditions EmulateNetworkConditionsCommandFn
        SetCacheDisabled SetCacheDisabledCommandFn
        SetBypassServiceWorker SetBypassServiceWorkerCommandFn
        SetDataSizeLimitsForTest SetDataSizeLimitsForTestCommandFn
        GetCertificate GetCertificateCommandFn
    }
}
func NewAgent(conn *shared.Connection) *NetworkAgent {
    agent := &NetworkAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *NetworkAgent) Name() string {
    return "Network"
}

func (agent *NetworkAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "setUserAgentOverride":
            var out SetUserAgentOverrideCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetUserAgentOverride.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setExtraHTTPHeaders":
            var out SetExtraHTTPHeadersCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetExtraHTTPHeaders.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getResponseBody":
            var out GetResponseBodyCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetResponseBody.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "addBlockedURL":
            var out AddBlockedURLCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.AddBlockedURL.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "removeBlockedURL":
            var out RemoveBlockedURLCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RemoveBlockedURL.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "replayXHR":
            var out ReplayXHRCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ReplayXHR.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setMonitoringXHREnabled":
            var out SetMonitoringXHREnabledCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetMonitoringXHREnabled.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "canClearBrowserCache":
            var out CanClearBrowserCacheCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CanClearBrowserCache.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "clearBrowserCache":
            var out ClearBrowserCacheCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ClearBrowserCache.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "canClearBrowserCookies":
            var out CanClearBrowserCookiesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CanClearBrowserCookies.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "clearBrowserCookies":
            var out ClearBrowserCookiesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ClearBrowserCookies.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getCookies":
            var out GetCookiesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetCookies.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getAllCookies":
            var out GetAllCookiesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetAllCookies.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "deleteCookie":
            var out DeleteCookieCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.DeleteCookie.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setCookie":
            var out SetCookieCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetCookie.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "canEmulateNetworkConditions":
            var out CanEmulateNetworkConditionsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CanEmulateNetworkConditions.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "emulateNetworkConditions":
            var out EmulateNetworkConditionsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.EmulateNetworkConditions.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setCacheDisabled":
            var out SetCacheDisabledCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetCacheDisabled.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setBypassServiceWorker":
            var out SetBypassServiceWorkerCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetBypassServiceWorker.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setDataSizeLimitsForTest":
            var out SetDataSizeLimitsForTestCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetDataSizeLimitsForTest.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getCertificate":
            var out GetCertificateCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetCertificate.Load()
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
func (agent *NetworkAgent) FireResourceChangedPriority(event ResourceChangedPriorityEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.resourceChangedPriority",
        Params: event,
    })
}
func (agent *NetworkAgent) FireResourceChangedPriorityOnTarget(targetId string,event ResourceChangedPriorityEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.resourceChangedPriority",
        Params: event,
    })
}
func (agent *NetworkAgent) FireRequestWillBeSent(event RequestWillBeSentEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.requestWillBeSent",
        Params: event,
    })
}
func (agent *NetworkAgent) FireRequestWillBeSentOnTarget(targetId string,event RequestWillBeSentEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.requestWillBeSent",
        Params: event,
    })
}
func (agent *NetworkAgent) FireRequestServedFromCache(event RequestServedFromCacheEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.requestServedFromCache",
        Params: event,
    })
}
func (agent *NetworkAgent) FireRequestServedFromCacheOnTarget(targetId string,event RequestServedFromCacheEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.requestServedFromCache",
        Params: event,
    })
}
func (agent *NetworkAgent) FireResponseReceived(event ResponseReceivedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.responseReceived",
        Params: event,
    })
}
func (agent *NetworkAgent) FireResponseReceivedOnTarget(targetId string,event ResponseReceivedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.responseReceived",
        Params: event,
    })
}
func (agent *NetworkAgent) FireDataReceived(event DataReceivedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.dataReceived",
        Params: event,
    })
}
func (agent *NetworkAgent) FireDataReceivedOnTarget(targetId string,event DataReceivedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.dataReceived",
        Params: event,
    })
}
func (agent *NetworkAgent) FireLoadingFinished(event LoadingFinishedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.loadingFinished",
        Params: event,
    })
}
func (agent *NetworkAgent) FireLoadingFinishedOnTarget(targetId string,event LoadingFinishedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.loadingFinished",
        Params: event,
    })
}
func (agent *NetworkAgent) FireLoadingFailed(event LoadingFailedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.loadingFailed",
        Params: event,
    })
}
func (agent *NetworkAgent) FireLoadingFailedOnTarget(targetId string,event LoadingFailedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.loadingFailed",
        Params: event,
    })
}
func (agent *NetworkAgent) FireWebSocketWillSendHandshakeRequest(event WebSocketWillSendHandshakeRequestEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.webSocketWillSendHandshakeRequest",
        Params: event,
    })
}
func (agent *NetworkAgent) FireWebSocketWillSendHandshakeRequestOnTarget(targetId string,event WebSocketWillSendHandshakeRequestEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.webSocketWillSendHandshakeRequest",
        Params: event,
    })
}
func (agent *NetworkAgent) FireWebSocketHandshakeResponseReceived(event WebSocketHandshakeResponseReceivedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.webSocketHandshakeResponseReceived",
        Params: event,
    })
}
func (agent *NetworkAgent) FireWebSocketHandshakeResponseReceivedOnTarget(targetId string,event WebSocketHandshakeResponseReceivedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.webSocketHandshakeResponseReceived",
        Params: event,
    })
}
func (agent *NetworkAgent) FireWebSocketCreated(event WebSocketCreatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.webSocketCreated",
        Params: event,
    })
}
func (agent *NetworkAgent) FireWebSocketCreatedOnTarget(targetId string,event WebSocketCreatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.webSocketCreated",
        Params: event,
    })
}
func (agent *NetworkAgent) FireWebSocketClosed(event WebSocketClosedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.webSocketClosed",
        Params: event,
    })
}
func (agent *NetworkAgent) FireWebSocketClosedOnTarget(targetId string,event WebSocketClosedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.webSocketClosed",
        Params: event,
    })
}
func (agent *NetworkAgent) FireWebSocketFrameReceived(event WebSocketFrameReceivedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.webSocketFrameReceived",
        Params: event,
    })
}
func (agent *NetworkAgent) FireWebSocketFrameReceivedOnTarget(targetId string,event WebSocketFrameReceivedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.webSocketFrameReceived",
        Params: event,
    })
}
func (agent *NetworkAgent) FireWebSocketFrameError(event WebSocketFrameErrorEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.webSocketFrameError",
        Params: event,
    })
}
func (agent *NetworkAgent) FireWebSocketFrameErrorOnTarget(targetId string,event WebSocketFrameErrorEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.webSocketFrameError",
        Params: event,
    })
}
func (agent *NetworkAgent) FireWebSocketFrameSent(event WebSocketFrameSentEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.webSocketFrameSent",
        Params: event,
    })
}
func (agent *NetworkAgent) FireWebSocketFrameSentOnTarget(targetId string,event WebSocketFrameSentEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.webSocketFrameSent",
        Params: event,
    })
}
func (agent *NetworkAgent) FireEventSourceMessageReceived(event EventSourceMessageReceivedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.eventSourceMessageReceived",
        Params: event,
    })
}
func (agent *NetworkAgent) FireEventSourceMessageReceivedOnTarget(targetId string,event EventSourceMessageReceivedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.eventSourceMessageReceived",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *NetworkAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *NetworkAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *NetworkAgent) SetSetUserAgentOverrideHandler(handler func(SetUserAgentOverrideCommand)) {
    agent.commandHandlers.SetUserAgentOverride.Store(handler)
}
func (agent *NetworkAgent) SetSetExtraHTTPHeadersHandler(handler func(SetExtraHTTPHeadersCommand)) {
    agent.commandHandlers.SetExtraHTTPHeaders.Store(handler)
}
func (agent *NetworkAgent) SetGetResponseBodyHandler(handler func(GetResponseBodyCommand)) {
    agent.commandHandlers.GetResponseBody.Store(handler)
}
func (agent *NetworkAgent) SetAddBlockedURLHandler(handler func(AddBlockedURLCommand)) {
    agent.commandHandlers.AddBlockedURL.Store(handler)
}
func (agent *NetworkAgent) SetRemoveBlockedURLHandler(handler func(RemoveBlockedURLCommand)) {
    agent.commandHandlers.RemoveBlockedURL.Store(handler)
}
func (agent *NetworkAgent) SetReplayXHRHandler(handler func(ReplayXHRCommand)) {
    agent.commandHandlers.ReplayXHR.Store(handler)
}
func (agent *NetworkAgent) SetSetMonitoringXHREnabledHandler(handler func(SetMonitoringXHREnabledCommand)) {
    agent.commandHandlers.SetMonitoringXHREnabled.Store(handler)
}
func (agent *NetworkAgent) SetCanClearBrowserCacheHandler(handler func(CanClearBrowserCacheCommand)) {
    agent.commandHandlers.CanClearBrowserCache.Store(handler)
}
func (agent *NetworkAgent) SetClearBrowserCacheHandler(handler func(ClearBrowserCacheCommand)) {
    agent.commandHandlers.ClearBrowserCache.Store(handler)
}
func (agent *NetworkAgent) SetCanClearBrowserCookiesHandler(handler func(CanClearBrowserCookiesCommand)) {
    agent.commandHandlers.CanClearBrowserCookies.Store(handler)
}
func (agent *NetworkAgent) SetClearBrowserCookiesHandler(handler func(ClearBrowserCookiesCommand)) {
    agent.commandHandlers.ClearBrowserCookies.Store(handler)
}
func (agent *NetworkAgent) SetGetCookiesHandler(handler func(GetCookiesCommand)) {
    agent.commandHandlers.GetCookies.Store(handler)
}
func (agent *NetworkAgent) SetGetAllCookiesHandler(handler func(GetAllCookiesCommand)) {
    agent.commandHandlers.GetAllCookies.Store(handler)
}
func (agent *NetworkAgent) SetDeleteCookieHandler(handler func(DeleteCookieCommand)) {
    agent.commandHandlers.DeleteCookie.Store(handler)
}
func (agent *NetworkAgent) SetSetCookieHandler(handler func(SetCookieCommand)) {
    agent.commandHandlers.SetCookie.Store(handler)
}
func (agent *NetworkAgent) SetCanEmulateNetworkConditionsHandler(handler func(CanEmulateNetworkConditionsCommand)) {
    agent.commandHandlers.CanEmulateNetworkConditions.Store(handler)
}
func (agent *NetworkAgent) SetEmulateNetworkConditionsHandler(handler func(EmulateNetworkConditionsCommand)) {
    agent.commandHandlers.EmulateNetworkConditions.Store(handler)
}
func (agent *NetworkAgent) SetSetCacheDisabledHandler(handler func(SetCacheDisabledCommand)) {
    agent.commandHandlers.SetCacheDisabled.Store(handler)
}
func (agent *NetworkAgent) SetSetBypassServiceWorkerHandler(handler func(SetBypassServiceWorkerCommand)) {
    agent.commandHandlers.SetBypassServiceWorker.Store(handler)
}
func (agent *NetworkAgent) SetSetDataSizeLimitsForTestHandler(handler func(SetDataSizeLimitsForTestCommand)) {
    agent.commandHandlers.SetDataSizeLimitsForTest.Store(handler)
}
func (agent *NetworkAgent) SetGetCertificateHandler(handler func(GetCertificateCommand)) {
    agent.commandHandlers.GetCertificate.Store(handler)
}
func init() {

}
