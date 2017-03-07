package runtime


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type RuntimeAgent struct {
    conn *shared.Connection
    commandChans struct {
        Evaluate []chan<- EvaluateCommand
        AwaitPromise []chan<- AwaitPromiseCommand
        CallFunctionOn []chan<- CallFunctionOnCommand
        GetProperties []chan<- GetPropertiesCommand
        ReleaseObject []chan<- ReleaseObjectCommand
        ReleaseObjectGroup []chan<- ReleaseObjectGroupCommand
        RunIfWaitingForDebugger []chan<- RunIfWaitingForDebuggerCommand
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        DiscardConsoleEntries []chan<- DiscardConsoleEntriesCommand
        SetCustomObjectFormatterEnabled []chan<- SetCustomObjectFormatterEnabledCommand
        CompileScript []chan<- CompileScriptCommand
        RunScript []chan<- RunScriptCommand
    }
}
func NewAgent(conn *shared.Connection) *RuntimeAgent {
    agent := &RuntimeAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *RuntimeAgent) Name() string {
    return "Runtime"
}

func (agent *RuntimeAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "evaluate":
            var out EvaluateCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Evaluate {
                c <-out
            }
        case "awaitPromise":
            var out AwaitPromiseCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.AwaitPromise {
                c <-out
            }
        case "callFunctionOn":
            var out CallFunctionOnCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CallFunctionOn {
                c <-out
            }
        case "getProperties":
            var out GetPropertiesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetProperties {
                c <-out
            }
        case "releaseObject":
            var out ReleaseObjectCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ReleaseObject {
                c <-out
            }
        case "releaseObjectGroup":
            var out ReleaseObjectGroupCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ReleaseObjectGroup {
                c <-out
            }
        case "runIfWaitingForDebugger":
            var out RunIfWaitingForDebuggerCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RunIfWaitingForDebugger {
                c <-out
            }
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
        case "discardConsoleEntries":
            var out DiscardConsoleEntriesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.DiscardConsoleEntries {
                c <-out
            }
        case "setCustomObjectFormatterEnabled":
            var out SetCustomObjectFormatterEnabledCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetCustomObjectFormatterEnabled {
                c <-out
            }
        case "compileScript":
            var out CompileScriptCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CompileScript {
                c <-out
            }
        case "runScript":
            var out RunScriptCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RunScript {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *RuntimeAgent) FireExecutionContextCreated(event ExecutionContextCreatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Runtime.executionContextCreated",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExecutionContextCreatedOnTarget(targetId string, event ExecutionContextCreatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Runtime.executionContextCreated",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExecutionContextDestroyed(event ExecutionContextDestroyedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Runtime.executionContextDestroyed",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExecutionContextDestroyedOnTarget(targetId string, event ExecutionContextDestroyedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Runtime.executionContextDestroyed",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExecutionContextsCleared(event ExecutionContextsClearedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Runtime.executionContextsCleared",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExecutionContextsClearedOnTarget(targetId string, event ExecutionContextsClearedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Runtime.executionContextsCleared",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExceptionThrown(event ExceptionThrownEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Runtime.exceptionThrown",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExceptionThrownOnTarget(targetId string, event ExceptionThrownEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Runtime.exceptionThrown",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExceptionRevoked(event ExceptionRevokedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Runtime.exceptionRevoked",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireExceptionRevokedOnTarget(targetId string, event ExceptionRevokedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Runtime.exceptionRevoked",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireConsoleAPICalled(event ConsoleAPICalledEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Runtime.consoleAPICalled",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireConsoleAPICalledOnTarget(targetId string, event ConsoleAPICalledEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Runtime.consoleAPICalled",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireInspectRequested(event InspectRequestedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Runtime.inspectRequested",
        Params: event,
    })
}
func (agent *RuntimeAgent) FireInspectRequestedOnTarget(targetId string, event InspectRequestedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Runtime.inspectRequested",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *RuntimeAgent) EvaluateNotify() <-chan EvaluateCommand {
    outChan := make(chan EvaluateCommand)
    agent.commandChans.Evaluate = append(agent.commandChans.Evaluate, outChan)
    return outChan
}
func (agent *RuntimeAgent) AwaitPromiseNotify() <-chan AwaitPromiseCommand {
    outChan := make(chan AwaitPromiseCommand)
    agent.commandChans.AwaitPromise = append(agent.commandChans.AwaitPromise, outChan)
    return outChan
}
func (agent *RuntimeAgent) CallFunctionOnNotify() <-chan CallFunctionOnCommand {
    outChan := make(chan CallFunctionOnCommand)
    agent.commandChans.CallFunctionOn = append(agent.commandChans.CallFunctionOn, outChan)
    return outChan
}
func (agent *RuntimeAgent) GetPropertiesNotify() <-chan GetPropertiesCommand {
    outChan := make(chan GetPropertiesCommand)
    agent.commandChans.GetProperties = append(agent.commandChans.GetProperties, outChan)
    return outChan
}
func (agent *RuntimeAgent) ReleaseObjectNotify() <-chan ReleaseObjectCommand {
    outChan := make(chan ReleaseObjectCommand)
    agent.commandChans.ReleaseObject = append(agent.commandChans.ReleaseObject, outChan)
    return outChan
}
func (agent *RuntimeAgent) ReleaseObjectGroupNotify() <-chan ReleaseObjectGroupCommand {
    outChan := make(chan ReleaseObjectGroupCommand)
    agent.commandChans.ReleaseObjectGroup = append(agent.commandChans.ReleaseObjectGroup, outChan)
    return outChan
}
func (agent *RuntimeAgent) RunIfWaitingForDebuggerNotify() <-chan RunIfWaitingForDebuggerCommand {
    outChan := make(chan RunIfWaitingForDebuggerCommand)
    agent.commandChans.RunIfWaitingForDebugger = append(agent.commandChans.RunIfWaitingForDebugger, outChan)
    return outChan
}
func (agent *RuntimeAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *RuntimeAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *RuntimeAgent) DiscardConsoleEntriesNotify() <-chan DiscardConsoleEntriesCommand {
    outChan := make(chan DiscardConsoleEntriesCommand)
    agent.commandChans.DiscardConsoleEntries = append(agent.commandChans.DiscardConsoleEntries, outChan)
    return outChan
}
func (agent *RuntimeAgent) SetCustomObjectFormatterEnabledNotify() <-chan SetCustomObjectFormatterEnabledCommand {
    outChan := make(chan SetCustomObjectFormatterEnabledCommand)
    agent.commandChans.SetCustomObjectFormatterEnabled = append(agent.commandChans.SetCustomObjectFormatterEnabled, outChan)
    return outChan
}
func (agent *RuntimeAgent) CompileScriptNotify() <-chan CompileScriptCommand {
    outChan := make(chan CompileScriptCommand)
    agent.commandChans.CompileScript = append(agent.commandChans.CompileScript, outChan)
    return outChan
}
func (agent *RuntimeAgent) RunScriptNotify() <-chan RunScriptCommand {
    outChan := make(chan RunScriptCommand)
    agent.commandChans.RunScript = append(agent.commandChans.RunScript, outChan)
    return outChan
}
func init() {

}
