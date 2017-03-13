package main

import (
    "os"
    "strconv"
    "net/http"
    "./protocol/shared"
    "./proxies/debugger"
    "./proxies/page"
    "./proxies/runtime"
    "./dbgClient"
)

type handler struct {
    File string
    Args []string
}

func (h handler) handleDebugRequest(conn *shared.Connection) {
    client := &dbgClient.Client{
        File: h.File,
        Args: h.Args,
    }
    go client.Start()

    runtimeProxy := runtime.NewProxy(conn, client)
    go runtimeProxy.Start()
    go debugger.NewProxy(conn, client).Start(runtimeProxy)
    go page.NewProxy(conn, client).Start()
}

func main() {
    // TODO This is ugly and should be cleaned up.
    port := "9922"
    file := "./helloworld.go"
    args := os.Args[1:]
    if len(os.Args) > 1 {
        if len(os.Args) > 2 {
            file = os.Args[2]
        }
        if _, err := strconv.Atoi(os.Args[1]); err != nil {
            panic("Expected first argument to be a valid port number if it exists")
        }
        port = os.Args[1]
        args = os.Args[2:]
    }
    h := handler{
        File: file,
        Args: args,
    }
    http.Handle("/", shared.Handler(h.handleDebugRequest))
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}
