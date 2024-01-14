package wordsearch

func exist(board [][]byte, word string) bool {
	var m, n int = len(board), len(board[0])
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			// init visited
			var visited [][]bool = make([][]bool, m)
			for i := 0; i < m; i++ {
				visited[i] = make([]bool, n)
			}

			// dfs find word
			if dfs(board, visited, []int{r, c}, word) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, visited [][]bool, point []int, word string) bool {
	// check the word has been found
	if len(word) == 0 {
		return true
	}

	// check the boundary
	var r, c int = point[0], point[1]
	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
		return false
	}

	// check if the letter is expected
	expC, word := word[0], word[1:]
	if board[r][c] != expC {
		return false
	}

	// check the point has been visited
	if visited[r][c] {
		return false
	}
	visited[r][c] = true

	// continue to search the letter of remaining word
	var dirs [][]int = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, dir := range dirs {
		if dfs(board, visited, []int{r + dir[0], c + dir[1]}, word) {
			return true
		}
	}

	visited[r][c] = false
	return false
}
