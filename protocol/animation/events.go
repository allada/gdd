package animation


type AnimationCreatedEvent struct {
    Id string `json:"id"`// Id of the animation that was created.
}
type AnimationStartedEvent struct {
    Animation Animation `json:"animation"`// Animation that was started.
}
type AnimationCanceledEvent struct {
    Id string `json:"id"`// Id of the animation that was cancelled.
}