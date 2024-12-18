package evalrpn

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int32
	operator := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	for _, token := range tokens {
		if operator[token] {
			// pop operand
			var num1, num2 int32
			num2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			num1, stack = stack[len(stack)-1], stack[:len(stack)-1]

			var num int32
			switch token {
			case "+":
				num = num1 + num2
			case "-":
				num = num1 - num2
			case "*":
				num = num1 * num2
			case "/":
				num = num1 / num2
			}

			stack = append(stack, num)
		} else {
			// number
			num, _ := strconv.ParseInt(token, 10, 8)
			stack = append(stack, int32(num))
		}
	}

	return int(stack[0])
}
