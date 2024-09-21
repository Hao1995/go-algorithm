# 153. Find Minimum in Rotated Sorted Array

## Intuition
Binary search

## Approach
// Check if the array is rotated
// nums=[3,4,5,1,2]
// if it's a rotated array
//      start binary search
//          if nums[midd] > nums[-1]
//              left = midd + 1
//          else
//              right = midd