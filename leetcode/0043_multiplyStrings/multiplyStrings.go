package multiplystrings

import "fmt"

func multiply(num1 string, num2 string) string {
	n, m := len(num1), len(num2)
	size := n + m
	arr := make([]int, size)
	for i2 := m - 1; i2 >= 0; i2-- {
		number2 := int(num2[i2] - '0')
		for i1 := n - 1; i1 >= 0; i1-- {
			number1 := int(num1[i1] - '0')
			idx := size - (m - 1 - i2) - (n - 1 - i1)
			total := number2*number1 + arr[idx-1]
			q, r := total/10, total%10
			arr[idx-2] += q
			arr[idx-1] = r
		}
	}

	var ans string
	hasNonZero := false
	for _, num := range arr {
		if num != 0 {
			hasNonZero = true
		}

		if !hasNonZero {
			continue
		}

		ans = fmt.Sprintf("%s%d", ans, num)
	}

	if len(ans) == 0 {
		return "0"
	}

	return ans
}
