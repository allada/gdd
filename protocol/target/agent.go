package target


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type TargetAgent struct {
    conn *shared.Connection
    commandChans struct {
        SetDiscoverTargets []chan<- SetDiscoverTargetsCommand
        SetAutoAttach []chan<- SetAutoAttachCommand
        SetAttachToFrames []chan<- SetAttachToFramesCommand
        SetRemoteLocations []chan<- SetRemoteLocationsCommand
        SendMessageToTarget []chan<- SendMessageToTargetCommand
        GetTargetInfo []chan<- GetTargetInfoCommand
        ActivateTarget []chan<- ActivateTargetCommand
        CloseTarget []chan<- CloseTargetCommand
        AttachToTarget []chan<- AttachToTargetCommand
        DetachFromTarget []chan<- DetachFromTargetCommand
        CreateBrowserContext []chan<- CreateBrowserContextCommand
        DisposeBrowserContext []chan<- DisposeBrowserContextCommand
        CreateTarget []chan<- CreateTargetCommand
        GetTargets []chan<- GetTargetsCommand
    }
}
func NewAgent(conn *shared.Connection) *TargetAgent {
    agent := &TargetAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *TargetAgent) Name() string {
    return "Target"
}

func (agent *TargetAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "setDiscoverTargets":
            var out SetDiscoverTargetsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetDiscoverTargets {
                c <-out
            }
        case "setAutoAttach":
            var out SetAutoAttachCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetAutoAttach {
                c <-out
            }
        case "setAttachToFrames":
            var out SetAttachToFramesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetAttachToFrames {
                c <-out
            }
        case "setRemoteLocations":
            var out SetRemoteLocationsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetRemoteLocations {
                c <-out
            }
        case "sendMessageToTarget":
            var out SendMessageToTargetCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SendMessageToTarget {
                c <-out
            }
        case "getTargetInfo":
            var out GetTargetInfoCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetTargetInfo {
                c <-out
            }
        case "activateTarget":
            var out ActivateTargetCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ActivateTarget {
                c <-out
            }
        case "closeTarget":
            var out CloseTargetCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CloseTarget {
                c <-out
            }
        case "attachToTarget":
            var out AttachToTargetCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.AttachToTarget {
                c <-out
            }
        case "detachFromTarget":
            var out DetachFromTargetCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.DetachFromTarget {
                c <-out
            }
        case "createBrowserContext":
            var out CreateBrowserContextCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CreateBrowserContext {
                c <-out
            }
        case "disposeBrowserContext":
            var out DisposeBrowserContextCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.DisposeBrowserContext {
                c <-out
            }
        case "createTarget":
            var out CreateTargetCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CreateTarget {
                c <-out
            }
        case "getTargets":
            var out GetTargetsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetTargets {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *TargetAgent) FireTargetCreated(event TargetCreatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Target.targetCreated",
        Params: event,
    })
}
func (agent *TargetAgent) FireTargetCreatedOnTarget(targetId string, event TargetCreatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Target.targetCreated",
        Params: event,
    })
}
func (agent *TargetAgent) FireTargetDestroyed(event TargetDestroyedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Target.targetDestroyed",
        Params: event,
    })
}
func (agent *TargetAgent) FireTargetDestroyedOnTarget(targetId string, event TargetDestroyedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Target.targetDestroyed",
        Params: event,
    })
}
func (agent *TargetAgent) FireAttachedToTarget(event AttachedToTargetEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Target.attachedToTarget",
        Params: event,
    })
}
func (agent *TargetAgent) FireAttachedToTargetOnTarget(targetId string, event AttachedToTargetEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Target.attachedToTarget",
        Params: event,
    })
}
func (agent *TargetAgent) FireDetachedFromTarget(event DetachedFromTargetEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Target.detachedFromTarget",
        Params: event,
    })
}
func (agent *TargetAgent) FireDetachedFromTargetOnTarget(targetId string, event DetachedFromTargetEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Target.detachedFromTarget",
        Params: event,
    })
}
func (agent *TargetAgent) FireReceivedMessageFromTarget(event ReceivedMessageFromTargetEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Target.receivedMessageFromTarget",
        Params: event,
    })
}
func (agent *TargetAgent) FireReceivedMessageFromTargetOnTarget(targetId string, event ReceivedMessageFromTargetEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Target.receivedMessageFromTarget",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *TargetAgent) SetDiscoverTargetsNotify() <-chan SetDiscoverTargetsCommand {
    outChan := make(chan SetDiscoverTargetsCommand)
    agent.commandChans.SetDiscoverTargets = append(agent.commandChans.SetDiscoverTargets, outChan)
    return outChan
}
func (agent *TargetAgent) SetAutoAttachNotify() <-chan SetAutoAttachCommand {
    outChan := make(chan SetAutoAttachCommand)
    agent.commandChans.SetAutoAttach = append(agent.commandChans.SetAutoAttach, outChan)
    return outChan
}
func (agent *TargetAgent) SetAttachToFramesNotify() <-chan SetAttachToFramesCommand {
    outChan := make(chan SetAttachToFramesCommand)
    agent.commandChans.SetAttachToFrames = append(agent.commandChans.SetAttachToFrames, outChan)
    return outChan
}
func (agent *TargetAgent) SetRemoteLocationsNotify() <-chan SetRemoteLocationsCommand {
    outChan := make(chan SetRemoteLocationsCommand)
    agent.commandChans.SetRemoteLocations = append(agent.commandChans.SetRemoteLocations, outChan)
    return outChan
}
func (agent *TargetAgent) SendMessageToTargetNotify() <-chan SendMessageToTargetCommand {
    outChan := make(chan SendMessageToTargetCommand)
    agent.commandChans.SendMessageToTarget = append(agent.commandChans.SendMessageToTarget, outChan)
    return outChan
}
func (agent *TargetAgent) GetTargetInfoNotify() <-chan GetTargetInfoCommand {
    outChan := make(chan GetTargetInfoCommand)
    agent.commandChans.GetTargetInfo = append(agent.commandChans.GetTargetInfo, outChan)
    return outChan
}
func (agent *TargetAgent) ActivateTargetNotify() <-chan ActivateTargetCommand {
    outChan := make(chan ActivateTargetCommand)
    agent.commandChans.ActivateTarget = append(agent.commandChans.ActivateTarget, outChan)
    return outChan
}
func (agent *TargetAgent) CloseTargetNotify() <-chan CloseTargetCommand {
    outChan := make(chan CloseTargetCommand)
    agent.commandChans.CloseTarget = append(agent.commandChans.CloseTarget, outChan)
    return outChan
}
func (agent *TargetAgent) AttachToTargetNotify() <-chan AttachToTargetCommand {
    outChan := make(chan AttachToTargetCommand)
    agent.commandChans.AttachToTarget = append(agent.commandChans.AttachToTarget, outChan)
    return outChan
}
func (agent *TargetAgent) DetachFromTargetNotify() <-chan DetachFromTargetCommand {
    outChan := make(chan DetachFromTargetCommand)
    agent.commandChans.DetachFromTarget = append(agent.commandChans.DetachFromTarget, outChan)
    return outChan
}
func (agent *TargetAgent) CreateBrowserContextNotify() <-chan CreateBrowserContextCommand {
    outChan := make(chan CreateBrowserContextCommand)
    agent.commandChans.CreateBrowserContext = append(agent.commandChans.CreateBrowserContext, outChan)
    return outChan
}
func (agent *TargetAgent) DisposeBrowserContextNotify() <-chan DisposeBrowserContextCommand {
    outChan := make(chan DisposeBrowserContextCommand)
    agent.commandChans.DisposeBrowserContext = append(agent.commandChans.DisposeBrowserContext, outChan)
    return outChan
}
func (agent *TargetAgent) CreateTargetNotify() <-chan CreateTargetCommand {
    outChan := make(chan CreateTargetCommand)
    agent.commandChans.CreateTarget = append(agent.commandChans.CreateTarget, outChan)
    return outChan
}
func (agent *TargetAgent) GetTargetsNotify() <-chan GetTargetsCommand {
    outChan := make(chan GetTargetsCommand)
    agent.commandChans.GetTargets = append(agent.commandChans.GetTargets, outChan)
    return outChan
}
func init() {

}
