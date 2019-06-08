package tree

type tree struct {
	root Node
	Len  int
}

// Node is a single piece of data in a List
type Node struct {
	Value    interface{}
	Children []Node
}
