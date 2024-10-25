package threesum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var ans [][]int
	for ai, anum := range nums {
		if ai > 0 && nums[ai-1] == anum {
			continue
		}

		var l, r int = ai + 1, len(nums) - 1
		for l < r {
			if l > ai+1 && nums[l] == nums[l-1] {
				// same value, skip
				l++
				continue
			}

			tmp := anum + nums[l] + nums[r]
			if tmp == 0 {
				ans = append(ans, []int{anum, nums[l], nums[r]})
				l++
				r--
			} else if tmp < 0 {
				// move right
				l++
			} else {
				// move left
				r--
			}
		}
	}

	return ans
}
