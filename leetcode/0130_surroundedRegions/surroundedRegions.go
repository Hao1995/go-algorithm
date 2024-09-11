package surroundedregions

func solve(board [][]byte) {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	// mark the regions on the edged to be visited
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if board[r][c] == 'O' &&
				(r == len(board)-1 || r == 0 || c == len(board[0])-1 || c == 0) {
				dfs(board, r, c, visited, true)
			}
		}
	}

	// mark the regions surrounded by 'X' be 'X'
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if board[r][c] == 'O' {
				dfs(board, r, c, visited, false)
			}
		}
	}
}

func dfs(board [][]byte, r, c int, visited [][]bool, isEdge bool) {
	if r >= len(board) || r < 0 || c >= len(board[0]) || c < 0 {
		return
	}

	if visited[r][c] {
		return
	}
	visited[r][c] = true

	if board[r][c] == 'O' {
		if !isEdge {
			board[r][c] = 'X'
		}

		// explore the adjacent 'O' cells
		dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		for _, dir := range dirs {
			dfs(board, r+dir[0], c+dir[1], visited, isEdge)
		}
	}
}
