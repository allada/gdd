package runtime

import (
    "fmt"
    "bufio"
    "strconv"
    "strings"
    "reflect"
    "../../dbgClient"
    "../../protocol/shared"
    runtimeAgent "../../protocol/runtime"
)

type proxy struct {
    agent *runtimeAgent.RuntimeAgent
    client *dbgClient.Client
}

func NewProxy(conn *shared.Connection, client *dbgClient.Client) *proxy {
    agent := runtimeAgent.NewAgent(conn)
    return &proxy{
        agent: agent,
        client: client,
    }
}

func (p *proxy) Start() {
    // Wait until we are enabled.
    enableChan := p.agent.EnableNotify()
    command := <-enableChan
    command.Respond()

    go p.handleNotifications()
    go p.handleStdout()
    go p.handleStderr()
}

func (p *proxy) CreateContext() {
    p.agent.FireExecutionContextCreated(runtimeAgent.ExecutionContextCreatedEvent{
        Context: runtimeAgent.ExecutionContextDescription{
            Id: 1,
            Origin: "://",
            Name: "Self",
        },
    })
}

func (p *proxy) handleNotifications() {
    enable       := p.agent.EnableNotify()
    getProperies := p.agent.GetPropertiesNotify()
    compileScript := p.agent.CompileScriptNotify()
    for {
        select {
        case command := <-enable:
            command.Respond()
        case command := <-getProperies:
            go p.getPropertiesAndRespond(command)
        case command := <-compileScript:
            command.Respond(nil)
        }
    }
}

func (p *proxy) getPropertiesAndRespond(command runtimeAgent.GetPropertiesCommand) {
    objectId := string(command.ObjectId)
    if strings.HasPrefix(objectId, "local:") {
        var goroutineID int
        var err error
        if command.DestinationTargetID == "" {
            state, err := p.client.GetState()
            if err != nil {
                panic(err)
            }
            if state.SelectedGoroutine == nil {
                panic("No active goroutine")
            }
            goroutineID = int(state.SelectedGoroutine.ID)
        } else {
            goroutineID, err = strconv.Atoi(command.DestinationTargetID)
            if err != nil {
                panic(err)
            }
        }
        frameId, err := strconv.Atoi(objectId[len("local:"):])
        if err != nil {
            panic(err)
        }
        variables, err := p.client.ListLocalVariables(dbgClient.EvalScope{
            GoroutineID: goroutineID,
            Frame: frameId,
        }, dbgClient.LoadConfig{
            FollowPointers: true,
            MaxVariableRecurse: 1,
            MaxStringLen: 500,
            MaxArrayValues: 1,
            MaxStructFields: 1,
        })
        if err != nil {
            panic(err)
        }
        properties := []runtimeAgent.PropertyDescriptor{}
        for _, variable := range variables {
            remoteObject := p.MakeRemoteObject(variable)
            properties = append(properties, runtimeAgent.PropertyDescriptor{
                Name: variable.Name,
                Value: &remoteObject,
            })
        }
        command.Respond(&runtimeAgent.GetPropertiesReturn{
            Result: properties,
        })
    }
}

func (p *proxy) handleStdout() {
    stdout, err := p.client.GetStdout()
    if err != nil {
        panic(err)
    }
    reader := bufio.NewReader(stdout)
    for {
        data, _, err := reader.ReadLine()
        if err != nil {
            panic(err)
        }
        fmt.Println("Stdout: ", string(data))
        p.agent.FireConsoleAPICalled(runtimeAgent.ConsoleAPICalledEvent{
            Type: runtimeAgent.ConsoleAPICalledTypeLog,
            Args: []runtimeAgent.RemoteObject{
                {
                    Type: runtimeAgent.RemoteObjectTypeString,
                    Value: string(data),
                },
            },
            Timestamp: 10000000,
            ExecutionContextId: 5,
        })
    }
}

func (p *proxy) handleStderr() {
    stderr, err := p.client.GetStderr()
    if err != nil {
        panic(err)
    }
    reader := bufio.NewReader(stderr)
    for {
        data, notDone, err := reader.ReadLine()
        if !notDone {
            fmt.Println("stdout too big!")
            return
        }
        if err != nil {
            panic(err)
        }
        p.agent.FireConsoleAPICalled(runtimeAgent.ConsoleAPICalledEvent{
            Type: runtimeAgent.ConsoleAPICalledTypeLog,
            Args: []runtimeAgent.RemoteObject{
                {
                    Type: runtimeAgent.RemoteObjectTypeString,
                    Value: string(data),
                },
            },
            ExecutionContextId: 0,
        })
    }
}

func (p *proxy) MakeRemoteObject(variable dbgClient.Variable) runtimeAgent.RemoteObject {
    //name := variable.Name
    kind := variable.Kind
    outKind := runtimeAgent.RemoteObjectTypeSymbol
    var subType runtimeAgent.RemoteObjectSubtypeEnum
    var subTypePtr *runtimeAgent.RemoteObjectSubtypeEnum

    previewType := runtimeAgent.ObjectPreviewTypeSymbol
    var previewSubType runtimeAgent.ObjectPreviewSubtypeEnum
    var previewSubTypePtr *runtimeAgent.ObjectPreviewSubtypeEnum
    if kind == reflect.Bool {
        outKind = runtimeAgent.RemoteObjectTypeBoolean
    } else if kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 ||
              kind == reflect.Int64 || kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 ||
              kind == reflect.Uint32 || kind == reflect.Uint64 || kind == reflect.Uintptr || kind == reflect.Float32 ||
              kind == reflect.Float64 {
        outKind = runtimeAgent.RemoteObjectTypeNumber
    } else if kind == reflect.Complex64 || kind == reflect.Complex128 {
    } else if kind == reflect.Array {
        outKind = runtimeAgent.RemoteObjectTypeObject
        subType = runtimeAgent.RemoteObjectSubtypeArray
        subTypePtr = &subType

        previewType = runtimeAgent.ObjectPreviewTypeObject
        previewSubType = runtimeAgent.ObjectPreviewSubtypeArray
        previewSubTypePtr = &previewSubType
    } else if kind == reflect.Chan {
    } else if kind == reflect.Func {
        outKind = runtimeAgent.RemoteObjectTypeFunction
    } else if kind == reflect.Interface {
    } else if kind == reflect.Map {
        outKind = runtimeAgent.RemoteObjectTypeObject
        subType = runtimeAgent.RemoteObjectSubtypeMap
        subTypePtr = &subType
    } else if kind == reflect.Ptr {
    } else if kind == reflect.Slice {
        outKind = runtimeAgent.RemoteObjectTypeObject
        subType = runtimeAgent.RemoteObjectSubtypeArray
        subTypePtr = &subType
    } else if kind == reflect.String {
        outKind = runtimeAgent.RemoteObjectTypeString
        previewType = runtimeAgent.ObjectPreviewTypeString
    } else if kind == reflect.Struct {
        outKind = runtimeAgent.RemoteObjectTypeObject
    } else if kind == reflect.UnsafePointer {
    }

    return runtimeAgent.RemoteObject{
        Type: outKind,
        Subtype: subTypePtr,
        Value: variable.Value,
        Preview: &runtimeAgent.ObjectPreview{
            Type: previewType,
            Subtype: previewSubTypePtr,
            Overflow: false,
            Properties: []runtimeAgent.PropertyPreview{},
            // Entries: &[]runtimeAgent.EntryPreview{},
        },
    }
}
