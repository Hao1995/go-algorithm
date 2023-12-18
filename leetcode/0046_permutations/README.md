# 46. Permutations

Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]
Example 3:

Input: nums = [1]
Output: [[1]]
 

Constraints:

1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.

## Intuition
We need an algorithm that can list all possible combinations, e.g., using depth-first search (DFS) or breadth-first search (BFS).

Then, I observed a situation:

Case: [1,2,3,4]

When I start from 1, the next one would be 2, and so on until reaching 4.
```
[1,2,3,4]
```

Then, I will backtrack to change 3 to 4 and continue filling the remaining number,
```
[1,2,4,3]
```

And so on. If there are no remaining numbers, continue backward to change the numbers.
```
[1,3,2,4],[1,3,4,2]
[1,4,2,3],[1,4,3,2]
[2,1,3,4],[2,1,4,3]
[2,3,1,4],[2,3,4,1]
...
```

## Approach
1. Use DFS.
2. Keep track of all existing numbers, and if a number already exists, skip it.
3. If the length of the list is the same as the length of the input, add it to the answer list.
