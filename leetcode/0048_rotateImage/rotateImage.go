package rotateimage

func rotate(matrix [][]int) {
	leftLimit, rightLimit, upperLimit, bottomLimit := 0, len(matrix[0])-1, 0, len(matrix)-1
	var p1, p2, p3, p4 []int = []int{upperLimit, leftLimit}, []int{upperLimit, rightLimit}, []int{bottomLimit, rightLimit}, []int{bottomLimit, leftLimit}
	for {
		matrix[p2[0]][p2[1]], matrix[p3[0]][p3[1]], matrix[p4[0]][p4[1]], matrix[p1[0]][p1[1]] = matrix[p1[0]][p1[1]], matrix[p2[0]][p2[1]], matrix[p3[0]][p3[1]], matrix[p4[0]][p4[1]]

		p1[1]++
		p2[0]++
		p3[1]--
		p4[0]--
		if p1[1] == rightLimit {
			leftLimit++
			rightLimit--
			upperLimit++
			bottomLimit--
			if leftLimit > rightLimit || upperLimit > bottomLimit {
				// check if finished?
				break
			}
			p1, p2, p3, p4 = []int{upperLimit, leftLimit}, []int{upperLimit, rightLimit}, []int{bottomLimit, rightLimit}, []int{bottomLimit, leftLimit}
		} else if p1[1] > rightLimit {
			// check if finished?
			break
		}
	}
}
