package linked_lists

// Write code to remove duplicates from an unsorted linked list. FOLLOW UP
// How would you solve this problem if a temporary buffer is not allowed?

type node struct {
	next *node
	data int
}

func newLinkedList(data ...int) node {
	if len(data) < 1 {
		panic("need items for a list")
	}
	root := node{
		data: data[0],
	}
	cur := &root
	for i := 1; i < len(data); i++ {
		cur.next = &node{
			data: data[i],
		}
		cur = cur.next
	}
	return root
}
