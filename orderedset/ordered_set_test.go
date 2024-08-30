package orderedset_test

import (
	"slices"
	"testing"

	"github.com/Jamlie/set/orderedset"
)

type Joke struct {
	joke     string
	setup    string
	delivery string
}

func TestSetInsert(t *testing.T) {
	test := struct {
		set    *orderedset.OrderedSet[int]
		expect []int
	}{
		set:    orderedset.New[int](),
		expect: []int{1, 2, 3, 4},
	}

	test.set.Insert(1)
	test.set.Insert(2)
	test.set.Insert(3)
	test.set.Insert(4)

	if !slices.Equal(test.set.Keys(), test.expect) {
		t.Fatalf("Expected: %s, Got: %v", test.set, test.expect)
	}
}

func TestSetDelete(t *testing.T) {
	test := struct {
		set    *orderedset.OrderedSet[int]
		expect []int
	}{
		set:    orderedset.New[int](),
		expect: []int{1, 2, 4},
	}

	test.set.Insert(1)
	test.set.Insert(2)
	test.set.Insert(3)
	test.set.Insert(4)

	test.set.Delete(3)

	if !slices.Equal(test.set.Keys(), test.expect) {
		t.Fatalf("Expected: %s, Got: %v", test.set, test.expect)
	}
}

func TestSetContains(t *testing.T) {
	tests := []struct {
		set      *orderedset.OrderedSet[int]
		contains int
		expect   bool
	}{
		{
			set:      orderedset.New[int](),
			contains: 3,
			expect:   true,
		},
		{
			set:      orderedset.New[int](),
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
