package storage


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type StorageAgent struct {
    conn *shared.Connection
    commandChans struct {
        ClearDataForOrigin []chan<- ClearDataForOriginCommand
    }
}
func NewAgent(conn *shared.Connection) *StorageAgent {
    agent := &StorageAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *StorageAgent) Name() string {
    return "Storage"
}

func (agent *StorageAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "clearDataForOrigin":
            var out ClearDataForOriginCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ClearDataForOrigin {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events

// Commands Sent From Frontend
func (agent *StorageAgent) ClearDataForOriginNotify() <-chan ClearDataForOriginCommand {
    outChan := make(chan ClearDataForOriginCommand)
    agent.commandChans.ClearDataForOrigin = append(agent.commandChans.ClearDataForOrigin, outChan)
    return outChan
}
func init() {

}
