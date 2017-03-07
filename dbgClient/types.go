package dbgClient

import (
    "github.com/derekparker/delve/service/api"
)

type EvalScope api.EvalScope

func (a EvalScope) conv() api.EvalScope {
    return api.EvalScope(a)
}

type Location api.Location

func (a Location) conv() api.Location {
    return api.Location(a)
}

type Breakpoint api.Breakpoint

func (a Breakpoint) conv() api.Breakpoint {
    return api.Breakpoint(a)
}

type DebuggerState api.DebuggerState

func (a DebuggerState) conv() api.DebuggerState {
    return api.DebuggerState(a)
}

type Goroutine api.Goroutine

func (a Goroutine) conv() api.Goroutine {
    return api.Goroutine(a)
}

type LoadConfig api.LoadConfig

func (a LoadConfig) conv() api.LoadConfig {
    return api.LoadConfig(a)
}

type Stackframe api.Stackframe

func (a Stackframe) conv() api.Stackframe {
    return api.Stackframe(a)
}

type Thread api.Thread

func (a Thread) conv() api.Thread {
    return api.Thread(a)
}
