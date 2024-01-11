package longestpalindromicsubstring

func longestPalindrome(s string) string {
	var maxSize int = 0
	var maxPalindromeStr string
	for i := 0; i < len(s); i++ {
		str, oddSize := extendPalindrome(s, i, i)
		if oddSize > maxSize {
			maxSize = oddSize
			maxPalindromeStr = str
		}

		str, evenSize := extendPalindrome(s, i, i+1)
		if evenSize > maxSize {
			maxSize = evenSize
			maxPalindromeStr = str
		}
	}
	return maxPalindromeStr
}

func extendPalindrome(s string, left int, right int) (string, int) {
	if left < 0 && right >= len(s) {
		return "", -1
	}

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	// correct left and right
	left++
	right--

	return s[left : right+1], right - left + 1
}
