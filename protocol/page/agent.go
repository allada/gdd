package page


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type PageAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        AddScriptToEvaluateOnLoad []chan<- AddScriptToEvaluateOnLoadCommand
        RemoveScriptToEvaluateOnLoad []chan<- RemoveScriptToEvaluateOnLoadCommand
        SetAutoAttachToCreatedPages []chan<- SetAutoAttachToCreatedPagesCommand
        Reload []chan<- ReloadCommand
        Navigate []chan<- NavigateCommand
        StopLoading []chan<- StopLoadingCommand
        GetNavigationHistory []chan<- GetNavigationHistoryCommand
        NavigateToHistoryEntry []chan<- NavigateToHistoryEntryCommand
        GetCookies []chan<- GetCookiesCommand
        DeleteCookie []chan<- DeleteCookieCommand
        GetResourceTree []chan<- GetResourceTreeCommand
        GetResourceContent []chan<- GetResourceContentCommand
        SearchInResource []chan<- SearchInResourceCommand
        SetDocumentContent []chan<- SetDocumentContentCommand
        SetDeviceMetricsOverride []chan<- SetDeviceMetricsOverrideCommand
        ClearDeviceMetricsOverride []chan<- ClearDeviceMetricsOverrideCommand
        SetGeolocationOverride []chan<- SetGeolocationOverrideCommand
        ClearGeolocationOverride []chan<- ClearGeolocationOverrideCommand
        SetDeviceOrientationOverride []chan<- SetDeviceOrientationOverrideCommand
        ClearDeviceOrientationOverride []chan<- ClearDeviceOrientationOverrideCommand
        SetTouchEmulationEnabled []chan<- SetTouchEmulationEnabledCommand
        CaptureScreenshot []chan<- CaptureScreenshotCommand
        PrintToPDF []chan<- PrintToPDFCommand
        StartScreencast []chan<- StartScreencastCommand
        StopScreencast []chan<- StopScreencastCommand
        ScreencastFrameAck []chan<- ScreencastFrameAckCommand
        HandleJavaScriptDialog []chan<- HandleJavaScriptDialogCommand
        SetColorPickerEnabled []chan<- SetColorPickerEnabledCommand
        ConfigureOverlay []chan<- ConfigureOverlayCommand
        GetAppManifest []chan<- GetAppManifestCommand
        RequestAppBanner []chan<- RequestAppBannerCommand
        SetControlNavigations []chan<- SetControlNavigationsCommand
        ProcessNavigation []chan<- ProcessNavigationCommand
        GetLayoutMetrics []chan<- GetLayoutMetricsCommand
    }
}
func NewAgent(conn *shared.Connection) *PageAgent {
    agent := &PageAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *PageAgent) Name() string {
    return "Page"
}

