package inspector


type DetachedEvent struct {
    Reason string `json:"reason"`// The reason why connection has been terminated.
}
type TargetCrashedEvent struct {
}