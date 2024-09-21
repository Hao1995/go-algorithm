package kclosest

import "container/heap"

func distance(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

type MaxHeap [][]int

func (this MaxHeap) Len() int           { return len(this) }
func (this MaxHeap) Less(i, j int) bool { return distance(this[i]) > distance(this[j]) }
func (this MaxHeap) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this *MaxHeap) Push(item any)     { *this = append(*this, item.([]int)) }
func (this *MaxHeap) Pop() any {
	item := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return item
}

func kClosestV2(points [][]int, k int) [][]int {
	var pq MaxHeap
	for _, point := range points {
		heap.Push(&pq, point)
		if len(pq) > k {
			heap.Pop(&pq)
		}
	}
	return pq
}
