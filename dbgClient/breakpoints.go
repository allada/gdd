package dbgClient

import (
    "unsafe"
    "github.com/allada/delve/service/api"
)

// PC: Program counter
func (c *Client) CreateBreakpointAtPC(pc uint64) {
    c.rpcClient.CreateBreakpoint(&api.Breakpoint{
        Addr: pc,
    })
}

func (c *Client) CreateBreakpointAtLine(file string, line int, name string) (*Breakpoint, error) {
    breakpoint, err := c.rpcClient.CreateBreakpoint(&api.Breakpoint{
        File: file,
        Line: line,
        Name: name,
    })
    return (*Breakpoint)(breakpoint), err
}

func (c *Client) ListAllBreakpoints() ([]*Breakpoint, error) {
    breakpoints, err := c.rpcClient.ListBreakpoints()
    // This pattern is here because we cannot convert between slices of same underlying types but different toplevel types.
    return *(*[]*Breakpoint)(unsafe.Pointer(&breakpoints)), err
}

func (c *Client) ClearAllBreakpoints() {
    breakpoints, err := c.rpcClient.ListBreakpoints()
    if err != nil {
        panic(err)
    }
    for _, breakpoint := range breakpoints {
        err := c.ClearBreakpoint(breakpoint.ID)
        if err != nil {
            panic(err)
        }
    }
}

func (c *Client) ClearBreakpoint(id int) error {
    _, err := c.rpcClient.ClearBreakpoint(id)
    return err
}

func (c *Client) ClearBreakpointByName(name string) error {
    _, err := c.rpcClient.ClearBreakpointByName(name)
    return err
}
