package powxn

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	var isNegative bool = false
	if n < 0 {
		isNegative = true
		n *= -1
	}

	ans := divideTwo(x, n)
	if isNegative {
		return 1.0 / ans
	}
	return ans
}

func divideTwo(x float64, n int) float64 {
	// Because the divisor is 2
	// remainder must within [0,1]
	var quotient, remainder int = n / 2, n % 2
	if quotient < 2 {
		// final division, quotient less then 2, it should not be divided by 2 again.
		var ans float64 = 1
		if quotient == 1 {
			ans *= x * x
		}

		if remainder == 1 {
			ans *= x
		}

		return ans
	}

	ans := divideTwo(x*x, quotient)
	if remainder == 0 {
		ans *= x
	}
	return ans
}
