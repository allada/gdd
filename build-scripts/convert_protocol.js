var fs = require('fs');

String.prototype.capitalize = function() {
    return this.charAt(0).toUpperCase() + this.slice(1);
}

var browser_protocol = JSON.parse(fs.readFileSync('./build-scripts/protocol/browser_protocol.json').toString());
// JS protocol overrides browser_protocol stuff.
var js_protocol = JSON.parse(fs.readFileSync('./build-scripts/protocol/js_protocol.json').toString());

var domains = new Map();
browser_protocol.domains.forEach(data => domains.set(data.domain, data));
js_protocol.domains.forEach(data => domains.set(data.domain, data));

var out = [];
var enums = new Map();

for (var domain of domains.values()) {
    var events = [];
    var types = [];
    var commands = [];

    var domainName = domain.domain;
    // if (domainName !== 'Debugger')
    //     continue;
    out.push(`package protocol${domainName}`);

    var eventsImports = new Set();
    if ('events' in domain) {
        for (var event of domain.events) {
            //console.log(event.name);
            events.push(`type ${event.name.capitalize()}Event struct {`);
            if ('parameters' in event) {
                for (var param of event.parameters) {
                    events.push(`    ${buildProperty(domainName, event.name, param, eventsImports)}`);
                }
            }
            events.push(`}`);
        }
    }

    var typesImports = new Set();
    if ('types' in domain) {
        for (var type of domain.types) {
            //console.log(event.name);
            var typeData = [];
            if ('properties' in type) {
                typeData.push ('struct {');
                for (var property of type.properties) {
                    typeData.push(`    ${buildProperty(domainName, type.id, property, typesImports)}`);
                }
                typeData.push ('}');
            } else {
                typeData.push(findType(domainName, type.id, type, undefined, true, typesImports));
            }
            types.push(`type ${type.id.capitalize()} ${typeData.join('\n')}`);
            types.push('');
        }
    }

    var commandsImports = new Set();
    var commandsNeedsShared = false;
    if ('commands' in domain) {
        for (var command of domain.commands) {
            //console.log(command.name);
            var commandName = command.name;

            var parameterData = [];
            parameterData.push ('struct {');
            parameterData.push ('    DestinationTargetID string');
            parameterData.push ('    responseId int64');
            parameterData.push ('    conn *shared.Connection');
            if ('parameters' in command) {
                for (var parameter of command.parameters) {
                    parameterData.push(`    ${buildProperty(domainName, commandName.capitalize(), parameter, commandsImports)}`);
                }
            }
            parameterData.push ('}');
            commands.push(`type ${command.name.capitalize()}Command ${parameterData.join('\n')}`);

            var respondData = '';

            var returnData = [];
            returnData.push ('struct {');
            commandsNeedsShared = true;
            if ('returns' in command) {
                for (var ret of command.returns) {
                    returnData.push(`    ${buildProperty(domainName, commandName.capitalize(), ret, commandsImports)}`);
                }
                respondData = `r *${command.name.capitalize()}Return`;
            }
            returnData.push ('}');
            commands.push(`type ${command.name.capitalize()}Return ${returnData.join('\n')}`);
            commands.push('');
            // go send(ws, protocol.ErrorResult{Id: id, Status: protocol.StatusCodeError, Error: protocol.ErrorResponse{
            //     Code: protocol.ErrorCodeMethodNotFound,
            // }})
            commands.push(`func (c *${command.name.capitalize()}Command) RespondWithError(code shared.ResponseErrorCodes, message string) {`);
            commands.push(`    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)`);
            commands.push(`}`);
            commands.push('');
            commands.push(`func (c *${command.name.capitalize()}Command) Respond(${respondData}) {`);
            if (respondData) {
                commands.push(`    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)`);
            } else {
                commands.push(`    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})`);
            }
            commands.push('}');
            commands.push('');
        }
    }
    // Must be last.
    var enumData = enums.get(domainName);
    // if (enumData) {
    //     console.log(enumData.join('\n'));
    // }

    try {
        fs.mkdirSync('./protocol/');
    } catch(e) {}
    try {
        fs.mkdirSync('./protocol/' + domainName.toLowerCase() + '/');
    } catch(e) {}
    // var typeData = [];
    // if (enumData) {
    //     typeData.push(enumData.join('\n'));
    // }
    // typeData.push(events.join('\n'));
    // typeData.push(types.join('\n'));
    // typeData.push(commands.join('\n'));
    var packageData = [
        'package ' + domainName.toLowerCase(),
        '',
        ''
    ];

    var sharedTypesHeaders = [packageData.join('\n')];
    if (typesImports.size) {
        sharedTypesHeaders.push('import (');
        for (var v of typesImports.values()) {
            sharedTypesHeaders.push(`    "../${v}"`);
        }
        sharedTypesHeaders.push(')');
        sharedTypesHeaders.push('');
    }
    fs.writeFileSync('./protocol/' + domainName.toLowerCase() + '/sharedTypes.go', [sharedTypesHeaders.join('\n'), enumData ? enumData.join('\n') : '', types.join('\n')].join(''));

    var eventHeaders = [packageData.join('\n')];
    if (eventsImports.size) {
        eventHeaders.push('import (');
        for (var v of eventsImports.values()) {
            eventHeaders.push(`    "../${v}"`);
        }
        eventHeaders.push(')');
        eventHeaders.push('');
    }
    fs.writeFileSync('./protocol/' + domainName.toLowerCase() + '/events.go', [eventHeaders.join('\n'), events.join('\n')].join('\n'));

    var commandHeaders = [packageData.join('\n')];
    if (commandsImports.size || commandsNeedsShared) {
        commandHeaders.push('import (');
        if (commandsNeedsShared) {
            commandHeaders.push(`    "../shared"`);
        }
        for (var v of commandsImports.values()) {
            commandHeaders.push(`    "../${v}"`);
        }
        commandHeaders.push(')');
        commandHeaders.push('');
    }
    fs.writeFileSync('./protocol/' + domainName.toLowerCase() + '/commands.go', [commandHeaders.join('\n'), commands.join('\n')].join('\n'));

    var commandSlices = [];
    var switchData = [];
    switchData.push(`    switch (funcName) {`)
    if ('commands' in domain) {
        commandSlices.push(`    commandChans struct {`);
        for (var command of domain.commands) {
            var commandName = command.name.capitalize();
            switchData.push(`        case "${command.name}":`)
            switchData.push(`            var out ${commandName}Command`);
            switchData.push(`            if data != nil {`);
            switchData.push(`                if err := json.Unmarshal(*data, &out); err != nil {`);
            switchData.push(`                    fmt.Println("Closing connection because malformed json was sent.")`);
            switchData.push(`                    agent.conn.Close()`);
            switchData.push(`                    return`);
            switchData.push(`                }`);
            switchData.push(`            }`);
            switchData.push(`            out.DestinationTargetID = targetId`);
            switchData.push(`            out.responseId = id`);
            switchData.push(`            out.conn = agent.conn`);
            switchData.push(`            if len(agent.commandChans.${commandName}) == 0 {`);
            switchData.push(`                out.RespondWithError(shared.ErrorCodeMethodNotFound, "")`);
            switchData.push(`                break`);
            switchData.push(`            }`);
            switchData.push(`            for _, fn := range agent.commandChans.${commandName} {`);
            switchData.push(`                fn(out)`);
            switchData.push(`            }`);
            commandSlices.push(`        ${commandName} []func(${commandName}Command)`);
        }
        commandSlices.push(`    }`);
    }
    switchData.push(`        default:`);
    switchData.push(`            fmt.Printf("Command %s unknown\\n", funcName)`);
    switchData.push(`            agent.conn.SendErrorResult(id, targetId, shared.ErrorCodeMethodNotFound, "")`);
    switchData.push(`    }`);

    var helperData = [

`type ${domainName}Agent struct {
    conn *shared.Connection
${commandSlices.join('\n')}
}`,
`func NewAgent(conn *shared.Connection) *${domainName}Agent {
    agent := &${domainName}Agent{
        conn: conn,
    }
    conn.RegisterAgent(agent)
    return agent
}`,
'',
`func (agent *${domainName}Agent) Name() string {
    return "${domainName}"
}`,
'',
`func (agent *${domainName}Agent) ProcessCommand(id int64, targetId string, funcName string, data *json.RawMessage) {
    defer func() {
        data := recover()
        switch data.(type) {
        case nil:
            return
        case shared.Warning:
            fmt.Println(data)
        case shared.Error:
            fmt.Println("Closing websocket because of following Error:")
            fmt.Println(data)
            agent.conn.Close()
        case error:
            fmt.Println("Closing websocket because of following Error:")
            fmt.Println(data)
            agent.conn.Close()
        default:
            fmt.Println("Should probably use shared.Warning or shared.Error instead to panic()")
            panic(data)
        }
    }()
${switchData.join('\n')}
}`,
'',
'// Dispatchable Events'];
    if ('events' in domain) {
        for (var event of domain.events) {
            var eventName = event.name.capitalize();
            var paramData = '';
            var outParamsData = '';
            if (event.parameters) {
                paramData = `event ${eventName}Event`;
                outParamsData = `\n        Params: event,`
            }
            helperData.push(
`func (agent *${domainName}Agent) Fire${eventName}(${paramData}) {
    agent.conn.Send(shared.Notification{
        Method: "${domainName}.${event.name}",${outParamsData}
    })
}`);
            helperData.push(
`func (agent *${domainName}Agent) Fire${eventName}OnTarget(targetId string${paramData ? ',' + paramData : ''}) {
    agent.conn.SendToTarget(targetId, shared.Notification{
        Method: "${domainName}.${event.name}",${outParamsData}
    })
}`);
        }
    }
    helperData.push('', '// Commands Sent From Frontend');
    if ('commands' in domain) {
        for (var command of domain.commands) {
            var commandName = command.name.capitalize();
            helperData.push(
`func (agent *${domainName}Agent) ${commandName}Handler(handler func(${commandName}Command)) {
    agent.commandChans.${commandName} = append(agent.commandChans.${commandName}, handler)
}`);
        }
    }
    helperData.push(
`func init() {

}`,
'')
    // var headerData = ['import ('];
    // headerData.push(`    "../shared"`)
    // if ('dependencies' in domain) {
    //     for (var dependency of domain.dependencies) {
    //         headerData.push(`    "../${dependency.toLowerCase()}"`)
    //     }
    // }
    // headerData.push(')');
    // headerData.push('');
    var sharedTypesHeaders = [packageData.join('\n'), 'import (', '    "../shared"', '    "encoding/json"', '    "fmt"', ')'];
    // if ('events' in domain && domain.events.length) {
    //     sharedTypesHeaders.push(headerData.join('\n'));
    // }
    fs.writeFileSync('./protocol/' + domainName.toLowerCase() + '/agent.go', [sharedTypesHeaders.join('\n'), helperData.join('\n')].join('\n'));
}
try {
    fs.mkdirSync('./protocol/shared/');
} catch(e) {}
fs.writeFileSync('./protocol/shared/types.go',
`package shared

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
    Code ResponseErrorCodes \`json:"code"\`
    Message string \`json:"message"\`
}

type Notification struct {
    Method string \`json:"method"\`
    Params interface{} \`json:"params,omitempty"\`
}

type Response struct {
    Id int64 \`json:"id"\`
    Result interface{} \`json:"result"\`
    Status ResponseStatusCodes \`json:"status,omitempty"\`
    Error *ResponseError \`json:"error,omitempty"\`
}

type ReceivedCommand struct {
    Id int64 \`json:"id"\`
    Method string \`json:"method"\`
    Params *json.RawMessage \`json:"params,omitempty"\`
}

type Error struct {
    msg string
}

func (e Error) Error() string {
    return e.msg
}

type Warning struct {
    msg string
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

type Connection struct {
    ws *websocket.Conn
    agents map[string]Agenter
    isClosed bool
}

func (c *Connection) Close() {
    if c.isClosed {
        return
    }
    c.isClosed = true
    c.ws.Close()
}

func (c *Connection) RegisterAgent(agent Agenter) {
    // TODO This is not thread safe.
    c.agents[agent.Name()] = agent
}

func (c *Connection) Closed() bool {
    return c.isClosed
}

func (c *Connection) SendToTarget(targetId string, event interface{}) {
    if c.isClosed {
        return
    }
    data := struct {
        TargetId string \`json:"targetId"\`
        Message interface{} \`json:"message"\`
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
    if c.isClosed {
        return
    }
    websocket.JSON.Send(c.ws, event)
}

func (c *Connection) SendErrorResult(id int64, targetId string, errorCode ResponseErrorCodes, message string) {
    if c.isClosed {
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
    if c.isClosed {
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

func (c *Connection) socketListener() {
    defer c.Close()
    for {
        var message []byte
        err := websocket.Message.Receive(c.ws, &message)
        if c.isClosed {
            return
        } else if err == io.EOF {
            fmt.Println("Connection closed by remote")
            return
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
                TargetId string \`json:"targetId"\`
                Message string \`json:"message"\`
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
            // fmt.Printf("Agent '%s' not registered.\\n", agentName)
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

`);
//console.log(out);

