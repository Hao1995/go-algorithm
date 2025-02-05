package cheapestflightswithinkstops

import "math"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// O(n+1)
	prices := make([]int, n+1)
	for i, _ := range prices {
		prices[i] = math.MaxInt
	}
	prices[src] = 0

	// O(k+1)
	for i := 0; i < k+1; i++ {
		tmpPrices := make([]int, len(prices))
		copy(tmpPrices, prices)

		// O(E): all edges
		for _, flight := range flights {
			s, d, p := flight[0], flight[1], flight[2]
			if prices[s] == math.MaxInt {
				continue
			}

			tmpPrices[d] = min(tmpPrices[d], prices[s]+p)
		}

		prices = tmpPrices
	}

	if prices[dst] == math.MaxInt {
		return -1
	} else {
		return prices[dst]
	}
}
