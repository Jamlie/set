// Package orderedset provides a generic implementation of an ordered set.
//
// A OrderedSet is a collection of unique elements in the same order they have been inserted in,
// implemented using Go's built-in map type.
// The Set is parameterized with a type T, which must be comparable.
package orderedset

import (
	"fmt"
	"iter"
)

// An `OrderedSet` is implemented as a `map[T]struct{}` and `[]T`.
//
// As with maps, a Set requires T to be a comparable, meaning it can
// accept structs if and only if they don't have a type
// like a slice/map/anything that is not comparable
//
// Examples:
//
//	package main
//
//	import (
//		"fmt"
//
//		"github.com/Jamlie/set/orderedset"
//	)
//
//	type Person struct {
//		Id   int
//		Name string
//		Age  int
//	}
//
//	func main() {
//		intsSet := orderedset.New[int]()
//		intsSet.Insert(1)
//		intsSet.Insert(2)
//		intsSet.Insert(3)
//		intsSet.Delete(1)
//
//		fmt.Println(intsSet.Len())
//		fmt.Println(intsSet)
//		if intsSet.Contains(2) {
//			fmt.Println("Set contains number 2")
//		}
//
//		uniquePeople := orderedset.New[Person]()
//		uniquePeople.Insert(Person{Id: 21, Name: "John", Age:30})
//		uniquePeople.Insert(Person{Id: 22, Name: "Jane", Age:30})
//		uniquePeople.Insert(Person{Id: 23, Name: "Roland", Age:30})
//
//		newUnique := uniquePeople.Clone()
//
//		if !newUnique.Empty() {
//			newUnique.Clear()
//		}
//
//		uniquePeople = uniquePeople.
//			Iter().
//			Map(func(k Person) Person {
//				return Person{
//					Id:   k.Id * 3,
//					Name: k.Name,
//					Age:  k.Age,
//				}
//			}).
//			Filter(func(k Person) bool {
//				return k.Id%2 == 1
//			}).
//			Collect()
//		fmt.Println(uniquePeople)
//	}
type OrderedSet[T comparable] struct {
	items []T
	set   map[T]struct{}
}

// Create a new instance of OrderedSet with Go's default capacity.
//
// Examples:
//
//	package main
//
//	import "github.com/Jamlie/set/orderedset"
//
//	func main() {
//		v := orderedset.New[int]()
//		_ = v
//	}
func New[T comparable]() *OrderedSet[T] {
	return &OrderedSet[T]{
		items: []T{},
		set:   make(map[T]struct{}),
	}
}

// Create a new instance of OrderedSet with a specified capacity
//
// The set will be able to hold at least `capacity` without reallocating
// until it's full. This function will panic if capacity is negative.
//
// Examples:
//
//	package main
//
//	import "github.com/Jamlie/orderedset"
//
//	func main() {
//		v := orderedset.WithCapacity[int](10)
//		_ = v
//	}
func WithCapacity[T comparable](capacity int) *OrderedSet[T] {
	if capacity < 0 {
		panic("Cannot allocate with a negative capacity")
	}

	if capacity == 0 {
		return New[T]()
	}

	return &OrderedSet[T]{
		set:   make(map[T]struct{}, capacity),
		items: make([]T, 0, capacity),
	}
}

// Adds a value to the set.
//
// Inserting the same value more than once won't change the set
//
// Examples:
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set/orderedset"
//	)
//
//	func main() {
//		v := orderset.New[int]()
//		v.Insert(1)
//		v.Insert(1)
//		assert.Assert(v.Len() == 1, "Should not insert the same value more than once")
//	}
func (s *OrderedSet[T]) Insert(k T) {
	if _, exists := s.set[k]; !exists {
		s.set[k] = struct{}{}
		s.items = append(s.items, k)
	}
}

// Removes a value from the set.
//
// Removeing a value that does not exists will result in nothing.
//
// Examples:
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set/orderedset"
//	)
//
//	func main() {
//		v := orderedset.New[int]()
//		v.Insert(1)
//		v.Insert(2)
//		v.Delete(1)
//		v.Delete(3)
//		assert.Assert(v.Len() == 1, "Delete should remove at the value if exists")
//	}
func (s *OrderedSet[T]) Delete(k T) {
	if _, exists := s.set[k]; exists {
		delete(s.set, k)
		for i, item := range s.items {
			if item == k {
				s.items = append(s.items[:i], s.items[i+1:]...)
				break
			}
		}
	}
}

// The number of elements the set currently has.
//
// Examples:
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/orderedset"
//	)
//
//	func main() {
//		v := orderedset.New[int]()
//		v.Insert(1)
//		v.Insert(2)
//		v.Insert(3)
//		assert.Assert(v.Len() == 3, "Gets the number of elements")
//	}
func (s *OrderedSet[T]) Len() int {
	return len(s.items)
}

// Returns `true` if the set contains a value.
//
// Examples:
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set/orderedset"
//	)
//
//	func main() {
//		v := orderedset.New[int]()
//		v.Insert(1)
//		v.Insert(2)
//		v.Insert(4)
//		assert.Assert(v.Contains(3) == false, "Number doesn't exist")
//		assert.Assert(v.Contains(4) == true, "Number exist")
//	}
func (s *OrderedSet[T]) Contains(k T) bool {
	_, exists := s.set[k]
	return exists
}

