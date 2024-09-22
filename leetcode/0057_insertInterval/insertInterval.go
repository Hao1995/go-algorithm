package insertinterval

func insert(intervals [][]int, newInterval []int) [][]int {
	isOverlapping := func(a []int, b []int) bool {
		// the begining of a within b
		if a[0] >= b[0] && a[0] <= b[1] {
			return true
		}

		// the begining of b within a
		if b[0] >= a[0] && b[0] <= a[1] {
			return true
		}

		return false
	}

	merge := func(a []int, b []int) []int {
		return []int{min(a[0], b[0]), max(a[1], b[1])}
	}

	ans := make([][]int, 0, len(intervals))

	// iterate intervals
	for i, interval := range intervals {
		if isOverlapping(interval, newInterval) {
			newInterval = merge(interval, newInterval)
		} else {
			if interval[1] < newInterval[0] {
				ans = append(ans, interval)
			} else {
				ans = append(ans, newInterval)
				ans = append(ans, intervals[i:]...)
				return ans
			}
		}
	}

	// after iteration, insert newInterval
	ans = append(ans, newInterval)
	return ans
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
