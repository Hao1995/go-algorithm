# 84. Largest Rectangle in Histogram

## Intuition

## Approach
Ref: https://youtu.be/zx5Sw9130L0?si=sDi1ixq4tLJSwHsp

Three cases:
1. increasing, heights=[1,2,3,4,5]
2. same, heights=[2,2,2,2,2]
3. decreasing, heights=[5,4,3,2,1]

Steps:
1. Iterate `heights` to calculate the start index.
2. If the bar starts to decrease, calculate the previous area and simultaneously update the start index of current bar.
3. Calculate the area for the remaining increasing bars.

Simulation:
// heights=[2,1,5,6,2,3]
//       *
//     * *    
//     * *    
//     * *   *
// *   * * * *
// * * * * * *
// 0 1 2 3 4 5

// Calculate the area of decreasing bars, and push increasing bars into the stack.
// 0, currH=2, stack is empty, append into stack=[(0,2)]
// 1, currH=1 is lower than prevH=2, the prev area=height*width=2*1=2, maxArea=2, append into stack=[(1,1)]
// 2, currH=5 is higher than prevH=1, append into stack=[(1,1), (2,5)]
// 3, currH=6 is higher than prevH=5, append into stack=[(1,1), (2,5), (3,6)]
// 4, currH=2 is lower than prevH=6, the prev area=6*1=6, maxArea=6, stack=[(1,1), (2,5)]
//                                   the prev area=5*2=10, maxArea=10, stack=[(1,1)]
//    currH=2 is higher than prevH=1, append into stack=[(1,1), (2,2)]
// 3, currH=3 is higher than prevH=1, append into stack=[(1,1), (2,2), (5,3)]

// The remaining items can be extended infinitely to the right, because the stack is increasing order.
// We can use the size of the original input as width to calculate their areas.
// Iterate stack = [(1,1), (2,2), (5,3)], len=6
// (1,1). area=(6-1)*1=6, maxArea still same = 10
// (2,2). area=(6-2)*2=8, maxArea still same = 10
// (5,3). area=(6-5)*3=3, maxArea still same = 10

// return 10
