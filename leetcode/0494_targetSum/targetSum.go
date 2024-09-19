package targetsum

func findTargetSumWays(nums []int, target int) int {
	return dfs(nums, target, 0)
}

func dfs(nums []int, target int, idx int) int {
	if idx == len(nums) {
		if target == 0 {
			return 1
		} else {
			return 0
		}
	}

	return dfs(nums, target-nums[idx], idx+1) + dfs(nums, target+nums[idx], idx+1)
}
