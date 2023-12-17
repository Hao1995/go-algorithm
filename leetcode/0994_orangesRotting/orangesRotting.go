package orangesrotting

func orangesRotting(grid [][]int) int {
	var m int = len(grid)
	var n int = len(grid[0])

	// init seen grid
	var seen [][]bool = make([][]bool, m)
	for i := 0; i < len(seen); i++ {
		seen[i] = make([]bool, n)
	}

	var queue, nextQueue [][]int // [[0,0],[0,1]...]

	// Get all rotten orange to queue and mark it has seen
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, []int{r, c})
				seen[r][c] = true
			}
		}
	}

	var count int  // count the minimum number
	var curr []int // current orange position
	// BFS marks the nearby orange to be rotten
	for len(queue) > 0 || len(nextQueue) > 0 {
		if len(queue) == 0 {
			queue = nextQueue
			nextQueue = [][]int{}
			count++
		}
		curr, queue = queue[0], queue[1:]
		r, c := curr[0], curr[1]

		grid[r][c] = 2

		var dirs [][]int = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, dir := range dirs {
			var nr, nc int = r + dir[0], c + dir[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n && !seen[nr][nc] && grid[nr][nc] == 1 {
				seen[nr][nc] = true
				nextQueue = append(nextQueue, []int{nr, nc})
				continue
			}
		}
	}

	// Check is there still exist any fresh orange?
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				return -1
			}
		}
	}

	return count
}
