package emulation


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type EmulationAgent struct {
    conn *shared.Connection
    commandChans struct {
        SetDeviceMetricsOverride []chan<- SetDeviceMetricsOverrideCommand
        ClearDeviceMetricsOverride []chan<- ClearDeviceMetricsOverrideCommand
        ForceViewport []chan<- ForceViewportCommand
        ResetViewport []chan<- ResetViewportCommand
        ResetPageScaleFactor []chan<- ResetPageScaleFactorCommand
        SetPageScaleFactor []chan<- SetPageScaleFactorCommand
        SetVisibleSize []chan<- SetVisibleSizeCommand
        SetScriptExecutionDisabled []chan<- SetScriptExecutionDisabledCommand
        SetGeolocationOverride []chan<- SetGeolocationOverrideCommand
        ClearGeolocationOverride []chan<- ClearGeolocationOverrideCommand
        SetTouchEmulationEnabled []chan<- SetTouchEmulationEnabledCommand
        SetEmulatedMedia []chan<- SetEmulatedMediaCommand
        SetCPUThrottlingRate []chan<- SetCPUThrottlingRateCommand
        CanEmulate []chan<- CanEmulateCommand
        SetVirtualTimePolicy []chan<- SetVirtualTimePolicyCommand
        SetDefaultBackgroundColorOverride []chan<- SetDefaultBackgroundColorOverrideCommand
    }
}
func NewAgent(conn *shared.Connection) *EmulationAgent {
    agent := &EmulationAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *EmulationAgent) Name() string {
    return "Emulation"
}

