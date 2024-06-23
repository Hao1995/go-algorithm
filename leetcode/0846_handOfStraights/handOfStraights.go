package handofstraights

import "container/heap"

type MinHeap []int

func (this MinHeap) Len() int {
	return len(this)
}

func (this MinHeap) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this MinHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *MinHeap) Push(item any) {
	*this = append(*this, item.(int))
}

func (this *MinHeap) Pop() any {
	old := *this
	item := old[len(old)-1]
	*this = old[:len(old)-1]
	return item
}

func isNStraightHand(hand []int, groupSize int) bool {
	if remain := len(hand) % groupSize; remain != 0 {
		return false
	}

	// convert to a hashMap >> O(n)
	countMap := make(map[int]int)
	for _, card := range hand {
		countMap[card]++
	}

	// convert the key of hashMap to minHeap >> O(n)
	var pq MinHeap
	for k, _ := range countMap {
		pq = append(pq, k)
	}
	heap.Init(&pq)

	// check if those cards can be arranged to new groups with consecutive cards. >> O(nlogn)
	for pq.Len() > 0 {
		firstCard := pq[0]
		for i := firstCard; i < firstCard+groupSize; i++ {
			count, ok := countMap[i]
			if !ok || count <= 0 {
				return false
			}
			count--
			countMap[i] = count

			if count == 0 && pq.Len() > 0 {
				if pq[0] == i {
					heap.Pop(&pq)
				} else {
					return false
				}
			}
		}
	}

	return true
}
