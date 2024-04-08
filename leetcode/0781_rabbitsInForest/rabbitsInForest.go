package rabbitsinforest

func numRabbits(answers []int) int {
	sameCount := make(map[int]int)
	var result int
	for _, ans := range answers {
		if count, ok := sameCount[ans]; ok && count > 0 {
			sameCount[ans]--
			continue
		}

		result += ans + 1
		sameCount[ans] = ans
	}
	return result
}
