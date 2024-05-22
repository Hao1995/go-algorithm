# 153. Find Minimum in Rotated Sorted Array

## Intuition
Binary search

## Approach
Always compare to the most right number.
If current number greater than the most right number, it means the minimum number at the right side.
If current number less than the most right number, it means the minimum number at the left side.

// nums=[4,5,1,2,3]
// l=0, r=4 (len-1)
// midd=(0+4)/2=2
// middNum=1
// middNum < mostRight (1 < 3)
//      r = m = 2
// midd=(0+2)/2=1
// middNum=5
// middNum > mostRight (5 > 3)
//      l = m + 1 = 2
// l == r >> break, return nums[2]

// nums=[1,2,3,4,5]
// l=0, r=4 (len-1)
// midd=(0+4)/2=2
// middNum=3
// middNum < mostRight (3 < 4)
//      r = m = 2
// midd=(0+2)/2=1
// middNum=1
// middNum < mostRight (2 < 4)
//      r = m = 1
// midd=(0+1)/2=0
// middNum=1
// middNum < mostRight (1 < 4)
//      r = m = 0
// l == r >> break, return nums[0]

// nums=[5,1,2,3,4]
// l=0, r=4 (len-1)
// midd=(0+4)/2=2
// middNum=2
// middNum < mostRight (2 < 4)
//      r = m = 2
// midd=(0+2)/2=1
// middNum=1
// middNum < mostRight (1 < 4)
//      r = m = 1
// l == r >> break, return nums[1]=1

// nums=[2,3,4,5,1]
// l=0, r=4 (len-1)
// midd=(0+4)/2=2
// middNum=4
// middNum > mostRight (4 > 1) >> min num is at right side
//      l = m + 1 = 3
// midd=(3+4)/2=3
// middNum=5
// middNum > mostRight (5 > 1)
//      l = m + 1 = 4
// l == r >> break, return nums[4]=1

// nums=[3,4,5,1,2]
// l=0, r=4 (len-1)
// midd=(0+4)/2=2
// middNum=5
// middNum > mostRight (5 > 2) >> min num is at right side
//      l = m + 1 = 3
// midd=(3+4)/2=3
// middNum=1
// middNum < mostRight (1 < 2) >> is valid, find left data
//      r = m = 3
// l < r >> break