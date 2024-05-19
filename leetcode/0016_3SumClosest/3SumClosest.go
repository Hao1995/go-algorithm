package sumclosest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var ans int = nums[0] + nums[1] + nums[2]
	var minDiff int = 1 << 31
	for i := 0; i < len(nums)-2; i++ {
		var left, right int = i + 1, len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if target == sum {
				return sum
			}
			if sum < target {
				left++
			} else {
				right--
			}

			diff := int(math.Abs(float64(sum - target)))
			if diff < minDiff {
				minDiff = diff
				ans = sum
			}
		}
	}

	return ans
}
