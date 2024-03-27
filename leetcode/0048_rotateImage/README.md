# 48. Rotate Image

## Intuition
Start exchanging value from the four outermost corners of the square, then rotate clockwise, exchanging one by one.
After exchanging, go inward and repeat the same action.

Time Complexity: O(n)