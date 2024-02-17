package heapordered

import "testing"

func makeBigTree(root *Tree[int], branchFactor, depth int) {
	if depth <= 0 {
		return
	}
	for i := 0; i < branchFactor; i++ {
		child := root.NewChild(-i, float64(-i))
		makeBigTree(child, branchFactor, depth-1)
	}
}

var BenchResult int

func BenchmarkTree(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		root := NewTree[int](0, 0)
		makeBigTree(root, 10, 4)
		n = root.Len()
	}
	BenchResult = n
}
