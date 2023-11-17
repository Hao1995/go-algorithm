package countsubislands

var (
	isSubIsland = true
)

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	// init seen grid
	var seen [][]int = make([][]int, len(grid2))
	for i := range seen {
		seen[i] = make([]int, len(grid2[i]))
	}

	var count int
	for r := 0; r < len(grid2); r++ {
		for c := 0; c < len(grid2[0]); c++ {
			if seen[r][c] == 1 {
				continue
			}

			if grid2[r][c] == 1 {
				isSubIsland = true
				dfs(grid1, grid2, seen, r, c)
				if isSubIsland {
					count++
				}
			}
		}
	}
	return count
}

func dfs(grid1, grid2, seen [][]int, r, c int) {
	if r < 0 || r == len(grid2) || c < 0 || c == len(grid2[0]) {
		return
	}

	if seen[r][c] == 1 {
		return
	}
	seen[r][c] = 1

	if grid2[r][c] == 0 {
		return
	}

	// keep going, because we need to mark whole island of grid2
	// then next time we would not try the remaining cells
	if grid1[r][c] == 0 {
		isSubIsland = false
	}

	dfs(grid1, grid2, seen, r+1, c)
	dfs(grid1, grid2, seen, r-1, c)
	dfs(grid1, grid2, seen, r, c+1)
	dfs(grid1, grid2, seen, r, c-1)

	return
}
