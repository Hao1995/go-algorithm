# 347. Top K Frequent Elements

## Intuition
// better than O(nlogn)

// Approach 1: Not better than O(nlogn)
// The first loop: count all element >> O(n)
// Sort: sort the frequency in decreasing order. >> O(nlogn)
// Retrieve the top k element: O(1)

// Approach 2:
// The first loop: count all element >> O(n)
// Init a MaxHeap: O(n)
// Retrieve the top k element: O(klogn)

## Approach
Use approach 2.