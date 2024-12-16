# 684. Redundant Connection

## Intuition
Undirected graph?

## Approach
Ref: NeetCode, https://youtu.be/FXWRE67PLL0?si=g3Debl1moX2Q5TP9
Algorithm: Union Find

n: the number of edges
m: the number of vertices

T: O(n)
S: O(m)

// Case1
// [[1,2],[1,3],[2,3]]
// parents=[1,2,3]
// rank=[1,1,1]
// [1,2], parents[1,1,3], rank=[2,1,1]
// [1,3], parents[1,1,1], rank=[3,1,1]
// [2,3], parents[1,1,1], since parents[2] == parents[3], it is formed a cycle, return it

// Case2
// [[1,2],[2,3],[3,4],[1,4],[1,5]]
// parents=[1,2,3,4,5]
// ranks=[1,1,1,1,1]
// [1,2], parents=[1,1,3,4,5], ranks=[2,1,1,1,1]
// [2,3], parents=[1,1,1,4,5], ranks=[3,1,1,1,1]
// [3,4], parents=[1,1,1,1,5], ranks=[4,1,1,1,1]
// [1,4], parents=[1,1,1,1,5], since parents[1] == parents[4], return [1,4]

// Case3
// [[1,2],[2,3],[4,5],[5,6],[1,6],[3,4]]
// parents=[1,2,3,4,5,6]
// ranks=[1,1,1,1,1,1]
// [1,2], parents=[1,1,3,4,5,6], ranks=[2,1,1,1,1,1]
// [2,3], parents=[1,1,1,4,5,6], ranks=[3,1,1,1,1,1]
// [4,5], parents=[1,1,1,4,4,6], ranks=[3,1,1,2,1,1]
// [5,6], parents=[1,1,1,4,4,4], ranks=[3,1,1,3,1,1]
// [1,6], parents=[1,1,1,1,4,4], ranks=[7,1,1,3,1,1]
// [3,4], parents=[1,1,1,1,4,4], since parents[3] == parents[4], return [3,4]

//    - - - - - - - - - - -
//  /                      \
// 1 -- 2 -- 3 -- 4 -- 5 -- 6
