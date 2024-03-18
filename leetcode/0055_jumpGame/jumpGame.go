package jumpgame

func canJump(nums []int) bool {
	return dfs(nums, 0, make([]bool, len(nums)))
}

func dfs(nums []int, currIdx int, visited []bool) bool {
	if currIdx >= len(nums) {
		return false
	}

	if currIdx == len(nums)-1 {
		return true
	}

	if visited[currIdx] {
		return false
	}

	for i := 1; i <= nums[currIdx]; i++ {
		if dfs(nums, currIdx+i, visited) {
			return true
		}
	}

	visited[currIdx] = true
	return false
}
