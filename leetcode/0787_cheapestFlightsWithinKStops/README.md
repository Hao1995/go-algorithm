# 787. Cheapest Flights Within K Stops

## Intuition
Convert flights to a HashMap, then we can find the relationship between all the cities.
Start travel from `src`, continue to find the next destination until arrive the `dest`.
>> Time Limit Exceeded.

Time complexity(worst case): O(E^K)
Because it will go through all possible paths.

## Approach
Refer: [Bellman-Ford - Cheapest Flights within K Stops - Leetcode 787 - Python](https://youtu.be/5eIK3zUdYmE?si=1YOel_zrcwMbqDsB)
Using `Bellman-Ford` to find the answer with time complexity `O(E*K)`. It's only go through the path with less price.
E: Number of edges.  
K: Maximum number of edges.

E = n*(n-1)/2 = O(n^2)  
K < n  
O(E\*K) = O (n^2\*n) = O(n^3)