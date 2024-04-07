package minheap

type Item struct {
	val int
}

type PQ []Item

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].val < pq[j].val }
func (pq PQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x any)        { *pq = append(*pq, x.(Item)) }
func (pq *PQ) Pop() any {
	old := *pq
	l := len(old)
	x := old[l-1]
	*pq = old[:l-1]
	return x
}
