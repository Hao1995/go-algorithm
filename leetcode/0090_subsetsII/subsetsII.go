package subsetsii

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var ans [][]int

	// O(nlogn)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// include or skip >> 2, n elements >> 2^n, O(2^n)
	// finally we're going to copy the list to ans >> O(n)
	// O(n*2^n)
	var dfs func(list []int, tmpNums []int)
	dfs = func(list []int, tmpNums []int) {
		if len(tmpNums) == 0 {
			tmp := make([]int, len(list))
			copy(tmp, list)
			ans = append(ans, tmp)
			return
		}

		num := tmpNums[0]

		// include
		dfs(append(list, num), tmpNums[1:])

		// skip
		var i int = 1
		for i < len(tmpNums) && tmpNums[i] == num {
			i++
		}
		dfs(list, tmpNums[i:])
	}

	dfs([]int{}, nums)

	return ans
}
