package heapordered

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type p float64

func (p p) Priority() float64 { return float64(p) }

func TestNewTreeNoChildren(t *testing.T) {
	got := NewTree(p(0))
	want := &Tree[p]{elem: p(0)}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestNewTreeNoChildren(): got diff:\n%s", diff)
	}
}

func TestNewTree(t *testing.T) {
	got := NewTree(p(0), NewTree(p(6), NewTree(p(7))), NewTree(p(5)), NewTree(p(4)), NewTree(p(3)), NewTree(p(2)), NewTree(p(1)))
	want := &Tree[p]{
		parent:    nil,
		heapIndex: 0,
		elem:      p(0),
	}
	want.children = []*Tree[p]{
		{parent: want, heapIndex: 0, elem: p(1)},
		{parent: want, heapIndex: 1, elem: p(2)},
		{parent: want, heapIndex: 2, elem: p(4)},
		{parent: want, heapIndex: 3, elem: p(3)},
		{parent: want, heapIndex: 4, elem: p(5)},
		{parent: want, heapIndex: 5, elem: p(6)},
	}
	want.children[5].children = []*Tree[p]{
		{parent: want.children[5], heapIndex: 0, elem: p(7)},
	}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestNewTree(): got diff:\n%s", diff)
	}
}

func TestNewChildNoChildren(t *testing.T) {
	n := NewTree(p(0))
	gotChild := n.NewChild(p(1))
	wantChild := &Tree[p]{parent: n, elem: p(1)}
	if diff := cmp.Diff(*wantChild, *gotChild, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestNewChildNoChildren(): got child diff:\n%s", diff)
	}
}

func TestNewChild(t *testing.T) {
	got := NewTree(p(0), NewTree(p(5)), NewTree(p(4)), NewTree(p(3)))
	got.NewChild(p(2))
	got.NewChild(p(1))
	subTree := NewTree(p(6), NewTree(p(7)))
	got.NewChildTree(subTree)
	want := &Tree[p]{
		parent:    nil,
		heapIndex: 0,
		elem:      p(0),
	}
	want.children = []*Tree[p]{
		{parent: want, heapIndex: 0, elem: p(1)},
		{parent: want, heapIndex: 1, elem: p(2)},
		{parent: want, heapIndex: 2, elem: p(5)},
		{parent: want, heapIndex: 3, elem: p(4)},
		{parent: want, heapIndex: 4, elem: p(3)},
		{parent: want, heapIndex: 5, elem: p(6)},
	}
	want.children[5].children = []*Tree[p]{
		{parent: want.children[5], heapIndex: 0, elem: p(7)},
	}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestNewChild(): got diff:\n%s", diff)
	}
}

func TestLenNil(t *testing.T) {
	var n *Tree[p]
	got := n.Len()
	if want := 0; want != got {
		t.Errorf("TestLenNil(): got Len %d, want %d", got, want)
	}
}

func TestLen(t *testing.T) {
	n := NewTree[p](p(0), NewTree(p(1)), NewTree(p(2)), NewTree(p(3)))
	got := n.Len()
	if want := 3; want != got {
		t.Errorf("TestLen(): got Len %d, want %d", got, want)
	}
}

func TestParentNil(t *testing.T) {
	var n *Tree[p]
	got := n.Parent()
	if want := (*Tree[p])(nil); want != got {
		t.Errorf("TestParentNil(): got Len %v, want %v", got, want)
	}
}

func TestParent(t *testing.T) {
	n := NewTree[p](p(0), NewTree(p(1)), NewTree(p(2)), NewTree(p(3)))
	got := n.children[0].Parent()
	if want := n; want != got {
		t.Errorf("TestParent(): got Len %v, want %v", got, want)
	}
}

func TestElemNil(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("TestElemNil(): expected panic")
		}
	}()
	var n *Tree[p]
	n.Elem()
}

func TestElem(t *testing.T) {
	n := NewTree(p(0))
	got := n.Elem()
	if diff := cmp.Diff(p(0), got); diff != "" {
		t.Errorf("TestElem(): got diff:\n%s", diff)
	}
}

func TestMinNil(t *testing.T) {
	var n *Tree[p]
	got := n.Min()
	want := (*Tree[p])(nil)
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestMinNil(): got diff:\n%s", diff)
	}
}

func TestMinEmpty(t *testing.T) {
	n := NewTree(p(0))
	got := n.Min()
	want := (*Tree[p])(nil)
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestMinEmpty(): got diff:\n%s", diff)
	}
}

