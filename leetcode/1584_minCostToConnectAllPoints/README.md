# 1584. Min Cost to Connect All Points

## Intuition
Construct the MST using Prim's algorithm.

## Approach
1. Generate all edges with manhattan distance >> O(n^2)
2. Find the MST using Prim's algorithm (MinHeap) >> O(n^2 * log(n^2))

* Why my Prim's algorithm is O(n^2 * log(n^2)) not typically O(n^2 * logn) ?
I think it's because some method will update the origin vertex instead of adding new vertex. (not sure...)
https://chatgpt.com/share/66efd23d-a764-8001-94c6-caf3e59ab9b1