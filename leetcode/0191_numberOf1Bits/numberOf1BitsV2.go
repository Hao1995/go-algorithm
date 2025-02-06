package numberof1bits

func hammingWeightV2(n int) int {
	var count int
	for n > 0 {
		if (n & 1) == 1 {
			count++
		}
		n >>= 1
	}
	return count
}
