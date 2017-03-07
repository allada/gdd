package log


import (
    "../runtime"
    "../network"
)
type LogEntrySourceEnum string
const (
    LogEntrySourceXml LogEntrySourceEnum = "xml"
    LogEntrySourceJavascript LogEntrySourceEnum = "javascript"
    LogEntrySourceNetwork LogEntrySourceEnum = "network"
    LogEntrySourceStorage LogEntrySourceEnum = "storage"
    LogEntrySourceAppcache LogEntrySourceEnum = "appcache"
    LogEntrySourceRendering LogEntrySourceEnum = "rendering"
    LogEntrySourceSecurity LogEntrySourceEnum = "security"
    LogEntrySourceDeprecation LogEntrySourceEnum = "deprecation"
    LogEntrySourceWorker LogEntrySourceEnum = "worker"
    LogEntrySourceViolation LogEntrySourceEnum = "violation"
    LogEntrySourceIntervention LogEntrySourceEnum = "intervention"
    LogEntrySourceOther LogEntrySourceEnum = "other"
)

type LogEntryLevelEnum string
const (
    LogEntryLevelVerbose LogEntryLevelEnum = "verbose"
    LogEntryLevelInfo LogEntryLevelEnum = "info"
    LogEntryLevelWarning LogEntryLevelEnum = "warning"
    LogEntryLevelError LogEntryLevelEnum = "error"
)

type ViolationSettingNameEnum string
const (
    ViolationSettingNameLongTask ViolationSettingNameEnum = "longTask"
    ViolationSettingNameLongLayout ViolationSettingNameEnum = "longLayout"
    ViolationSettingNameBlockedEvent ViolationSettingNameEnum = "blockedEvent"
    ViolationSettingNameBlockedParser ViolationSettingNameEnum = "blockedParser"
    ViolationSettingNameDiscouragedAPIUse ViolationSettingNameEnum = "discouragedAPIUse"
    ViolationSettingNameHandler ViolationSettingNameEnum = "handler"
    ViolationSettingNameRecurringHandler ViolationSettingNameEnum = "recurringHandler"
)
type LogEntry struct {
    Source LogEntrySourceEnum `json:"source"`// Log entry source.
    Level LogEntryLevelEnum `json:"level"`// Log entry severity.
    Text string `json:"text"`// Logged text.
    Timestamp runtime.Timestamp `json:"timestamp"`// Timestamp when this entry was added.
    Url *string `json:"url,omitempty"`// URL of the resource if known.
    LineNumber *int64 `json:"lineNumber,omitempty"`// Line number in the resource.
    StackTrace *runtime.StackTrace `json:"stackTrace,omitempty"`// JavaScript stack trace.
    NetworkRequestId *network.RequestId `json:"networkRequestId,omitempty"`// Identifier of the network request associated with this entry.
    WorkerId *string `json:"workerId,omitempty"`// Identifier of the worker associated with this entry.
}

type ViolationSetting struct {
    Name ViolationSettingNameEnum `json:"name"`// Violation type.
    Threshold float64 `json:"threshold"`// Time threshold to trigger upon.
}
