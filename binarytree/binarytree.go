package binarytree

// Node is a single piece of data in a List
type Node struct {
	Value interface{}
	Left  *Node
	Right *Node
}

// New returns an initialized list.
func New(val interface{}) *Node {
	newNode := new(Node).Init()
	newNode.Value = val
	return newNode
}

// Init initializes or clears list l.
func (b *Node) Init() *Node {
	b.Value = 0
	b.Left = nil
	b.Right = nil
	return b
}

func (b *Node) AddLeft(val interface{}) *Node {
	newNode := New(val)
	b.Left = newNode
	return newNode
}

func (b *Node) AddRight(val interface{}) *Node {
	newNode := New(val)
	b.Right = newNode
	return newNode
}
