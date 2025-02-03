# 238. Product Except Self

# Intuition

# Approach
Ref: https://youtu.be/bNvIQI2wAjk?si=0xQPRXxTFX0gtXNX

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
prefix >> 1
postfix >> 1
ouput >> n
1+1+n >> O(n)