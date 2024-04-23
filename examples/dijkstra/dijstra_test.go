package dijkstra

import "fmt"

func Example_NormalCase() {
	graph := map[int][]Edge{
		0: {{1, 4}, {2, 1}},
		1: {{3, 1}},
		2: {{1, 2}, {3, 5}},
		3: {{4, 3}},
		4: {},
	}

	start := 0
	distances := dijkstra(graph, start)

	for vertex, dist := range distances {
		fmt.Printf("Vertex %d -> Distance: %d\n", vertex, dist)
	}
}
