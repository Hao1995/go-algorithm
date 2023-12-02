/*
150. Evaluate Reverse Polish Notation
https://leetcode.com/problems/evaluate-reverse-polish-notation/
You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:

The valid operators are '+', '-', '*', and '/'.
Each operand may be an integer or another expression.
The division between two integers always truncates toward zero.
There will not be any division by zero.
The input represents a valid arithmetic expression in a reverse polish notation.
The answer and all the intermediate calculations can be represented in a 32-bit integer.
*/
package evalrpn

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		if isOperand(token) {
			var num1, num2 int
			num1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			num2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			newNum := calculate(token, num2, num1)
			stack = append(stack, newNum)
			continue
		}
		num, _ := strconv.Atoi(token)
		stack = append(stack, num)
	}
	return stack[0]
}

func isOperand(str string) bool {
	return str == "*" || str == "/" || str == "-" || str == "+"
}

func calculate(operand string, a, b int) int {
	switch operand {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return -1
	}
}
