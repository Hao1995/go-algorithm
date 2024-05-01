# 658. Find K Closest Elements

## Intuition
Binary search.
Use binary search to find the value is most closet to `x`, then use two pointer to find the boundary.
The problem is if we can't find the `x` within the array, or if we find the x in the array but it is located at the most right side and without enough k numbers.
EX:
arr=[1,2,3,4,5], k=2, x=-1 >> We can't find `-1` in the array
arr=[1,2,3,4,5], k=3, x=4 >> We can't find enough k elements since start from `4`.

## Approach
Ref: NeetCode https://youtu.be/o-YDQzHoaKM?si=qpG50PULypewUNEL

We don't need to find the `x` within the array, we must find the most left element of the window from the `arr`.

Since we extend from the most closet number of `x` to both left and right.
Example: arr=[1,2,3,4,5], k=3, x=3
start from the most closet value of `3` from array >> [3]
extend to left and right side >> [2,3,4]
So we should find the `2` value using binary search.

1. Keep enough window: right boundary should be `(len(arr) - 1) - (k - 1) = len(arr) - k`
2. The left and right number from the middle number must have a similar difference, x=2, x-2=2 and 4-x=2
3. If left and right numbers have same difference, chose smaller index, index 2 is better choice than index 4.
4. Suppose the middle value if the start number of this window, it should has smaller difference than the index `start + k`
5. Then we can find the valid numbers.