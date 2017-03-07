package types

type ConsoleMessageSourceEnum string
const (
    ConsoleMessageSourceXml ConsoleMessageSourceEnum = "xml"
    ConsoleMessageSourceJavascript ConsoleMessageSourceEnum = "javascript"
    ConsoleMessageSourceNetwork ConsoleMessageSourceEnum = "network"
    ConsoleMessageSourceConsoleDashapi ConsoleMessageSourceEnum = "console-api"
    ConsoleMessageSourceStorage ConsoleMessageSourceEnum = "storage"
    ConsoleMessageSourceAppcache ConsoleMessageSourceEnum = "appcache"
    ConsoleMessageSourceRendering ConsoleMessageSourceEnum = "rendering"
    ConsoleMessageSourceSecurity ConsoleMessageSourceEnum = "security"
    ConsoleMessageSourceOther ConsoleMessageSourceEnum = "other"
    ConsoleMessageSourceDeprecation ConsoleMessageSourceEnum = "deprecation"
    ConsoleMessageSourceWorker ConsoleMessageSourceEnum = "worker"
)

type ConsoleMessageLevelEnum string
const (
    ConsoleMessageLevelLog ConsoleMessageLevelEnum = "log"
    ConsoleMessageLevelWarning ConsoleMessageLevelEnum = "warning"
    ConsoleMessageLevelError ConsoleMessageLevelEnum = "error"
    ConsoleMessageLevelDebug ConsoleMessageLevelEnum = "debug"
    ConsoleMessageLevelInfo ConsoleMessageLevelEnum = "info"
)
type ConsoleMessage struct {
    Source ConsoleMessageSourceEnum `json:"source"`// Message source.
    Level ConsoleMessageLevelEnum `json:"level"`// Message severity.
    Text string `json:"text"`// Message text.
    Url *string `json:"url,omitempty"`// URL of the message origin.
    Line *int64 `json:"line,omitempty"`// Line number in the resource that generated this message (1-based).
    Column *int64 `json:"column,omitempty"`// Column number in the resource that generated this message (1-based).
}
