package burstballoons

func maxCoins(nums []int) int {
	nums = append([]int{1}, append(nums, 1)...)
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums))
	}

	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if l > r {
			return 0
		}

		if dp[l][r] != 0 {
			return dp[l][r]
		}

		for i := l; i < r+1; i++ {
			coins := nums[l-1] * nums[i] * nums[r+1]
			coins += dfs(l, i-1) + dfs(i+1, r)
			dp[l][r] = max(dp[l][r], coins)
		}

		return dp[l][r]
	}

	return dfs(1, len(nums)-2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