function buildProperty(domainName, namespace, data, importsMap) {
    var jsonOptions = [data.name];
    var type = findType(domainName, namespace, data, undefined, undefined, importsMap);
    var isPointer = false;
    if (('optional' in data && data.optional) && type !== "interface{}") {
        jsonOptions.push('omitempty');
        isPointer = true;
    }
    var comments = [];
    if ('experimental' in data && data.experimental) {
        comments.push('[Experimental]');
    }
    if (data.description) {
        comments.push(data.description);
    }
    var comment = '';
    if (comments.length) {
        comment = '// ' + comments.join(' ');
    }
    return `${data.name.capitalize()} ${isPointer ? '*' : ''}${type} \`json:"${jsonOptions.join(',')}"\`${comment}`;
}

function findType(domainName, namespace, obj, previousNamespace, isNamedType, importsMap) {
    if ('type' in obj) {
        if (obj.type === 'array') {
            return '[]' + findType(domainName, namespace, obj.items, obj.id || obj.name, false, importsMap);
        }
        var type = protocolTypeToGoPrimitive(domainName, namespace, obj.type, obj, previousNamespace || obj.id || obj.name, isNamedType);
        if (!type) {
            console.error("Unknown type: " + obj.type);
        }
        if (isNamedType && obj.enum) {
            return 'string';
        }
        return type;
    } else if ('$ref' in obj) {
        var refData = obj.$ref.split('.');
        if (refData.length === 1) {
            return refData.join('');
        }
        console.assert(refData.length === 2, "$ref Expected max 1 period.");
        if (refData[0].toLowerCase() !== domainName.toLowerCase()) {
            importsMap.add(refData[0].toLowerCase());
        }
        return refData[0].toLowerCase() + '.' + refData[1];
    }
    console.error("Expected 'type' or 'ref' in " + JSON.stringify(obj));
}

