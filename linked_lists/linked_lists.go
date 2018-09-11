package linked_lists

type node struct {
	next *node
	data int
}

func listFromSlice(data ...int) *node {
	if len(data) < 1 {
		panic("need items for a list")
	}
	root := &node{
		data: data[0],
	}
	cur := root
	for i := 1; i < len(data); i++ {
		cur.next = &node{
			data: data[i],
		}
		cur = cur.next
	}
	return root
}

func sliceFromList(n *node) []int {
	var slice []int
	cur := n
	for cur != nil {
		slice = append(slice, cur.data)
		cur = cur.next
	}
	return slice
}

// 2.1
// Write code to remove duplicates from an unsorted linked list.
func removeDuplicates(n *node) {
	seen := make(map[int]bool)
	var prev *node

	for n != nil {
		if _, ok := seen[n.data]; ok {
			// remove node
			prev.next = n.next
		} else {
			seen[n.data] = true
			prev = n
		}
		n = n.next
	}
}

// 2.1.b
// How would you solve this problem if a temporary buffer is not allowed?
