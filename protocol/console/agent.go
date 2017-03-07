package console


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type ConsoleAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        ClearMessages []chan<- ClearMessagesCommand
    }
}
func NewAgent(conn *shared.Connection) *ConsoleAgent {
    agent := &ConsoleAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *ConsoleAgent) Name() string {
    return "Console"
}

func (agent *ConsoleAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "clearMessages":
            var out ClearMessagesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ClearMessages {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *ConsoleAgent) FireMessageAdded(event MessageAddedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Console.messageAdded",
        Params: event,
    })
}
func (agent *ConsoleAgent) FireMessageAddedOnTarget(targetId string, event MessageAddedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Console.messageAdded",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *ConsoleAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *ConsoleAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *ConsoleAgent) ClearMessagesNotify() <-chan ClearMessagesCommand {
    outChan := make(chan ClearMessagesCommand)
    agent.commandChans.ClearMessages = append(agent.commandChans.ClearMessages, outChan)
    return outChan
}
func init() {

}
