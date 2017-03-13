package dom


type DocumentUpdatedEvent struct {
}
type InspectNodeRequestedEvent struct {
    BackendNodeId BackendNodeId `json:"backendNodeId"`// Id of the node to inspect.
}
type SetChildNodesEvent struct {
    ParentId NodeId `json:"parentId"`// Parent node id to populate with children.
    Nodes []Node `json:"nodes"`// Child nodes array.
}
type AttributeModifiedEvent struct {
    NodeId NodeId `json:"nodeId"`// Id of the node that has changed.
    Name string `json:"name"`// Attribute name.
    Value string `json:"value"`// Attribute value.
}
type AttributeRemovedEvent struct {
    NodeId NodeId `json:"nodeId"`// Id of the node that has changed.
    Name string `json:"name"`// A ttribute name.
}
type InlineStyleInvalidatedEvent struct {
    NodeIds []NodeId `json:"nodeIds"`// Ids of the nodes for which the inline styles have been invalidated.
}
type CharacterDataModifiedEvent struct {
    NodeId NodeId `json:"nodeId"`// Id of the node that has changed.
    CharacterData string `json:"characterData"`// New text value.
}
type ChildNodeCountUpdatedEvent struct {
    NodeId NodeId `json:"nodeId"`// Id of the node that has changed.
    ChildNodeCount int64 `json:"childNodeCount"`// New node count.
}
type ChildNodeInsertedEvent struct {
    ParentNodeId NodeId `json:"parentNodeId"`// Id of the node that has changed.
    PreviousNodeId NodeId `json:"previousNodeId"`// If of the previous siblint.
    Node Node `json:"node"`// Inserted node data.
}
type ChildNodeRemovedEvent struct {
    ParentNodeId NodeId `json:"parentNodeId"`// Parent id.
    NodeId NodeId `json:"nodeId"`// Id of the node that has been removed.
}
type ShadowRootPushedEvent struct {
    HostId NodeId `json:"hostId"`// Host element id.
    Root Node `json:"root"`// Shadow root.
}
type ShadowRootPoppedEvent struct {
    HostId NodeId `json:"hostId"`// Host element id.
    RootId NodeId `json:"rootId"`// Shadow root id.
}
type PseudoElementAddedEvent struct {
    ParentId NodeId `json:"parentId"`// Pseudo element's parent element id.
    PseudoElement Node `json:"pseudoElement"`// The added pseudo element.
}
type PseudoElementRemovedEvent struct {
    ParentId NodeId `json:"parentId"`// Pseudo element's parent element id.
    PseudoElementId NodeId `json:"pseudoElementId"`// The removed pseudo element id.
}
type DistributedNodesUpdatedEvent struct {
    InsertionPointId NodeId `json:"insertionPointId"`// Insertion point where distrubuted nodes were updated.
    DistributedNodes []BackendNode `json:"distributedNodes"`// Distributed nodes for given insertion point.
}
type NodeHighlightRequestedEvent struct {
    NodeId NodeId `json:"nodeId"`
}