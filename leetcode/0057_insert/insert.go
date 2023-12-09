/*
57. Insert Interval
You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti. You are also given an interval newInterval = [start, end] that represents the start and end of another interval.
Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).
Return intervals after the insertion.
*/
package insert

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	var ans [][]int
	var tmpIntvl []int = newInterval
	var isMerged bool = false
	for _, intvl := range intervals {
		if isOverlapped(tmpIntvl, intvl) && !isMerged {
			tmpIntvl = mergeIntvls(tmpIntvl, intvl)
			continue
		}

		if !isMerged && isTmpAtFront(tmpIntvl, intvl) {
			ans = append(ans, tmpIntvl)
			isMerged = true
		}
		ans = append(ans, intvl)
	}

	if !isMerged {
		ans = append(ans, tmpIntvl)
	}

	return ans
}

func isOverlapped(a []int, b []int) bool {
	// [(]) || [()]
	if a[0] >= b[0] && a[0] <= b[1] {
		return true
	}
	// ([)]
	if a[1] >= b[0] && a[1] <= b[1] {
		return true
	}
	// ([])
	if a[0] <= b[0] && a[1] >= b[1] {
		return true
	}
	return false
}

func mergeIntvls(a []int, b []int) []int {
	start := a[0]
	if b[0] < start {
		start = b[0]
	}

	end := a[1]
	if b[1] > end {
		end = b[1]
	}

	return []int{start, end}
}

func isTmpAtFront(tmpIntvl []int, intvl []int) bool {
	return tmpIntvl[0] < intvl[0]
}
