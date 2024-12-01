package swiminrisingwater

import "container/heap"

// {time, r, c}
type MinHeap [][3]int

func (mh MinHeap) Len() int           { return len(mh) }
func (mh MinHeap) Less(i, j int) bool { return mh[i][0] < mh[j][0] }
func (mh MinHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh *MinHeap) Push(item any)     { *mh = append(*mh, item.([3]int)) }
func (mh *MinHeap) Pop() any {
	item := (*mh)[mh.Len()-1]
	*mh = (*mh)[:mh.Len()-1]
	return item
}

func swimInWater(grid [][]int) int {
	// init a visited grid
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}

	// init a min heap
	mh := MinHeap{{grid[0][0], 0, 0}}
	heap.Init(&mh)

	// Dijkstra' Algo
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for mh.Len() > 0 {
		item := heap.Pop(&mh).([3]int)
		t, r, c := item[0], item[1], item[2]

		if visited[r][c] {
			continue
		}

		visited[r][c] = true

		if r == len(grid)-1 && c == len(grid[0])-1 {
			return t
		}

		for _, dir := range dirs {
			dr, dc := dir[0], dir[1]
			nr, nc := r+dr, c+dc
			if nr < 0 || nr == len(grid) || nc < 0 || nc == len(grid[0]) {
				// out of the boundary
				continue
			}
			heap.Push(&mh, [3]int{max(t, grid[nr][nc]), nr, nc})
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
