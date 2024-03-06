package longestincreasingsubsequence

func lengthOfLIS(nums []int) int {
	n := len(nums)
	type item struct {
		len int
		val int
	}
	dp := make([]item, n+1)
	dp[0] = item{len: 0, val: -1 << 31}
	for i := 0; i < n; i++ {
		var localMax int
		for j := i; j >= 0; j-- {
			if nums[i] > dp[j].val {
				localMax = max(localMax, dp[j].len)
			}
		}
		dp[i+1] = item{len: localMax + 1, val: nums[i]}
	}

	var ans int
	for _, item := range dp[1:] {
		if item.len > ans {
			ans = item.len
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
