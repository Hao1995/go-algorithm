package rotateimage

func rotate(matrix [][]int) {
	var SIZE int = len(matrix)

	var layer int = 0
	for layer < SIZE-1-layer {
		var boundary int = SIZE - 1 - layer
		for idx := 0; layer+idx < boundary; idx++ {
			matrix[idx+layer][boundary], matrix[boundary][boundary-idx], matrix[boundary-idx][layer], matrix[layer][layer+idx] =
				matrix[layer][layer+idx], matrix[layer+idx][boundary], matrix[boundary][boundary-idx], matrix[boundary-idx][layer]
		}
		layer++
	}
}
