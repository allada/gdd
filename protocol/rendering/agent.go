package rendering


import (
    "../shared"
    "encoding/json"
    "fmt"
)
type RenderingAgent struct {
    conn *shared.Connection
    commandChans struct {
        SetShowPaintRects []chan<- SetShowPaintRectsCommand
        SetShowDebugBorders []chan<- SetShowDebugBordersCommand
        SetShowFPSCounter []chan<- SetShowFPSCounterCommand
        SetShowScrollBottleneckRects []chan<- SetShowScrollBottleneckRectsCommand
        SetShowViewportSizeOnResize []chan<- SetShowViewportSizeOnResizeCommand
    }
}
func NewAgent(conn *shared.Connection) *RenderingAgent {
    agent := &RenderingAgent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}

func (agent *RenderingAgent) Name() string {
    return "Rendering"
}

func (agent *RenderingAgent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    switch (funcName) {
        case "setShowPaintRects":
            var out SetShowPaintRectsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetShowPaintRects {
                c <-out
            }
        case "setShowDebugBorders":
            var out SetShowDebugBordersCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetShowDebugBorders {
                c <-out
            }
        case "setShowFPSCounter":
            var out SetShowFPSCounterCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetShowFPSCounter {
                c <-out
            }
        case "setShowScrollBottleneckRects":
            var out SetShowScrollBottleneckRectsCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetShowScrollBottleneckRects {
                c <-out
            }
        case "setShowViewportSizeOnResize":
            var out SetShowViewportSizeOnResizeCommand
            if data != nil {
                if err := json.Unmarshal(*data, &out); err != nil {
                    panic(err)
                }
            }
            out.DestinationTargetID = targetId
            out.responseId = id
            out.conn = agent.conn
            for _, c := range agent.commandChans.SetShowViewportSizeOnResize {
                c <-out
            }
        default:
            fmt.Printf("Command %s unknown\n", funcName)
    }
}

// Dispatchable Events

// Commands Sent From Frontend
func (agent *RenderingAgent) SetShowPaintRectsNotify() <-chan SetShowPaintRectsCommand {
    outChan := make(chan SetShowPaintRectsCommand)
    agent.commandChans.SetShowPaintRects = append(agent.commandChans.SetShowPaintRects, outChan)
    return outChan
}
func (agent *RenderingAgent) SetShowDebugBordersNotify() <-chan SetShowDebugBordersCommand {
    outChan := make(chan SetShowDebugBordersCommand)
    agent.commandChans.SetShowDebugBorders = append(agent.commandChans.SetShowDebugBorders, outChan)
    return outChan
}
func (agent *RenderingAgent) SetShowFPSCounterNotify() <-chan SetShowFPSCounterCommand {
    outChan := make(chan SetShowFPSCounterCommand)
    agent.commandChans.SetShowFPSCounter = append(agent.commandChans.SetShowFPSCounter, outChan)
    return outChan
}
func (agent *RenderingAgent) SetShowScrollBottleneckRectsNotify() <-chan SetShowScrollBottleneckRectsCommand {
    outChan := make(chan SetShowScrollBottleneckRectsCommand)
    agent.commandChans.SetShowScrollBottleneckRects = append(agent.commandChans.SetShowScrollBottleneckRects, outChan)
    return outChan
}
func (agent *RenderingAgent) SetShowViewportSizeOnResizeNotify() <-chan SetShowViewportSizeOnResizeCommand {
    outChan := make(chan SetShowViewportSizeOnResizeCommand)
    agent.commandChans.SetShowViewportSizeOnResize = append(agent.commandChans.SetShowViewportSizeOnResize, outChan)
    return outChan
}
func init() {

}
