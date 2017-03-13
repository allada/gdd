package css


import (
    "../dom"
    "../page"
)
const (
    StyleSheetOriginInjected StyleSheetOrigin = "injected"
    StyleSheetOriginUserDashagent StyleSheetOrigin = "user-agent"
    StyleSheetOriginInspector StyleSheetOrigin = "inspector"
    StyleSheetOriginRegular StyleSheetOrigin = "regular"
)

type CSSMediaSourceEnum string
const (
    CSSMediaSourceMediaRule CSSMediaSourceEnum = "mediaRule"
    CSSMediaSourceImportRule CSSMediaSourceEnum = "importRule"
    CSSMediaSourceLinkedSheet CSSMediaSourceEnum = "linkedSheet"
    CSSMediaSourceInlineSheet CSSMediaSourceEnum = "inlineSheet"
)

type ForcePseudoStateForcedPseudoClassesEnum string
const (
    ForcePseudoStateForcedPseudoClassesActive ForcePseudoStateForcedPseudoClassesEnum = "active"
    ForcePseudoStateForcedPseudoClassesFocus ForcePseudoStateForcedPseudoClassesEnum = "focus"
    ForcePseudoStateForcedPseudoClassesHover ForcePseudoStateForcedPseudoClassesEnum = "hover"
    ForcePseudoStateForcedPseudoClassesVisited ForcePseudoStateForcedPseudoClassesEnum = "visited"
)
type StyleSheetId string

type StyleSheetOrigin string

type PseudoElementMatches struct {
    PseudoType dom.PseudoType `json:"pseudoType"`// Pseudo element type.
    Matches []RuleMatch `json:"matches"`// Matches of CSS rules applicable to the pseudo style.
}

type InheritedStyleEntry struct {
    InlineStyle *CSSStyle `json:"inlineStyle,omitempty"`// The ancestor node's inline style, if any, in the style inheritance chain.
    MatchedCSSRules []RuleMatch `json:"matchedCSSRules"`// Matches of CSS rules matching the ancestor node in the style inheritance chain.
}

type RuleMatch struct {
    Rule CSSRule `json:"rule"`// CSS rule in the match.
    MatchingSelectors []int64 `json:"matchingSelectors"`// Matching selector indices in the rule's selectorList selectors (0-based).
}

type Value struct {
    Text string `json:"text"`// Value text.
    Range *SourceRange `json:"range,omitempty"`// Value range in the underlying resource (if available).
}

type SelectorList struct {
    Selectors []Value `json:"selectors"`// Selectors in the list.
    Text string `json:"text"`// Rule selector text.
}

type CSSStyleSheetHeader struct {
    StyleSheetId StyleSheetId `json:"styleSheetId"`// The stylesheet identifier.
    FrameId page.FrameId `json:"frameId"`// Owner frame identifier.
    SourceURL string `json:"sourceURL"`// Stylesheet resource URL.
    SourceMapURL *string `json:"sourceMapURL,omitempty"`// URL of source map associated with the stylesheet (if any).
    Origin StyleSheetOrigin `json:"origin"`// Stylesheet origin.
    Title string `json:"title"`// Stylesheet title.
    OwnerNode *dom.BackendNodeId `json:"ownerNode,omitempty"`// The backend id for the owner node of the stylesheet.
    Disabled bool `json:"disabled"`// Denotes whether the stylesheet is disabled.
    HasSourceURL *bool `json:"hasSourceURL,omitempty"`// Whether the sourceURL field value comes from the sourceURL comment.
    IsInline bool `json:"isInline"`// Whether this stylesheet is created for STYLE tag by parser. This flag is not set for document.written STYLE tags.
    StartLine float64 `json:"startLine"`// Line offset of the stylesheet within the resource (zero based).
    StartColumn float64 `json:"startColumn"`// Column offset of the stylesheet within the resource (zero based).
}

type CSSRule struct {
    StyleSheetId *StyleSheetId `json:"styleSheetId,omitempty"`// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
    SelectorList SelectorList `json:"selectorList"`// Rule selector data.
    Origin StyleSheetOrigin `json:"origin"`// Parent stylesheet's origin.
    Style CSSStyle `json:"style"`// Associated style declaration.
    Media *[]CSSMedia `json:"media,omitempty"`// Media list array (for rules involving media queries). The array enumerates media queries starting with the innermost one, going outwards.
}

type RuleUsage struct {
    StyleSheetId StyleSheetId `json:"styleSheetId"`// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
    Range SourceRange `json:"range"`// Style declaration range in the enclosing stylesheet (if available).
    Used bool `json:"used"`// Indicates whether the rule was actually used by some element in the page.
}

type SourceRange struct {
    StartLine int64 `json:"startLine"`// Start line of range.
    StartColumn int64 `json:"startColumn"`// Start column of range (inclusive).
    EndLine int64 `json:"endLine"`// End line of range
    EndColumn int64 `json:"endColumn"`// End column of range (exclusive).
}

type ShorthandEntry struct {
    Name string `json:"name"`// Shorthand name.
    Value string `json:"value"`// Shorthand value.
    Important *bool `json:"important,omitempty"`// Whether the property has "!important" annotation (implies <code>false</code> if absent).
}

