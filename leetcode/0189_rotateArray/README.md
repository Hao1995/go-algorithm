# 189. Rotate Array

## Intuition
Input like: nums=[1,2,3,4,5,6,7], k=3.

Split `nums` into 2 parts:
The first part is [5,6,7], which will become the leftPart of the new nums.
The second part is [1,2,3,4], which will become the rightPart of the new nums.

## Approach
### My Approach
1. Calculate the leftPart and rightPart for the new `nums`.
2. Write the leftPart and rightPart back to `nums`.

Time Complexity: O(n)
Space Complexity: O(n)

### danny6514's Approach
https://leetcode.com/problems/rotate-array/solutions/54250/easy-to-read-java-solution
1. After rotation, the final k elements will move to the front. So we first reverse the whole `nums`.
2. Reverse the first k elements to maintain the correct order.
3. Reverse the remaining elements to maintain the correct order.