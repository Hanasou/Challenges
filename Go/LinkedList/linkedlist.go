package linkedlist

// SingleNode in a singly linked list
type SingleNode struct {
	val  int
	next *SingleNode
}

// DoubleNode is a doubly linked list
type DoubleNode struct {
	val  int
	next *DoubleNode
	prev *DoubleNode
}

func reverse(n *SingleNode) SingleNode {
	prev := &SingleNode{}
	curr := n

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return *prev
}

func nthToLast(n int, head *SingleNode) SingleNode {
	reversed := reverse(head)
	p := &reversed
	for n > 1 {
		p = p.next
		n--
	}
	return *p
}
