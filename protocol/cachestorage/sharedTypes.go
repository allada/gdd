package cachestorage

type CacheId string

type DataEntry struct {
    Request string `json:"request"`// Request url spec.
    Response string `json:"response"`// Response stataus text.
}

type Cache struct {
    CacheId CacheId `json:"cacheId"`// An opaque unique id of the cache.
    SecurityOrigin string `json:"securityOrigin"`// Security origin of the cache.
    CacheName string `json:"cacheName"`// The name of the cache.
}
