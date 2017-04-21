package dbgClient

import (
    "unsafe"
    "github.com/allada/delve/service/api"
)

func (c *Client) ListGoroutines() ([]*Goroutine, error) {
    goRoutines, err := c.rpcClient.ListGoroutines()
    return *(*[]*Goroutine)(unsafe.Pointer(&goRoutines)), err
}

func (c *Client) ListThreads() ([]*Thread, error) {
    threads, err := c.rpcClient.ListThreads()
    return *(*[]*Thread)(unsafe.Pointer(&threads)), err
}

func (c *Client) Stacktrace(goroutineID int, depth int, cfg *LoadConfig) ([]Stackframe, error) {
    stack, err := c.rpcClient.Stacktrace(goroutineID, depth, (*api.LoadConfig)(cfg))
    return *(*[]Stackframe)(unsafe.Pointer(&stack)), err
}

func (c *Client) GetState() (*DebuggerState, error) {
    debuggerState, err := c.rpcClient.GetState()
    return (*DebuggerState)(debuggerState), err
}

func (c *Client) Continue() <-chan *DebuggerState {
    debuggerState := c.rpcClient.Continue()
    return *(*<-chan *DebuggerState)(unsafe.Pointer(&debuggerState))
}

func (c *Client) Next() (*DebuggerState, error) {
    debuggerState, err := c.rpcClient.Next()
    return (*DebuggerState)(debuggerState), err
}

func (c *Client) Step() (*DebuggerState, error) {
    debuggerState, err := c.rpcClient.Step()
    return (*DebuggerState)(debuggerState), err
}

func (c *Client) StepOut() (*DebuggerState, error) {
    debuggerState, err := c.rpcClient.StepOut()
    return (*DebuggerState)(debuggerState), err
}

func (c *Client) SwitchGoroutine(goroutineID int) (*DebuggerState, error) {
    debuggerState, err := c.rpcClient.SwitchGoroutine(goroutineID)
    return (*DebuggerState)(debuggerState), err
}

func (c *Client) SwitchThread(threadID int) (*DebuggerState, error) {
    debuggerState, err := c.rpcClient.SwitchThread(threadID)
    return (*DebuggerState)(debuggerState), err
}

func (c *Client) ListLocalVariables(scope EvalScope, cfg LoadConfig) ([]Variable, error) {
    variables, err := c.rpcClient.ListLocalVariables(api.EvalScope(scope), api.LoadConfig(cfg))
    return *(*[]Variable)(unsafe.Pointer(&variables)), err
}

func (c *Client) EvalVariable(scope EvalScope, expr string, cfg LoadConfig) (*Variable, error) {
    variable, err := c.rpcClient.EvalVariable(api.EvalScope(scope), expr, api.LoadConfig(cfg))
    return (*Variable)(variable), err
}
