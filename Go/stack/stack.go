package stack

// Stack is a stack. Deal with it.
type stack []interface{}

func sliceIsEmpty(s []int) bool {
	return len(s) == 0
}

func sumFunc(a, b int) int {
	return a + b
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(item interface{}) {
	*s = append(*s, item)
}

func (s *stack) pop() interface{} {
	if s.isEmpty() {
		return false
	}
	el := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return el
}

func (s *stack) peek() interface{} {
	return (*s)[len(*s)-1]
}

func (s *stack) size() int {
	return len(*s)
}
