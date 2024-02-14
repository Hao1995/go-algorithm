# 134. Gas Station

## Intuition
1. Use `Binary Search` to find the `start` gas station >> O(logn)
2. Check gas tank every travel >> O(n)

Total >> O(nlogn)

But we can't find the `start` gas station when using `Binary Search`.
Because the input `gas` can't be order.

## Approach
Calculate the gas tank every travel, if it's burned out, find next possible gas stations.

// gas=[1,2,3,4,5], cost=[1,2,3,5,4], isSet=false
// 0, tank=1-1=0, total=0+0=0 >> ok >> isSet==false >> ans=0, isSet=true
// 1, tank=0+2-2=0, total=0+0=0 >> ok >> isSet==true
// 2, tank=0+3-3=0, total=0+0=0 >> ok >> isSet==true
// 3, tank=0+4-5=-1, total=0-1=-1 >> failed >> reset, tank=0, isSet=false
// 4, tank=5-4=1, total=-1+1=0 >> ok >> isSet==false >> ans=4, isSet=true
// total >= 0 >> return ans