package memory


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type MemoryAgent struct {
    conn *shared.Connection
    commandChans struct {
        GetDOMCounters []chan<- GetDOMCountersCommand
        SetPressureNotificationsSuppressed []chan<- SetPressureNotificationsSuppressedCommand
        SimulatePressureNotification []chan<- SimulatePressureNotificationCommand
    }
}
func NewAgent(conn *shared.Connection) *MemoryAgent {
    agent := &MemoryAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *MemoryAgent) Name() string {
    return "Memory"
}

func (agent *MemoryAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "getDOMCounters":
            var out GetDOMCountersCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetDOMCounters {
                c <-out
            }
        case "setPressureNotificationsSuppressed":
            var out SetPressureNotificationsSuppressedCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetPressureNotificationsSuppressed {
                c <-out
            }
        case "simulatePressureNotification":
            var out SimulatePressureNotificationCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SimulatePressureNotification {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events

// Commands Sent From Frontend
func (agent *MemoryAgent) GetDOMCountersNotify() <-chan GetDOMCountersCommand {
    outChan := make(chan GetDOMCountersCommand)
    agent.commandChans.GetDOMCounters = append(agent.commandChans.GetDOMCounters, outChan)
    return outChan
}
func (agent *MemoryAgent) SetPressureNotificationsSuppressedNotify() <-chan SetPressureNotificationsSuppressedCommand {
    outChan := make(chan SetPressureNotificationsSuppressedCommand)
    agent.commandChans.SetPressureNotificationsSuppressed = append(agent.commandChans.SetPressureNotificationsSuppressed, outChan)
    return outChan
}
func (agent *MemoryAgent) SimulatePressureNotificationNotify() <-chan SimulatePressureNotificationCommand {
    outChan := make(chan SimulatePressureNotificationCommand)
    agent.commandChans.SimulatePressureNotification = append(agent.commandChans.SimulatePressureNotification, outChan)
    return outChan
}
func init() {

}
