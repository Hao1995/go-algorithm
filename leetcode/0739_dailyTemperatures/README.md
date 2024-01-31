# 739. Daily Temperatures

## Intuition
Iterate through each element from the beginning.
If the current element is greater than the previous element, I will calculate the 'number of days'.

However, it doesn't work as expected. For example: [30,35,34,33,36].
I will iterate through the elements until '36'.
Then, I will go back to calculate the previous answers.
But I don't know when to stop!
For example: I have to stop when I encounter '35', but it's difficult to determine.

So, if I have a stack to record the previous elements that I need to calculate, it's much easier.


## Approach

// [73,74,75,71,69,72,76,73]
// stack=[]
// 0=73, emptyStack=next, ans=[0,0,0,0,0,0,0,0], stack=[0]
// 1=74, 74>73,emptyStack=next, ans=[1,0,0,0,0,0,0,0], stack=[1]
// 2=75, 75>74,emptyStack=next, ans=[1,1,0,0,0,0,0,0], stack=[2]
// 3=71, 71<=74, ans=[1,1,0,0,0,0,0,0], stack=[2,3]
// 4=69, 69<=71, ans=[1,1,0,0,0,0,0,0], stack=[2,3,4]
// 5=72, 72>69=1,72>71=2,72<=75=next, ans=[1,1,0,2,1,0,0,0], stack=[2,5]
// 6=76, 76>72=1,76>75=4,emptyStack=next, ans=[1,1,4,2,1,1,0,0], stack=[6]
// 7=73, 73<=76=next, ans=[1,1,0,1,1,0,0,0], stack=[6]