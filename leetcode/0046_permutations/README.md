# 46. Permutations

## Intuition
1. Convert to hash map in order to try the remaining nums.
(If we are using an array, the first permutation must start by [1,2], but when we want to try [1,3], we are hard to know  we still have a [2] not yet be considered.)
2. Try different num
3. Recursive backtracking function to try different possibility within the hash map.
4. Restore the num back to the hash map.
5. Repeat 2~4.

## Approach
// nums=[1,2,3]
// convert it to a hash map=[1:1, 2:1, 3:1]
// iterate the hash map
// 1, permutation=[1], map=[2:1, 3:1]
//      2, permutation=[1, 2], map=[3:1]
//          3, permutation=[1, 2, 3], map=[]
//              append permutation to ans=[[1, 2, 3]]
//          restore 3, map=[3:1]
//      restore 2, map=[2:1, 3:1]
//      3, permutation=[1, 3], map=[2:1]
//          2, permutation=[1, 3, 2], map=[]
//              append permutation to ans=[[1, 2, 3], [1, 3, 2]]
// ...

// T: O(n!)
// S: O(n)