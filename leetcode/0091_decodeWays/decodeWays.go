package decodeways

import "strconv"

func numDecodings(s string) int {
	cache := make([]int, len(s))
	for i := range cache {
		cache[i] = -1
	}

	return dp(s, 0, cache)
}

func dp(s string, idx int, cache []int) int {
	if idx >= len(s) {
		return 1
	}

	if cache[idx] != -1 {
		return cache[idx]
	}

	cache[idx] = 0
	if isValid(int(s[idx]-'0'), 1) {
		cache[idx] = dp(s, idx+1, cache)
	}

	if idx < len(s)-1 {
		codeStr := s[idx : idx+2]
		code, _ := strconv.Atoi(codeStr)
		if isValid(code, 2) {
			cache[idx] += dp(s, idx+2, cache)
		}
	}
	return cache[idx]
}

func isValid(code int, length int) bool {
	if length == 1 {
		return code >= 1 && code <= 9
	}
	return code >= 10 && code <= 26
}
