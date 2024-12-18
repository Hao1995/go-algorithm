# 150. Eval RPN

## Intuition
Stack
T: O(n)
S: O(n)

## Approach
// ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
// 10, stack=[10]
// 6, stack=[10, 6]
// 9, stack=[10, 6, 9]
// 3, stack=[10, 6, 9, 3]
// +, pop from stack, 9 + 3 = 12, push into the stack=[10, 6, 12]
// -11, stack=[10, 6, 12, -11]
// *, pop from stack, 12 * -11 = -132, push into the stack=[10, 6, -132]
// /, pop from stack, 6 / -132 = 0, push into the stack=[10, 0]
// *, pop from stack, 10 * 0 = 0, push into the stack=[0]
// 17, stack=[0, 17]
// +, pop from stack, 0 + 17 = 17, push into the stack=[17]
// 5, stack=[17, 5]
// +, pop from stack, 17 + 5 = 22, push into the stack=[22]