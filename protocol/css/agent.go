package css


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type CSSAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        GetMatchedStylesForNode []chan<- GetMatchedStylesForNodeCommand
        GetInlineStylesForNode []chan<- GetInlineStylesForNodeCommand
        GetComputedStyleForNode []chan<- GetComputedStyleForNodeCommand
        GetPlatformFontsForNode []chan<- GetPlatformFontsForNodeCommand
        GetStyleSheetText []chan<- GetStyleSheetTextCommand
        CollectClassNames []chan<- CollectClassNamesCommand
        SetStyleSheetText []chan<- SetStyleSheetTextCommand
        SetRuleSelector []chan<- SetRuleSelectorCommand
        SetKeyframeKey []chan<- SetKeyframeKeyCommand
        SetStyleTexts []chan<- SetStyleTextsCommand
        SetMediaText []chan<- SetMediaTextCommand
        CreateStyleSheet []chan<- CreateStyleSheetCommand
        AddRule []chan<- AddRuleCommand
        ForcePseudoState []chan<- ForcePseudoStateCommand
        GetMediaQueries []chan<- GetMediaQueriesCommand
        SetEffectivePropertyValueForNode []chan<- SetEffectivePropertyValueForNodeCommand
        GetBackgroundColors []chan<- GetBackgroundColorsCommand
        GetLayoutTreeAndStyles []chan<- GetLayoutTreeAndStylesCommand
        StartRuleUsageTracking []chan<- StartRuleUsageTrackingCommand
        StopRuleUsageTracking []chan<- StopRuleUsageTrackingCommand
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
        case "getMatchedStylesForNode":
            var out GetMatchedStylesForNodeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetMatchedStylesForNode {
                c <-out
            }
        case "getInlineStylesForNode":
            var out GetInlineStylesForNodeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetInlineStylesForNode {
                c <-out
            }
        case "getComputedStyleForNode":
            var out GetComputedStyleForNodeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetComputedStyleForNode {
                c <-out
            }
        case "getPlatformFontsForNode":
            var out GetPlatformFontsForNodeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetPlatformFontsForNode {
                c <-out
            }
        case "getStyleSheetText":
            var out GetStyleSheetTextCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetStyleSheetText {
                c <-out
            }
        case "collectClassNames":
            var out CollectClassNamesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CollectClassNames {
                c <-out
            }
        case "setStyleSheetText":
            var out SetStyleSheetTextCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetStyleSheetText {
                c <-out
            }
        case "setRuleSelector":
            var out SetRuleSelectorCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetRuleSelector {
                c <-out
            }
        case "setKeyframeKey":
            var out SetKeyframeKeyCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetKeyframeKey {
                c <-out
            }
        case "setStyleTexts":
            var out SetStyleTextsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetStyleTexts {
                c <-out
            }
        case "setMediaText":
            var out SetMediaTextCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetMediaText {
                c <-out
            }
        case "createStyleSheet":
            var out CreateStyleSheetCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CreateStyleSheet {
                c <-out
            }
        case "addRule":
            var out AddRuleCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.AddRule {
                c <-out
            }
        case "forcePseudoState":
            var out ForcePseudoStateCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ForcePseudoState {
                c <-out
            }
        case "getMediaQueries":
            var out GetMediaQueriesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetMediaQueries {
                c <-out
            }
        case "setEffectivePropertyValueForNode":
            var out SetEffectivePropertyValueForNodeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetEffectivePropertyValueForNode {
                c <-out
            }
        case "getBackgroundColors":
            var out GetBackgroundColorsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetBackgroundColors {
                c <-out
            }
        case "getLayoutTreeAndStyles":
            var out GetLayoutTreeAndStylesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetLayoutTreeAndStyles {
                c <-out
            }
        case "startRuleUsageTracking":
            var out StartRuleUsageTrackingCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StartRuleUsageTracking {
                c <-out
            }
        case "stopRuleUsageTracking":
            var out StopRuleUsageTrackingCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StopRuleUsageTracking {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *CSSAgent) FireMediaQueryResultChanged(event MediaQueryResultChangedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "CSS.mediaQueryResultChanged",
        Params: event,
    })
}
func (agent *CSSAgent) FireMediaQueryResultChangedOnTarget(targetId string, event MediaQueryResultChangedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "CSS.mediaQueryResultChanged",
        Params: event,
    })
}
func (agent *CSSAgent) FireFontsUpdated(event FontsUpdatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "CSS.fontsUpdated",
        Params: event,
    })
}
func (agent *CSSAgent) FireFontsUpdatedOnTarget(targetId string, event FontsUpdatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "CSS.fontsUpdated",
        Params: event,
    })
}
func (agent *CSSAgent) FireStyleSheetChanged(event StyleSheetChangedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "CSS.styleSheetChanged",
        Params: event,
    })
}
func (agent *CSSAgent) FireStyleSheetChangedOnTarget(targetId string, event StyleSheetChangedEvent) {
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
func (agent *CSSAgent) FireStyleSheetAddedOnTarget(targetId string, event StyleSheetAddedEvent) {
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
func (agent *CSSAgent) FireStyleSheetRemovedOnTarget(targetId string, event StyleSheetRemovedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "CSS.styleSheetRemoved",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *CSSAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *CSSAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *CSSAgent) GetMatchedStylesForNodeNotify() <-chan GetMatchedStylesForNodeCommand {
    outChan := make(chan GetMatchedStylesForNodeCommand)
    agent.commandChans.GetMatchedStylesForNode = append(agent.commandChans.GetMatchedStylesForNode, outChan)
    return outChan
}
func (agent *CSSAgent) GetInlineStylesForNodeNotify() <-chan GetInlineStylesForNodeCommand {
    outChan := make(chan GetInlineStylesForNodeCommand)
    agent.commandChans.GetInlineStylesForNode = append(agent.commandChans.GetInlineStylesForNode, outChan)
    return outChan
}
func (agent *CSSAgent) GetComputedStyleForNodeNotify() <-chan GetComputedStyleForNodeCommand {
    outChan := make(chan GetComputedStyleForNodeCommand)
    agent.commandChans.GetComputedStyleForNode = append(agent.commandChans.GetComputedStyleForNode, outChan)
    return outChan
}
func (agent *CSSAgent) GetPlatformFontsForNodeNotify() <-chan GetPlatformFontsForNodeCommand {
    outChan := make(chan GetPlatformFontsForNodeCommand)
    agent.commandChans.GetPlatformFontsForNode = append(agent.commandChans.GetPlatformFontsForNode, outChan)
    return outChan
}
func (agent *CSSAgent) GetStyleSheetTextNotify() <-chan GetStyleSheetTextCommand {
    outChan := make(chan GetStyleSheetTextCommand)
    agent.commandChans.GetStyleSheetText = append(agent.commandChans.GetStyleSheetText, outChan)
    return outChan
}
func (agent *CSSAgent) CollectClassNamesNotify() <-chan CollectClassNamesCommand {
    outChan := make(chan CollectClassNamesCommand)
    agent.commandChans.CollectClassNames = append(agent.commandChans.CollectClassNames, outChan)
    return outChan
}
func (agent *CSSAgent) SetStyleSheetTextNotify() <-chan SetStyleSheetTextCommand {
    outChan := make(chan SetStyleSheetTextCommand)
    agent.commandChans.SetStyleSheetText = append(agent.commandChans.SetStyleSheetText, outChan)
    return outChan
}
func (agent *CSSAgent) SetRuleSelectorNotify() <-chan SetRuleSelectorCommand {
    outChan := make(chan SetRuleSelectorCommand)
    agent.commandChans.SetRuleSelector = append(agent.commandChans.SetRuleSelector, outChan)
    return outChan
}
func (agent *CSSAgent) SetKeyframeKeyNotify() <-chan SetKeyframeKeyCommand {
    outChan := make(chan SetKeyframeKeyCommand)
    agent.commandChans.SetKeyframeKey = append(agent.commandChans.SetKeyframeKey, outChan)
    return outChan
}
func (agent *CSSAgent) SetStyleTextsNotify() <-chan SetStyleTextsCommand {
    outChan := make(chan SetStyleTextsCommand)
    agent.commandChans.SetStyleTexts = append(agent.commandChans.SetStyleTexts, outChan)
    return outChan
}
func (agent *CSSAgent) SetMediaTextNotify() <-chan SetMediaTextCommand {
    outChan := make(chan SetMediaTextCommand)
    agent.commandChans.SetMediaText = append(agent.commandChans.SetMediaText, outChan)
    return outChan
}
func (agent *CSSAgent) CreateStyleSheetNotify() <-chan CreateStyleSheetCommand {
    outChan := make(chan CreateStyleSheetCommand)
    agent.commandChans.CreateStyleSheet = append(agent.commandChans.CreateStyleSheet, outChan)
    return outChan
}
func (agent *CSSAgent) AddRuleNotify() <-chan AddRuleCommand {
    outChan := make(chan AddRuleCommand)
    agent.commandChans.AddRule = append(agent.commandChans.AddRule, outChan)
    return outChan
}
func (agent *CSSAgent) ForcePseudoStateNotify() <-chan ForcePseudoStateCommand {
    outChan := make(chan ForcePseudoStateCommand)
    agent.commandChans.ForcePseudoState = append(agent.commandChans.ForcePseudoState, outChan)
    return outChan
}
func (agent *CSSAgent) GetMediaQueriesNotify() <-chan GetMediaQueriesCommand {
    outChan := make(chan GetMediaQueriesCommand)
    agent.commandChans.GetMediaQueries = append(agent.commandChans.GetMediaQueries, outChan)
    return outChan
}
func (agent *CSSAgent) SetEffectivePropertyValueForNodeNotify() <-chan SetEffectivePropertyValueForNodeCommand {
    outChan := make(chan SetEffectivePropertyValueForNodeCommand)
    agent.commandChans.SetEffectivePropertyValueForNode = append(agent.commandChans.SetEffectivePropertyValueForNode, outChan)
    return outChan
}
func (agent *CSSAgent) GetBackgroundColorsNotify() <-chan GetBackgroundColorsCommand {
    outChan := make(chan GetBackgroundColorsCommand)
    agent.commandChans.GetBackgroundColors = append(agent.commandChans.GetBackgroundColors, outChan)
    return outChan
}
func (agent *CSSAgent) GetLayoutTreeAndStylesNotify() <-chan GetLayoutTreeAndStylesCommand {
    outChan := make(chan GetLayoutTreeAndStylesCommand)
    agent.commandChans.GetLayoutTreeAndStyles = append(agent.commandChans.GetLayoutTreeAndStyles, outChan)
    return outChan
}
func (agent *CSSAgent) StartRuleUsageTrackingNotify() <-chan StartRuleUsageTrackingCommand {
    outChan := make(chan StartRuleUsageTrackingCommand)
    agent.commandChans.StartRuleUsageTracking = append(agent.commandChans.StartRuleUsageTracking, outChan)
    return outChan
}
func (agent *CSSAgent) StopRuleUsageTrackingNotify() <-chan StopRuleUsageTrackingCommand {
    outChan := make(chan StopRuleUsageTrackingCommand)
    agent.commandChans.StopRuleUsageTracking = append(agent.commandChans.StopRuleUsageTracking, outChan)
    return outChan
}
func init() {

}
