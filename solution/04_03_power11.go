package solution

func Power11KKDay(N int) int {
	// write your code in Go 1.4

	fibArr := fibnacci(N)
	count := 0
	overflow := 0
	for i := len(fibArr) - 1; i > -1; i-- {
		val := fibArr[i] + overflow
		val1 := val % 10 // ex: 17%10 = 7
		overflow = val / 10
		if val1 == 1 {
			count++
		}
	}
	if overflow == 1 {
		count++
	}
	return count
}

func fibnacci(N int) []int {

	if N == 0 {
		return []int{1}
	}

	if N == 1 {
		return []int{1, 1}
	}

	res := []int{}
	lastItem := fibnacci(N - 1)
	// fmt.Printf("CN=%v, LastArr=%v\n", N, lastItem)
	for i := 0; i <= N; i++ {
		if i == 0 || i == N {
			res = append(res, 1)
			continue
		}
		res = append(res, lastItem[i-1]+lastItem[i])
	}
	return res
}
