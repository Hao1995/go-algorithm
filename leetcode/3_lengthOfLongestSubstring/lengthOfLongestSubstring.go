/*
3. Longest Substring Without Repeating Characters
Given a string s, find the length of the longest
substring without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
package lengthoflongestsubstring

func lengthOfLongestSubstring(s string) int {
	var longestLen int = 0
	checkMap := make(map[rune]int)
	for idx, r := range s {
		prevIdx, ok := checkMap[r]
		if ok {
			checkMap = initCheckMap(s, prevIdx, idx)
		} else {
			checkMap[r] = idx
		}
		longestLen = max(longestLen, len(checkMap))
	}

	return longestLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func initCheckMap(s string, prevIdx int, idx int) map[rune]int {
	checkMap := make(map[rune]int)
	for i := prevIdx + 1; i < idx+1; i++ {
		checkMap[rune(s[i])] = i
	}
	return checkMap
}
