package debugger


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type DebuggerAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        SetBreakpointsActive []chan<- SetBreakpointsActiveCommand
        SetSkipAllPauses []chan<- SetSkipAllPausesCommand
        SetBreakpointByUrl []chan<- SetBreakpointByUrlCommand
        SetBreakpoint []chan<- SetBreakpointCommand
        RemoveBreakpoint []chan<- RemoveBreakpointCommand
        GetPossibleBreakpoints []chan<- GetPossibleBreakpointsCommand
        ContinueToLocation []chan<- ContinueToLocationCommand
        StepOver []chan<- StepOverCommand
        StepInto []chan<- StepIntoCommand
        StepOut []chan<- StepOutCommand
        Pause []chan<- PauseCommand
        Resume []chan<- ResumeCommand
        SearchInContent []chan<- SearchInContentCommand
        SetScriptSource []chan<- SetScriptSourceCommand
        RestartFrame []chan<- RestartFrameCommand
        GetScriptSource []chan<- GetScriptSourceCommand
        SetPauseOnExceptions []chan<- SetPauseOnExceptionsCommand
        EvaluateOnCallFrame []chan<- EvaluateOnCallFrameCommand
        SetVariableValue []chan<- SetVariableValueCommand
        SetAsyncCallStackDepth []chan<- SetAsyncCallStackDepthCommand
        SetBlackboxPatterns []chan<- SetBlackboxPatternsCommand
        SetBlackboxedRanges []chan<- SetBlackboxedRangesCommand
    }
}
func NewAgent(conn *shared.Connection) *DebuggerAgent {
    agent := &DebuggerAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *DebuggerAgent) Name() string {
    return "Debugger"
}

