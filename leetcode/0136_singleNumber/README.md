# 136. Single Number

## Intuition
All solutions in my minds:
1. sort >> O(nlogn), O(1)
2. hash map >> O(n), O(n)
3. greedy >> O(n^2), O(1)

But The question requires time complexity: O(n) and space complexity: O(1).

The question mentions "every elements appears twice except for one". So I thought maybe I can try "XOR".

## Approach

XOR output false when both of input are same:
```
true XOR true = false
true XOR false = true
false XOR false = false
false XOR true = true
```

So if the input are number:
1. 4 XOR 4 = 0
2. 4 XOR 4 XOR 2 = 2
3. 4 XOR 4 XOR 2 XOR 2 = 0
4. 4 XOR 4 XOR 2 XOR 2 XOR 3 = 3
