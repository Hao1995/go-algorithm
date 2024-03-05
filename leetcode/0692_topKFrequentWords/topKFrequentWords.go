package topkfrequentwords

import "sort"

func topKFrequent(words []string, k int) []string {
	// count by hash map O(n)
	countMap := make(map[string]int)
	for _, word := range words {
		countMap[word]++
	}

	// Put to priority queue O(n)+O(nlogn),O(n)
	// (order by: freq desc, lexicographical asc)
	type strFreq struct {
		str  string
		freq int
	}
	var pq []strFreq
	for k, v := range countMap {
		newItem := strFreq{
			str:  k,
			freq: v,
		}
		pq = append(pq, newItem)
	}
	sort.Slice(pq, func(i, j int) bool {
		if pq[i].freq == pq[j].freq {
			return pq[i].str < pq[j].str
		}
		return pq[i].freq > pq[j].freq
	})

	// Retrieve top k data from queue, O(k),O(k)
	var ans []string = make([]string, 0, k)
	for _, item := range pq[:k] {
		ans = append(ans, item.str)
	}

	return ans
}
