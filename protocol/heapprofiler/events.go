package heapprofiler


type AddHeapSnapshotChunkEvent struct {
    Chunk string `json:"chunk"`
}
type ResetProfilesEvent struct {
}
type ReportHeapSnapshotProgressEvent struct {
    Done int64 `json:"done"`
    Total int64 `json:"total"`
    Finished *bool `json:"finished,omitempty"`
}
type LastSeenObjectIdEvent struct {
    LastSeenObjectId int64 `json:"lastSeenObjectId"`
    Timestamp float64 `json:"timestamp"`
}
type HeapStatsUpdateEvent struct {
    StatsUpdate []int64 `json:"statsUpdate"`// An array of triplets. Each triplet describes a fragment. The first integer is the fragment index, the second integer is a total count of objects for the fragment, the third integer is a total size of the objects for the fragment.
}