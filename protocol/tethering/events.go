package tethering


type AcceptedEvent struct {
    Port int64 `json:"port"`// Port number that was successfully bound.
    ConnectionId string `json:"connectionId"`// Connection id to be used.
}