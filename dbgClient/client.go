package dbgClient

import (
    "io"
    "os/exec"
    "fmt"
    "../config"
    "unsafe"
    "sync"
    "github.com/derekparker/delve/service/rpc2"
)

const API_PORT_LISTENING_STRING string = "API server listening at: 127.0.0.1:%d\n"

type Client struct{
    File string
    Args []string
    isReady bool
    isReadyLock sync.Mutex
    rpcClient *rpc2.RPCClient
    cmd *exec.Cmd
    stdout io.ReadCloser
    stderr io.ReadCloser
}

func (c *Client) Start() {
    c.isReadyLock.Lock()
    defer c.isReadyLock.Unlock()
    args := []string{
        "debug",
        "--headless",
        "--build-flags=-gcflags='-N -l'",
        "--api-version=2",
        "--listen=localhost:0",
        c.File,
    }
    if len(args) > 0 {
        args = append(args, "--")
        args = append(args, c.Args...)
    }
    c.cmd = exec.Command(config.DVL_PATH, args...)

    var err error
    c.stdout, err = c.cmd.StdoutPipe()
    if err != nil {
        panic(err)
    }

    c.stderr, err = c.cmd.StderrPipe()
    if err != nil {
        panic(err)
    }

    if err := c.cmd.Start(); err != nil {
        panic(err)
    }

    var port int
    fmt.Fscanf(c.stdout, API_PORT_LISTENING_STRING, &port)
    if port == 0 {
        panic("Did not get good output from dvl or 0 for port number.")
    }

    c.rpcClient = rpc2.NewClient("localhost:" + fmt.Sprintf("%d", port))
    c.setupBreakOnStart()
    // Continue to newly created breakpoint. We should now be at first instruction in main().
    <-c.Continue()
    breakpoints, err := c.ListAllBreakpoints()
    if err != nil {
        panic(err)
    }
    for _, breakpoint := range breakpoints {
        c.ClearBreakpoint((*breakpoint).ID)
    }
    c.isReady = true
}

func (c *Client) GetStdout() (io.ReadCloser, error) {
    c.BlockUntilReady()
    // This is done so it errors properly.
    defer func() {
        c.stdout = nil
    }()
    return c.stdout, nil
}

func (c *Client) GetStderr() (io.ReadCloser, error) {
    c.BlockUntilReady()
    // This is done so it errors properly.
    defer func() {
        c.stderr = nil
    }()
    return c.stderr, nil
}

func (c *Client) setupBreakOnStart() {
    scope, err := c.FindLocation(EvalScope{
        GoroutineID: -1,
        Frame: 0,
    }, "main.main")
    if err != nil {
        panic(err)
    }
    if len(scope) < 1 {
        panic("Expected scope to have at least 1 item in it.")
    }
    c.CreateBreakpointAtPC(scope[0].PC)
}

func (c *Client) BlockUntilReady() {
    if c.isReady {
        return
    }
    c.isReadyLock.Lock()
    c.isReadyLock.Unlock()
}

func (c *Client) FindLocation(scope EvalScope, location string) ([]Location, error) {
    loc, err := c.rpcClient.FindLocation(scope.conv(), location)
    // This pattern is here because we cannot convert between slices of same underlying types but different toplevel types.
    return *(*[]Location)(unsafe.Pointer(&loc)), err
}

func (c *Client) ListSources() ([]string, error) {
    return c.rpcClient.ListSources("")
}