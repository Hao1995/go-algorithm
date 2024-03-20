# 22. Generate Parentheses

## Intuition
I am going to try any possible combination of parentheses.
So it must be `backtracking`.

Backtracking:
To find all or some solutions to a computational problem.
Builds a solution incrementally, discarding partial solutions that fail to satisfy the problem constraints.

## Approach
// n=3, store=[3,3]
// put a '(', store=[2,3], tmp=['(']
// if store '(' >= ')'
//   put a '(', store=[1,3], tmp=['((']
//     if store '(' >= ')'
//       put a '(', store=[0,3], tmp=['(((']
//         if store '(' > 0 and >= ')'
//           break
//         if store ')' > 0 and >= '('
//           put a ')', store=[0,2], tmp=['((()']
// if store '(' >= ')'
//   put a ')', store=[2,2], tmp=['()']
//     ...
