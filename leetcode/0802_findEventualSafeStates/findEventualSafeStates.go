package findeventualsafestates

func eventualSafeNodes(graph [][]int) []int {
	safeMap := make(map[int]bool)

	// O(E+V)
	var dfs func(i int) bool
	dfs = func(i int) bool {
		if safe, ok := safeMap[i]; ok {
			return safe
		}
		safeMap[i] = false

		for _, neig := range graph[i] {
			if !dfs(neig) {
				return safeMap[i]
			}
		}

		safeMap[i] = true
		return safeMap[i]
	}

	// O(v)
	var ans []int
	for i := 0; i < len(graph); i++ {
		if dfs(i) {
			ans = append(ans, i)
		}
	}

	return ans
}
