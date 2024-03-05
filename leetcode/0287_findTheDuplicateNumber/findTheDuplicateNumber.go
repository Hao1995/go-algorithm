package findtheduplicatenumber

func findDuplicate(nums []int) int {
	var low, high int = 1, len(nums)
	for low < high {
		var midd int = (low + high) / 2

		var count int
		for _, num := range nums {
			if midd >= num {
				count++
			}
		}

		if count > midd {
			high = midd
		} else {
			low = midd + 1
		}
	}
	return low
}
