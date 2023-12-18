package permutations

func permute(nums []int) [][]int {
	var ans [][]int
	dfs(&ans, nums, []int{}, []int{})
	return ans
}

func dfs(ans *[][]int, nums []int, list []int, seenNums []int) {
	if len(list) == len(nums) {
		cpList := make([]int, len(list))
		copy(cpList, list)
		*ans = append(*ans, cpList)
		return
	}

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if numIsExist(num, seenNums) {
			continue
		}
		dfs(ans, nums, append(list, num), append(seenNums, num))
	}
}

func numIsExist(num int, seenNums []int) bool {
	for _, seenNum := range seenNums {
		if num == seenNum {
			return true
		}
	}
	return false
}
