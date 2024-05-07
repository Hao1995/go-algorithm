package numberofrestrictedpathsfromfirsttolastnode

import "container/heap"

// [city, weight]
type PQ [][]int

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i][1] < pq[j][1] }
func (pq PQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(item any)     { *pq = append(*pq, item.([]int)) }
func (pq *PQ) Pop() any {
	l := len(*pq)
	old := *pq
	item := old[l-1]
	*pq = old[:l-1]
	return item
}

func countRestrictedPaths(n int, edges [][]int) int {
	graph := make(map[int][][]int)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		graph[u] = append(graph[u], []int{v, w})
		graph[v] = append(graph[v], []int{u, w})
	}

	cache := make([]int, n+1)
	for i := 1; i < len(cache); i++ {
		cache[i] = 1 << 31
	}

	var pq PQ
	var item []int
	heap.Init(&pq)
	heap.Push(&pq, []int{n, 0})
	for pq.Len() > 0 {
		item = heap.Pop(&pq).([]int)
		if item[1] >= cache[item[0]] {
			continue
		}
		cache[item[0]] = item[1]

		for _, neig := range graph[item[0]] {
			w := item[1] + neig[1]
			if w < cache[neig[0]] {
				heap.Push(&pq, []int{neig[0], w})
			}
		}
	}

	pathCache := make([]int, len(cache))
	for i := 0; i < len(pathCache); i++ {
		pathCache[i] = -1
	}

	return dfs(graph, cache, n, 1, pathCache)
}

func dfs(graph map[int][][]int, cache []int, n int, curr int, pathCache []int) int {
	if curr == n {
		return 1
	}

	if pathCache[curr] != -1 {
		return pathCache[curr]
	}

	var ans int
	for _, neig := range graph[curr] {
		if cache[curr] > cache[neig[0]] {
			ans += dfs(graph, cache, n, neig[0], pathCache)
			ans %= 1000_000_007
		}
	}

	pathCache[curr] = ans
	return ans
}
