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
func removeDuplicatesNoBuffer(n *node) {
	if n == nil {
		return
	}

	prev, cur := n, n.next
	for cur != nil {
		runner := n
		for runner != cur {
			if runner.data == cur.data {
				// remove cur
				prev.next, cur = cur.next, cur.next
				break
			}
			runner = runner.next
		}
		if runner == cur {
			prev, cur = cur, cur.next
		}
	}
}

// 2.2
// Implement an algorithm to find the nth to last element of a singly linked list.

func findNthFromLast(head *node, n int) *node {
	if head == nil {
		return nil
	}

	var length int
	for cur := head; cur != nil; {
		cur = cur.next
		length++
	}

	if n > length {
		return nil
	}

	for cur, i := head, 0; i <= length-n; {
		if i == length-n {
			return cur
		}
		cur = cur.next
		i++
	}

	return nil
}
