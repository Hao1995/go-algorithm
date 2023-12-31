# 416. Partition Equal Subset Sum

Given an integer array nums, return true if you can partition the array into two subsets such that the sum of the elements in both subsets is equal or false otherwise.

Example 1:
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].

Example 2:
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.

## Intuition
First, sort the array in descending order.
Declare a variable named `sum` and initialize it with the value at the rightmost position. Start subtracting values from the second position from the right.
If the result is greater than 0, continue subtracting; otherwise, add the value starting from the left side.

However, this approach does not always work since the `sum` might be composed of more than one number in certain cases:
```
[3,3,3,4,5] = [3,3,3] + [4,5]
```

## Approach
// [1,2,3,8], sum=14, sum/2=7
// dp[0][0]=true
// dp[1][0]=dp[0][0]=true
// dp[1][1]=dp[0][0]=true
// dp[1][2]=false
// ...
// dp[1][7]=false
// dp[2][0]=dp[1][0]=true
// dp[2][1]=dp[1][1]=true
// dp[2][2]=dp[1][0]=true
// dp[2][3]=dp[1][1]=true
// dp[2][4]=dp[1][4]||dp[1][2]=false
// ...
// dp[2][7]=false
// dp[3][0]=dp[2][0]=true
// dp[3][1]=dp[2][1]=true
// dp[3][2]=dp[2][2]=true
// dp[3][3]=dp[2][3]=true
// dp[3][4]=dp[2][1]=true
// dp[3][5]=dp[2][2]=true
// dp[3][6]=dp[2][3]=true
// dp[3][7]=dp[2][7]||dp[2][4]=false
// dp[8][0]=dp[3][0]=true
// dp[8][1]=dp[3][1]=true
// dp[8][2]=dp[3][2]=true
// dp[8][3]=dp[3][3]=true
// dp[8][4]=dp[3][4]=true
// dp[8][5]=dp[3][5]=true
// dp[8][6]=dp[3][6]=true
// dp[8][7]=dp[3][7] || 8>=7=false