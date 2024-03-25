package maximalsquare

func maximalSquare(matrix [][]byte) int {
	// init dp
	var dp [][]int = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}

	var ans int

	// init column
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			ans = max(ans, dp[0][i])
		}
	}

	// init row
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			ans = max(ans, dp[i][0])
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
				ans = max(ans, dp[i][j]*dp[i][j])
			}
		}
	}

	return ans
}
