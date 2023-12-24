# 56. Merge Intervals
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

Example 1:
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

Example 2:
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.

## Intuition
This problem is not like a medium problem, it is nearly simple.
Just for loop all interval then check, merge them.

## Approach
1. Sorted the `intervals` by start
2. For each
3. Check if it's overlapping then merge them.
4. Update result to `ans` list.
5. Return `ans`