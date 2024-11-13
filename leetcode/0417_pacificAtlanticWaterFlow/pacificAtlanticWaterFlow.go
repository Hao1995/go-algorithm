package pacificatlanticwaterflow

func pacificAtlantic(heights [][]int) [][]int {
	var rl, cl int = len(heights), len(heights[0])

	var dfs func(r, c, prevHeight int, visited [][]bool)
	dfs = func(r, c, prevHeight int, visited [][]bool) {
		if r < 0 || c < 0 || r == rl || c == cl || visited[r][c] || heights[r][c] < prevHeight {
			return
		}

		visited[r][c] = true

		dfs(r+1, c, heights[r][c], visited)
		dfs(r-1, c, heights[r][c], visited)
		dfs(r, c+1, heights[r][c], visited)
		dfs(r, c-1, heights[r][c], visited)
	}

	pac, atl := make([][]bool, rl), make([][]bool, rl)
	for r := 0; r < rl; r++ {
		pac[r] = make([]bool, cl)
		atl[r] = make([]bool, cl)
	}

	for c := 0; c < cl; c++ {
		dfs(0, c, heights[0][c], pac)
		dfs(rl-1, c, heights[rl-1][c], atl)
	}

	for r := 0; r < rl; r++ {
		dfs(r, 0, heights[r][0], pac)
		dfs(r, cl-1, heights[r][cl-1], atl)
	}

	var ans [][]int
	for r := 0; r < rl; r++ {
		for c := 0; c < cl; c++ {
			if pac[r][c] && atl[r][c] {
				ans = append(ans, []int{r, c})
			}
		}
	}

	return ans
}
