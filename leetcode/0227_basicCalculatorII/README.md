# 227. Basic Calculator II

## Intuition
Because we need to calculate multiplication and division first, then handle addition and subtraction afterward.
There are totally two loop.
But it exceeds the time limit.

## Approach
[Cracking FAANG](https://youtu.be/W3Rg4HVSZ9k?si=44FnxfVq5MJ4QMvO)

Use a single loop to process each result, keeping track of the previous number.
This allow us to recalculate the result when encountering a multiplication or division operator.
