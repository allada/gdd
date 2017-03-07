package applicationcache


import (
    "../page"
)
type ApplicationCacheResource struct {
    Url string `json:"url"`// Resource url.
    Size int64 `json:"size"`// Resource size.
    Type string `json:"type"`// Resource type.
}

type ApplicationCache struct {
    ManifestURL string `json:"manifestURL"`// Manifest URL.
    Size float64 `json:"size"`// Application cache size.
    CreationTime float64 `json:"creationTime"`// Application cache creation time.
    UpdateTime float64 `json:"updateTime"`// Application cache update time.
    Resources []ApplicationCacheResource `json:"resources"`// Application cache resources.
}

type FrameWithManifest struct {
    FrameId page.FrameId `json:"frameId"`// Frame identifier.
    ManifestURL string `json:"manifestURL"`// Manifest URL.
    Status int64 `json:"status"`// Application cache status.
}
