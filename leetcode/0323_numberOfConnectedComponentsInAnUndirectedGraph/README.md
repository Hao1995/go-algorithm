# 323. Number of Connected Components in an Undirected Graph

## Intuition
Those nodes are from 0 to n(not included). So my first though is I need to iterate all nodes within the graph.
But how to calculate the number of components? Every time we iterate a new component, that number has definitely not been traversed yet.
So, whenever that number hasn't traversed, we can increment the count by one.

## Approach
DFS + Visited Array
Time Complexity: O(n+m) (m is the number of edges)
Space Complexity: O(n+n)