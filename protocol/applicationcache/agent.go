package applicationcache


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type ApplicationCacheAgent struct {
    conn *shared.Connection
    commandChans struct {
        GetFramesWithManifests []chan<- GetFramesWithManifestsCommand
        Enable []chan<- EnableCommand
        GetManifestForFrame []chan<- GetManifestForFrameCommand
        GetApplicationCacheForFrame []chan<- GetApplicationCacheForFrameCommand
    }
}
func NewAgent(conn *shared.Connection) *ApplicationCacheAgent {
    agent := &ApplicationCacheAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *ApplicationCacheAgent) Name() string {
    return "ApplicationCache"
}

func (agent *ApplicationCacheAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "getFramesWithManifests":
            var out GetFramesWithManifestsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetFramesWithManifests {
                c <-out
            }
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
        case "getManifestForFrame":
            var out GetManifestForFrameCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetManifestForFrame {
                c <-out
            }
        case "getApplicationCacheForFrame":
            var out GetApplicationCacheForFrameCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetApplicationCacheForFrame {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *ApplicationCacheAgent) FireApplicationCacheStatusUpdated(event ApplicationCacheStatusUpdatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "ApplicationCache.applicationCacheStatusUpdated",
        Params: event,
    })
}
func (agent *ApplicationCacheAgent) FireApplicationCacheStatusUpdatedOnTarget(targetId string, event ApplicationCacheStatusUpdatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "ApplicationCache.applicationCacheStatusUpdated",
        Params: event,
    })
}
func (agent *ApplicationCacheAgent) FireNetworkStateUpdated(event NetworkStateUpdatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "ApplicationCache.networkStateUpdated",
        Params: event,
    })
}
func (agent *ApplicationCacheAgent) FireNetworkStateUpdatedOnTarget(targetId string, event NetworkStateUpdatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "ApplicationCache.networkStateUpdated",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *ApplicationCacheAgent) GetFramesWithManifestsNotify() <-chan GetFramesWithManifestsCommand {
    outChan := make(chan GetFramesWithManifestsCommand)
    agent.commandChans.GetFramesWithManifests = append(agent.commandChans.GetFramesWithManifests, outChan)
    return outChan
}
func (agent *ApplicationCacheAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *ApplicationCacheAgent) GetManifestForFrameNotify() <-chan GetManifestForFrameCommand {
    outChan := make(chan GetManifestForFrameCommand)
    agent.commandChans.GetManifestForFrame = append(agent.commandChans.GetManifestForFrame, outChan)
    return outChan
}
func (agent *ApplicationCacheAgent) GetApplicationCacheForFrameNotify() <-chan GetApplicationCacheForFrameCommand {
    outChan := make(chan GetApplicationCacheForFrameCommand)
    agent.commandChans.GetApplicationCacheForFrame = append(agent.commandChans.GetApplicationCacheForFrame, outChan)
    return outChan
}
func init() {

}
