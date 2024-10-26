package houseRobberII

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var dp func(nums []int) int
	dp = func(nums []int) int {
		var rob1, rob2 int = 0, 0
		for _, num := range nums {
			tmp := max(rob1+num, rob2)
			rob1, rob2 = rob2, tmp
		}
		return rob2
	}

	return max(dp(nums[1:]), dp(nums[:len(nums)-1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
