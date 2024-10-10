package permutations

func permute(nums []int) [][]int {
	// convert nums to a hash map >> O(n)
	numCount := make(map[int]int)
	for _, num := range nums {
		numCount[num]++
	}

	var ans [][]int

	// explore all the possible permutations >>  O(n!)
	var backtracking func(path []int)
	backtracking = func(path []int) {
		if len(path) == len(nums) {
			ans = append(ans, path)
			return
		}

		for num, count := range numCount {
			if count == 0 {
				continue
			}

			numCount[num]--
			backtracking(append(path, num))
			numCount[num]++
		}
	}

	backtracking([]int{})

	return ans
}
