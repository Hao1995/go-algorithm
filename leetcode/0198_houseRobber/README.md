# 198. House Robber

## Intuition
No idea.

## Approach
### DP with O(1) Space Complexity
Follow https://leetcode.com/problems/house-robber/solutions/55693/c-1ms-o-1-space-very-simple-solution

It's easy to understand.
Simply iterate through the numbers and calculate the sum with a one-night interval.

The edge case is when 'the interval of the maximum robbed amount is more than one night.'
In this case, the program will maintain the current maximum robbed amount until obtaining a higher one.

EX:
// [2,7,3,1,9] >> [7,9]
// 0=2, even, a=2, b=0
// 1=7, odd, b=7, a=2
// 2=3, even, a=5>>7, b=7
// 3=1, odd, b=8, a=7
// 4=9, even, a=16, b=8
// return 16

### Pure DP
// nums=[1,2,3,1]
// dp=[0,0,0,0]
// i=0, num=1, dp[0]=num=1, dp=[1,0,0,0]
// i=1, num=2, dp[1]=max(dp[0], num)=2, dp=[1,2,0,0]
// i=2, num=3, dp[2]=max(dp[1], num + dp[0])=max(2,3+1)=4, dp=[1,2,4,0]
// i=3, num=1, dp[3]=max(dp[2], num + dp[1])=max(4,1+2)=4, dp=[1,2,4,4]
// return 4