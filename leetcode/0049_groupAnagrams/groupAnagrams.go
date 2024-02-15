package groupanagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	sortedMap := make(map[string][]string)
	for _, s := range strs {
		v := []byte(s)
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		vStr := string(v)
		sortedMap[vStr] = append(sortedMap[vStr], s)
	}

	var ans [][]string
	for _, v := range sortedMap {
		ans = append(ans, v)
	}

	return ans
}
