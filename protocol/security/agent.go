package security


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type SecurityAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        ShowCertificateViewer []chan<- ShowCertificateViewerCommand
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
        case "showCertificateViewer":
            var out ShowCertificateViewerCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ShowCertificateViewer {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *SecurityAgent) FireSecurityStateChanged(event SecurityStateChangedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Security.securityStateChanged",
        Params: event,
    })
}
func (agent *SecurityAgent) FireSecurityStateChangedOnTarget(targetId string, event SecurityStateChangedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Security.securityStateChanged",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *SecurityAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *SecurityAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *SecurityAgent) ShowCertificateViewerNotify() <-chan ShowCertificateViewerCommand {
    outChan := make(chan ShowCertificateViewerCommand)
    agent.commandChans.ShowCertificateViewer = append(agent.commandChans.ShowCertificateViewer, outChan)
    return outChan
}
func init() {

}
