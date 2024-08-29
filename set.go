package set

import (
	"fmt"
	"maps"
)

// A `Set` is implemented as a `map[T]struct{}`.
//
// As with maps, a Set requires T to be a comparable, meaning it can
// accept structs if and only if they don't have a type
// like a slice/map/anything that is not comparable
//
//	# Examples
//
//	```go
//
//	package main
//
//	import (
//		"fmt"
//
//		"github.com/Jamlie/set"
//	)
//
//	type Person struct {
//		Id   int
//		Name string
//		Age  int
//	}
//
//	func main() {
//		intsSet := set.New[int]()
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
//		uniquePeople := set.New[Person]()
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
//
// ```
type Set[T comparable] struct {
	set map[T]struct{}
}

// Create a new instance of Set with Go's default capacity.
//
//	# Examples
//
//	```go
//
//	package main
//
//	import "github.com/Jamlie/set"
//
//	func main() {
//		v := set.New[int]()
//		_ = v
//	}
//
//	```
func New[T comparable]() *Set[T] {
	return &Set[T]{
		set: make(map[T]struct{}),
	}
}

// Create a new instance of Set with a specified capacity
//
// The set will be able to hold at least `capacity` without reallocating
// until it's full. This function will panic if capacity is negative.
//
//	# Examples
//
//	```go
//
//	package main
//
//	import "github.com/Jamlie/set"
//
//	func main() {
//		v := set.WithCapacity[int](10)
//		_ = v
//	}
//
//	```
func WithCapacity[T comparable](capacity int) *Set[T] {
	if capacity < 0 {
		panic("Cannot allocate with a negative capacity")
	}

	if capacity == 0 {
		return New[T]()
	}

	return &Set[T]{
		set: make(map[T]struct{}, capacity),
	}
}

// Adds a value to the set.
//
// Inserting the same value more than once won't change the set
//
//	# Examples
//
//	```go
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set"
//	)
//
//	func main() {
//		v := set.New[int]()
//		v.Insert(1)
//		v.Insert(1)
//		assert.Assert(v.Len() == 1, "Should not insert the same value more than once")
//	}
//
//	```
func (s *Set[T]) Insert(k T) {
	s.set[k] = struct{}{}
}

// Removes a value from the set.
//
// Removeing a value that does not exists will resolt in nothing
//
//	# Examples
//
//	```go
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set"
//	)
//
//	func main() {
//		v := set.New[int]()
//		v.Insert(1)
//		v.Insert(2)
//		v.Delete(1)
//		v.Delete(3)
//		assert.Assert(v.Len() == 1, "Delete should remove at the value if exists")
//	}
//
//	```
func (s *Set[T]) Delete(k T) {
	delete(s.set, k)
}

// The number of elements the set currently has.
//
//	# Examples
//
//	```go
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set"
//	)
//
//	func main() {
//		v := set.New[int]()
//		v.Insert(1)
//		v.Insert(2)
//		v.Insert(3)
//		assert.Assert(v.Len() == 3, "Gets the number of elements")
//	}
//
//	```
func (s *Set[T]) Len() int {
	return len(s.set)
}

// Returns `true` if the set contains a value.
//
//	# Examples
//
//	```go
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set"
//	)
//
//	func main() {
//		v := set.New[int]()
//		v.Insert(1)
//		v.Insert(2)
//		v.Insert(4)
//		assert.Assert(v.Contains(3) == false, "Number doesn't exist")
//		assert.Assert(v.Contains(4) == true, "Number exist")
//	}
//
//	```
func (s *Set[T]) Contains(k T) bool {
	_, ok := s.set[k]
	return ok
}

// Returns a clone of the setA.
//
//	# Examples
//
//	```go
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set"
//	)
//
//	func main() {
//		v := set.New[int]()
//		v.Insert(1)
//		v.Insert(2)
//		v.Insert(4)
//		clone := v.Clone()
//		assert.Assert(clone.Len() == 3, "Should have the same elements and the same length")
//	}
//
//	```
func (s *Set[T]) Clone() *Set[T] {
	return &Set[T]{
		set: maps.Clone(s.set),
	}
}

