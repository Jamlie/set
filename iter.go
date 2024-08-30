package set

import "github.com/Jamlie/set/internal"

type setIter[T comparable] struct {
	set         *Set[T]
	internalSet *Set[T]
}

func (it *setIter[T]) Map(fn internal.MapIterFn[T]) *setIter[T] {
	newSet := New[T]()

	for k := range it.internalSet.set {
		newKey := fn(k)
		newSet.set[newKey] = struct{}{}
	}

	it.internalSet.set = newSet.set
	return it.internalSet.Iter()
}

func (it *setIter[T]) Filter(fn internal.FilterIterFn[T]) *setIter[T] {
	newSet := New[T]()

	for k := range it.internalSet.set {
		if fn(k) {
			newSet.set[k] = struct{}{}
		}
	}

	it.internalSet.set = newSet.set
	return it.internalSet.Iter()
}

func (it *setIter[T]) ForEach(fn internal.ForEachIterFn[T]) {
	for k := range it.internalSet.set {
		fn(k)
	}
}

func (it *setIter[T]) Collect() {
	it.set.set = it.internalSet.set
}
