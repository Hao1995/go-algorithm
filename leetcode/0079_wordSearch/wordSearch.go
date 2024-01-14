package wordsearch

var (
	dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

func exist(board [][]byte, word string) bool {
	var m, n int = len(board), len(board[0])
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] != word[0] {
				continue
			}

			if dfs(board, word, r, c, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, r, c, widx int) bool {
	// check the word has been found
	if widx == len(word) {
		return true
	}

	// check the boundary
	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
		return false
	}

	// check if the letter is expected
	expC := word[widx]
	if board[r][c] != expC {
		return false
	}

	// check the point has been visited
	if board[r][c] == '*' {
		return false
	}
	board[r][c] = '*'

	// continue to search the letter of remaining word
	widx++
	for _, dir := range dirs {
		if dfs(board, word, r+dir[0], c+dir[1], widx) {
			return true
		}
	}

	board[r][c] = expC
	return false
}
