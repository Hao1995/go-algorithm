# 74. Search a 2D Matrix

## Intuition
Each row of the matrix is in non-decreasing order, and the last value in each row is less than the first value of the next row.
Furthermore, the question requires the time complexity should be O(log(m*n)). So my first though is binary search.

## Approach
Imaging the matrix is a 1 dimensional array, we just need to concatenate each row.
Then calculate the index by `r` and `c`.

// matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// m is the number of rows = len(matrix) = 3
// n is the number of columns = len(matrix[0]) = 4
// vIdx=0*n+0=0, r=vIdx/4=0, c=vIdx%4=0, (0,0)
// vIdx=0*n+1=1, r=vIdx/4=0, c=vIdx%4=1, (0,1)
// vIdx=0*n+2=2, r=vIdx/4=0, c=vIdx%4=2, (0,2)
// vIdx=0*n+3=3, r=vIdx/4=0, c=vIdx%4=3, (0,3)
// vIdx=1*n+0=4, r=vIdx/4=1, c=vIdx%4=0, (1,0)
// vIdx=1*n+1=5, r=vIdx/4=1, c=vIdx%4=1, (1,1)
// ... 