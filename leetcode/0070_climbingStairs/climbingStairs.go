package climbingstairs

func climbStairs(n int) int {
	cache := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		cache[i] = -1
	}
	return dp(n, cache)
}

func dp(n int, cache []int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if cache[n] != -1 {
		return cache[n]
	}

	cache[n] = dp(n-1, cache) + dp(n-2, cache)

	return cache[n]
}
