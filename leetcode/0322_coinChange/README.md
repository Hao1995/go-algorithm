# 322 CoinChange

You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.

## Approach
It's a DP problem. We calculate the fewest number of coins that we need, then update to the cache array.
Next time we're going to find the remaining coin within the cache array.

// Input
// [1,2,5], 11
//  00,01,02,03,04,05,06,07,08,09,10,11
// [ 0, n, n, n, n, n, n, n, n, n, n, n]

// round 01
// 01-1=0, dp[0]=0, dp[1]=dp[0]+1=1
// 01-2 < 0, pass
// [ 0, 1, n, n, n, n, n, n, n, n, n, n]
// round 02
// 02-1=1, dp[1]=1, dp[2]=dp[1]+1=2
// 02-2=0, dp[0]=0, dp[2]=dp[0]+1=1
// 02-5<0, pass
// [ 0, 1, 1, n, n, n, n, n, n, n, n, n]
// round 03
// 03-1=2, dp[2]=1, dp[3]=dp[2]+1=3
// 03-2=1, dp[1]=1, dp[3]=dp[1]+1=2
// 03-5<0, pass
// [ 0, 1, 1, 2, n, n, n, n, n, n, n, n]
// round 04
// 04-1=3, dp[3]=2, dp[4]=dp[3]+1=3
// 04-2=2, dp[2]=1, dp[4]=dp[2]+1=2
// 04-5<0, pass
// [ 0, 1, 1, 2, 2, n, n, n, n, n, n, n]
// round 05
// 05-1=4, dp[4]=2, dp[5]=dp[4]+1=3
// ...
// 05-5=0, dp[0]=0, dp[5]=dp[0]+1=1
// [ 0, 1, 1, 2, 2, 1, n, n, n, n, n, n]
// ...
// round 11
// 11-1=10, dp[10]=2, dp[11]=dp[10]+1=3
// ...
// 11-5=6, dp[6]=2, dp[11]=dp[6]+1=3
// [ 0, 1, 1, 2, 2, 1, 2, 2, 3, 3, 2, 3]
