# 90. Subsets II

## Intuition
DFS to explore any possibilities.

## Approach
// [1,2,2]
// 1, include, [1], [2,2]
//      2, include, [1,2], [2]
//          include [1,2,2], []
//              is empty, ans=[[1,2,2]]
//          not include [1,2], []
//              is empty, ans=[[1,2,2], [1,2]]
//      2, skip all 2, [1], []
//          is empty, ans=[[1,2,2], [1,2], [1]]
// 1, skip, [], [2,2]
//      2, include, [2], [2]
//          2, include, [2,2], []
//              is empty, ans=[[1,2,2], [1,2], [1], [2,2]]
//          2, skip, [2], []
//              is empty, ans=[[1,2,2], [1,2], [1], [2,2], [2]]
//      2, skip all 2, []
//          is empty, ans=[[1,2,2], [1,2], [1], [2,2], [2], []]

Time complexity: O(n*2^n)
Space complexity: O(n*2^n) // we need to return all subsets