func TestMin(t *testing.T) {
	n := NewTree(p(0), NewTree(p(-1)), NewTree(p(-2)), NewTree(p(-1)))
	got := n.Min()
	want := &Tree[p]{
		parent:    n,
		heapIndex: 0,
		elem:      p(-2),
	}
	if diff := cmp.Diff(*want, *got, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestMin(): got diff:\n%s", diff)
	}
}

func TestReplaceElem(t *testing.T) {
	n := NewTree(p(0), NewTree(p(5)), NewTree(p(4)), NewTree(p(3)), NewTree(p(2)), NewTree(p(1)))
	// p(1)->p(6).
	if got := n.children[0].ReplaceElem(p(6)); got != p(1) {
		t.Fatalf("TestReplaceElem(): ReplaceElem got %v, want %v", got, p(1))
	}
	// p(5)->p(0).
	if got := n.children[3].ReplaceElem(p(0)); got != p(5) {
		t.Fatalf("TestReplaceElem(): ReplaceElem got %v, want %v", got, p(5))
	}
	want := &Tree[p]{
		parent:    nil,
		heapIndex: 0,
		elem:      p(0),
	}
	want.children = []*Tree[p]{
		{parent: want, heapIndex: 0, elem: p(0)},
		{parent: want, heapIndex: 1, elem: p(2)},
		{parent: want, heapIndex: 2, elem: p(3)},
		{parent: want, heapIndex: 3, elem: p(4)},
		{parent: want, heapIndex: 4, elem: p(6)},
	}
	if diff := cmp.Diff(want, n, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestReplaceElem(): got diff:\n%s", diff)
	}
}

func TestRemoveNoParent(t *testing.T) {
	got := NewTree(p(0))
	got.Remove()
	want := &Tree[p]{parent: nil, heapIndex: 0, elem: p(0)}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestRemoveNoParent(): got diff:\n%s", diff)
	}
}

func TestRemove(t *testing.T) {
	n := NewTree(p(0), NewTree(p(5)), NewTree(p(4)), NewTree(p(3)), NewTree(p(2)), NewTree(p(1)))
	// Remove p(1).
	got := n.children[0]
	n.children[0].Remove()
	want := &Tree[p]{parent: nil, heapIndex: 0, elem: p(1)}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestRemove(): got Remove() #0 diff:\n%s", diff)
	}
	// Remove p(5).
	got = n.children[3]
	n.children[3].Remove()
	want = &Tree[p]{parent: nil, heapIndex: 0, elem: p(5)}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestRemove(): got Remove() #1 diff:\n%s", diff)
	}
	// Check remaining elements.
	want = &Tree[p]{
		parent:    nil,
		heapIndex: 0,
		elem:      p(0),
	}
	want.children = []*Tree[p]{
		{parent: want, heapIndex: 0, elem: p(2)},
		{parent: want, heapIndex: 1, elem: p(4)},
		{parent: want, heapIndex: 2, elem: p(3)},
	}
	if diff := cmp.Diff(want, n, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestRemove(): got result diff:\n%s", diff)
	}
}

func TestPopEmpty(t *testing.T) {
	n := NewTree(p(0))
	want := (*Tree[p])(nil)
	got := n.Pop()
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestPopEmpty(): got diff:\n%s", diff)
	}
}

func TestPop(t *testing.T) {
	n := NewTree(p(0), NewTree(p(5)), NewTree(p(4)), NewTree(p(3)), NewTree(p(2)), NewTree(p(1)))
	// Pop p(1).
	got := n.Pop()
	want := &Tree[p]{parent: nil, heapIndex: 0, elem: p(1)}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestPop(): got Pop()#0 diff:\n%s", diff)
	}
	// Pop p(2).
	got = n.Pop()
	want = &Tree[p]{parent: nil, heapIndex: 0, elem: p(2)}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestPop(): got Pop()#1 diff:\n%s", diff)
	}
	// Pop p(3).
	got = n.Pop()
	want = &Tree[p]{parent: nil, heapIndex: 0, elem: p(3)}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestPop(): got Pop()#2 diff:\n%s", diff)
	}
	want = &Tree[p]{
		parent:    nil,
		heapIndex: 0,
		elem:      p(0),
	}
	want.children = []*Tree[p]{
		{parent: want, heapIndex: 0, elem: p(4)},
		{parent: want, heapIndex: 1, elem: p(5)},
	}
	if diff := cmp.Diff(want, n, cmp.AllowUnexported(Tree[p]{})); diff != "" {
		t.Errorf("TestRemove(): got result diff:\n%s", diff)
	}
}
