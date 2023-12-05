package coinchange

const (
	MAX_INT = 1 << 31
)

func coinChange(coins []int, amount int) int {
	var dp []int = make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = MAX_INT
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if rem := i - coin; rem >= 0 && dp[rem] != MAX_INT {
				dp[i] = min(dp[i], dp[rem]+1)
			}
		}
	}

	if dp[amount] == MAX_INT {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
