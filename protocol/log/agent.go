package log


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type LogAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        Clear []chan<- ClearCommand
        StartViolationsReport []chan<- StartViolationsReportCommand
        StopViolationsReport []chan<- StopViolationsReportCommand
    }
}
func NewAgent(conn *shared.Connection) *LogAgent {
    agent := &LogAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *LogAgent) Name() string {
    return "Log"
}

func (agent *LogAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "clear":
            var out ClearCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Clear {
                c <-out
            }
        case "startViolationsReport":
            var out StartViolationsReportCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StartViolationsReport {
                c <-out
            }
        case "stopViolationsReport":
            var out StopViolationsReportCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StopViolationsReport {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *LogAgent) FireEntryAdded(event EntryAddedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Log.entryAdded",
        Params: event,
    })
}
func (agent *LogAgent) FireEntryAddedOnTarget(targetId string, event EntryAddedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Log.entryAdded",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *LogAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *LogAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *LogAgent) ClearNotify() <-chan ClearCommand {
    outChan := make(chan ClearCommand)
    agent.commandChans.Clear = append(agent.commandChans.Clear, outChan)
    return outChan
}
func (agent *LogAgent) StartViolationsReportNotify() <-chan StartViolationsReportCommand {
    outChan := make(chan StartViolationsReportCommand)
    agent.commandChans.StartViolationsReport = append(agent.commandChans.StartViolationsReport, outChan)
    return outChan
}
func (agent *LogAgent) StopViolationsReportNotify() <-chan StopViolationsReportCommand {
    outChan := make(chan StopViolationsReportCommand)
    agent.commandChans.StopViolationsReport = append(agent.commandChans.StopViolationsReport, outChan)
    return outChan
}
func init() {

}
