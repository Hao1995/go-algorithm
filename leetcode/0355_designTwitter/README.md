# 355. Design Twitter

## Intuition
m: the number of users
n: the number of posts

PostTweet >> O(1)
GetNewsFeed >> Worst case: O(n) // copy all posts
Follow >> O(1)
Unfollow >> O(1)

Total space complexity: O(n+m^m) // each user follow all other users

## Approach
