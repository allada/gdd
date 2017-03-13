package security


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

type ShowCertificateViewerCommandFn struct {
    mux sync.RWMutex
    fn func(ShowCertificateViewerCommand)
}

func (a *ShowCertificateViewerCommandFn) Load() func(ShowCertificateViewerCommand) {
    a.mux.RLock()
    defer a.mux.RUnlock()
    return a.fn
}

func (a *ShowCertificateViewerCommandFn) Store(fn func(ShowCertificateViewerCommand)) {
    a.mux.Lock()
    defer a.mux.Unlock()
    a.fn = fn
}

type SecurityAgent struct {
    conn *shared.Connection
    commandHandlers struct {
        Enable EnableCommandFn
        Disable DisableCommandFn
        ShowCertificateViewer ShowCertificateViewerCommandFn
    }
}
func NewAgent(conn *shared.Connection) *SecurityAgent {
    agent := &SecurityAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *SecurityAgent) Name() string {
    return "Security"
}

func (agent *SecurityAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "showCertificateViewer":
            var out ShowCertificateViewerCommand
            agent.conn.SetupCommand(id, targetId, data, &out)
            fn := agent.commandHandlers.ShowCertificateViewer.Load()
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
func (agent *SecurityAgent) FireSecurityStateChanged(event SecurityStateChangedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Security.securityStateChanged",
        Params: event,
    })
}
func (agent *SecurityAgent) FireSecurityStateChangedOnTarget(targetId string,event SecurityStateChangedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Security.securityStateChanged",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *SecurityAgent) SetEnableHandler(handler func(EnableCommand)) {
    agent.commandHandlers.Enable.Store(handler)
}
func (agent *SecurityAgent) SetDisableHandler(handler func(DisableCommand)) {
    agent.commandHandlers.Disable.Store(handler)
}
func (agent *SecurityAgent) SetShowCertificateViewerHandler(handler func(ShowCertificateViewerCommand)) {
    agent.commandHandlers.ShowCertificateViewer.Store(handler)
}
func init() {

}
