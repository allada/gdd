package database


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type DatabaseAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        GetDatabaseTableNames []chan<- GetDatabaseTableNamesCommand
        ExecuteSQL []chan<- ExecuteSQLCommand
    }
}
func NewAgent(conn *shared.Connection) *DatabaseAgent {
    agent := &DatabaseAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *DatabaseAgent) Name() string {
    return "Database"
}

func (agent *DatabaseAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "getDatabaseTableNames":
            var out GetDatabaseTableNamesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetDatabaseTableNames {
                c <-out
            }
        case "executeSQL":
            var out ExecuteSQLCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ExecuteSQL {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *DatabaseAgent) FireAddDatabase(event AddDatabaseEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Database.addDatabase",
        Params: event,
    })
}
func (agent *DatabaseAgent) FireAddDatabaseOnTarget(targetId string, event AddDatabaseEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Database.addDatabase",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *DatabaseAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *DatabaseAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *DatabaseAgent) GetDatabaseTableNamesNotify() <-chan GetDatabaseTableNamesCommand {
    outChan := make(chan GetDatabaseTableNamesCommand)
    agent.commandChans.GetDatabaseTableNames = append(agent.commandChans.GetDatabaseTableNames, outChan)
    return outChan
}
func (agent *DatabaseAgent) ExecuteSQLNotify() <-chan ExecuteSQLCommand {
    outChan := make(chan ExecuteSQLCommand)
    agent.commandChans.ExecuteSQL = append(agent.commandChans.ExecuteSQL, outChan)
    return outChan
}
func init() {

}
