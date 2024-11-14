package minimumwindowsubstring

func minWindow(s string, t string) string {
	sCountMap, tCountMap := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tCountMap[t[i]]++
	}

	var have, need int = 0, len(tCountMap)

	minSet := [2]int{-1, -1}
	var minLen int = 1<<31 - 1

	var left int
	for right := 0; right < len(s); right++ {
		sCountMap[s[right]]++
		if sCountMap[s[right]] == tCountMap[s[right]] {
			have++
		}

		for have == need {
			// update min
			if right-left+1 < minLen {
				minLen = right - left + 1
				minSet = [2]int{left, right}
			}

			sCountMap[s[left]]--
			if _, ok := tCountMap[s[left]]; ok && sCountMap[s[left]] < tCountMap[s[left]] {
				have--
			}
			left++
		}
	}

	if minLen == 1<<31-1 {
		return ""
	}

	return s[minSet[0] : minSet[1]+1]
}
