/*
207. Course Schedule
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.
*/
package canfinish

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	for _, pre := range prerequisites {
		graph[pre[0]] = append(graph[pre[0]], pre[1])
	}

	visited := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if hasCycle(graph, visited, i) {
			return false
		}
	}

	return true
}

func hasCycle(graph [][]int, visited []int, node int) bool {
	if visited[node] == 1 {
		return true
	}

	if visited[node] == -1 {
		return false
	}

	visited[node] = 1

	for _, neighbor := range graph[node] {
		if hasCycle(graph, visited, neighbor) {
			return true
		}
	}

	visited[node] = -1
	return false
}
