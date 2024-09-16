package validparentheses

func isValid(s string) bool {
	mappingBracket := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	var stack []rune
	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		var lr rune
		lr, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if mr := mappingBracket[lr]; mr != r {
			return false
		}
	}

	return len(stack) == 0
}
