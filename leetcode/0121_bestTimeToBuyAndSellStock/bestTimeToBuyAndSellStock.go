package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	var minPrice, ans int = prices[0], 0
	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > ans {
			ans = profit
		}
	}
	return ans
}
