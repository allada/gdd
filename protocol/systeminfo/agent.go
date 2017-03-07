package systeminfo


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type SystemInfoAgent struct {
    conn *shared.Connection
    commandChans struct {
        GetInfo []chan<- GetInfoCommand
    }
}
func NewAgent(conn *shared.Connection) *SystemInfoAgent {
    agent := &SystemInfoAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *SystemInfoAgent) Name() string {
    return "SystemInfo"
}

func (agent *SystemInfoAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "getInfo":
            var out GetInfoCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetInfo {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events

// Commands Sent From Frontend
func (agent *SystemInfoAgent) GetInfoNotify() <-chan GetInfoCommand {
    outChan := make(chan GetInfoCommand)
    agent.commandChans.GetInfo = append(agent.commandChans.GetInfo, outChan)
    return outChan
}
func init() {

}
