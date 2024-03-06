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
https://www.youtube.com/live/Q6KyDl_xiIg?si=dx56uA5ouPNoIozy

// [10,9,2,5,3,7,101,18], tmp=[]
// 10, tmp=[10]
// 9, tmp=[9], 9<10, change
// 2, tmp=[2], 2<9, change
// 5, tmp=[2,5], 5>2, append
// 3, tmp=[2,3], 3<5, change
// 7, tmp=[2,3,7], 7>3, append
// 101, tmp=[2,3,7,101], 101>7, append
// 18, tmp=[2,3,7,18], 18<101, change
// return len(tmp)

The `tmp` array just like a value which store the "maximum length".

So even we have another number `8` at above example, the final answer would not change: ans = len([]int{2,3,7,8}) = 4

The "length" of `tmp` represents the final data.
Only when the next numbers completely replaces the excessively large number in `tmp` will it be possible to make the length of `tmp` longer.