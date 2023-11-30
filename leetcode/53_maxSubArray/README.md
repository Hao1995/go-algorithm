# Intuition
I've tried to implement well known algorithm on it.
First one is `Binary Search`, we will define a answer then check if any subarray has the same sum. But its time complexity would reach O(n^2), because I have to start from first element to last element (n) with the different sizes (n).

The second one is that I try to sum up each element from the beginning of the nums array, then I will remember the maximum ans, the current ans. If the current num is greater than current ans, replace it, and if the current ans is greater than the maximum ans, replact it.

# Approach
Normal Case:
// [-2,1,-3,4,-1,2,1,-5,4]
// start from -2, maxAns=min_int32, currAns=0
// -2: currAns=-2, maxAns=-2
// 1: currAns(-2) < 0 && num(1) > currAns(-2) >> currAns=1, maxAns=1
// -3: currAns=1-3=-2, maxAns=1
// 4, currAns(-2) < 0 && num(4) > currAns(-2) >> currAns=4, maxAns=4
// -1, currAns=3, maxAns=4
// 2, currAns=5, maxAns=5
// 1, currAns=6, maxAns=6
// -5, currAns=0, maxAns=6
// 4, currAns=4, maxAns=6
// ans=6

Negative Case:
// [-1,-2,-3]
// start from -1, maxAns=min_int32, currAns=0
// -1: currAns=-1, maxAns=-1
// -2: currAns(-1) < 0 && num(-2) > currAns(-1) >> currAns+=-2=-3, maxAns=-1
// -3: currAns(-3) < 0 && num(-3) > currAns(-3) >> currAns+=-3=-6, maxAns=-1
// ans=-1

Asc Negative Case:
// [-3,-2,-1]
// start from -3, maxAns=min_int32, currAns=0
// -3, currAns=-3, maxAns=-3
// -2: currAns(-3) < 0 && num(-2) > currAns(-3) >> currAns=-2, maxAns=-2
// -1: currAns(-2) < 0 && num(-1) > currAns(-2) >> currAns=-1, maxAns=-1


# Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```
func maxSubArray(nums []int) int {
    var ans int = -1 << 31
    var localSum int = 0
    for _, num := range nums {
        localSum += num
        localSum = max(localSum, num)
        ans = max(ans, localSum)
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```