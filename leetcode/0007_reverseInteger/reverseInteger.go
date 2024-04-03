package reverseinteger

func reverse(x int) int {
	const MAX, MIN int = 1<<31 - 1, -1 * (1 << 31)
	var ans int
	for x != 0 {
		ans = ans*10 + x%10

		// check overflow
		if ans > MAX || ans < MIN {
			return 0
		}

		x /= 10
	}
	return ans
}
