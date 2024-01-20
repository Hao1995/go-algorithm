package minimumheighttrees

func findMinHeightTrees(n int, edges [][]int) []int {
	// topology
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{0, 1}
	}

	// build graph, calculate degree(count the number of adjacent nodes)
	graph := make(map[int][]int)
	degree := make([]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
		degree[a]++
		degree[b]++
	}

	// BFS to remove the outermost nodes
	var queue []int

	// init queue
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}

	// start removing
	for n > 2 {
		size := len(queue)
		n -= size

		for i := 0; i < size; i++ {
			node := queue[i]
			degree[node]--
			for _, nextNode := range graph[node] {
				degree[nextNode]--
				if degree[nextNode] == 1 {
					queue = append(queue, nextNode)
				}
			}
		}

		queue = queue[size:]
	}

	return queue
}
