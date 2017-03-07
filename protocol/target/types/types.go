package types

type TargetID string

type BrowserContextID string

type TargetInfo struct {
    TargetId TargetID `json:"targetId"`
    Type string `json:"type"`
    Title string `json:"title"`
    Url string `json:"url"`
}

type RemoteLocation struct {
    Host string `json:"host"`
    Port int64 `json:"port"`
}
