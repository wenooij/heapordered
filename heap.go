package heapordered

type minHeap[E any] []*Tree[E]

func (a minHeap[E]) Len() int           { return len(a) }
func (a minHeap[E]) Swap(i, j int)      { a[i], a[j] = a[j], a[i]; a[i].heapIndex, a[j].heapIndex = i, j }
func (a minHeap[E]) Less(i, j int) bool { return a[i].Priority < a[j].Priority }
func (a *minHeap[E]) Push(x any)        { *a = append(*a, x.(*Tree[E])) }
func (a *minHeap[E]) Pop() any          { n := len(*a) - 1; x := (*a)[n]; *a = (*a)[:n]; return x }