// Returns a slice containing the keys of the set in an arbitrary ordered.
//
//	# Examples
//
//	```go
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set"
//	)
//
//	func main() {
//		v := set.New[int]()
//		v.Insert(1)
//		v.Insert(2)
//		v.Insert(4)
//		keys := v.Keys()
//		assert.Assert(len(keys) == 3, "Should have the same elements and the same length")
//		assert.Equals(sameSlice(keys, []int{2,1,4}), "Should have the same elements and the same length")
//	}
//
//	// check https://stackoverflow.com/questions/36000487/check-for-equality-on-slices-without-order for source code
//	func sameSlice[T comparable](x, y []T) bool
//
//	```
func (s *Set[T]) Keys() []T {
	keys := make([]T, 0, s.Len())

	for k := range s.set {
		keys = append(keys, k)
	}

	return keys
}

// Clears the set, removing all values.
//
//	# Examples
//
//	```go
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set"
//	)
//
//	func main() {
//		v := set.New[string]()
//		v.Insert("first")
//		v.Insert("second")
//		v.Insert("third")
//		v.Clear()
//		assert.Assert(v.Len() == 0, "Should have all elements removed")
//	}
//
//	```
func (s *Set[T]) Clear() {
	clear(s.set)
}

// Returns `true` if the set contains no elements.
//
//	# Examples
//
//	```go
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set"
//	)
//
//	func main() {
//		v := set.New[int]()
//		assert.Assert(v.Empty(), "Empty set");
//		v.Add(1);
//		assert.Assert(!v.is_empty(), "Set should be empty");
//	}
//
//	```
func (s *Set[T]) Empty() bool {
	return s.Len() == 0
}

// Returns a stringified version of the set with elements in an arbitrary order
//
//	# Examples
//
//	```go
//
//	package main
//
//	import "fmt"
//
//	func main() {
//		v := set.New[int]()
//		v.Insert(1)
//		v.Insert(2)
//		v.Insert(3)
//		fmt.Println(v)
//	}
//
//	```
func (s Set[T]) String() string {
	return fmt.Sprint(s.Keys())
}

// An iterator visiting all elements in arbitrary order.
//
//	# Examples
//
//	```go
//
//	package main
//
//	import (
//		"github.com/Jamlie/assert"
//		"github.com/Jamlie/set"
//	)
//
//	func main() {
//		v := set.New[string]()
//		v.Insert("first")
//		v.Insert("second")
//		v.Insert("third")
//
//		v = v.Iter().Map(...).Filter(...).Collect()
//	}
//
//	```
func (s *Set[T]) Iter() *setIter[T] {
	return &setIter[T]{
		internalSet: s,
	}
}

// Converts a slice into a set
//
//	# Examples
//
//	```go
//
//	package main
//
//	import (
//		"fmt"
//
//		"github.com/Jamlie/set"
//	)
//
//	func main() {
//		arr := []string{"first", "second", "last"}
//
//		v := set.FromSlice(arr)
//
//		fmt.Println(v)
//	}
//
//	```
func FromSlice[T comparable](v []T) *Set[T] {
	s := WithCapacity[T](len(v))

	for _, k := range v {
		s.Insert(k)
	}

	return s
}

// Converts a map into a set
//
//	# Examples
//
//	```go
//
//	package main
//
//	import (
//		"fmt"
//
//		"github.com/Jamlie/set"
//	)
//
//	func main() {
//		arr := map[string]int{
//			"first":  1,
//			"second": 2,
//			"last":   3,
//		}
//
//		v := set.FromMap(arr)
//
//		fmt.Println(v)
//	}
//
//	```
func FromMap[K comparable, V any](v map[K]V) *Set[K] {
	s := WithCapacity[K](len(v))

	for k := range v {
		s.Insert(k)
	}

	return s
}
