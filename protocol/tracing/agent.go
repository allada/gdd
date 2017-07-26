package tracing


import (
    "github.com/allada/gdd/protocol/shared"
    "sync"
    "encoding/json"
    "fmt"
)
type StartCommandFn struct {
    mux sync.RWMutex
    fn func(StartCommand)
}

func (a *StartCommandFn) Load() func(StartCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *StartCommandFn) Store(fn func(StartCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type EndCommandFn struct {
    mux sync.RWMutex
    fn func(EndCommand)
}

func (a *EndCommandFn) Load() func(EndCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *EndCommandFn) Store(fn func(EndCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type GetCategoriesCommandFn struct {
    mux sync.RWMutex
    fn func(GetCategoriesCommand)
}

func (a *GetCategoriesCommandFn) Load() func(GetCategoriesCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *GetCategoriesCommandFn) Store(fn func(GetCategoriesCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RequestMemoryDumpCommandFn struct {
    mux sync.RWMutex
    fn func(RequestMemoryDumpCommand)
}

func (a *RequestMemoryDumpCommandFn) Load() func(RequestMemoryDumpCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RequestMemoryDumpCommandFn) Store(fn func(RequestMemoryDumpCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type RecordClockSyncMarkerCommandFn struct {
    mux sync.RWMutex
    fn func(RecordClockSyncMarkerCommand)
}

func (a *RecordClockSyncMarkerCommandFn) Load() func(RecordClockSyncMarkerCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *RecordClockSyncMarkerCommandFn) Store(fn func(RecordClockSyncMarkerCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type TracingAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Start StartCommandFn
        End EndCommandFn
        GetCategories GetCategoriesCommandFn
        RequestMemoryDump RequestMemoryDumpCommandFn
        RecordClockSyncMarker RecordClockSyncMarkerCommandFn
    }
}
func NewAgent(conn *shared.Connection) *TracingAgent {
    agent := &TracingAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *TracingAgent) Name() string {
    return "Tracing"
}

func (agent *TracingAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    defer shared.TryRecoverFromPanic(agent.conn)
    switch (funcName) {
        case "start":
            var out StartCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.Start.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "end":
            var out EndCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.End.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "getCategories":
            var out GetCategoriesCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.GetCategories.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "requestMemoryDump":
            var out RequestMemoryDumpCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RequestMemoryDump.Load()
            if fn == nil {
                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")
                break
            }
            fn(out)
        case "recordClockSyncMarker":
            var out RecordClockSyncMarkerCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.RecordClockSyncMarker.Load()
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
func (agent *TracingAgent) FireDataCollected(event DataCollectedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Tracing.dataCollected",
        Params: event,
    })
}
func (agent *TracingAgent) FireDataCollectedOnTarget(targetId string,event DataCollectedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Tracing.dataCollected",
        Params: event,
    })
}
func (agent *TracingAgent) FireTracingComplete(event TracingCompleteEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Tracing.tracingComplete",
        Params: event,
    })
}
func (agent *TracingAgent) FireTracingCompleteOnTarget(targetId string,event TracingCompleteEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Tracing.tracingComplete",
        Params: event,
    })
}
func (agent *TracingAgent) FireBufferUsage(event BufferUsageEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Tracing.bufferUsage",
        Params: event,
    })
}
func (agent *TracingAgent) FireBufferUsageOnTarget(targetId string,event BufferUsageEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Tracing.bufferUsage",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *TracingAgent) SetStartHandler(handler func(StartCommand)) {
    agent.commandHandlers.Start.Store(handler)
}
func (agent *TracingAgent) SetEndHandler(handler func(EndCommand)) {
    agent.commandHandlers.End.Store(handler)
}
func (agent *TracingAgent) SetGetCategoriesHandler(handler func(GetCategoriesCommand)) {
    agent.commandHandlers.GetCategories.Store(handler)
}
func (agent *TracingAgent) SetRequestMemoryDumpHandler(handler func(RequestMemoryDumpCommand)) {
    agent.commandHandlers.RequestMemoryDump.Store(handler)
}
func (agent *TracingAgent) SetRecordClockSyncMarkerHandler(handler func(RecordClockSyncMarkerCommand)) {
    agent.commandHandlers.RecordClockSyncMarker.Store(handler)
}
func init() {

}
