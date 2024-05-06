# 215. Kth Largest Element in an Array

## Intuition
MaxHeap(Priority Queue)
Even the question requirement is without sorting, I can push the num to a priority queue to achieve the same functionality.

## Approach
1. Store nums to a priority queue. Time complexity is O(n*logn)
2. Retrieve top k elements from the priority queue. Time complexity is O(k*logn)

Total time complexity is O((k+n)logn).
