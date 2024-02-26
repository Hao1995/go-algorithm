package maximumproductsubarray

func maxProduct(nums []int) int {
	var ans int = nums[0]
	var minVal, maxVal int = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxVal, minVal = minVal, maxVal
		}
		maxVal = max(nums[i], maxVal*nums[i])
		minVal = min(nums[i], minVal*nums[i])

		ans = max(ans, maxVal)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
