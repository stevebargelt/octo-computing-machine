package main

type stack struct {
	list *List
}

func (s stack) Push(v interface{}) stack {
	s.list.PushFront(v)
	return s
}

func (s stack) Pop() interface{} {
	v := s.list.Front().Value
	s.list.Remove(s.list.Front())
	return v
}

// NewStack creates and returns a stack object
func NewStack() *stack {
	stk := new(stack)
	stk.list = New()
	return stk
}
