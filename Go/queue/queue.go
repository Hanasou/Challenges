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
