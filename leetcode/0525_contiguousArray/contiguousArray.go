package contiguousarray

func findMaxLength(nums []int) int {
	var zero, one int
	var ans int

	// diff: index
	diffMap := make(map[int]int)

	for i, num := range nums {
		if num == 0 {
			zero++
		} else {
			one++
		}

		diff := one - zero
		if _, ok := diffMap[diff]; !ok {
			diffMap[diff] = i
		}

		if one == zero {
			ans = one + zero
		} else {
			idx := diffMap[diff]
			ans = max(ans, i-idx)
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