type CSSComputedStyleProperty struct {
    Name string `json:"name"`// Computed style property name.
    Value string `json:"value"`// Computed style property value.
}

type CSSStyle struct {
    StyleSheetId *StyleSheetId `json:"styleSheetId,omitempty"`// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
    CssProperties []CSSProperty `json:"cssProperties"`// CSS properties in the style.
    ShorthandEntries []ShorthandEntry `json:"shorthandEntries"`// Computed values for all shorthands found in the style.
    CssText *string `json:"cssText,omitempty"`// Style declaration text (if available).
    Range *SourceRange `json:"range,omitempty"`// Style declaration range in the enclosing stylesheet (if available).
}

type CSSProperty struct {
    Name string `json:"name"`// The property name.
    Value string `json:"value"`// The property value.
    Important *bool `json:"important,omitempty"`// Whether the property has "!important" annotation (implies <code>false</code> if absent).
    Implicit *bool `json:"implicit,omitempty"`// Whether the property is implicit (implies <code>false</code> if absent).
    Text *string `json:"text,omitempty"`// The full property text as specified in the style.
    ParsedOk *bool `json:"parsedOk,omitempty"`// Whether the property is understood by the browser (implies <code>true</code> if absent).
    Disabled *bool `json:"disabled,omitempty"`// Whether the property is disabled by the user (present for source-based properties only).
    Range *SourceRange `json:"range,omitempty"`// The entire property range in the enclosing style declaration (if available).
}

type CSSMedia struct {
    Text string `json:"text"`// Media query text.
    Source CSSMediaSourceEnum `json:"source"`// Source of the media query: "mediaRule" if specified by a @media rule, "importRule" if specified by an @import rule, "linkedSheet" if specified by a "media" attribute in a linked stylesheet's LINK tag, "inlineSheet" if specified by a "media" attribute in an inline stylesheet's STYLE tag.
    SourceURL *string `json:"sourceURL,omitempty"`// URL of the document containing the media query description.
    Range *SourceRange `json:"range,omitempty"`// The associated rule (@media or @import) header range in the enclosing stylesheet (if available).
    StyleSheetId *StyleSheetId `json:"styleSheetId,omitempty"`// Identifier of the stylesheet containing this object (if exists).
    MediaList *[]MediaQuery `json:"mediaList,omitempty"`// [Experimental] Array of media queries.
}

type MediaQuery struct {
    Expressions []MediaQueryExpression `json:"expressions"`// Array of media query expressions.
    Active bool `json:"active"`// Whether the media query condition is satisfied.
}

type MediaQueryExpression struct {
    Value float64 `json:"value"`// Media query expression value.
    Unit string `json:"unit"`// Media query expression units.
    Feature string `json:"feature"`// Media query expression feature.
    ValueRange *SourceRange `json:"valueRange,omitempty"`// The associated range of the value text in the enclosing stylesheet (if available).
    ComputedLength *float64 `json:"computedLength,omitempty"`// Computed length of media query expression (if applicable).
}

type PlatformFontUsage struct {
    FamilyName string `json:"familyName"`// Font's family name reported by platform.
    IsCustomFont bool `json:"isCustomFont"`// Indicates if the font was downloaded or resolved locally.
    GlyphCount float64 `json:"glyphCount"`// Amount of glyphs that were rendered with this font.
}

type CSSKeyframesRule struct {
    AnimationName Value `json:"animationName"`// Animation name.
    Keyframes []CSSKeyframeRule `json:"keyframes"`// List of keyframes.
}

type CSSKeyframeRule struct {
    StyleSheetId *StyleSheetId `json:"styleSheetId,omitempty"`// The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
    Origin StyleSheetOrigin `json:"origin"`// Parent stylesheet's origin.
    KeyText Value `json:"keyText"`// Associated key text.
    Style CSSStyle `json:"style"`// Associated style declaration.
}

type StyleDeclarationEdit struct {
    StyleSheetId StyleSheetId `json:"styleSheetId"`// The css style sheet identifier.
    Range SourceRange `json:"range"`// The range of the style text in the enclosing stylesheet.
    Text string `json:"text"`// New style text.
}

type InlineTextBox struct {
    BoundingBox dom.Rect `json:"boundingBox"`// The absolute position bounding box.
    StartCharacterIndex int64 `json:"startCharacterIndex"`// The starting index in characters, for this post layout textbox substring.
    NumCharacters int64 `json:"numCharacters"`// The number of characters in this post layout textbox substring.
}

type LayoutTreeNode struct {
    NodeId dom.NodeId `json:"nodeId"`// The id of the related DOM node matching one from DOM.GetDocument.
    BoundingBox dom.Rect `json:"boundingBox"`// The absolute position bounding box.
    LayoutText *string `json:"layoutText,omitempty"`// Contents of the LayoutText if any
    InlineTextNodes *[]InlineTextBox `json:"inlineTextNodes,omitempty"`// The post layout inline text nodes, if any.
    StyleIndex *int64 `json:"styleIndex,omitempty"`// Index into the computedStyles array returned by getLayoutTreeAndStyles.
}

type ComputedStyle struct {
    Properties []CSSComputedStyleProperty `json:"properties"`
}
