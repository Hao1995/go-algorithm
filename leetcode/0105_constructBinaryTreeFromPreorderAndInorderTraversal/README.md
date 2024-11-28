# 105. Construct Binary Tree from Preorder and Inorder Traversal

## Intuition
In preorder traversal, the root node is printed first, followed by the left node and then the right node.
In inorder traversal, the left node is printed first, and after traversing back, the root node is printed, follow by the right node.

Thus, we can determine the current root node by selecting the first element in the preorder array.
Next, we identify the left subtree by finding the root node position in the inorder array.
We then recursively pass the preorder and the inorder array of the left subtree to the DFS and do the same for the right subtree. Eventually, this constructs the original binary tree.

## Approach