package basiccalculatorii

import "fmt"

func calculate(s string) int {
	var ans, prev int
	var operator byte = '+'
	var i int
	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			var curr int = int(s[i] - '0')
			i++
			for i < len(s) && (s[i] >= '0' && s[i] <= '9') {
				curr = curr*10 + int(s[i]-'0')
				i++
			}
			i--

			switch operator {
			case '+':
				ans += curr
				prev = curr
			case '-':
				ans -= curr
				prev = -curr
			case '*':
				ans -= prev
				ans += prev * curr
				prev = prev * curr
			case '/':
				ans -= prev
				ans += prev / curr
				prev = prev / curr
			default:
				fmt.Println("not a valid operator")
			}
		} else if s[i] != ' ' {
			operator = s[i]
		}
		i++
	}

	return ans
}
