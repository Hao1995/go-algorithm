package longestrepeatingcharacterreplacement

func characterReplacement(s string, k int) int {
	countMap := make(map[byte]int)
	var l, maxF, ans int

	for r := 0; r < len(s); r++ {
		countMap[s[r]]++

		maxF = max(maxF, countMap[s[r]])

		if (r-l+1)-maxF > k {
			countMap[s[l]]--
			l++
		}

		ans = max(ans, r-l+1)
	}

	return ans
}

func easyCharacterReplacement(s string, k int) int {
	countMap := make(map[byte]int)
	var l, r int
	countMap[s[r]]++
	var ans int
	for r < len(s) {

		var maxHz, totalHz int
		for _, v := range countMap {
			maxHz = max(maxHz, v)
			totalHz += v
		}
		otherHz := totalHz - maxHz

		if otherHz <= k {
			ans = max(ans, r-l+1)
			r++
			if r < len(s) {
				countMap[s[r]]++
			} else {
				break
			}
		} else {
			countMap[s[l]]--
			l++
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
