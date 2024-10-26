# 213. 213. House Robber II

## Intuition
DP. But I was struggled by adjacent problem when I check the final number.

## Approach
Ref: https://youtu.be/rWAJCfYYOvM?si=kqlq2_vyMPVsgz60

// 0,1,2,3,4,5
// 1,2,3,1,2,2, x=0, o=0
//
// x is the one before the previos number
// o is the max value of previous combination
// c
// 0,1,2,3,4,5
// 1,2,3,1,2,2
// idx=0, house 1, localMaxRub=max(x+1,o)=(1,0)=1, x=0, o=1
//
//   c
// 0,1,2,3,4,5
// 1,2,3,1,2,2
// idx=1, house 2, localMaxRub=max(0+2,1)=2, x=1, o=2
//
//     c
// 0,1,2,3,4,5
// 1,2,3,1,2,2
// idx=2, house 3, localMaxRub=max(1+3,2)=4, x=2, o=4
//
//       c
// 0,1,2,3,4,5
// 1,2,3,1,2,2
// idx=3, house 1, localMaxRub=max(2+1,4)=4, x=4, o=4
//
//         c
// 0,1,2,3,4,5
// 1,2,3,1,2,2
// idx=4, house 2, localMaxRub=max(2+4,4)=6, x=4, o=6
//
//           c
// 0,1,2,3,4,5
// 1,2,3,1,2,2
// idx=5, house 2, localMaxRub=max(4+2,6)=6, x=6, o=6

// for filter the edge cases
// avoid the adjacent, remove the last element, because we won't not if the curr+x would include the first element, nums=[:-1] >> [1,2,3,1,2,_]
// then you mighe miss the possible combination like indexes 1,3,5, so we could try different combination with3 nums[1:] >> [_,2,3,1,2,2]
// if the length of nums is 1, return it directly.