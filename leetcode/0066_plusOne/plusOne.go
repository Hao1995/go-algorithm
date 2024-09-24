package plusone

func plusOne(digits []int) []int {
	var quotient int = 1
	for i := len(digits) - 1; i >= 0; i-- {
		if quotient == 0 {
			break
		}

		nd := digits[i] + quotient
		quotient = nd / 10
		digits[i] = nd % 10
	}

	if quotient != 0 {
		return append([]int{quotient}, digits...)
	}

	return digits
}
