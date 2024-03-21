// This example demonstrates an integer heap built using the heap interface.
package maxheap

import (
	"fmt"
)

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func Example_MaxHeap() {
	heap := &MaxHeap{}
	heap.insert(0)
	heap.print()
	heap.insert(1)
	heap.print()
	heap.insert(2)
	heap.print()
	heap.insert(3)
	heap.print()
	heap.insert(4)
	heap.print()
	heap.insert(5)
	heap.print()
	heap.insert(6)
	heap.print()
	heap.insert(7)
	heap.print()
	heap.insert(8)
	heap.print()
	heap.insert(9)
	heap.print()
	heap.insert(10)
	heap.print()
	heap.insert(11)
	heap.print()
	heap.insert(12)
	heap.print()

	fmt.Printf("Extract=%v\n", heap.extract()) // 12
	heap.print()
	fmt.Printf("Extract=%v\n", heap.extract()) // 11
	heap.print()
	fmt.Printf("Extract=%v\n", heap.extract()) // 10
	heap.print()
	fmt.Printf("Extract=%v\n", heap.extract()) // 9
	heap.print()
	fmt.Printf("Extract=%v\n", heap.extract()) // 8
	heap.print()
	fmt.Printf("Extract=%v\n", heap.extract()) // 7
	heap.print()
	fmt.Printf("Extract=%v\n", heap.extract()) // 6
	heap.print()
	fmt.Printf("Extract=%v\n", heap.extract()) // 5
	heap.print()
	fmt.Printf("Extract=%v\n", heap.extract()) // 4
	heap.print()
	fmt.Printf("Extract=%v\n", heap.extract()) // 3
	heap.print()
	fmt.Printf("Extract=%v\n", heap.extract()) // 2
	heap.print()
	fmt.Printf("Extract=%v\n", heap.extract()) // 1
	heap.print()
	fmt.Printf("Extract=%v\n", heap.extract()) // 0
	heap.print()
}
