# 662. Maximum Width of Binary Tree

## Intuition
Use BFS to calculate the width of each layer.
The problem is I need to know when the index of the start node and the end node, so I need to create another structure to record the index data.

## Approach

// layer=0,         1
// layer=1,    2          3
// layer=2,  n   5     6     7
// layer=3, 8 n n 11 12 13 14 n

// left idx = lastIdx * 2, right idx = lastIdx * 2
// layer=0, 1:0
// layer=1, 2:0, 3:1
// layer=2, n:0, 5:1, 6:2, 7:3
// layer=2, 8:0, n:1, n:2, 11:3, ... 14:6, n:7 >> startId=0, endIdx=6, ans=6-0+1
