package maxareaofisland

func maxAreaOfIsland(grid [][]int) int {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
			return 0
		}

		if grid[i][j] == 0 {
			return 0
		}

		grid[i][j] = 0
		var area int = 1 // current land

		// try different directions
		dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, dir := range dirs {
			area += dfs(i+dir[0], j+dir[1])
		}

		return area
	}

	var ans int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				ans = max(ans, dfs(i, j))
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