// Returns a clone of the set.
//
// Examples:
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set/orderedset"
//	)
//
//	func main() {
//		v := orderedset.New[int]()
//		v.Insert(1)
//		v.Insert(2)
//		v.Insert(4)
//		clone := v.Clone()
//		assert.Assert(clone.Len() == 3, "Should have the same elements and the same length")
//	}
func (s *OrderedSet[T]) Clone() *OrderedSet[T] {
	clone := New[T]()
	for _, item := range s.items {
		clone.Insert(item)
	}
	return clone
}

// Returns a slice containing the keys of the set in an the order the items where inserted in.
//
// Examples:
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set/orderedset"
//	)
//
//	func main() {
//		v := orderedset.New[int]()
//		v.Insert(1)
//		v.Insert(2)
//		v.Insert(4)
//		keys := v.Keys()
//		assert.Assert(len(keys) == 3, "Should have the same elements and the same length")
//		assert.Equals(keys, []int{1, 2, 4}, "Should have the same elements and the same length")
//	}
func (s *OrderedSet[T]) Keys() []T {
	return s.items
}

// Clears the set, removing all values.
//
// Examples:
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set/orderedset"
//	)
//
//	func main() {
//		v := orderedset.New[string]()
//		v.Insert("first")
//		v.Insert("second")
//		v.Insert("third")
//		v.Clear()
//		assert.Assert(v.Len() == 0, "Should have all elements removed")
//	}
func (s *OrderedSet[T]) Clear() {
	clear(s.items)
	clear(s.set)
}

// Returns `true` if the set contains no elements.
//
// Examples:
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set/orderedset"
//	)
//
//	func main() {
//		v := orderedset.New[int]()
//		assert.Assert(v.Empty(), "Empty set");
//		v.Add(1);
//		assert.Assert(!v.is_empty(), "Set should be empty");
//	}
func (s *OrderedSet[T]) Empty() bool {
	return len(s.items) == 0
}

// Returns a stringified version of the set with elements in the same order
//
// Examples:
//
//	package main
//
//	import (
//		"fmt"
//
//		"github.com/Jamlie/set/orderedset"
//	)
//
//	func main() {
//		v := orderedset.New[int]()
//		v.Insert(1)
//		v.Insert(2)
//		v.Insert(3)
//		fmt.Println(v)
//	}
func (s OrderedSet[T]) String() string {
	return fmt.Sprint(s.items)
}

// A way to iterate through OrderedSet using a range-loop
//
// Examples:
//
//	package main
//
//	import (
//		"log"
//
//		"github.com/Jamlie/set/orderedset"
//	)
//
//	func main() {
//		v := orderedset.New[int]()
//		v.Insert(3)
//		v.Insert(2)
//		v.Insert(1)
//
//		for k := range v.All() {
//			log.Println(k)
//		}
//	}
func (s *OrderedSet[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, k := range s.items {
			if !yield(k) {
				return
			}
		}
	}
}

// Collect allows passing any `iter.Seq[T]` and replaces all values in the existing set.
// Note: Collect changes the whole set.
//
// Examples:
//
//	package main
//
//	import (
//		"log"
//
//		"github.com/Jamlie/set/orderedset"
//	)
//
//	func main() {
//		v := orderedset.New[int]()
//		v.Insert(3)
//		v.Insert(2)
//		v.Insert(1)
//
//		newSet := orderedset.New[int]()
//		newSet.Insert(5)
//		newSet.Collect(v.All())
//		log.Println(newSet) // [3 2 1]
//	}
func (s *OrderedSet[T]) Collect(seq iter.Seq[T]) {
	newSet := WithCapacity[T](s.Len())
	newSet.InsertSeq(seq)
	s.set = newSet.set
	s.items = newSet.items
}

// InsertSeq allows entering any `iter.Seq[T]` and appends all values into the existing set.
//
// Examples:
//
//	package main
//
//	import (
//		"log"
//
//		"github.com/Jamlie/set/orderedset"
//	)
//
//	func main() {
//		v := set/orderedset.New[int]()
//		v.Insert(3)
//		v.Insert(2)
//		v.Insert(1)
//
//		newSet := set/orderedset.New[int]()
//		newSet.Insert(4)
//		newSet.InsertSeq(v.All())
//		log.Println(newSet) // [3 2 1 4]
//	}
func (s *OrderedSet[T]) InsertSeq(seq iter.Seq[T]) {
	for k := range seq {
		s.items = append(s.items, k)
		s.set[k] = struct{}{}
	}
}

// Converts a slice into a set
//
// Examples:
//
//	package main
//
//	import (
//		"fmt"
//
//		"github.com/Jamlie/set/orderedset"
//	)
//
//	func main() {
//		arr := []string{"first", "second", "last"}
//
//		v := orderedset.FromSlice(arr)
//
//		fmt.Println(v)
//	}
func FromSlice[Slice ~[]T, T comparable](v Slice) *OrderedSet[T] {
	s := WithCapacity[T](len(v))

	for _, k := range v {
		s.Insert(k)
	}

	return s
}

// Converts a map into a set
//
// Examples:
//
//	package main
//
//	import (
//		"fmt"
//
//		"github.com/Jamlie/set/orderedset"
//	)
//
//	func main() {
//		m := map[string]int{
//			"first":  1,
//			"second": 2,
//			"last":   3,
//		}
//
//		v := orderedset.FromMap(m)
//
//		fmt.Println(v)
//	}
func FromMap[Map ~map[K]V, K comparable, V any](v Map) *OrderedSet[K] {
	s := WithCapacity[K](len(v))

	for k := range v {
		s.Insert(k)
	}

	return s
}
