package kthlargestelementinanarray

import "container/heap"

type MaxHeap struct {
	data []int
}

func (this MaxHeap) Len() int           { return len(this.data) }
func (this MaxHeap) Less(i, j int) bool { return this.data[i] > this.data[j] }
func (this MaxHeap) Swap(i, j int)      { this.data[i], this.data[j] = this.data[j], this.data[i] }
func (this *MaxHeap) Push(item any)     { this.data = append(this.data, item.(int)) }
func (this *MaxHeap) Pop() any {
	var item int
	item, this.data = this.data[len(this.data)-1], this.data[:len(this.data)-1]
	return item
}

func findKthLargest(nums []int, k int) int {
	pq := &MaxHeap{}
	heap.Init(pq)

	// O(n*logn)
	for _, num := range nums {
		heap.Push(pq, num)
	}

	// O(k*logn)
	var ans int
	for i := 0; i < k; i++ {
		ans = heap.Pop(pq).(int)
	}

	return ans
}
