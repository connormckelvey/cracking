package linked_lists

import "testing"

func TestNew(t *testing.T) {
	tests := []struct {
		in       []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, false},
	}

	for _, test := range tests {
		actual := newLinkedList(test.in...)
		cur := &actual

		for i := 0; i < len(test.in); i++ {
			if cur.data != test.in[i] {
				t.Fail()
			}
			cur = cur.next
			t.Logf("%#v", cur)
		}
	}
}
