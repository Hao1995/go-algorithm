package wordbreak

func wordBreakV2(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true

	for i := len(s) - 1; i >= 0; i-- {
		for _, word := range wordDict {
			if len(word) <= len(s[i:]) && word == s[i:i+len(word)] && dp[i+len(word)] {
				dp[i] = true
			}

			if dp[i] {
				break
			}
		}
	}

	return dp[0]
}
