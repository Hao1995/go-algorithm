package findminimuminrotatedsortedarray

func findMin(nums []int) int {
	var length int = len(nums)
	var l, r int = 0, length - 1
	for l < r {
		var midd int = (l + r) / 2
		if nums[midd] > nums[length-1] {
			// min at right
			l = midd + 1
		} else {
			// min at left
			r = midd
		}
	}

	return nums[l]
}
