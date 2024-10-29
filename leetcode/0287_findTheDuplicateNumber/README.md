# 287. Find the Duplicate Number

## Intuition
1. Sort `nums`, and then for each to find if the number is not match the corresponding index.
  >> Can't modify the array `nums`
2. Create an array first. For loop `nums` and store number to corresponding index. Then we can find the one already exist.
  >> Only allow constant extra space.
3. Hash map. Iterate the entire `nums` to check if the number already exist.
  >> Only allow constant extra space.

## Approach - Binary Search
Ref: https://www.youtube.com/live/86co28GuZ5U?si=89fwUvb71vD9FH4S&t=1185

The above Youtube video introduce 2 approaches:
1. Indexing Sort: Swap numbers to correct index >> The video is created at 2021. But the question was updated now.(2023)
2. Binary Search: Using BinarySearch to guess the repeated number, each guess we will for loop `nums` to check if it's possible one.

// [1,3,4,2,2]
// low=0, high=n-1=4 (must one repeated number, so maximum must be n-1)
// midd=2, count=3, 3>2, lower, low=0, high=2
// midd=1, count=1, 1==1, higher, low=midd+1=1+1=2, high=2
// break, return 2

// [3,1,3,4,2]
// low=0, high=4
// midd=2, count=2, 2==2, higher, low=midd+1=2+1=3, high=4
// midd=3, count=4, 4>3, lower, low=3, high=3,
// break, return 3

## Approach - Floyd's Cycle Detection.
Ref: https://youtu.be/wjYnzkAhcNk?si=PRyegN5hCs9oC8VV

Imagine the `nums` is a linked-list, each number means the next index.

According to the Floyd's algorithm, we can find the intersection by the slow and fast pointers.
(The detail explanation should check the video.)
