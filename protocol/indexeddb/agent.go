package indexeddb


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type IndexedDBAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        RequestDatabaseNames []chan<- RequestDatabaseNamesCommand
        RequestDatabase []chan<- RequestDatabaseCommand
        RequestData []chan<- RequestDataCommand
        ClearObjectStore []chan<- ClearObjectStoreCommand
        DeleteDatabase []chan<- DeleteDatabaseCommand
    }
}
func NewAgent(conn *shared.Connection) *IndexedDBAgent {
    agent := &IndexedDBAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *IndexedDBAgent) Name() string {
    return "IndexedDB"
}

func (agent *IndexedDBAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "requestDatabaseNames":
            var out RequestDatabaseNamesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RequestDatabaseNames {
                c <-out
            }
        case "requestDatabase":
            var out RequestDatabaseCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RequestDatabase {
                c <-out
            }
        case "requestData":
            var out RequestDataCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RequestData {
                c <-out
            }
        case "clearObjectStore":
            var out ClearObjectStoreCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ClearObjectStore {
                c <-out
            }
        case "deleteDatabase":
            var out DeleteDatabaseCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.DeleteDatabase {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events

// Commands Sent From Frontend
func (agent *IndexedDBAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *IndexedDBAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *IndexedDBAgent) RequestDatabaseNamesNotify() <-chan RequestDatabaseNamesCommand {
    outChan := make(chan RequestDatabaseNamesCommand)
    agent.commandChans.RequestDatabaseNames = append(agent.commandChans.RequestDatabaseNames, outChan)
    return outChan
}
func (agent *IndexedDBAgent) RequestDatabaseNotify() <-chan RequestDatabaseCommand {
    outChan := make(chan RequestDatabaseCommand)
    agent.commandChans.RequestDatabase = append(agent.commandChans.RequestDatabase, outChan)
    return outChan
}
func (agent *IndexedDBAgent) RequestDataNotify() <-chan RequestDataCommand {
    outChan := make(chan RequestDataCommand)
    agent.commandChans.RequestData = append(agent.commandChans.RequestData, outChan)
    return outChan
}
func (agent *IndexedDBAgent) ClearObjectStoreNotify() <-chan ClearObjectStoreCommand {
    outChan := make(chan ClearObjectStoreCommand)
    agent.commandChans.ClearObjectStore = append(agent.commandChans.ClearObjectStore, outChan)
    return outChan
}
func (agent *IndexedDBAgent) DeleteDatabaseNotify() <-chan DeleteDatabaseCommand {
    outChan := make(chan DeleteDatabaseCommand)
    agent.commandChans.DeleteDatabase = append(agent.commandChans.DeleteDatabase, outChan)
    return outChan
}
func init() {

}
