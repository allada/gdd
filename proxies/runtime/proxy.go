package runtime

import (
    "fmt"
    "bufio"
    "strconv"
    "strings"
    "reflect"
    "sync/atomic"
    "github.com/allada/gdd/dbgClient"
    "github.com/allada/gdd/protocol/shared"
    runtimeAgent "github.com/allada/gdd/protocol/runtime"
)

type proxy struct {
    agent *runtimeAgent.RuntimeAgent
    client *dbgClient.Client
    conn *shared.Connection

    enabled int32 // Since Go does not have atomic_flag I use int32
}

func NewProxy(conn *shared.Connection, client *dbgClient.Client) *proxy {
    agent := runtimeAgent.NewAgent(conn)
    return &proxy{
        conn: conn,
        agent: agent,
        client: client,
    }
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

func (p *proxy) Start() {
    // Wait until we are enabled.
    p.agent.SetEnableHandler(p.enableAndRespond)
}

func (p *proxy) enableAndRespond(command runtimeAgent.EnableCommand) {
    command.Respond()

    // This ensures we only execute stuff after this if statment once per agent.
    if !atomic.CompareAndSwapInt32(&p.enabled, 0, 1) {
        return
    }

    p.agent.SetGetPropertiesHandler(p.getPropertiesAndRespond)
    p.agent.SetCompileScriptHandler(p.compileScriptAndRespond)

    go shared.WrapFunctionForPanicRecover(p.handleStdout, p.conn)()
    go shared.WrapFunctionForPanicRecover(p.handleStderr, p.conn)()
}

func (p *proxy) compileScriptAndRespond(command runtimeAgent.CompileScriptCommand) {
    command.Respond(nil)
}

func (p *proxy) getPropertiesAndRespond(command runtimeAgent.GetPropertiesCommand) {
    objectId := string(command.ObjectId)
    if strings.HasPrefix(objectId, "local:") {
        var goroutineID int
        var err error
        if command.DestinationTargetID == "" {
            state, err := p.client.GetState()
            if err != nil {
                shared.ThrowError(err.Error())
            }
            if state.SelectedGoroutine == nil {
                shared.ThrowError("No active goroutine")
            }
            goroutineID = int(state.SelectedGoroutine.ID)
        } else {
            goroutineID, err = strconv.Atoi(command.DestinationTargetID)
            if err != nil {
                shared.ThrowError(err.Error())
            }
        }
        frameId, err := strconv.Atoi(objectId[len("local:"):])
        if err != nil {
            shared.ThrowError(err.Error())
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
            shared.ThrowError(err.Error())
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

func (p *proxy) getAndSyncCloseState() bool {
    if p.conn.Closed() {
        p.client.Kill()
        return true
    } else if p.client.Killed() {
        p.conn.Close()
        return true
    }
    return false
}

func (p *proxy) handleStdout() {
    stdout, err := p.client.GetStdout()
    if err != nil {
        shared.ThrowError(err.Error())
    }
    reader := bufio.NewReader(stdout)
    for {
        if p.getAndSyncCloseState() {
            return
        }
        totalData := []byte{}
        for notDone := true; notDone; {
            var data []byte
            var err error
            data, notDone, err = reader.ReadLine()
            totalData = append(totalData, data...)
            if len(totalData) > 5e+6 { // 5 megabytes
                fmt.Println("stdout too big!")
                return
            }
            if err != nil {
                shared.ThrowError(err.Error())
            }
        }

        fmt.Println("Stdout: ", string(totalData))
        p.agent.FireConsoleAPICalled(runtimeAgent.ConsoleAPICalledEvent{
            Type: runtimeAgent.ConsoleAPICalledTypeLog,
            Args: []runtimeAgent.RemoteObject{
                {
                    Type: runtimeAgent.RemoteObjectTypeString,
                    Value: string(totalData),
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
        shared.ThrowError(err.Error())
    }
    reader := bufio.NewReader(stderr)
    for {
        if p.getAndSyncCloseState() {
            return
        }
        totalData := []byte{}
        for notDone := true; notDone; {
            var data []byte
            var err error
            data, notDone, err = reader.ReadLine()
            totalData = append(totalData, data...)
            if len(totalData) > 5e+6 { // 5 megabytes
                fmt.Println("stderr too big!")
                return
            }
            if err != nil {
                shared.ThrowError(err.Error())
            }
        }
        fmt.Println("Stderr: ", string(totalData))
        p.agent.FireConsoleAPICalled(runtimeAgent.ConsoleAPICalledEvent{
            Type: runtimeAgent.ConsoleAPICalledTypeLog,
            Args: []runtimeAgent.RemoteObject{
                {
                    Type: runtimeAgent.RemoteObjectTypeString,
                    Value: string(totalData),
                },
            },
            ExecutionContextId: 0,
        })
    }
}

// TODO This is horrible, must fix!
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
