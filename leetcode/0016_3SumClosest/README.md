# 16. 3Sum Closest

## Intuition
Brute force, DFS to traversal each combinations, but it's time complexity is O(n^3).

## Approach
[Nikhi Lohia](https://youtu.be/uSpJQa6MRZ8?si=G8TdHRc6zEuSb2LW)

Sort the nums with ascending order, that in order to find a bigger or smaller sum result.

// nums=[-1,2,1,-4], target=1
// sort=[-4,-1,0,1,2,5]
// fix -4, left=-1, right=5, sum=0
//          if 0 == target { return sum }


// nums=[1,1,1,0], target=-100
// sort=[0,1,1,1], sum=2
// fixed 0
//      left=1, right=3, sum=2
//          2 > target ? right-- >> right=2
//      left=1, right=2, sum=2
//          2 > target ? right-- >> right = 1
//      break because left == right.
// return 2