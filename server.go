package main

import (
    //"fmt"
    "net/http"
    "./protocol/shared"
    "./proxies/debugger"
    "./proxies/page"
    "./proxies/runtime"
    "./dbgClient"
)

func handleDebugRequest(conn *shared.Connection) {
    client := &dbgClient.Client{
        File: "./helloworld.go",
    }
    runtime := runtime.NewProxy(conn, client)
    go client.Start()
    go runtime.Start()
    go debugger.NewProxy(conn, client).Start(runtime.Agent())
    go page.NewProxy(conn, client).Start()
    // go log.NewProxy(conn, client).Start()
}

func main() {
    http.Handle("/", shared.Handler(handleDebugRequest))
    err := http.ListenAndServe(":9922", nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}
