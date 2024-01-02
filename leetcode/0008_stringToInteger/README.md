# 8. String to Integer (atoi)

## Intuition
Check boundaries.
* check the second parts number were skipped. EX: "42 dog 13" >> 42
* check if the string is starting beyond '-', '+', ' ', return 0. EX: "dog 42" >> 0

## Approach

