package combinationsumiv

func combinationSum4(nums []int, target int) int {
	cache := make([]int, target+1)
	for i := 0; i < len(cache); i++ {
		cache[i] = -1
	}
	cache[0] = 1

	return dfs(nums, target, cache)
}

func dfs(nums []int, target int, cache []int) int {
	if cache[target] != -1 {
		return cache[target]
	}

	var result int
	for _, num := range nums {
		if newTarget := target - num; target == 0 {
			result++
		} else if newTarget < 0 {
			continue
		} else {
			result += dfs(nums, newTarget, cache)
		}
	}

	cache[target] = result
	return cache[target]
}
