package partitionequalsubsetsum

func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	n := len(nums) + 1 // should insert 0 to the beginning
	m := sum/2 + 1
	var dp [][]bool = make([][]bool, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, m)
	}

	dp[0][0] = true
	nums = append([]int{0}, nums...)
	for i := 1; i < n; i++ {
		num := nums[i]
		for s := 0; s < m; s++ {
			dp[i][s] = dp[i-1][s] || s >= num && dp[i-1][s-num]
		}
	}

	return dp[n-1][sum/2]
}
