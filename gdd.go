package main

import (
    "os"
    "net"
    "net/http"
    "time"
    "github.com/allada/gdd/protocol/shared"
    "github.com/allada/gdd/proxies/debugger"
    "github.com/allada/gdd/proxies/page"
    "github.com/allada/gdd/proxies/runtime"
    "github.com/allada/gdd/dbgClient"
    "github.com/allada/gdd/config"
    "fmt"
)

type handler struct {
    Config config.Config
}

func (h handler) handleDebugRequest(conn *shared.Connection) {
    client := &dbgClient.Client{
        File: h.Config.DebugSession.File,
        Args: h.Config.DebugSession.Args,
    }
    go client.Start(h.Config)

    go func() {
        for {
            // This just ensures we syncronize our sockets if either dies, they both die.
            if client.Killed() || conn.Closed() {
                client.Kill()
                conn.Close()
                return
            }
            time.Sleep(300 * time.Millisecond)
        }
    }()

    runtimeProxy := runtime.NewProxy(conn, client)
    go runtimeProxy.Start()
    go debugger.NewProxy(conn, client).Start(runtimeProxy)
    go page.NewProxy(conn, client).Start()
}

func printHelp() {

}

func configureFromArgs(args []string) {

}

func main() {
    config := config.NewFromArgs(os.Args[1:])
    if config == nil {
        return
    }
    h := handler{
        Config: *config,
    }

    http.Handle("/", shared.Handler(h.handleDebugRequest))
    listener, err := net.Listen("tcp", ":" + (*config).Port);
    if err != nil {
        panic("Listen: " + err.Error())
    }
    fmt.Println("https://chrome-devtools-frontend.appspot.com/serve_file/@3f9be266524354849e17f6a123ee20dd6e7d53e0/inspector.html?experiments=true&ws=localhost:9922/go-devtools-debug")
    err = http.Serve(listener, nil)
    if err != nil {
        panic("Serve: " + err.Error())
    }
}