function protocolTypeToGoPrimitive(domainName, namespace, type, obj, previousNamespace, isNamedType) {
    switch (type) {
        case 'boolean':
            return 'bool';
        case 'string':
            if ('enum' in obj) {
                var name = namespace;
                if (!isNamedType) {
                    name = namespace + previousNamespace.capitalize();
                }
                return registerEnum(domainName, name.capitalize(), obj.enum, isNamedType);
            }
            return 'string';
        case 'integer':
            return 'int64';
        case 'number':
            return 'float64';
        case 'object':
            return 'map[string]string';
        case 'any':
            return 'interface{}';
    }
    return null;
}

function registerEnum(domainName, namespace, enumData, isNamedType) {
    var enumResult = enums.get(domainName);
    if (!enumResult)
        enumResult = [];
    var typeName = namespace;
    if (!isNamedType) {
        typeName = `${namespace}Enum`;
        enumResult.push(`type ${typeName} string`);
    }
    enumResult.push('const (');
    for (var value of enumData) {
        var enumName = cleanEnumName(namespace + value.capitalize());
        enumResult.push(`    ${enumName} ${typeName} = "${value}"`);
        //enumResult.push(namespace, namespace.capitalize() + value.capitalize());
    }
    enumResult.push(')');
    enumResult.push('');
    enums.set(domainName, enumResult);
    return typeName;
}

function cleanEnumName(data) {
    return data.replace('-', 'Dash').replace(/[^a-zA-Z0-9_]+/g, '');
}