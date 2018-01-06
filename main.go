package main

import (
	"fmt"
)

func main() {

	fmt.Println("---------- Linked List ----------")
	// Create a new list and put some numbers in it.
	l := New()
	n0 := l.PushFront(1)
	n2 := l.InsertValue(2, n0)
	l.InsertValue(3, n2)
	l.PushFront(0)

	// Iterate through list and print its contents.
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Println(n.Value)
	}

	fmt.Println("----------")
	l.remove(n2)

	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Println(n.Value)
	}

	fmt.Println("---------- Stack ----------")

	stk := NewStack()
	stk.Push(0)
	stk.Push(1)
	stk.Push(2)

	fmt.Printf("POP 1 = %v\n", stk.Pop())
	fmt.Printf("POP 2 = %v\n", stk.Pop())
	fmt.Printf("POP 3 = %v\n", stk.Pop())

}
