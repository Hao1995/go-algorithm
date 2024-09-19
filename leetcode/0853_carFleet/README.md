# 853. Car Fleet

## Intuition
I have no idea...

## Approach
Refer solution: https://youtu.be/Pr6T-3yB9RM?si=a19ePkArh6nk8tL3

// EX1:
// p=[10,8,0,5,3], s=[2,4,1,1,3]
// sorted=[(10,2),(8,4),(5,1),(3,3),(0,1)]
// car1, 10 + 2T = 12, 2T = 2, T = 1
// car2, 8 + 4T = 12, 4T = 4, T = 1
// car4, 5 + 1T = 12, T = 7
// car5, 3 + 3T = 12, T = 3
// car3, 0 + 1T = 12, T = 12

// EX2: (float)
// p=[6,8], s=[3,2], target=10
// sorted=[(8,3),(6,2)]
// car2, 8+3T=10, T=2/3=1.5
// car1, 6+2T=10, T=2