# 40. Combination Sum II

## Intuition
Backtracking, try different combination and append to answer list.

## Approach
But I failed to remove the duplicated sets at the first time. I got solution from: https://youtu.be/FOyRpNUSFeA?si=oOqTBYAOB4_tXQ__
Just simply image that it's a binary tree, include current element or not.

// [2,5,2,1,2], target=5
// sort=[1,2,2,2,5]
// 1, sum=1, 1 < target, comb=[1]
//      2, sum=3, 3 < target, comb=[1,2]
//          2, sum=5, 5 == target, comb=[1,2,2], combs=[[1,2,2]], return
//          5, sum=8, 8 > target, return
//      2, same, pass
//      2, same, pass
//      5, sum=6, 6 > target, return
// 2, sum=2, 2 < target, comb=[2]
//      2, sum=4, 4 < target, comb=[2,2]
//          2, sum=6, 6 > target, return
//          5, sum=9, 9 > target, return
//      5, sum=7, 7 > target, return
// 5, sum=5, 5 == target, comb=[5], combs=[[1,2,2],[5]]