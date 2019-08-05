package solution

import "strings"

// Test Cases :
// ("nice", "cein")
// ("lbaba", "lbbbb")
// ("llabe", "lalbe")

func CombWords(S, T string) string {

	if len(S) == len(T)+1 || len(S)+1 == len(T) {
		// S include T; REMOVE
		if strings.Contains(S, T) {
			return "REMOVE " + strings.Replace(S, T, "", -1)
		}
		// T include S; INSERT
		if strings.Contains(T, S) {
			return "INSERT " + strings.Replace(T, S, "", -1)
		}
	}

	// SWAP
	if len(S) == len(T) {
		var c byte
		for i := 0; i < len(S); i++ {
			if S[i] != T[i] {
				if c == 0 {
					c = S[i]
				} else if T[i] == c {
					return "SWAP " + string(c) + " " + string(S[i])
				}
			}
		}
	}

	return "IMPOSSIBLE"
}
