package lengthoflongestsubstring

func lengthOfLongestSubstring(s string) int {
	var ans int

	exist := make(map[byte]bool)

	var l, r int
	for l <= r && r < len(s) {
		c := s[r]
		if _, ok := exist[c]; ok {
			for l < r {
				lc := s[l]
				delete(exist, lc)
				l++
				if lc == c {
					break
				}
			}
		}
		exist[c] = true
		ans = max(ans, r-l+1)
		r++
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
