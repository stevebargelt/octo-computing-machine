package queue

import "github.com/stevebargelt/missionInterview/linkedlist"

type queue struct {
	list *linkedlist.List
}

func (q *queue) Enqueue(v interface{}) *queue {
	q.list.PushBack(v)
	return q
}

func (q *queue) Dequeue() interface{} {
	v := q.Peek()
	q.list.Remove(q.list.Front())
	return v
}

func (q queue) Peek() interface{} {
	return q.list.Front()
}

// NewStack creates and returns a stack object
func NewQueue() *queue {
	kew := new(queue)
	kew.list = linkedlist.New()
	return kew
}
