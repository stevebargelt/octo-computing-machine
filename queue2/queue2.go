package queue2

import (
	"github.com/stevebargelt/missionInterview/stack"
)

type queue2 struct {
	inbox  *stack.Stack
	outbox *stack.Stack
}

func (q *queue2) Enqueue(v interface{}) *queue2 {
	q.inbox.Push(v)
	return q
}

func (q *queue2) Dequeue() interface{} {
	if q.outbox.IsEmpty {
		for !q.inbox.IsEmpty {
			q.outbox.Push(q.inbox.Pop())
		}

	}
	return q.outbox.Pop()
}

func (q queue2) Peek() interface{} {
	return q.inbox.Peek()
}

// NewStack creates and returns a stack object
func NewQueue2() *queue2 {
	kew2 := new(queue2)
	kew2.inbox = stack.NewStack()
	kew2.outbox = stack.NewStack()
	return kew2
}
