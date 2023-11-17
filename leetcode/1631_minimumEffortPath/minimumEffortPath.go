package minimumeffortpath

import "math"

func minimumEffortPath(heights [][]int) int {
	var start int
	var end int = 1 << 32
	var ans int
	for start < end {
		var midd int = (start + end) / 2
		if canReachDist(heights, midd) {
			end = midd
			ans = midd
		} else {
			start = midd + 1
		}
	}
	return ans
}

func canReachDist(heights [][]int, threshold int) bool {
	y := len(heights)
	x := len(heights[0])

	visitedGrid := make([][]bool, y)
	for i := 0; i < y; i++ {
		visitedGrid[i] = make([]bool, x)
	}

	return dfs(heights, 0, 0, visitedGrid, threshold)
}

func dfs(heights [][]int, r, c int, visitedGrid [][]bool, threshold int) bool {
	y := len(heights)
	x := len(heights[0])

	// bottom-right point, reach the distination
	if r == y-1 && c == x-1 {
		return true
	}

	visitedGrid[r][c] = true

	directions := []int{0, 1, 0, -1, 0}
	for i := 0; i < 4; i++ {
		nr, nc := r+directions[i], c+directions[i+1]
		// skip if exceed the boundary or already visited
		if nr < 0 || nr == y || nc < 0 || nc == x || visitedGrid[nr][nc] {
			continue
		}

		if abs(heights[nr][nc]-heights[r][c]) <= threshold && dfs(heights, nr, nc, visitedGrid, threshold) {
			return true
		}
	}

	return false
}

func abs(val int) int {
	return int(math.Abs(float64(val)))
}
