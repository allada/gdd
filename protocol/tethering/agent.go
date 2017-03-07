package tethering


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type TetheringAgent struct {
    conn *shared.Connection
    commandChans struct {
        Bind []chan<- BindCommand
        Unbind []chan<- UnbindCommand
    }
}
func NewAgent(conn *shared.Connection) *TetheringAgent {
    agent := &TetheringAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *TetheringAgent) Name() string {
    return "Tethering"
}

func (agent *TetheringAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "bind":
            var out BindCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Bind {
                c <-out
            }
        case "unbind":
            var out UnbindCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Unbind {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *TetheringAgent) FireAccepted(event AcceptedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Tethering.accepted",
        Params: event,
    })
}
func (agent *TetheringAgent) FireAcceptedOnTarget(targetId string, event AcceptedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Tethering.accepted",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *TetheringAgent) BindNotify() <-chan BindCommand {
    outChan := make(chan BindCommand)
    agent.commandChans.Bind = append(agent.commandChans.Bind, outChan)
    return outChan
}
func (agent *TetheringAgent) UnbindNotify() <-chan UnbindCommand {
    outChan := make(chan UnbindCommand)
    agent.commandChans.Unbind = append(agent.commandChans.Unbind, outChan)
    return outChan
}
func init() {

}
