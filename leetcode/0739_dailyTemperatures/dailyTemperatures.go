package dailytemperatures

func dailyTemperatures(temperatures []int) []int {
	// store the index of temperatures
	stack := make([]int, 0, len(temperatures))
	output := make([]int, len(temperatures))

	for idx, temp := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, idx)
			continue
		}

		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			output[stack[len(stack)-1]] = idx - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, idx)
	}

	return output
}
