package set

import "iter"

type setIter[T comparable] struct {
	internalSet *Set[T]
}

type (
	mapIterFn[T comparable]     func(k T) T
	filterIterFn[T comparable]  func(k T) bool
	forEachIterFn[T comparable] func(k T)
)

func (it *setIter[T]) Map(fn mapIterFn[T]) *setIter[T] {
	newSet := New[T]()

	for k := range it.internalSet.set {
		newKey := fn(k)
		newSet.set[newKey] = struct{}{}
	}

	it.internalSet = newSet
	return it.internalSet.Iter()
}

func (it *setIter[T]) Filter(fn filterIterFn[T]) *setIter[T] {
	newSet := New[T]()

	for k := range it.internalSet.set {
		if fn(k) {
			newSet.set[k] = struct{}{}
		}
	}

	it.internalSet = newSet
	return it.internalSet.Iter()
}

func (it *setIter[T]) ForEach(fn forEachIterFn[T]) {
	for k := range it.internalSet.set {
		fn(k)
	}
}

func (it *setIter[T]) Collect() *Set[T] {
	return it.internalSet
}

func Iter[K comparable](m *Set[K]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range m.set {
			if !yield(k) {
				return
			}
		}
	}
}

func Insert[K comparable](s *Set[K], seq iter.Seq[K]) {
	for k := range seq {
		s.set[k] = struct{}{}
	}
}

func Collect[K comparable](seq iter.Seq[K]) *Set[K] {
	// s := make(map[K]V)
	s := New[K]()
	Insert(s, seq)
	return s
}
