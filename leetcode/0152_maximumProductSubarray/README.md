# 152. Maximum Product Subarray

## Intuition
Must consider positive and negative situations.
Remember the minimum value and maximum value of each step to calculate the next min, max values.

## Approach
// [2,3,-2,-4]
// 2, minVal=2, maxVal=2, ans=2
// 3, minVal=3, maxVal=6, ans=6
// -2, minVal=-12, maxVal=-6, ans=6
// -4, minVal=24, maxVal=48, ans=48
// maxAns=48

T: O(n)
S: O(1)