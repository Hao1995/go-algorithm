# 20. Valid Parentheses

## Intuition
Push open brackets to stack then compare the current closed bracket each time.

## Approach
EX. 1
// "([])"
// '(', stack=['(']
// '[', stack=['(','[']
// ']', get '[' from stack, stack=['('], '[' and ']' are a pair, continue
// ')', get '(' from stack, stack=[], '(' and ')' are a pair, continue

EX. 2
// "([)]"
// '(', stack=['(']
// '[', stack=['(','[']
// ')', get '[' from stack, '[' and ')' are not a pair, return false