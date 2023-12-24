package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	var ans [][]int

	// sorted by start with asc
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		a := ans[len(ans)-1]
		b := intervals[i]
		if isOverlapped(a, b) {
			tmp := mergeIntervals(a, b)
			ans[len(ans)-1] = tmp
		} else {
			ans = append(ans, b)
		}
	}

	return ans
}

// [1,3],[2,4]
// [1,5],[2,4]
// [1,6],[1,3]
// [1,6],[8,10]
func isOverlapped(a []int, b []int) bool {
	if b[0] >= a[0] && b[0] <= a[1] {
		return true
	}
	return false
}

func mergeIntervals(a []int, b []int) []int {
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
