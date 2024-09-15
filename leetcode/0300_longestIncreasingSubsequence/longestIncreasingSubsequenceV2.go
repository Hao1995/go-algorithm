package longestincreasingsubsequence

func lengthOfLISV2(nums []int) int {
	// init
	tmp := []int{nums[0]}

	for _, num := range nums[1:] {
		// BS, num greater than tmp[j] and less than tmp[j+1]
		var start, end int = 0, len(tmp)
		for start < end {
			var midd int = (start + end) / 2
			if num > tmp[midd] {
				start = midd + 1
			} else {
				end = midd
			}
		}

		if start == len(tmp) {
			// append
			tmp = append(tmp, num)
		} else {
			// change
			tmp[start] = num
		}
	}

	return len(tmp)
}
