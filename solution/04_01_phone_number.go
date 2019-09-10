package solution

import "strings"

func PhoneNumber(S string) string {
	// write your code in Go 1.4
	// remove " ", "-"

	S = strings.Replace(strings.Replace(S, " ", "", -1), "-", "", -1)

	rS := ""
	for idx, c := range S {
		if len(S)-idx == 2 && idx%3 == 2 {
			rS += "-"
		} else if idx != len(S)-1 && idx != 0 && idx%3 == 0 {
			rS += "-"
		}
		rS += string(c)
	}

	return rS
}
