# 437. Path Sum III

## Intuition
The question wants to know the number of paths where the sum of the numbers along with the path equals to `targetSum`.
It gives us a tree, so my first thought is DFS.
They want to calculate the sum no matter start from which node. Just only match the go downwards condition. It must be a `prefixSum` problem.

# Approach
Create a hash map to record the count of the sum.
If current sum exists in the map, it's a valid answer.
At the end of the DFS function, I will remove the current sum from the map to avoid record the sum of the different paths.

// question. targetSum=8, sum=0, map[0:1]
// num=10, sum+=10=10, diff=10-8=2, diff not exist, map[0:1,10:1]
//      left, num=5, sum+=5=15, diff=15-8=7, diff not exist, map[0:1,10:1,15:1]
//          left, num=3, sum+=3=18, diff=18-8=10, diff is exist=1, ans+=1=1, map[0:1,10:1,15:1,18:1]
//              left, num=3, sum+=3=21, diff=21-8=13, diff not exist, map[0:1,10:1,15:1,18:1,21:1]
//                  left, return
//                  right, return
//                  remove itself from map=[0:1,10:1,15:1,18:1]
//              right, num=-2, sum+=-2=16, diff=16-8=8, diff not exist, map[0:1,10:1,15:1,16:1,18:1]
//                  left, return
//                  right, return
//                  remove itself from map=[0:1,10:1,15:1,18:1]
//              remove itself from map=[0:1,10:1,15:1]
//          right, num=2, sum+=2=17, diff=17-8=9, diff not exist, map[0:1,10:1,15:1,17:1]
//              left, return
//              rightm, num=1, sum+=1=18, diff=18-8=10, diff is exit=1, ans+=1=2, map[0:1,10:1,15:1,17:1,18:1]
//              ...