# 51. N-Queens

## Intuition

## Approach
Ref: https://youtu.be/Ph95IHmRp5M?si=D9opJumxuM6s3wxn

The difficult part for me is "I don't know how the queen move", therefore I refer the explanation from the NeetCode.

1. Each row at most has one queen because the entire row and col can not exist another queen.
2. Use `col`, `postSet`, and `negSet` to records the prohibited locations.
3. Continue try different row and col.

Time Complexity: O(n^2)

Illustration.

## Correct Case
// Step 1. Try first one, mark entire row, col, positive diagonal path and negative diagonal path.
// Q x x x
// x x - -
// x - x -
// x - - x

// Step 2. Try second row, the mark.
// Q x x x
// x x Q x
// x x x x
// x - x x

// Step 3. Try third row, but can not find a position, pass.
// Q x x x
// x x Q x
// x x x x
// x - x x

// Step 4. Back to second row, choose col 3
// Q x x x
// x x x Q
// x - x x
// x x - x

// Step 5. Try third row, then mark.
// Q x x x
// x x x Q
// x Q x x
// x x x x

// Step 6. Try fourth row, but can not find a position, pass.
// Q x x x
// x x x Q
// x Q x x
// x x x x

// Step 7. Back to first row, choose col 1
// x Q x x
// x x x -
// - x - x
// - x - -

// Step 8. Try second row, then mark
// x Q x x
// x x x Q
// - x x x
// - x - x

// Step 8. Try third row, then mark
// x Q x x
// x x x Q
// Q x x x
// x x - x

// Step 9. Try fourth row, then mark
// x Q x x
// x x x Q
// Q x x x
// x x Q x

// Step 10. append to result

// Step 11. Back to previous row and continue to try until end.