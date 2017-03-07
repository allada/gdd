package cachestorage


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type CacheStorageAgent struct {
    conn *shared.Connection
    commandChans struct {
        RequestCacheNames []chan<- RequestCacheNamesCommand
        RequestEntries []chan<- RequestEntriesCommand
        DeleteCache []chan<- DeleteCacheCommand
        DeleteEntry []chan<- DeleteEntryCommand
    }
}
func NewAgent(conn *shared.Connection) *CacheStorageAgent {
    agent := &CacheStorageAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *CacheStorageAgent) Name() string {
    return "CacheStorage"
}

func (agent *CacheStorageAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "requestCacheNames":
            var out RequestCacheNamesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RequestCacheNames {
                c <-out
            }
        case "requestEntries":
            var out RequestEntriesCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RequestEntries {
                c <-out
            }
        case "deleteCache":
            var out DeleteCacheCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.DeleteCache {
                c <-out
            }
        case "deleteEntry":
            var out DeleteEntryCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.DeleteEntry {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events

// Commands Sent From Frontend
func (agent *CacheStorageAgent) RequestCacheNamesNotify() <-chan RequestCacheNamesCommand {
    outChan := make(chan RequestCacheNamesCommand)
    agent.commandChans.RequestCacheNames = append(agent.commandChans.RequestCacheNames, outChan)
    return outChan
}
func (agent *CacheStorageAgent) RequestEntriesNotify() <-chan RequestEntriesCommand {
    outChan := make(chan RequestEntriesCommand)
    agent.commandChans.RequestEntries = append(agent.commandChans.RequestEntries, outChan)
    return outChan
}
func (agent *CacheStorageAgent) DeleteCacheNotify() <-chan DeleteCacheCommand {
    outChan := make(chan DeleteCacheCommand)
    agent.commandChans.DeleteCache = append(agent.commandChans.DeleteCache, outChan)
    return outChan
}
func (agent *CacheStorageAgent) DeleteEntryNotify() <-chan DeleteEntryCommand {
    outChan := make(chan DeleteEntryCommand)
    agent.commandChans.DeleteEntry = append(agent.commandChans.DeleteEntry, outChan)
    return outChan
}
func init() {

}
