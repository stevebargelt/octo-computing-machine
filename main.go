package main

import (
	"fmt"
)

func main() {

	// Create a new list and put some numbers in it.
	l := New()
	n0 := l.PushFront(0)
	fmt.Print(n0)
	// n1 := l.insert(&Node{Value: 1}, n0)
	// n2 := l.insert(&Node{Value: 2}, n1)
	// l.InsertValue(33, n2)

	// Iterate through list and print its contents.
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Println(n.Value)
	}

	//l.remove(n3)

	// // Iterate through list and print its contents.
	// for n := l.Front(); n != nil; n = n.Next() {
	// 	fmt.Println(n.Value)
	// }

}
