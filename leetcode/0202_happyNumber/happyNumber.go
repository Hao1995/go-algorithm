package happynumber

func isHappy(n int) bool {
	existMap := make(map[int]bool)
	for {
		// Repeat the process
		var quotient, remainder int = n, 0
		n = 0
		for quotient > 0 {
			remainder = quotient % 10
			n += remainder * remainder
			quotient /= 10
		}

		// endless check
		if _, ok := existMap[n]; ok {
			return false
		}
		existMap[n] = true

		// happy number check
		if n == 1 {
			return true
		}
	}
}
