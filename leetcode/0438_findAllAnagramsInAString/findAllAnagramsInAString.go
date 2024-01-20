package findallanagramsinastring

func findAnagrams(s string, p string) []int {
	// O(n)
	pMap := make(map[byte]int32)
	for i := 0; i < len(p); i++ {
		pMap[p[i]] += 1
	}

	var ans []int
	var i, j int
	for j < len(s) {
		c := s[j]
		count, ok := pMap[c]
		if !ok {
			for i < j {
				pMap[s[i]] += 1
				i++
			}
			j++
			i++
		} else if count == 0 {
			pMap[s[i]] += 1
			i++
		} else {
			pMap[c] -= 1
			if pMap[c] == 0 && (j-i+1) == len(p) {
				ans = append(ans, i)
			}
			j++
		}
	}

	return ans
}
