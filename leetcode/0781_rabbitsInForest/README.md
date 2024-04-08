# 781. Rabbits In Forest

## Intuition
Choose `Greedy` topic from wisdompeak.
We don't need to know which rabbits are in the same group.
We just need to know which rabbits `possibly` have the same group.

EX: answers=[1,1,2,1,1]
The 1th, 2nd, 4th, and 5th rabbits said there is 1 rabbit has same color as them.
So we have the following combinations:
solution1 = (1th, 2nd), (4th, 5th)
solution1 = (1th, 4th), (2nd, 5th)
solution2 = (1th, 5th), (2nd, 4th)

After grouping, the result would be [1,2,1], then summing them up with themselves = (1+1) + (2+1) + (1+1) = 7

## Approach
1. Add to countMap if the `ans` not exist, then add `ans + 1` to the result
2. If `ans` is present, decrement it by one.