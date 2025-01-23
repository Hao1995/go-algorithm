# 1899. Merge Triplets to Form Target Triplet

## Intuition
Backtracking.
But this solution is too complicated, we're going to try endless possibilities because we can do `max` operation any number of times.


## Approach
Ref: NeetCode https://youtu.be/kShkQLQZ9K4?si=Z1kEYAK6cbA2InsU

// triplets = [[2,5,3],[1,8,4],[1,7,5]], target = [2,7,5]
// target[0]=2, it can be found at triplets[0][0]
// target[1]=7, it can be found at triplets[2][1]
// target[2]=5, it can be found at triplets[2][2]
// triplets[1][1] is greater than target[1], so skip it.