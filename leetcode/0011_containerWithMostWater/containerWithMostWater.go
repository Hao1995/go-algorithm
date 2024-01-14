package containerwithmostwater

func maxArea(height []int) int {
	var start, end int = 0, len(height) - 1
	var ans int
	for start < end {
		var w int = end - start
		ans = max(ans, min(height[start], height[end])*w)
		if height[start] < height[end] {
			// if startH lt endH, that mean the max height is restricted by startH
			// so we gonna to inc start, in order to get higher max height
			// then we may get a higher amount of water
			start++
		} else {
			// vice versa
			end--
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
