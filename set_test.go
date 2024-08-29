package set_test

import (
	"testing"

	"github.com/Jamlie/set"
)

type Joke struct {
	joke     string
	setup    string
	delivery string
}

func TestSetInsert(t *testing.T) {
	test := struct {
		set    *set.Set[int]
		expect []int
	}{
		set:    set.New[int](),
		expect: []int{1, 2, 3, 4},
	}

	test.set.Insert(1)
	test.set.Insert(2)
	test.set.Insert(3)
	test.set.Insert(4)

	if !sameSlice(test.set.Keys(), test.expect) {
		t.Fatalf("Expected: %s, Got: %v", test.set, test.expect)
	}
}

func TestSetDelete(t *testing.T) {
	test := struct {
		set    *set.Set[int]
		expect []int
	}{
		set:    set.New[int](),
		expect: []int{1, 2, 4},
	}

	test.set.Insert(1)
	test.set.Insert(2)
	test.set.Insert(3)
	test.set.Insert(4)

	test.set.Delete(3)

	if !sameSlice(test.set.Keys(), test.expect) {
		t.Fatalf("Expected: %s, Got: %v", test.set, test.expect)
	}
}

func TestSetContains(t *testing.T) {
	tests := []struct {
		set      *set.Set[int]
		contains int
		expect   bool
	}{
		{
			set:      set.New[int](),
			contains: 3,
			expect:   true,
		},
		{
			set:      set.New[int](),
			contains: 5,
			expect:   false,
		},
	}

	for i, test := range tests {
		test.set.Insert(1)
		test.set.Insert(2)
		test.set.Insert(3)
		test.set.Insert(4)

		if test.set.Contains(test.contains) != test.expect {
			t.Fatalf("Index: %d, Expected: %v, Got: %v", i, test.expect, !test.expect)
		}
	}
}

func TestSetMap(t *testing.T) {
	test := struct {
		set    *set.Set[int]
		expect []int
	}{
		set:    set.New[int](),
		expect: []int{2, 4, 6, 8},
	}

	test.set.Insert(1)
	test.set.Insert(2)
	test.set.Insert(3)
	test.set.Insert(4)

	test.set = test.set.Iter().Map(func(k int) int {
		return k * 2
	}).Collect()

	if !sameSlice(test.set.Keys(), test.expect) {
		t.Fatalf("Expected: %s, Got: %v", test.set, test.expect)
	}
}

func TestSetFilter(t *testing.T) {
	test := struct {
		set    *set.Set[int]
		expect []int
	}{
		set:    set.New[int](),
		expect: []int{1, 3},
	}

	test.set.Insert(1)
	test.set.Insert(2)
	test.set.Insert(3)
	test.set.Insert(4)

	test.set = test.set.Iter().Filter(func(k int) bool {
		return k%2 == 1
	}).Collect()

	if !sameSlice(test.set.Keys(), test.expect) {
		t.Fatalf("Expected: %s, Got: %v", test.set, test.expect)
	}
}

// check https://stackoverflow.com/questions/36000487/check-for-equality-on-slices-without-order for source code
func sameSlice[T comparable](x, y []T) bool {
	if len(x) != len(y) {
		return false
	}

	diff := make(map[T]int, len(x))
	for _, _x := range x {
		diff[_x]++
	}

	for _, _y := range y {
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y]--
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	return len(diff) == 0
}
