# 50. Pow(x,n)

## Intuition
Due to the value of `n` being between -2^31 and 2^31, if we use a foreach loop to multiply `ans`, it will exceed the time limitation.
To decrease the time complexity to O(logn), we can divide the power by 2 each round.

## Approach
// x=2, n=10 >> 2^10
// 10/2=5, (2*2)^5=4^5, x=4, n=5 >> return 1024
// 5/2=2..1, (4*4)^2*4=16^2*4, x=16, n=2, remainder=4 >> return 256 * 4
// 2/2=1, 1 < 2 >> return 16*16

// x=2, n=-2 >> 2^(-2)
// ifNegative=true, n=2, x=2 >> return 1 / 4 = 0.25
// 2/2=1, 1< 2 >> return 2*2 = 4