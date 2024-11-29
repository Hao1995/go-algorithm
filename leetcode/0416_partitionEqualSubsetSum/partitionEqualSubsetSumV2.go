package partitionequalsubsetsum

func canPartitionV2(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}

	var target int = sum / 2

	dp := make(map[int]bool)
	for _, num := range nums {
		tmpDP := make(map[int]bool)
		for key, _ := range dp {
			tmpDP[key+num] = true
			tmpDP[key] = true
		}
		tmpDP[num] = true
		dp = tmpDP
	}

	return dp[target]
}
