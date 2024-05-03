package subarraysumequalsk

func subarraySum(nums []int, k int) int {
	// sum:count
	prefixSumMap := make(map[int]int)
	prefixSumMap[0]++

	var sum, ans int
	for _, num := range nums {
		sum += num
		diff := sum - k
		count, ok := prefixSumMap[diff]
		if ok {
			ans += count
		}
		prefixSumMap[sum]++
	}

	return ans
}
