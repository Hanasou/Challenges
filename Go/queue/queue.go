package queue

type Queue []interface{}

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
			_, ele1 := Dequeue(q1)
			Enqueue(combined, ele1)
		}
		if !IsEmpty(q2) {
			_, ele2 := Dequeue(q2)
			Enqueue(combined, ele2)
		}
	}

	// Add the rest of the elements from either queue into combined
	// if one of them still has elements
	for !IsEmpty(q1) {
		_, ele := Dequeue(q1)
		Enqueue(combined, ele)
	}
	for !IsEmpty(q2) {
		_, ele := Dequeue(q2)
		Enqueue(combined, ele)
	}
	return combined
}