func (agent *EmulationAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "setDeviceMetricsOverride":
            var out SetDeviceMetricsOverrideCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetDeviceMetricsOverride {
                c <-out
            }
        case "clearDeviceMetricsOverride":
            var out ClearDeviceMetricsOverrideCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ClearDeviceMetricsOverride {
                c <-out
            }
        case "forceViewport":
            var out ForceViewportCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ForceViewport {
                c <-out
            }
        case "resetViewport":
            var out ResetViewportCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ResetViewport {
                c <-out
            }
        case "resetPageScaleFactor":
            var out ResetPageScaleFactorCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ResetPageScaleFactor {
                c <-out
            }
        case "setPageScaleFactor":
            var out SetPageScaleFactorCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetPageScaleFactor {
                c <-out
            }
        case "setVisibleSize":
            var out SetVisibleSizeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetVisibleSize {
                c <-out
            }
        case "setScriptExecutionDisabled":
            var out SetScriptExecutionDisabledCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetScriptExecutionDisabled {
                c <-out
            }
        case "setGeolocationOverride":
            var out SetGeolocationOverrideCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetGeolocationOverride {
                c <-out
            }
        case "clearGeolocationOverride":
            var out ClearGeolocationOverrideCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ClearGeolocationOverride {
                c <-out
            }
        case "setTouchEmulationEnabled":
            var out SetTouchEmulationEnabledCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetTouchEmulationEnabled {
                c <-out
            }
        case "setEmulatedMedia":
            var out SetEmulatedMediaCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetEmulatedMedia {
                c <-out
            }
        case "setCPUThrottlingRate":
            var out SetCPUThrottlingRateCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetCPUThrottlingRate {
                c <-out
            }
        case "canEmulate":
            var out CanEmulateCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.CanEmulate {
                c <-out
            }
        case "setVirtualTimePolicy":
            var out SetVirtualTimePolicyCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetVirtualTimePolicy {
                c <-out
            }
        case "setDefaultBackgroundColorOverride":
            var out SetDefaultBackgroundColorOverrideCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetDefaultBackgroundColorOverride {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *EmulationAgent) FireVirtualTimeBudgetExpired(event VirtualTimeBudgetExpiredEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Emulation.virtualTimeBudgetExpired",
        Params: event,
    })
}
func (agent *EmulationAgent) FireVirtualTimeBudgetExpiredOnTarget(targetId string, event VirtualTimeBudgetExpiredEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Emulation.virtualTimeBudgetExpired",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *EmulationAgent) SetDeviceMetricsOverrideNotify() <-chan SetDeviceMetricsOverrideCommand {
    outChan := make(chan SetDeviceMetricsOverrideCommand)
    agent.commandChans.SetDeviceMetricsOverride = append(agent.commandChans.SetDeviceMetricsOverride, outChan)
    return outChan
}
func (agent *EmulationAgent) ClearDeviceMetricsOverrideNotify() <-chan ClearDeviceMetricsOverrideCommand {
    outChan := make(chan ClearDeviceMetricsOverrideCommand)
    agent.commandChans.ClearDeviceMetricsOverride = append(agent.commandChans.ClearDeviceMetricsOverride, outChan)
    return outChan
}
func (agent *EmulationAgent) ForceViewportNotify() <-chan ForceViewportCommand {
    outChan := make(chan ForceViewportCommand)
    agent.commandChans.ForceViewport = append(agent.commandChans.ForceViewport, outChan)
    return outChan
}
func (agent *EmulationAgent) ResetViewportNotify() <-chan ResetViewportCommand {
    outChan := make(chan ResetViewportCommand)
    agent.commandChans.ResetViewport = append(agent.commandChans.ResetViewport, outChan)
    return outChan
}
func (agent *EmulationAgent) ResetPageScaleFactorNotify() <-chan ResetPageScaleFactorCommand {
    outChan := make(chan ResetPageScaleFactorCommand)
    agent.commandChans.ResetPageScaleFactor = append(agent.commandChans.ResetPageScaleFactor, outChan)
    return outChan
}
func (agent *EmulationAgent) SetPageScaleFactorNotify() <-chan SetPageScaleFactorCommand {
    outChan := make(chan SetPageScaleFactorCommand)
    agent.commandChans.SetPageScaleFactor = append(agent.commandChans.SetPageScaleFactor, outChan)
    return outChan
}
func (agent *EmulationAgent) SetVisibleSizeNotify() <-chan SetVisibleSizeCommand {
    outChan := make(chan SetVisibleSizeCommand)
    agent.commandChans.SetVisibleSize = append(agent.commandChans.SetVisibleSize, outChan)
    return outChan
}
func (agent *EmulationAgent) SetScriptExecutionDisabledNotify() <-chan SetScriptExecutionDisabledCommand {
    outChan := make(chan SetScriptExecutionDisabledCommand)
    agent.commandChans.SetScriptExecutionDisabled = append(agent.commandChans.SetScriptExecutionDisabled, outChan)
    return outChan
}
func (agent *EmulationAgent) SetGeolocationOverrideNotify() <-chan SetGeolocationOverrideCommand {
    outChan := make(chan SetGeolocationOverrideCommand)
    agent.commandChans.SetGeolocationOverride = append(agent.commandChans.SetGeolocationOverride, outChan)
    return outChan
}
func (agent *EmulationAgent) ClearGeolocationOverrideNotify() <-chan ClearGeolocationOverrideCommand {
    outChan := make(chan ClearGeolocationOverrideCommand)
    agent.commandChans.ClearGeolocationOverride = append(agent.commandChans.ClearGeolocationOverride, outChan)
    return outChan
}
func (agent *EmulationAgent) SetTouchEmulationEnabledNotify() <-chan SetTouchEmulationEnabledCommand {
    outChan := make(chan SetTouchEmulationEnabledCommand)
    agent.commandChans.SetTouchEmulationEnabled = append(agent.commandChans.SetTouchEmulationEnabled, outChan)
    return outChan
}
func (agent *EmulationAgent) SetEmulatedMediaNotify() <-chan SetEmulatedMediaCommand {
    outChan := make(chan SetEmulatedMediaCommand)
    agent.commandChans.SetEmulatedMedia = append(agent.commandChans.SetEmulatedMedia, outChan)
    return outChan
}
func (agent *EmulationAgent) SetCPUThrottlingRateNotify() <-chan SetCPUThrottlingRateCommand {
    outChan := make(chan SetCPUThrottlingRateCommand)
    agent.commandChans.SetCPUThrottlingRate = append(agent.commandChans.SetCPUThrottlingRate, outChan)
    return outChan
}
func (agent *EmulationAgent) CanEmulateNotify() <-chan CanEmulateCommand {
    outChan := make(chan CanEmulateCommand)
    agent.commandChans.CanEmulate = append(agent.commandChans.CanEmulate, outChan)
    return outChan
}
func (agent *EmulationAgent) SetVirtualTimePolicyNotify() <-chan SetVirtualTimePolicyCommand {
    outChan := make(chan SetVirtualTimePolicyCommand)
    agent.commandChans.SetVirtualTimePolicy = append(agent.commandChans.SetVirtualTimePolicy, outChan)
    return outChan
}
func (agent *EmulationAgent) SetDefaultBackgroundColorOverrideNotify() <-chan SetDefaultBackgroundColorOverrideCommand {
    outChan := make(chan SetDefaultBackgroundColorOverrideCommand)
    agent.commandChans.SetDefaultBackgroundColorOverride = append(agent.commandChans.SetDefaultBackgroundColorOverride, outChan)
    return outChan
}
func init() {

}
