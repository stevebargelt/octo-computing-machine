package main

import (
	"fmt"

	"github.com/stevebargelt/missionInterview/binarytree"
	"github.com/stevebargelt/missionInterview/fibonacci"
	"github.com/stevebargelt/missionInterview/linkedlist"
	"github.com/stevebargelt/missionInterview/queue"
	"github.com/stevebargelt/missionInterview/stack"
)

func main() {

	fmt.Println("---------- Linked List ----------")
	// Create a new list and put some numbers in it.
	l := linkedlist.New()
	n0 := l.PushFront(1)
	n2 := l.InsertValue(2, n0)
	l.InsertValue(3, n2)
	l.PushFront(0)

	// Iterate through list and print its contents.
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Println(n.Value)
	}

	fmt.Println("----------")
	l.Remove(n2)

	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Println(n.Value)
	}

	fmt.Println("---------- Stack ----------")

	stk := stack.NewStack()
	stk.Push(0)
	stk.Push(1)
	stk.Push(2)

	fmt.Printf("POP 1 = %v\n", stk.Pop())
	fmt.Printf("POP 2 = %v\n", stk.Pop())
	fmt.Printf("POP 3 = %v\n", stk.Pop())

	fmt.Println("---------- Queue ----------")

	q := queue.NewQueue()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)

	fmt.Printf("Dequeue 1 = %v\n", q.Dequeue())
	fmt.Printf("Dequeue 2 = %v\n", q.Dequeue())
	fmt.Printf("Dequeue 3 = %v\n", q.Dequeue())

	// fmt.Println("---------- Queue2 ----------")

	// q2 := queue2.NewQueue2()
	// q2.Enqueue(0)
	// q2.Enqueue(1)
	// q2.Enqueue(2)

	// fmt.Printf("Dequeue 1 = %d\n", q2.Dequeue())
	// fmt.Printf("Dequeue 2 = %d\n", q2.Dequeue())
	// fmt.Printf("Dequeue 3 = %d\n", q2.Dequeue())

	fmt.Println("---------- Binary Tree ----------")

	root := binarytree.New(1)
	r1 := root.AddRight(2)
	l1 := root.AddLeft(3)

	r1.AddRight(4).AddRight(5)
	l1.AddLeft(7)
	fmt.Print(root)

	fmt.Printf("\n\n---------- Fibonacci ----------\n")

	fmt.Printf("Sequence: %v\n", fibonacci.Fibs(20))
	fmt.Printf("Nth Fib: %v\n", fibonacci.Fib(20))

}
