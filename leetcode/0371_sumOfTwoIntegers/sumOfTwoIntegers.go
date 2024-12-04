package sumoftwointegers

func getSum(a int, b int) int {
	for b != 0 {
		var tmp int = (a & b) << 1
		a = a ^ b
		b = tmp
	}
	return a
}
