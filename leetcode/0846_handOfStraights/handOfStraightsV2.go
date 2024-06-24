package handofstraights

import "container/heap"

// n: lenght of hand
// m: number of unique cards
// k: group size
// Time complexity: O(n+mlogm)
func isNStraightHand(hand []int, groupSize int) bool {
	// convert to count map >> O(n)
	countMap := make(map[int]int)
	for _, card := range hand {
		countMap[card]++
	}

	// init min priority queue >> O(m)
	var pq MinHeap
	for k, _ := range countMap {
		pq = append(pq, k)
	}
	heap.Init(&pq)

	// keep retrieve the min card >> O(m*logm)
	var start []int
	var lastChecked, opened int = -1, 0
	for pq.Len() > 0 {
		firstCard := heap.Pop(&pq).(int)

		// There are still unfilled groups but the card number not consecutive.
		if opened > 0 && firstCard > lastChecked+1 {
			return false
		}

		// The number of the card can't fulfill all the group.
		count := countMap[firstCard]
		if opened > count {
			return false
		}

		start = append(start, count-opened)
		lastChecked = firstCard
		opened = count

		// Close opened groups
		if len(start) == groupSize {
			var firstGroupNum int
			firstGroupNum, start = start[0], start[1:]
			opened -= firstGroupNum
		}
	}

	return opened == 0
}
