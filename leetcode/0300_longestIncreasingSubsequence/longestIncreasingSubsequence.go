package longestincreasingsubsequence

func lengthOfLIS(nums []int) int {
	// t:O(n), s:O(n)
	cache := make([]int, len(nums))
	for i := 0; i < len(cache); i++ {
		cache[i] = -1
	}

	// t:O(n^2)
	for i := len(nums) - 1; i >= 0; i-- {
		dp(nums, i, cache)
	}

	// t:O(n), s:O(1)
	var ans int
	for _, tmp := range cache {
		ans = max(ans, tmp)
	}

	return ans
}

func dp(nums []int, idx int, cache []int) int {
	if idx >= len(nums) {
		return 0
	}

	var maxLen int
	for i := idx + 1; i < len(nums); i++ {
		// check it's increasingly
		if nums[i] > nums[idx] {
			// check if the following sequence has calcualted
			if cache[i] != -1 {
				maxLen = max(maxLen, cache[i])
			} else {
				maxLen = max(maxLen, dp(nums, i, cache))
			}
		}
	}

	// 1 means the current length
	cache[idx] = maxLen + 1

	return cache[idx]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
