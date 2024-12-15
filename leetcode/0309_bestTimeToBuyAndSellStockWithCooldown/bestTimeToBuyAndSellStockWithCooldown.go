package besttimetobuyandsellstockwithcooldown

func maxProfit(prices []int) int {
	dp := make([][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = [2]int{-1, -1}
	}

	var dfs func(idx int, operator int) int
	dfs = func(idx int, operator int) int {
		if idx >= len(prices) {
			return 0
		}

		if val := dp[idx][operator]; val != -1 {
			return val
		}

		var res int
		if operator == 1 {
			buy := dfs(idx+1, 0) - prices[idx]
			skip := dfs(idx+1, 1)
			res = max(buy, skip)
		} else {
			sell := dfs(idx+2, 1) + prices[idx]
			hold := dfs(idx+1, 0)
			res = max(sell, hold)
		}

		dp[idx][operator] = res
		return res
	}

	return dfs(0, 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
