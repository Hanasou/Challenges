package queue

import "sandbox/stack"

// Note that the implementation here is not objected oriented
// The implementation of the stack is though
type Queue []interface{}

type StackQueue struct {
	s1 *stack.Stack
	s2 *stack.Stack
}

func (sq *StackQueue) Enqueue(ele interface{}) {
	sq.s1.Push(ele)
}

func (sq *StackQueue) Dequeue() {
	for !sq.s1.IsEmpty() {
		ele := sq.s1.Pop()
		sq.s2.Push(ele)
	}
	sq.s2.Pop()
}

func Enqueue(q Queue, x interface{}) Queue {
	new := append(q, x)
	return new
}

func Dequeue(q Queue) (Queue, interface{}) {
	if len(q) == 0 {
		return q, nil
	}
	lastIndex := len(q) - 1
	return q[:lastIndex], q[lastIndex]
}

func IsEmpty(q Queue) bool {
	return len(q) == 0
}

func Peek(q Queue) interface{} {
	return q[len(q)-1]
}

func Weave(q1 Queue, q2 Queue) Queue {
	combined := Queue{}
	// Alternate between both queues while they both have elements
	for !IsEmpty(q1) && !IsEmpty(q2) {
		if !IsEmpty(q1) {
			var ele1 interface{}
			q1, ele1 = Dequeue(q1)
			combined = Enqueue(combined, ele1)
		}
		if !IsEmpty(q2) {
			var ele2 interface{}
			q2, ele2 = Dequeue(q2)
			combined = Enqueue(combined, ele2)
		}
	}

	// Add the rest of the elements from either queue into combined
	// if one of them still has elements
	for !IsEmpty(q1) {
		var ele interface{}
		q1, ele = Dequeue(q1)
		combined = Enqueue(combined, ele)
	}
	for !IsEmpty(q2) {
		var ele interface{}
		q2, ele = Dequeue(q2)
		combined = Enqueue(combined, ele)
	}
	return combined
}

func WeaveTest() Queue {
	q1 := Queue{1, 2, 3, 4, 5}
	q2 := Queue{"Hi", "Bye", "Nothing Left"}
	return Weave(q1, q2)
}
