# 543. Diameter of Binary Tree

# Intuition
DFS.
The longest diameter must is the longest left path + the longest right path.

# Approach
// dfs
// get the length from left
// get the length from right
// check if leftLength + rightLength >= ans, update it
// return the current longest length back to their parents
