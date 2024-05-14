# 179. Largest Number

## Intuition
Sort.
The problem is how to sort?

There are so many cases that hard to sort.

Case1: It bigger when the next char greater than previous chars.
34567 > 345
34531 < 345
// nums = [3,33,345,34567,3453,34531,34543,5,9] >> [5,9,34567,345,34543,3,33]

Case2: When the middle number is smaller, the size is determined by the first one. 
300 > 30000
300 > 3004
3004 > 3000
3003 > 300
300 > 3002
// nums = [3,33,300,30000,3004,3003,3002,33335,5,9] >> [5,9,33335,3,33,3004,3003,300,3002,30000]

## Approach
Use sorting the same way.
Directly use the concatenated strings to decide the order.

34567 v.s. 345
return "34567345" > "34534567"

3003 v.s. 300
return "3003300" > "3003003"