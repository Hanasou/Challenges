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

func (this *SingleNode) PointerTest() {
	for this != nil {
		this = this.Next
	}
}

func (this *SingleNode) Swap(a *SingleNode, b *SingleNode) {
	// Find nodes a and b
	// If they don't exist, then we can return an error or something but let's not worry about that for now
	for this != nil {
		this = this.Next
	}
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

func Midpoint(head *SingleNode) *SingleNode {
	slow := head
	fast := head

	for fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
