# 1905. Count Sub Islands

## Intuition
1. Find the points on the same y axis.
2. According to the previous x points, find the corresponding diagonal points with same length.
3. Find the points on the same x axis.
4. Do multiplication to get the number of combination.

## Approach
My intuition is a little inefficient because we need to check the different corner of a square step by step.
Ref: https://youtu.be/bahebearrDc?si=r-O2UYo0rvJW9FBP

According to the NeetCode's solution, we just need to iterate all points to find the diagonal points first, which cost O(n).
