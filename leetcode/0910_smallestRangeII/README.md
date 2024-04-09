# 910. Smallest Range II

## Intuition
We choose the `Greedy` topic from wisdompeak.

After sorting `nums` with ascending order.
The number on the far right is expected to be the largest.
So the number on the far right will only be subtracted by k, not added by k.

Similarity, the number on the far left is the smallest one, it should always added by k, not subtracted by k.


## Approach
// Case: k is less than the biggest number
// nums=[1,2,3,4,5], k=2
// 1, maxVal=max(1+2=3, 5-2=3)=3, minVal=(1+2=3, 2-2=0)=0, diff=3
// 2, maxVal=max(2+2=4, 3)=4, minVal=(3, 3-2=1)=1, diff=2
// 3, maxVal=max(3+2=5, 3)=5, minVal=(3, 4-2=2)=2, diff=3
// 4, maxVal=max(4+2=6, 3)=6, minVal=(3, 5-2=3)=3, diff=3

// Case: k is greater than the biggest number >> The largest number will become on the left side, not the right side.
// nums=[1,2,3,4,5], k=10
// 1, maxVal=max(1+10=11, 5-10=-5)=11, minVal=(1+10=11, 2-10=-8)=-8, diff=19
// 2, maxVal=max(2+10=12, -5)=12, minVal=(11, 3-10=-7)=-7, diff=19
// 3, maxVal=max(3+10=13, -5)=13, minVal=(11, 4-10=-6)=-6, diff=19
// 4, maxVal=max(4+10=14, -5)=14, minVal=(11, 5-10=-5)=-5, diff=19