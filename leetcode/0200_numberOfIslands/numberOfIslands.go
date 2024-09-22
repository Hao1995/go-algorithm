package numberofislands

func numIslands(grid [][]byte) int {
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r == len(grid) || c < 0 || c == len(grid[0]) {
			return
		}

		if grid[r][c] == '1' {
			grid[r][c] = '2'

			dires := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
			for _, dire := range dires {
				dfs(r+dire[0], c+dire[1])
			}
		}
	}

	var ans int
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '1' {
				// mark all adjacent lands to be '2'
				dfs(r, c)
				// count 1
				ans++
			}
		}
	}

	return ans
}
