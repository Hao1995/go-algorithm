# 787. Cheapest Flights Within K Stops

## Intuition

## Approach
Refer: https://youtu.be/5eIK3zUdYmE?si=PQ7VqgcYjx74SFui

Using `Bellman-Ford` to find the answer with time complexity `O(E*K)`.
E: Number of edges.  
K: Maximum number of edges.
T: O(N+E*K)
S: O(N)