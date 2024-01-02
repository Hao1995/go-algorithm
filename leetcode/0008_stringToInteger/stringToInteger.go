package stringtointeger

func myAtoi(s string) int {
	// ignore whitespace
	var idx int = 0
	for idx < len(s) && s[idx] == ' ' {
		idx++
	}

	// only allowed '-', '+'
	var sign int = 1
	if idx < len(s) && (s[idx] == '-' || s[idx] == '+') {
		if s[idx] == '-' {
			sign = -1
		}
		idx++
	}

	// check number
	var MAX, MIN int = 1<<31 - 1, -1 << 31
	var ans int
	for i := idx; i < len(s); i++ {
		num := int(s[i] - '0')
		if num <= 9 {
			tmpAns := ans*10 + num
			if sign > 0 && (tmpAns < ans || tmpAns > MAX) {
				return MAX
			}
			if sign < 0 && (tmpAns < ans || (tmpAns*-1) < MIN) {
				return MIN
			}
			ans = tmpAns
		} else {
			break
		}
	}

	if sign > 0 {
		return ans
	}

	return ans * -1
}
