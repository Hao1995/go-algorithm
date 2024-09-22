# Intuition
Dynamic Programming.

# Approach
But I figured out I don't need to remember previous result.
So I only keep the current sum and final answer.

// nums=[-2,1,-3,4,-1,2,1,-5,4], ans=-Max, sum=0
// -2, sum=-2, ans=-2
// 1, sum=max(-2+1=-1,1)=1, ans=2
// -3, sum=max(1-3=-2,-3)=-2, ans=2
// 4, sum=max(-2+4=2,4)=4, ans=4
// -1, sum=max(4-1=3,-1)=3, ans=4
// 2, sum=max(3+2=5,2)=5, ans=5
// 1, sum=max(5+1=6,1)=6, ans=6
// -5, sum=max(6-5=1,-5)=1, ans=6
// 4, sum=max(1+4=5,4)=5, ans=6