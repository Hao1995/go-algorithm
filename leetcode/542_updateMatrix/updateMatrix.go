/*
542. 01 Matrix
Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
The distance between two adjacent cells is 1.
*/
package updatematrix

const (
	MAX_INT int = 1<<31 - 1
)

func updateMatrix(mat [][]int) [][]int {
	var rl int = len(mat)
	var cl int = len(mat[0])

	var queue [][]int = make([][]int, 0)
	for r := 0; r < rl; r++ {
		for c := 0; c < cl; c++ {
			if mat[r][c] == 0 {
				queue = append(queue, []int{r, c})
			} else {
				mat[r][c] = MAX_INT
			}
		}
	}

	var cell []int
	for len(queue) > 0 {
		cell, queue = queue[0], queue[1:]
		r, c := cell[0], cell[1]

		dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, dir := range dirs {
			nr, nc := r+dir[0], c+dir[1]
			if nr >= 0 && nr < rl && nc >= 0 && nc < cl && mat[nr][nc] > mat[r][c]+1 {
				queue = append(queue, []int{nr, nc})
				mat[nr][nc] = mat[r][c] + 1
			}
		}
	}
	return mat
}
