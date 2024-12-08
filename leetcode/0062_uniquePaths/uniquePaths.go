package uniquepaths

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for r := 0; r < m; r++ {
		dp[r] = make([]int, n)
		for c := 0; c < n; c++ {
			dp[r][c] = -1
		}
	}

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || c < 0 || r == m || c == n {
			return 0
		}

		if r == m-1 && c == n-1 {
			return 1
		}

		if dp[r][c] != -1 {
			return dp[r][c]
		}

		dp[r][c] = dfs(r+1, c) + dfs(r, c+1)
		return dp[r][c]
	}

	return dfs(0, 0)
}
