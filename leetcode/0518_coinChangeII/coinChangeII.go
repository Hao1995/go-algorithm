package coinchangeii

func change(amount int, coins []int) int {
	cache := make([][]int, len(coins))
	for i := 0; i < len(coins); i++ {
		cache[i] = make([]int, amount+1)
		for j := 0; j < amount+1; j++ {
			cache[i][j] = -1
		}
	}

	var dfs func(idx int, a int) int
	dfs = func(idx int, a int) int {
		if a == 0 {
			return 1
		}

		if a < 0 {
			return 0
		}

		if idx == len(coins) {
			return 0
		}

		if cache[idx][a] != -1 {
			return cache[idx][a]
		}

		cache[idx][a] = dfs(idx, a-coins[idx]) + dfs(idx+1, a)
		return cache[idx][a]
	}

	ans := dfs(0, amount)
	return ans
}
