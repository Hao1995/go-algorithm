package maxsubarray

func maxSubArray(nums []int) int {
	var ans, sum int = -1 << 32, 0
	for _, num := range nums {
		sum = max(sum+num, num)
		ans = max(ans, sum)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
