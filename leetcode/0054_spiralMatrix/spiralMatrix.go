package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	var m, n int = len(matrix), len(matrix[0])
	var ans []int

	var top, right, bottom, left int = -1, n, m, -1
	var r, c int = top + 1, left + 1
	for {
		// go to right until the boundary
		for c < right {
			ans = append(ans, matrix[r][c])
			c++
		}
		c--
		top++
		r++
		if !isWithinBoundary(r, c, top, bottom, right, left) {
			break
		}

		// go to bottom until the boundary
		for r < bottom {
			ans = append(ans, matrix[r][c])
			r++
		}
		r--
		right--
		c--
		if !isWithinBoundary(r, c, top, bottom, right, left) {
			break
		}

		// go to left until the boundary
		for c > left {
			ans = append(ans, matrix[r][c])
			c--
		}
		c++
		bottom--
		r--
		if !isWithinBoundary(r, c, top, bottom, right, left) {
			break
		}

		// go to top until the boundary
		for r > top {
			ans = append(ans, matrix[r][c])
			r--
		}
		r++
		left++
		c++
		if !isWithinBoundary(r, c, top, bottom, right, left) {
			break
		}
	}

	return ans
}

func isWithinBoundary(r, c int, top, bottom, right, left int) bool {
	return r > top && r < bottom && c > left && c < right
}
