package main

// Doubly linked list
// Similar to singly linked list but each node point to Previous node in addition to next node

// Values stored in non-contiguous memory - can grow and shrink as needed

// Downsides:

// traverse list to get to data
// Retrieval is always O(n)

// List is a doubly linked list
type List struct {
	root Node
	len  int
}

// Node is a single piece of data in a List
type Node struct {
	Value      interface{}
	next, prev *Node
	list       *List
}

func (l *List) insert(n, at *Node) *Node {
	next := at.next
	at.next = n
	n.prev = at
	n.next = next
	next.prev = n
	n.list = l
	l.len++
	return n
}

// InsertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List) InsertValue(val interface{}, at *Node) *Node {
	return l.insert(&Node{Value: val}, at)
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List) PushFront(v interface{}) *Node {
	l.lazyInit()
	return l.InsertValue(v, &l.root)
}

func (l *List) remove(n *Node) *Node {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil // avoid memory leaks
	n.prev = nil // avoid memory leaks
	n.list = nil
	l.len--
	return n
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
func (l *List) Remove(n *Node) interface{} {
	if n.list == l {
		l.remove(n)
	}
	return n.Value
}

// lazyInit lazily initializes a zero List value.
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// Init initializes or clears list l.
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// Front returns the first element of list l or nil.
func (l *List) Front() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Next returns the next list element or nil.
func (n *Node) Next() *Node {
	if p := n.next; n.list != nil && p != &n.list.root {
		return p
	}
	return nil
}

// Prev returns the previous list element or nil.
func (n *Node) Prev() *Node {
	if p := n.prev; n.list != nil && n != &n.list.root {
		return p
	}
	return nil
}

// New returns an initialized list.
func New() *List { return new(List).Init() }
