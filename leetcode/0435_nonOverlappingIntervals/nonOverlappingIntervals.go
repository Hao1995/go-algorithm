package nonoverlappingintervals

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	// O(nlogn)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	// O(n)
	var end int = intervals[0][1]
	var ans int
	for _, intv := range intervals[1:] {
		if intv[0] >= end && intv[1] >= end {
			end = intv[1]
			continue
		}
		ans++
	}

	return ans
}
