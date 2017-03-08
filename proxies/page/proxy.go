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
    command := <-p.agent.EnableNotify()
    command.Respond()

    resourceTreeCommand := <-p.agent.GetResourceTreeNotify()
    resourceTreeCommand.Respond(&pageAgent.GetResourceTreeReturn{
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
