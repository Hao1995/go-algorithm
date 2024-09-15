# 300. Longest Increasing Subsequence

## Intuition
Create a dp array to store the relationship between the number and the maximum subsequence length.
For each number in nums, calculate the maximum subsequence length using the values in the dp list before the current index.

Time Complexity
Best case: O(n) (ascending order)
Worse case: O(n^2) (descending order)

Space Complexity
O(n)

## Approach
### Method 1: Dynamic Programming O(n)~O(n^2)
// [0,1,0,3,2,3], n=6, dp=make([]int,7), dp[0]={0,-max}
// i=0, num=0
// j=i=0, j>=0, j--
// j=0, 0>dp[0].val(-max), dp[i+1=1]={len:0+1=1, val:0}, break
// i=1, num=1
// j=i=1, j>=0, j--
// j=1, 1>dp[1].val(0), dp[i+1=2]={len:1+1=2, val:1}, break
// i=2, num=0
// j=i=2, j>=0, j--
// j=2, 0<dp[2].val(1) pass
// j=1, 0==dp[1].val(0) pass
// j=0, 0>dp[0].val(-max), dp[i+1=3]={len:0+1=1, val:0}, break
// i=3, num=3
// j=i=3, j>=0, j--
// j=3, 3>dp[3].val(0), dp[i+1=4]={len:0+1=1, val:3}, break

// [10,9,2,5,3,7,101,18],n=8, dp=make([]int,9), dp[0]={0,-max}
// i=0, num=10
// j=i=0, j>=0, j--
// j=0, 10>dp[0].val(-max), dp[i+1=1]={dp[0].len+1, 10}
// j=-1, break
// i=1, num=9
// j=i=1, j>=0, j--
// j=1, 9<dp[1](10), pass
// j=0, 9<dp[0](-max), dp[i+1=2]={dp[0].len+1, 10}
// ...

### Method 2: Binary Search (follow up requirement O(nlogn))
https://leetcode.com/problems/longest-increasing-subsequence/solutions/74824/java-python-binary-search-o-nlogn-time-with-explanation
https://www.youtube.com/live/Q6KyDl_xiIg?si=lAd1j5RqYyp9-Fb6&t=467

Even the real longest subsequence is [1,4,5,9,10,20] not [1,5,6,9,10,20], the answer still is `6`.
Because ...
1. The current length of `tmp` represents what has been reached in the past.
2. Changing the past num in the `tmp` is to keep the flexibility of extension in the future.
3. Appending means the current num can become the largest number of any combination in the past.

// [1,8,4,6,9,2,5,10,20], tmp=[]
// 1,                   initial, tmp=[1]
// 8, 8 > 1,            append, tmp=[1,8]
// 4, 4 > 1 && 4 < 8,   change 8, tmp=[1,4]
// 6, 6 > 4,            append, tmp=[1,4,6]
// 9, 9 > 6,            append, tmp=[1,4,6,9]
// 2, 2 > 1 && 2 < 4,   change 4, tmp=[1,2,6,9], this combination not really exist, but it means further potential cases.
// 5, 5 > 2 && 5 < 6,   change 2, tmp=[1,5,6,9], there is impossible to have a combination like [...,2,5,...] because the former combination [1,4,6,9] already longer than it.
// 10, 10 > 9,          append, tmp=[1,5,6,9,10]
// 20, 20 > 10,         append, tmp=[1,5,6,9,10,20], size=6