package linkedlist

import (
	"strconv"
)

// SingleNode in a singly linked list
type SingleNode struct {
	Val  int
	Next *SingleNode
}

// DoubleNode is a doubly linked list
type DoubleNode struct {
	val  int
	next *DoubleNode
	prev *DoubleNode
}

func (this *SingleNode) InsertFirst(val int) {
	temp := *this
	this.Val = val
	this.Next = &temp
}

func (this *SingleNode) Append(val int) {
	new := &SingleNode{
		Val:  val,
		Next: nil,
	}
	//p := this
	for this != nil {
		if this.Next == nil {
			this.Next = new
			break
		}
		this = this.Next
	}
}

func (this *SingleNode) RemoveFirst() int {
	temp := *this
	this = this.Next
	return temp.Val
}

func (this *SingleNode) String() string {
	s := ""
	p := this
	for p != nil {
		s += strconv.Itoa(p.Val) + " -> "
		if p.Next == nil {
			s += "end"
		}
		p = p.Next
	}
	return s
}

func reverse(n *SingleNode) SingleNode {
	prev := &SingleNode{}
	curr := n

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return *prev
}

func nthToLast(n int, head *SingleNode) SingleNode {
	reversed := reverse(head)
	p := &reversed
	for n > 1 {
		p = p.Next
		n--
	}
	return *p
}
