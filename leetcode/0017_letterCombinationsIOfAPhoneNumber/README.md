# 17. Letter Combinations of a Phone Number

## Intuition
The question requires us to list all possible permutations and combinations.
My first thought is using backtracking.

## Approach
// cands=[2,3]
// cands=[2,3], cand=[a,b,c]

// >> tmpList=[a], cands=[3]
// cands=[3], cand=[d,e,f]
// >> tmpList=[a,d], cands=[]
// >> tmpList=[a,e], cands=[]
// >> tmpList=[a,f], cands=[]

// >> tmpList=[b], cands=[3]
// cands=[3], cand=[d,e,f]
// >> tmpList=[b,d], cands=[]
// >> tmpList=[b,e], cands=[]
// >> tmpList=[b,f], cands=[]

// ...
