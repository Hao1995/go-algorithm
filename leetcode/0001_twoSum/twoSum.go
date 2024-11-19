package twosum

func twoSum(nums []int, target int) []int {
	// remaining number : index
	remainingMap := make(map[int]int)
	for idx, num := range nums {
		remaining := target - num
		remainingIdx, ok := remainingMap[remaining]
		if ok {
			return []int{idx, remainingIdx}
		} else {
			remainingMap[num] = idx
		}
	}
	return []int{}
}
