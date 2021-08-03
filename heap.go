package main

type HandHeap []Hand

func (h HandHeap) Len() int           { return len(h) }
func (h HandHeap) Less(i, j int) bool { return h[i].getValue() < h[j].getValue() }
func (h HandHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HandHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Hand))
}

func (h *HandHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
