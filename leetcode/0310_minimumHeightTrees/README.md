# 310. Minimum Height Trees

## Intuition
Start from each node and use BFS to walk outward layer by layer until the end.
Finally, we will know which nodes have the smallest distance to the outermost layer.

But it will spend O(n^2).

So I find a solution using Topology Sort and BFS only need O(n).

## Approach
1. Record the degree of each node >> O(n)
2. Starting to remove the outermost nodes that have only 1 degree.
3. Stop the loop when we find the possible nodes (n <= 2)
