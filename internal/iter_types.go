package internal

type (
	MapIterFn[T comparable]     func(k T) T
	FilterIterFn[T comparable]  func(k T) bool
	ForEachIterFn[T comparable] func(k T)
)
