package missingnumber

func missingNumber(nums []int) int {
	checkArr := make([]bool, len(nums)+1)

	for _, num := range nums {
		checkArr[num] = true
	}

	for i, exist := range checkArr {
		if !exist {
			return i
		}
	}

	return -1
}
