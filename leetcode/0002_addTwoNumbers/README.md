# 2. Add Two Numbers

## Intuition
Just keep finding next node from `l1` and `l2`.
Be careful the overflow problem and the different node's length problem.

## Approach
Simulation process:
//  l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//  9+9=18,[8,1]
//  1+9+9=19,[8,9,1]
//  1+9+9=19,[8,9,9,1]
//  1+9+9=19,[8,9,9,9,1]
//  1+9=10,[8,9,9,9,0,1]
//  1+9=10,[8,9,9,9,0,0,1]
//  1+9=10,[8,9,9,9,0,0,0,1]