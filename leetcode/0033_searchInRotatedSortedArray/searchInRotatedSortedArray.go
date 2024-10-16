package searchinrotatedsortedarray

func search(nums []int, target int) int {
	// find start
	var start int
	// find start idx in a rotated array
	if nums[0] > nums[len(nums)-1] {
		var left, right int = 0, len(nums)
		for left < right {
			var midd int = (left + right) / 2
			if nums[midd] >= nums[0] {
				left = midd + 1
			} else {
				right = midd
			}
		}
		start = left
	}

	if target >= nums[start] && target <= nums[len(nums)-1] {
		// find right side
		return bs(nums, target, start, len(nums))
	} else {
		// find left side
		return bs(nums, target, 0, start)
	}
}

func bs(nums []int, target int, left, right int) int {
	for left < right {
		var midd int = (left + right) / 2
		if nums[midd] == target {
			return midd
		}
		if nums[midd] > target {
			right = midd
		} else {
			left = midd + 1
		}
	}
	return -1
}
