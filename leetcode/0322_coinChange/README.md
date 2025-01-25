# 322. CoinChange

## Intuition
DP

## Approach
// [1,3,4,5], amount=7
// dp[0]=0
//
// 1, dp[1]=1+dp[1-1]=1
// 3, too big, pass
// 4, too big, pass
// 5, too big, pass
//
// 1, dp[2]=1+dp[2-1]=2
// 3, too big, pass
// 4, too big, pass
// 5, too big, pass
//
// 1, dp[3]=1+dp[3-1]=1+2=3
// 3, dp[3]=min(dp[3], 1+dp[3-3])=min(3, 1+0)=1
// 4, too big, pass
// 5, too big, pass
//
// 1, dp[4]=min(dp[4], 1+dp[4-1])=min(~, 2)=2
// 3, dp[4]=min(dp[4], 1+dp[4-3])=min(2, 1+1)=2
// 4, dp[4]=min(dp[4], 1+dp[4-4])=min(2, 1+0)=1
// 5, too big, pass
//
// 1, dp[5]=min(dp[5], 1+dp[5-1])=min(~, 1+1)=2
// 3, dp[5]=min(dp[5], 1+dp[5-3])=min(2, 1+2)=2
// 4, dp[5]=min(dp[5], 1+dp[5-4])=min(2, 1+1)=2
// 5, dp[5]=min(dp[5], 1+dp[5-5])=min(2, 1+0)=1
//
// 1, dp[6]=min(dp[6], 1+dp[6-1])=min(~, 1+1)=2
// 3, dp[6]=min(dp[6], 1+dp[6-3])=min(2, 1+1)=2
// 4, dp[6]=min(dp[6], 1+dp[6-4])=min(2, 1+1)=2
// 5, dp[6]=min(dp[6], 1+dp[6-5])=min(2, 1+1)=2
//
// 1, dp[7]=min(dp[7], 1+dp[7-1])=min(~, 1+2)=3
// 3, dp[7]=min(dp[7], 1+dp[7-3])=min(3, 1+1)=2
// 4, dp[7]=min(dp[7], 1+dp[7-4])=min(2, 1+1)=2
// 5, dp[7]=min(dp[7], 1+dp[7-5])=min(2, 2+1)=2
//
// return dp[7]=2