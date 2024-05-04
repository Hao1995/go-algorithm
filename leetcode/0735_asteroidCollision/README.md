# 735. Asteroid Collision

## Intuition
Stack.
Push valid asteroid into the stack.
If the next asteroid will cause collision, trigger the collision process.

## Approach
Time Complexity: O(n)

// [5,5,5,5,-10,6]
// 5, sign=right, [5]
// 5, sign=right, [5,5]
// 5, sign=right, [5,5,5]
// 5, sign=right, [5,5,5,5]
// -10, sign=left, state is from right to left, changed, start coliising
// for
//      pop asteroid from the stash, aestriod=5
//      if |-10| == 5
//          break
//      elseif |-10|>5
//          win, continue
//      else // less
//          lose, push 5 back, then brake.