# 494. Target Sum

## Intuition
Try different symbols until the end of nums, and count the valid combinations.

## Approach
// nums=[1,1,1,1,1], target=3
// +1, return 3
//      +1, return 2
//          +1, return 1
//              +1, return 0
//                  +1=5, return 0
//                  -1=4, return 0
//              -1, return 1
//                  +1=4, return 0
//                  -1=3, return 1
//          -1, return 1
//              +1, return 1
//                  +1=3, return 1
//                  -1=2, return 0
//              -1, return 0
//                  +1=2, return 0
//                  -1=1, return 0
//      -1, return 1
//          +1, return 1
//              +1, return 1
//                  +1=3, return 1
//                  -1=2, return 0
//              -1, return 0
//                  +1=2, return 0
//                  -1=1, return 0
//          -1, return 0
//              +1, return 0
//                  +1=1, return 0
//                  -1=-1, return 0
//              -1, return 0
//                  +1=-1, return 0
//                  -1=-2, return 0
// -1, ...