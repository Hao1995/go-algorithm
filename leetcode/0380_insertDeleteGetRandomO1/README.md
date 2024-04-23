# 380. Insert Delete Get Random O(1)

## Intuitive
My intuitive is using a hash Map, because the question require O(1) for insert, remove, getRandom.
The question is "how to get the random value?".
We can't use the build-in function `rand.Intn` with the hash map, because the hash map without specific range.
There must have another list for calculating the random value, then connecting the hash map to the list.