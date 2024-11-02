package longestincreasingpathinamatrix

func longestIncreasingPath(matrix [][]int) int {
	var m, n int = len(matrix), len(matrix[0])

	// init dp
	dp := make([][]int, m)
	for r := 0; r < m; r++ {
		dp[r] = make([]int, n)
	}

	var ans int

	// dfs
	var dfs func(r, c, prev int) int
	dfs = func(r, c, prev int) int {
		if r == m || c == n || r < 0 || c < 0 || matrix[r][c] <= prev {
			return 0
		}

		if dp[r][c] != 0 {
			return dp[r][c]
		}

		res := 1 + dfs(r+1, c, matrix[r][c])
		res = max(res, 1+dfs(r-1, c, matrix[r][c]))
		res = max(res, 1+dfs(r, c+1, matrix[r][c]))
		res = max(res, 1+dfs(r, c-1, matrix[r][c]))
		dp[r][c] = res
		ans = max(ans, dp[r][c])
		return dp[r][c]
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			dfs(r, c, -1)
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
