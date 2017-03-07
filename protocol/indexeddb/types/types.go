package types


import (
    "../runtime/types"
)
type KeyTypeEnum string
const (
    KeyTypeNumber KeyTypeEnum = "number"
    KeyTypeString KeyTypeEnum = "string"
    KeyTypeDate KeyTypeEnum = "date"
    KeyTypeArray KeyTypeEnum = "array"
)

type KeyPathTypeEnum string
const (
    KeyPathTypeNull KeyPathTypeEnum = "null"
    KeyPathTypeString KeyPathTypeEnum = "string"
    KeyPathTypeArray KeyPathTypeEnum = "array"
)
type DatabaseWithObjectStores struct {
    Name string `json:"name"`// Database name.
    Version int64 `json:"version"`// Database version.
    ObjectStores []ObjectStore `json:"objectStores"`// Object stores in this database.
}

type ObjectStore struct {
    Name string `json:"name"`// Object store name.
    KeyPath KeyPath `json:"keyPath"`// Object store key path.
    AutoIncrement bool `json:"autoIncrement"`// If true, object store has auto increment flag set.
    Indexes []ObjectStoreIndex `json:"indexes"`// Indexes in this object store.
}

type ObjectStoreIndex struct {
    Name string `json:"name"`// Index name.
    KeyPath KeyPath `json:"keyPath"`// Index key path.
    Unique bool `json:"unique"`// If true, index is unique.
    MultiEntry bool `json:"multiEntry"`// If true, index allows multiple entries for a key.
}

type Key struct {
    Type KeyTypeEnum `json:"type"`// Key type.
    Number *float64 `json:"number,omitempty"`// Number value.
    String *string `json:"string,omitempty"`// String value.
    Date *float64 `json:"date,omitempty"`// Date value.
    Array *[]Key `json:"array,omitempty"`// Array value.
}

type KeyRange struct {
    Lower *Key `json:"lower,omitempty"`// Lower bound.
    Upper *Key `json:"upper,omitempty"`// Upper bound.
    LowerOpen bool `json:"lowerOpen"`// If true lower bound is open.
    UpperOpen bool `json:"upperOpen"`// If true upper bound is open.
}

type DataEntry struct {
    Key runtime.RemoteObject `json:"key"`// Key object.
    PrimaryKey runtime.RemoteObject `json:"primaryKey"`// Primary key object.
    Value runtime.RemoteObject `json:"value"`// Value object.
}

type KeyPath struct {
    Type KeyPathTypeEnum `json:"type"`// Key path type.
    String *string `json:"string,omitempty"`// String value.
    Array *[]string `json:"array,omitempty"`// Array value.
}
