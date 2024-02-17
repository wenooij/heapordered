package heapordered

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewTreeNoChildren(t *testing.T) {
	got := NewTree[int](0, 0)
	want := &Tree[int]{E: 0, Priority: 0}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[int]{})); diff != "" {
		t.Errorf("TestNewTreeNoChildren(): got diff:\n%s", diff)
	}
}

func TestNewTree(t *testing.T) {
	got := NewTree(0, 0, NewTree(6, 6, NewTree(7, 7)), NewTree(5, 5), NewTree(4, 4), NewTree(3, 3), NewTree(2, 2), NewTree(1, 1))
	want := &Tree[int]{
		parent:    nil,
		heapIndex: 0,
		E:         0,
		Priority:  0,
	}
	want.children = []*Tree[int]{
		{parent: want, heapIndex: 0, E: 1, Priority: 1},
		{parent: want, heapIndex: 1, E: 2, Priority: 2},
		{parent: want, heapIndex: 2, E: 4, Priority: 4},
		{parent: want, heapIndex: 3, E: 3, Priority: 3},
		{parent: want, heapIndex: 4, E: 5, Priority: 5},
		{parent: want, heapIndex: 5, E: 6, Priority: 6},
	}
	want.children[5].children = []*Tree[int]{
		{parent: want.children[5], heapIndex: 0, E: 7, Priority: 7},
	}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[int]{})); diff != "" {
		t.Errorf("TestNewTree(): got diff:\n%s", diff)
	}
}

func TestNewChildNoChildren(t *testing.T) {
	n := NewTree(0, 0)
	gotChild := n.NewChild(1, 1)
	wantChild := &Tree[int]{parent: n, E: 1, Priority: 1}
	if diff := cmp.Diff(*wantChild, *gotChild, cmp.AllowUnexported(Tree[int]{})); diff != "" {
		t.Errorf("TestNewChildNoChildren(): got child diff:\n%s", diff)
	}
}

func TestNewChild(t *testing.T) {
	got := NewTree(0, 0, NewTree(5, 5), NewTree(4, 4), NewTree(3, 3))
	got.NewChild(2, 2)
	got.NewChild(1, 1)
	subTree := NewTree(6, 6, NewTree(7, 7))
	got.NewChildTree(subTree)
	want := &Tree[int]{
		parent:    nil,
		heapIndex: 0,
		E:         0,
		Priority:  0,
	}
	want.children = []*Tree[int]{
		{parent: want, heapIndex: 0, E: 1, Priority: 1},
		{parent: want, heapIndex: 1, E: 2, Priority: 2},
		{parent: want, heapIndex: 2, E: 5, Priority: 5},
		{parent: want, heapIndex: 3, E: 4, Priority: 4},
		{parent: want, heapIndex: 4, E: 3, Priority: 3},
		{parent: want, heapIndex: 5, E: 6, Priority: 6},
	}
	want.children[5].children = []*Tree[int]{
		{parent: want.children[5], heapIndex: 0, E: 7, Priority: 7},
	}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[int]{})); diff != "" {
		t.Errorf("TestNewChild(): got diff:\n%s", diff)
	}
}

func TestLenNil(t *testing.T) {
	var n *Tree[int]
	got := n.Len()
	if want := 0; want != got {
		t.Errorf("TestLenNil(): got Len %d, want %d", got, want)
	}
}

func TestLen(t *testing.T) {
	n := NewTree[int](0, 0, NewTree(1, 1), NewTree(2, 2), NewTree(3, 3))
	got := n.Len()
	if want := 3; want != got {
		t.Errorf("TestLen(): got Len %d, want %d", got, want)
	}
}

func TestParentNil(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("TestParentNil(): expected panic, got no panic")
		}
	}()
	_ = (*Tree[int])(nil).Parent()
}

func TestParent(t *testing.T) {
	n := NewTree[int](0, 0, NewTree(1, 1), NewTree(2, 2), NewTree(3, 3))
	got := n.children[0].Parent()
	if want := n; want != got {
		t.Errorf("TestParent(): got Len %v, want %v", got, want)
	}
}

func TestMinNil(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("TestMinNil(): expected panic, got no panic")
		}
	}()
	_ = (*Tree[int])(nil).Min()
}

func TestMinEmpty(t *testing.T) {
	n := NewTree(0, 0)
	got := n.Min()
	want := (*Tree[int])(nil)
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[int]{})); diff != "" {
		t.Errorf("TestMinEmpty(): got diff:\n%s", diff)
	}
}

func TestMin(t *testing.T) {
	n := NewTree(0, 0, NewTree(-1, -1), NewTree(-2, -2), NewTree(-1, -1))
	got := n.Min()
	want := &Tree[int]{
		parent:    n,
		heapIndex: 0,
		E:         -2,
		Priority:  -2,
	}
	if diff := cmp.Diff(*want, *got, cmp.AllowUnexported(Tree[int]{})); diff != "" {
		t.Errorf("TestMin(): got diff:\n%s", diff)
	}
}

