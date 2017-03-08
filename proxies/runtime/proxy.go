package runtime

import (
    "fmt"
    "bufio"
    "../../dbgClient"
    "../../protocol/shared"
    runtimeAgent "../../protocol/runtime"
)

type proxy struct {
    agent *runtimeAgent.RuntimeAgent
    client *dbgClient.Client
}

func NewProxy(conn *shared.Connection, client *dbgClient.Client) *proxy {
    agent := runtimeAgent.NewAgent(conn)
    return &proxy{
        agent: agent,
        client: client,
    }
}

func (p *proxy) Agent() *runtimeAgent.RuntimeAgent {
    return p.agent
}

func (p *proxy) Start() {
    // Wait until we are enabled.
    command := <-p.agent.EnableNotify()
    command.Respond()

    go p.handleStdout()
    go p.handleStderr()
}

func (p *proxy) handleStdout() {
    stdout, err := p.client.GetStdout()
    if err != nil {
        panic(err)
    }
    reader := bufio.NewReader(stdout)
    for {
        data, _, err := reader.ReadLine()
        if err != nil {
            panic(err)
        }
        fmt.Println("Stdout: ", string(data))
        p.agent.FireConsoleAPICalled(runtimeAgent.ConsoleAPICalledEvent{
            Type: runtimeAgent.ConsoleAPICalledTypeLog,
            Args: []runtimeAgent.RemoteObject{
                {
                    Type: runtimeAgent.RemoteObjectTypeString,
                    Value: string(data),
                },
            },
            Timestamp: 10000000,
            ExecutionContextId: 5,
        })
    }
}

func (p *proxy) handleStderr() {
    stderr, err := p.client.GetStderr()
    if err != nil {
        panic(err)
    }
    reader := bufio.NewReader(stderr)
    for {
        data, notDone, err := reader.ReadLine()
        if !notDone {
            fmt.Println("stdout too big!")
            return
        }
        if err != nil {
            panic(err)
        }
        p.agent.FireConsoleAPICalled(runtimeAgent.ConsoleAPICalledEvent{
            Type: runtimeAgent.ConsoleAPICalledTypeLog,
            Args: []runtimeAgent.RemoteObject{
                {
                    Type: runtimeAgent.RemoteObjectTypeString,
                    Value: string(data),
                },
            },
            ExecutionContextId: 0,
        })
    }
}
