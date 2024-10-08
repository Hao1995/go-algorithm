# 121. Best Time to Buy and Sell Stock

## Intuition
Iterate each price, then remember the minimum price and update the maximum profit.

## Approach
// [7,1,5,3,6,4]
// Init minPrice=prices[0]=7, ans=0
// 1, 1 < 7, update min >> minPrice=1
// 5, profit=5-1=4, ans=max(4,0)=4
// 3, profit=3-1=2, ans=max(4,2)=4
// 6, profit=6-1=5, ans=max(4,5)=5
// 4, profit=4-1=3, ans=max(5,3)=5

// If the prices is decreasing order. We will return default ans `0` directly.