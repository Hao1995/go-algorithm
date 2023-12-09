package productexceptself

func productExceptSelf(nums []int) []int {
	var ans []int = make([]int, len(nums))
	for i := 0; i < len(ans); i++ {
		ans[i] = 1
	}

	var prefix int = 1
	for i := 0; i < len(ans); i++ {
		ans[i] *= prefix
		prefix *= nums[i]
	}

	var postfix int = 1
	for i := len(ans) - 1; i >= 0; i-- {
		ans[i] *= postfix
		postfix *= nums[i]
	}

	return ans
}
