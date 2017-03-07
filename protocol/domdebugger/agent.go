package domdebugger


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type DOMDebuggerAgent struct {
    conn *shared.Connection
    commandChans struct {
        SetDOMBreakpoint []chan<- SetDOMBreakpointCommand
        RemoveDOMBreakpoint []chan<- RemoveDOMBreakpointCommand
        SetEventListenerBreakpoint []chan<- SetEventListenerBreakpointCommand
        RemoveEventListenerBreakpoint []chan<- RemoveEventListenerBreakpointCommand
        SetInstrumentationBreakpoint []chan<- SetInstrumentationBreakpointCommand
        RemoveInstrumentationBreakpoint []chan<- RemoveInstrumentationBreakpointCommand
        SetXHRBreakpoint []chan<- SetXHRBreakpointCommand
        RemoveXHRBreakpoint []chan<- RemoveXHRBreakpointCommand
        GetEventListeners []chan<- GetEventListenersCommand
    }
}
func NewAgent(conn *shared.Connection) *DOMDebuggerAgent {
    agent := &DOMDebuggerAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *DOMDebuggerAgent) Name() string {
    return "DOMDebugger"
}

func (agent *DOMDebuggerAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "setDOMBreakpoint":
            var out SetDOMBreakpointCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetDOMBreakpoint {
                c <-out
            }
        case "removeDOMBreakpoint":
            var out RemoveDOMBreakpointCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RemoveDOMBreakpoint {
                c <-out
            }
        case "setEventListenerBreakpoint":
            var out SetEventListenerBreakpointCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetEventListenerBreakpoint {
                c <-out
            }
        case "removeEventListenerBreakpoint":
            var out RemoveEventListenerBreakpointCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RemoveEventListenerBreakpoint {
                c <-out
            }
        case "setInstrumentationBreakpoint":
            var out SetInstrumentationBreakpointCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetInstrumentationBreakpoint {
                c <-out
            }
        case "removeInstrumentationBreakpoint":
            var out RemoveInstrumentationBreakpointCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RemoveInstrumentationBreakpoint {
                c <-out
            }
        case "setXHRBreakpoint":
            var out SetXHRBreakpointCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetXHRBreakpoint {
                c <-out
            }
        case "removeXHRBreakpoint":
            var out RemoveXHRBreakpointCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.RemoveXHRBreakpoint {
                c <-out
            }
        case "getEventListeners":
            var out GetEventListenersCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.GetEventListeners {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events

// Commands Sent From Frontend
func (agent *DOMDebuggerAgent) SetDOMBreakpointNotify() <-chan SetDOMBreakpointCommand {
    outChan := make(chan SetDOMBreakpointCommand)
    agent.commandChans.SetDOMBreakpoint = append(agent.commandChans.SetDOMBreakpoint, outChan)
    return outChan
}
func (agent *DOMDebuggerAgent) RemoveDOMBreakpointNotify() <-chan RemoveDOMBreakpointCommand {
    outChan := make(chan RemoveDOMBreakpointCommand)
    agent.commandChans.RemoveDOMBreakpoint = append(agent.commandChans.RemoveDOMBreakpoint, outChan)
    return outChan
}
func (agent *DOMDebuggerAgent) SetEventListenerBreakpointNotify() <-chan SetEventListenerBreakpointCommand {
    outChan := make(chan SetEventListenerBreakpointCommand)
    agent.commandChans.SetEventListenerBreakpoint = append(agent.commandChans.SetEventListenerBreakpoint, outChan)
    return outChan
}
func (agent *DOMDebuggerAgent) RemoveEventListenerBreakpointNotify() <-chan RemoveEventListenerBreakpointCommand {
    outChan := make(chan RemoveEventListenerBreakpointCommand)
    agent.commandChans.RemoveEventListenerBreakpoint = append(agent.commandChans.RemoveEventListenerBreakpoint, outChan)
    return outChan
}
func (agent *DOMDebuggerAgent) SetInstrumentationBreakpointNotify() <-chan SetInstrumentationBreakpointCommand {
    outChan := make(chan SetInstrumentationBreakpointCommand)
    agent.commandChans.SetInstrumentationBreakpoint = append(agent.commandChans.SetInstrumentationBreakpoint, outChan)
    return outChan
}
func (agent *DOMDebuggerAgent) RemoveInstrumentationBreakpointNotify() <-chan RemoveInstrumentationBreakpointCommand {
    outChan := make(chan RemoveInstrumentationBreakpointCommand)
    agent.commandChans.RemoveInstrumentationBreakpoint = append(agent.commandChans.RemoveInstrumentationBreakpoint, outChan)
    return outChan
}
func (agent *DOMDebuggerAgent) SetXHRBreakpointNotify() <-chan SetXHRBreakpointCommand {
    outChan := make(chan SetXHRBreakpointCommand)
    agent.commandChans.SetXHRBreakpoint = append(agent.commandChans.SetXHRBreakpoint, outChan)
    return outChan
}
func (agent *DOMDebuggerAgent) RemoveXHRBreakpointNotify() <-chan RemoveXHRBreakpointCommand {
    outChan := make(chan RemoveXHRBreakpointCommand)
    agent.commandChans.RemoveXHRBreakpoint = append(agent.commandChans.RemoveXHRBreakpoint, outChan)
    return outChan
}
func (agent *DOMDebuggerAgent) GetEventListenersNotify() <-chan GetEventListenersCommand {
    outChan := make(chan GetEventListenersCommand)
    agent.commandChans.GetEventListeners = append(agent.commandChans.GetEventListeners, outChan)
    return outChan
}
func init() {

}
