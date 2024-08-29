package concurrentset

import "github.com/Jamlie/set/internal"

type concurrentSetIter[T comparable] struct {
	internalSet *ConcurrentSet[T]
}

func (it *concurrentSetIter[T]) Map(fn internal.MapIterFn[T]) *concurrentSetIter[T] {
	newSet := New[T]()

	it.internalSet.mu.RLock()
	defer it.internalSet.mu.RUnlock()

	for k := range it.internalSet.set {
		newKey := fn(k)
		newSet.set[newKey] = struct{}{}
	}

	it.internalSet = newSet
	return it.internalSet.Iter()
}

func (it *concurrentSetIter[T]) Filter(fn internal.FilterIterFn[T]) *concurrentSetIter[T] {
	newSet := New[T]()

	it.internalSet.mu.RLock()
	defer it.internalSet.mu.RUnlock()

	for k := range it.internalSet.set {
		if fn(k) {
			newSet.set[k] = struct{}{}
		}
	}

	it.internalSet = newSet
	return it.internalSet.Iter()
}

func (it *concurrentSetIter[T]) ForEach(fn internal.ForEachIterFn[T]) {
	it.internalSet.mu.RLock()
	defer it.internalSet.mu.RUnlock()

	for k := range it.internalSet.set {
		fn(k)
	}
}

func (it *concurrentSetIter[T]) Collect() *ConcurrentSet[T] {
	return it.internalSet
}
