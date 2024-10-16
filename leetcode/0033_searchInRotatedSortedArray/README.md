# 0033. Search in Rotated Sorted Array
https://leetcode.com/problems/search-in-rotated-sorted-array/description/

## Intuition
The question limits the time complexity to O(log(n)), requiring the use of Binary Search.
However, the `nums` array might have been rotated, so it is essential to determine the pivot index first.

## Approach
1. Check if the `nums` array has been rotated.
2. If it is, find the original start number.
3. Find the target on different side of the `nums`
