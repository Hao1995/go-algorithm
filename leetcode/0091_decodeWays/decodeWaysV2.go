package decodeways

import "strconv"

func numDecodingsV2(s string) int {
	dp := make([]int, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		// check invalid char in advance
		if s[i] == '0' {
			continue
		}

		for _, length := range []int{1, 2} {
			// avoid out of size
			if i >= len(s)-length+1 {
				continue
			}

			num, _ := strconv.Atoi(s[i : i+length])
			if num >= 1 && num <= 26 {
				if i == len(s)-length {
					dp[i] += 1 // last n chars should default add 1
				} else {
					dp[i] += dp[i+length]
				}
			}
		}
	}

	return dp[0]
}
