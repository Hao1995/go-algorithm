package wordladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// *ot:[hot, dot, lot]
	neg := make(map[string][]string)
	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			pattern := word[:i] + "*" + word[i+1:]
			neg[pattern] = append(neg[pattern], word)
		}
	}

	// BFS
	queue := []string{beginWord}
	exist := make(map[string]bool)
	var ans int = 1
	for len(queue) > 0 {
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			var word string
			word, queue = queue[0], queue[1:]
			if word == endWord {
				return ans
			}

			for j := 0; j < len(word); j++ {
				pattern := word[:j] + "*" + word[j+1:]
				for _, neiWord := range neg[pattern] {
					if _, ok := exist[neiWord]; !ok {
						exist[neiWord] = true
						queue = append(queue, neiWord)
					}
				}
			}
		}
		ans++
	}

	return 0
}
