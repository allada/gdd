package types

const (
    StorageTypeAppcache StorageType = "appcache"
    StorageTypeCookies StorageType = "cookies"
    StorageTypeFile_systems StorageType = "file_systems"
    StorageTypeIndexeddb StorageType = "indexeddb"
    StorageTypeLocal_storage StorageType = "local_storage"
    StorageTypeShader_cache StorageType = "shader_cache"
    StorageTypeWebsql StorageType = "websql"
    StorageTypeService_workers StorageType = "service_workers"
    StorageTypeCache_storage StorageType = "cache_storage"
    StorageTypeAll StorageType = "all"
)
type StorageType string
