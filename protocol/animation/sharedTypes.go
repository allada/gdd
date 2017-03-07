package animation


import (
    "../dom"
)
type AnimationTypeEnum string
const (
    AnimationTypeCSSTransition AnimationTypeEnum = "CSSTransition"
    AnimationTypeCSSAnimation AnimationTypeEnum = "CSSAnimation"
    AnimationTypeWebAnimation AnimationTypeEnum = "WebAnimation"
)
type Animation struct {
    Id string `json:"id"`// <code>Animation</code>'s id.
    Name string `json:"name"`// <code>Animation</code>'s name.
    PausedState bool `json:"pausedState"`// [Experimental] <code>Animation</code>'s internal paused state.
    PlayState string `json:"playState"`// <code>Animation</code>'s play state.
    PlaybackRate float64 `json:"playbackRate"`// <code>Animation</code>'s playback rate.
    StartTime float64 `json:"startTime"`// <code>Animation</code>'s start time.
    CurrentTime float64 `json:"currentTime"`// <code>Animation</code>'s current time.
    Source AnimationEffect `json:"source"`// <code>Animation</code>'s source animation node.
    Type AnimationTypeEnum `json:"type"`// Animation type of <code>Animation</code>.
    CssId *string `json:"cssId,omitempty"`// A unique ID for <code>Animation</code> representing the sources that triggered this CSS animation/transition.
}

type AnimationEffect struct {
    Delay float64 `json:"delay"`// <code>AnimationEffect</code>'s delay.
    EndDelay float64 `json:"endDelay"`// <code>AnimationEffect</code>'s end delay.
    IterationStart float64 `json:"iterationStart"`// <code>AnimationEffect</code>'s iteration start.
    Iterations float64 `json:"iterations"`// <code>AnimationEffect</code>'s iterations.
    Duration float64 `json:"duration"`// <code>AnimationEffect</code>'s iteration duration.
    Direction string `json:"direction"`// <code>AnimationEffect</code>'s playback direction.
    Fill string `json:"fill"`// <code>AnimationEffect</code>'s fill mode.
    BackendNodeId dom.BackendNodeId `json:"backendNodeId"`// <code>AnimationEffect</code>'s target node.
    KeyframesRule *KeyframesRule `json:"keyframesRule,omitempty"`// <code>AnimationEffect</code>'s keyframes.
    Easing string `json:"easing"`// <code>AnimationEffect</code>'s timing function.
}

type KeyframesRule struct {
    Name *string `json:"name,omitempty"`// CSS keyframed animation's name.
    Keyframes []KeyframeStyle `json:"keyframes"`// List of animation keyframes.
}

type KeyframeStyle struct {
    Offset string `json:"offset"`// Keyframe's time offset.
    Easing string `json:"easing"`// <code>AnimationEffect</code>'s timing function.
}
