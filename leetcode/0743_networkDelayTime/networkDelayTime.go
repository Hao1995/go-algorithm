package networkdelaytime

import "container/heap"

type Node struct {
	V int
	W int
}

type MinHeap []Node

func (mh MinHeap) Len() int           { return len(mh) }
func (mh MinHeap) Less(i, j int) bool { return mh[i].W < mh[j].W }
func (mh MinHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh *MinHeap) Push(item any)     { *mh = append(*mh, item.(Node)) }
func (mh *MinHeap) Pop() any {
	item := (*mh)[mh.Len()-1]
	(*mh) = (*mh)[:mh.Len()-1]
	return item
}

func networkDelayTime(times [][]int, n int, k int) int {
	// construct the graph
	graph := make(map[int][]Node)
	for _, time := range times {
		graph[time[0]] = append(graph[time[0]], Node{time[1], time[2]})
	}

	// find the shortest path using dijkstra's algo
	visited := make(map[int]bool)
	neighbors, found := graph[k]
	if !found {
		return -1
	}
	visited[k] = true

	mh := MinHeap(neighbors)
	heap.Init(&mh)

	var ans int
	for mh.Len() > 0 {
		node := heap.Pop(&mh).(Node)
		if _, found := visited[node.V]; found {
			continue
		}
		visited[node.V] = true
		ans = max(ans, node.W)

		neighbors, found := graph[node.V]
		if !found {
			continue
		}

		for _, neighbor := range neighbors {
			heap.Push(&mh, Node{neighbor.V, node.W + neighbor.W})
		}
	}

	if len(visited) == n {
		return ans
	}
	return -1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
