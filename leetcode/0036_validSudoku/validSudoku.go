package validsudoku

func isValidSudoku(board [][]byte) bool {
	rowBoard := initBoard(9, 9)
	colBoard := initBoard(9, 9)
	subBoard := initBoard(9, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c == '.' {
				continue
			}

			num := int(board[i][j]-'0') - 1
			subIdx := j/3 + int(i/3)*3
			if rowBoard[i][num] || colBoard[j][num] || subBoard[subIdx][num] {
				return false
			}
			rowBoard[i][num] = true
			colBoard[j][num] = true
			subBoard[subIdx][num] = true
		}
	}

	return true
}

func initBoard(m, n int) [][]bool {
	board := make([][]bool, m)
	for i := 0; i < m; i++ {
		board[i] = make([]bool, n)
	}
	return board
}
