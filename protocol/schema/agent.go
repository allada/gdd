package schema


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type SchemaAgent struct {
    conn *shared.Connection
    commandChans struct {
        GetDomains []chan<- GetDomainsCommand
    }
}
func NewAgent(conn *shared.Connection) *SchemaAgent {
    agent := &SchemaAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *SchemaAgent) Name() string {
    return "Schema"
}

func (agent *SchemaAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "getDomains":
            var out GetDomainsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetDomains {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events

// Commands Sent From Frontend
func (agent *SchemaAgent) GetDomainsNotify() <-chan GetDomainsCommand {
    outChan := make(chan GetDomainsCommand)
    agent.commandChans.GetDomains = append(agent.commandChans.GetDomains, outChan)
    return outChan
}
func init() {

}
