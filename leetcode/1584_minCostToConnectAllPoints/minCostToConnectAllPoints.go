package mincosttoconnectallpoints

import "container/heap"

type MinHeap [][2]int

func (this MinHeap) Len() int           { return len(this) }
func (this MinHeap) Less(i, j int) bool { return this[i][0] < this[j][0] }
func (this MinHeap) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this *MinHeap) Push(item any)     { *this = append(*this, item.([2]int)) }
func (this *MinHeap) Pop() any {
	item := (*this)[this.Len()-1]
	(*this) = (*this)[:this.Len()-1]
	return item
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)

	// O(n^2)
	edges := make([][][2]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])
			edges[i] = append(edges[i], [2]int{dist, j})
			edges[j] = append(edges[j], [2]int{dist, i})
		}
	}

	// Prim's Algo: O(n^2*log(n^2))
	var ans int
	pq := MinHeap([][2]int{{0, 0}})
	heap.Init(&pq)
	visited := make(map[int]bool)
	for len(visited) < n {
		point := heap.Pop(&pq).([2]int)

		// already exist, pass
		if _, ok := visited[point[1]]; ok {
			continue
		}

		visited[point[1]] = true
		ans += point[0]

		for _, edge := range edges[point[1]] {
			heap.Push(&pq, edge)
		}
	}

	return ans
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
