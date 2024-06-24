package handofstraights

import "container/heap"

// n: lenght of hand
// m: number of unique cards
// k: group size
// Time complexity: O(n+mk+mlogm)
func isNStraightHandV1(hand []int, groupSize int) bool {
	if remain := len(hand) % groupSize; remain != 0 {
		return false
	}

	// convert to a hashMap >> O(n)
	countMap := make(map[int]int)
	for _, card := range hand {
		countMap[card]++
	}

	// convert the key of hashMap to minHeap >> O(m)
	var pq MinHeap
	for k, _ := range countMap {
		pq = append(pq, k)
	}
	heap.Init(&pq)

	// check if those cards can be arranged to new groups with consecutive cards. >> O(mk+mlogm)
	for pq.Len() > 0 {
		firstCard := pq[0]

		// We expected all the card can be calculated, so it should be O(mk)
		for i := firstCard; i < firstCard+groupSize; i++ {
			count, ok := countMap[i]
			if !ok || count <= 0 {
				return false
			}
			count--
			countMap[i] = count

			if count == 0 && pq.Len() > 0 {
				if pq[0] == i {
					// Eventually this line will only be executed O(mlogm)
					heap.Pop(&pq)
				} else {
					return false
				}
			}
		}
	}

	return true
}
