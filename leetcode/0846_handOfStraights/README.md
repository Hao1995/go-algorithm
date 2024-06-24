# 846. Hand of Straights

## Intuition
Hash Map, Min Heap

## Approach

### V1: Check the consecutive group exist.
This question aims to rearrange the cards into new groups with consecutive numbers. A hash map is created to easily identify consecutive numbers and check their quantities.

The key issue is how to find the minimum card? A minimum heap can solve this problem. Once the quantity of a card decreases to zero, it is removed from the minimum heap. If the card with lowest value in the heap is not the card that has just run out, that indicates that the remaining cards could not be consisted to consecutive numbers.

Time complexity is O(n+mk+mlogm)
Space Complexity is O(n).

### V2: Check the consecutive group exist.
// For example
// hand=[1,2,3,2,3,4], groupSize=3, count=[1:1,2:2,3:2,4:1], lastChecked=-1, opened=0

// Get min one 1=1:
// open a new group, push 1-0=1 to start, start=[1]
// lastChecked=1, opened=1

// Get min one 2=2:
// open a new group, push 2-1=1 to start, start=[1,1]
// lastChecked=2, opened=2

// Get min one 3=2:
// push 2-2=0 to start, start=[1,1,0]
// lastChecked=3, opened=2
// it match current opened groups.
// Enough cards, close first groups=1, start=[1,0], opened=opened-1=1

// Get min one 4=1:
// push 1-1=0 to start, start=[1,0,0]
// lastChecked=4, opened=1
// it match current opened groups.
// Enough cards, close first groups=1, start=[0,0], opened=opened-1=0

// return opened == 0;

Time complexity: O(n+mlogm)
Space complexity: O(m+k)