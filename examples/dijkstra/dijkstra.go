package dijkstra

import (
	"container/heap"
	"math"
)

type Edge struct {
	neighbor int
	weight   int
}

type Node struct {
	vertex int
	dist   int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Time Complexity: O(V+ElogE)
func dijkstra(graph map[int][]Edge, start int) map[int]int {
	// Make distance array >> O(V)
	dist := make(map[int]int)
	for vertex := range graph {
		dist[vertex] = math.MaxInt64
	}
	dist[start] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Node{start, 0})

	// Worst case: push all edges to the PQ >> O(E*LogE).
	for len(pq) > 0 {
		node := heap.Pop(&pq).(*Node) // O(logE)
		u := node.vertex

		for _, edge := range graph[u] {
			v := edge.neighbor
			weight := edge.weight

			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(&pq, &Node{v, dist[v]})
			}
		}
	}

	return dist
}