func (agent *DebuggerAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "setBreakpointsActive":
            var out SetBreakpointsActiveCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetBreakpointsActive {
                c <-out
            }
        case "setSkipAllPauses":
            var out SetSkipAllPausesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetSkipAllPauses {
                c <-out
            }
        case "setBreakpointByUrl":
            var out SetBreakpointByUrlCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetBreakpointByUrl {
                c <-out
            }
        case "setBreakpoint":
            var out SetBreakpointCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetBreakpoint {
                c <-out
            }
        case "removeBreakpoint":
            var out RemoveBreakpointCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RemoveBreakpoint {
                c <-out
            }
        case "getPossibleBreakpoints":
            var out GetPossibleBreakpointsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetPossibleBreakpoints {
                c <-out
            }
        case "continueToLocation":
            var out ContinueToLocationCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ContinueToLocation {
                c <-out
            }
        case "stepOver":
            var out StepOverCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StepOver {
                c <-out
            }
        case "stepInto":
            var out StepIntoCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StepInto {
                c <-out
            }
        case "stepOut":
            var out StepOutCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StepOut {
                c <-out
            }
        case "pause":
            var out PauseCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Pause {
                c <-out
            }
        case "resume":
            var out ResumeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Resume {
                c <-out
            }
        case "searchInContent":
            var out SearchInContentCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SearchInContent {
                c <-out
            }
        case "setScriptSource":
            var out SetScriptSourceCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetScriptSource {
                c <-out
            }
        case "restartFrame":
            var out RestartFrameCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RestartFrame {
                c <-out
            }
        case "getScriptSource":
            var out GetScriptSourceCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetScriptSource {
                c <-out
            }
        case "setPauseOnExceptions":
            var out SetPauseOnExceptionsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetPauseOnExceptions {
                c <-out
            }
        case "evaluateOnCallFrame":
            var out EvaluateOnCallFrameCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.EvaluateOnCallFrame {
                c <-out
            }
        case "setVariableValue":
            var out SetVariableValueCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetVariableValue {
                c <-out
            }
        case "setAsyncCallStackDepth":
            var out SetAsyncCallStackDepthCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetAsyncCallStackDepth {
                c <-out
            }
        case "setBlackboxPatterns":
            var out SetBlackboxPatternsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetBlackboxPatterns {
                c <-out
            }
        case "setBlackboxedRanges":
            var out SetBlackboxedRangesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetBlackboxedRanges {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *DebuggerAgent) FireScriptParsed(event ScriptParsedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Debugger.scriptParsed",
        Params: event,
    })
}
func (agent *DebuggerAgent) FireScriptParsedOnTarget(targetId string, event ScriptParsedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Debugger.scriptParsed",
        Params: event,
    })
}
func (agent *DebuggerAgent) FireScriptFailedToParse(event ScriptFailedToParseEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Debugger.scriptFailedToParse",
        Params: event,
    })
}
func (agent *DebuggerAgent) FireScriptFailedToParseOnTarget(targetId string, event ScriptFailedToParseEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Debugger.scriptFailedToParse",
        Params: event,
    })
}
func (agent *DebuggerAgent) FireBreakpointResolved(event BreakpointResolvedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Debugger.breakpointResolved",
        Params: event,
    })
}
func (agent *DebuggerAgent) FireBreakpointResolvedOnTarget(targetId string, event BreakpointResolvedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Debugger.breakpointResolved",
        Params: event,
    })
}
func (agent *DebuggerAgent) FirePaused(event PausedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Debugger.paused",
        Params: event,
    })
}
func (agent *DebuggerAgent) FirePausedOnTarget(targetId string, event PausedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Debugger.paused",
        Params: event,
    })
}
func (agent *DebuggerAgent) FireResumed(event ResumedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Debugger.resumed",
        Params: event,
    })
}
func (agent *DebuggerAgent) FireResumedOnTarget(targetId string, event ResumedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Debugger.resumed",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *DebuggerAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *DebuggerAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *DebuggerAgent) SetBreakpointsActiveNotify() <-chan SetBreakpointsActiveCommand {
    outChan := make(chan SetBreakpointsActiveCommand)
    agent.commandChans.SetBreakpointsActive = append(agent.commandChans.SetBreakpointsActive, outChan)
    return outChan
}
func (agent *DebuggerAgent) SetSkipAllPausesNotify() <-chan SetSkipAllPausesCommand {
    outChan := make(chan SetSkipAllPausesCommand)
    agent.commandChans.SetSkipAllPauses = append(agent.commandChans.SetSkipAllPauses, outChan)
    return outChan
}
func (agent *DebuggerAgent) SetBreakpointByUrlNotify() <-chan SetBreakpointByUrlCommand {
    outChan := make(chan SetBreakpointByUrlCommand)
    agent.commandChans.SetBreakpointByUrl = append(agent.commandChans.SetBreakpointByUrl, outChan)
    return outChan
}
func (agent *DebuggerAgent) SetBreakpointNotify() <-chan SetBreakpointCommand {
    outChan := make(chan SetBreakpointCommand)
    agent.commandChans.SetBreakpoint = append(agent.commandChans.SetBreakpoint, outChan)
    return outChan
}
func (agent *DebuggerAgent) RemoveBreakpointNotify() <-chan RemoveBreakpointCommand {
    outChan := make(chan RemoveBreakpointCommand)
    agent.commandChans.RemoveBreakpoint = append(agent.commandChans.RemoveBreakpoint, outChan)
    return outChan
}
func (agent *DebuggerAgent) GetPossibleBreakpointsNotify() <-chan GetPossibleBreakpointsCommand {
    outChan := make(chan GetPossibleBreakpointsCommand)
    agent.commandChans.GetPossibleBreakpoints = append(agent.commandChans.GetPossibleBreakpoints, outChan)
    return outChan
}
func (agent *DebuggerAgent) ContinueToLocationNotify() <-chan ContinueToLocationCommand {
    outChan := make(chan ContinueToLocationCommand)
    agent.commandChans.ContinueToLocation = append(agent.commandChans.ContinueToLocation, outChan)
    return outChan
}
func (agent *DebuggerAgent) StepOverNotify() <-chan StepOverCommand {
    outChan := make(chan StepOverCommand)
    agent.commandChans.StepOver = append(agent.commandChans.StepOver, outChan)
    return outChan
}
func (agent *DebuggerAgent) StepIntoNotify() <-chan StepIntoCommand {
    outChan := make(chan StepIntoCommand)
    agent.commandChans.StepInto = append(agent.commandChans.StepInto, outChan)
    return outChan
}
func (agent *DebuggerAgent) StepOutNotify() <-chan StepOutCommand {
    outChan := make(chan StepOutCommand)
    agent.commandChans.StepOut = append(agent.commandChans.StepOut, outChan)
    return outChan
}
func (agent *DebuggerAgent) PauseNotify() <-chan PauseCommand {
    outChan := make(chan PauseCommand)
    agent.commandChans.Pause = append(agent.commandChans.Pause, outChan)
    return outChan
}
func (agent *DebuggerAgent) ResumeNotify() <-chan ResumeCommand {
    outChan := make(chan ResumeCommand)
    agent.commandChans.Resume = append(agent.commandChans.Resume, outChan)
    return outChan
}
func (agent *DebuggerAgent) SearchInContentNotify() <-chan SearchInContentCommand {
    outChan := make(chan SearchInContentCommand)
    agent.commandChans.SearchInContent = append(agent.commandChans.SearchInContent, outChan)
    return outChan
}
func (agent *DebuggerAgent) SetScriptSourceNotify() <-chan SetScriptSourceCommand {
    outChan := make(chan SetScriptSourceCommand)
    agent.commandChans.SetScriptSource = append(agent.commandChans.SetScriptSource, outChan)
    return outChan
}
func (agent *DebuggerAgent) RestartFrameNotify() <-chan RestartFrameCommand {
    outChan := make(chan RestartFrameCommand)
    agent.commandChans.RestartFrame = append(agent.commandChans.RestartFrame, outChan)
    return outChan
}
func (agent *DebuggerAgent) GetScriptSourceNotify() <-chan GetScriptSourceCommand {
    outChan := make(chan GetScriptSourceCommand)
    agent.commandChans.GetScriptSource = append(agent.commandChans.GetScriptSource, outChan)
    return outChan
}
func (agent *DebuggerAgent) SetPauseOnExceptionsNotify() <-chan SetPauseOnExceptionsCommand {
    outChan := make(chan SetPauseOnExceptionsCommand)
    agent.commandChans.SetPauseOnExceptions = append(agent.commandChans.SetPauseOnExceptions, outChan)
    return outChan
}
func (agent *DebuggerAgent) EvaluateOnCallFrameNotify() <-chan EvaluateOnCallFrameCommand {
    outChan := make(chan EvaluateOnCallFrameCommand)
    agent.commandChans.EvaluateOnCallFrame = append(agent.commandChans.EvaluateOnCallFrame, outChan)
    return outChan
}
func (agent *DebuggerAgent) SetVariableValueNotify() <-chan SetVariableValueCommand {
    outChan := make(chan SetVariableValueCommand)
    agent.commandChans.SetVariableValue = append(agent.commandChans.SetVariableValue, outChan)
    return outChan
}
func (agent *DebuggerAgent) SetAsyncCallStackDepthNotify() <-chan SetAsyncCallStackDepthCommand {
    outChan := make(chan SetAsyncCallStackDepthCommand)
    agent.commandChans.SetAsyncCallStackDepth = append(agent.commandChans.SetAsyncCallStackDepth, outChan)
    return outChan
}
func (agent *DebuggerAgent) SetBlackboxPatternsNotify() <-chan SetBlackboxPatternsCommand {
    outChan := make(chan SetBlackboxPatternsCommand)
    agent.commandChans.SetBlackboxPatterns = append(agent.commandChans.SetBlackboxPatterns, outChan)
    return outChan
}
func (agent *DebuggerAgent) SetBlackboxedRangesNotify() <-chan SetBlackboxedRangesCommand {
    outChan := make(chan SetBlackboxedRangesCommand)
    agent.commandChans.SetBlackboxedRanges = append(agent.commandChans.SetBlackboxedRanges, outChan)
    return outChan
}
func init() {

}
