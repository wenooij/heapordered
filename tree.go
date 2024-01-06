// Package heapordered provides Tree: a generic min-heap-ordered tree.
package heapordered

import "container/heap"

// Priority is an interface for min-heap priority.
type Prioirty interface {
	// Prioirty returns the value used for min-heap priority.
	Prioirty() float64
}

// Tree represents a node in the tree among heap-ordered children.
//
// Tree keeps track of its own index in the child slice so we can
// call Fix when the Priority changes.
type Tree[E Prioirty] struct {
	parent    *Tree[E]
	children  minHeap[E]
	heapIndex int // Index in the parent heap.
	elem      E   // Payload.
}

// NewTree creates a new tree node with children.
//
// NewTree initializes a heap for the children.
func NewTree[E Prioirty](e E, children ...*Tree[E]) *Tree[E] {
	n := &Tree[E]{elem: e, children: make(minHeap[E], 0, len(children))}
	for _, e := range children {
		n.link(e)
		n.children = append(n.children, e)
	}
	heap.Init((*minHeap[E])(&n.children))
	return n
}

func (parent *Tree[E]) link(n *Tree[E]) {
	n.parent = parent
	n.heapIndex = parent.children.Len()
}

func (n *Tree[E]) unlink() {
	n.parent = nil
	n.heapIndex = 0
}

// NewChild creates a new child node in the parent.
//
// NewChild places the child on the heap.
func (parent *Tree[E]) NewChild(e E) *Tree[E] {
	n := &Tree[E]{elem: e}
	parent.NewChildTree(n)
	return n
}

// NewChild creates a new child node in the parent.
//
// NewChild places the child on the heap.
func (parent *Tree[E]) NewChildTree(n *Tree[E]) {
	parent.link(n)
	heap.Push(&parent.children, n)
}

// Elem returns the Prioirty element data for this node.
func (n *Tree[E]) Elem() E { return n.elem }

// Min returns the minimum element or false if none.
func (n *Tree[E]) Min() *Tree[E] {
	if n == nil || n.children.Len() == 0 {
		return nil
	}
	return n.children[0]
}

// ReplaceElem replaces the element for the current node and calls Fix to repair the heap.
// Returns the old element.
func (n *Tree[E]) ReplaceElem(e E) (old E) {
	old = n.elem
	n.elem = e
	n.Fix()
	return old
}

// Fix repairs the heap property for the current node in the parent's child heap.
//
// Fix should be called when the node's Prioirty value changes.
// Prefer ReplaceElem when possible.
func (n *Tree[E]) Fix() {
	if n.parent != nil {
		heap.Fix(&n.parent.children, n.heapIndex)
	}
}

// Remove the current node from the parent's child heap.
func (n *Tree[E]) Remove() {
	if n.parent != nil {
		heap.Remove(&n.parent.children, n.heapIndex)
		n.unlink()
	}
}

// Pop the best node from among this node's child heap.
func (n *Tree[E]) Pop() (min *Tree[E]) {
	if n.children.Len() == 0 {
		return nil
	}
	child := heap.Pop(&n.children).(*Tree[E])
	child.unlink()
	return child
}
