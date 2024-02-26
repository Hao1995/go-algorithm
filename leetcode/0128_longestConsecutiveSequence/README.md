# 128. Longest Consecutive Sequence

## Intuition
Store each num into a boolean array to represents if the num was exist. 

// [100,4,200,1,3,2], markArr=make([]false, 2 * math.Pow(10,9))
// 100, markArr[100] = true
// 4, markArr[4] = true
// 200, markArr[200] = true
// ...
// markArry = [false, true, true, true, true, false, ..., true, ..., true]

// foreach to calculate the max length of the markArr
// prevIsExist=false, tmpLen=0, maxLen=0
// 1, isExist=true, prevIsExist=, tmpLen=1, maxLen=1
// 2, isExist=true, prevIsExist=true, tmpLen=2, maxLen=2
// 3, isExist=true, prevIsExist=true, tmpLen=3, maxLen=3
// 4, isExist=true, prevIsExist=true, tmpLen=4, maxLen=4
// return 4

But I got an error `cannot allocate memory`.

I must find a solution that doesn't require so much memory.


## Approach
Refer to `https://leetcode.com/problems/longest-consecutive-sequence/solutions/41057/simple-o-n-with-explanation-just-walk-each-streak`

Just insert all numbers to a hash map.
Find the minimum number of each number group.
Then start counting the maximum consecutive numbers.

// [100,4,200,1,3,2]
// map >> [100:true,4:true,200:true,1:true,3:true,2:true], max len = 0

// 100, find 99 >> not found! The `100` is the min number of this group.
// find 101 >> not found >> tmp len = 1, max len = 1

// 4, find 3 >> find 2 >> find 1 >> find 0 >> not found! The `1` is the min number of this group.
// find 2, tmp len = 2, max len = 2
// find 3, tmp len = 3, max len = 3
// find 4, tmp len = 4, max len = 4
// find 5, not found, break

// ... finally we will got the length of max consecutive numbers.