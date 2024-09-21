# 875. Koko Eating Bananas

## Intuition
No.

## Approach
Ref: https://youtu.be/U2SozAs9RzA?si=dsTQQ8mobQVM30In

// piles=[30,11,23,4,20], h=5
// k is from 1/h ~ max(p)=30/h
// binary search from 1/h ~ 30/h to reduce find the minimum eating rate.
// * invalid case use "l=midd+1", valid case uses "r=midd"