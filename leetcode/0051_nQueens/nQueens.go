package nqueens

func solveNQueens(n int) [][]string {
	col := make(map[int]bool)
	postDiag := make(map[int]bool) // r+c
	negDiag := make(map[int]bool)  // r-c

	board := make([][]byte, n)
	for r := 0; r < n; r++ {
		board[r] = make([]byte, n)
		for c := 0; c < n; c++ {
			board[r][c] = '.'
		}
	}

	var ans [][]string

	var backtracking func(r int)
	backtracking = func(r int) {
		if r == n {
			tmp := make([]string, n)
			for r := 0; r < n; r++ {
				tmp[r] = string(board[r])
			}
			ans = append(ans, tmp)
			return
		}

		for c := 0; c < n; c++ {
			if col[c] || postDiag[r+c] || negDiag[r-c] {
				continue
			}

			board[r][c] = 'Q'
			col[c] = true
			postDiag[r+c] = true
			negDiag[r-c] = true

			backtracking(r + 1)

			board[r][c] = '.'
			delete(col, c)
			delete(postDiag, r+c)
			delete(negDiag, r-c)
		}
	}

	backtracking(0)

	return ans
}
