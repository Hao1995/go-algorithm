package laststoneweight

import "container/heap"

type MaxHeap []int

func (m MaxHeap) Len() int           { return len(m) }
func (m MaxHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m MaxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *MaxHeap) Push(item any)     { (*m) = append((*m), item.(int)) }
func (m *MaxHeap) Pop() any {
	item := (*m)[m.Len()-1]
	(*m) = (*m)[:m.Len()-1]
	return item
}

func lastStoneWeight(stones []int) int {
	pq := MaxHeap(stones)
	heap.Init(&pq)

	for pq.Len() > 1 {
		item1, item2 := heap.Pop(&pq).(int), heap.Pop(&pq).(int)

		if item1 != item2 {
			newItem := item1 - item2
			heap.Push(&pq, newItem)
		}
	}

	if pq.Len() == 0 {
		return 0
	}

	return heap.Pop(&pq).(int)
}
