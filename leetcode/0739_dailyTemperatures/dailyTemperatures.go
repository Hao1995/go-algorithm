package dailytemperatures

func dailyTemperatures(temperatures []int) []int {
	var ans []int = make([]int, len(temperatures))
	var stack []int = make([]int, 0, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else if temperatures[i] <= temperatures[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 {
				lastIdx := stack[len(stack)-1]
				if temperatures[i] > temperatures[lastIdx] {
					ans[lastIdx] = i - lastIdx   // cal the num of days
					stack = stack[:len(stack)-1] // pop
				} else {
					break
				}
			}
			stack = append(stack, i)
		}
	}

	return ans
}
