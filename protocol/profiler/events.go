package profiler


import (
    "../debugger"
)

type ConsoleProfileStartedEvent struct {
    Id string `json:"id"`
    Location debugger.Location `json:"location"`// Location of console.profile().
    Title *string `json:"title,omitempty"`// Profile title passed as an argument to console.profile().
}
type ConsoleProfileFinishedEvent struct {
    Id string `json:"id"`
    Location debugger.Location `json:"location"`// Location of console.profileEnd().
    Profile Profile `json:"profile"`
    Title *string `json:"title,omitempty"`// Profile title passed as an argument to console.profile().
}