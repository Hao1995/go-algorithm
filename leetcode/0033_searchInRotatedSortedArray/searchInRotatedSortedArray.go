package searchinrotatedsortedarray

func search(nums []int, target int) int {
	if nums[0] > nums[len(nums)-1] {
		var minOfLeft int = nums[0]

		pivot := getPivot(nums)
		pivotNum := nums[pivot]
		if target >= minOfLeft && target <= pivotNum {
			return bs(nums[0:pivot+1], target)
		} else {
			idx := bs(nums[pivot+1:], target)
			if idx == -1 {
				return -1
			} else {
				return pivot + 1 + idx
			}
		}
	} else {
		return bs(nums, target)
	}

	return -1
}

func getPivot(nums []int) int {
	var minOfLeft int = nums[0]
	var start, end int = 0, len(nums)
	var ans int
	for start < end {
		var midd int = (start + end) / 2
		num := nums[midd]
		if num < minOfLeft {
			end = midd
		} else {
			ans = midd
			start = midd + 1
		}
	}
	return ans
}

func bs(nums []int, target int) int {
	var start, end int = 0, len(nums)
	for start < end {
		var midd int = (start + end) / 2
		num := nums[midd]
		if num == target {
			return midd
		} else if num > target {
			end = midd
		} else if num < target {
			start = midd + 1
		}
	}
	return -1
}
