package heapordered

func swap[E any](h []*Tree[E], i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIndex, h[j].heapIndex = i, j
}

// Code adapted from containers/heap.go.

func initHeap[E any](h []*Tree[E]) {
	n := len(h)
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

func push[E any](h *[]*Tree[E], e *Tree[E]) {
	*h = append(*h, e) // Push(x)
	up(*h, len(*h)-1)
}

func pop[E any](h *[]*Tree[E]) (e *Tree[E]) {
	n := len(*h) - 1
	swap(*h, 0, n)
	down(*h, 0, n)
	e = (*h)[n]
	*h = (*h)[:n] // Pop()
	return e
}

func remove[E any](h *[]*Tree[E], i int) (e E) {
	n := len(*h) - 1
	if n != i {
		swap(*h, i, n)
		if !down(*h, i, n) {
			up(*h, i)
		}
	}
	e = (*h)[n].E
	*h = (*h)[:n] // Pop()
	return e
}

func fix[E any](h []*Tree[E], i int) {
	if !down(h, i, len(h)) {
		up(h, i)
	}
}

func up[E any](h []*Tree[E], j int) {
	for {
		i := (j - 1) / 2                              // Parent.
		if i == j || h[j].Priority >= h[i].Priority { // !Less(j, i)
			break
		}
		swap(h, i, j)
		j = i
	}
}

func down[E any](h []*Tree[E], i0 int, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1                                                      // Left child.
		if j2 := j1 + 1; j2 < n && h[j2].Priority < h[j1].Priority { // Less(j2, j1)
			j = j2 // = 2*i + 2  // right child
		}
		if h[j].Priority >= h[i].Priority { // !Less(j, i)
			break
		}
		swap(h, i, j)
		i = j
	}
	return i > i0
}
