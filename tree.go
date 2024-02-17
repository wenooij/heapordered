// Package heapordered provides Tree: a generic min-heap-ordered tree.
package heapordered

import (
	"slices"
)

// Tree represents a node in the tree among heap-ordered children.
//
// Tree keeps track of its own index in the child slice so we can
// call Fix when the Priority changes.
type Tree[E any] struct {
	parent    *Tree[E]
	children  []*Tree[E] // Min heap.
	heapIndex int        // Index in the parent heap.
	// Priority value for the node.
	//
	// Fix, Down or UpdatePriority should be called when the Prioirty value changes.
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
	n.children = make([]*Tree[E], 0, len(children))
	for _, e := range children {
		e.link(n)
		n.children = append(n.children, e)
	}
	n.Init()
	return n
}

func (n *Tree[E]) link(parent *Tree[E]) {
	n.parent = parent
	n.heapIndex = len(parent.children)
}

func (n *Tree[E]) unlink() {
	n.parent = nil
	n.heapIndex = 0
}

// Init fixes the child heap for this node.
func (n *Tree[E]) Init() { initHeap(n.children) }

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
	push(&parent.children, n)
}

// Grow ensures capacity for the given number of additional children.
func (n *Tree[E]) Grow(cap int) { n.children = slices.Grow(n.children, cap) }

// Len returns the number of children for this node.
func (n *Tree[E]) Len() int {
	if n == nil {
		return 0
	}
	return len(n.children)
}

// Parent returns the parent for this node or nil.
func (n *Tree[E]) Parent() *Tree[E] { return n.parent }

// At returns the element at i.
func (n *Tree[E]) At(i int) *Tree[E] { return n.children[i] }

// Min returns the minimum element or nil if none is available.
func (n *Tree[E]) Min() *Tree[E] {
	if len(n.children) == 0 {
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
func (n *Tree[E]) Fix() { fix(n.parent.children, n.heapIndex) }

// Down repairs the heap property when the priority value has increased.
//
// Fix or ReplaceElem should be called when the node's Prioirty value changes.
func (n *Tree[E]) Down() { down(n.parent.children, 0, len(n.parent.children)) }

// Remove the current node from the parent's child heap.
//
// Remove panics if n is not part of a parent heap.
func (n *Tree[E]) Remove() { remove(&n.parent.children, n.heapIndex); n.unlink() }

// Pop removes the best node from the child heap.
//
// Pop panics if n has no children.
func (n *Tree[E]) Pop() (min *Tree[E]) { child := pop(&n.children); child.unlink(); return child }

// EachChild calls f on each child of n.
func (n *Tree[E]) EachChild(f func(*Tree[E])) {
	for _, child := range n.children {
		f(child)
	}
}
