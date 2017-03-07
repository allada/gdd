package main

import (
    //"fmt"
    "net/http"
    "./protocol/shared"
    "./proxies/debugger"
    "./dbgClient"
)

func handleDebugRequest(conn *shared.Connection) {
    client := &dbgClient.Client{
        File: "./helloworld.go",
    }
    go client.Start()
    go debugger.NewProxy(conn, client).Start()
}

func main() {
    http.Handle("/", shared.Handler(handleDebugRequest))
    err := http.ListenAndServe(":9922", nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}