# 155 Min Stack
https://leetcode.com/problems/min-stack/

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
You must implement a solution with O(1) time complexity for each function.

## Intuition
Because the stack is LIFO, we just need to remember the minimum value each time we push.

## Approach
Create a data structure to store the current value and the minimum value since the first push:
```
type Node struct {
    Min int
    Val int
    Next *Node
}
```

So the code will run as follows:
// push 5 >> node(5)=(5,5,nil) >> heap->node(5)
// push 3 >> node(3)=(3,3,node(5)) >> heap->node(3)->node(5)
// push 4 >> node(4)=(4,3,node(3)) >> heap->node(4)->node(3)->node(5)
// push 2 >> node(2)=(2,2,node(4)) >> heap->node(2)->node(4)->node(3)->node(5)
// getMin >> heap->node(2)->min >> 2
// pop >> heap.node(2).next=node(4) >> heap->node(4)->node(3)->node(5)
// getMin >> heap->node(4)->min >> 3