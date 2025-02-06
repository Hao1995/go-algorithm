package numberof1bits

func hammingWeight(n int) int {
	var count int

	for n >= 2 {
		remainder := n % 2
		n /= 2

		if remainder == 1 {
			count++
		}
	}

	if n == 1 {
		count++
	}

	return count
}
