# 199 Binary Tree Right Side View

# Intuition
Keep pushing right nodes to the `ans` array.
But if some cases's node without a right node, the `ans` will be wrong.
EX:
//    1
//  2   3
// n 4 5 n

# Approach
Using BFS to find each layer's nodes.
Then push the final element to the `ans` array.
