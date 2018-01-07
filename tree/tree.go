package tree

// List is a doubly linked list
type tree struct {
	root TreeNode
	Len  int
}

// TreeNode is a single piece of data in a List
type TreeNode struct {
	Value    interface{}
	Children []TreeNode
}
