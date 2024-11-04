# 338. Counting Bits

## Intuition
Brute force. Iterate each number and calculate the number of 1. >> O(nlogn)

## Approach
Ref: https://youtu.be/RyBM56RIWrM?si=rxotv2Esgsvp0keb

T: O(n), S: O(n)
We can observe a pattern from the bits value of each number.
// offset=1, dp=[5]
// 1, offset*2=2 not equal 1, dp[1]=1+dp[1-1]=1+dp[0]=1, dp=[0,1,0,0,0],0
// 2, offset*2=2 equal to 2, offset=2, dp[2]=1+dp[2-2]=1+dp[0]=1, dp=[0,1,1,0,0,0]
// 3, offset*2=4 not equal to 3, dp[3]=1+dp[3-2]=1+dp[1]=2, dp=[0,1,1,2,0,0]
// 4, offset*2=4 equal to 4, offset=4, dp[4]=1+dp[4-4]=1, dp=[0,1,1,2,1,0]
// 5, offset*2=8 not equal to 5, dp[5]=1+dp[5-4]=1, dp=[0,1,1,2,1,2]