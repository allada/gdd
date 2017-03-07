package log


type EntryAddedEvent struct {
    Entry LogEntry `json:"entry"`// The entry.
}