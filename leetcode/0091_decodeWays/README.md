# 91. Decode Ways

## Intuitions
DFS, but it exceed the time limit.

## Approach
Dynamic Programming.
Refer from: https://www.youtube.com/watch?v=6aEyTjOwlJU
Start from idx 0, then keep search to the final one.

Explanation:
// input="121"
// start from 0
//      start from 1
//          start from 2
//              start from 3
//              exceed length of input, return 1
//          cache[2]=1
//          try next 2, it is exceeded.
//      cache[1]=1
//      try next 2, it's work.
//      cache[1]+=cache[3]=1, cache[1]=2
//  cache[0]=dp[1]=2
//  try next 2, it's work.
//  cache[0]+=cache[2]=1, cache[0]=2+1=3
//  cache=3

// 1
// 1,2
// 1,2,1, count+=1=1
// 1,21, count+=1=2
// 12
// 12,1, count+=1=3