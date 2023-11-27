/*
2930. Number of Strings Which Can Be Rearranged to Contain Substring

You are given an integer n.

A string s is called good if it contains only lowercase English characters and it is possible to rearrange the characters of s such that the new string contains "leet" as a substring.

For example:

The string "lteer" is good because we can rearrange it to form "leetr" .
"letl" is not good because we cannot rearrange it to contain "leet" as a substring.
Return the total number of good strings of length n.

Since the answer may be large, return it modulo 109 + 7.

A substring is a contiguous sequence of characters within a string.

https://leetcode.com/problems/number-of-strings-which-can-be-rearranged-to-contain-substring/submissions/1103460112/
*/

package stringcount

const (
	MOD = 1e9 + 7
)

func stringCount(n int) int {
	//     n e l t
	var dp [][][][]int64 = make([][][][]int64, n+1)
	for i := range dp {
		// e
		dp[i] = make([][][]int64, 3)
		for j := range dp[i] {
			// l
			dp[i][j] = make([][]int64, 2)
			for k := range dp[i][j] {
				// t
				dp[i][j][k] = make([]int64, 2)
				for v := range dp[i][j][k] {
					dp[i][j][k][v] = -1
				}
			}
		}
	}

	return int(find(dp, 0, n, 0, 0, 0) % MOD)
}

func find(dp [][][][]int64, idx int, n int, e int, l int, t int) int64 {
	if idx == n {
		if e >= 2 && l >= 1 && t >= 1 {
			return 1
		} else {
			return 0
		}
	}

	if dp[idx][e][l][t] != -1 {
		return dp[idx][e][l][t]
	}

	var pick int64 = 0
	for i := 0; i < 26; i++ {
		if i == 4 && e < 2 {
			// char is e
			pick += find(dp, idx+1, n, e+1, l, t) % MOD
		} else if i == 11 && l < 1 {
			// char is l
			pick += find(dp, idx+1, n, e, l+1, t) % MOD
		} else if i == 19 && t < 1 {
			// char is t
			pick += find(dp, idx+1, n, e, l, t+1) % MOD
		} else {
			// char is others
			pick += find(dp, idx+1, n, e, l, t) % MOD
		}
	}

	dp[idx][e][l][t] = pick
	return dp[idx][e][l][t]
}
