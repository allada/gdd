package profiler


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type ProfilerAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        SetSamplingInterval []chan<- SetSamplingIntervalCommand
        Start []chan<- StartCommand
        Stop []chan<- StopCommand
        StartPreciseCoverage []chan<- StartPreciseCoverageCommand
        StopPreciseCoverage []chan<- StopPreciseCoverageCommand
        TakePreciseCoverage []chan<- TakePreciseCoverageCommand
        GetBestEffortCoverage []chan<- GetBestEffortCoverageCommand
    }
}
func NewAgent(conn *shared.Connection) *ProfilerAgent {
    agent := &ProfilerAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *ProfilerAgent) Name() string {
    return "Profiler"
}

func (agent *ProfilerAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "setSamplingInterval":
            var out SetSamplingIntervalCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetSamplingInterval {
                c <-out
            }
        case "start":
            var out StartCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Start {
                c <-out
            }
        case "stop":
            var out StopCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.Stop {
                c <-out
            }
        case "startPreciseCoverage":
            var out StartPreciseCoverageCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StartPreciseCoverage {
                c <-out
            }
        case "stopPreciseCoverage":
            var out StopPreciseCoverageCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.StopPreciseCoverage {
                c <-out
            }
        case "takePreciseCoverage":
            var out TakePreciseCoverageCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.TakePreciseCoverage {
                c <-out
            }
        case "getBestEffortCoverage":
            var out GetBestEffortCoverageCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetBestEffortCoverage {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *ProfilerAgent) FireConsoleProfileStarted(event ConsoleProfileStartedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Profiler.consoleProfileStarted",
        Params: event,
    })
}
func (agent *ProfilerAgent) FireConsoleProfileStartedOnTarget(targetId string, event ConsoleProfileStartedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Profiler.consoleProfileStarted",
        Params: event,
    })
}
func (agent *ProfilerAgent) FireConsoleProfileFinished(event ConsoleProfileFinishedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Profiler.consoleProfileFinished",
        Params: event,
    })
}
func (agent *ProfilerAgent) FireConsoleProfileFinishedOnTarget(targetId string, event ConsoleProfileFinishedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Profiler.consoleProfileFinished",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *ProfilerAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *ProfilerAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *ProfilerAgent) SetSamplingIntervalNotify() <-chan SetSamplingIntervalCommand {
    outChan := make(chan SetSamplingIntervalCommand)
    agent.commandChans.SetSamplingInterval = append(agent.commandChans.SetSamplingInterval, outChan)
    return outChan
}
func (agent *ProfilerAgent) StartNotify() <-chan StartCommand {
    outChan := make(chan StartCommand)
    agent.commandChans.Start = append(agent.commandChans.Start, outChan)
    return outChan
}
func (agent *ProfilerAgent) StopNotify() <-chan StopCommand {
    outChan := make(chan StopCommand)
    agent.commandChans.Stop = append(agent.commandChans.Stop, outChan)
    return outChan
}
func (agent *ProfilerAgent) StartPreciseCoverageNotify() <-chan StartPreciseCoverageCommand {
    outChan := make(chan StartPreciseCoverageCommand)
    agent.commandChans.StartPreciseCoverage = append(agent.commandChans.StartPreciseCoverage, outChan)
    return outChan
}
func (agent *ProfilerAgent) StopPreciseCoverageNotify() <-chan StopPreciseCoverageCommand {
    outChan := make(chan StopPreciseCoverageCommand)
    agent.commandChans.StopPreciseCoverage = append(agent.commandChans.StopPreciseCoverage, outChan)
    return outChan
}
func (agent *ProfilerAgent) TakePreciseCoverageNotify() <-chan TakePreciseCoverageCommand {
    outChan := make(chan TakePreciseCoverageCommand)
    agent.commandChans.TakePreciseCoverage = append(agent.commandChans.TakePreciseCoverage, outChan)
    return outChan
}
func (agent *ProfilerAgent) GetBestEffortCoverageNotify() <-chan GetBestEffortCoverageCommand {
    outChan := make(chan GetBestEffortCoverageCommand)
    agent.commandChans.GetBestEffortCoverage = append(agent.commandChans.GetBestEffortCoverage, outChan)
    return outChan
}
func init() {

}
