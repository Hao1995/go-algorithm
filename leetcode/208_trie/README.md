# Intuition
Using a tree to remember the child characters.

# Approach
We also need an end flag to check if it's a word or just a prefix.
```
// There is the tree when inserted = ["apple", "app", "banana"]
     a.      b
  p  c  u.   a
  p  t  n.   n
 l * *  t.   a
 e.     *    n
 *           a
             *
```

# Complexity
- Time complexity:
Insert: O(n)
Search: O(n)
Prefix: O(n)

- Space complexity:?