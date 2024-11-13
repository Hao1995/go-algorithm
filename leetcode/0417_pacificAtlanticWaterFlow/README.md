# 417. Pacific Atlantic Water Flow

## Intuition
Use DFS to iterate each cell and a DP to cache the result.
The problem is that if a new cell can flow into the adjacent cells but they happen to be cached, then the data will not be updated.

## Approach
Ref: https://youtu.be/s-VkcjHqkGI?si=eJem_vUTQdkf6vzt
T: O(n*m)
S: O(n*m)
