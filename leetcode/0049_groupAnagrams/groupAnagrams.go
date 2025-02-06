package groupanagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	// T:O(k) * O(nlogn), S:O(k)
	sortedMap := make(map[string][]string)
	for _, s := range strs {
		v := []byte(s)
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		vStr := string(v)
		sortedMap[vStr] = append(sortedMap[vStr], s)
	}

	// T: O(v)
	var ans [][]string
	for _, v := range sortedMap {
		ans = append(ans, v)
	}

	return ans
}
