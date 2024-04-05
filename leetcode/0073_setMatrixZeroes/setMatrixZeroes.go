package setmatrixzeroes

func setZeroes(matrix [][]int) {
	// mark start and end of zero element.
	zeroRow := make(map[int]bool)
	zeroCol := make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				if !zeroRow[i] {
					zeroRow[i] = true
				}
				if !zeroCol[j] {
					zeroCol[j] = true
				}
			}
		}
	}

	// mark zero row and col
	for r := range zeroRow {
		for c := 0; c < len(matrix[r]); c++ {
			matrix[r][c] = 0
		}
	}

	for c := range zeroCol {
		for r := 0; r < len(matrix); r++ {
			matrix[r][c] = 0
		}
	}
}
