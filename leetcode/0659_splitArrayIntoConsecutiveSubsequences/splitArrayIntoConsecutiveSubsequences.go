package splitarrayintoconsecutivesubsequences

func isPossible(nums []int) bool {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	endMap := make(map[int]int)
	for _, num := range nums {
		if count := countMap[num]; count > 0 {
			countMap[num]--

			// continue append to exist subsequence
			if eTimes := endMap[num-1]; eTimes > 0 {
				endMap[num-1]--
				endMap[num]++
				continue
			}

			// start a new subsequence
			count1 := countMap[num+1]
			count2 := countMap[num+2]
			if count1 > 0 && count2 > 0 {
				countMap[num+1]--
				countMap[num+2]--
				endMap[num+2]++
				continue
			}

			// illegal
			return false
		}
	}

	return true
}
