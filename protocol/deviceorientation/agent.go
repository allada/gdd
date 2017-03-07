package deviceorientation


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type DeviceOrientationAgent struct {
    conn *shared.Connection
    commandChans struct {
        SetDeviceOrientationOverride []chan<- SetDeviceOrientationOverrideCommand
        ClearDeviceOrientationOverride []chan<- ClearDeviceOrientationOverrideCommand
    }
}
func NewAgent(conn *shared.Connection) *DeviceOrientationAgent {
    agent := &DeviceOrientationAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *DeviceOrientationAgent) Name() string {
    return "DeviceOrientation"
}

func (agent *DeviceOrientationAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "setDeviceOrientationOverride":
            var out SetDeviceOrientationOverrideCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetDeviceOrientationOverride {
                c <-out
            }
        case "clearDeviceOrientationOverride":
            var out ClearDeviceOrientationOverrideCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ClearDeviceOrientationOverride {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events

// Commands Sent From Frontend
func (agent *DeviceOrientationAgent) SetDeviceOrientationOverrideNotify() <-chan SetDeviceOrientationOverrideCommand {
    outChan := make(chan SetDeviceOrientationOverrideCommand)
    agent.commandChans.SetDeviceOrientationOverride = append(agent.commandChans.SetDeviceOrientationOverride, outChan)
    return outChan
}
func (agent *DeviceOrientationAgent) ClearDeviceOrientationOverrideNotify() <-chan ClearDeviceOrientationOverrideCommand {
    outChan := make(chan ClearDeviceOrientationOverrideCommand)
    agent.commandChans.ClearDeviceOrientationOverride = append(agent.commandChans.ClearDeviceOrientationOverride, outChan)
    return outChan
}
func init() {

}
