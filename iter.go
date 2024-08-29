package set

import "github.com/Jamlie/set/internal"

type setIter[T comparable] struct {
	internalSet *Set[T]
}

func (it *setIter[T]) Map(fn internal.MapIterFn[T]) *setIter[T] {
	newSet := New[T]()

	for k := range it.internalSet.set {
		newKey := fn(k)
		newSet.set[newKey] = struct{}{}
	}

	it.internalSet = newSet
	return it.internalSet.Iter()
}

func (it *setIter[T]) Filter(fn internal.FilterIterFn[T]) *setIter[T] {
	newSet := New[T]()

	for k := range it.internalSet.set {
		if fn(k) {
			newSet.set[k] = struct{}{}
		}
	}

	it.internalSet = newSet
	return it.internalSet.Iter()
}

func (it *setIter[T]) ForEach(fn internal.ForEachIterFn[T]) {
	for k := range it.internalSet.set {
		fn(k)
	}
}

func (it *setIter[T]) Collect() *Set[T] {
	return it.internalSet
}
