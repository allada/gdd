package page


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

type AddScriptToEvaluateOnLoadCommandFn struct {
    mux sync.RWMutex
    fn func(AddScriptToEvaluateOnLoadCommand)
}

func (a *AddScriptToEvaluateOnLoadCommandFn) Load() func(AddScriptToEvaluateOnLoadCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *AddScriptToEvaluateOnLoadCommandFn) Store(fn func(AddScriptToEvaluateOnLoadCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RemoveScriptToEvaluateOnLoadCommandFn struct {
    mux sync.RWMutex
    fn func(RemoveScriptToEvaluateOnLoadCommand)
}

func (a *RemoveScriptToEvaluateOnLoadCommandFn) Load() func(RemoveScriptToEvaluateOnLoadCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RemoveScriptToEvaluateOnLoadCommandFn) Store(fn func(RemoveScriptToEvaluateOnLoadCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetAutoAttachToCreatedPagesCommandFn struct {
    mux sync.RWMutex
    fn func(SetAutoAttachToCreatedPagesCommand)
}

func (a *SetAutoAttachToCreatedPagesCommandFn) Load() func(SetAutoAttachToCreatedPagesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetAutoAttachToCreatedPagesCommandFn) Store(fn func(SetAutoAttachToCreatedPagesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ReloadCommandFn struct {
    mux sync.RWMutex
    fn func(ReloadCommand)
}

func (a *ReloadCommandFn) Load() func(ReloadCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ReloadCommandFn) Store(fn func(ReloadCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type NavigateCommandFn struct {
    mux sync.RWMutex
    fn func(NavigateCommand)
}

func (a *NavigateCommandFn) Load() func(NavigateCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *NavigateCommandFn) Store(fn func(NavigateCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StopLoadingCommandFn struct {
    mux sync.RWMutex
    fn func(StopLoadingCommand)
}

func (a *StopLoadingCommandFn) Load() func(StopLoadingCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StopLoadingCommandFn) Store(fn func(StopLoadingCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetNavigationHistoryCommandFn struct {
    mux sync.RWMutex
    fn func(GetNavigationHistoryCommand)
}

func (a *GetNavigationHistoryCommandFn) Load() func(GetNavigationHistoryCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetNavigationHistoryCommandFn) Store(fn func(GetNavigationHistoryCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type NavigateToHistoryEntryCommandFn struct {
    mux sync.RWMutex
    fn func(NavigateToHistoryEntryCommand)
}

func (a *NavigateToHistoryEntryCommandFn) Load() func(NavigateToHistoryEntryCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *NavigateToHistoryEntryCommandFn) Store(fn func(NavigateToHistoryEntryCommand)) {
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

type GetResourceTreeCommandFn struct {
    mux sync.RWMutex
    fn func(GetResourceTreeCommand)
}

func (a *GetResourceTreeCommandFn) Load() func(GetResourceTreeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetResourceTreeCommandFn) Store(fn func(GetResourceTreeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetResourceContentCommandFn struct {
    mux sync.RWMutex
    fn func(GetResourceContentCommand)
}

func (a *GetResourceContentCommandFn) Load() func(GetResourceContentCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetResourceContentCommandFn) Store(fn func(GetResourceContentCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SearchInResourceCommandFn struct {
    mux sync.RWMutex
    fn func(SearchInResourceCommand)
}

func (a *SearchInResourceCommandFn) Load() func(SearchInResourceCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SearchInResourceCommandFn) Store(fn func(SearchInResourceCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetDocumentContentCommandFn struct {
    mux sync.RWMutex
    fn func(SetDocumentContentCommand)
}

func (a *SetDocumentContentCommandFn) Load() func(SetDocumentContentCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetDocumentContentCommandFn) Store(fn func(SetDocumentContentCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

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

type CaptureScreenshotCommandFn struct {
    mux sync.RWMutex
    fn func(CaptureScreenshotCommand)
}

func (a *CaptureScreenshotCommandFn) Load() func(CaptureScreenshotCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CaptureScreenshotCommandFn) Store(fn func(CaptureScreenshotCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type PrintToPDFCommandFn struct {
    mux sync.RWMutex
    fn func(PrintToPDFCommand)
}

func (a *PrintToPDFCommandFn) Load() func(PrintToPDFCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *PrintToPDFCommandFn) Store(fn func(PrintToPDFCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StartScreencastCommandFn struct {
    mux sync.RWMutex
    fn func(StartScreencastCommand)
}

func (a *StartScreencastCommandFn) Load() func(StartScreencastCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StartScreencastCommandFn) Store(fn func(StartScreencastCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StopScreencastCommandFn struct {
    mux sync.RWMutex
    fn func(StopScreencastCommand)
}

func (a *StopScreencastCommandFn) Load() func(StopScreencastCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StopScreencastCommandFn) Store(fn func(StopScreencastCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ScreencastFrameAckCommandFn struct {
    mux sync.RWMutex
    fn func(ScreencastFrameAckCommand)
}

func (a *ScreencastFrameAckCommandFn) Load() func(ScreencastFrameAckCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ScreencastFrameAckCommandFn) Store(fn func(ScreencastFrameAckCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type HandleJavaScriptDialogCommandFn struct {
    mux sync.RWMutex
    fn func(HandleJavaScriptDialogCommand)
}

func (a *HandleJavaScriptDialogCommandFn) Load() func(HandleJavaScriptDialogCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *HandleJavaScriptDialogCommandFn) Store(fn func(HandleJavaScriptDialogCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetColorPickerEnabledCommandFn struct {
    mux sync.RWMutex
    fn func(SetColorPickerEnabledCommand)
}

func (a *SetColorPickerEnabledCommandFn) Load() func(SetColorPickerEnabledCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetColorPickerEnabledCommandFn) Store(fn func(SetColorPickerEnabledCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ConfigureOverlayCommandFn struct {
    mux sync.RWMutex
    fn func(ConfigureOverlayCommand)
}

func (a *ConfigureOverlayCommandFn) Load() func(ConfigureOverlayCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ConfigureOverlayCommandFn) Store(fn func(ConfigureOverlayCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetAppManifestCommandFn struct {
    mux sync.RWMutex
    fn func(GetAppManifestCommand)
}

func (a *GetAppManifestCommandFn) Load() func(GetAppManifestCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetAppManifestCommandFn) Store(fn func(GetAppManifestCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RequestAppBannerCommandFn struct {
    mux sync.RWMutex
    fn func(RequestAppBannerCommand)
}

func (a *RequestAppBannerCommandFn) Load() func(RequestAppBannerCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RequestAppBannerCommandFn) Store(fn func(RequestAppBannerCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetControlNavigationsCommandFn struct {
    mux sync.RWMutex
    fn func(SetControlNavigationsCommand)
}

func (a *SetControlNavigationsCommandFn) Load() func(SetControlNavigationsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetControlNavigationsCommandFn) Store(fn func(SetControlNavigationsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ProcessNavigationCommandFn struct {
    mux sync.RWMutex
    fn func(ProcessNavigationCommand)
}

func (a *ProcessNavigationCommandFn) Load() func(ProcessNavigationCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ProcessNavigationCommandFn) Store(fn func(ProcessNavigationCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetLayoutMetricsCommandFn struct {
    mux sync.RWMutex
    fn func(GetLayoutMetricsCommand)
}

func (a *GetLayoutMetricsCommandFn) Load() func(GetLayoutMetricsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetLayoutMetricsCommandFn) Store(fn func(GetLayoutMetricsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type PageAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
        AddScriptToEvaluateOnLoad AddScriptToEvaluateOnLoadCommandFn
        RemoveScriptToEvaluateOnLoad RemoveScriptToEvaluateOnLoadCommandFn
        SetAutoAttachToCreatedPages SetAutoAttachToCreatedPagesCommandFn
        Reload ReloadCommandFn
        Navigate NavigateCommandFn
        StopLoading StopLoadingCommandFn
        GetNavigationHistory GetNavigationHistoryCommandFn
        NavigateToHistoryEntry NavigateToHistoryEntryCommandFn
        DeleteCookie DeleteCookieCommandFn
        GetResourceTree GetResourceTreeCommandFn
        GetResourceContent GetResourceContentCommandFn
        SearchInResource SearchInResourceCommandFn
        SetDocumentContent SetDocumentContentCommandFn
        SetDeviceMetricsOverride SetDeviceMetricsOverrideCommandFn
        ClearDeviceMetricsOverride ClearDeviceMetricsOverrideCommandFn
        SetGeolocationOverride SetGeolocationOverrideCommandFn
        ClearGeolocationOverride ClearGeolocationOverrideCommandFn
        SetDeviceOrientationOverride SetDeviceOrientationOverrideCommandFn
        ClearDeviceOrientationOverride ClearDeviceOrientationOverrideCommandFn
        SetTouchEmulationEnabled SetTouchEmulationEnabledCommandFn
        CaptureScreenshot CaptureScreenshotCommandFn
        PrintToPDF PrintToPDFCommandFn
        StartScreencast StartScreencastCommandFn
        StopScreencast StopScreencastCommandFn
        ScreencastFrameAck ScreencastFrameAckCommandFn
        HandleJavaScriptDialog HandleJavaScriptDialogCommandFn
        SetColorPickerEnabled SetColorPickerEnabledCommandFn
        ConfigureOverlay ConfigureOverlayCommandFn
        GetAppManifest GetAppManifestCommandFn
        RequestAppBanner RequestAppBannerCommandFn
        SetControlNavigations SetControlNavigationsCommandFn
        ProcessNavigation ProcessNavigationCommandFn
        GetLayoutMetrics GetLayoutMetricsCommandFn
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
        case "addScriptToEvaluateOnLoad":
            var out AddScriptToEvaluateOnLoadCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.AddScriptToEvaluateOnLoad.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "removeScriptToEvaluateOnLoad":
            var out RemoveScriptToEvaluateOnLoadCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RemoveScriptToEvaluateOnLoad.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setAutoAttachToCreatedPages":
            var out SetAutoAttachToCreatedPagesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetAutoAttachToCreatedPages.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "reload":
            var out ReloadCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Reload.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "navigate":
            var out NavigateCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Navigate.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "stopLoading":
            var out StopLoadingCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StopLoading.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getNavigationHistory":
            var out GetNavigationHistoryCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetNavigationHistory.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "navigateToHistoryEntry":
            var out NavigateToHistoryEntryCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.NavigateToHistoryEntry.Load()
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
        case "getResourceTree":
            var out GetResourceTreeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetResourceTree.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getResourceContent":
            var out GetResourceContentCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetResourceContent.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "searchInResource":
            var out SearchInResourceCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SearchInResource.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setDocumentContent":
            var out SetDocumentContentCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetDocumentContent.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
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
        case "setTouchEmulationEnabled":
            var out SetTouchEmulationEnabledCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetTouchEmulationEnabled.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "captureScreenshot":
            var out CaptureScreenshotCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CaptureScreenshot.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "printToPDF":
            var out PrintToPDFCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.PrintToPDF.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "startScreencast":
            var out StartScreencastCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StartScreencast.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "stopScreencast":
            var out StopScreencastCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StopScreencast.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "screencastFrameAck":
            var out ScreencastFrameAckCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ScreencastFrameAck.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "handleJavaScriptDialog":
            var out HandleJavaScriptDialogCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.HandleJavaScriptDialog.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setColorPickerEnabled":
            var out SetColorPickerEnabledCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetColorPickerEnabled.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "configureOverlay":
            var out ConfigureOverlayCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ConfigureOverlay.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getAppManifest":
            var out GetAppManifestCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetAppManifest.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "requestAppBanner":
            var out RequestAppBannerCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RequestAppBanner.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setControlNavigations":
            var out SetControlNavigationsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetControlNavigations.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "processNavigation":
            var out ProcessNavigationCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ProcessNavigation.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getLayoutMetrics":
            var out GetLayoutMetricsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetLayoutMetrics.Load()
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
func (agent *PageAgent) FireDomContentEventFired(event DomContentEventFiredEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.domContentEventFired",
        Params: event,
    })
}
func (agent *PageAgent) FireDomContentEventFiredOnTarget(targetId string,event DomContentEventFiredEvent) {
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
func (agent *PageAgent) FireLoadEventFiredOnTarget(targetId string,event LoadEventFiredEvent) {
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
func (agent *PageAgent) FireFrameAttachedOnTarget(targetId string,event FrameAttachedEvent) {
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
func (agent *PageAgent) FireFrameNavigatedOnTarget(targetId string,event FrameNavigatedEvent) {
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
func (agent *PageAgent) FireFrameDetachedOnTarget(targetId string,event FrameDetachedEvent) {
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
func (agent *PageAgent) FireFrameStartedLoadingOnTarget(targetId string,event FrameStartedLoadingEvent) {
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
func (agent *PageAgent) FireFrameStoppedLoadingOnTarget(targetId string,event FrameStoppedLoadingEvent) {
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
func (agent *PageAgent) FireFrameScheduledNavigationOnTarget(targetId string,event FrameScheduledNavigationEvent) {
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
func (agent *PageAgent) FireFrameClearedScheduledNavigationOnTarget(targetId string,event FrameClearedScheduledNavigationEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.frameClearedScheduledNavigation",
        Params: event,
    })
}
func (agent *PageAgent) FireFrameResized() {
    agent.conn.Send(shared.Notification{
        Method: "Page.frameResized",
    })
}
func (agent *PageAgent) FireFrameResizedOnTarget(targetId string) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.frameResized",
    })
}
func (agent *PageAgent) FireJavascriptDialogOpening(event JavascriptDialogOpeningEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.javascriptDialogOpening",
        Params: event,
    })
}
func (agent *PageAgent) FireJavascriptDialogOpeningOnTarget(targetId string,event JavascriptDialogOpeningEvent) {
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
func (agent *PageAgent) FireJavascriptDialogClosedOnTarget(targetId string,event JavascriptDialogClosedEvent) {
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
func (agent *PageAgent) FireScreencastFrameOnTarget(targetId string,event ScreencastFrameEvent) {
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
func (agent *PageAgent) FireScreencastVisibilityChangedOnTarget(targetId string,event ScreencastVisibilityChangedEvent) {
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
func (agent *PageAgent) FireColorPickedOnTarget(targetId string,event ColorPickedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.colorPicked",
        Params: event,
    })
}
func (agent *PageAgent) FireInterstitialShown() {
    agent.conn.Send(shared.Notification{
        Method: "Page.interstitialShown",
    })
}
func (agent *PageAgent) FireInterstitialShownOnTarget(targetId string) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.interstitialShown",
    })
}
func (agent *PageAgent) FireInterstitialHidden() {
    agent.conn.Send(shared.Notification{
        Method: "Page.interstitialHidden",
    })
}
func (agent *PageAgent) FireInterstitialHiddenOnTarget(targetId string) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.interstitialHidden",
    })
}
func (agent *PageAgent) FireNavigationRequested(event NavigationRequestedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Page.navigationRequested",
        Params: event,
    })
}
func (agent *PageAgent) FireNavigationRequestedOnTarget(targetId string,event NavigationRequestedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Page.navigationRequested",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *PageAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *PageAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *PageAgent) SetAddScriptToEvaluateOnLoadHandler(handler func(AddScriptToEvaluateOnLoadCommand)) {
    agent.commandHandlers.AddScriptToEvaluateOnLoad.Store(handler)
}
func (agent *PageAgent) SetRemoveScriptToEvaluateOnLoadHandler(handler func(RemoveScriptToEvaluateOnLoadCommand)) {
    agent.commandHandlers.RemoveScriptToEvaluateOnLoad.Store(handler)
}
func (agent *PageAgent) SetSetAutoAttachToCreatedPagesHandler(handler func(SetAutoAttachToCreatedPagesCommand)) {
    agent.commandHandlers.SetAutoAttachToCreatedPages.Store(handler)
}
func (agent *PageAgent) SetReloadHandler(handler func(ReloadCommand)) {
    agent.commandHandlers.Reload.Store(handler)
}
func (agent *PageAgent) SetNavigateHandler(handler func(NavigateCommand)) {
    agent.commandHandlers.Navigate.Store(handler)
}
func (agent *PageAgent) SetStopLoadingHandler(handler func(StopLoadingCommand)) {
    agent.commandHandlers.StopLoading.Store(handler)
}
func (agent *PageAgent) SetGetNavigationHistoryHandler(handler func(GetNavigationHistoryCommand)) {
    agent.commandHandlers.GetNavigationHistory.Store(handler)
}
func (agent *PageAgent) SetNavigateToHistoryEntryHandler(handler func(NavigateToHistoryEntryCommand)) {
    agent.commandHandlers.NavigateToHistoryEntry.Store(handler)
}
func (agent *PageAgent) SetDeleteCookieHandler(handler func(DeleteCookieCommand)) {
    agent.commandHandlers.DeleteCookie.Store(handler)
}
func (agent *PageAgent) SetGetResourceTreeHandler(handler func(GetResourceTreeCommand)) {
    agent.commandHandlers.GetResourceTree.Store(handler)
}
func (agent *PageAgent) SetGetResourceContentHandler(handler func(GetResourceContentCommand)) {
    agent.commandHandlers.GetResourceContent.Store(handler)
}
func (agent *PageAgent) SetSearchInResourceHandler(handler func(SearchInResourceCommand)) {
    agent.commandHandlers.SearchInResource.Store(handler)
}
func (agent *PageAgent) SetSetDocumentContentHandler(handler func(SetDocumentContentCommand)) {
    agent.commandHandlers.SetDocumentContent.Store(handler)
}
func (agent *PageAgent) SetSetDeviceMetricsOverrideHandler(handler func(SetDeviceMetricsOverrideCommand)) {
    agent.commandHandlers.SetDeviceMetricsOverride.Store(handler)
}
func (agent *PageAgent) SetClearDeviceMetricsOverrideHandler(handler func(ClearDeviceMetricsOverrideCommand)) {
    agent.commandHandlers.ClearDeviceMetricsOverride.Store(handler)
}
func (agent *PageAgent) SetSetGeolocationOverrideHandler(handler func(SetGeolocationOverrideCommand)) {
    agent.commandHandlers.SetGeolocationOverride.Store(handler)
}
func (agent *PageAgent) SetClearGeolocationOverrideHandler(handler func(ClearGeolocationOverrideCommand)) {
    agent.commandHandlers.ClearGeolocationOverride.Store(handler)
}
func (agent *PageAgent) SetSetDeviceOrientationOverrideHandler(handler func(SetDeviceOrientationOverrideCommand)) {
    agent.commandHandlers.SetDeviceOrientationOverride.Store(handler)
}
func (agent *PageAgent) SetClearDeviceOrientationOverrideHandler(handler func(ClearDeviceOrientationOverrideCommand)) {
    agent.commandHandlers.ClearDeviceOrientationOverride.Store(handler)
}
func (agent *PageAgent) SetSetTouchEmulationEnabledHandler(handler func(SetTouchEmulationEnabledCommand)) {
    agent.commandHandlers.SetTouchEmulationEnabled.Store(handler)
}
func (agent *PageAgent) SetCaptureScreenshotHandler(handler func(CaptureScreenshotCommand)) {
    agent.commandHandlers.CaptureScreenshot.Store(handler)
}
func (agent *PageAgent) SetPrintToPDFHandler(handler func(PrintToPDFCommand)) {
    agent.commandHandlers.PrintToPDF.Store(handler)
}
func (agent *PageAgent) SetStartScreencastHandler(handler func(StartScreencastCommand)) {
    agent.commandHandlers.StartScreencast.Store(handler)
}
func (agent *PageAgent) SetStopScreencastHandler(handler func(StopScreencastCommand)) {
    agent.commandHandlers.StopScreencast.Store(handler)
}
func (agent *PageAgent) SetScreencastFrameAckHandler(handler func(ScreencastFrameAckCommand)) {
    agent.commandHandlers.ScreencastFrameAck.Store(handler)
}
func (agent *PageAgent) SetHandleJavaScriptDialogHandler(handler func(HandleJavaScriptDialogCommand)) {
    agent.commandHandlers.HandleJavaScriptDialog.Store(handler)
}
func (agent *PageAgent) SetSetColorPickerEnabledHandler(handler func(SetColorPickerEnabledCommand)) {
    agent.commandHandlers.SetColorPickerEnabled.Store(handler)
}
func (agent *PageAgent) SetConfigureOverlayHandler(handler func(ConfigureOverlayCommand)) {
    agent.commandHandlers.ConfigureOverlay.Store(handler)
}
func (agent *PageAgent) SetGetAppManifestHandler(handler func(GetAppManifestCommand)) {
    agent.commandHandlers.GetAppManifest.Store(handler)
}
func (agent *PageAgent) SetRequestAppBannerHandler(handler func(RequestAppBannerCommand)) {
    agent.commandHandlers.RequestAppBanner.Store(handler)
}
func (agent *PageAgent) SetSetControlNavigationsHandler(handler func(SetControlNavigationsCommand)) {
    agent.commandHandlers.SetControlNavigations.Store(handler)
}
func (agent *PageAgent) SetProcessNavigationHandler(handler func(ProcessNavigationCommand)) {
    agent.commandHandlers.ProcessNavigation.Store(handler)
}
func (agent *PageAgent) SetGetLayoutMetricsHandler(handler func(GetLayoutMetricsCommand)) {
    agent.commandHandlers.GetLayoutMetrics.Store(handler)
}
func init() {

}
