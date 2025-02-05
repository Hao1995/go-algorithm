package binarysearch

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		var midd int = (l + r) / 2

		if nums[midd] == target {
			return midd
		}

		if nums[midd] > target {
			r = midd - 1
		} else {
			l = midd + 1
		}
	}
	return -1
}
