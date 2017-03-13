package dbgClient

import (
    "io"
    "os/exec"
    "fmt"
    "../config"
    "unsafe"
    "sync"
    "sync/atomic"
    "github.com/derekparker/delve/service/rpc2"
)

const API_PORT_LISTENING_STRING string = "API server listening at: 127.0.0.1:%d\n"

type ProcessState int32
const (
    StartingState = iota
    RunningState
    KilledState
)

type Client struct{
    File string
    Args []string
    isReadyLock sync.Mutex
    rpcClient *rpc2.RPCClient
    cmd *exec.Cmd
    runningState ProcessState // Since Go does not have atomic_flag I use int32
    stdout io.ReadCloser
    stderr io.ReadCloser
    stdin io.WriteCloser
}

func (c *Client) Start() {
    defer func() {
        if r := recover(); r != nil {
            c.Kill()
        }
        c.isReadyLock.Unlock()
    }()
    c.isReadyLock.Lock()
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
    // Looks like we were told to die before we started. Lets die now.
    if c.Killed() {
        return
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

    c.stdin, err = c.cmd.StdinPipe()
    if err != nil {
        panic(err)
    }

    if err = c.cmd.Start(); err != nil {
        panic(err)
    }

    var port int
    fmt.Fscanf(c.stdout, API_PORT_LISTENING_STRING, &port)
    if port == 0 {
        panic("Did not get good output from dvl or 0 for port number.")
    }

    c.rpcClient = rpc2.NewClient("localhost:" + fmt.Sprintf("%d", port))
    if ProcessState(atomic.SwapInt32((*int32)(&c.runningState), RunningState)) == KilledState {
        // Looks like we got a kill signal while we were starting.
        c.Kill()
        return
    }

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
}

func (c *Client) Starting() bool {
    return ProcessState(atomic.LoadInt32((*int32)(&c.runningState))) == StartingState
}

func (c *Client) Killed() bool {
    return ProcessState(atomic.LoadInt32((*int32)(&c.runningState))) == KilledState
}

func (c *Client) Kill() {
    // This ensures we only execute stuff after this if statment once per agent.
    if atomic.SwapInt32((*int32)(&c.runningState), KilledState) != RunningState {
        return
    }
    fmt.Println("Killing debugging process")
    c.Detach(false)
    c.stdout.Close()
    c.stderr.Close()
    c.stdin.Close()
    c.cmd.Process.Kill()
    c.cmd.Wait()
    fmt.Println("Killed debugging process")
}

func (c *Client) GetStdout() (io.ReadCloser, error) {
    c.BlockUntilReady()
    return c.stdout, nil
}

func (c *Client) GetStderr() (io.ReadCloser, error) {
    c.BlockUntilReady()
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