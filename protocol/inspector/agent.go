package inspector


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type InspectorAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
    }
}
func NewAgent(conn *shared.Connection) *InspectorAgent {
    agent := &InspectorAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *InspectorAgent) Name() string {
    return "Inspector"
}

func (agent *InspectorAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *InspectorAgent) FireDetached(event DetachedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Inspector.detached",
        Params: event,
    })
}
func (agent *InspectorAgent) FireDetachedOnTarget(targetId string, event DetachedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Inspector.detached",
        Params: event,
    })
}
func (agent *InspectorAgent) FireTargetCrashed(event TargetCrashedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Inspector.targetCrashed",
        Params: event,
    })
}
func (agent *InspectorAgent) FireTargetCrashedOnTarget(targetId string, event TargetCrashedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Inspector.targetCrashed",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *InspectorAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *InspectorAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func init() {

}
