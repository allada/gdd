package css


type MediaQueryResultChangedEvent struct {
}
type FontsUpdatedEvent struct {
}
type StyleSheetChangedEvent struct {
    StyleSheetId StyleSheetId `json:"styleSheetId"`
}
type StyleSheetAddedEvent struct {
    Header CSSStyleSheetHeader `json:"header"`// Added stylesheet metainfo.
}
type StyleSheetRemovedEvent struct {
    StyleSheetId StyleSheetId `json:"styleSheetId"`// Identifier of the removed stylesheet.
}