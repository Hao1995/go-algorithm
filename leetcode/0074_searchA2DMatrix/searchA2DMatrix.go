package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	var m, n int = len(matrix), len(matrix[0])
	var start, end int = 0, m * n
	for start < end {
		var midd int = (start + end) / 2
		var r, c int = midd / n, midd % n
		num := matrix[r][c]
		if num == target {
			return true
		}
		if num >= target {
			end = midd
		} else {
			start = midd + 1
		}
	}
	return false
}
