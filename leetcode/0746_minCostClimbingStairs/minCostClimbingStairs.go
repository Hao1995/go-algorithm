package mincostclimbingstairs

func minCostClimbingStairs(cost []int) int {
	var dp []int = make([]int, len(cost))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1 << 31
	}

	for i := 0; i < len(cost); i++ {
		if i < 2 {
			dp[i] = cost[i]
			continue
		}

		dp[i] = min(dp[i-1]+cost[i], dp[i-2]+cost[i])
	}

	return min(dp[len(dp)-1], dp[len(dp)-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
