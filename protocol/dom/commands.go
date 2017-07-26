package dom


import (
    "github.com/allada/gdd/protocol/shared"
    "github.com/allada/gdd/protocol/runtime"
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

type GetDocumentCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Depth *int64 `json:"depth,omitempty"`// [Experimental] The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
    Pierce *bool `json:"pierce,omitempty"`// [Experimental] Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
}
type GetDocumentReturn struct {
    Root Node `json:"root"`// Resulting node.
}

func (c *GetDocumentCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetDocumentCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetDocumentCommand) Respond(r *GetDocumentReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetFlattenedDocumentCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Depth *int64 `json:"depth,omitempty"`// [Experimental] The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
    Pierce *bool `json:"pierce,omitempty"`// [Experimental] Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
}
type GetFlattenedDocumentReturn struct {
    Nodes []Node `json:"nodes"`// Resulting node.
}

func (c *GetFlattenedDocumentCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetFlattenedDocumentCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetFlattenedDocumentCommand) Respond(r *GetFlattenedDocumentReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type CollectClassNamesFromSubtreeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to collect class names.
}
type CollectClassNamesFromSubtreeReturn struct {
    ClassNames []string `json:"classNames"`// Class name list.
}

func (c *CollectClassNamesFromSubtreeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *CollectClassNamesFromSubtreeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CollectClassNamesFromSubtreeCommand) Respond(r *CollectClassNamesFromSubtreeReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type RequestChildNodesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to get children for.
    Depth *int64 `json:"depth,omitempty"`// [Experimental] The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
    Pierce *bool `json:"pierce,omitempty"`// [Experimental] Whether or not iframes and shadow roots should be traversed when returning the sub-tree (default is false).
}
type RequestChildNodesReturn struct {
}

func (c *RequestChildNodesCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RequestChildNodesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RequestChildNodesCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type QuerySelectorCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to query upon.
    Selector string `json:"selector"`// Selector string.
}
type QuerySelectorReturn struct {
    NodeId NodeId `json:"nodeId"`// Query selector result.
}

func (c *QuerySelectorCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *QuerySelectorCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *QuerySelectorCommand) Respond(r *QuerySelectorReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type QuerySelectorAllCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to query upon.
    Selector string `json:"selector"`// Selector string.
}
type QuerySelectorAllReturn struct {
    NodeIds []NodeId `json:"nodeIds"`// Query selector result.
}

func (c *QuerySelectorAllCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *QuerySelectorAllCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *QuerySelectorAllCommand) Respond(r *QuerySelectorAllReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetNodeNameCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to set name for.
    Name string `json:"name"`// New node's name.
}
type SetNodeNameReturn struct {
    NodeId NodeId `json:"nodeId"`// New node's id.
}

func (c *SetNodeNameCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetNodeNameCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetNodeNameCommand) Respond(r *SetNodeNameReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetNodeValueCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to set value for.
    Value string `json:"value"`// New node's value.
}
type SetNodeValueReturn struct {
}

func (c *SetNodeValueCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetNodeValueCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetNodeValueCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type RemoveNodeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to remove.
}
type RemoveNodeReturn struct {
}

func (c *RemoveNodeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RemoveNodeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RemoveNodeCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetAttributeValueCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the element to set attribute for.
    Name string `json:"name"`// Attribute name.
    Value string `json:"value"`// Attribute value.
}
type SetAttributeValueReturn struct {
}

func (c *SetAttributeValueCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetAttributeValueCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetAttributeValueCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetAttributesAsTextCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the element to set attributes for.
    Text string `json:"text"`// Text with a number of attributes. Will parse this text using HTML parser.
    Name *string `json:"name,omitempty"`// Attribute name to replace with new attributes derived from text in case text parsed successfully.
}
type SetAttributesAsTextReturn struct {
}

func (c *SetAttributesAsTextCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetAttributesAsTextCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetAttributesAsTextCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type RemoveAttributeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the element to remove attribute from.
    Name string `json:"name"`// Name of the attribute to remove.
}
type RemoveAttributeReturn struct {
}

func (c *RemoveAttributeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RemoveAttributeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RemoveAttributeCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetOuterHTMLCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to get markup for.
}
type GetOuterHTMLReturn struct {
    OuterHTML string `json:"outerHTML"`// Outer HTML markup.
}

func (c *GetOuterHTMLCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetOuterHTMLCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetOuterHTMLCommand) Respond(r *GetOuterHTMLReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetOuterHTMLCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to set markup for.
    OuterHTML string `json:"outerHTML"`// Outer HTML markup to set.
}
type SetOuterHTMLReturn struct {
}

