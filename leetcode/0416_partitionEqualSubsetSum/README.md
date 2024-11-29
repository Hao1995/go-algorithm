# 416. Partition Equal Subset Sum

## Intuition
First, sort the array in descending order.
Declare a variable named `sum` and initialize it with the value at the rightmost position. Start subtracting values from the second position from the right.
If the result is greater than 0, continue subtracting; otherwise, add the value starting from the left side.

However, this approach does not always work since the `sum` might be composed of more than one number in certain cases:
```
[3,3,3,4,5] = [3,3,3] + [4,5]
```

## Approach
### Approach 1 - DP (T:O(n\*m), S:O(n\*m))
Check if current subset is valid at current target.
(The vertical axis represents different targets)
(The horizontal axis represents the input `nums`)

   0, 1, 2, 3, 4
 [ 0, 1, 5,11, 5]
 0 y  y  y  y  y 
 1    y  y  y  y 
 2               
 3               
 4               
 5       y  y  y 
 6       y  y  y 
 7               
 8               
 9               
10             y 
11          y  y 

### Approach 2 - Hash Map (T:O(n^2), S:O(n^2))
The hash map would store any combination without `target` limit.