func (agent *PageAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "addScriptToEvaluateOnLoad":
            var out AddScriptToEvaluateOnLoadCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.AddScriptToEvaluateOnLoad {
                c <-out
            }
        case "removeScriptToEvaluateOnLoad":
            var out RemoveScriptToEvaluateOnLoadCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RemoveScriptToEvaluateOnLoad {
                c <-out
            }
        case "setAutoAttachToCreatedPages":
            var out SetAutoAttachToCreatedPagesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetAutoAttachToCreatedPages {
                c <-out
            }
        case "reload":
            var out ReloadCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Reload {
                c <-out
            }
        case "navigate":
            var out NavigateCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Navigate {
                c <-out
            }
        case "stopLoading":
            var out StopLoadingCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StopLoading {
                c <-out
            }
        case "getNavigationHistory":
            var out GetNavigationHistoryCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetNavigationHistory {
                c <-out
            }
        case "navigateToHistoryEntry":
            var out NavigateToHistoryEntryCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.NavigateToHistoryEntry {
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
        case "getResourceTree":
            var out GetResourceTreeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetResourceTree {
                c <-out
            }
        case "getResourceContent":
            var out GetResourceContentCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetResourceContent {
                c <-out
            }
        case "searchInResource":
            var out SearchInResourceCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SearchInResource {
                c <-out
            }
        case "setDocumentContent":
            var out SetDocumentContentCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetDocumentContent {
                c <-out
            }
        case "setDeviceMetricsOverride":
            var out SetDeviceMetricsOverrideCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetDeviceMetricsOverride {
                c <-out
            }
        case "clearDeviceMetricsOverride":
            var out ClearDeviceMetricsOverrideCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ClearDeviceMetricsOverride {
                c <-out
            }
        case "setGeolocationOverride":
            var out SetGeolocationOverrideCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetGeolocationOverride {
                c <-out
            }
        case "clearGeolocationOverride":
            var out ClearGeolocationOverrideCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ClearGeolocationOverride {
                c <-out
            }
        case "setDeviceOrientationOverride":
            var out SetDeviceOrientationOverrideCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetDeviceOrientationOverride {
                c <-out
            }
        case "clearDeviceOrientationOverride":
            var out ClearDeviceOrientationOverrideCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ClearDeviceOrientationOverride {
                c <-out
            }
        case "setTouchEmulationEnabled":
            var out SetTouchEmulationEnabledCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetTouchEmulationEnabled {
                c <-out
            }
        case "captureScreenshot":
            var out CaptureScreenshotCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CaptureScreenshot {
                c <-out
            }
        case "printToPDF":
            var out PrintToPDFCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.PrintToPDF {
                c <-out
            }
        case "startScreencast":
            var out StartScreencastCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StartScreencast {
                c <-out
            }
        case "stopScreencast":
            var out StopScreencastCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StopScreencast {
                c <-out
            }
        case "screencastFrameAck":
            var out ScreencastFrameAckCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ScreencastFrameAck {
                c <-out
            }
        case "handleJavaScriptDialog":
            var out HandleJavaScriptDialogCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.HandleJavaScriptDialog {
                c <-out
            }
        case "setColorPickerEnabled":
            var out SetColorPickerEnabledCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetColorPickerEnabled {
                c <-out
            }
        case "configureOverlay":
            var out ConfigureOverlayCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ConfigureOverlay {
                c <-out
            }
        case "getAppManifest":
            var out GetAppManifestCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetAppManifest {
                c <-out
            }
        case "requestAppBanner":
            var out RequestAppBannerCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RequestAppBanner {
                c <-out
            }
        case "setControlNavigations":
            var out SetControlNavigationsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetControlNavigations {
                c <-out
            }
        case "processNavigation":
            var out ProcessNavigationCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ProcessNavigation {
                c <-out
            }
        case "getLayoutMetrics":
            var out GetLayoutMetricsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetLayoutMetrics {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *PageAgent) FireDomContentEventFired(event DomContentEventFiredEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.domContentEventFired",
        Params: event,
    })
}
func (agent *PageAgent) FireDomContentEventFiredOnTarget(targetId string, event DomContentEventFiredEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.domContentEventFired",
        Params: event,
    })
}
func (agent *PageAgent) FireLoadEventFired(event LoadEventFiredEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.loadEventFired",
        Params: event,
    })
}
func (agent *PageAgent) FireLoadEventFiredOnTarget(targetId string, event LoadEventFiredEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.loadEventFired",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameAttached(event FrameAttachedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.frameAttached",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameAttachedOnTarget(targetId string, event FrameAttachedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.frameAttached",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameNavigated(event FrameNavigatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.frameNavigated",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameNavigatedOnTarget(targetId string, event FrameNavigatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.frameNavigated",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameDetached(event FrameDetachedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.frameDetached",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameDetachedOnTarget(targetId string, event FrameDetachedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.frameDetached",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameStartedLoading(event FrameStartedLoadingEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.frameStartedLoading",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameStartedLoadingOnTarget(targetId string, event FrameStartedLoadingEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.frameStartedLoading",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameStoppedLoading(event FrameStoppedLoadingEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.frameStoppedLoading",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameStoppedLoadingOnTarget(targetId string, event FrameStoppedLoadingEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.frameStoppedLoading",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameScheduledNavigation(event FrameScheduledNavigationEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.frameScheduledNavigation",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameScheduledNavigationOnTarget(targetId string, event FrameScheduledNavigationEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.frameScheduledNavigation",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameClearedScheduledNavigation(event FrameClearedScheduledNavigationEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.frameClearedScheduledNavigation",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameClearedScheduledNavigationOnTarget(targetId string, event FrameClearedScheduledNavigationEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.frameClearedScheduledNavigation",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameResized(event FrameResizedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.frameResized",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameResizedOnTarget(targetId string, event FrameResizedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.frameResized",
        Params: event,
    })
}
func (agent *PageAgent) FireJavascriptDialogOpening(event JavascriptDialogOpeningEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.javascriptDialogOpening",
        Params: event,
    })
}
func (agent *PageAgent) FireJavascriptDialogOpeningOnTarget(targetId string, event JavascriptDialogOpeningEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.javascriptDialogOpening",
        Params: event,
    })
}
func (agent *PageAgent) FireJavascriptDialogClosed(event JavascriptDialogClosedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.javascriptDialogClosed",
        Params: event,
    })
}
func (agent *PageAgent) FireJavascriptDialogClosedOnTarget(targetId string, event JavascriptDialogClosedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.javascriptDialogClosed",
        Params: event,
    })
}
func (agent *PageAgent) FireScreencastFrame(event ScreencastFrameEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.screencastFrame",
        Params: event,
    })
}
func (agent *PageAgent) FireScreencastFrameOnTarget(targetId string, event ScreencastFrameEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.screencastFrame",
        Params: event,
    })
}
func (agent *PageAgent) FireScreencastVisibilityChanged(event ScreencastVisibilityChangedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.screencastVisibilityChanged",
        Params: event,
    })
}
func (agent *PageAgent) FireScreencastVisibilityChangedOnTarget(targetId string, event ScreencastVisibilityChangedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.screencastVisibilityChanged",
        Params: event,
    })
}
func (agent *PageAgent) FireColorPicked(event ColorPickedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.colorPicked",
        Params: event,
    })
}
func (agent *PageAgent) FireColorPickedOnTarget(targetId string, event ColorPickedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.colorPicked",
        Params: event,
    })
}
func (agent *PageAgent) FireInterstitialShown(event InterstitialShownEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.interstitialShown",
        Params: event,
    })
}
func (agent *PageAgent) FireInterstitialShownOnTarget(targetId string, event InterstitialShownEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.interstitialShown",
        Params: event,
    })
}
func (agent *PageAgent) FireInterstitialHidden(event InterstitialHiddenEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.interstitialHidden",
        Params: event,
    })
}
func (agent *PageAgent) FireInterstitialHiddenOnTarget(targetId string, event InterstitialHiddenEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.interstitialHidden",
        Params: event,
    })
}
func (agent *PageAgent) FireNavigationRequested(event NavigationRequestedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.navigationRequested",
        Params: event,
    })
}
func (agent *PageAgent) FireNavigationRequestedOnTarget(targetId string, event NavigationRequestedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.navigationRequested",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *PageAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *PageAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *PageAgent) AddScriptToEvaluateOnLoadNotify() <-chan AddScriptToEvaluateOnLoadCommand {
    outChan := make(chan AddScriptToEvaluateOnLoadCommand)
    agent.commandChans.AddScriptToEvaluateOnLoad = append(agent.commandChans.AddScriptToEvaluateOnLoad, outChan)
    return outChan
}
func (agent *PageAgent) RemoveScriptToEvaluateOnLoadNotify() <-chan RemoveScriptToEvaluateOnLoadCommand {
    outChan := make(chan RemoveScriptToEvaluateOnLoadCommand)
    agent.commandChans.RemoveScriptToEvaluateOnLoad = append(agent.commandChans.RemoveScriptToEvaluateOnLoad, outChan)
    return outChan
}
func (agent *PageAgent) SetAutoAttachToCreatedPagesNotify() <-chan SetAutoAttachToCreatedPagesCommand {
    outChan := make(chan SetAutoAttachToCreatedPagesCommand)
    agent.commandChans.SetAutoAttachToCreatedPages = append(agent.commandChans.SetAutoAttachToCreatedPages, outChan)
    return outChan
}
func (agent *PageAgent) ReloadNotify() <-chan ReloadCommand {
    outChan := make(chan ReloadCommand)
    agent.commandChans.Reload = append(agent.commandChans.Reload, outChan)
    return outChan
}
func (agent *PageAgent) NavigateNotify() <-chan NavigateCommand {
    outChan := make(chan NavigateCommand)
    agent.commandChans.Navigate = append(agent.commandChans.Navigate, outChan)
    return outChan
}
func (agent *PageAgent) StopLoadingNotify() <-chan StopLoadingCommand {
    outChan := make(chan StopLoadingCommand)
    agent.commandChans.StopLoading = append(agent.commandChans.StopLoading, outChan)
    return outChan
}
func (agent *PageAgent) GetNavigationHistoryNotify() <-chan GetNavigationHistoryCommand {
    outChan := make(chan GetNavigationHistoryCommand)
    agent.commandChans.GetNavigationHistory = append(agent.commandChans.GetNavigationHistory, outChan)
    return outChan
}
func (agent *PageAgent) NavigateToHistoryEntryNotify() <-chan NavigateToHistoryEntryCommand {
    outChan := make(chan NavigateToHistoryEntryCommand)
    agent.commandChans.NavigateToHistoryEntry = append(agent.commandChans.NavigateToHistoryEntry, outChan)
    return outChan
}
func (agent *PageAgent) GetCookiesNotify() <-chan GetCookiesCommand {
    outChan := make(chan GetCookiesCommand)
    agent.commandChans.GetCookies = append(agent.commandChans.GetCookies, outChan)
    return outChan
}
func (agent *PageAgent) DeleteCookieNotify() <-chan DeleteCookieCommand {
    outChan := make(chan DeleteCookieCommand)
    agent.commandChans.DeleteCookie = append(agent.commandChans.DeleteCookie, outChan)
    return outChan
}
func (agent *PageAgent) GetResourceTreeNotify() <-chan GetResourceTreeCommand {
    outChan := make(chan GetResourceTreeCommand)
    agent.commandChans.GetResourceTree = append(agent.commandChans.GetResourceTree, outChan)
    return outChan
}
func (agent *PageAgent) GetResourceContentNotify() <-chan GetResourceContentCommand {
    outChan := make(chan GetResourceContentCommand)
    agent.commandChans.GetResourceContent = append(agent.commandChans.GetResourceContent, outChan)
    return outChan
}
func (agent *PageAgent) SearchInResourceNotify() <-chan SearchInResourceCommand {
    outChan := make(chan SearchInResourceCommand)
    agent.commandChans.SearchInResource = append(agent.commandChans.SearchInResource, outChan)
    return outChan
}
func (agent *PageAgent) SetDocumentContentNotify() <-chan SetDocumentContentCommand {
    outChan := make(chan SetDocumentContentCommand)
    agent.commandChans.SetDocumentContent = append(agent.commandChans.SetDocumentContent, outChan)
    return outChan
}
func (agent *PageAgent) SetDeviceMetricsOverrideNotify() <-chan SetDeviceMetricsOverrideCommand {
    outChan := make(chan SetDeviceMetricsOverrideCommand)
    agent.commandChans.SetDeviceMetricsOverride = append(agent.commandChans.SetDeviceMetricsOverride, outChan)
    return outChan
}
func (agent *PageAgent) ClearDeviceMetricsOverrideNotify() <-chan ClearDeviceMetricsOverrideCommand {
    outChan := make(chan ClearDeviceMetricsOverrideCommand)
    agent.commandChans.ClearDeviceMetricsOverride = append(agent.commandChans.ClearDeviceMetricsOverride, outChan)
    return outChan
}
func (agent *PageAgent) SetGeolocationOverrideNotify() <-chan SetGeolocationOverrideCommand {
    outChan := make(chan SetGeolocationOverrideCommand)
    agent.commandChans.SetGeolocationOverride = append(agent.commandChans.SetGeolocationOverride, outChan)
    return outChan
}
func (agent *PageAgent) ClearGeolocationOverrideNotify() <-chan ClearGeolocationOverrideCommand {
    outChan := make(chan ClearGeolocationOverrideCommand)
    agent.commandChans.ClearGeolocationOverride = append(agent.commandChans.ClearGeolocationOverride, outChan)
    return outChan
}
func (agent *PageAgent) SetDeviceOrientationOverrideNotify() <-chan SetDeviceOrientationOverrideCommand {
    outChan := make(chan SetDeviceOrientationOverrideCommand)
    agent.commandChans.SetDeviceOrientationOverride = append(agent.commandChans.SetDeviceOrientationOverride, outChan)
    return outChan
}
func (agent *PageAgent) ClearDeviceOrientationOverrideNotify() <-chan ClearDeviceOrientationOverrideCommand {
    outChan := make(chan ClearDeviceOrientationOverrideCommand)
    agent.commandChans.ClearDeviceOrientationOverride = append(agent.commandChans.ClearDeviceOrientationOverride, outChan)
    return outChan
}
func (agent *PageAgent) SetTouchEmulationEnabledNotify() <-chan SetTouchEmulationEnabledCommand {
    outChan := make(chan SetTouchEmulationEnabledCommand)
    agent.commandChans.SetTouchEmulationEnabled = append(agent.commandChans.SetTouchEmulationEnabled, outChan)
    return outChan
}
func (agent *PageAgent) CaptureScreenshotNotify() <-chan CaptureScreenshotCommand {
    outChan := make(chan CaptureScreenshotCommand)
    agent.commandChans.CaptureScreenshot = append(agent.commandChans.CaptureScreenshot, outChan)
    return outChan
}
func (agent *PageAgent) PrintToPDFNotify() <-chan PrintToPDFCommand {
    outChan := make(chan PrintToPDFCommand)
    agent.commandChans.PrintToPDF = append(agent.commandChans.PrintToPDF, outChan)
    return outChan
}
func (agent *PageAgent) StartScreencastNotify() <-chan StartScreencastCommand {
    outChan := make(chan StartScreencastCommand)
    agent.commandChans.StartScreencast = append(agent.commandChans.StartScreencast, outChan)
    return outChan
}
func (agent *PageAgent) StopScreencastNotify() <-chan StopScreencastCommand {
    outChan := make(chan StopScreencastCommand)
    agent.commandChans.StopScreencast = append(agent.commandChans.StopScreencast, outChan)
    return outChan
}
func (agent *PageAgent) ScreencastFrameAckNotify() <-chan ScreencastFrameAckCommand {
    outChan := make(chan ScreencastFrameAckCommand)
    agent.commandChans.ScreencastFrameAck = append(agent.commandChans.ScreencastFrameAck, outChan)
    return outChan
}
func (agent *PageAgent) HandleJavaScriptDialogNotify() <-chan HandleJavaScriptDialogCommand {
    outChan := make(chan HandleJavaScriptDialogCommand)
    agent.commandChans.HandleJavaScriptDialog = append(agent.commandChans.HandleJavaScriptDialog, outChan)
    return outChan
}
func (agent *PageAgent) SetColorPickerEnabledNotify() <-chan SetColorPickerEnabledCommand {
    outChan := make(chan SetColorPickerEnabledCommand)
    agent.commandChans.SetColorPickerEnabled = append(agent.commandChans.SetColorPickerEnabled, outChan)
    return outChan
}
func (agent *PageAgent) ConfigureOverlayNotify() <-chan ConfigureOverlayCommand {
    outChan := make(chan ConfigureOverlayCommand)
    agent.commandChans.ConfigureOverlay = append(agent.commandChans.ConfigureOverlay, outChan)
    return outChan
}
func (agent *PageAgent) GetAppManifestNotify() <-chan GetAppManifestCommand {
    outChan := make(chan GetAppManifestCommand)
    agent.commandChans.GetAppManifest = append(agent.commandChans.GetAppManifest, outChan)
    return outChan
}
func (agent *PageAgent) RequestAppBannerNotify() <-chan RequestAppBannerCommand {
    outChan := make(chan RequestAppBannerCommand)
    agent.commandChans.RequestAppBanner = append(agent.commandChans.RequestAppBanner, outChan)
    return outChan
}
func (agent *PageAgent) SetControlNavigationsNotify() <-chan SetControlNavigationsCommand {
    outChan := make(chan SetControlNavigationsCommand)
    agent.commandChans.SetControlNavigations = append(agent.commandChans.SetControlNavigations, outChan)
    return outChan
}
func (agent *PageAgent) ProcessNavigationNotify() <-chan ProcessNavigationCommand {
    outChan := make(chan ProcessNavigationCommand)
    agent.commandChans.ProcessNavigation = append(agent.commandChans.ProcessNavigation, outChan)
    return outChan
}
func (agent *PageAgent) GetLayoutMetricsNotify() <-chan GetLayoutMetricsCommand {
    outChan := make(chan GetLayoutMetricsCommand)
    agent.commandChans.GetLayoutMetrics = append(agent.commandChans.GetLayoutMetrics, outChan)
    return outChan
}
func init() {

}
