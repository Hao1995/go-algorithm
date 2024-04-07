package minheap

import (
	"container/heap"
	"testing"
)

// TestMinHeapWhenInsertDataWithAscOrder
// When inserting data with ascending order
// Should get the correctly result
func TestMinHeapWhenInsertDataWithAscOrder(t *testing.T) {
	pq := PQ{}
	heap.Init(&pq)

	heap.Push(&pq, Item{5})
	heap.Push(&pq, Item{4})
	heap.Push(&pq, Item{3})
	heap.Push(&pq, Item{2})
	heap.Push(&pq, Item{1})

	var item any
	for i := 1; i <= 5; i++ {
		item = heap.Pop(&pq)
		if val := item.(Item).val; val != i {
			t.Fatalf("item value should be %d, but got the %v", i, val)
		}
	}
}

// TestMinHeapWhenInsertDataWithDescOrder
// When inserting data with descending order
// Should get the correctly result
func TestMinHeapWhenInsertDataWithDescOrder(t *testing.T) {
	pq := PQ{}
	heap.Init(&pq)

	heap.Push(&pq, Item{1})
	heap.Push(&pq, Item{2})
	heap.Push(&pq, Item{3})
	heap.Push(&pq, Item{4})
	heap.Push(&pq, Item{5})

	var item any
	for i := 1; i <= 5; i++ {
		item = heap.Pop(&pq)
		if val := item.(Item).val; val != i {
			t.Fatalf("item value should be %d, but got the %v", i, val)
		}
	}
}

// TestMinHeapWhenInsertDataWithSOrder
// When inserting data with S order (big to small to big to small ...)
// Should get the correctly result
func TestMinHeapWhenInsertDataWithSOrder(t *testing.T) {
	pq := PQ{}
	heap.Init(&pq)

	heap.Push(&pq, Item{5})
	heap.Push(&pq, Item{1})
	heap.Push(&pq, Item{4})
	heap.Push(&pq, Item{2})
	heap.Push(&pq, Item{3})

	var item any
	for i := 1; i <= 5; i++ {
		item = heap.Pop(&pq)
		if val := item.(Item).val; val != i {
			t.Fatalf("item value should be %d, but got the %v", i, val)
		}
	}
}
