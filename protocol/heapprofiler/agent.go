package heapprofiler


import (
    "../shared"
    "sync"
    "encoding/json"
    "fmt"
)
type EnableCommandFn struct {
    mux sync.RWMutex
    fn func(EnableCommand)
}

func (a *EnableCommandFn) Load() func(EnableCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *EnableCommandFn) Store(fn func(EnableCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type DisableCommandFn struct {
    mux sync.RWMutex
    fn func(DisableCommand)
}

func (a *DisableCommandFn) Load() func(DisableCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *DisableCommandFn) Store(fn func(DisableCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StartTrackingHeapObjectsCommandFn struct {
    mux sync.RWMutex
    fn func(StartTrackingHeapObjectsCommand)
}

func (a *StartTrackingHeapObjectsCommandFn) Load() func(StartTrackingHeapObjectsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StartTrackingHeapObjectsCommandFn) Store(fn func(StartTrackingHeapObjectsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StopTrackingHeapObjectsCommandFn struct {
    mux sync.RWMutex
    fn func(StopTrackingHeapObjectsCommand)
}

func (a *StopTrackingHeapObjectsCommandFn) Load() func(StopTrackingHeapObjectsCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StopTrackingHeapObjectsCommandFn) Store(fn func(StopTrackingHeapObjectsCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type TakeHeapSnapshotCommandFn struct {
    mux sync.RWMutex
    fn func(TakeHeapSnapshotCommand)
}

func (a *TakeHeapSnapshotCommandFn) Load() func(TakeHeapSnapshotCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *TakeHeapSnapshotCommandFn) Store(fn func(TakeHeapSnapshotCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type CollectGarbageCommandFn struct {
    mux sync.RWMutex
    fn func(CollectGarbageCommand)
}

func (a *CollectGarbageCommandFn) Load() func(CollectGarbageCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *CollectGarbageCommandFn) Store(fn func(CollectGarbageCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetObjectByHeapObjectIdCommandFn struct {
    mux sync.RWMutex
    fn func(GetObjectByHeapObjectIdCommand)
}

func (a *GetObjectByHeapObjectIdCommandFn) Load() func(GetObjectByHeapObjectIdCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetObjectByHeapObjectIdCommandFn) Store(fn func(GetObjectByHeapObjectIdCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type AddInspectedHeapObjectCommandFn struct {
    mux sync.RWMutex
    fn func(AddInspectedHeapObjectCommand)
}

func (a *AddInspectedHeapObjectCommandFn) Load() func(AddInspectedHeapObjectCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *AddInspectedHeapObjectCommandFn) Store(fn func(AddInspectedHeapObjectCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetHeapObjectIdCommandFn struct {
    mux sync.RWMutex
    fn func(GetHeapObjectIdCommand)
}

func (a *GetHeapObjectIdCommandFn) Load() func(GetHeapObjectIdCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetHeapObjectIdCommandFn) Store(fn func(GetHeapObjectIdCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StartSamplingCommandFn struct {
    mux sync.RWMutex
    fn func(StartSamplingCommand)
}

func (a *StartSamplingCommandFn) Load() func(StartSamplingCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StartSamplingCommandFn) Store(fn func(StartSamplingCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type StopSamplingCommandFn struct {
    mux sync.RWMutex
    fn func(StopSamplingCommand)
}

func (a *StopSamplingCommandFn) Load() func(StopSamplingCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StopSamplingCommandFn) Store(fn func(StopSamplingCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type HeapProfilerAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
        StartTrackingHeapObjects StartTrackingHeapObjectsCommandFn
        StopTrackingHeapObjects StopTrackingHeapObjectsCommandFn
        TakeHeapSnapshot TakeHeapSnapshotCommandFn
        CollectGarbage CollectGarbageCommandFn
        GetObjectByHeapObjectId GetObjectByHeapObjectIdCommandFn
        AddInspectedHeapObject AddInspectedHeapObjectCommandFn
        GetHeapObjectId GetHeapObjectIdCommandFn
        StartSampling StartSamplingCommandFn
        StopSampling StopSamplingCommandFn
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
    defer func() {
        data := recover()
        switch data.(type) {
        case nil:
            return
        case shared.Warning:
            fmt.Println(data)
        case shared.Error:
            fmt.Println("Closing websocket because of following Error:")
            fmt.Println(data)
            agent.conn.Close()
        case error:
            fmt.Println("Closing websocket because of following Error:")
            fmt.Println(data)
            agent.conn.Close()
        default:
            fmt.Println("Should probably use shared.Warning or shared.Error instead to panic()")
            panic(data)
        }
    }()
    switch (funcName) {
        case "enable":
            var out EnableCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Enable.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "disable":
            var out DisableCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Disable.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "startTrackingHeapObjects":
            var out StartTrackingHeapObjectsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StartTrackingHeapObjects.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "stopTrackingHeapObjects":
            var out StopTrackingHeapObjectsCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StopTrackingHeapObjects.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "takeHeapSnapshot":
            var out TakeHeapSnapshotCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.TakeHeapSnapshot.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "collectGarbage":
            var out CollectGarbageCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.CollectGarbage.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getObjectByHeapObjectId":
            var out GetObjectByHeapObjectIdCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetObjectByHeapObjectId.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "addInspectedHeapObject":
            var out AddInspectedHeapObjectCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.AddInspectedHeapObject.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getHeapObjectId":
            var out GetHeapObjectIdCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetHeapObjectId.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "startSampling":
            var out StartSamplingCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StartSampling.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "stopSampling":
            var out StopSamplingCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.StopSampling.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        default:
            fmt.Printf("Command %s unknown\n", funcName)
            agent.conn.SendErrorResult(id, targetId, shared.ErrorCodeMethodNotFound, "")
    }
}

// Dispatchable Events
func (agent *HeapProfilerAgent) FireAddHeapSnapshotChunk(event AddHeapSnapshotChunkEvent) {
    agent.conn.Send(shared.Notification{
        Method: "HeapProfiler.addHeapSnapshotChunk",
        Params: event,
    })
}
func (agent *HeapProfilerAgent) FireAddHeapSnapshotChunkOnTarget(targetId string,event AddHeapSnapshotChunkEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "HeapProfiler.addHeapSnapshotChunk",
        Params: event,
    })
}
func (agent *HeapProfilerAgent) FireResetProfiles() {
    agent.conn.Send(shared.Notification{
        Method: "HeapProfiler.resetProfiles",
    })
}
func (agent *HeapProfilerAgent) FireResetProfilesOnTarget(targetId string) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "HeapProfiler.resetProfiles",
    })
}
func (agent *HeapProfilerAgent) FireReportHeapSnapshotProgress(event ReportHeapSnapshotProgressEvent) {
    agent.conn.Send(shared.Notification{
        Method: "HeapProfiler.reportHeapSnapshotProgress",
        Params: event,
    })
}
func (agent *HeapProfilerAgent) FireReportHeapSnapshotProgressOnTarget(targetId string,event ReportHeapSnapshotProgressEvent) {
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
func (agent *HeapProfilerAgent) FireLastSeenObjectIdOnTarget(targetId string,event LastSeenObjectIdEvent) {
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
func (agent *HeapProfilerAgent) FireHeapStatsUpdateOnTarget(targetId string,event HeapStatsUpdateEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "HeapProfiler.heapStatsUpdate",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *HeapProfilerAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *HeapProfilerAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *HeapProfilerAgent) SetStartTrackingHeapObjectsHandler(handler func(StartTrackingHeapObjectsCommand)) {
    agent.commandHandlers.StartTrackingHeapObjects.Store(handler)
}
func (agent *HeapProfilerAgent) SetStopTrackingHeapObjectsHandler(handler func(StopTrackingHeapObjectsCommand)) {
    agent.commandHandlers.StopTrackingHeapObjects.Store(handler)
}
func (agent *HeapProfilerAgent) SetTakeHeapSnapshotHandler(handler func(TakeHeapSnapshotCommand)) {
    agent.commandHandlers.TakeHeapSnapshot.Store(handler)
}
func (agent *HeapProfilerAgent) SetCollectGarbageHandler(handler func(CollectGarbageCommand)) {
    agent.commandHandlers.CollectGarbage.Store(handler)
}
func (agent *HeapProfilerAgent) SetGetObjectByHeapObjectIdHandler(handler func(GetObjectByHeapObjectIdCommand)) {
    agent.commandHandlers.GetObjectByHeapObjectId.Store(handler)
}
func (agent *HeapProfilerAgent) SetAddInspectedHeapObjectHandler(handler func(AddInspectedHeapObjectCommand)) {
    agent.commandHandlers.AddInspectedHeapObject.Store(handler)
}
func (agent *HeapProfilerAgent) SetGetHeapObjectIdHandler(handler func(GetHeapObjectIdCommand)) {
    agent.commandHandlers.GetHeapObjectId.Store(handler)
}
func (agent *HeapProfilerAgent) SetStartSamplingHandler(handler func(StartSamplingCommand)) {
    agent.commandHandlers.StartSampling.Store(handler)
}
func (agent *HeapProfilerAgent) SetStopSamplingHandler(handler func(StopSamplingCommand)) {
    agent.commandHandlers.StopSampling.Store(handler)
}
func init() {

}
