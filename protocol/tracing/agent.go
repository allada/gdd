package tracing


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type TracingAgent struct {
    conn *shared.Connection
    commandChans struct {
        Start []chan<- StartCommand
        End []chan<- EndCommand
        GetCategories []chan<- GetCategoriesCommand
        RequestMemoryDump []chan<- RequestMemoryDumpCommand
        RecordClockSyncMarker []chan<- RecordClockSyncMarkerCommand
    }
}
func NewAgent(conn *shared.Connection) *TracingAgent {
    agent := &TracingAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *TracingAgent) Name() string {
    return "Tracing"
}

func (agent *TracingAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "start":
            var out StartCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Start {
                c <-out
            }
        case "end":
            var out EndCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.End {
                c <-out
            }
        case "getCategories":
            var out GetCategoriesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetCategories {
                c <-out
            }
        case "requestMemoryDump":
            var out RequestMemoryDumpCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RequestMemoryDump {
                c <-out
            }
        case "recordClockSyncMarker":
            var out RecordClockSyncMarkerCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RecordClockSyncMarker {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *TracingAgent) FireDataCollected(event DataCollectedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Tracing.dataCollected",
        Params: event,
    })
}
func (agent *TracingAgent) FireDataCollectedOnTarget(targetId string, event DataCollectedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Tracing.dataCollected",
        Params: event,
    })
}
func (agent *TracingAgent) FireTracingComplete(event TracingCompleteEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Tracing.tracingComplete",
        Params: event,
    })
}
func (agent *TracingAgent) FireTracingCompleteOnTarget(targetId string, event TracingCompleteEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Tracing.tracingComplete",
        Params: event,
    })
}
func (agent *TracingAgent) FireBufferUsage(event BufferUsageEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Tracing.bufferUsage",
        Params: event,
    })
}
func (agent *TracingAgent) FireBufferUsageOnTarget(targetId string, event BufferUsageEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Tracing.bufferUsage",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *TracingAgent) StartNotify() <-chan StartCommand {
    outChan := make(chan StartCommand)
    agent.commandChans.Start = append(agent.commandChans.Start, outChan)
    return outChan
}
func (agent *TracingAgent) EndNotify() <-chan EndCommand {
    outChan := make(chan EndCommand)
    agent.commandChans.End = append(agent.commandChans.End, outChan)
    return outChan
}
func (agent *TracingAgent) GetCategoriesNotify() <-chan GetCategoriesCommand {
    outChan := make(chan GetCategoriesCommand)
    agent.commandChans.GetCategories = append(agent.commandChans.GetCategories, outChan)
    return outChan
}
func (agent *TracingAgent) RequestMemoryDumpNotify() <-chan RequestMemoryDumpCommand {
    outChan := make(chan RequestMemoryDumpCommand)
    agent.commandChans.RequestMemoryDump = append(agent.commandChans.RequestMemoryDump, outChan)
    return outChan
}
func (agent *TracingAgent) RecordClockSyncMarkerNotify() <-chan RecordClockSyncMarkerCommand {
    outChan := make(chan RecordClockSyncMarkerCommand)
    agent.commandChans.RecordClockSyncMarker = append(agent.commandChans.RecordClockSyncMarker, outChan)
    return outChan
}
func init() {

}
