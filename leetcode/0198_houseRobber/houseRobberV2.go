package houserobber

func robV2(nums []int) int {
	dp := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			dp[i] = num
		} else if i == 1 {
			dp[i] = max(dp[i-1], num)
		} else {
			dp[i] = max(dp[i-1], dp[i-2]+num)
		}
	}
	return dp[len(dp)-1]
}
