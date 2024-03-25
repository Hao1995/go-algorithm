# 221. Maximal Square

## Intuition
My Intuition is DFS or BFS, because the question is related to graphics.
I will for each each position then calculate the area.
So if all elements in the matrix are '1', the time complexity would be O(m^2*n^2).
Sure enough, after the calculation, a warning of "time limit exceeded" popped up.

I can't imagine how to cache the data I've walked through before.

## Approach
Solution by Huifeng Guan: "Dynamic Programming https://www.youtube.com/live/eg6g6cNvsTs?si=YZqMovs0xrfke9MO"

1. Initial first row and first column.
2. Calculate next position by left, upper, left and upper position.

For example:
// [0, 1, 1]
// [1, 1, 1]
// [0, 1, 1]

1. Initial DP's first row and column.
// [0, 1, 1]
// [1, x, x]
// [0, x, x]

2. Calculate the next row by left, upper, left and upper position
// [0, 1, 1]
// [1, 1, 2]
// [0, x, x]

2. Continue to calculate the next row...
// [0, 1, 1]
// [1, 1, 2]
// [0, 1, 2]
