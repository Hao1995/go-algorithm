package encodeanddecodestrings

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		length := utf8.RuneCountInString(str)
		sb.WriteString(strconv.Itoa(length))
		sb.WriteString("#")
		sb.WriteString(str)
	}
	return sb.String()
}

func decode(input string) []string {
	ans := make([]string, 0)

	var i int
	for i < len(input) {
		// find #
		j := i
		for j < len(input) && input[j] != '#' {
			j++
		}

		// find length
		length, _ := strconv.Atoi(input[i:j])

		// skip #
		j++

		// find string
		runeArr := []rune(input[j:])
		runeStr := runeArr[:length]
		str := string(runeStr)

		ans = append(ans, str)

		i = j + len([]byte(str))
	}

	return ans
}
