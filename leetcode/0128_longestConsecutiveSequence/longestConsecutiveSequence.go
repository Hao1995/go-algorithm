package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}

	var ans int
	for key := range numMap {
		// start since the min val in that group
		_, ok := numMap[key-1]
		if !ok {
			var count int = 1
			var idx int = key + 1
			for {
				if _, ok := numMap[idx]; !ok {
					break
				}
				count++
				idx++
			}
			ans = max(ans, count)
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
