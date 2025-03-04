package validparenthesisstring

import "fmt"

func checkValidString(s string) bool {
	var leftMin, leftMax int

	for _, c := range s {
		switch c {
		case rune('('):
			leftMin++
			leftMax++
		case rune(')'):
			leftMin--
			leftMax--
		case rune('*'):
			leftMin--
			leftMax++
		default:
			fmt.Println("wrong symbol")
			return false
		}

		if leftMax < 0 {
			return false
		}

		if leftMin < 0 {
			leftMin = 0
		}
	}

	return leftMin == 0
}
