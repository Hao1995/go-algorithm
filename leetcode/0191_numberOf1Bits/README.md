# 191. Number of 1 Bits

## Intuition

## Approach
1. Keep dividing by two until the quotient is less than 2, and count how many times the remainder is 1 during the process.
Worst case is O(32) >> O(1)
// 11(1101)
// 11/2=5...1, count=1
// 5/2=2...1, count=2
// 2/2=1...0, count=3

2. Use right shift (>>) and and (&&) operators to count if the current bit is set.
Worst case is O(32) >> O(1)
// 11=1011
// 1101 & 1 = 1, count=1, n=110
// 110 & 1 = 0, count=1, n=11
// 11 & 1 = 1, count=2, n=1
// 1 & 1 = 1, count=3, n=0
