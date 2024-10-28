package taskscheduler

import "container/heap"

type LetterCount struct {
	Letter byte
	Count  int
}

type MaxHeap []LetterCount

func (this MaxHeap) Len() int           { return len(this) }
func (this MaxHeap) Less(i, j int) bool { return this[i].Count > this[j].Count }
func (this MaxHeap) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this *MaxHeap) Push(item any)     { *this = append(*this, item.(LetterCount)) }
func (this *MaxHeap) Pop() any {
	item := (*this)[this.Len()-1]
	*this = (*this)[:this.Len()-1]
	return item
}

func leastInterval(tasks []byte, n int) int {
	countMap := make(map[byte]int)
	for _, t := range tasks {
		countMap[t]++
	}

	pq := MaxHeap(make([]LetterCount, 0, len(countMap)))
	heap.Init(&pq)
	for c, count := range countMap {
		heap.Push(&pq, LetterCount{
			Letter: c,
			Count:  count,
		})
	}

	var ans int
	for pq.Len() > 0 {
		var tmpPQ MaxHeap

		var count int
		slots := n + 1
		for i := 0; i < slots && pq.Len() > 0; i++ {
			item := heap.Pop(&pq).(LetterCount)
			item.Count--
			if item.Count > 0 {
				tmpPQ = append(tmpPQ, item)
			}
			count++
		}
		ans += count

		// idle: add the remaining number
		if tmpPQ.Len() > 0 {
			ans += slots - count
		}

		for _, item := range tmpPQ {
			heap.Push(&pq, item)
		}
	}

	return ans
}
