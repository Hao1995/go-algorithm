package maximumproductsubarray

func maxProductV2(nums []int) int {
	var ans int = nums[0]
	var maxProd, minProd int = nums[0], nums[0]
	for _, num := range nums[1:] {
		maxProd, minProd = max(max(maxProd*num, minProd*num), num), min(min(maxProd*num, minProd*num), num)
		ans = max(ans, maxProd)
	}
	return ans
}
