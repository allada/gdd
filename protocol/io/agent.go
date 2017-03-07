package io


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type IOAgent struct {
    conn *shared.Connection
    commandChans struct {
        Read []chan<- ReadCommand
        Close []chan<- CloseCommand
    }
}
func NewAgent(conn *shared.Connection) *IOAgent {
    agent := &IOAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *IOAgent) Name() string {
    return "IO"
}

func (agent *IOAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "read":
            var out ReadCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Read {
                c <-out
            }
        case "close":
            var out CloseCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Close {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events

// Commands Sent From Frontend
func (agent *IOAgent) ReadNotify() <-chan ReadCommand {
    outChan := make(chan ReadCommand)
    agent.commandChans.Read = append(agent.commandChans.Read, outChan)
    return outChan
}
func (agent *IOAgent) CloseNotify() <-chan CloseCommand {
    outChan := make(chan CloseCommand)
    agent.commandChans.Close = append(agent.commandChans.Close, outChan)
    return outChan
}
func init() {

}
