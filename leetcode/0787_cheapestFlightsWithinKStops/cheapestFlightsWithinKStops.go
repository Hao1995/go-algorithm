package cheapestflightswithinkstops

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	var prices []int = make([]int, n)
	for i := range prices {
		prices[i] = 1 << 31
	}
	prices[src] = 0

	for i := 0; i < k+1; i++ {
		tmpPrices := make([]int, n)
		copy(tmpPrices, prices)

		for _, flight := range flights {
			s, d, p := flight[0], flight[1], flight[2]
			if prices[s] == 1<<31 {
				continue
			}
			if prices[s]+p < tmpPrices[d] {
				tmpPrices[d] = prices[s] + p
			}
		}

		prices = tmpPrices
	}

	if prices[dst] == 1<<31 {
		return -1
	}
	return prices[dst]
}
