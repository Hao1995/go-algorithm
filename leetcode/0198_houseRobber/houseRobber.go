package houserobber

func rob(nums []int) int {
	var a, b int

	for i, num := range nums {
		if i%2 == 0 {
			a = max(a+num, b)
		} else {
			b = max(b+num, a)
		}
	}

	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
