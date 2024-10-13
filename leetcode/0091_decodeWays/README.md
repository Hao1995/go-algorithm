# 91. Decode Ways

## Intuitions

## Approach
### Approach 1: Dynamic Programming - Recursive
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

### Approach 2: Dynamic Programming - For loop
// s="11106", len=5
// idx=4,   len=1, 6 is valid, len=2, out of boundary, dp[4]=1, dp=[0,0,0,0,1]
// idx=3,   0 is invalid, pass
// idx=2,   len=1, 1 is valid, len=2, 10 is valid, dp[2]=dp[3]+dp[4]=1, dp=[0,0,1,0,1]
// idx=1,   len=1, 1 is valid, len=2, 11 is valid, dp[1]=dp[2]+dp[3]=1, dp=[0,1,1,0,1]
// idx=0,   len=1, 1 is valid, len=2, 11 is valid, dp[0]=dp[1]+dp[2]=2, dp=[2,1,1,0,1]
// return 2

// T: O(n)
// S: O(n)