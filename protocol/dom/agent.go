package dom


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

type GetDocumentCommandFn struct {
    mux sync.RWMutex
    fn func(GetDocumentCommand)
}

func (a *GetDocumentCommandFn) Load() func(GetDocumentCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetDocumentCommandFn) Store(fn func(GetDocumentCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetFlattenedDocumentCommandFn struct {
    mux sync.RWMutex
    fn func(GetFlattenedDocumentCommand)
}

func (a *GetFlattenedDocumentCommandFn) Load() func(GetFlattenedDocumentCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetFlattenedDocumentCommandFn) Store(fn func(GetFlattenedDocumentCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CollectClassNamesFromSubtreeCommandFn struct {
    mux sync.RWMutex
    fn func(CollectClassNamesFromSubtreeCommand)
}

func (a *CollectClassNamesFromSubtreeCommandFn) Load() func(CollectClassNamesFromSubtreeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CollectClassNamesFromSubtreeCommandFn) Store(fn func(CollectClassNamesFromSubtreeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RequestChildNodesCommandFn struct {
    mux sync.RWMutex
    fn func(RequestChildNodesCommand)
}

func (a *RequestChildNodesCommandFn) Load() func(RequestChildNodesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RequestChildNodesCommandFn) Store(fn func(RequestChildNodesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type QuerySelectorCommandFn struct {
    mux sync.RWMutex
    fn func(QuerySelectorCommand)
}

func (a *QuerySelectorCommandFn) Load() func(QuerySelectorCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *QuerySelectorCommandFn) Store(fn func(QuerySelectorCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type QuerySelectorAllCommandFn struct {
    mux sync.RWMutex
    fn func(QuerySelectorAllCommand)
}

func (a *QuerySelectorAllCommandFn) Load() func(QuerySelectorAllCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *QuerySelectorAllCommandFn) Store(fn func(QuerySelectorAllCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetNodeNameCommandFn struct {
    mux sync.RWMutex
    fn func(SetNodeNameCommand)
}

func (a *SetNodeNameCommandFn) Load() func(SetNodeNameCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetNodeNameCommandFn) Store(fn func(SetNodeNameCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetNodeValueCommandFn struct {
    mux sync.RWMutex
    fn func(SetNodeValueCommand)
}

func (a *SetNodeValueCommandFn) Load() func(SetNodeValueCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetNodeValueCommandFn) Store(fn func(SetNodeValueCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RemoveNodeCommandFn struct {
    mux sync.RWMutex
    fn func(RemoveNodeCommand)
}

func (a *RemoveNodeCommandFn) Load() func(RemoveNodeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RemoveNodeCommandFn) Store(fn func(RemoveNodeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetAttributeValueCommandFn struct {
    mux sync.RWMutex
    fn func(SetAttributeValueCommand)
}

func (a *SetAttributeValueCommandFn) Load() func(SetAttributeValueCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetAttributeValueCommandFn) Store(fn func(SetAttributeValueCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetAttributesAsTextCommandFn struct {
    mux sync.RWMutex
    fn func(SetAttributesAsTextCommand)
}

func (a *SetAttributesAsTextCommandFn) Load() func(SetAttributesAsTextCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetAttributesAsTextCommandFn) Store(fn func(SetAttributesAsTextCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RemoveAttributeCommandFn struct {
    mux sync.RWMutex
    fn func(RemoveAttributeCommand)
}

func (a *RemoveAttributeCommandFn) Load() func(RemoveAttributeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RemoveAttributeCommandFn) Store(fn func(RemoveAttributeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetOuterHTMLCommandFn struct {
    mux sync.RWMutex
    fn func(GetOuterHTMLCommand)
}

func (a *GetOuterHTMLCommandFn) Load() func(GetOuterHTMLCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetOuterHTMLCommandFn) Store(fn func(GetOuterHTMLCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetOuterHTMLCommandFn struct {
    mux sync.RWMutex
    fn func(SetOuterHTMLCommand)
}

func (a *SetOuterHTMLCommandFn) Load() func(SetOuterHTMLCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetOuterHTMLCommandFn) Store(fn func(SetOuterHTMLCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type PerformSearchCommandFn struct {
    mux sync.RWMutex
    fn func(PerformSearchCommand)
}

func (a *PerformSearchCommandFn) Load() func(PerformSearchCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *PerformSearchCommandFn) Store(fn func(PerformSearchCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetSearchResultsCommandFn struct {
    mux sync.RWMutex
    fn func(GetSearchResultsCommand)
}

func (a *GetSearchResultsCommandFn) Load() func(GetSearchResultsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetSearchResultsCommandFn) Store(fn func(GetSearchResultsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DiscardSearchResultsCommandFn struct {
    mux sync.RWMutex
    fn func(DiscardSearchResultsCommand)
}

func (a *DiscardSearchResultsCommandFn) Load() func(DiscardSearchResultsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *DiscardSearchResultsCommandFn) Store(fn func(DiscardSearchResultsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RequestNodeCommandFn struct {
    mux sync.RWMutex
    fn func(RequestNodeCommand)
}

func (a *RequestNodeCommandFn) Load() func(RequestNodeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RequestNodeCommandFn) Store(fn func(RequestNodeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetInspectModeCommandFn struct {
    mux sync.RWMutex
    fn func(SetInspectModeCommand)
}

func (a *SetInspectModeCommandFn) Load() func(SetInspectModeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetInspectModeCommandFn) Store(fn func(SetInspectModeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type HighlightRectCommandFn struct {
    mux sync.RWMutex
    fn func(HighlightRectCommand)
}

func (a *HighlightRectCommandFn) Load() func(HighlightRectCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *HighlightRectCommandFn) Store(fn func(HighlightRectCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type HighlightQuadCommandFn struct {
    mux sync.RWMutex
    fn func(HighlightQuadCommand)
}

func (a *HighlightQuadCommandFn) Load() func(HighlightQuadCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *HighlightQuadCommandFn) Store(fn func(HighlightQuadCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type HighlightNodeCommandFn struct {
    mux sync.RWMutex
    fn func(HighlightNodeCommand)
}

func (a *HighlightNodeCommandFn) Load() func(HighlightNodeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *HighlightNodeCommandFn) Store(fn func(HighlightNodeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type HideHighlightCommandFn struct {
    mux sync.RWMutex
    fn func(HideHighlightCommand)
}

func (a *HideHighlightCommandFn) Load() func(HideHighlightCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *HideHighlightCommandFn) Store(fn func(HideHighlightCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type HighlightFrameCommandFn struct {
    mux sync.RWMutex
    fn func(HighlightFrameCommand)
}

func (a *HighlightFrameCommandFn) Load() func(HighlightFrameCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *HighlightFrameCommandFn) Store(fn func(HighlightFrameCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type PushNodeByPathToFrontendCommandFn struct {
    mux sync.RWMutex
    fn func(PushNodeByPathToFrontendCommand)
}

func (a *PushNodeByPathToFrontendCommandFn) Load() func(PushNodeByPathToFrontendCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *PushNodeByPathToFrontendCommandFn) Store(fn func(PushNodeByPathToFrontendCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type PushNodesByBackendIdsToFrontendCommandFn struct {
    mux sync.RWMutex
    fn func(PushNodesByBackendIdsToFrontendCommand)
}

func (a *PushNodesByBackendIdsToFrontendCommandFn) Load() func(PushNodesByBackendIdsToFrontendCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *PushNodesByBackendIdsToFrontendCommandFn) Store(fn func(PushNodesByBackendIdsToFrontendCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetInspectedNodeCommandFn struct {
    mux sync.RWMutex
    fn func(SetInspectedNodeCommand)
}

func (a *SetInspectedNodeCommandFn) Load() func(SetInspectedNodeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetInspectedNodeCommandFn) Store(fn func(SetInspectedNodeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type ResolveNodeCommandFn struct {
    mux sync.RWMutex
    fn func(ResolveNodeCommand)
}

func (a *ResolveNodeCommandFn) Load() func(ResolveNodeCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ResolveNodeCommandFn) Store(fn func(ResolveNodeCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetAttributesCommandFn struct {
    mux sync.RWMutex
    fn func(GetAttributesCommand)
}

func (a *GetAttributesCommandFn) Load() func(GetAttributesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetAttributesCommandFn) Store(fn func(GetAttributesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CopyToCommandFn struct {
    mux sync.RWMutex
    fn func(CopyToCommand)
}

func (a *CopyToCommandFn) Load() func(CopyToCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CopyToCommandFn) Store(fn func(CopyToCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type MoveToCommandFn struct {
    mux sync.RWMutex
    fn func(MoveToCommand)
}

func (a *MoveToCommandFn) Load() func(MoveToCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *MoveToCommandFn) Store(fn func(MoveToCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type UndoCommandFn struct {
    mux sync.RWMutex
    fn func(UndoCommand)
}

func (a *UndoCommandFn) Load() func(UndoCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *UndoCommandFn) Store(fn func(UndoCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RedoCommandFn struct {
    mux sync.RWMutex
    fn func(RedoCommand)
}

func (a *RedoCommandFn) Load() func(RedoCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RedoCommandFn) Store(fn func(RedoCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type MarkUndoableStateCommandFn struct {
    mux sync.RWMutex
    fn func(MarkUndoableStateCommand)
}

func (a *MarkUndoableStateCommandFn) Load() func(MarkUndoableStateCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *MarkUndoableStateCommandFn) Store(fn func(MarkUndoableStateCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type FocusCommandFn struct {
    mux sync.RWMutex
    fn func(FocusCommand)
}

func (a *FocusCommandFn) Load() func(FocusCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *FocusCommandFn) Store(fn func(FocusCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SetFileInputFilesCommandFn struct {
    mux sync.RWMutex
    fn func(SetFileInputFilesCommand)
}

func (a *SetFileInputFilesCommandFn) Load() func(SetFileInputFilesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *SetFileInputFilesCommandFn) Store(fn func(SetFileInputFilesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetBoxModelCommandFn struct {
    mux sync.RWMutex
    fn func(GetBoxModelCommand)
}

func (a *GetBoxModelCommandFn) Load() func(GetBoxModelCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetBoxModelCommandFn) Store(fn func(GetBoxModelCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetNodeForLocationCommandFn struct {
    mux sync.RWMutex
    fn func(GetNodeForLocationCommand)
}

func (a *GetNodeForLocationCommandFn) Load() func(GetNodeForLocationCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetNodeForLocationCommandFn) Store(fn func(GetNodeForLocationCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetRelayoutBoundaryCommandFn struct {
    mux sync.RWMutex
    fn func(GetRelayoutBoundaryCommand)
}

func (a *GetRelayoutBoundaryCommandFn) Load() func(GetRelayoutBoundaryCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetRelayoutBoundaryCommandFn) Store(fn func(GetRelayoutBoundaryCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetHighlightObjectForTestCommandFn struct {
    mux sync.RWMutex
    fn func(GetHighlightObjectForTestCommand)
}

func (a *GetHighlightObjectForTestCommandFn) Load() func(GetHighlightObjectForTestCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetHighlightObjectForTestCommandFn) Store(fn func(GetHighlightObjectForTestCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DOMAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
        GetDocument GetDocumentCommandFn
        GetFlattenedDocument GetFlattenedDocumentCommandFn
        CollectClassNamesFromSubtree CollectClassNamesFromSubtreeCommandFn
        RequestChildNodes RequestChildNodesCommandFn
        QuerySelector QuerySelectorCommandFn
        QuerySelectorAll QuerySelectorAllCommandFn
        SetNodeName SetNodeNameCommandFn
        SetNodeValue SetNodeValueCommandFn
        RemoveNode RemoveNodeCommandFn
        SetAttributeValue SetAttributeValueCommandFn
        SetAttributesAsText SetAttributesAsTextCommandFn
        RemoveAttribute RemoveAttributeCommandFn
        GetOuterHTML GetOuterHTMLCommandFn
        SetOuterHTML SetOuterHTMLCommandFn
        PerformSearch PerformSearchCommandFn
        GetSearchResults GetSearchResultsCommandFn
        DiscardSearchResults DiscardSearchResultsCommandFn
        RequestNode RequestNodeCommandFn
        SetInspectMode SetInspectModeCommandFn
        HighlightRect HighlightRectCommandFn
        HighlightQuad HighlightQuadCommandFn
        HighlightNode HighlightNodeCommandFn
        HideHighlight HideHighlightCommandFn
        HighlightFrame HighlightFrameCommandFn
        PushNodeByPathToFrontend PushNodeByPathToFrontendCommandFn
        PushNodesByBackendIdsToFrontend PushNodesByBackendIdsToFrontendCommandFn
        SetInspectedNode SetInspectedNodeCommandFn
        ResolveNode ResolveNodeCommandFn
        GetAttributes GetAttributesCommandFn
        CopyTo CopyToCommandFn
        MoveTo MoveToCommandFn
        Undo UndoCommandFn
        Redo RedoCommandFn
        MarkUndoableState MarkUndoableStateCommandFn
        Focus FocusCommandFn
        SetFileInputFiles SetFileInputFilesCommandFn
        GetBoxModel GetBoxModelCommandFn
        GetNodeForLocation GetNodeForLocationCommandFn
        GetRelayoutBoundary GetRelayoutBoundaryCommandFn
        GetHighlightObjectForTest GetHighlightObjectForTestCommandFn
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
    defer func() {
        data := recover()
        switch data.(type) {
        case nil:
            return
        case shared.Warning:
            fmt.Println(data)
        case shared.Error:
            fmt.Println("Closing websocket because of following Error:")
            fmt.Println(data)
            agent.conn.Close()
        case error:
            fmt.Println("Closing websocket because of following Error:")
            fmt.Println(data)
            agent.conn.Close()
        default:
            fmt.Println("Should probably use shared.Warning or shared.Error instead to panic()")
            panic(data)
        }
    }()
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
        case "getDocument":
            var out GetDocumentCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetDocument.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getFlattenedDocument":
            var out GetFlattenedDocumentCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetFlattenedDocument.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "collectClassNamesFromSubtree":
            var out CollectClassNamesFromSubtreeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CollectClassNamesFromSubtree.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "requestChildNodes":
            var out RequestChildNodesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RequestChildNodes.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "querySelector":
            var out QuerySelectorCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.QuerySelector.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "querySelectorAll":
            var out QuerySelectorAllCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.QuerySelectorAll.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setNodeName":
            var out SetNodeNameCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetNodeName.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setNodeValue":
            var out SetNodeValueCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetNodeValue.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "removeNode":
            var out RemoveNodeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RemoveNode.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setAttributeValue":
            var out SetAttributeValueCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetAttributeValue.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setAttributesAsText":
            var out SetAttributesAsTextCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetAttributesAsText.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "removeAttribute":
            var out RemoveAttributeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RemoveAttribute.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getOuterHTML":
            var out GetOuterHTMLCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetOuterHTML.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setOuterHTML":
            var out SetOuterHTMLCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetOuterHTML.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "performSearch":
            var out PerformSearchCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.PerformSearch.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getSearchResults":
            var out GetSearchResultsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetSearchResults.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "discardSearchResults":
            var out DiscardSearchResultsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.DiscardSearchResults.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "requestNode":
            var out RequestNodeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RequestNode.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setInspectMode":
            var out SetInspectModeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetInspectMode.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "highlightRect":
            var out HighlightRectCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.HighlightRect.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "highlightQuad":
            var out HighlightQuadCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.HighlightQuad.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "highlightNode":
            var out HighlightNodeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.HighlightNode.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "hideHighlight":
            var out HideHighlightCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.HideHighlight.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "highlightFrame":
            var out HighlightFrameCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.HighlightFrame.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "pushNodeByPathToFrontend":
            var out PushNodeByPathToFrontendCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.PushNodeByPathToFrontend.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "pushNodesByBackendIdsToFrontend":
            var out PushNodesByBackendIdsToFrontendCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.PushNodesByBackendIdsToFrontend.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setInspectedNode":
            var out SetInspectedNodeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetInspectedNode.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "resolveNode":
            var out ResolveNodeCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ResolveNode.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getAttributes":
            var out GetAttributesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetAttributes.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "copyTo":
            var out CopyToCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CopyTo.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "moveTo":
            var out MoveToCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.MoveTo.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "undo":
            var out UndoCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Undo.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "redo":
            var out RedoCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Redo.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "markUndoableState":
            var out MarkUndoableStateCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.MarkUndoableState.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "focus":
            var out FocusCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Focus.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "setFileInputFiles":
            var out SetFileInputFilesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.SetFileInputFiles.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getBoxModel":
            var out GetBoxModelCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetBoxModel.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getNodeForLocation":
            var out GetNodeForLocationCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetNodeForLocation.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getRelayoutBoundary":
            var out GetRelayoutBoundaryCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetRelayoutBoundary.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getHighlightObjectForTest":
            var out GetHighlightObjectForTestCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetHighlightObjectForTest.Load()
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
func (agent *DOMAgent) FireDocumentUpdated() {
    agent.conn.Send(shared.Notification{
        Method: "DOM.documentUpdated",
    })
}
func (agent *DOMAgent) FireDocumentUpdatedOnTarget(targetId string) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.documentUpdated",
    })
}
func (agent *DOMAgent) FireInspectNodeRequested(event InspectNodeRequestedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "DOM.inspectNodeRequested",
        Params: event,
    })
}
func (agent *DOMAgent) FireInspectNodeRequestedOnTarget(targetId string,event InspectNodeRequestedEvent) {
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
func (agent *DOMAgent) FireSetChildNodesOnTarget(targetId string,event SetChildNodesEvent) {
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
func (agent *DOMAgent) FireAttributeModifiedOnTarget(targetId string,event AttributeModifiedEvent) {
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
func (agent *DOMAgent) FireAttributeRemovedOnTarget(targetId string,event AttributeRemovedEvent) {
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
func (agent *DOMAgent) FireInlineStyleInvalidatedOnTarget(targetId string,event InlineStyleInvalidatedEvent) {
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
func (agent *DOMAgent) FireCharacterDataModifiedOnTarget(targetId string,event CharacterDataModifiedEvent) {
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
func (agent *DOMAgent) FireChildNodeCountUpdatedOnTarget(targetId string,event ChildNodeCountUpdatedEvent) {
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
func (agent *DOMAgent) FireChildNodeInsertedOnTarget(targetId string,event ChildNodeInsertedEvent) {
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
func (agent *DOMAgent) FireChildNodeRemovedOnTarget(targetId string,event ChildNodeRemovedEvent) {
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
func (agent *DOMAgent) FireShadowRootPushedOnTarget(targetId string,event ShadowRootPushedEvent) {
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
func (agent *DOMAgent) FireShadowRootPoppedOnTarget(targetId string,event ShadowRootPoppedEvent) {
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
func (agent *DOMAgent) FirePseudoElementAddedOnTarget(targetId string,event PseudoElementAddedEvent) {
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
func (agent *DOMAgent) FirePseudoElementRemovedOnTarget(targetId string,event PseudoElementRemovedEvent) {
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
func (agent *DOMAgent) FireDistributedNodesUpdatedOnTarget(targetId string,event DistributedNodesUpdatedEvent) {
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
func (agent *DOMAgent) FireNodeHighlightRequestedOnTarget(targetId string,event NodeHighlightRequestedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "DOM.nodeHighlightRequested",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *DOMAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *DOMAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *DOMAgent) SetGetDocumentHandler(handler func(GetDocumentCommand)) {
    agent.commandHandlers.GetDocument.Store(handler)
}
func (agent *DOMAgent) SetGetFlattenedDocumentHandler(handler func(GetFlattenedDocumentCommand)) {
    agent.commandHandlers.GetFlattenedDocument.Store(handler)
}
func (agent *DOMAgent) SetCollectClassNamesFromSubtreeHandler(handler func(CollectClassNamesFromSubtreeCommand)) {
    agent.commandHandlers.CollectClassNamesFromSubtree.Store(handler)
}
func (agent *DOMAgent) SetRequestChildNodesHandler(handler func(RequestChildNodesCommand)) {
    agent.commandHandlers.RequestChildNodes.Store(handler)
}
func (agent *DOMAgent) SetQuerySelectorHandler(handler func(QuerySelectorCommand)) {
    agent.commandHandlers.QuerySelector.Store(handler)
}
func (agent *DOMAgent) SetQuerySelectorAllHandler(handler func(QuerySelectorAllCommand)) {
    agent.commandHandlers.QuerySelectorAll.Store(handler)
}
func (agent *DOMAgent) SetSetNodeNameHandler(handler func(SetNodeNameCommand)) {
    agent.commandHandlers.SetNodeName.Store(handler)
}
func (agent *DOMAgent) SetSetNodeValueHandler(handler func(SetNodeValueCommand)) {
    agent.commandHandlers.SetNodeValue.Store(handler)
}
func (agent *DOMAgent) SetRemoveNodeHandler(handler func(RemoveNodeCommand)) {
    agent.commandHandlers.RemoveNode.Store(handler)
}
func (agent *DOMAgent) SetSetAttributeValueHandler(handler func(SetAttributeValueCommand)) {
    agent.commandHandlers.SetAttributeValue.Store(handler)
}
func (agent *DOMAgent) SetSetAttributesAsTextHandler(handler func(SetAttributesAsTextCommand)) {
    agent.commandHandlers.SetAttributesAsText.Store(handler)
}
func (agent *DOMAgent) SetRemoveAttributeHandler(handler func(RemoveAttributeCommand)) {
    agent.commandHandlers.RemoveAttribute.Store(handler)
}
func (agent *DOMAgent) SetGetOuterHTMLHandler(handler func(GetOuterHTMLCommand)) {
    agent.commandHandlers.GetOuterHTML.Store(handler)
}
func (agent *DOMAgent) SetSetOuterHTMLHandler(handler func(SetOuterHTMLCommand)) {
    agent.commandHandlers.SetOuterHTML.Store(handler)
}
func (agent *DOMAgent) SetPerformSearchHandler(handler func(PerformSearchCommand)) {
    agent.commandHandlers.PerformSearch.Store(handler)
}
func (agent *DOMAgent) SetGetSearchResultsHandler(handler func(GetSearchResultsCommand)) {
    agent.commandHandlers.GetSearchResults.Store(handler)
}
func (agent *DOMAgent) SetDiscardSearchResultsHandler(handler func(DiscardSearchResultsCommand)) {
    agent.commandHandlers.DiscardSearchResults.Store(handler)
}
func (agent *DOMAgent) SetRequestNodeHandler(handler func(RequestNodeCommand)) {
    agent.commandHandlers.RequestNode.Store(handler)
}
func (agent *DOMAgent) SetSetInspectModeHandler(handler func(SetInspectModeCommand)) {
    agent.commandHandlers.SetInspectMode.Store(handler)
}
func (agent *DOMAgent) SetHighlightRectHandler(handler func(HighlightRectCommand)) {
    agent.commandHandlers.HighlightRect.Store(handler)
}
func (agent *DOMAgent) SetHighlightQuadHandler(handler func(HighlightQuadCommand)) {
    agent.commandHandlers.HighlightQuad.Store(handler)
}
func (agent *DOMAgent) SetHighlightNodeHandler(handler func(HighlightNodeCommand)) {
    agent.commandHandlers.HighlightNode.Store(handler)
}
func (agent *DOMAgent) SetHideHighlightHandler(handler func(HideHighlightCommand)) {
    agent.commandHandlers.HideHighlight.Store(handler)
}
func (agent *DOMAgent) SetHighlightFrameHandler(handler func(HighlightFrameCommand)) {
    agent.commandHandlers.HighlightFrame.Store(handler)
}
func (agent *DOMAgent) SetPushNodeByPathToFrontendHandler(handler func(PushNodeByPathToFrontendCommand)) {
    agent.commandHandlers.PushNodeByPathToFrontend.Store(handler)
}
func (agent *DOMAgent) SetPushNodesByBackendIdsToFrontendHandler(handler func(PushNodesByBackendIdsToFrontendCommand)) {
    agent.commandHandlers.PushNodesByBackendIdsToFrontend.Store(handler)
}
func (agent *DOMAgent) SetSetInspectedNodeHandler(handler func(SetInspectedNodeCommand)) {
    agent.commandHandlers.SetInspectedNode.Store(handler)
}
func (agent *DOMAgent) SetResolveNodeHandler(handler func(ResolveNodeCommand)) {
    agent.commandHandlers.ResolveNode.Store(handler)
}
func (agent *DOMAgent) SetGetAttributesHandler(handler func(GetAttributesCommand)) {
    agent.commandHandlers.GetAttributes.Store(handler)
}
func (agent *DOMAgent) SetCopyToHandler(handler func(CopyToCommand)) {
    agent.commandHandlers.CopyTo.Store(handler)
}
func (agent *DOMAgent) SetMoveToHandler(handler func(MoveToCommand)) {
    agent.commandHandlers.MoveTo.Store(handler)
}
func (agent *DOMAgent) SetUndoHandler(handler func(UndoCommand)) {
    agent.commandHandlers.Undo.Store(handler)
}
func (agent *DOMAgent) SetRedoHandler(handler func(RedoCommand)) {
    agent.commandHandlers.Redo.Store(handler)
}
func (agent *DOMAgent) SetMarkUndoableStateHandler(handler func(MarkUndoableStateCommand)) {
    agent.commandHandlers.MarkUndoableState.Store(handler)
}
func (agent *DOMAgent) SetFocusHandler(handler func(FocusCommand)) {
    agent.commandHandlers.Focus.Store(handler)
}
func (agent *DOMAgent) SetSetFileInputFilesHandler(handler func(SetFileInputFilesCommand)) {
    agent.commandHandlers.SetFileInputFiles.Store(handler)
}
func (agent *DOMAgent) SetGetBoxModelHandler(handler func(GetBoxModelCommand)) {
    agent.commandHandlers.GetBoxModel.Store(handler)
}
func (agent *DOMAgent) SetGetNodeForLocationHandler(handler func(GetNodeForLocationCommand)) {
    agent.commandHandlers.GetNodeForLocation.Store(handler)
}
func (agent *DOMAgent) SetGetRelayoutBoundaryHandler(handler func(GetRelayoutBoundaryCommand)) {
    agent.commandHandlers.GetRelayoutBoundary.Store(handler)
}
func (agent *DOMAgent) SetGetHighlightObjectForTestHandler(handler func(GetHighlightObjectForTestCommand)) {
    agent.commandHandlers.GetHighlightObjectForTest.Store(handler)
}
func init() {

}
