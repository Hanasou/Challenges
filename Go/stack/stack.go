package stack

// Stack is a stack. Deal with it.
type Stack []interface{}

func sliceIsEmpty(s []int) bool {
	return len(s) == 0
}

func sumFunc(a, b int) int {
	return a + b
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return false
	}
	el := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return el
}

func (s *Stack) peek() interface{} {
	return (*s)[len(*s)-1]
}

func (s *Stack) size() int {
	return len(*s)
}
