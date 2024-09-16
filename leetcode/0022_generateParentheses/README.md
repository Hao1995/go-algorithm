# 22. Generate Parentheses

## Intuition
I am going to try any possible combination of parentheses.
So it must be `backtracking`.

Backtracking:
To find all or some solutions to a computational problem.
Builds a solution incrementally, discarding partial solutions that fail to satisfy the problem constraints.

## Approach
// n=3
// "", open=0, close=0, combs=[]
// add "(" >> "(", open=1, close=0, combs=[]
//  add "(" >> "((", open=2, close=0, combs=[]
//      add "(" >> "(((", open=3, close=0, combs=[]
//          add ")" >> "((()", open=3, close=1, combs=[]
//              add ")" >> "((())", open=3, close=2, combs=[]
//                  add ")" >> "((()))", open=3, close=3, combs=["((()))"]
//      add ")" >> "(()", open=2, close=1, combs=["((()))"]
//          add "(" >> "(()(", open=3, close=1, combs=["((()))"]
//              add ")" >> "(()()", open=3, close=2, combs=["((()))"]
//                  add ")" >> "(()())", open=3, close=3, combs=["((()))", "(()())"]
//          add ")" >> "(())", open=2, close=2, combs=["((()))", "(()())"]
//              add "(" >> "(())(", open=3, close=2, combs=["((()))", "(()())"]
//                  add ")" >> "(())()", open=3, close=3, combs=["((()))", "(()())", "(())()"]
// ...
