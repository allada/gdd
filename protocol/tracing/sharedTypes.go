package tracing

type TraceConfigRecordModeEnum string
const (
    TraceConfigRecordModeRecordUntilFull TraceConfigRecordModeEnum = "recordUntilFull"
    TraceConfigRecordModeRecordContinuously TraceConfigRecordModeEnum = "recordContinuously"
    TraceConfigRecordModeRecordAsMuchAsPossible TraceConfigRecordModeEnum = "recordAsMuchAsPossible"
    TraceConfigRecordModeEchoToConsole TraceConfigRecordModeEnum = "echoToConsole"
)

type StartTransferModeEnum string
const (
    StartTransferModeReportEvents StartTransferModeEnum = "ReportEvents"
    StartTransferModeReturnAsStream StartTransferModeEnum = "ReturnAsStream"
)
type MemoryDumpConfig map[string]string

type TraceConfig struct {
    RecordMode *TraceConfigRecordModeEnum `json:"recordMode,omitempty"`// Controls how the trace buffer stores data.
    EnableSampling *bool `json:"enableSampling,omitempty"`// Turns on JavaScript stack sampling.
    EnableSystrace *bool `json:"enableSystrace,omitempty"`// Turns on system tracing.
    EnableArgumentFilter *bool `json:"enableArgumentFilter,omitempty"`// Turns on argument filter.
    IncludedCategories *[]string `json:"includedCategories,omitempty"`// Included category filters.
    ExcludedCategories *[]string `json:"excludedCategories,omitempty"`// Excluded category filters.
    SyntheticDelays *[]string `json:"syntheticDelays,omitempty"`// Configuration to synthesize the delays in tracing.
    MemoryDumpConfig *MemoryDumpConfig `json:"memoryDumpConfig,omitempty"`// Configuration for memory dump triggers. Used only when "memory-infra" category is enabled.
}
