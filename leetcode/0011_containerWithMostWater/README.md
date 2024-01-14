# 11. Container With Most Water

## Intuition
I attempted a binary search, but the problem is that the data is not sorted, so I wasn't sure which boundary to set.

## Approach
Just start from the most left side and end to most right side to calculate the amount of water.
If the left idx is less than the right idx, that mean the max height is restricted by the left idx.
So we gonna to increment the left idx to get the higher max height, and vice versa.