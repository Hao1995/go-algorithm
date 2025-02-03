package productexceptself

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))

	prefix := 1
	for idx, num := range nums {
		output[idx] = prefix
		prefix *= num
	}

	suffix := 1
	for idx := len(nums) - 1; idx >= 0; idx-- {
		num := nums[idx]
		output[idx] *= suffix
		suffix *= num
	}

	return output
}
