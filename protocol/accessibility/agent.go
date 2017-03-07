package accessibility


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type AccessibilityAgent struct {
    conn *shared.Connection
    commandChans struct {
        GetPartialAXTree []chan<- GetPartialAXTreeCommand
    }
}
func NewAgent(conn *shared.Connection) *AccessibilityAgent {
    agent := &AccessibilityAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *AccessibilityAgent) Name() string {
    return "Accessibility"
}

func (agent *AccessibilityAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "getPartialAXTree":
            var out GetPartialAXTreeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetPartialAXTree {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events

// Commands Sent From Frontend
func (agent *AccessibilityAgent) GetPartialAXTreeNotify() <-chan GetPartialAXTreeCommand {
    outChan := make(chan GetPartialAXTreeCommand)
    agent.commandChans.GetPartialAXTree = append(agent.commandChans.GetPartialAXTree, outChan)
    return outChan
}
func init() {

}
