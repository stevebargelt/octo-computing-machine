package stack

import "github.com/stevebargelt/missionInterview/linkedlist"

type Stack struct {
	list    *linkedlist.List
	IsEmpty bool
}

func (s *Stack) Push(v interface{}) *Stack {
	s.list.PushFront(v)
	s.IsEmpty = false
	return s
}

func (s *Stack) Pop() interface{} {
	v := s.Peek()
	s.list.Remove(s.list.Front())
	if s.list.Len == 0 {
		s.IsEmpty = true
	}
	return v
}

func (s *Stack) Peek() interface{} {
	return s.list.Front()
}

// NewStack creates and returns a stack object
func NewStack() *Stack {
	stk := new(Stack)
	stk.list = linkedlist.New()
	stk.IsEmpty = true
	return stk
}
