package redundantconnection

func findRedundantConnection(edges [][]int) []int {
	// [1,2,3,...,n]
	parents := make([]int, len(edges)+1)
	for i := 1; i < len(parents); i++ {
		parents[i] = i
	}

	// [1,1,1,...,1]
	ranks := make([]int, len(edges)+1)

	findParent := func(node int) int {
		p := parents[node]
		for p != parents[p] {
			parents[p] = parents[parents[p]]
			p = parents[p]
		}
		return p
	}

	union := func(n1, n2 int) bool {
		p1, p2 := findParent(n1), findParent(n2)

		if p1 == p2 {
			return false
		}

		if ranks[p1] >= ranks[p2] {
			parents[p2] = p1
			ranks[p1] += ranks[p2]
		} else {
			parents[p1] = p2
			ranks[p2] += ranks[p1]
		}

		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge
		}
	}

	return []int{0, 0}
}
