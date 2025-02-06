package groupanagrams

import "fmt"

func groupAnagramsV2(strs []string) [][]string {
	generateKey := func(str string) string {
		letters := make([]int, 26)

		// O(n)
		for i := 0; i < len(str); i++ {
			letters[int(str[i]-'a')]++
		}

		// O(1)
		var ans string
		for _, count := range letters {
			// %3d because length size <= 100
			ans = fmt.Sprintf("%s%3d", ans, count)
		}

		return ans
	}

	// O(k) * O(n)
	groups := make(map[string][]string)
	for _, str := range strs {
		key := generateKey(str)
		groups[key] = append(groups[key], str)
	}

	// O(v)
	ans := make([][]string, 0, len(groups))
	for _, group := range groups {
		ans = append(ans, group)
	}

	// O(kn+v)
	return ans
}
