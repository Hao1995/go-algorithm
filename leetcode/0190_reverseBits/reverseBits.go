package reversebits

func reverseBits(num uint32) uint32 {
	var ans uint32
	var quotient, remainder uint32 = num, 0
	var count int
	for quotient > 0 {
		remainder = quotient % 2
		quotient /= 2
		ans *= 2
		ans += remainder
		count++
	}

	for i := count + 1; i <= 32; i++ {
		ans *= 2
	}

	return ans
}
