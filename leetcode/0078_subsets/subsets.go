package subsets

func subsets(nums []int) [][]int {
	var ans [][]int
	backtracking(&ans, []int{}, nums, 0)
	return ans
}

func backtracking(result *[][]int, tmpList []int, nums []int, start int) {
	var cpList []int = make([]int, len(tmpList))
	copy(cpList, tmpList)
	*result = append(*result, cpList)

	for i := start; i < len(nums); i++ {
		tmpList = append(tmpList, nums[i])
		backtracking(result, tmpList, nums, i+1)
		tmpList = tmpList[:len(tmpList)-1]
	}
}
