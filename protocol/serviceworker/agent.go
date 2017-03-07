package serviceworker


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type ServiceWorkerAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        Unregister []chan<- UnregisterCommand
        UpdateRegistration []chan<- UpdateRegistrationCommand
        StartWorker []chan<- StartWorkerCommand
        SkipWaiting []chan<- SkipWaitingCommand
        StopWorker []chan<- StopWorkerCommand
        InspectWorker []chan<- InspectWorkerCommand
        SetForceUpdateOnPageLoad []chan<- SetForceUpdateOnPageLoadCommand
        DeliverPushMessage []chan<- DeliverPushMessageCommand
        DispatchSyncEvent []chan<- DispatchSyncEventCommand
    }
}
func NewAgent(conn *shared.Connection) *ServiceWorkerAgent {
    agent := &ServiceWorkerAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *ServiceWorkerAgent) Name() string {
    return "ServiceWorker"
}

func (agent *ServiceWorkerAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "unregister":
            var out UnregisterCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Unregister {
                c <-out
            }
        case "updateRegistration":
            var out UpdateRegistrationCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.UpdateRegistration {
                c <-out
            }
        case "startWorker":
            var out StartWorkerCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StartWorker {
                c <-out
            }
        case "skipWaiting":
            var out SkipWaitingCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SkipWaiting {
                c <-out
            }
        case "stopWorker":
            var out StopWorkerCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StopWorker {
                c <-out
            }
        case "inspectWorker":
            var out InspectWorkerCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.InspectWorker {
                c <-out
            }
        case "setForceUpdateOnPageLoad":
            var out SetForceUpdateOnPageLoadCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetForceUpdateOnPageLoad {
                c <-out
            }
        case "deliverPushMessage":
            var out DeliverPushMessageCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.DeliverPushMessage {
                c <-out
            }
        case "dispatchSyncEvent":
            var out DispatchSyncEventCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.DispatchSyncEvent {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *ServiceWorkerAgent) FireWorkerRegistrationUpdated(event WorkerRegistrationUpdatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "ServiceWorker.workerRegistrationUpdated",
        Params: event,
    })
}
func (agent *ServiceWorkerAgent) FireWorkerRegistrationUpdatedOnTarget(targetId string, event WorkerRegistrationUpdatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "ServiceWorker.workerRegistrationUpdated",
        Params: event,
    })
}
func (agent *ServiceWorkerAgent) FireWorkerVersionUpdated(event WorkerVersionUpdatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "ServiceWorker.workerVersionUpdated",
        Params: event,
    })
}
func (agent *ServiceWorkerAgent) FireWorkerVersionUpdatedOnTarget(targetId string, event WorkerVersionUpdatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "ServiceWorker.workerVersionUpdated",
        Params: event,
    })
}
func (agent *ServiceWorkerAgent) FireWorkerErrorReported(event WorkerErrorReportedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "ServiceWorker.workerErrorReported",
        Params: event,
    })
}
func (agent *ServiceWorkerAgent) FireWorkerErrorReportedOnTarget(targetId string, event WorkerErrorReportedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "ServiceWorker.workerErrorReported",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *ServiceWorkerAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *ServiceWorkerAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *ServiceWorkerAgent) UnregisterNotify() <-chan UnregisterCommand {
    outChan := make(chan UnregisterCommand)
    agent.commandChans.Unregister = append(agent.commandChans.Unregister, outChan)
    return outChan
}
func (agent *ServiceWorkerAgent) UpdateRegistrationNotify() <-chan UpdateRegistrationCommand {
    outChan := make(chan UpdateRegistrationCommand)
    agent.commandChans.UpdateRegistration = append(agent.commandChans.UpdateRegistration, outChan)
    return outChan
}
func (agent *ServiceWorkerAgent) StartWorkerNotify() <-chan StartWorkerCommand {
    outChan := make(chan StartWorkerCommand)
    agent.commandChans.StartWorker = append(agent.commandChans.StartWorker, outChan)
    return outChan
}
func (agent *ServiceWorkerAgent) SkipWaitingNotify() <-chan SkipWaitingCommand {
    outChan := make(chan SkipWaitingCommand)
    agent.commandChans.SkipWaiting = append(agent.commandChans.SkipWaiting, outChan)
    return outChan
}
func (agent *ServiceWorkerAgent) StopWorkerNotify() <-chan StopWorkerCommand {
    outChan := make(chan StopWorkerCommand)
    agent.commandChans.StopWorker = append(agent.commandChans.StopWorker, outChan)
    return outChan
}
func (agent *ServiceWorkerAgent) InspectWorkerNotify() <-chan InspectWorkerCommand {
    outChan := make(chan InspectWorkerCommand)
    agent.commandChans.InspectWorker = append(agent.commandChans.InspectWorker, outChan)
    return outChan
}
func (agent *ServiceWorkerAgent) SetForceUpdateOnPageLoadNotify() <-chan SetForceUpdateOnPageLoadCommand {
    outChan := make(chan SetForceUpdateOnPageLoadCommand)
    agent.commandChans.SetForceUpdateOnPageLoad = append(agent.commandChans.SetForceUpdateOnPageLoad, outChan)
    return outChan
}
func (agent *ServiceWorkerAgent) DeliverPushMessageNotify() <-chan DeliverPushMessageCommand {
    outChan := make(chan DeliverPushMessageCommand)
    agent.commandChans.DeliverPushMessage = append(agent.commandChans.DeliverPushMessage, outChan)
    return outChan
}
func (agent *ServiceWorkerAgent) DispatchSyncEventNotify() <-chan DispatchSyncEventCommand {
    outChan := make(chan DispatchSyncEventCommand)
    agent.commandChans.DispatchSyncEvent = append(agent.commandChans.DispatchSyncEvent, outChan)
    return outChan
}
func init() {

}
