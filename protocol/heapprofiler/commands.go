package heapprofiler


import (
    "../shared"
    "../runtime"
)

type EnableCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type EnableReturn struct {
}

func (c *EnableCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *EnableCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type DisableCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type DisableReturn struct {
}

func (c *DisableCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DisableCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type StartTrackingHeapObjectsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    TrackAllocations *bool `json:"trackAllocations,omitempty"`
}
type StartTrackingHeapObjectsReturn struct {
}

func (c *StartTrackingHeapObjectsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StartTrackingHeapObjectsCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type StopTrackingHeapObjectsCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ReportProgress *bool `json:"reportProgress,omitempty"`// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken when the tracking is stopped.
}
type StopTrackingHeapObjectsReturn struct {
}

func (c *StopTrackingHeapObjectsCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StopTrackingHeapObjectsCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type TakeHeapSnapshotCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ReportProgress *bool `json:"reportProgress,omitempty"`// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
}
type TakeHeapSnapshotReturn struct {
}

func (c *TakeHeapSnapshotCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *TakeHeapSnapshotCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type CollectGarbageCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type CollectGarbageReturn struct {
}

func (c *CollectGarbageCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *CollectGarbageCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetObjectByHeapObjectIdCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ObjectId HeapSnapshotObjectId `json:"objectId"`
    ObjectGroup *string `json:"objectGroup,omitempty"`// Symbolic group name that can be used to release multiple objects.
}
type GetObjectByHeapObjectIdReturn struct {
    Result runtime.RemoteObject `json:"result"`// Evaluation result.
}

func (c *GetObjectByHeapObjectIdCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetObjectByHeapObjectIdCommand) Respond(r *GetObjectByHeapObjectIdReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type AddInspectedHeapObjectCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    HeapObjectId HeapSnapshotObjectId `json:"heapObjectId"`// Heap snapshot object id to be accessible by means of $x command line API.
}
type AddInspectedHeapObjectReturn struct {
}

func (c *AddInspectedHeapObjectCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *AddInspectedHeapObjectCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type GetHeapObjectIdCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ObjectId runtime.RemoteObjectId `json:"objectId"`// Identifier of the object to get heap object id for.
}
type GetHeapObjectIdReturn struct {
    HeapSnapshotObjectId HeapSnapshotObjectId `json:"heapSnapshotObjectId"`// Id of the heap snapshot object corresponding to the passed remote object id.
}

func (c *GetHeapObjectIdCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *GetHeapObjectIdCommand) Respond(r *GetHeapObjectIdReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}

type StartSamplingCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    SamplingInterval *float64 `json:"samplingInterval,omitempty"`// Average sample interval in bytes. Poisson distribution is used for the intervals. The default value is 32768 bytes.
}
type StartSamplingReturn struct {
}

func (c *StartSamplingCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StartSamplingCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type StopSamplingCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type StopSamplingReturn struct {
    Profile SamplingHeapProfile `json:"profile"`// Recorded sampling heap profile.
}

func (c *StopSamplingCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StopSamplingCommand) Respond(r *StopSamplingReturn) {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, r)
}
