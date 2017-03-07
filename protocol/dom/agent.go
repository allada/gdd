package dom


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type DOMAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        GetDocument []chan<- GetDocumentCommand
        GetFlattenedDocument []chan<- GetFlattenedDocumentCommand
        CollectClassNamesFromSubtree []chan<- CollectClassNamesFromSubtreeCommand
        RequestChildNodes []chan<- RequestChildNodesCommand
        QuerySelector []chan<- QuerySelectorCommand
        QuerySelectorAll []chan<- QuerySelectorAllCommand
        SetNodeName []chan<- SetNodeNameCommand
        SetNodeValue []chan<- SetNodeValueCommand
        RemoveNode []chan<- RemoveNodeCommand
        SetAttributeValue []chan<- SetAttributeValueCommand
        SetAttributesAsText []chan<- SetAttributesAsTextCommand
        RemoveAttribute []chan<- RemoveAttributeCommand
        GetOuterHTML []chan<- GetOuterHTMLCommand
        SetOuterHTML []chan<- SetOuterHTMLCommand
        PerformSearch []chan<- PerformSearchCommand
        GetSearchResults []chan<- GetSearchResultsCommand
        DiscardSearchResults []chan<- DiscardSearchResultsCommand
        RequestNode []chan<- RequestNodeCommand
        SetInspectMode []chan<- SetInspectModeCommand
        HighlightRect []chan<- HighlightRectCommand
        HighlightQuad []chan<- HighlightQuadCommand
        HighlightNode []chan<- HighlightNodeCommand
        HideHighlight []chan<- HideHighlightCommand
        HighlightFrame []chan<- HighlightFrameCommand
        PushNodeByPathToFrontend []chan<- PushNodeByPathToFrontendCommand
        PushNodesByBackendIdsToFrontend []chan<- PushNodesByBackendIdsToFrontendCommand
        SetInspectedNode []chan<- SetInspectedNodeCommand
        ResolveNode []chan<- ResolveNodeCommand
        GetAttributes []chan<- GetAttributesCommand
        CopyTo []chan<- CopyToCommand
        MoveTo []chan<- MoveToCommand
        Undo []chan<- UndoCommand
        Redo []chan<- RedoCommand
        MarkUndoableState []chan<- MarkUndoableStateCommand
        Focus []chan<- FocusCommand
        SetFileInputFiles []chan<- SetFileInputFilesCommand
        GetBoxModel []chan<- GetBoxModelCommand
        GetNodeForLocation []chan<- GetNodeForLocationCommand
        GetRelayoutBoundary []chan<- GetRelayoutBoundaryCommand
        GetHighlightObjectForTest []chan<- GetHighlightObjectForTestCommand
    }
}
func NewAgent(conn *shared.Connection) *DOMAgent {
    agent := &DOMAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *DOMAgent) Name() string {
    return "DOM"
}

