package shared

import (
    "io"
    "net/http"
    "golang.org/x/net/websocket"
    "strings"
    "encoding/json"
    "fmt"
)

func init() {
}

type ResponseErrorCodes int64
const (
    ErrorCodeParseError = ResponseErrorCodes(-32700)
    ErrorCodeInvalidRequest = ResponseErrorCodes(-32600)
    ErrorCodeMethodNotFound = ResponseErrorCodes(-32601)
    ErrorCodeInvalidParams = ResponseErrorCodes(-32602)
    ErrorCodeInternalError = ResponseErrorCodes(-32603)
    ErrorCodeServerError = ResponseErrorCodes(-32000)
)

type ResponseStatusCodes int64
const (
    StatusCodeSuccess = ResponseStatusCodes(0)
    StatusCodeError = ResponseStatusCodes(1)
    StatusCodeFallThrough = ResponseStatusCodes(2)
    StatusCodeAsync = ResponseStatusCodes(3)
)

type ResponseError struct {
    Code ResponseErrorCodes `json:"code"`
    Message string `json:"message"`
}

type Notification struct {
    Method string `json:"method"`
    Params interface{} `json:"params,omitempty"`
}

type Response struct {
    Id int64 `json:"id"`
    Result interface{} `json:"result"`
    Status ResponseStatusCodes `json:"status,omitempty"`
    Error *ResponseError `json:"error,omitempty"`
}

type ReceivedCommand struct {
    Id int64 `json:"id"`
    Method string `json:"method"`
    Params *json.RawMessage `json:"params,omitempty"`
}

func (rc *ReceivedCommand) agentAndCommand() (agent string, command string, ok bool) {
    parts := strings.Split(rc.Method, ".")
    if len(parts) < 2 {
        return "", "", false
    }
    return parts[0], parts[1], true
}

type Agenter interface {
    Name() string
    ProcessCommand(int64, string, string, *json.RawMessage)
}

type Connection struct {
    ws *websocket.Conn
    agents map[string]Agenter
}

func (c *Connection) RegisterAgent(agent Agenter) {
    // TODO This is not thread safe.
    c.agents[agent.Name()] = agent
}

func (c *Connection) SendToTarget(targetId string, event interface{}) {
    //message, err := json.Marshal(event)
    //if err != nil {
    //    panic(err)
    //}
    data := struct {
        TargetId string `json:"targetId"`
        Message interface{} `json:"message"`
    }{
        TargetId: targetId,
        Message: event,
    }
    c.Send(Notification{
        Method: "Target.receivedMessageFromTarget",
        Params: data,
    })
}

func (c *Connection) Send(event Notification) {
    websocket.JSON.Send(c.ws, event)
}

func (c *Connection) SendErrorResult(id int64, targetId string, errorCode ResponseErrorCodes, message string) {
    data := Response{
        Id: id,
        Status: StatusCodeError,
        Error: &ResponseError{
            Code: errorCode,
            Message: message,
        },
    }
    if targetId == "" {
        websocket.JSON.Send(c.ws, data)
        return
    }
    // If there is a target send it via receivedMessageFromTarget.
    c.SendToTarget(targetId, data)
}

func (c *Connection) SendResult(id int64, targetId string, result interface{}) {
    data := Response{
        Id: id,
        Result: result,
    }
    if targetId == "" {
        websocket.JSON.Send(c.ws, data)
        return
    }
    // If there is a target send it via receivedMessageFromTarget.
    c.SendToTarget(targetId, data)
}

func (c *Connection) socketListener() {
    defer c.ws.Close()
    for {
        var message []byte
        err := websocket.Message.Receive(c.ws, &message)
        if err == io.EOF {
            fmt.Println("Connection closed by remote")
            return
        } else if err != nil {
            panic(err)
        }
        fmt.Print("Recv: ")
        fmt.Println(string(message))

        var receivedCommand ReceivedCommand
        err = json.Unmarshal(message, &receivedCommand)
        if err != nil {
            panic(err)
            return
        }

        params := receivedCommand.Params
        targetId := ""
        // Special exception for Target.
        if receivedCommand.Method == "Target.sendMessageToTarget" {
            var targetData struct {
                TargetId string `json:"targetId"`
                Message string `json:"message"`
            }
            if receivedCommand.Params == nil {
                panic("Expected params to be sent in json message.")
            }
            if err := json.Unmarshal(*params, &targetData); err != nil {
                panic(err)
            }
            // TODO Send a "received" command back.
            targetId = targetData.TargetId
            err = json.Unmarshal([]byte(targetData.Message), &receivedCommand)
            if err != nil {
                panic(err)
            }
            // fmt.Print("Recv Target: ")
            // fmt.Println(targetData.Message)
        }

        agentName, command, ok := receivedCommand.agentAndCommand()
        if !ok {
            fmt.Println("Unknown Agent and/or Command")
            return;
        }
        var agent Agenter
        agent, ok = c.agents[agentName]
        if !ok {
            fmt.Printf("Agent '%s' not registered.\n", agentName)
            continue
        }
        go agent.ProcessCommand(receivedCommand.Id, targetId, command, params)
    }
}

func Handler(fn func(*Connection)) http.Handler {
    return websocket.Handler(func(ws *websocket.Conn) {
        conn := &Connection{
            ws: ws,
            agents: make(map[string]Agenter),
        }
        fn(conn)
        conn.socketListener()
    })
}

