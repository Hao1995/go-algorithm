# 1. Two Sum

## Intuition
Sort the nums and use two pointer to find the corresponding index. (Approach2)

## Approach
// Approach1
// New array stors number and index at the same time >> O(n)
// Sort array in order to use two pointers >> O(nlogn)
// Two pointer find the corresponding index >> O(n)
// >> total: O(nlogn)

// Approach2
// For loop to store remaining data into the map, return the corresponding index when the remaining number exist >> O(n)
// total: O(n)