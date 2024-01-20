package taskscheduler

import "sort"

func leastInterval(tasks []byte, n int) int {
	// O(m)
	countMap := make(map[byte]int)
	for _, c := range tasks {
		countMap[c] += 1
	}

	// O(m)
	pq := make([]int, len(countMap))
	var i int
	for _, count := range countMap {
		pq[i] = count
		i++
	}

	// O(mlonm)
	sort.Slice(pq, func(i, j int) bool { return pq[i] > pq[j] })

	var ans int
	for len(pq) > 0 {
		var tmpPQ []int
		for i := 0; i <= n; i++ {
			if len(pq) > 0 {
				var p int
				p, pq = pq[0], pq[1:]
				tmpPQ = append(tmpPQ, p-1)
			}
			ans++
			if len(pq) == 0 {
				break
			}
		}

		for _, v := range tmpPQ {
			if v > 0 {
				pq = append(pq, v)
			}
		}
		sort.Slice(pq, func(i, j int) bool { return pq[i] > pq[j] })

		// If n == len(unique(tasks)), there must be idle time. So we need to add it.
		// If len(pq) == 0, it means there is no idle time that needs to be added.
		// EX: tasks=["A","A","A","B","B","B"], n=2, `A -> B -> idle`
		if len(pq) > 0 {
			ans += n - len(tmpPQ) + 1
		}
	}

	return ans
}
