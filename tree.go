// Package heapordered provides Tree: a generic min-heap-ordered tree.
package heapordered

import (
	"container/heap"
	"slices"
)

// Tree represents a node in the tree among heap-ordered children.
//
// Tree keeps track of its own index in the child slice so we can
// call Fix when the Priority changes.
type Tree[E any] struct {
	parent    *Tree[E]
	children  minHeap[E]
	heapIndex int // Index in the parent heap.
	// Priority value for the node.
	//
	// Fix or ReplaceElem should be called when the Prioirty value changes.
	Priority float64
	// E is a generic element payload.
	E E
}

// NewTree creates a new tree node with children.
//
// NewTree initializes a heap for the children.
func NewTree[E any](e E, priority float64, children ...*Tree[E]) *Tree[E] {
	n := &Tree[E]{E: e, Priority: priority}
	if len(children) == 0 {
		return n
	}
	n.children = make(minHeap[E], 0, len(children))
	for _, e := range children {
		e.link(n)
		n.children = append(n.children, e)
	}
	n.Init()
	return n
}

func (n *Tree[E]) link(parent *Tree[E]) {
	n.parent = parent
	n.heapIndex = parent.children.Len()
}

func (n *Tree[E]) unlink() {
	n.parent = nil
	n.heapIndex = 0
}

// Init fixes the child heap for this node.
func (n *Tree[E]) Init() { heap.Init(&n.children) }

// NewChild creates a new child node in the parent.
//
// NewChild places the child on the heap.
func (parent *Tree[E]) NewChild(e E, priority float64) *Tree[E] {
	n := &Tree[E]{E: e, Priority: priority}
	parent.NewChildTree(n)
	return n
}

// NewChild creates a new child node in the parent.
//
// NewChild places the child on the heap.
func (parent *Tree[E]) NewChildTree(n *Tree[E]) {
	n.link(parent)
	heap.Push(&parent.children, n)
}

// Grow ensures capacity for the given number of additional children.
func (n *Tree[E]) Grow(cap int) { n.children = slices.Grow(n.children, cap) }

// Len returns the number of children for this node.
func (n *Tree[E]) Len() int {
	if n == nil {
		return 0
	}
	return n.children.Len()
}

// Parent returns the parent for this node or nil.
func (n *Tree[E]) Parent() *Tree[E] {
	if n == nil {
		return nil
	}
	return n.parent
}

// Min returns the minimum element or nil if none is available.
func (n *Tree[E]) Min() *Tree[E] {
	if n == nil || n.children.Len() == 0 {
		return nil
	}
	return n.children[0]
}

// UpdatePriority replaces the priority for the current node and calls Fix to repair the heap property.
// Returns the old element.
func (n *Tree[E]) UpdatePriority(v float64) (oldValue float64) {
	oldValue = n.Priority
	n.Priority = v
	n.Fix()
	return oldValue
}

// Fix repairs the heap property for the current node in the parent's child heap.
//
// Fix or ReplaceElem should be called when the node's Prioirty value changes.
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

// Children returns the children of the current node.
//
// Note that modifying the list require corresponding calls to Fix and Init.
func (n *Tree[E]) Children() []*Tree[E] { return n.children }