func (c *SetOuterHTMLCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetOuterHTMLCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetOuterHTMLCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type PerformSearchCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Query string `json:"query"`// Plain text or query selector or XPath search query.
    IncludeUserAgentShadowDOM *bool `json:"includeUserAgentShadowDOM,omitempty"`// [Experimental] True to search in user agent shadow DOM.
}
type PerformSearchReturn struct {
    SearchId string `json:"searchId"`// Unique search session identifier.
    ResultCount int64 `json:"resultCount"`// Number of search results.
}

func (c *PerformSearchCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *PerformSearchCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *PerformSearchCommand) Respond(r *PerformSearchReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetSearchResultsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    SearchId string `json:"searchId"`// Unique search session identifier.
    FromIndex int64 `json:"fromIndex"`// Start index of the search result to be returned.
    ToIndex int64 `json:"toIndex"`// End index of the search result to be returned.
}
type GetSearchResultsReturn struct {
    NodeIds []NodeId `json:"nodeIds"`// Ids of the search result nodes.
}

func (c *GetSearchResultsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetSearchResultsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetSearchResultsCommand) Respond(r *GetSearchResultsReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type DiscardSearchResultsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    SearchId string `json:"searchId"`// Unique search session identifier.
}
type DiscardSearchResultsReturn struct {
}

func (c *DiscardSearchResultsCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *DiscardSearchResultsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DiscardSearchResultsCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type RequestNodeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ObjectId runtime.RemoteObjectId `json:"objectId"`// JavaScript object id to convert into node.
}
type RequestNodeReturn struct {
    NodeId NodeId `json:"nodeId"`// Node id for given object.
}

func (c *RequestNodeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RequestNodeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RequestNodeCommand) Respond(r *RequestNodeReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetInspectModeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Mode InspectMode `json:"mode"`// Set an inspection mode.
    HighlightConfig *HighlightConfig `json:"highlightConfig,omitempty"`// A descriptor for the highlight appearance of hovered-over nodes. May be omitted if <code>enabled == false</code>.
}
type SetInspectModeReturn struct {
}

func (c *SetInspectModeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetInspectModeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetInspectModeCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type HighlightRectCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    X int64 `json:"x"`// X coordinate
    Y int64 `json:"y"`// Y coordinate
    Width int64 `json:"width"`// Rectangle width
    Height int64 `json:"height"`// Rectangle height
    Color *RGBA `json:"color,omitempty"`// The highlight fill color (default: transparent).
    OutlineColor *RGBA `json:"outlineColor,omitempty"`// The highlight outline color (default: transparent).
}
type HighlightRectReturn struct {
}

func (c *HighlightRectCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *HighlightRectCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *HighlightRectCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type HighlightQuadCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Quad Quad `json:"quad"`// Quad to highlight
    Color *RGBA `json:"color,omitempty"`// The highlight fill color (default: transparent).
    OutlineColor *RGBA `json:"outlineColor,omitempty"`// The highlight outline color (default: transparent).
}
type HighlightQuadReturn struct {
}

func (c *HighlightQuadCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *HighlightQuadCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *HighlightQuadCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type HighlightNodeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    HighlightConfig HighlightConfig `json:"highlightConfig"`// A descriptor for the highlight appearance.
    NodeId *NodeId `json:"nodeId,omitempty"`// Identifier of the node to highlight.
    BackendNodeId *BackendNodeId `json:"backendNodeId,omitempty"`// Identifier of the backend node to highlight.
    ObjectId *runtime.RemoteObjectId `json:"objectId,omitempty"`// [Experimental] JavaScript object id of the node to be highlighted.
}
type HighlightNodeReturn struct {
}

