# 115. Distinct Subsequences

## Intuition
We need to list all possibility, but some combinations could be duplicated during the exploring, we must use Dynamic Programming to cache the possibility.

## Approach
Solution: https://youtu.be/-RDzMJ33nx8?si=eDAAoMcWGvpUbHRp

// s="rabbbit"(7), t="rabbit"(6)
// t[0]='r', s[0]='r', equal >> return dp[1][1]+dp[0][1] = 3+0 = 3, dp[0][0]=3
//      t[1]='a', s[1]='a', equal >> return dp[2][2]+dp[2][3] = 2+1 = 3, dp[1][1]=3
//          t[2]='b', s[2]='b', equal >> return dp[3][3]+dp[3][4]=2, dp[2][2]=2
//              t[3]='b', s[3]='b', equal >> return 1, dp[3][3]=1
//                  t[4]='i', s[4]='b', not same >> return 1, dp[4][4]=1
//                      t[4]='i', s[5]='i', equal >> return 1+0=1, dp[4][5]=1
//                          t[5]='t', s[6]='t', equal >> return 1, dp[5][6]=1
//                          t[5]='t', s[7] is out >> return 0
//                      t[4]='i', s[6]='t', not same >> return 0, dp[4][6]=0
//                          t[4]='t', s[7] is out >> return 0
//              t[3]='b', s[4]='b', equal >> return 1, dp[3][4]=1
//                  t[4]='i', s[5]='i', dp[4][5]=1 >> return 1
//                  t[3]='b', s[5]='i', not same >> return 0, dp[3][5]=0
//                      t[3]='b', s[6]='t', not same >> return 0, dp[3][6]=0
//          t[2]='b', s[3]='b', equal >> return 1, dp[2][3]=1
//              t[3]='b', s[4]='b', dp[3][4]=1 >> return dp[3][4]=1
//      t[0]='a', s[1]='a', not same >> return dp[0][1]=0
// t[0]='r', s[1]='a', not same >> return 0, dp[0][1]=0
//      t[0]='r', s[2]='b', not same >> return 0, dp[0][2]=0
//          t[0]='r', s[3]='b', not same >> return 0, dp[0][3]=0
//              t[0]='r', s[4]='b', not same >> return 0, dp[0][4]=0
//                  t[0]='r', s[5]='i', not same >> return 0, dp[0][5]=0
//                      t[0]='r', s[6]='t', not same >> return 0, dp[0][6]=0