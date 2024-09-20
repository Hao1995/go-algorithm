package distinctsubsequences

func numDistinct(s string, t string) int {
	// init dp
	dp := make([][]int, len(t))
	for ti := 0; ti < len(t); ti++ {
		dp[ti] = make([]int, len(s))
		for si := 0; si < len(s); si++ {
			dp[ti][si] = -1
		}
	}

	// search resulte
	return dfs(t, 0, s, 0, dp)
}

func dfs(t string, ti int, s string, si int, dp [][]int) int {
	if ti == len(t) {
		return 1
	}

	if si == len(s) {
		return 0
	}

	if dp[ti][si] != -1 {
		return dp[ti][si]
	}

	if t[ti] == s[si] {
		dp[ti][si] = dfs(t, ti+1, s, si+1, dp) + dfs(t, ti, s, si+1, dp)
		return dp[ti][si]
	}

	dp[ti][si] = dfs(t, ti, s, si+1, dp)
	return dp[ti][si]
}
