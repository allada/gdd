package input


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type InputAgent struct {
    conn *shared.Connection
    commandChans struct {
        DispatchKeyEvent []chan<- DispatchKeyEventCommand
        DispatchMouseEvent []chan<- DispatchMouseEventCommand
        DispatchTouchEvent []chan<- DispatchTouchEventCommand
        EmulateTouchFromMouseEvent []chan<- EmulateTouchFromMouseEventCommand
        SynthesizePinchGesture []chan<- SynthesizePinchGestureCommand
        SynthesizeScrollGesture []chan<- SynthesizeScrollGestureCommand
        SynthesizeTapGesture []chan<- SynthesizeTapGestureCommand
    }
}
func NewAgent(conn *shared.Connection) *InputAgent {
    agent := &InputAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *InputAgent) Name() string {
    return "Input"
}

func (agent *InputAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "dispatchKeyEvent":
            var out DispatchKeyEventCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.DispatchKeyEvent {
                c <-out
            }
        case "dispatchMouseEvent":
            var out DispatchMouseEventCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.DispatchMouseEvent {
                c <-out
            }
        case "dispatchTouchEvent":
            var out DispatchTouchEventCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.DispatchTouchEvent {
                c <-out
            }
        case "emulateTouchFromMouseEvent":
            var out EmulateTouchFromMouseEventCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.EmulateTouchFromMouseEvent {
                c <-out
            }
        case "synthesizePinchGesture":
            var out SynthesizePinchGestureCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SynthesizePinchGesture {
                c <-out
            }
        case "synthesizeScrollGesture":
            var out SynthesizeScrollGestureCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SynthesizeScrollGesture {
                c <-out
            }
        case "synthesizeTapGesture":
            var out SynthesizeTapGestureCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SynthesizeTapGesture {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events

// Commands Sent From Frontend
func (agent *InputAgent) DispatchKeyEventNotify() <-chan DispatchKeyEventCommand {
    outChan := make(chan DispatchKeyEventCommand)
    agent.commandChans.DispatchKeyEvent = append(agent.commandChans.DispatchKeyEvent, outChan)
    return outChan
}
func (agent *InputAgent) DispatchMouseEventNotify() <-chan DispatchMouseEventCommand {
    outChan := make(chan DispatchMouseEventCommand)
    agent.commandChans.DispatchMouseEvent = append(agent.commandChans.DispatchMouseEvent, outChan)
    return outChan
}
func (agent *InputAgent) DispatchTouchEventNotify() <-chan DispatchTouchEventCommand {
    outChan := make(chan DispatchTouchEventCommand)
    agent.commandChans.DispatchTouchEvent = append(agent.commandChans.DispatchTouchEvent, outChan)
    return outChan
}
func (agent *InputAgent) EmulateTouchFromMouseEventNotify() <-chan EmulateTouchFromMouseEventCommand {
    outChan := make(chan EmulateTouchFromMouseEventCommand)
    agent.commandChans.EmulateTouchFromMouseEvent = append(agent.commandChans.EmulateTouchFromMouseEvent, outChan)
    return outChan
}
func (agent *InputAgent) SynthesizePinchGestureNotify() <-chan SynthesizePinchGestureCommand {
    outChan := make(chan SynthesizePinchGestureCommand)
    agent.commandChans.SynthesizePinchGesture = append(agent.commandChans.SynthesizePinchGesture, outChan)
    return outChan
}
func (agent *InputAgent) SynthesizeScrollGestureNotify() <-chan SynthesizeScrollGestureCommand {
    outChan := make(chan SynthesizeScrollGestureCommand)
    agent.commandChans.SynthesizeScrollGesture = append(agent.commandChans.SynthesizeScrollGesture, outChan)
    return outChan
}
func (agent *InputAgent) SynthesizeTapGestureNotify() <-chan SynthesizeTapGestureCommand {
    outChan := make(chan SynthesizeTapGestureCommand)
    agent.commandChans.SynthesizeTapGesture = append(agent.commandChans.SynthesizeTapGesture, outChan)
    return outChan
}
func init() {

}
