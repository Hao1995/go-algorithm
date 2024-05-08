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

func hashFindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// O(n)
	const MAX = 1 << 31
	price := make([]int, n)
	for i := 0; i < n; i++ {
		price[i] = MAX
	}
	price[src] = 0

	// O(E)
	// src:destination[s,w]
	destMap := make(map[int][][]int)
	for _, flight := range flights {
		s, d, p := flight[0], flight[1], flight[2]
		destMap[s] = append(destMap[s], []int{d, p})
	}

	// O(k+1)
	var step int
	for step <= k {
		tmpPrice := make([]int, len(price))
		copy(tmpPrice, price)

		// O(n)
		for s, sp := range price {
			if sp == MAX {
				continue
			}

			for _, dest := range destMap[s] {
				d, dp := dest[0], dest[1]
				tmpPrice[d] = min(tmpPrice[d], sp+dp)
			}
		}

		step++
		price = tmpPrice
	}

	if price[dst] == MAX {
		return -1
	}
	return price[dst]
}
