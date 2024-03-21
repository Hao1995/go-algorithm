package main

func QuickSort(data []int) {
	if len(data) == 0 {
		return
	}
	quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, left, right int) {
	if left >= right {
		return
	}

	pivot := (left + right) / 2

	pivot = sort(data, left, pivot, right)

	quickSort(data, left, pivot-1)
	quickSort(data, pivot+1, right)
}

func sort(data []int, left, pivot, right int) int {
	data[pivot], data[right] = data[right], data[pivot]
	pivot = right

	right = right - 1

	for {
		for {
			if left == right {
				break
			}

			if data[left] > data[pivot] {
				break
			}

			left++
		}

		for {
			if left == right {
				break
			}

			if data[right] < data[pivot] {
				break
			}

			right--
		}

		if left == right {
			break
		}

		// case: left and right could not meet each other.
		//       l  r           p
		// [10, 53, 20, 80, 60, 50]
		// left 53 is greater than pivot 50
		// right 20 is less than pivot 50
		//
		// So we have to change left and pivot value
		//       l  r           p
		// [10, 50, 20, 80, 60, 53]
		// Then left value will less than pivot value, left++
		//          lr          p
		// [10, 50, 20, 80, 60, 53]
		data[left], data[pivot] = data[pivot], data[left]
	}

	if data[left] > data[pivot] {
		data[left], data[pivot] = data[pivot], data[left]
		pivot = left
	}

	return pivot
}
