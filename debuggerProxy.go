package main

import (
    "./protocol/shared"
    "./protocol/debugger"
)

type proxy struct {
    agent Agent
}

func newDebuggerProxy(conn *shared.Connection) *proxy {
    agent := debugger.NewAgent(conn)
    return &proxy{
        agent: agent,
    }
}

func (p *proxy) Start() {
    // Wait until we are enabled
    <-p.agent.EnableNotify()
}
