package longestrepeatingcharacterreplacement

func characterReplacement(s string, k int) int {
	var ans int
	countMap := make(map[byte]int)

	maxFreqNum := func() int {
		var ans int
		for _, count := range countMap {
			ans = max(ans, count)
		}
		return ans
	}

	var l, r int = 0, 0

	for r < len(s) {
		countMap[s[r]]++
		num := maxFreqNum()
		for l < r && (r-l+1-num) > k {
			countMap[s[l]]--
			l++
		}
		ans = max(ans, r-l+1)
		r++
	}

	return ans
}
