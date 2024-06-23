# 846. Hand of Straights

## Intuition
Hash Map, Min Heap

## Approach
This question aims to rearrange the cards into new groups with consecutive numbers. A hash map is created to easily identify consecutive numbers and check their quantities.

The key issue is how to find the minimum card? A minimum heap can solve this problem. Once the quantity of a card decreases to zero, it is removed from the minimum heap. If the card with lowest value in the heap is not the card that has just run out, that indicates that the remaining cards could not be consisted to consecutive numbers.

Time Complexity is O(nlogn).
Space Complexity is O(n).