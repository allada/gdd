package css


import (
    "../shared"
    "../dom"
    "../page"
)

type EnableCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type EnableReturn struct {
}

func (c *EnableCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *EnableCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *EnableCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type DisableCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type DisableReturn struct {
}

func (c *DisableCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *DisableCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DisableCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetMatchedStylesForNodeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId dom.NodeId `json:"nodeId"`
}
type GetMatchedStylesForNodeReturn struct {
    InlineStyle *CSSStyle `json:"inlineStyle,omitempty"`// Inline style for the specified DOM node.
    AttributesStyle *CSSStyle `json:"attributesStyle,omitempty"`// Attribute-defined element style (e.g. resulting from "width=20 height=100%").
    MatchedCSSRules *[]RuleMatch `json:"matchedCSSRules,omitempty"`// CSS rules matching this node, from all applicable stylesheets.
    PseudoElements *[]PseudoElementMatches `json:"pseudoElements,omitempty"`// Pseudo style matches for this node.
    Inherited *[]InheritedStyleEntry `json:"inherited,omitempty"`// A chain of inherited styles (from the immediate node parent up to the DOM tree root).
    CssKeyframesRules *[]CSSKeyframesRule `json:"cssKeyframesRules,omitempty"`// A list of CSS keyframed animations matching this node.
}

func (c *GetMatchedStylesForNodeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetMatchedStylesForNodeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetMatchedStylesForNodeCommand) Respond(r *GetMatchedStylesForNodeReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetInlineStylesForNodeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId dom.NodeId `json:"nodeId"`
}
type GetInlineStylesForNodeReturn struct {
    InlineStyle *CSSStyle `json:"inlineStyle,omitempty"`// Inline style for the specified DOM node.
    AttributesStyle *CSSStyle `json:"attributesStyle,omitempty"`// Attribute-defined element style (e.g. resulting from "width=20 height=100%").
}

func (c *GetInlineStylesForNodeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetInlineStylesForNodeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetInlineStylesForNodeCommand) Respond(r *GetInlineStylesForNodeReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetComputedStyleForNodeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId dom.NodeId `json:"nodeId"`
}
type GetComputedStyleForNodeReturn struct {
    ComputedStyle []CSSComputedStyleProperty `json:"computedStyle"`// Computed style for the specified DOM node.
}

func (c *GetComputedStyleForNodeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetComputedStyleForNodeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetComputedStyleForNodeCommand) Respond(r *GetComputedStyleForNodeReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetPlatformFontsForNodeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId dom.NodeId `json:"nodeId"`
}
type GetPlatformFontsForNodeReturn struct {
    Fonts []PlatformFontUsage `json:"fonts"`// Usage statistics for every employed platform font.
}

func (c *GetPlatformFontsForNodeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetPlatformFontsForNodeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetPlatformFontsForNodeCommand) Respond(r *GetPlatformFontsForNodeReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetStyleSheetTextCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    StyleSheetId StyleSheetId `json:"styleSheetId"`
}
type GetStyleSheetTextReturn struct {
    Text string `json:"text"`// The stylesheet text.
}

func (c *GetStyleSheetTextCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetStyleSheetTextCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetStyleSheetTextCommand) Respond(r *GetStyleSheetTextReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type CollectClassNamesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    StyleSheetId StyleSheetId `json:"styleSheetId"`
}
type CollectClassNamesReturn struct {
    ClassNames []string `json:"classNames"`// Class name list.
}

func (c *CollectClassNamesCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *CollectClassNamesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CollectClassNamesCommand) Respond(r *CollectClassNamesReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetStyleSheetTextCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    StyleSheetId StyleSheetId `json:"styleSheetId"`
    Text string `json:"text"`
}
type SetStyleSheetTextReturn struct {
    SourceMapURL *string `json:"sourceMapURL,omitempty"`// URL of source map associated with script (if any).
}

func (c *SetStyleSheetTextCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetStyleSheetTextCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetStyleSheetTextCommand) Respond(r *SetStyleSheetTextReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetRuleSelectorCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    StyleSheetId StyleSheetId `json:"styleSheetId"`
    Range SourceRange `json:"range"`
    Selector string `json:"selector"`
}
type SetRuleSelectorReturn struct {
    SelectorList SelectorList `json:"selectorList"`// The resulting selector list after modification.
}

func (c *SetRuleSelectorCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetRuleSelectorCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetRuleSelectorCommand) Respond(r *SetRuleSelectorReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetKeyframeKeyCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    StyleSheetId StyleSheetId `json:"styleSheetId"`
    Range SourceRange `json:"range"`
    KeyText string `json:"keyText"`
}
type SetKeyframeKeyReturn struct {
    KeyText Value `json:"keyText"`// The resulting key text after modification.
}

func (c *SetKeyframeKeyCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetKeyframeKeyCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetKeyframeKeyCommand) Respond(r *SetKeyframeKeyReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetStyleTextsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Edits []StyleDeclarationEdit `json:"edits"`
}
type SetStyleTextsReturn struct {
    Styles []CSSStyle `json:"styles"`// The resulting styles after modification.
}

func (c *SetStyleTextsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetStyleTextsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetStyleTextsCommand) Respond(r *SetStyleTextsReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetMediaTextCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    StyleSheetId StyleSheetId `json:"styleSheetId"`
    Range SourceRange `json:"range"`
    Text string `json:"text"`
}
type SetMediaTextReturn struct {
    Media CSSMedia `json:"media"`// The resulting CSS media rule after modification.
}

func (c *SetMediaTextCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetMediaTextCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetMediaTextCommand) Respond(r *SetMediaTextReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type CreateStyleSheetCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    FrameId page.FrameId `json:"frameId"`// Identifier of the frame where "via-inspector" stylesheet should be created.
}
type CreateStyleSheetReturn struct {
    StyleSheetId StyleSheetId `json:"styleSheetId"`// Identifier of the created "via-inspector" stylesheet.
}

func (c *CreateStyleSheetCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *CreateStyleSheetCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CreateStyleSheetCommand) Respond(r *CreateStyleSheetReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type AddRuleCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    StyleSheetId StyleSheetId `json:"styleSheetId"`// The css style sheet identifier where a new rule should be inserted.
    RuleText string `json:"ruleText"`// The text of a new rule.
    Location SourceRange `json:"location"`// Text position of a new rule in the target style sheet.
}
type AddRuleReturn struct {
    Rule CSSRule `json:"rule"`// The newly created rule.
}

func (c *AddRuleCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *AddRuleCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *AddRuleCommand) Respond(r *AddRuleReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type ForcePseudoStateCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId dom.NodeId `json:"nodeId"`// The element id for which to force the pseudo state.
    ForcedPseudoClasses []ForcePseudoStateForcedPseudoClassesEnum `json:"forcedPseudoClasses"`// Element pseudo classes to force when computing the element's style.
}
type ForcePseudoStateReturn struct {
}

func (c *ForcePseudoStateCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ForcePseudoStateCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ForcePseudoStateCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetMediaQueriesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type GetMediaQueriesReturn struct {
    Medias []CSSMedia `json:"medias"`
}

func (c *GetMediaQueriesCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetMediaQueriesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetMediaQueriesCommand) Respond(r *GetMediaQueriesReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetEffectivePropertyValueForNodeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId dom.NodeId `json:"nodeId"`// The element id for which to set property.
    PropertyName string `json:"propertyName"`
    Value string `json:"value"`
}
type SetEffectivePropertyValueForNodeReturn struct {
}

func (c *SetEffectivePropertyValueForNodeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetEffectivePropertyValueForNodeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetEffectivePropertyValueForNodeCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetBackgroundColorsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId dom.NodeId `json:"nodeId"`// Id of the node to get background colors for.
}
type GetBackgroundColorsReturn struct {
    BackgroundColors *[]string `json:"backgroundColors,omitempty"`// The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load).
}

func (c *GetBackgroundColorsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetBackgroundColorsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetBackgroundColorsCommand) Respond(r *GetBackgroundColorsReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetLayoutTreeAndStylesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ComputedStyleWhitelist []string `json:"computedStyleWhitelist"`// Whitelist of computed styles to return.
}
type GetLayoutTreeAndStylesReturn struct {
    LayoutTreeNodes []LayoutTreeNode `json:"layoutTreeNodes"`
    ComputedStyles []ComputedStyle `json:"computedStyles"`
}

func (c *GetLayoutTreeAndStylesCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetLayoutTreeAndStylesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetLayoutTreeAndStylesCommand) Respond(r *GetLayoutTreeAndStylesReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type StartRuleUsageTrackingCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type StartRuleUsageTrackingReturn struct {
}

func (c *StartRuleUsageTrackingCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *StartRuleUsageTrackingCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StartRuleUsageTrackingCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type StopRuleUsageTrackingCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type StopRuleUsageTrackingReturn struct {
    RuleUsage []RuleUsage `json:"ruleUsage"`
}

func (c *StopRuleUsageTrackingCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *StopRuleUsageTrackingCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StopRuleUsageTrackingCommand) Respond(r *StopRuleUsageTrackingReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
