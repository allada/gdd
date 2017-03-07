package layertree


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type LayerTreeAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        CompositingReasons []chan<- CompositingReasonsCommand
        MakeSnapshot []chan<- MakeSnapshotCommand
        LoadSnapshot []chan<- LoadSnapshotCommand
        ReleaseSnapshot []chan<- ReleaseSnapshotCommand
        ProfileSnapshot []chan<- ProfileSnapshotCommand
        ReplaySnapshot []chan<- ReplaySnapshotCommand
        SnapshotCommandLog []chan<- SnapshotCommandLogCommand
    }
}
func NewAgent(conn *shared.Connection) *LayerTreeAgent {
    agent := &LayerTreeAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *LayerTreeAgent) Name() string {
    return "LayerTree"
}

func (agent *LayerTreeAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "compositingReasons":
            var out CompositingReasonsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CompositingReasons {
                c <-out
            }
        case "makeSnapshot":
            var out MakeSnapshotCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.MakeSnapshot {
                c <-out
            }
        case "loadSnapshot":
            var out LoadSnapshotCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.LoadSnapshot {
                c <-out
            }
        case "releaseSnapshot":
            var out ReleaseSnapshotCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ReleaseSnapshot {
                c <-out
            }
        case "profileSnapshot":
            var out ProfileSnapshotCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ProfileSnapshot {
                c <-out
            }
        case "replaySnapshot":
            var out ReplaySnapshotCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ReplaySnapshot {
                c <-out
            }
        case "snapshotCommandLog":
            var out SnapshotCommandLogCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SnapshotCommandLog {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *LayerTreeAgent) FireLayerTreeDidChange(event LayerTreeDidChangeEvent) {
    agent.conn.Send(shared.Notification{
        Method: "LayerTree.layerTreeDidChange",
        Params: event,
    })
}
func (agent *LayerTreeAgent) FireLayerTreeDidChangeOnTarget(targetId string, event LayerTreeDidChangeEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "LayerTree.layerTreeDidChange",
        Params: event,
    })
}
func (agent *LayerTreeAgent) FireLayerPainted(event LayerPaintedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "LayerTree.layerPainted",
        Params: event,
    })
}
func (agent *LayerTreeAgent) FireLayerPaintedOnTarget(targetId string, event LayerPaintedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "LayerTree.layerPainted",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *LayerTreeAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *LayerTreeAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *LayerTreeAgent) CompositingReasonsNotify() <-chan CompositingReasonsCommand {
    outChan := make(chan CompositingReasonsCommand)
    agent.commandChans.CompositingReasons = append(agent.commandChans.CompositingReasons, outChan)
    return outChan
}
func (agent *LayerTreeAgent) MakeSnapshotNotify() <-chan MakeSnapshotCommand {
    outChan := make(chan MakeSnapshotCommand)
    agent.commandChans.MakeSnapshot = append(agent.commandChans.MakeSnapshot, outChan)
    return outChan
}
func (agent *LayerTreeAgent) LoadSnapshotNotify() <-chan LoadSnapshotCommand {
    outChan := make(chan LoadSnapshotCommand)
    agent.commandChans.LoadSnapshot = append(agent.commandChans.LoadSnapshot, outChan)
    return outChan
}
func (agent *LayerTreeAgent) ReleaseSnapshotNotify() <-chan ReleaseSnapshotCommand {
    outChan := make(chan ReleaseSnapshotCommand)
    agent.commandChans.ReleaseSnapshot = append(agent.commandChans.ReleaseSnapshot, outChan)
    return outChan
}
func (agent *LayerTreeAgent) ProfileSnapshotNotify() <-chan ProfileSnapshotCommand {
    outChan := make(chan ProfileSnapshotCommand)
    agent.commandChans.ProfileSnapshot = append(agent.commandChans.ProfileSnapshot, outChan)
    return outChan
}
func (agent *LayerTreeAgent) ReplaySnapshotNotify() <-chan ReplaySnapshotCommand {
    outChan := make(chan ReplaySnapshotCommand)
    agent.commandChans.ReplaySnapshot = append(agent.commandChans.ReplaySnapshot, outChan)
    return outChan
}
func (agent *LayerTreeAgent) SnapshotCommandLogNotify() <-chan SnapshotCommandLogCommand {
    outChan := make(chan SnapshotCommandLogCommand)
    agent.commandChans.SnapshotCommandLog = append(agent.commandChans.SnapshotCommandLog, outChan)
    return outChan
}
func init() {

}
