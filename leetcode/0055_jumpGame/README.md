# 55. Jump Game

## Intuition
I don't understand the question.
I originally thought that the question said that the first idx number was my starting position, so I couldn't understand the operating logic of the subsequent example cases at all.

In fact, the first idx number also is the maximum jump length.
So if the input is `nums=[2,3,1,1,4]`, then I am going to start at index `0` and the maximum jump length is `2`.

## Approach
// nums=[2,3,1,1,4]
// idx=0, 2~1
// idx=2, 1~1
// idx=3, 1~1,
// idx=4, true


We should record the `visited` array because we may jump into the same location that we met before.
// nums=[2,2,1,2,1,1,2,2,2,2,2,4]
// idx=0, 2~1
// [idx=0, jump=2] idx=2, 1~1
// [idx=0, jump=2] idx=3, 2~1, visited[3]=true
// [idx=0, jump=2] idx=5, 0, false
// [idx=0, jump=1] idx=1, 2~1
// [idx=0, jump=1] idx=3, visited[3]==true, pass

So the solution is `Dynamic Programming`
