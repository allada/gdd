package css


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

type GetMatchedStylesForNodeCommandFn struct {
    mux sync.RWMutex
    fn func(GetMatchedStylesForNodeCommand)
}

func (a *GetMatchedStylesForNodeCommandFn) Load() func(GetMatchedStylesForNodeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetMatchedStylesForNodeCommandFn) Store(fn func(GetMatchedStylesForNodeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetInlineStylesForNodeCommandFn struct {
    mux sync.RWMutex
    fn func(GetInlineStylesForNodeCommand)
}

func (a *GetInlineStylesForNodeCommandFn) Load() func(GetInlineStylesForNodeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetInlineStylesForNodeCommandFn) Store(fn func(GetInlineStylesForNodeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetComputedStyleForNodeCommandFn struct {
    mux sync.RWMutex
    fn func(GetComputedStyleForNodeCommand)
}

func (a *GetComputedStyleForNodeCommandFn) Load() func(GetComputedStyleForNodeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetComputedStyleForNodeCommandFn) Store(fn func(GetComputedStyleForNodeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetPlatformFontsForNodeCommandFn struct {
    mux sync.RWMutex
    fn func(GetPlatformFontsForNodeCommand)
}

func (a *GetPlatformFontsForNodeCommandFn) Load() func(GetPlatformFontsForNodeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetPlatformFontsForNodeCommandFn) Store(fn func(GetPlatformFontsForNodeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetStyleSheetTextCommandFn struct {
    mux sync.RWMutex
    fn func(GetStyleSheetTextCommand)
}

func (a *GetStyleSheetTextCommandFn) Load() func(GetStyleSheetTextCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetStyleSheetTextCommandFn) Store(fn func(GetStyleSheetTextCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CollectClassNamesCommandFn struct {
    mux sync.RWMutex
    fn func(CollectClassNamesCommand)
}

func (a *CollectClassNamesCommandFn) Load() func(CollectClassNamesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CollectClassNamesCommandFn) Store(fn func(CollectClassNamesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetStyleSheetTextCommandFn struct {
    mux sync.RWMutex
    fn func(SetStyleSheetTextCommand)
}

func (a *SetStyleSheetTextCommandFn) Load() func(SetStyleSheetTextCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetStyleSheetTextCommandFn) Store(fn func(SetStyleSheetTextCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetRuleSelectorCommandFn struct {
    mux sync.RWMutex
    fn func(SetRuleSelectorCommand)
}

func (a *SetRuleSelectorCommandFn) Load() func(SetRuleSelectorCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetRuleSelectorCommandFn) Store(fn func(SetRuleSelectorCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetKeyframeKeyCommandFn struct {
    mux sync.RWMutex
    fn func(SetKeyframeKeyCommand)
}

func (a *SetKeyframeKeyCommandFn) Load() func(SetKeyframeKeyCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetKeyframeKeyCommandFn) Store(fn func(SetKeyframeKeyCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetStyleTextsCommandFn struct {
    mux sync.RWMutex
    fn func(SetStyleTextsCommand)
}

func (a *SetStyleTextsCommandFn) Load() func(SetStyleTextsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetStyleTextsCommandFn) Store(fn func(SetStyleTextsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetMediaTextCommandFn struct {
    mux sync.RWMutex
    fn func(SetMediaTextCommand)
}

func (a *SetMediaTextCommandFn) Load() func(SetMediaTextCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetMediaTextCommandFn) Store(fn func(SetMediaTextCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CreateStyleSheetCommandFn struct {
    mux sync.RWMutex
    fn func(CreateStyleSheetCommand)
}

func (a *CreateStyleSheetCommandFn) Load() func(CreateStyleSheetCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CreateStyleSheetCommandFn) Store(fn func(CreateStyleSheetCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type AddRuleCommandFn struct {
    mux sync.RWMutex
    fn func(AddRuleCommand)
}

func (a *AddRuleCommandFn) Load() func(AddRuleCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *AddRuleCommandFn) Store(fn func(AddRuleCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ForcePseudoStateCommandFn struct {
    mux sync.RWMutex
    fn func(ForcePseudoStateCommand)
}

func (a *ForcePseudoStateCommandFn) Load() func(ForcePseudoStateCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ForcePseudoStateCommandFn) Store(fn func(ForcePseudoStateCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetMediaQueriesCommandFn struct {
    mux sync.RWMutex
    fn func(GetMediaQueriesCommand)
}

func (a *GetMediaQueriesCommandFn) Load() func(GetMediaQueriesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetMediaQueriesCommandFn) Store(fn func(GetMediaQueriesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetEffectivePropertyValueForNodeCommandFn struct {
    mux sync.RWMutex
    fn func(SetEffectivePropertyValueForNodeCommand)
}

func (a *SetEffectivePropertyValueForNodeCommandFn) Load() func(SetEffectivePropertyValueForNodeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetEffectivePropertyValueForNodeCommandFn) Store(fn func(SetEffectivePropertyValueForNodeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetBackgroundColorsCommandFn struct {
    mux sync.RWMutex
    fn func(GetBackgroundColorsCommand)
}

func (a *GetBackgroundColorsCommandFn) Load() func(GetBackgroundColorsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetBackgroundColorsCommandFn) Store(fn func(GetBackgroundColorsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetLayoutTreeAndStylesCommandFn struct {
    mux sync.RWMutex
    fn func(GetLayoutTreeAndStylesCommand)
}

func (a *GetLayoutTreeAndStylesCommandFn) Load() func(GetLayoutTreeAndStylesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetLayoutTreeAndStylesCommandFn) Store(fn func(GetLayoutTreeAndStylesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StartRuleUsageTrackingCommandFn struct {
    mux sync.RWMutex
    fn func(StartRuleUsageTrackingCommand)
}

func (a *StartRuleUsageTrackingCommandFn) Load() func(StartRuleUsageTrackingCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StartRuleUsageTrackingCommandFn) Store(fn func(StartRuleUsageTrackingCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StopRuleUsageTrackingCommandFn struct {
    mux sync.RWMutex
    fn func(StopRuleUsageTrackingCommand)
}

func (a *StopRuleUsageTrackingCommandFn) Load() func(StopRuleUsageTrackingCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StopRuleUsageTrackingCommandFn) Store(fn func(StopRuleUsageTrackingCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CSSAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
        GetMatchedStylesForNode GetMatchedStylesForNodeCommandFn
        GetInlineStylesForNode GetInlineStylesForNodeCommandFn
        GetComputedStyleForNode GetComputedStyleForNodeCommandFn
        GetPlatformFontsForNode GetPlatformFontsForNodeCommandFn
        GetStyleSheetText GetStyleSheetTextCommandFn
        CollectClassNames CollectClassNamesCommandFn
        SetStyleSheetText SetStyleSheetTextCommandFn
        SetRuleSelector SetRuleSelectorCommandFn
        SetKeyframeKey SetKeyframeKeyCommandFn
        SetStyleTexts SetStyleTextsCommandFn
        SetMediaText SetMediaTextCommandFn
        CreateStyleSheet CreateStyleSheetCommandFn
        AddRule AddRuleCommandFn
        ForcePseudoState ForcePseudoStateCommandFn
        GetMediaQueries GetMediaQueriesCommandFn
        SetEffectivePropertyValueForNode SetEffectivePropertyValueForNodeCommandFn
        GetBackgroundColors GetBackgroundColorsCommandFn
        GetLayoutTreeAndStyles GetLayoutTreeAndStylesCommandFn
        StartRuleUsageTracking StartRuleUsageTrackingCommandFn
        StopRuleUsageTracking StopRuleUsageTrackingCommandFn
    }
}
func NewAgent(conn *shared.Connection) *CSSAgent {
    agent := &CSSAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *CSSAgent) Name() string {
    return "CSS"
}

func (agent *CSSAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "getMatchedStylesForNode":
            var out GetMatchedStylesForNodeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetMatchedStylesForNode.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getInlineStylesForNode":
            var out GetInlineStylesForNodeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetInlineStylesForNode.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getComputedStyleForNode":
            var out GetComputedStyleForNodeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetComputedStyleForNode.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getPlatformFontsForNode":
            var out GetPlatformFontsForNodeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetPlatformFontsForNode.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getStyleSheetText":
            var out GetStyleSheetTextCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetStyleSheetText.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "collectClassNames":
            var out CollectClassNamesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CollectClassNames.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setStyleSheetText":
            var out SetStyleSheetTextCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetStyleSheetText.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setRuleSelector":
            var out SetRuleSelectorCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetRuleSelector.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setKeyframeKey":
            var out SetKeyframeKeyCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetKeyframeKey.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setStyleTexts":
            var out SetStyleTextsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetStyleTexts.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setMediaText":
            var out SetMediaTextCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetMediaText.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "createStyleSheet":
            var out CreateStyleSheetCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CreateStyleSheet.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "addRule":
            var out AddRuleCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.AddRule.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "forcePseudoState":
            var out ForcePseudoStateCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ForcePseudoState.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getMediaQueries":
            var out GetMediaQueriesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetMediaQueries.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setEffectivePropertyValueForNode":
            var out SetEffectivePropertyValueForNodeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetEffectivePropertyValueForNode.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getBackgroundColors":
            var out GetBackgroundColorsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetBackgroundColors.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getLayoutTreeAndStyles":
            var out GetLayoutTreeAndStylesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetLayoutTreeAndStyles.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "startRuleUsageTracking":
            var out StartRuleUsageTrackingCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StartRuleUsageTracking.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "stopRuleUsageTracking":
            var out StopRuleUsageTrackingCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StopRuleUsageTracking.Load()
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
func (agent *CSSAgent) FireMediaQueryResultChanged() {
    agent.conn.Send(shared.Notification{
        Method: "CSS.mediaQueryResultChanged",
    })
}
func (agent *CSSAgent) FireMediaQueryResultChangedOnTarget(targetId string) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "CSS.mediaQueryResultChanged",
    })
}
func (agent *CSSAgent) FireFontsUpdated() {
    agent.conn.Send(shared.Notification{
        Method: "CSS.fontsUpdated",
    })
}
func (agent *CSSAgent) FireFontsUpdatedOnTarget(targetId string) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "CSS.fontsUpdated",
    })
}
func (agent *CSSAgent) FireStyleSheetChanged(event StyleSheetChangedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "CSS.styleSheetChanged",
        Params: event,
    })
}
func (agent *CSSAgent) FireStyleSheetChangedOnTarget(targetId string,event StyleSheetChangedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "CSS.styleSheetChanged",
        Params: event,
    })
}
func (agent *CSSAgent) FireStyleSheetAdded(event StyleSheetAddedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "CSS.styleSheetAdded",
        Params: event,
    })
}
func (agent *CSSAgent) FireStyleSheetAddedOnTarget(targetId string,event StyleSheetAddedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "CSS.styleSheetAdded",
        Params: event,
    })
}
func (agent *CSSAgent) FireStyleSheetRemoved(event StyleSheetRemovedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "CSS.styleSheetRemoved",
        Params: event,
    })
}
func (agent *CSSAgent) FireStyleSheetRemovedOnTarget(targetId string,event StyleSheetRemovedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "CSS.styleSheetRemoved",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *CSSAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *CSSAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *CSSAgent) SetGetMatchedStylesForNodeHandler(handler func(GetMatchedStylesForNodeCommand)) {
    agent.commandHandlers.GetMatchedStylesForNode.Store(handler)
}
func (agent *CSSAgent) SetGetInlineStylesForNodeHandler(handler func(GetInlineStylesForNodeCommand)) {
    agent.commandHandlers.GetInlineStylesForNode.Store(handler)
}
func (agent *CSSAgent) SetGetComputedStyleForNodeHandler(handler func(GetComputedStyleForNodeCommand)) {
    agent.commandHandlers.GetComputedStyleForNode.Store(handler)
}
func (agent *CSSAgent) SetGetPlatformFontsForNodeHandler(handler func(GetPlatformFontsForNodeCommand)) {
    agent.commandHandlers.GetPlatformFontsForNode.Store(handler)
}
func (agent *CSSAgent) SetGetStyleSheetTextHandler(handler func(GetStyleSheetTextCommand)) {
    agent.commandHandlers.GetStyleSheetText.Store(handler)
}
func (agent *CSSAgent) SetCollectClassNamesHandler(handler func(CollectClassNamesCommand)) {
    agent.commandHandlers.CollectClassNames.Store(handler)
}
func (agent *CSSAgent) SetSetStyleSheetTextHandler(handler func(SetStyleSheetTextCommand)) {
    agent.commandHandlers.SetStyleSheetText.Store(handler)
}
func (agent *CSSAgent) SetSetRuleSelectorHandler(handler func(SetRuleSelectorCommand)) {
    agent.commandHandlers.SetRuleSelector.Store(handler)
}
func (agent *CSSAgent) SetSetKeyframeKeyHandler(handler func(SetKeyframeKeyCommand)) {
    agent.commandHandlers.SetKeyframeKey.Store(handler)
}
func (agent *CSSAgent) SetSetStyleTextsHandler(handler func(SetStyleTextsCommand)) {
    agent.commandHandlers.SetStyleTexts.Store(handler)
}
func (agent *CSSAgent) SetSetMediaTextHandler(handler func(SetMediaTextCommand)) {
    agent.commandHandlers.SetMediaText.Store(handler)
}
func (agent *CSSAgent) SetCreateStyleSheetHandler(handler func(CreateStyleSheetCommand)) {
    agent.commandHandlers.CreateStyleSheet.Store(handler)
}
func (agent *CSSAgent) SetAddRuleHandler(handler func(AddRuleCommand)) {
    agent.commandHandlers.AddRule.Store(handler)
}
func (agent *CSSAgent) SetForcePseudoStateHandler(handler func(ForcePseudoStateCommand)) {
    agent.commandHandlers.ForcePseudoState.Store(handler)
}
func (agent *CSSAgent) SetGetMediaQueriesHandler(handler func(GetMediaQueriesCommand)) {
    agent.commandHandlers.GetMediaQueries.Store(handler)
}
func (agent *CSSAgent) SetSetEffectivePropertyValueForNodeHandler(handler func(SetEffectivePropertyValueForNodeCommand)) {
    agent.commandHandlers.SetEffectivePropertyValueForNode.Store(handler)
}
func (agent *CSSAgent) SetGetBackgroundColorsHandler(handler func(GetBackgroundColorsCommand)) {
    agent.commandHandlers.GetBackgroundColors.Store(handler)
}
func (agent *CSSAgent) SetGetLayoutTreeAndStylesHandler(handler func(GetLayoutTreeAndStylesCommand)) {
    agent.commandHandlers.GetLayoutTreeAndStyles.Store(handler)
}
func (agent *CSSAgent) SetStartRuleUsageTrackingHandler(handler func(StartRuleUsageTrackingCommand)) {
    agent.commandHandlers.StartRuleUsageTracking.Store(handler)
}
func (agent *CSSAgent) SetStopRuleUsageTrackingHandler(handler func(StopRuleUsageTrackingCommand)) {
    agent.commandHandlers.StopRuleUsageTracking.Store(handler)
}
func init() {

}