func (c *HighlightNodeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *HighlightNodeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *HighlightNodeCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type HideHighlightCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type HideHighlightReturn struct {
}

func (c *HideHighlightCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *HideHighlightCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *HideHighlightCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type HighlightFrameCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    FrameId string `json:"frameId"`// Identifier of the frame to highlight.
    ContentColor *RGBA `json:"contentColor,omitempty"`// The content box highlight fill color (default: transparent).
    ContentOutlineColor *RGBA `json:"contentOutlineColor,omitempty"`// The content box highlight outline color (default: transparent).
}
type HighlightFrameReturn struct {
}

func (c *HighlightFrameCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *HighlightFrameCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *HighlightFrameCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type PushNodeByPathToFrontendCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Path string `json:"path"`// Path to node in the proprietary format.
}
type PushNodeByPathToFrontendReturn struct {
    NodeId NodeId `json:"nodeId"`// Id of the node for given path.
}

func (c *PushNodeByPathToFrontendCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *PushNodeByPathToFrontendCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *PushNodeByPathToFrontendCommand) Respond(r *PushNodeByPathToFrontendReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type PushNodesByBackendIdsToFrontendCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    BackendNodeIds []BackendNodeId `json:"backendNodeIds"`// The array of backend node ids.
}
type PushNodesByBackendIdsToFrontendReturn struct {
    NodeIds []NodeId `json:"nodeIds"`// The array of ids of pushed nodes that correspond to the backend ids specified in backendNodeIds.
}

func (c *PushNodesByBackendIdsToFrontendCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *PushNodesByBackendIdsToFrontendCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *PushNodesByBackendIdsToFrontendCommand) Respond(r *PushNodesByBackendIdsToFrontendReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type SetInspectedNodeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// DOM node id to be accessible by means of $x command line API.
}
type SetInspectedNodeReturn struct {
}

func (c *SetInspectedNodeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetInspectedNodeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetInspectedNodeCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type ResolveNodeCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to resolve.
    ObjectGroup *string `json:"objectGroup,omitempty"`// Symbolic group name that can be used to release multiple objects.
}
type ResolveNodeReturn struct {
    Object runtime.RemoteObject `json:"object"`// JavaScript object wrapper for given node.
}

func (c *ResolveNodeCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *ResolveNodeCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *ResolveNodeCommand) Respond(r *ResolveNodeReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetAttributesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to retrieve attibutes for.
}
type GetAttributesReturn struct {
    Attributes []string `json:"attributes"`// An interleaved array of node attribute names and values.
}

func (c *GetAttributesCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetAttributesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetAttributesCommand) Respond(r *GetAttributesReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type CopyToCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to copy.
    TargetNodeId NodeId `json:"targetNodeId"`// Id of the element to drop the copy into.
    InsertBeforeNodeId *NodeId `json:"insertBeforeNodeId,omitempty"`// Drop the copy before this node (if absent, the copy becomes the last child of <code>targetNodeId</code>).
}
type CopyToReturn struct {
    NodeId NodeId `json:"nodeId"`// Id of the node clone.
}

func (c *CopyToCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *CopyToCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CopyToCommand) Respond(r *CopyToReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type MoveToCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to move.
    TargetNodeId NodeId `json:"targetNodeId"`// Id of the element to drop the moved node into.
    InsertBeforeNodeId *NodeId `json:"insertBeforeNodeId,omitempty"`// Drop node before this one (if absent, the moved node becomes the last child of <code>targetNodeId</code>).
}
type MoveToReturn struct {
    NodeId NodeId `json:"nodeId"`// New id of the moved node.
}

func (c *MoveToCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *MoveToCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *MoveToCommand) Respond(r *MoveToReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type UndoCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type UndoReturn struct {
}

func (c *UndoCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *UndoCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *UndoCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type RedoCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type RedoReturn struct {
}

func (c *RedoCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *RedoCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *RedoCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type MarkUndoableStateCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type MarkUndoableStateReturn struct {
}

func (c *MarkUndoableStateCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *MarkUndoableStateCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *MarkUndoableStateCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type FocusCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to focus.
}
type FocusReturn struct {
}

func (c *FocusCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *FocusCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *FocusCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetFileInputFilesCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the file input node to set files for.
    Files []string `json:"files"`// Array of file paths to set.
}
type SetFileInputFilesReturn struct {
}

func (c *SetFileInputFilesCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *SetFileInputFilesCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetFileInputFilesCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetBoxModelCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to get box model for.
}
type GetBoxModelReturn struct {
    Model BoxModel `json:"model"`// Box model for the node.
}

func (c *GetBoxModelCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetBoxModelCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetBoxModelCommand) Respond(r *GetBoxModelReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetNodeForLocationCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    X int64 `json:"x"`// X coordinate.
    Y int64 `json:"y"`// Y coordinate.
}
type GetNodeForLocationReturn struct {
    NodeId NodeId `json:"nodeId"`// Id of the node at given coordinates.
}

func (c *GetNodeForLocationCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetNodeForLocationCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetNodeForLocationCommand) Respond(r *GetNodeForLocationReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetRelayoutBoundaryCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node.
}
type GetRelayoutBoundaryReturn struct {
    NodeId NodeId `json:"nodeId"`// Relayout boundary node id for the given node.
}

func (c *GetRelayoutBoundaryCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetRelayoutBoundaryCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetRelayoutBoundaryCommand) Respond(r *GetRelayoutBoundaryReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type GetHighlightObjectForTestCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    NodeId NodeId `json:"nodeId"`// Id of the node to get highlight object for.
}
type GetHighlightObjectForTestReturn struct {
    Highlight map[string]string `json:"highlight"`// Highlight data for the node.
}

func (c *GetHighlightObjectForTestCommand) Initalize(targetId string, responseId int64, conn *shared.Connection) {
    c.DestinationTargetID = targetId
    c.responseId = responseId
    c.conn = conn
}

func (c *GetHighlightObjectForTestCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetHighlightObjectForTestCommand) Respond(r *GetHighlightObjectForTestReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
