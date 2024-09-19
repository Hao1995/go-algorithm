package carfleet

import "sort"

func carFleet(target int, position []int, speed []int) int {
	sorted := make([][2]int, len(position))
	for i, p := range position {
		sorted[i] = [2]int{p, speed[i]}
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] > sorted[j][0]
	})

	var stack []float32
	for _, c := range sorted {
		stack = append(stack, float32(target-c[0])/float32(c[1]))
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)
}
