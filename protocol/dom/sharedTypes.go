package dom


import (
    "../page"
)
const (
    PseudoTypeFirstDashline PseudoType = "first-line"
    PseudoTypeFirstDashletter PseudoType = "first-letter"
    PseudoTypeBefore PseudoType = "before"
    PseudoTypeAfter PseudoType = "after"
    PseudoTypeBackdrop PseudoType = "backdrop"
    PseudoTypeSelection PseudoType = "selection"
    PseudoTypeFirstDashlineinherited PseudoType = "first-line-inherited"
    PseudoTypeScrollbar PseudoType = "scrollbar"
    PseudoTypeScrollbarDashthumb PseudoType = "scrollbar-thumb"
    PseudoTypeScrollbarDashbutton PseudoType = "scrollbar-button"
    PseudoTypeScrollbarDashtrack PseudoType = "scrollbar-track"
    PseudoTypeScrollbarDashtrackpiece PseudoType = "scrollbar-track-piece"
    PseudoTypeScrollbarDashcorner PseudoType = "scrollbar-corner"
    PseudoTypeResizer PseudoType = "resizer"
    PseudoTypeInputDashlistbutton PseudoType = "input-list-button"
)

const (
    ShadowRootTypeUserDashagent ShadowRootType = "user-agent"
    ShadowRootTypeOpen ShadowRootType = "open"
    ShadowRootTypeClosed ShadowRootType = "closed"
)

const (
    InspectModeSearchForNode InspectMode = "searchForNode"
    InspectModeSearchForUAShadowDOM InspectMode = "searchForUAShadowDOM"
    InspectModeNone InspectMode = "none"
)
type NodeId int64

type BackendNodeId int64

type BackendNode struct {
    NodeType int64 `json:"nodeType"`// <code>Node</code>'s nodeType.
    NodeName string `json:"nodeName"`// <code>Node</code>'s nodeName.
    BackendNodeId BackendNodeId `json:"backendNodeId"`
}

type PseudoType string

type ShadowRootType string

type Node struct {
    NodeId NodeId `json:"nodeId"`// Node identifier that is passed into the rest of the DOM messages as the <code>nodeId</code>. Backend will only push node with given <code>id</code> once. It is aware of all requested nodes and will only fire DOM events for nodes known to the client.
    ParentId *NodeId `json:"parentId,omitempty"`// [Experimental] The id of the parent node if any.
    BackendNodeId BackendNodeId `json:"backendNodeId"`// [Experimental] The BackendNodeId for this node.
    NodeType int64 `json:"nodeType"`// <code>Node</code>'s nodeType.
    NodeName string `json:"nodeName"`// <code>Node</code>'s nodeName.
    LocalName string `json:"localName"`// <code>Node</code>'s localName.
    NodeValue string `json:"nodeValue"`// <code>Node</code>'s nodeValue.
    ChildNodeCount *int64 `json:"childNodeCount,omitempty"`// Child count for <code>Container</code> nodes.
    Children *[]Node `json:"children,omitempty"`// Child nodes of this node when requested with children.
    Attributes *[]string `json:"attributes,omitempty"`// Attributes of the <code>Element</code> node in the form of flat array <code>[name1, value1, name2, value2]</code>.
    DocumentURL *string `json:"documentURL,omitempty"`// Document URL that <code>Document</code> or <code>FrameOwner</code> node points to.
    BaseURL *string `json:"baseURL,omitempty"`// [Experimental] Base URL that <code>Document</code> or <code>FrameOwner</code> node uses for URL completion.
    PublicId *string `json:"publicId,omitempty"`// <code>DocumentType</code>'s publicId.
    SystemId *string `json:"systemId,omitempty"`// <code>DocumentType</code>'s systemId.
    InternalSubset *string `json:"internalSubset,omitempty"`// <code>DocumentType</code>'s internalSubset.
    XmlVersion *string `json:"xmlVersion,omitempty"`// <code>Document</code>'s XML version in case of XML documents.
    Name *string `json:"name,omitempty"`// <code>Attr</code>'s name.
    Value *string `json:"value,omitempty"`// <code>Attr</code>'s value.
    PseudoType *PseudoType `json:"pseudoType,omitempty"`// Pseudo element type for this node.
    ShadowRootType *ShadowRootType `json:"shadowRootType,omitempty"`// Shadow root type.
    FrameId *page.FrameId `json:"frameId,omitempty"`// [Experimental] Frame ID for frame owner elements.
    ContentDocument *Node `json:"contentDocument,omitempty"`// Content document for frame owner elements.
    ShadowRoots *[]Node `json:"shadowRoots,omitempty"`// [Experimental] Shadow root list for given element host.
    TemplateContent *Node `json:"templateContent,omitempty"`// [Experimental] Content document fragment for template elements.
    PseudoElements *[]Node `json:"pseudoElements,omitempty"`// [Experimental] Pseudo elements associated with this node.
    ImportedDocument *Node `json:"importedDocument,omitempty"`// Import document for the HTMLImport links.
    DistributedNodes *[]BackendNode `json:"distributedNodes,omitempty"`// [Experimental] Distributed nodes for given insertion point.
    IsSVG *bool `json:"isSVG,omitempty"`// [Experimental] Whether the node is SVG.
}

