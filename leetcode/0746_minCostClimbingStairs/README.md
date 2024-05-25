# 746. Min Cost Climbing Stairs

## Intuition
DP, because the problem can be breakdown to small problem.

## Approach
// cost=[10,15,20], dp=[x,x,x]
// dp[0]=curr=10
// dp[1]=curr=15
// dp[2]=min(10+20, 15+20)=min(30, 35)=30
// return min(dp[len(dp)-1], dp[len(dp)-2])=min(30,15)

// cost=[1,100,1,1,1,100,1,1,100,1], dp=[x,x,x,x,x,x,x,x,x,x]
// dp[0]=1
// dp[1]=100
// dp[2]=min(1+1, 100+1)=2
// dp[3]=min(100+1, 2+1)=3
// dp[4]=min(2+1, 3+1)=3
// dp[5]=min(3+100, 3+100)=103
// dp[6]=min(3+1, 103+1)=4
// dp[7]=min(103+1, 4+1)=5
// dp[8]=min(4+100, 5+100)=104
// dp[9]=min(5+1, 104+1)=6
// return min(dp[len(dp)-1], dp[len(dp)-2])=min(6,104)