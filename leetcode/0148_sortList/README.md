# 148. Sort List

## Intuition
No.

## Approach
[Sort List - Merge Sort - Leetcode 148](https://youtu.be/TGveA1oFhrc?si=mENShzcM4CwkiYPz)

// time complexity = O(nlogn), space complexity = O(logn)
// merge sort
// -1,5,3,4,0
// left=-1,5,3, right=4,0
//      left=-1,5, right=3
//          left=-1, right=5
//          after merge=-1,5
//      after merge=-1,3,5
//      left=4, right=0
//      after merge=0,4
//  after merge=-1,0,3,4,5
