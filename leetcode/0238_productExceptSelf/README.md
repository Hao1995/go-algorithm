# Intuition
My first thought is calculate the total product from the all elements, the the second loop will divide by current num, then the result is the answer.
But the `division` is not allowed in this problem. So I referred to [_V3N0M_](https://leetcode.com/problems/product-of-array-except-self/solutions/1694007/golang-easy-implementation-of-solution/)'s solution.

# Approach
First for loop, calculate the prefix product then stroe at current index.
The second loop, calculate the postfix product, then storing to current index.

EX:
// num=[1,2,3,4], ans=[1,1,1,1]
// Fisrt loop: prefix
// prefix=1
// store previous product to 0, [1,1,1,1], prefix *= curr num = 1*1 = 1
// store previous product to 1, [1,1,1,1], prefix *= curr num = 1*2 = 2
// store previous product to 2, [1,1,2,1], prefix *= curr num = 2*3 = 6
// store previous product to 3, [1,1,2,6], prefix *= curr num = 6*4 = 24
// Fisrt loop: postfix
// store previous product to 3, [1,1,2,6], postfix *= curr num = 1*4 = 4
// store previous product to 2, [1,1,8,6], postfix *= curr num = 4*3 = 12
// store previous product to 1, [1,12,8,6], postfix *= curr num = 12*2 = 24
// store previous product to 0, [24,12,8,6], postfix *= curr num = 24*1 = 24

# Complexity
- Time complexity:
n+n=2n >> O(n)

- Space complexity:
ans >> n
prefix >> 1
postfix >> 1
n+1+1=O(n)

# Code
```
func productExceptSelf(nums []int) []int {
    var ans []int = make([]int, len(nums))
    for i := 0; i < len(ans); i++ {
        ans[i] = 1
    }

    var prefix int = 1
    for i := 0; i < len(ans); i++ {
        ans[i] *= prefix
        prefix *= nums[i]
    }

    var postfix int = 1
    for i := len(ans) - 1; i >= 0; i-- {
        ans[i] *= postfix
        postfix *= nums[i]
    }

    return ans
}
```