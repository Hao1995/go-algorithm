package validanagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// hash map, record each letter, and compared them.
	sMap := make(map[rune]int)
	for _, c := range s {
		sMap[c]++
	}

	tMap := make(map[rune]int)
	for _, c := range t {
		tMap[c]++
	}

	// compare them
	for k, sc := range sMap {
		tc := tMap[k]
		if tc != sc {
			return false
		}
	}

	return true
}