type RGBA struct {
    R int64 `json:"r"`// The red component, in the [0-255] range.
    G int64 `json:"g"`// The green component, in the [0-255] range.
    B int64 `json:"b"`// The blue component, in the [0-255] range.
    A *float64 `json:"a,omitempty"`// The alpha component, in the [0-1] range (default: 1).
}

type Quad []float64

type BoxModel struct {
    Content Quad `json:"content"`// Content box
    Padding Quad `json:"padding"`// Padding box
    Border Quad `json:"border"`// Border box
    Margin Quad `json:"margin"`// Margin box
    Width int64 `json:"width"`// Node width
    Height int64 `json:"height"`// Node height
    ShapeOutside *ShapeOutsideInfo `json:"shapeOutside,omitempty"`// Shape outside coordinates
}

type ShapeOutsideInfo struct {
    Bounds Quad `json:"bounds"`// Shape bounds
    Shape []interface{} `json:"shape"`// Shape coordinate details
    MarginShape []interface{} `json:"marginShape"`// Margin shape bounds
}

type Rect struct {
    X float64 `json:"x"`// X coordinate
    Y float64 `json:"y"`// Y coordinate
    Width float64 `json:"width"`// Rectangle width
    Height float64 `json:"height"`// Rectangle height
}

type HighlightConfig struct {
    ShowInfo *bool `json:"showInfo,omitempty"`// Whether the node info tooltip should be shown (default: false).
    ShowRulers *bool `json:"showRulers,omitempty"`// Whether the rulers should be shown (default: false).
    ShowExtensionLines *bool `json:"showExtensionLines,omitempty"`// Whether the extension lines from node to the rulers should be shown (default: false).
    DisplayAsMaterial *bool `json:"displayAsMaterial,omitempty"`// [Experimental]
    ContentColor *RGBA `json:"contentColor,omitempty"`// The content box highlight fill color (default: transparent).
    PaddingColor *RGBA `json:"paddingColor,omitempty"`// The padding highlight fill color (default: transparent).
    BorderColor *RGBA `json:"borderColor,omitempty"`// The border highlight fill color (default: transparent).
    MarginColor *RGBA `json:"marginColor,omitempty"`// The margin highlight fill color (default: transparent).
    EventTargetColor *RGBA `json:"eventTargetColor,omitempty"`// [Experimental] The event target element highlight fill color (default: transparent).
    ShapeColor *RGBA `json:"shapeColor,omitempty"`// [Experimental] The shape outside fill color (default: transparent).
    ShapeMarginColor *RGBA `json:"shapeMarginColor,omitempty"`// [Experimental] The shape margin fill color (default: transparent).
    SelectorList *string `json:"selectorList,omitempty"`// Selectors to highlight relevant nodes.
}

type InspectMode string
