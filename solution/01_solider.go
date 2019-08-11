package solution

func Solider(ranks []int) int {

	ranksMap := make(map[int]int)
	for _, rank := range ranks {
		ranksMap[rank]++
	}

	total := 0
	for rank, time := range ranksMap {
		if _, ok := ranksMap[rank+1]; ok {
			total += time
		}
	}

	return total
}
