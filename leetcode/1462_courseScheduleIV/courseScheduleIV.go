package coursescheduleiv

// n=numCourses
// len(prerequisites)=n^2
// len(queries)=m
// Total:O(n^2+n+m) >> O(n^2+m)
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	// course:[]prerequistes
	// O(n^2)
	graph := make(map[int][]int)
	for _, prereq := range prerequisites {
		graph[prereq[1]] = append(graph[prereq[1]], prereq[0])
	}

	// find all prerequesties
	// course:[]all preprequistes
	// O(n): Because we walk throgh all course only one time.
	prereqMap := make(map[int]map[int]bool)
	for i := 0; i < numCourses; i++ {
		dfs(graph, prereqMap, i)
	}

	// O(m)
	ans := make([]bool, 0, len(queries))
	for _, query := range queries {
		// if u is prereq of v?
		u, v := query[0], query[1]
		ans = append(ans, prereqMap[v][u])
	}

	return ans
}

func dfs(graph map[int][]int, prereqMap map[int]map[int]bool, course int) map[int]bool {
	if v, f := prereqMap[course]; f {
		return v
	}

	for _, prereq := range graph[course] {
		if _, ok := prereqMap[course]; !ok {
			prereqMap[course] = make(map[int]bool)
		}

		originalMap := dfs(graph, prereqMap, prereq)
		for k, v := range originalMap {
			prereqMap[course][k] = v
		}
		prereqMap[course][prereq] = true
	}

	return prereqMap[course]
}
