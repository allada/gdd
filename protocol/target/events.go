package target


type TargetCreatedEvent struct {
    TargetInfo TargetInfo `json:"targetInfo"`
}
type TargetDestroyedEvent struct {
    TargetId TargetID `json:"targetId"`
}
type AttachedToTargetEvent struct {
    TargetInfo TargetInfo `json:"targetInfo"`
    WaitingForDebugger bool `json:"waitingForDebugger"`
}
type DetachedFromTargetEvent struct {
    TargetId TargetID `json:"targetId"`
}
type ReceivedMessageFromTargetEvent struct {
    TargetId TargetID `json:"targetId"`
    Message string `json:"message"`
}