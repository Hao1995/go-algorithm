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
	idxMap := make(map[rune]int)
	var maxCount int
	var count int
	for currIdx, c := range s {
		if repeatedIdx, ok := idxMap[c]; ok {
			count = len(s[repeatedIdx+1 : currIdx+1])

			// initial idxMap
			idxMap = make(map[rune]int)
			for i := repeatedIdx + 1; i < currIdx+1; i++ {
				idxMap[rune(s[i])] = i
			}
			continue
		}

		count++
		idxMap[c] = currIdx
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}
