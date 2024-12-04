# 371. Sum of Two Integers

## Intuition
Bit manipulation.

## Approach
Ref: https://youtu.be/gVUrDV4tZfY?si=L1DiEHSE0FI5ZrBr

// a=1001(9), b=1011(11)
//  round 1
//      a^b     = 0010
//      (a&b)<<1=10010
//      -
//      a= 0010
//      b=10010
//  round 2
//      a^b     =10000
//      (a&b)<<1=00100
//      -
//      a=10000
//      b=00100
//  round 3
//      a^b     =10100
//      (a&b)<<1=00000
//      -
//      a=10100
//      b=00000
//  return 10100(16+4=20)

// T: Usually O(1)
// S: O(1)