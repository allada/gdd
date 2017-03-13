package shared

import (
    "io"
    "net/http"
    "golang.org/x/net/websocket"
    "strings"
    "encoding/json"
    "sync/atomic"
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

type Error struct {
    msg string
}

func ThrowError(msg string) {
    panic(Error{msg})
}

func (e Error) Error() string {
    return e.msg
}

type Warning struct {
    msg string
}

func ThrowWarning(msg string) {
    panic(Warning{msg})
}

func (w Warning) Error() string {
    return w.msg
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

type Commander interface {
    Initalize(targetId string, responseId int64, conn *Connection)
}

type Connection struct {
    ws *websocket.Conn
    agents map[string]Agenter
    closed int32 // Since Go does not have atomic_flag I use int32
}

func (c *Connection) Close() {
    // This ensures we only execute stuff after this if statment once per agent.
    if !atomic.CompareAndSwapInt32(&c.closed, 0, 1) {
        return
    }
    fmt.Println("Closing socket")
    c.ws.Close()
}

func (c *Connection) RegisterAgent(agent Agenter) {
    // TODO This is not thread safe.
    c.agents[agent.Name()] = agent
}

func (c *Connection) Closed() bool {
    return atomic.LoadInt32(&c.closed) == 1
}

func (c *Connection) SetupCommand(id int64, targetId string, data *json.RawMessage, out Commander) {
    if data != nil {
        if err := json.Unmarshal(*data, out); err != nil {
            fmt.Println("Closing connection because malformed json was sent.")
            c.Close()
            return
        }
    }
    out.Initalize(targetId, id, c)
}

func (c *Connection) SendToTarget(targetId string, event interface{}) {
    if c.Closed() {
        return
    }
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
    if c.Closed() {
        return
    }
    websocket.JSON.Send(c.ws, event)
}

func (c *Connection) SendErrorResult(id int64, targetId string, errorCode ResponseErrorCodes, message string) {
    if c.Closed() {
        return
    }
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
    if c.Closed() {
        return
    }
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

func TryRecoverFromPanic(conn *Connection) {
    data := recover()
    switch data.(type) {
    case nil:
        return
    case Warning:
        fmt.Println(data)
    case Error:
        fmt.Println("Closing websocket because of following Error:")
        fmt.Println(data)
        conn.Close()
    case error:
        fmt.Println("Closing websocket because of following Error:")
        fmt.Println(data)
        conn.Close()
    default:
        fmt.Println("Should probably use shared.Warning or shared.Error instead to panic()")
        panic(data)
    }
}


func WrapFunctionForPanicRecover(fn func(), conn *Connection) func() {
    return func() {
        defer TryRecoverFromPanic(conn)
        fn()
    }
}

func (c *Connection) socketListener() {
    for {
        var message []byte
        err := websocket.Message.Receive(c.ws, &message)
        if c.Closed() {
            return
        } else if err == io.EOF {
            panic(Warning{"Connection closed by remote"})
        } else if err != nil {
            panic(err)
        }
        // fmt.Print("Recv: ")
        // fmt.Println(string(message))

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
            // fmt.Printf("Agent '%s' not registered.\n", agentName)
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
        // Because we cannot call recover() inside another function to recover we do them seperate.
        defer conn.Close()
        defer TryRecoverFromPanic(conn)
        fn(conn)
        conn.socketListener()
    })
}
