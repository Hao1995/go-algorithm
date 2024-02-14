package nextpermutation

import "sort"

func nextPermutation(nums []int) {
	n := len(nums)

	// Find peakIdx
	var peakIdx int = n - 1
	for ; peakIdx > 0; peakIdx-- {
		if nums[peakIdx-1] < nums[peakIdx] {
			break
		}
	}

	// Swap the largeNum and peakNum
	var largeIdx int = n - 1
	if peakIdx != 0 {
		for ; largeIdx >= peakIdx; largeIdx-- {
			if nums[peakIdx-1] < nums[largeIdx] {
				nums[peakIdx-1], nums[largeIdx] = nums[largeIdx], nums[peakIdx-1]
				break
			}
		}
	}

	// Sort nums gte peakIdx
	sort.Slice(nums[peakIdx:], func(i, j int) bool {
		return nums[peakIdx+i] < nums[peakIdx+j]
	})
}
