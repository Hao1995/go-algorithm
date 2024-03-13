package coursescheduleii

func findOrder(numCourses int, prerequisites [][]int) []int {
	childMap := make(map[int][]int)
	countParent := make([]int, numCourses)
	for _, pre := range prerequisites {
		curr, parent := pre[0], pre[1]
		childMap[parent] = append(childMap[parent], curr)
		countParent[curr]++
	}

	// Find heads
	var queue []int
	for num, count := range countParent {
		if count == 0 {
			queue = append(queue, num)
		}
	}

	var ans []int
	var curr int
	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		ans = append(ans, curr)

		for _, child := range childMap[curr] {
			countParent[child]--
			if countParent[child] == 0 {
				queue = append(queue, child)
			}
		}
	}

	if len(ans) == numCourses {
		return ans
	}

	return []int{}
}
