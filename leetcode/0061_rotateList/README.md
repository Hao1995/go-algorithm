# 61. Rotate List

## Intuition
// find length and lastNode, O(n)
// calculate rotate steps, k%length, O(1)
// find parentNode and targetNode, O(n)
// remove next node of parentNode, O(1)
// connect lastNode to head node, O(1)
// return targetNode, O(1)

## Approach
// head=[1,2,3,4,5], k=1
// len=5, k=1, rotateSteps=1%5=1, lastNode=5
// seekCount=5-1=4
// round 1, parentNode=1, targetNode=2
// round 2, parentNode=2, targetNode=3
// round 3, parentNode=3, targetNode=4
// round 4, parentNode=4, targetNode=5
// parentNode 4.Next=nil
// lastNoe 5.Next=head=1
// ans=[5,1,2,3,4]

// head=[1,2,3,4,5], k=2
// len=5, k=2, rotateSteps=2%5=2, lastNode=5
// seekCount=5-2=3
// parentNode=1, targetNode=2, round 1
// parentNode=2, targetNode=3, round 2
// parentNode=3, targetNode=4, round 3
// parentNode 3.Next=nil
// lastNode 5.Next=head=1
// ans=[4,5,1,2,3]

// head=[1,2,3,4,5], k=8
// len=5, k=8, rotateSteps=8%5=3, lastNode=5
// ...
// same process
// ...