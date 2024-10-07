package largestrectangleinhistogram

func largestRectangleArea(heights []int) int {
	var ans int

	// (startIdx, height)
	stack := make([][2]int, 0, len(heights))

	// Start ot calcualte
	for i, height := range heights {
		start := i
		var item [2]int
		for len(stack) > 0 && stack[len(stack)-1][1] > height {
			item, stack = stack[len(stack)-1], stack[:len(stack)-1]
			start = item[0]

			ans = max(ans, (i-item[0])*item[1])
		}

		stack = append(stack, [2]int{start, height})
	}

	// Calculate the remaining areas
	for _, item := range stack {
		ans = max(ans, (len(heights)-item[0])*item[1])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
