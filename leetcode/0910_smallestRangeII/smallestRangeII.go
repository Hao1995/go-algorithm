package smallestrangeii

import "sort"

func smallestRangeII(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var ans int = nums[len(nums)-1] - nums[0]
	for i := 0; i < len(nums)-1; i++ {
		maxVal := max(nums[i]+k, nums[len(nums)-1]-k)
		minVal := min(nums[0]+k, nums[i+1]-k)
		ans = min(ans, maxVal-minVal)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
