# 973. K Closest Points to Origin

## Intuition
Sort or Priority Queue

## Approach
// Approach.1
// Calculate all Euclidean distance: O(n)
// Sort the distance array with increading order: O(nlogn)
// Return the top K closest points: O(k)
// Overall: O(k+nlogn)

// Approach.2
// Calculate all Euclidean distance and store them into a MaxHeap: O(n*logn)
// Return the top K closest points: O(k*logn)
// Overall: O((k+n)logn)

// Approach.3 (my V2 solution)
// Iterate input:
//      Insert to the MaxHeap: O(nlogn)
//      Remove the max item from the MaxHeap if length more than k: O((n-k)logn)
// Overall: O(nlogn)

// Approach. 4 (my V1 solution)
// Sort the points with build-in distance calculation.
// Overall: O(nlogn)