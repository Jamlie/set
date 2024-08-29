# Set

An implementation of a generic Set data structure in Go with an `Iter` function to simulate Rust's iterators and has Go's version of iterator functions, leveraging the best of both worlds.

## Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/Jamlie/set"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	intsSet := set.New[int]()
	intsSet.Insert(1)
	intsSet.Insert(2)
	intsSet.Insert(3)
	intsSet.Delete(1)

	fmt.Println(intsSet.Len())
	fmt.Println(intsSet)
	if intsSet.Contains(2) {
		fmt.Println("Set contains number 2")
	}

	uniquePeople := set.New[Person]()
	uniquePeople.Insert(Person{Id: 21, Name: "John", Age: 30})
	uniquePeople.Insert(Person{Id: 22, Name: "Jane", Age: 31})
	uniquePeople.Insert(Person{Id: 23, Name: "Roland", Age: 32})

	newUnique := uniquePeople.Clone()

	if !newUnique.Empty() {
		newUnique.Clear()
	}

	uniquePeople = uniquePeople.
		Iter().
		Map(func(k Person) Person {
			return Person{
				Id:   k.Id * 3,
				Name: k.Name,
				Age:  k.Age,
			}
		}).
		Filter(func(k Person) bool {
			return k.Id%2 == 1
		}).
		Collect()

	for k := range uniquePeople.All() {
		log.Println(k)
	}

	fmt.Println(uniquePeople)

	newPeople := set.New[Person]()

	newPeople.Collect(uniquePeople.All())

	log.Println(newPeople)
}
```
