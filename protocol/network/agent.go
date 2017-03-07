package network


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type NetworkAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        SetUserAgentOverride []chan<- SetUserAgentOverrideCommand
        SetExtraHTTPHeaders []chan<- SetExtraHTTPHeadersCommand
        GetResponseBody []chan<- GetResponseBodyCommand
        AddBlockedURL []chan<- AddBlockedURLCommand
        RemoveBlockedURL []chan<- RemoveBlockedURLCommand
        ReplayXHR []chan<- ReplayXHRCommand
        SetMonitoringXHREnabled []chan<- SetMonitoringXHREnabledCommand
        CanClearBrowserCache []chan<- CanClearBrowserCacheCommand
        ClearBrowserCache []chan<- ClearBrowserCacheCommand
        CanClearBrowserCookies []chan<- CanClearBrowserCookiesCommand
        ClearBrowserCookies []chan<- ClearBrowserCookiesCommand
        GetCookies []chan<- GetCookiesCommand
        GetAllCookies []chan<- GetAllCookiesCommand
        DeleteCookie []chan<- DeleteCookieCommand
        SetCookie []chan<- SetCookieCommand
        CanEmulateNetworkConditions []chan<- CanEmulateNetworkConditionsCommand
        EmulateNetworkConditions []chan<- EmulateNetworkConditionsCommand
        SetCacheDisabled []chan<- SetCacheDisabledCommand
        SetBypassServiceWorker []chan<- SetBypassServiceWorkerCommand
        SetDataSizeLimitsForTest []chan<- SetDataSizeLimitsForTestCommand
        GetCertificate []chan<- GetCertificateCommand
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
    switch (funcName) {
        case "enable":
            var out EnableCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Enable {
                c <-out
            }
        case "disable":
            var out DisableCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Disable {
                c <-out
            }
        case "setUserAgentOverride":
            var out SetUserAgentOverrideCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetUserAgentOverride {
                c <-out
            }
        case "setExtraHTTPHeaders":
            var out SetExtraHTTPHeadersCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetExtraHTTPHeaders {
                c <-out
            }
        case "getResponseBody":
            var out GetResponseBodyCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetResponseBody {
                c <-out
            }
        case "addBlockedURL":
            var out AddBlockedURLCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.AddBlockedURL {
                c <-out
            }
        case "removeBlockedURL":
            var out RemoveBlockedURLCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RemoveBlockedURL {
                c <-out
            }
        case "replayXHR":
            var out ReplayXHRCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ReplayXHR {
                c <-out
            }
        case "setMonitoringXHREnabled":
            var out SetMonitoringXHREnabledCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetMonitoringXHREnabled {
                c <-out
            }
        case "canClearBrowserCache":
            var out CanClearBrowserCacheCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CanClearBrowserCache {
                c <-out
            }
        case "clearBrowserCache":
            var out ClearBrowserCacheCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ClearBrowserCache {
                c <-out
            }
        case "canClearBrowserCookies":
            var out CanClearBrowserCookiesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CanClearBrowserCookies {
                c <-out
            }
        case "clearBrowserCookies":
            var out ClearBrowserCookiesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ClearBrowserCookies {
                c <-out
            }
        case "getCookies":
            var out GetCookiesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetCookies {
                c <-out
            }
        case "getAllCookies":
            var out GetAllCookiesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetAllCookies {
                c <-out
            }
        case "deleteCookie":
            var out DeleteCookieCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.DeleteCookie {
                c <-out
            }
        case "setCookie":
            var out SetCookieCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetCookie {
                c <-out
            }
        case "canEmulateNetworkConditions":
            var out CanEmulateNetworkConditionsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CanEmulateNetworkConditions {
                c <-out
            }
        case "emulateNetworkConditions":
            var out EmulateNetworkConditionsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.EmulateNetworkConditions {
                c <-out
            }
        case "setCacheDisabled":
            var out SetCacheDisabledCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetCacheDisabled {
                c <-out
            }
        case "setBypassServiceWorker":
            var out SetBypassServiceWorkerCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetBypassServiceWorker {
                c <-out
            }
        case "setDataSizeLimitsForTest":
            var out SetDataSizeLimitsForTestCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetDataSizeLimitsForTest {
                c <-out
            }
        case "getCertificate":
            var out GetCertificateCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetCertificate {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *NetworkAgent) FireResourceChangedPriority(event ResourceChangedPriorityEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Network.resourceChangedPriority",
        Params: event,
    })
}
func (agent *NetworkAgent) FireResourceChangedPriorityOnTarget(targetId string, event ResourceChangedPriorityEvent) {
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
func (agent *NetworkAgent) FireRequestWillBeSentOnTarget(targetId string, event RequestWillBeSentEvent) {
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
func (agent *NetworkAgent) FireRequestServedFromCacheOnTarget(targetId string, event RequestServedFromCacheEvent) {
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
func (agent *NetworkAgent) FireResponseReceivedOnTarget(targetId string, event ResponseReceivedEvent) {
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
func (agent *NetworkAgent) FireDataReceivedOnTarget(targetId string, event DataReceivedEvent) {
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
func (agent *NetworkAgent) FireLoadingFinishedOnTarget(targetId string, event LoadingFinishedEvent) {
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
func (agent *NetworkAgent) FireLoadingFailedOnTarget(targetId string, event LoadingFailedEvent) {
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
func (agent *NetworkAgent) FireWebSocketWillSendHandshakeRequestOnTarget(targetId string, event WebSocketWillSendHandshakeRequestEvent) {
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
func (agent *NetworkAgent) FireWebSocketHandshakeResponseReceivedOnTarget(targetId string, event WebSocketHandshakeResponseReceivedEvent) {
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
func (agent *NetworkAgent) FireWebSocketCreatedOnTarget(targetId string, event WebSocketCreatedEvent) {
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
func (agent *NetworkAgent) FireWebSocketClosedOnTarget(targetId string, event WebSocketClosedEvent) {
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
func (agent *NetworkAgent) FireWebSocketFrameReceivedOnTarget(targetId string, event WebSocketFrameReceivedEvent) {
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
func (agent *NetworkAgent) FireWebSocketFrameErrorOnTarget(targetId string, event WebSocketFrameErrorEvent) {
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
func (agent *NetworkAgent) FireWebSocketFrameSentOnTarget(targetId string, event WebSocketFrameSentEvent) {
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
func (agent *NetworkAgent) FireEventSourceMessageReceivedOnTarget(targetId string, event EventSourceMessageReceivedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Network.eventSourceMessageReceived",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *NetworkAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *NetworkAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *NetworkAgent) SetUserAgentOverrideNotify() <-chan SetUserAgentOverrideCommand {
    outChan := make(chan SetUserAgentOverrideCommand)
    agent.commandChans.SetUserAgentOverride = append(agent.commandChans.SetUserAgentOverride, outChan)
    return outChan
}
func (agent *NetworkAgent) SetExtraHTTPHeadersNotify() <-chan SetExtraHTTPHeadersCommand {
    outChan := make(chan SetExtraHTTPHeadersCommand)
    agent.commandChans.SetExtraHTTPHeaders = append(agent.commandChans.SetExtraHTTPHeaders, outChan)
    return outChan
}
func (agent *NetworkAgent) GetResponseBodyNotify() <-chan GetResponseBodyCommand {
    outChan := make(chan GetResponseBodyCommand)
    agent.commandChans.GetResponseBody = append(agent.commandChans.GetResponseBody, outChan)
    return outChan
}
func (agent *NetworkAgent) AddBlockedURLNotify() <-chan AddBlockedURLCommand {
    outChan := make(chan AddBlockedURLCommand)
    agent.commandChans.AddBlockedURL = append(agent.commandChans.AddBlockedURL, outChan)
    return outChan
}
func (agent *NetworkAgent) RemoveBlockedURLNotify() <-chan RemoveBlockedURLCommand {
    outChan := make(chan RemoveBlockedURLCommand)
    agent.commandChans.RemoveBlockedURL = append(agent.commandChans.RemoveBlockedURL, outChan)
    return outChan
}
func (agent *NetworkAgent) ReplayXHRNotify() <-chan ReplayXHRCommand {
    outChan := make(chan ReplayXHRCommand)
    agent.commandChans.ReplayXHR = append(agent.commandChans.ReplayXHR, outChan)
    return outChan
}
func (agent *NetworkAgent) SetMonitoringXHREnabledNotify() <-chan SetMonitoringXHREnabledCommand {
    outChan := make(chan SetMonitoringXHREnabledCommand)
    agent.commandChans.SetMonitoringXHREnabled = append(agent.commandChans.SetMonitoringXHREnabled, outChan)
    return outChan
}
func (agent *NetworkAgent) CanClearBrowserCacheNotify() <-chan CanClearBrowserCacheCommand {
    outChan := make(chan CanClearBrowserCacheCommand)
    agent.commandChans.CanClearBrowserCache = append(agent.commandChans.CanClearBrowserCache, outChan)
    return outChan
}
func (agent *NetworkAgent) ClearBrowserCacheNotify() <-chan ClearBrowserCacheCommand {
    outChan := make(chan ClearBrowserCacheCommand)
    agent.commandChans.ClearBrowserCache = append(agent.commandChans.ClearBrowserCache, outChan)
    return outChan
}
func (agent *NetworkAgent) CanClearBrowserCookiesNotify() <-chan CanClearBrowserCookiesCommand {
    outChan := make(chan CanClearBrowserCookiesCommand)
    agent.commandChans.CanClearBrowserCookies = append(agent.commandChans.CanClearBrowserCookies, outChan)
    return outChan
}
func (agent *NetworkAgent) ClearBrowserCookiesNotify() <-chan ClearBrowserCookiesCommand {
    outChan := make(chan ClearBrowserCookiesCommand)
    agent.commandChans.ClearBrowserCookies = append(agent.commandChans.ClearBrowserCookies, outChan)
    return outChan
}
func (agent *NetworkAgent) GetCookiesNotify() <-chan GetCookiesCommand {
    outChan := make(chan GetCookiesCommand)
    agent.commandChans.GetCookies = append(agent.commandChans.GetCookies, outChan)
    return outChan
}
func (agent *NetworkAgent) GetAllCookiesNotify() <-chan GetAllCookiesCommand {
    outChan := make(chan GetAllCookiesCommand)
    agent.commandChans.GetAllCookies = append(agent.commandChans.GetAllCookies, outChan)
    return outChan
}
func (agent *NetworkAgent) DeleteCookieNotify() <-chan DeleteCookieCommand {
    outChan := make(chan DeleteCookieCommand)
    agent.commandChans.DeleteCookie = append(agent.commandChans.DeleteCookie, outChan)
    return outChan
}
func (agent *NetworkAgent) SetCookieNotify() <-chan SetCookieCommand {
    outChan := make(chan SetCookieCommand)
    agent.commandChans.SetCookie = append(agent.commandChans.SetCookie, outChan)
    return outChan
}
func (agent *NetworkAgent) CanEmulateNetworkConditionsNotify() <-chan CanEmulateNetworkConditionsCommand {
    outChan := make(chan CanEmulateNetworkConditionsCommand)
    agent.commandChans.CanEmulateNetworkConditions = append(agent.commandChans.CanEmulateNetworkConditions, outChan)
    return outChan
}
func (agent *NetworkAgent) EmulateNetworkConditionsNotify() <-chan EmulateNetworkConditionsCommand {
    outChan := make(chan EmulateNetworkConditionsCommand)
    agent.commandChans.EmulateNetworkConditions = append(agent.commandChans.EmulateNetworkConditions, outChan)
    return outChan
}
func (agent *NetworkAgent) SetCacheDisabledNotify() <-chan SetCacheDisabledCommand {
    outChan := make(chan SetCacheDisabledCommand)
    agent.commandChans.SetCacheDisabled = append(agent.commandChans.SetCacheDisabled, outChan)
    return outChan
}
func (agent *NetworkAgent) SetBypassServiceWorkerNotify() <-chan SetBypassServiceWorkerCommand {
    outChan := make(chan SetBypassServiceWorkerCommand)
    agent.commandChans.SetBypassServiceWorker = append(agent.commandChans.SetBypassServiceWorker, outChan)
    return outChan
}
func (agent *NetworkAgent) SetDataSizeLimitsForTestNotify() <-chan SetDataSizeLimitsForTestCommand {
    outChan := make(chan SetDataSizeLimitsForTestCommand)
    agent.commandChans.SetDataSizeLimitsForTest = append(agent.commandChans.SetDataSizeLimitsForTest, outChan)
    return outChan
}
func (agent *NetworkAgent) GetCertificateNotify() <-chan GetCertificateCommand {
    outChan := make(chan GetCertificateCommand)
    agent.commandChans.GetCertificate = append(agent.commandChans.GetCertificate, outChan)
    return outChan
}
func init() {

}
