package heapprofiler


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type HeapProfilerAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        StartTrackingHeapObjects []chan<- StartTrackingHeapObjectsCommand
        StopTrackingHeapObjects []chan<- StopTrackingHeapObjectsCommand
        TakeHeapSnapshot []chan<- TakeHeapSnapshotCommand
        CollectGarbage []chan<- CollectGarbageCommand
        GetObjectByHeapObjectId []chan<- GetObjectByHeapObjectIdCommand
        AddInspectedHeapObject []chan<- AddInspectedHeapObjectCommand
        GetHeapObjectId []chan<- GetHeapObjectIdCommand
        StartSampling []chan<- StartSamplingCommand
        StopSampling []chan<- StopSamplingCommand
    }
}
func NewAgent(conn *shared.Connection) *HeapProfilerAgent {
    agent := &HeapProfilerAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *HeapProfilerAgent) Name() string {
    return "HeapProfiler"
}

func (agent *HeapProfilerAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "startTrackingHeapObjects":
            var out StartTrackingHeapObjectsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StartTrackingHeapObjects {
                c <-out
            }
        case "stopTrackingHeapObjects":
            var out StopTrackingHeapObjectsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StopTrackingHeapObjects {
                c <-out
            }
        case "takeHeapSnapshot":
            var out TakeHeapSnapshotCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.TakeHeapSnapshot {
                c <-out
            }
        case "collectGarbage":
            var out CollectGarbageCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CollectGarbage {
                c <-out
            }
        case "getObjectByHeapObjectId":
            var out GetObjectByHeapObjectIdCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetObjectByHeapObjectId {
                c <-out
            }
        case "addInspectedHeapObject":
            var out AddInspectedHeapObjectCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.AddInspectedHeapObject {
                c <-out
            }
        case "getHeapObjectId":
            var out GetHeapObjectIdCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetHeapObjectId {
                c <-out
            }
        case "startSampling":
            var out StartSamplingCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StartSampling {
                c <-out
            }
        case "stopSampling":
            var out StopSamplingCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StopSampling {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *HeapProfilerAgent) FireAddHeapSnapshotChunk(event AddHeapSnapshotChunkEvent) {
    agent.conn.Send(shared.Notification{
        Method: "HeapProfiler.addHeapSnapshotChunk",
        Params: event,
    })
}
func (agent *HeapProfilerAgent) FireAddHeapSnapshotChunkOnTarget(targetId string, event AddHeapSnapshotChunkEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "HeapProfiler.addHeapSnapshotChunk",
        Params: event,
    })
}
func (agent *HeapProfilerAgent) FireResetProfiles(event ResetProfilesEvent) {
    agent.conn.Send(shared.Notification{
        Method: "HeapProfiler.resetProfiles",
        Params: event,
    })
}
func (agent *HeapProfilerAgent) FireResetProfilesOnTarget(targetId string, event ResetProfilesEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "HeapProfiler.resetProfiles",
        Params: event,
    })
}
func (agent *HeapProfilerAgent) FireReportHeapSnapshotProgress(event ReportHeapSnapshotProgressEvent) {
    agent.conn.Send(shared.Notification{
        Method: "HeapProfiler.reportHeapSnapshotProgress",
        Params: event,
    })
}
func (agent *HeapProfilerAgent) FireReportHeapSnapshotProgressOnTarget(targetId string, event ReportHeapSnapshotProgressEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "HeapProfiler.reportHeapSnapshotProgress",
        Params: event,
    })
}
func (agent *HeapProfilerAgent) FireLastSeenObjectId(event LastSeenObjectIdEvent) {
    agent.conn.Send(shared.Notification{
        Method: "HeapProfiler.lastSeenObjectId",
        Params: event,
    })
}
func (agent *HeapProfilerAgent) FireLastSeenObjectIdOnTarget(targetId string, event LastSeenObjectIdEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "HeapProfiler.lastSeenObjectId",
        Params: event,
    })
}
func (agent *HeapProfilerAgent) FireHeapStatsUpdate(event HeapStatsUpdateEvent) {
    agent.conn.Send(shared.Notification{
        Method: "HeapProfiler.heapStatsUpdate",
        Params: event,
    })
}
func (agent *HeapProfilerAgent) FireHeapStatsUpdateOnTarget(targetId string, event HeapStatsUpdateEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "HeapProfiler.heapStatsUpdate",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *HeapProfilerAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *HeapProfilerAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *HeapProfilerAgent) StartTrackingHeapObjectsNotify() <-chan StartTrackingHeapObjectsCommand {
    outChan := make(chan StartTrackingHeapObjectsCommand)
    agent.commandChans.StartTrackingHeapObjects = append(agent.commandChans.StartTrackingHeapObjects, outChan)
    return outChan
}
func (agent *HeapProfilerAgent) StopTrackingHeapObjectsNotify() <-chan StopTrackingHeapObjectsCommand {
    outChan := make(chan StopTrackingHeapObjectsCommand)
    agent.commandChans.StopTrackingHeapObjects = append(agent.commandChans.StopTrackingHeapObjects, outChan)
    return outChan
}
func (agent *HeapProfilerAgent) TakeHeapSnapshotNotify() <-chan TakeHeapSnapshotCommand {
    outChan := make(chan TakeHeapSnapshotCommand)
    agent.commandChans.TakeHeapSnapshot = append(agent.commandChans.TakeHeapSnapshot, outChan)
    return outChan
}
func (agent *HeapProfilerAgent) CollectGarbageNotify() <-chan CollectGarbageCommand {
    outChan := make(chan CollectGarbageCommand)
    agent.commandChans.CollectGarbage = append(agent.commandChans.CollectGarbage, outChan)
    return outChan
}
func (agent *HeapProfilerAgent) GetObjectByHeapObjectIdNotify() <-chan GetObjectByHeapObjectIdCommand {
    outChan := make(chan GetObjectByHeapObjectIdCommand)
    agent.commandChans.GetObjectByHeapObjectId = append(agent.commandChans.GetObjectByHeapObjectId, outChan)
    return outChan
}
func (agent *HeapProfilerAgent) AddInspectedHeapObjectNotify() <-chan AddInspectedHeapObjectCommand {
    outChan := make(chan AddInspectedHeapObjectCommand)
    agent.commandChans.AddInspectedHeapObject = append(agent.commandChans.AddInspectedHeapObject, outChan)
    return outChan
}
func (agent *HeapProfilerAgent) GetHeapObjectIdNotify() <-chan GetHeapObjectIdCommand {
    outChan := make(chan GetHeapObjectIdCommand)
    agent.commandChans.GetHeapObjectId = append(agent.commandChans.GetHeapObjectId, outChan)
    return outChan
}
func (agent *HeapProfilerAgent) StartSamplingNotify() <-chan StartSamplingCommand {
    outChan := make(chan StartSamplingCommand)
    agent.commandChans.StartSampling = append(agent.commandChans.StartSampling, outChan)
    return outChan
}
func (agent *HeapProfilerAgent) StopSamplingNotify() <-chan StopSamplingCommand {
    outChan := make(chan StopSamplingCommand)
    agent.commandChans.StopSampling = append(agent.commandChans.StopSampling, outChan)
    return outChan
}
func init() {

}
