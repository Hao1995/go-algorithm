# 235. Lowest Common Ancestor of a Binary Search Tree

## Intuition
The first though comes in my mind is either DFS or BFS, because the data structure is the BST.
The question wants to find the lowest "common" ancestor, so it must be DFS. Let's continue and see what the processing flow looks like.

## Approach
Continue to travel through all nodes and put the ancestor into the corresponding slice.
For example, before I found the `p` node, I will push all nodes into p's ancestor array, but if `p` can't be find at any subtree, those nodes would be remove directly.
Finally, we can find the lowest common ancestor from p and q ancestor array.

# Complexity
Time complexity: O(n)
Space complexity: O(n) (worst case)