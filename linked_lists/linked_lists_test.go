package linked_lists

import (
	"testing"
)

func TestListFromSlice(t *testing.T) {
	tests := []struct {
		in       []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, false},
	}

	for _, test := range tests {
		actualList := listFromSlice(test.in...)
		actualSlice := sliceFromList(actualList)

		if len(actualSlice) != len(test.in) {
			t.Fail()
		}

		for i := range test.in {
			if actualSlice[i] != test.in[i] {
				t.Fail()
			}
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{1, 3, 2, 3}, []int{1, 3, 2}},
	}

	for _, test := range tests {
		list := listFromSlice(test.in...)
		removeDuplicates(list)

		actualSlice := sliceFromList(list)

		if len(test.expected) != len(actualSlice) {
			t.Fail()
		}

		for i := range test.expected {
			if actualSlice[i] != test.expected[i] {
				t.Fail()
			}
		}
	}
}
