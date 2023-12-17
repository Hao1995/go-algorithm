# 0994 Rotting Oranges
https://leetcode.com/problems/rotting-oranges/

You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

## Intuition
Using BFS to rot the nearby oranges, then we can determine the time required for all oranges to be rotten.

## Approach
1. BFS: Start rotting from all rotted oranges, add 1 minute when they rotted the nearby oranges each time.
2. Visited Grid: Pass insert the new rotted orange to the nextQueue if it already inserted.
3. Finally, check if there has any orange has not been rotted.

1. **BFS**: Start rotting from all rotten oranges, adding 1 minute each time they rot nearby oranges.
2. **Visited Grid**: Insert the newly rotten orange into the nextQueue only if it has not been inserted before.
3. Finally, check if there are any oranges that have not yet rotted.