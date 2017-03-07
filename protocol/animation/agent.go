package animation


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type AnimationAgent struct {
    conn *shared.Connection
    commandChans struct {
        Enable []chan<- EnableCommand
        Disable []chan<- DisableCommand
        GetPlaybackRate []chan<- GetPlaybackRateCommand
        SetPlaybackRate []chan<- SetPlaybackRateCommand
        GetCurrentTime []chan<- GetCurrentTimeCommand
        SetPaused []chan<- SetPausedCommand
        SetTiming []chan<- SetTimingCommand
        SeekAnimations []chan<- SeekAnimationsCommand
        ReleaseAnimations []chan<- ReleaseAnimationsCommand
        ResolveAnimation []chan<- ResolveAnimationCommand
    }
}
func NewAgent(conn *shared.Connection) *AnimationAgent {
    agent := &AnimationAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *AnimationAgent) Name() string {
    return "Animation"
}

func (agent *AnimationAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
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
        case "getPlaybackRate":
            var out GetPlaybackRateCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetPlaybackRate {
                c <-out
            }
        case "setPlaybackRate":
            var out SetPlaybackRateCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetPlaybackRate {
                c <-out
            }
        case "getCurrentTime":
            var out GetCurrentTimeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetCurrentTime {
                c <-out
            }
        case "setPaused":
            var out SetPausedCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetPaused {
                c <-out
            }
        case "setTiming":
            var out SetTimingCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetTiming {
                c <-out
            }
        case "seekAnimations":
            var out SeekAnimationsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SeekAnimations {
                c <-out
            }
        case "releaseAnimations":
            var out ReleaseAnimationsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ReleaseAnimations {
                c <-out
            }
        case "resolveAnimation":
            var out ResolveAnimationCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.ResolveAnimation {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events
func (agent *AnimationAgent) FireAnimationCreated(event AnimationCreatedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Animation.animationCreated",
        Params: event,
    })
}
func (agent *AnimationAgent) FireAnimationCreatedOnTarget(targetId string, event AnimationCreatedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Animation.animationCreated",
        Params: event,
    })
}
func (agent *AnimationAgent) FireAnimationStarted(event AnimationStartedEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Animation.animationStarted",
        Params: event,
    })
}
func (agent *AnimationAgent) FireAnimationStartedOnTarget(targetId string, event AnimationStartedEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Animation.animationStarted",
        Params: event,
    })
}
func (agent *AnimationAgent) FireAnimationCanceled(event AnimationCanceledEvent) {
    agent.conn.Send(shared.Notification{
        Method: "Animation.animationCanceled",
        Params: event,
    })
}
func (agent *AnimationAgent) FireAnimationCanceledOnTarget(targetId string, event AnimationCanceledEvent) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "Animation.animationCanceled",
        Params: event,
    })
}

// Commands Sent From Frontend
func (agent *AnimationAgent) EnableNotify() <-chan EnableCommand {
    outChan := make(chan EnableCommand)
    agent.commandChans.Enable = append(agent.commandChans.Enable, outChan)
    return outChan
}
func (agent *AnimationAgent) DisableNotify() <-chan DisableCommand {
    outChan := make(chan DisableCommand)
    agent.commandChans.Disable = append(agent.commandChans.Disable, outChan)
    return outChan
}
func (agent *AnimationAgent) GetPlaybackRateNotify() <-chan GetPlaybackRateCommand {
    outChan := make(chan GetPlaybackRateCommand)
    agent.commandChans.GetPlaybackRate = append(agent.commandChans.GetPlaybackRate, outChan)
    return outChan
}
func (agent *AnimationAgent) SetPlaybackRateNotify() <-chan SetPlaybackRateCommand {
    outChan := make(chan SetPlaybackRateCommand)
    agent.commandChans.SetPlaybackRate = append(agent.commandChans.SetPlaybackRate, outChan)
    return outChan
}
func (agent *AnimationAgent) GetCurrentTimeNotify() <-chan GetCurrentTimeCommand {
    outChan := make(chan GetCurrentTimeCommand)
    agent.commandChans.GetCurrentTime = append(agent.commandChans.GetCurrentTime, outChan)
    return outChan
}
func (agent *AnimationAgent) SetPausedNotify() <-chan SetPausedCommand {
    outChan := make(chan SetPausedCommand)
    agent.commandChans.SetPaused = append(agent.commandChans.SetPaused, outChan)
    return outChan
}
func (agent *AnimationAgent) SetTimingNotify() <-chan SetTimingCommand {
    outChan := make(chan SetTimingCommand)
    agent.commandChans.SetTiming = append(agent.commandChans.SetTiming, outChan)
    return outChan
}
func (agent *AnimationAgent) SeekAnimationsNotify() <-chan SeekAnimationsCommand {
    outChan := make(chan SeekAnimationsCommand)
    agent.commandChans.SeekAnimations = append(agent.commandChans.SeekAnimations, outChan)
    return outChan
}
func (agent *AnimationAgent) ReleaseAnimationsNotify() <-chan ReleaseAnimationsCommand {
    outChan := make(chan ReleaseAnimationsCommand)
    agent.commandChans.ReleaseAnimations = append(agent.commandChans.ReleaseAnimations, outChan)
    return outChan
}
func (agent *AnimationAgent) ResolveAnimationNotify() <-chan ResolveAnimationCommand {
    outChan := make(chan ResolveAnimationCommand)
    agent.commandChans.ResolveAnimation = append(agent.commandChans.ResolveAnimation, outChan)
    return outChan
}
func init() {

}
