# 73. Set Matrix Zeroes

## Intuition
Iterate all elements and store the coordinates of zero elements in a list.
Iterate through the list of zero elements coordinates and set the entire row and column to zero.
However, in the worst case scenario where all the elements are zero, the time complexity would be O(n^2*m^2)

How can we optimize this?
We can't use an array to store which elements have already been marked as zero, because even if they were marked, they should be checked again.

Then I realized that I only need to know which rows and columns should be marked as zero.
So I can use a hash map to remember those rows and columns to keep the value distinct.
The time complexity would be O(n*m), and the space complexity would be O(n+m).

## Approach
1. Iterate through each elements to check if it is zero, then store the row and column indices in a hash map.
2. Iterate through the hash map of row and column indices, then set the corresponding rows and columns to zero.