func TestUpdatePriority(t *testing.T) {
	n := NewTree(0, 0, NewTree(5, 5), NewTree(4, 4), NewTree(3, 3), NewTree(2, 2), NewTree(1, 1))
	// Priority 1 -> 6.
	if got := n.children[0].UpdatePriority(6); got != 1 {
		t.Fatalf("TestUpdatePriority(): UpdatePriority got oldValue %v, want %v", got, 1)
	}
	// Priority 5 -> 0.
	if got := n.children[3].UpdatePriority(0); got != 5 {
		t.Fatalf("TestUpdatePriority(): UpdatePriority got oldValue %v, want %v", got, 5)
	}
	want := &Tree[int]{
		parent:    nil,
		heapIndex: 0,
		E:         0,
		Priority:  0,
	}
	want.children = []*Tree[int]{
		{parent: want, heapIndex: 0, E: 5, Priority: 0},
		{parent: want, heapIndex: 1, E: 2, Priority: 2},
		{parent: want, heapIndex: 2, E: 3, Priority: 3},
		{parent: want, heapIndex: 3, E: 4, Priority: 4},
		{parent: want, heapIndex: 4, E: 1, Priority: 6},
	}
	if diff := cmp.Diff(want, n, cmp.AllowUnexported(Tree[int]{})); diff != "" {
		t.Errorf("TestUpdatePriority(): got diff:\n%s", diff)
	}
}

func TestRemoveNoParent(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("TestRemoveNoParent(): Expected panic, got no panic")
		}
	}()
	NewTree(0, 0).Remove()
}

func TestRemove(t *testing.T) {
	n := NewTree(0, 0, NewTree(5, 5), NewTree(4, 4), NewTree(3, 3), NewTree(2, 2), NewTree(1, 1))
	// Remove 1.
	got := n.children[0]
	n.children[0].Remove()
	want := &Tree[int]{parent: nil, heapIndex: 0, E: 1, Priority: 1}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[int]{})); diff != "" {
		t.Errorf("TestRemove(): got Remove() #0 diff:\n%s", diff)
	}
	// Remove 5.
	got = n.children[3]
	n.children[3].Remove()
	want = &Tree[int]{parent: nil, heapIndex: 0, E: 5, Priority: 5}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[int]{})); diff != "" {
		t.Errorf("TestRemove(): got Remove() #1 diff:\n%s", diff)
	}
	// Check remaining elements.
	want = &Tree[int]{
		parent:    nil,
		heapIndex: 0,
		E:         0,
		Priority:  0,
	}
	want.children = []*Tree[int]{
		{parent: want, heapIndex: 0, E: 2, Priority: 2},
		{parent: want, heapIndex: 1, E: 4, Priority: 4},
		{parent: want, heapIndex: 2, E: 3, Priority: 3},
	}
	if diff := cmp.Diff(want, n, cmp.AllowUnexported(Tree[int]{})); diff != "" {
		t.Errorf("TestRemove(): got result diff:\n%s", diff)
	}
}

func TestPopEmpty(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("TestPopEmpty(): expected panic, got no panic")
		}
	}()
	_ = NewTree(0, 0).Pop()
}

func TestPop(t *testing.T) {
	n := NewTree(0, 0, NewTree(5, 5), NewTree(4, 4), NewTree(3, 3), NewTree(2, 2), NewTree(1, 1))
	// Pop p(1).
	got := n.Pop()
	want := &Tree[int]{parent: nil, heapIndex: 0, E: 1, Priority: 1}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[int]{})); diff != "" {
		t.Errorf("TestPop(): got Pop()#0 diff:\n%s", diff)
	}
	// Pop p(2).
	got = n.Pop()
	want = &Tree[int]{parent: nil, heapIndex: 0, E: 2, Priority: 2}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[int]{})); diff != "" {
		t.Errorf("TestPop(): got Pop()#1 diff:\n%s", diff)
	}
	// Pop p(3).
	got = n.Pop()
	want = &Tree[int]{parent: nil, heapIndex: 0, E: 3, Priority: 3}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[int]{})); diff != "" {
		t.Errorf("TestPop(): got Pop()#2 diff:\n%s", diff)
	}
	want = &Tree[int]{
		parent:    nil,
		heapIndex: 0,
		E:         0,
		Priority:  0,
	}
	want.children = []*Tree[int]{
		{parent: want, heapIndex: 0, E: 4, Priority: 4},
		{parent: want, heapIndex: 1, E: 5, Priority: 5},
	}
	if diff := cmp.Diff(want, n, cmp.AllowUnexported(Tree[int]{})); diff != "" {
		t.Errorf("TestRemove(): got result diff:\n%s", diff)
	}
}
