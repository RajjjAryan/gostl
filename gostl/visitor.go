package gostl

// Visitor is a function type for elements
type Visitor[T any] func(value T) bool

// KvVisitor is a function type for key-value type elements
type KvVisitor[K, V any] func(key K, value V) bool