func (agent *DOMAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "getDocument":
            var out GetDocumentCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetDocument {
                c <-out
            }
        case "getFlattenedDocument":
            var out GetFlattenedDocumentCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetFlattenedDocument {
                c <-out
            }
        case "collectClassNamesFromSubtree":
            var out CollectClassNamesFromSubtreeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CollectClassNamesFromSubtree {
                c <-out
            }
        case "requestChildNodes":
            var out RequestChildNodesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RequestChildNodes {
                c <-out
            }
        case "querySelector":
            var out QuerySelectorCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.QuerySelector {
                c <-out
            }
        case "querySelectorAll":
            var out QuerySelectorAllCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.QuerySelectorAll {
                c <-out
            }
        case "setNodeName":
            var out SetNodeNameCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetNodeName {
                c <-out
            }
        case "setNodeValue":
            var out SetNodeValueCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetNodeValue {
                c <-out
            }
        case "removeNode":
            var out RemoveNodeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RemoveNode {
                c <-out
            }
        case "setAttributeValue":
            var out SetAttributeValueCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetAttributeValue {
                c <-out
            }
        case "setAttributesAsText":
            var out SetAttributesAsTextCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetAttributesAsText {
                c <-out
            }
        case "removeAttribute":
            var out RemoveAttributeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RemoveAttribute {
                c <-out
            }
        case "getOuterHTML":
            var out GetOuterHTMLCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetOuterHTML {
                c <-out
            }
        case "setOuterHTML":
            var out SetOuterHTMLCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetOuterHTML {
                c <-out
            }
        case "performSearch":
            var out PerformSearchCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.PerformSearch {
                c <-out
            }
        case "getSearchResults":
            var out GetSearchResultsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetSearchResults {
                c <-out
            }
        case "discardSearchResults":
            var out DiscardSearchResultsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.DiscardSearchResults {
                c <-out
            }
        case "requestNode":
            var out RequestNodeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RequestNode {
                c <-out
            }
        case "setInspectMode":
            var out SetInspectModeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetInspectMode {
                c <-out
            }
        case "highlightRect":
            var out HighlightRectCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.HighlightRect {
                c <-out
            }
        case "highlightQuad":
            var out HighlightQuadCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.HighlightQuad {
                c <-out
            }
        case "highlightNode":
            var out HighlightNodeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.HighlightNode {
                c <-out
            }
        case "hideHighlight":
            var out HideHighlightCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.HideHighlight {
                c <-out
            }
        case "highlightFrame":
            var out HighlightFrameCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.HighlightFrame {
                c <-out
            }
        case "pushNodeByPathToFrontend":
            var out PushNodeByPathToFrontendCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.PushNodeByPathToFrontend {
                c <-out
            }
        case "pushNodesByBackendIdsToFrontend":
            var out PushNodesByBackendIdsToFrontendCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.PushNodesByBackendIdsToFrontend {
                c <-out
            }
        case "setInspectedNode":
            var out SetInspectedNodeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetInspectedNode {
                c <-out
            }
        case "resolveNode":
            var out ResolveNodeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ResolveNode {
                c <-out
            }
        case "getAttributes":
            var out GetAttributesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetAttributes {
                c <-out
            }
        case "copyTo":
            var out CopyToCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CopyTo {
                c <-out
            }
        case "moveTo":
            var out MoveToCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.MoveTo {
                c <-out
            }
        case "undo":
            var out UndoCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Undo {
                c <-out
            }
        case "redo":
            var out RedoCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Redo {
                c <-out
            }
        case "markUndoableState":
            var out MarkUndoableStateCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.MarkUndoableState {
                c <-out
            }
        case "focus":
            var out FocusCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Focus {
                c <-out
            }
        case "setFileInputFiles":
            var out SetFileInputFilesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetFileInputFiles {
                c <-out
            }
        case "getBoxModel":
            var out GetBoxModelCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetBoxModel {
                c <-out
            }
        case "getNodeForLocation":
            var out GetNodeForLocationCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetNodeForLocation {
                c <-out
            }
        case "getRelayoutBoundary":
            var out GetRelayoutBoundaryCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetRelayoutBoundary {
                c <-out
            }
        case "getHighlightObjectForTest":
            var out GetHighlightObjectForTestCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetHighlightObjectForTest {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *DOMAgent) FireDocumentUpdated(event DocumentUpdatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.documentUpdated",
        Params: event,
    })
}
func (agent *DOMAgent) FireDocumentUpdatedOnTarget(targetId string, event DocumentUpdatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.documentUpdated",
        Params: event,
    })
}
func (agent *DOMAgent) FireInspectNodeRequested(event InspectNodeRequestedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.inspectNodeRequested",
        Params: event,
    })
}
func (agent *DOMAgent) FireInspectNodeRequestedOnTarget(targetId string, event InspectNodeRequestedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.inspectNodeRequested",
        Params: event,
    })
}
func (agent *DOMAgent) FireSetChildNodes(event SetChildNodesEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.setChildNodes",
        Params: event,
    })
}
func (agent *DOMAgent) FireSetChildNodesOnTarget(targetId string, event SetChildNodesEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.setChildNodes",
        Params: event,
    })
}
func (agent *DOMAgent) FireAttributeModified(event AttributeModifiedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.attributeModified",
        Params: event,
    })
}
func (agent *DOMAgent) FireAttributeModifiedOnTarget(targetId string, event AttributeModifiedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.attributeModified",
        Params: event,
    })
}
func (agent *DOMAgent) FireAttributeRemoved(event AttributeRemovedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.attributeRemoved",
        Params: event,
    })
}
func (agent *DOMAgent) FireAttributeRemovedOnTarget(targetId string, event AttributeRemovedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.attributeRemoved",
        Params: event,
    })
}
func (agent *DOMAgent) FireInlineStyleInvalidated(event InlineStyleInvalidatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.inlineStyleInvalidated",
        Params: event,
    })
}
func (agent *DOMAgent) FireInlineStyleInvalidatedOnTarget(targetId string, event InlineStyleInvalidatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.inlineStyleInvalidated",
        Params: event,
    })
}
func (agent *DOMAgent) FireCharacterDataModified(event CharacterDataModifiedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.characterDataModified",
        Params: event,
    })
}
func (agent *DOMAgent) FireCharacterDataModifiedOnTarget(targetId string, event CharacterDataModifiedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.characterDataModified",
        Params: event,
    })
}
func (agent *DOMAgent) FireChildNodeCountUpdated(event ChildNodeCountUpdatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.childNodeCountUpdated",
        Params: event,
    })
}
func (agent *DOMAgent) FireChildNodeCountUpdatedOnTarget(targetId string, event ChildNodeCountUpdatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.childNodeCountUpdated",
        Params: event,
    })
}
func (agent *DOMAgent) FireChildNodeInserted(event ChildNodeInsertedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.childNodeInserted",
        Params: event,
    })
}
func (agent *DOMAgent) FireChildNodeInsertedOnTarget(targetId string, event ChildNodeInsertedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.childNodeInserted",
        Params: event,
    })
}
func (agent *DOMAgent) FireChildNodeRemoved(event ChildNodeRemovedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.childNodeRemoved",
        Params: event,
    })
}
func (agent *DOMAgent) FireChildNodeRemovedOnTarget(targetId string, event ChildNodeRemovedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.childNodeRemoved",
        Params: event,
    })
}
func (agent *DOMAgent) FireShadowRootPushed(event ShadowRootPushedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.shadowRootPushed",
        Params: event,
    })
}
func (agent *DOMAgent) FireShadowRootPushedOnTarget(targetId string, event ShadowRootPushedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.shadowRootPushed",
        Params: event,
    })
}
func (agent *DOMAgent) FireShadowRootPopped(event ShadowRootPoppedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.shadowRootPopped",
        Params: event,
    })
}
func (agent *DOMAgent) FireShadowRootPoppedOnTarget(targetId string, event ShadowRootPoppedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.shadowRootPopped",
        Params: event,
    })
}
func (agent *DOMAgent) FirePseudoElementAdded(event PseudoElementAddedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.pseudoElementAdded",
        Params: event,
    })
}
func (agent *DOMAgent) FirePseudoElementAddedOnTarget(targetId string, event PseudoElementAddedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.pseudoElementAdded",
        Params: event,
    })
}
func (agent *DOMAgent) FirePseudoElementRemoved(event PseudoElementRemovedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.pseudoElementRemoved",
        Params: event,
    })
}
func (agent *DOMAgent) FirePseudoElementRemovedOnTarget(targetId string, event PseudoElementRemovedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.pseudoElementRemoved",
        Params: event,
    })
}
func (agent *DOMAgent) FireDistributedNodesUpdated(event DistributedNodesUpdatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.distributedNodesUpdated",
        Params: event,
    })
}
func (agent *DOMAgent) FireDistributedNodesUpdatedOnTarget(targetId string, event DistributedNodesUpdatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.distributedNodesUpdated",
        Params: event,
    })
}
func (agent *DOMAgent) FireNodeHighlightRequested(event NodeHighlightRequestedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.nodeHighlightRequested",
        Params: event,
    })
}
func (agent *DOMAgent) FireNodeHighlightRequestedOnTarget(targetId string, event NodeHighlightRequestedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.nodeHighlightRequested",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *DOMAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *DOMAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *DOMAgent) GetDocumentNotify() <-chan GetDocumentCommand {
    outChan := make(chan GetDocumentCommand)
    agent.commandChans.GetDocument = append(agent.commandChans.GetDocument, outChan)
    return outChan
}
func (agent *DOMAgent) GetFlattenedDocumentNotify() <-chan GetFlattenedDocumentCommand {
    outChan := make(chan GetFlattenedDocumentCommand)
    agent.commandChans.GetFlattenedDocument = append(agent.commandChans.GetFlattenedDocument, outChan)
    return outChan
}
func (agent *DOMAgent) CollectClassNamesFromSubtreeNotify() <-chan CollectClassNamesFromSubtreeCommand {
    outChan := make(chan CollectClassNamesFromSubtreeCommand)
    agent.commandChans.CollectClassNamesFromSubtree = append(agent.commandChans.CollectClassNamesFromSubtree, outChan)
    return outChan
}
func (agent *DOMAgent) RequestChildNodesNotify() <-chan RequestChildNodesCommand {
    outChan := make(chan RequestChildNodesCommand)
    agent.commandChans.RequestChildNodes = append(agent.commandChans.RequestChildNodes, outChan)
    return outChan
}
func (agent *DOMAgent) QuerySelectorNotify() <-chan QuerySelectorCommand {
    outChan := make(chan QuerySelectorCommand)
    agent.commandChans.QuerySelector = append(agent.commandChans.QuerySelector, outChan)
    return outChan
}
func (agent *DOMAgent) QuerySelectorAllNotify() <-chan QuerySelectorAllCommand {
    outChan := make(chan QuerySelectorAllCommand)
    agent.commandChans.QuerySelectorAll = append(agent.commandChans.QuerySelectorAll, outChan)
    return outChan
}
func (agent *DOMAgent) SetNodeNameNotify() <-chan SetNodeNameCommand {
    outChan := make(chan SetNodeNameCommand)
    agent.commandChans.SetNodeName = append(agent.commandChans.SetNodeName, outChan)
    return outChan
}
func (agent *DOMAgent) SetNodeValueNotify() <-chan SetNodeValueCommand {
    outChan := make(chan SetNodeValueCommand)
    agent.commandChans.SetNodeValue = append(agent.commandChans.SetNodeValue, outChan)
    return outChan
}
func (agent *DOMAgent) RemoveNodeNotify() <-chan RemoveNodeCommand {
    outChan := make(chan RemoveNodeCommand)
    agent.commandChans.RemoveNode = append(agent.commandChans.RemoveNode, outChan)
    return outChan
}
func (agent *DOMAgent) SetAttributeValueNotify() <-chan SetAttributeValueCommand {
    outChan := make(chan SetAttributeValueCommand)
    agent.commandChans.SetAttributeValue = append(agent.commandChans.SetAttributeValue, outChan)
    return outChan
}
func (agent *DOMAgent) SetAttributesAsTextNotify() <-chan SetAttributesAsTextCommand {
    outChan := make(chan SetAttributesAsTextCommand)
    agent.commandChans.SetAttributesAsText = append(agent.commandChans.SetAttributesAsText, outChan)
    return outChan
}
func (agent *DOMAgent) RemoveAttributeNotify() <-chan RemoveAttributeCommand {
    outChan := make(chan RemoveAttributeCommand)
    agent.commandChans.RemoveAttribute = append(agent.commandChans.RemoveAttribute, outChan)
    return outChan
}
func (agent *DOMAgent) GetOuterHTMLNotify() <-chan GetOuterHTMLCommand {
    outChan := make(chan GetOuterHTMLCommand)
    agent.commandChans.GetOuterHTML = append(agent.commandChans.GetOuterHTML, outChan)
    return outChan
}
func (agent *DOMAgent) SetOuterHTMLNotify() <-chan SetOuterHTMLCommand {
    outChan := make(chan SetOuterHTMLCommand)
    agent.commandChans.SetOuterHTML = append(agent.commandChans.SetOuterHTML, outChan)
    return outChan
}
func (agent *DOMAgent) PerformSearchNotify() <-chan PerformSearchCommand {
    outChan := make(chan PerformSearchCommand)
    agent.commandChans.PerformSearch = append(agent.commandChans.PerformSearch, outChan)
    return outChan
}
func (agent *DOMAgent) GetSearchResultsNotify() <-chan GetSearchResultsCommand {
    outChan := make(chan GetSearchResultsCommand)
    agent.commandChans.GetSearchResults = append(agent.commandChans.GetSearchResults, outChan)
    return outChan
}
func (agent *DOMAgent) DiscardSearchResultsNotify() <-chan DiscardSearchResultsCommand {
    outChan := make(chan DiscardSearchResultsCommand)
    agent.commandChans.DiscardSearchResults = append(agent.commandChans.DiscardSearchResults, outChan)
    return outChan
}
func (agent *DOMAgent) RequestNodeNotify() <-chan RequestNodeCommand {
    outChan := make(chan RequestNodeCommand)
    agent.commandChans.RequestNode = append(agent.commandChans.RequestNode, outChan)
    return outChan
}
func (agent *DOMAgent) SetInspectModeNotify() <-chan SetInspectModeCommand {
    outChan := make(chan SetInspectModeCommand)
    agent.commandChans.SetInspectMode = append(agent.commandChans.SetInspectMode, outChan)
    return outChan
}
func (agent *DOMAgent) HighlightRectNotify() <-chan HighlightRectCommand {
    outChan := make(chan HighlightRectCommand)
    agent.commandChans.HighlightRect = append(agent.commandChans.HighlightRect, outChan)
    return outChan
}
func (agent *DOMAgent) HighlightQuadNotify() <-chan HighlightQuadCommand {
    outChan := make(chan HighlightQuadCommand)
    agent.commandChans.HighlightQuad = append(agent.commandChans.HighlightQuad, outChan)
    return outChan
}
func (agent *DOMAgent) HighlightNodeNotify() <-chan HighlightNodeCommand {
    outChan := make(chan HighlightNodeCommand)
    agent.commandChans.HighlightNode = append(agent.commandChans.HighlightNode, outChan)
    return outChan
}
func (agent *DOMAgent) HideHighlightNotify() <-chan HideHighlightCommand {
    outChan := make(chan HideHighlightCommand)
    agent.commandChans.HideHighlight = append(agent.commandChans.HideHighlight, outChan)
    return outChan
}
func (agent *DOMAgent) HighlightFrameNotify() <-chan HighlightFrameCommand {
    outChan := make(chan HighlightFrameCommand)
    agent.commandChans.HighlightFrame = append(agent.commandChans.HighlightFrame, outChan)
    return outChan
}
func (agent *DOMAgent) PushNodeByPathToFrontendNotify() <-chan PushNodeByPathToFrontendCommand {
    outChan := make(chan PushNodeByPathToFrontendCommand)
    agent.commandChans.PushNodeByPathToFrontend = append(agent.commandChans.PushNodeByPathToFrontend, outChan)
    return outChan
}
func (agent *DOMAgent) PushNodesByBackendIdsToFrontendNotify() <-chan PushNodesByBackendIdsToFrontendCommand {
    outChan := make(chan PushNodesByBackendIdsToFrontendCommand)
    agent.commandChans.PushNodesByBackendIdsToFrontend = append(agent.commandChans.PushNodesByBackendIdsToFrontend, outChan)
    return outChan
}
func (agent *DOMAgent) SetInspectedNodeNotify() <-chan SetInspectedNodeCommand {
    outChan := make(chan SetInspectedNodeCommand)
    agent.commandChans.SetInspectedNode = append(agent.commandChans.SetInspectedNode, outChan)
    return outChan
}
func (agent *DOMAgent) ResolveNodeNotify() <-chan ResolveNodeCommand {
    outChan := make(chan ResolveNodeCommand)
    agent.commandChans.ResolveNode = append(agent.commandChans.ResolveNode, outChan)
    return outChan
}
func (agent *DOMAgent) GetAttributesNotify() <-chan GetAttributesCommand {
    outChan := make(chan GetAttributesCommand)
    agent.commandChans.GetAttributes = append(agent.commandChans.GetAttributes, outChan)
    return outChan
}
func (agent *DOMAgent) CopyToNotify() <-chan CopyToCommand {
    outChan := make(chan CopyToCommand)
    agent.commandChans.CopyTo = append(agent.commandChans.CopyTo, outChan)
    return outChan
}
func (agent *DOMAgent) MoveToNotify() <-chan MoveToCommand {
    outChan := make(chan MoveToCommand)
    agent.commandChans.MoveTo = append(agent.commandChans.MoveTo, outChan)
    return outChan
}
func (agent *DOMAgent) UndoNotify() <-chan UndoCommand {
    outChan := make(chan UndoCommand)
    agent.commandChans.Undo = append(agent.commandChans.Undo, outChan)
    return outChan
}
func (agent *DOMAgent) RedoNotify() <-chan RedoCommand {
    outChan := make(chan RedoCommand)
    agent.commandChans.Redo = append(agent.commandChans.Redo, outChan)
    return outChan
}
func (agent *DOMAgent) MarkUndoableStateNotify() <-chan MarkUndoableStateCommand {
    outChan := make(chan MarkUndoableStateCommand)
    agent.commandChans.MarkUndoableState = append(agent.commandChans.MarkUndoableState, outChan)
    return outChan
}
func (agent *DOMAgent) FocusNotify() <-chan FocusCommand {
    outChan := make(chan FocusCommand)
    agent.commandChans.Focus = append(agent.commandChans.Focus, outChan)
    return outChan
}
func (agent *DOMAgent) SetFileInputFilesNotify() <-chan SetFileInputFilesCommand {
    outChan := make(chan SetFileInputFilesCommand)
    agent.commandChans.SetFileInputFiles = append(agent.commandChans.SetFileInputFiles, outChan)
    return outChan
}
func (agent *DOMAgent) GetBoxModelNotify() <-chan GetBoxModelCommand {
    outChan := make(chan GetBoxModelCommand)
    agent.commandChans.GetBoxModel = append(agent.commandChans.GetBoxModel, outChan)
    return outChan
}
func (agent *DOMAgent) GetNodeForLocationNotify() <-chan GetNodeForLocationCommand {
    outChan := make(chan GetNodeForLocationCommand)
    agent.commandChans.GetNodeForLocation = append(agent.commandChans.GetNodeForLocation, outChan)
    return outChan
}
func (agent *DOMAgent) GetRelayoutBoundaryNotify() <-chan GetRelayoutBoundaryCommand {
    outChan := make(chan GetRelayoutBoundaryCommand)
    agent.commandChans.GetRelayoutBoundary = append(agent.commandChans.GetRelayoutBoundary, outChan)
    return outChan
}
func (agent *DOMAgent) GetHighlightObjectForTestNotify() <-chan GetHighlightObjectForTestCommand {
    outChan := make(chan GetHighlightObjectForTestCommand)
    agent.commandChans.GetHighlightObjectForTest = append(agent.commandChans.GetHighlightObjectForTest, outChan)
    return outChan
}
func init() {

}
