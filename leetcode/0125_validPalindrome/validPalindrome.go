package validpalindrome

import "strings"

func isPalindrome(s string) bool {
	nsb := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9') {
			nsb = append(nsb, s[i])
		}
	}
	ns := strings.ToLower(string(nsb))

	var i, j int = 0, len(ns) - 1
	for i < j {
		if ns[i] != ns[j] {
			return false
		}
		i++
		j--
	}
	return true
}
