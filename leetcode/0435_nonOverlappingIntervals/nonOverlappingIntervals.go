package nonoverlappingintervals

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	var res, prevEnd int = 0, intervals[0][1]
	for _, interval := range intervals[1:] {
		if interval[0] < prevEnd {
			res++
			prevEnd = min(prevEnd, interval[1])
		} else {
			prevEnd = interval[1]
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
