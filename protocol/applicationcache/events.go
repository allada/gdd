package applicationcache


import (
    "github.com/allada/gdd/protocol/page"
)

type ApplicationCacheStatusUpdatedEvent struct {
    FrameId page.FrameId `json:"frameId"`// Identifier of the frame containing document whose application cache updated status.
    ManifestURL string `json:"manifestURL"`// Manifest URL.
    Status int64 `json:"status"`// Updated application cache status.
}
type NetworkStateUpdatedEvent struct {
    IsNowOnline bool `json:"isNowOnline"`
}