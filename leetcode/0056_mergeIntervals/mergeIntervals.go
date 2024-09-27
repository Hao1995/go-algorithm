package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	// O(nlogn)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0, len(intervals))
	ans = append(ans, intervals[0])

	// O(n)
	for _, interval := range intervals[1:] {
		if isOverlapping(ans[len(ans)-1], interval) {
			ans[len(ans)-1] = mergeTwoIntervals(ans[len(ans)-1], interval)
		} else {
			ans = append(ans, interval)
		}
	}

	return ans
}

func isOverlapping(a []int, b []int) bool {
	// both start and end less than the other interval
	// () []
	if a[0] < b[0] && a[1] < b[0] {
		return false
	}

	// [] ()
	if b[0] < a[0] && b[1] < a[0] {
		return false
	}

	return true
}

func mergeTwoIntervals(a []int, b []int) []int {
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
