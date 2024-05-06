package kthlargestelementinanarray

import "container/heap"

type MaxHeap []int

func (this MaxHeap) Len() int           { return len(this) }
func (this MaxHeap) Less(i, j int) bool { return this[i] > this[j] }
func (this MaxHeap) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this *MaxHeap) Push(item any)     { *this = append(*this, item.(int)) }
func (this *MaxHeap) Pop() any {
	old := *this
	var item int
	item, old = old[len(old)-1], old[:len(old)-1]
	*this = old
	return item
}

func findKthLargest(nums []int, k int) int {
	// O(n): https://pkg.go.dev/container/heap#Init
	pq := MaxHeap(nums)
	heap.Init(&pq)

	// O(k*logn)
	var ans int
	for i := 0; i < k; i++ {
		ans = heap.Pop(&pq).(int)
	}

	return ans
}
