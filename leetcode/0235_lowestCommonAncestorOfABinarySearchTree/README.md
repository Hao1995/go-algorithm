# 235. Lowest Common Ancestor of a Binary Search Tree

## Intuition
The first though comes in my mind is either DFS or BFS, because the data structure is the BST.
The question wants to find the lowest "common" ancestor, so it must be DFS. Let's continue and see what the processing flow looks like.

## Approach
### DFS
Continue to travel through all nodes and put the ancestor into the corresponding slice.
For example, before I found the `p` node, I will push all nodes into p's ancestor array, but if `p` can't be find at any subtree, those nodes would be remove directly.
Finally, we can find the lowest common ancestor from p and q ancestor array.

Time complexity: O(n)
Space complexity: O(n) (worst case)

### O(1) Space Solution
The key to this solution is the BST; all nodes on the left are smaller than the current node, and vice versa.
There are 3 cases:
1. P and Q are on the left and right side of the current node, respectively, then `(root-q) * (root-p)` must be negative, directly return current node.
2. One of P and Q is the LCA, so `(root-q) * (root-p)` must be 0, return current node directly.

Time Complexity: O(n)
Space Complexity: O(1)