package page

import (
    "../../dbgClient"
    "../../protocol/shared"
    pageAgent "../../protocol/page"
)

type proxy struct {
    agent *pageAgent.PageAgent
    client *dbgClient.Client
}

func NewProxy(conn *shared.Connection, client *dbgClient.Client) *proxy {
    agent := pageAgent.NewAgent(conn)
    return &proxy{
        agent: agent,
        client: client,
    }
}

func (p *proxy) Agent() *pageAgent.PageAgent {
    return p.agent
}

func (p *proxy) Start() {
    // Wait until we are enabled.
    p.agent.SetEnableHandler(p.enableAndRespond)
}

func (p *proxy) enableAndRespond(command pageAgent.EnableCommand) {
    command.Respond()

    p.agent.SetGetResourceTreeHandler(p.getResourceTreeAndRespond)
}

func (p *proxy) getResourceTreeAndRespond(command pageAgent.GetResourceTreeCommand) {
    command.Respond(&pageAgent.GetResourceTreeReturn{
        FrameTree: pageAgent.FrameResourceTree{
            Frame: pageAgent.Frame{
                Id: "1",
                LoaderId: "1",
                Url: "",
                SecurityOrigin: "",
                MimeType: "",
            },
            Resources: []pageAgent.FrameResource{},
        },
    })
}